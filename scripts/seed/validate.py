#!/usr/bin/env python3

import json
import re
from pathlib import Path


SEED_DIR = Path(__file__).resolve().parent
SQL_PATH = SEED_DIR / "postgres.sql"
ELASTICSEARCH_PATH = SEED_DIR / "elasticsearch.ndjson"

PRODUCT_FIELDS = {
    "id",
    "name",
    "description",
    "slug",
    "brand",
    "category",
    "subCategory",
    "imageUrl",
    "status",
    "createdAt",
    "updatedAt",
    "deletedAt",
}
SAMPLE_IMAGE_URL = "https://placehold.co/600x400/png?text=Demo+Product"


def split_tuple(line: str) -> list[object]:
    values: list[object] = []
    current: list[str] = []
    quoted = False
    index = 1

    while index < len(line) - 1:
        char = line[index]
        if char == "'":
            if quoted and index + 1 < len(line) - 1 and line[index + 1] == "'":
                current.append("'")
                index += 2
                continue
            quoted = not quoted
        elif char == "," and not quoted:
            values.append("".join(current).strip())
            current = []
        else:
            current.append(char)
        index += 1
    values.append("".join(current).strip())
    return values


def insert_rows(sql: str, table: str) -> list[list[object]]:
    match = re.search(
        rf"INSERT INTO {table}\([^)]+\) VALUES\n(?P<rows>.*?)\nON CONFLICT",
        sql,
        re.DOTALL,
    )
    if not match:
        raise AssertionError(f"{table} seed block not found")

    rows = []
    for line in match.group("rows").splitlines():
        line = line.strip().removesuffix(",")
        if line:
            rows.append(split_tuple(line))
    return rows


def inventory_rows(sql: str) -> list[list[object]]:
    match = re.search(
        r"WITH inventory_seed\([^)]+\) AS \(\nVALUES\n(?P<rows>.*?)\n\)\n"
        r"INSERT INTO inventory",
        sql,
        re.DOTALL,
    )
    if not match:
        raise AssertionError("Inventory seed block not found")
    return [
        split_tuple(line.strip().removesuffix(","))
        for line in match.group("rows").splitlines()
        if line.strip()
    ]


def read_elasticsearch_products() -> list[dict]:
    lines = ELASTICSEARCH_PATH.read_text().splitlines()
    if len(lines) % 2:
        raise AssertionError("Elasticsearch seed must contain action/document line pairs")

    products = []
    for offset in range(0, len(lines), 2):
        action = json.loads(lines[offset])
        product = json.loads(lines[offset + 1])
        expected_id = str(product["id"])
        if action.get("index", {}).get("_id") != expected_id:
            raise AssertionError(
                f"action ID {action.get('index', {}).get('_id')} does not match product {expected_id}"
            )
        products.append(product)
    return products


def main() -> None:
    sql = SQL_PATH.read_text()
    products = read_elasticsearch_products()
    sql_categories = insert_rows(sql, "categories")
    sql_subcategories = insert_rows(sql, "sub_categories")
    sql_brands = insert_rows(sql, "brands")
    sql_products = insert_rows(sql, "products")
    sql_keywords = insert_rows(sql, "keywords")
    sql_inventory = inventory_rows(sql)

    if len(products) != 105 or len(sql_products) != 105:
        raise AssertionError("expected 105 products in PostgreSQL and Elasticsearch")
    if (
        len(sql_categories) != 15
        or len(sql_subcategories) != 60
        or len(sql_brands) != 15
    ):
        raise AssertionError("expected 15 brands, categories, and 60 subcategories")
    if len(sql_inventory) != len(products):
        raise AssertionError("each product must have one inventory seed row")
    expected_keyword_names = {
        str(row[1]) for row in sql_categories + sql_subcategories
    }
    keyword_names = [str(row[0]) for row in sql_keywords]
    if len(keyword_names) != 71 or set(keyword_names) != expected_keyword_names:
        raise AssertionError("keywords must include each unique category and subcategory")
    if len(keyword_names) != len(set(keyword_names)):
        raise AssertionError("keyword seed contains duplicate names")

    category_ids = {int(row[0]) for row in sql_categories}
    brand_ids = {int(row[0]) for row in sql_brands}
    subcategory_categories = {int(row[0]): int(row[2]) for row in sql_subcategories}
    inventory_product_ids = {int(row[0]) for row in sql_inventory}

    if [int(row[0]) for row in sql_inventory] != list(range(1, 106)):
        raise AssertionError("inventory product IDs must cover 1 through 105")
    if any(row[2] == "NULL" for row in sql_inventory):
        raise AssertionError("inventory MRP must be populated")

    for product, row in zip(products, sql_products, strict=True):
        if set(product) != PRODUCT_FIELDS:
            raise AssertionError(f"product {product.get('id')} does not match the current model")
        if not isinstance(product["brand"], int) or not isinstance(product["category"], int):
            raise AssertionError(f"product {product['id']} contains legacy string references")
        if not isinstance(product["subCategory"], int):
            raise AssertionError(f"product {product['id']} contains a legacy subcategory")
        if product["status"] != 1:
            raise AssertionError(f"product {product['id']} is not active")
        if product["imageUrl"] != SAMPLE_IMAGE_URL:
            raise AssertionError(f"product {product['id']} does not use the sample image")
        if any(key in product for key in ("imageId", "weight")):
            raise AssertionError(f"product {product['id']} contains legacy fields")

        expected = [
            str(product["id"]),
            product["name"],
            product["description"],
            product["slug"],
            str(product["brand"]),
            str(product["category"]),
            str(product["subCategory"]),
            product["imageUrl"],
            str(product["status"]),
            "NOW()",
            "NOW()",
            "NULL",
        ]
        if row != expected:
            raise AssertionError(f"PostgreSQL product {product['id']} differs from Elasticsearch")
        if product["category"] not in category_ids:
            raise AssertionError(f"product {product['id']} has an unknown category")
        if product["brand"] not in brand_ids:
            raise AssertionError(f"product {product['id']} has an unknown brand")
        if subcategory_categories.get(product["subCategory"]) != product["category"]:
            raise AssertionError(f"product {product['id']} has an invalid subcategory/category pair")

    if set(range(1, 106)) != inventory_product_ids:
        raise AssertionError("inventory rows do not match product IDs")
    if (
        "INSERT INTO product_keywords(product_id, keyword_id)" not in sql
        or "k.name IN (c.name, sc.name)" not in sql
    ):
        raise AssertionError("product keywords must link category and subcategory keywords")
    print(
        "Seed validation passed: 105 model-aligned products, 15 brands, 71 keywords, "
        "and 210 product-keyword associations."
    )


if __name__ == "__main__":
    main()
