package pg

const (
	CATEGORY_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS Categories(
		Id SMALLSERIAL PRIMARY KEY,
        Name VARCHAR(255) UNIQUE NOT NULL,
		CreatedAt TIMESTAMP
	)`
	CATEGORY_INSERT_QUERY    = "INSERT INTO Categories(Name,CreatedAt) VALUES($1,$2) RETURNING Id"
	CATEGORY_GET_BY_ID_QUERY = "SELECT Id,Name,CreatedAt FROM Categories WHERE Id=$1"
	CATEGORY_GET_QUERY       = "SELECT Id,Name,CreatedAt FROM Categories"

	SUB_CATEGORY_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS SubCategories(
		Id SMALLSERIAL PRIMARY KEY,
        Name VARCHAR(255) NOT NULL,
		Category VARCHAR(255),
		CreatedAt TIMESTAMP,
		FOREIGN KEY (Category) REFERENCES Categories(Name),
		UNIQUE (Name,Category)
	)`
	SUB_CATEGORY_INSERT_QUERY    = "INSERT INTO SubCategories(Name,Category,CreatedAt) VALUES($1,$2,$3) RETURNING Id"
	SUB_CATEGORY_GET_BY_ID_QUERY = "SELECT Id,Name,Category,CreatedAt FROM SubCategories WHERE Id=$1"
	SUB_CATEGORY_GET_QUERY       = "SELECT Id,Name,Category,CreatedAt FROM SubCategories"

	PRODUCT_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS Products(
		Id BIGSERIAL PRIMARY KEY,
        Name        VARCHAR(255) NOT NULL,
		Description TEXT,
		Brand       VARCHAR(255) NOT NULL,
		Category    VARCHAR(255),
		SubCategory VARCHAR(255),
		ImageId     VARCHAR(255),
		Weight      FLOAT,
		CreatedAt   TIMESTAMP,
		UpdatedAt   TIMESTAMP,
		FOREIGN KEY (Category) REFERENCES Categories(Name),
		FOREIGN KEY (SubCategory, Category) REFERENCES SubCategories(Name, Category),
		UNIQUE (Name,Category,SubCategory)
	)`
	PRODUCT_INSERT_QUERY    = "INSERT INTO Products(Name,Description,Brand,Category,SubCategory,ImageId,Weight,CreatedAt,UpdatedAt) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING Id"
	PRODUCT_UPDATE_QUERY    = "UPDATE Products SET Name=$1,Description=$2,Brand=$3,Category=$4,SubCategory=$5,ImageId=$6,Weight=$7,UpdatedAt=$8 WHERE Id=$9"
	PRODUCT_GET_BY_ID_QUERY = "SELECT Id,Name,Description,Brand,Category,SubCategory,ImageId,Weight,CreatedAt,UpdatedAt FROM Products WHERE Id=$1"
	PRODUCT_GET_QUERY       = "SELECT Id,Name,Description,Brand,Category,SubCategory,ImageId,Weight,CreatedAt,UpdatedAt FROM Products"

	INVENTORY_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS Inventory(
		ProductId BIGINT NOT NULL PRIMARY KEY,
        SKU       		INT,
		PurchasePrice   INT,
		SalePrice 		INT,
        Quantity		INT NOT NULL DEFAULT 0 CHECK (Quantity >= 0),
        CreatedAt 		TIMESTAMP,
        UpdatedAt 		TIMESTAMP,
        FOREIGN KEY (ProductId) REFERENCES Products(Id)
	)`
	INVENTORY_ADD_QUANTITY_COLUMN_QUERY = "ALTER TABLE Inventory ADD COLUMN IF NOT EXISTS Quantity INT NOT NULL DEFAULT 0 CHECK (Quantity >= 0)"
	INVENTORY_INSERT_QUERY              = "INSERT INTO Inventory(ProductId,SKU,PurchasePrice,SalePrice,Quantity,CreatedAt,UpdatedAt) VALUES($1,$2,$3,$4,$5,$6,$7)"
	INVENTORY_UPDATE_QUERY              = "UPDATE Inventory SET SKU=$1,PurchasePrice=$2,SalePrice=$3,Quantity=$4,UpdatedAt=$5 WHERE ProductId=$6"
	INVENTORY_GET_BY_ID_QUERY           = "SELECT ProductId,SKU,PurchasePrice,SalePrice,Quantity,CreatedAt,UpdatedAt FROM Inventory WHERE ProductId=$1"
	INVENTORY_GET_QUERY                 = "SELECT ProductId,SKU,PurchasePrice,SalePrice,Quantity,CreatedAt,UpdatedAt FROM Inventory"
)
