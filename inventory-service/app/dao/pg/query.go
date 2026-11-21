package pg

const (
	CATEGORY_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS categories(
		id SMALLSERIAL PRIMARY KEY,
        name VARCHAR(255) UNIQUE NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)`
	CATEGORY_INSERT_QUERY    = "INSERT INTO categories(name,created_at) VALUES($1,$2) RETURNING id"
	CATEGORY_GET_BY_ID_QUERY = "SELECT id,name,created_at FROM categories WHERE id=$1"
	CATEGORY_GET_QUERY       = "SELECT id,name,created_at FROM categories"

	BRAND_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS brands(
		id SMALLSERIAL PRIMARY KEY,
		name VARCHAR(255) UNIQUE NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)`
	BRAND_GET_QUERY = "SELECT id,name,created_at FROM brands ORDER BY name"

	SUB_CATEGORY_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS sub_categories(
		id SMALLSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
		category_id SMALLINT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		FOREIGN KEY (category_id) REFERENCES categories(id),
		UNIQUE (name,category_id)
	)`
	SUB_CATEGORY_INSERT_QUERY    = "INSERT INTO sub_categories(name,category_id,created_at) VALUES($1,$2,$3) RETURNING id"
	SUB_CATEGORY_GET_BY_ID_QUERY = "SELECT id,name,category_id,created_at FROM sub_categories WHERE id=$1"
	SUB_CATEGORY_GET_QUERY       = "SELECT id,name,category_id,created_at FROM sub_categories"

	PRODUCT_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS products(
		id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
		description TEXT,
		slug VARCHAR(255),
		brand_id SMALLINT NOT NULL,
		category_id SMALLINT NOT NULL,
		sub_category_id SMALLINT NOT NULL,
		image_url VARCHAR(255),
		status BIGINT NOT NULL DEFAULT 0 CHECK (status IN (0, 1, 2)),
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		deleted_at TIMESTAMPTZ,
		CONSTRAINT products_brand_id_fkey FOREIGN KEY (brand_id) REFERENCES brands(id),
		FOREIGN KEY (category_id) REFERENCES categories(id),
		FOREIGN KEY (sub_category_id) REFERENCES sub_categories(id),
		UNIQUE (name,category_id,sub_category_id)
	)`
	PRODUCT_BRAND_CONSTRAINT_QUERY = `DO $$
	BEGIN
		IF NOT EXISTS (
			SELECT 1
			FROM pg_constraint
			WHERE conname = 'products_brand_id_fkey'
			  AND conrelid = 'products'::regclass
		) THEN
			ALTER TABLE products
				ADD CONSTRAINT products_brand_id_fkey
				FOREIGN KEY (brand_id) REFERENCES brands(id) NOT VALID;
		END IF;
	END
	$$`
	PRODUCT_STATUS_CONSTRAINT_QUERY = `DO $$
	BEGIN
		IF NOT EXISTS (
			SELECT 1
			FROM pg_constraint
			WHERE conname = 'products_status_check'
			  AND conrelid = 'products'::regclass
		) THEN
			ALTER TABLE products
				ADD CONSTRAINT products_status_check
				CHECK (status IS NOT NULL AND status IN (0, 1, 2)) NOT VALID;
		END IF;
	END
	$$`
	PRODUCT_INSERT_QUERY    = "INSERT INTO products(name,description,slug,brand_id,category_id,sub_category_id,image_url,status,created_at,updated_at,deleted_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id"
	PRODUCT_UPDATE_QUERY    = "UPDATE products SET name=$1,description=$2,slug=$3,brand_id=$4,category_id=$5,sub_category_id=$6,image_url=$7,status=$8,updated_at=$9,deleted_at=$10 WHERE id=$11"
	PRODUCT_GET_BY_ID_QUERY = "SELECT id,name,description,slug,brand_id,category_id,sub_category_id,image_url,status,created_at,updated_at,deleted_at FROM products WHERE id=$1"
	PRODUCT_GET_QUERY       = "SELECT id,name,description,slug,brand_id,category_id,sub_category_id,image_url,status,created_at,updated_at,deleted_at FROM products"

	KEYWORD_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS keywords(
		id BIGSERIAL PRIMARY KEY,
		name VARCHAR(255) UNIQUE NOT NULL
	)`
	KEYWORD_INSERT_QUERY    = "INSERT INTO keywords(name) VALUES($1) RETURNING id"
	KEYWORD_GET_BY_ID_QUERY = "SELECT id,name FROM keywords WHERE id=$1"
	KEYWORD_GET_QUERY       = "SELECT id,name FROM keywords"

	PRODUCT_KEYWORD_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS product_keywords(
		product_id BIGINT NOT NULL,
		keyword_id BIGINT NOT NULL,
		PRIMARY KEY (product_id,keyword_id),
		FOREIGN KEY (product_id) REFERENCES products(id),
		FOREIGN KEY (keyword_id) REFERENCES keywords(id)
	)`
	PRODUCT_KEYWORD_INSERT_QUERY            = "INSERT INTO product_keywords(product_id,keyword_id) VALUES($1,$2)"
	PRODUCT_KEYWORD_GET_BY_PRODUCT_ID_QUERY = "SELECT product_id,keyword_id FROM product_keywords WHERE product_id=$1"
	PRODUCT_KEYWORD_GET_BY_KEYWORD_ID_QUERY = "SELECT product_id,keyword_id FROM product_keywords WHERE keyword_id=$1"
	PRODUCT_KEYWORD_DELETE_QUERY            = "DELETE FROM product_keywords WHERE product_id=$1 AND keyword_id=$2"

	INVENTORY_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS inventory(
		product_id BIGINT NOT NULL PRIMARY KEY,
        sku INT,
		mrp INT,
		purchase_price INT,
		sale_price INT,
        quantity INT NOT NULL DEFAULT 0 CHECK (quantity >= 0),
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        FOREIGN KEY (product_id) REFERENCES products(id)
	)`
	INVENTORY_INSERT_QUERY    = "INSERT INTO inventory(product_id,sku,mrp,purchase_price,sale_price,quantity,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8)"
	INVENTORY_UPDATE_QUERY    = "UPDATE inventory SET sku=$1,mrp=$2,purchase_price=$3,sale_price=$4,quantity=$5,updated_at=$6 WHERE product_id=$7"
	INVENTORY_GET_BY_ID_QUERY = "SELECT product_id,sku,mrp,purchase_price,sale_price,quantity,created_at,updated_at FROM inventory WHERE product_id=$1"
	INVENTORY_GET_QUERY       = "SELECT product_id,sku,mrp,purchase_price,sale_price,quantity,created_at,updated_at FROM inventory"
)
