package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var Conn = NewMysqlDB()

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func NewMysqlDB() *gorm.DB {
	dbName := "abc"
	// docker mysql 127.0.0.1:3306连接不上可以 设置为如下
	host := "localhost:3306"
	userName := "root"
	password := "123456"
	charset := "utf8"
	//parseTime := true
	prefix := ""
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", userName, password, host, dbName, charset)

	//DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name

	log.Println("dsn", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		DSN:                       dsn,   // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Product{})
	if err != nil {
		panic(err.Error())
	}

	// Migrate the schema
	_ = db.AutoMigrate(&Product{})
	return db
}

func init() {

	// Create
	Conn.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	Conn.First(&product, 1)                 // find product with integer primary key
	Conn.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	Conn.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	Conn.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	Conn.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	Conn.Delete(&product, 1)
	//dbName := "abc"
	//host := "127.0.0.1:3306"
	//userName := "root"
	//password := "123456"
	//charset := "utf8"
	//parseTime := true
	//prefix := ""
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", userName, password, host, dbName, charset, parseTime)
	//
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       dsn,   // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
	//	DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
	//	DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
	//	DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
	//	DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
	//	SkipInitializeWithVersion: false, // smart configure based on used version
	//}), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		TablePrefix:   prefix,
	//		SingularTable: true,
	//	},
	//})
	//if err != nil {
	//	panic("failed to connect database: " + err.Error())
	//}
	//
	//err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Product{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//// Migrate the schema
	//_ = db.AutoMigrate(&Product{})

}
