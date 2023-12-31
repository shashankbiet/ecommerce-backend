package mysql

const (
	CATEGORY_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS Categories(
		Id SMALLINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
        Name VARCHAR(255) UNIQUE NOT NULL,
		CreatedAt DATETIME
	)`
	CATEGORY_INSERT_QUERY    = "INSERT INTO Categories(Name,CreatedAt) VALUES(?,?)"
	CATEGORY_GET_BY_ID_QUERY = "SELECT Id,Name,CreatedAt FROM Categories WHERE Id=?"
	CATEGORY_GET_QUERY       = "SELECT Id,Name,CreatedAt FROM Categories"

	SUB_CATEGORY_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS SubCategories(
		Id SMALLINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
        Name VARCHAR(255) NOT NULL,
		Category VARCHAR(255),
		CreatedAt DATETIME,
		FOREIGN KEY (Category) REFERENCES Categories(Name),
		UNIQUE KEY UniqueId (Name,Category)
	)`
	SUB_CATEGORY_INSERT_QUERY    = "INSERT INTO SubCategories(Name,Category,CreatedAt) VALUES(?,?,?)"
	SUB_CATEGORY_GET_BY_ID_QUERY = "SELECT Id,Name,Category,CreatedAt FROM SubCategories WHERE Id=?"
	SUB_CATEGORY_GET_QUERY       = "SELECT Id,Name,Category,CreatedAt FROM SubCategories"

	PRODUCT_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS Products(
		Id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
        Name        VARCHAR(255) NOT NULL,
		Description TEXT,
		Brand       VARCHAR(255) NOT NULL,
		Category    VARCHAR(255),
		SubCategory VARCHAR(255),
		ImageId     VARCHAR(255),
		Weight      FLOAT,
		CreatedAt   DATETIME,
		UpdatedAt   DATETIME,
		FOREIGN KEY (Category) REFERENCES Categories(Name),
		FOREIGN KEY (SubCategory) REFERENCES SubCategories(Name),
		UNIQUE KEY  UniqueId (Name,Category,SubCategory)
	)`
	PRODUCT_INSERT_QUERY    = "INSERT INTO Products(Name,Description,Brand,Category,SubCategory,ImageId,Weight,CreatedAt,UpdatedAt) VALUES(?,?,?,?,?,?,?,?,?)"
	PRODUCT_UPDATE_QUERY    = "UPDATE Products SET Name=?,Description=?,Brand=?,Category=?,SubCategory=?,ImageId=?,Weight=?,UpdatedAt=? WHERE Id=?"
	PRODUCT_GET_BY_ID_QUERY = "SELECT Id,Name,Description,Brand,Category,SubCategory,ImageId,Weight,CreatedAt,UpdatedAt FROM Products WHERE Id=?"
	PRODUCT_GET_QUERY       = "SELECT Id,Name,Description,Brand,Category,SubCategory,ImageId,Weight,CreatedAt,UpdatedAt FROM Products"

	INVENTORY_CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS Inventory(
		ProductId BIGINT NOT NULL PRIMARY KEY,
        SKU       		INT,
		PurchasePrice   INT,
		SalePrice 		INT,
		CreatedAt 		DATETIME,
		UpdatedAt 		DATETIME,
		FOREIGN KEY (ProductId) REFERENCES Products(Id)
	)`
	INVENTORY_INSERT_QUERY    = "INSERT INTO Inventory(ProductId,SKU,PurchasePrice,SalePrice,CreatedAt,UpdatedAt) VALUES(?,?,?,?,?,?)"
	INVENTORY_UPDATE_QUERY    = "UPDATE Inventory SET SKU=?,PurchasePrice=?,SalePrice=?,UpdatedAt=? WHERE ProductId=?"
	INVENTORY_GET_BY_ID_QUERY = "SELECT ProductId,SKU,PurchasePrice,SalePrice,CreatedAt,UpdatedAt FROM Inventory WHERE ProductId=?"
	INVENTORY_GET_QUERY       = "SELECT ProductId,SKU,PurchasePrice,SalePrice,CreatedAt,UpdatedAt FROM Inventory"
)
