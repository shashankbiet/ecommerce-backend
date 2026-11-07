#!/usr/bin/env python3
"""Check that service env files, docker-compose.yml, secrets, and code agree.

Run from anywhere: python3 scripts/check_env_drift.py
Exits with status 1 when any drift is found. Prints key names and paths only,
never values.
"""

import re
import sys
from pathlib import Path

import yaml

ROOT = Path(__file__).resolve().parents[1]
SERVICES = ["api-gateway", "user-service", "inventory-service", "search-service"]
SOURCE_SUFFIXES = {".js", ".go"}
ENV_LINE = re.compile(r"^([A-Za-z_][A-Za-z0-9_]*)=(.*)$")
TOKEN = re.compile(r"\b[A-Z][A-Z0-9_]+\b")

# Compose literals for the database images must equal the service .env values.
LITERAL_PAIRS = [
    ("mongodb", "MONGO_INITDB_ROOT_USERNAME", "user-service", "MONGODB_ROOT_USERNAME"),
    ("postgres", "POSTGRES_USER", "inventory-service", "SQL_CONFIG_USERNAME"),
    ("postgres", "POSTGRES_DB", "inventory-service", "SQL_CONFIG_DATABASE"),
]


def read_env(path):
    values = {}
    for line in path.read_text().splitlines():
        match = ENV_LINE.match(line)
        if match:
            values[match.group(1)] = match.group(2)
    return values


def compose_env(service):
    env = service.get("environment", {})
    if isinstance(env, list):
        return dict(item.partition("=")[::2] for item in env)
    return {key: str(value) for key, value in env.items()}


def flatten_yaml(node, prefix=""):
    """Map viper-style keys (nested keys joined by _, uppercased) from a YAML tree."""
    names = set()
    if isinstance(node, dict):
        for key, value in node.items():
            names |= flatten_yaml(value, f"{prefix}_{key}" if prefix else str(key))
    else:
        names.add(prefix.upper())
    return names


def consumed_names(service_dir):
    names = set()
    for path in service_dir.rglob("*"):
        if path.suffix in SOURCE_SUFFIXES and "node_modules" not in path.parts and path.is_file():
            names |= set(TOKEN.findall(path.read_text(errors="ignore")))
    default_yaml = service_dir / "conf" / "default.yaml"
    if default_yaml.exists():
        names |= flatten_yaml(yaml.safe_load(default_yaml.read_text()))
    return names


def is_consumed(name, consumed):
    base = name[: -len("_FILE")] if name.endswith("_FILE") else name
    return name in consumed or base in consumed


def main():
    problems = []
    compose = yaml.safe_load((ROOT / "docker-compose.yml").read_text())
    secrets = compose.get("secrets", {})
    for name, spec in secrets.items():
        if not (ROOT / spec["file"]).exists():
            problems.append(f"compose secret '{name}' file missing: {spec['file']}")

    for service_name in SERVICES:
        service_dir = ROOT / service_name
        example = service_dir / ".env.example"
        env_path = service_dir / ".env"
        example_keys = set(read_env(example))
        consumed = consumed_names(service_dir)
        compose_service = compose["services"][service_name]

        for key in sorted(example_keys):
            if not is_consumed(key, consumed):
                problems.append(f"{service_name}/.env.example: '{key}' is not read by the service")

        for key in sorted(compose_env(compose_service)):
            if not is_consumed(key, consumed):
                problems.append(f"docker-compose.yml {service_name}: '{key}' is not read by the service")

        for key, value in compose_env(compose_service).items():
            if value.startswith("/run/secrets/") and value.removeprefix("/run/secrets/") not in secrets:
                problems.append(f"docker-compose.yml {service_name}: '{key}' points to an undeclared secret")

        if env_path.exists():
            env_values = read_env(env_path)
            for key in sorted(set(env_values) ^ example_keys):
                problems.append(f"{service_name}: '{key}' differs between .env and .env.example")
            for key, value in env_values.items():
                if key.endswith("_FILE") and not (service_dir / value).exists():
                    problems.append(f"{service_name}/.env: '{key}' target missing: {value}")
        else:
            print(f"note: {service_name}/.env not found; skipped value checks", file=sys.stderr)

    for db_service, db_key, app_service, app_key in LITERAL_PAIRS:
        env_path = ROOT / app_service / ".env"
        if not env_path.exists():
            continue
        db_value = compose_env(compose["services"][db_service]).get(db_key)
        app_value = read_env(env_path).get(app_key)
        if db_value != app_value:
            problems.append(
                f"{db_service} {db_key} does not match {app_service} {app_key} (values differ)"
            )

    if problems:
        print("Env drift found:")
        for problem in problems:
            print(f"  - {problem}")
        return 1
    print("No env drift found.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
