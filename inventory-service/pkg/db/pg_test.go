package db

import (
	"net/url"
	"testing"

	"inventory-service/pkg/config"
)

func TestPostgresDSN(t *testing.T) {
	dsn := postgresDSN(config.SqlConfig{
		Username: "db user",
		Password: "p@ss:/?#",
		Host:     "localhost",
		Port:     "5432",
		Database: "inventory",
	})

	parsed, err := url.Parse(dsn)
	if err != nil {
		t.Fatalf("parse PostgreSQL DSN: %v", err)
	}

	username := parsed.User.Username()
	password, ok := parsed.User.Password()
	if username != "db user" || !ok || password != "p@ss:/?#" {
		t.Fatalf("unexpected credentials in DSN: user=%q passwordPresent=%t", username, ok)
	}
	if parsed.Hostname() != "localhost" || parsed.Port() != "5432" {
		t.Fatalf("unexpected host in DSN: %s", parsed.Host)
	}
	if parsed.Path != "/inventory" {
		t.Fatalf("unexpected database in DSN: %s", parsed.Path)
	}
	if parsed.Query().Get("sslmode") != "disable" {
		t.Fatalf("unexpected sslmode in DSN: %s", parsed.Query().Get("sslmode"))
	}
}
