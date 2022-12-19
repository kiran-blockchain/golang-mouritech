package repository

import (
	"log"
	
	//"os"

	"gorm.io/driver/sqlserver"
  "gorm.io/gorm"
	
)

//DB -> connection to Database
func DB() *gorm.DB {

	dsn := "sqlserver://SA:BrU@u2@@9StepvApp@206.189.129.40?database=demodb"
    db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to Database")
		return nil
	}

	//db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
	return db

}
