package main

import (
	"github.com/johnldev/integrador-mvc/controller"
	"github.com/johnldev/integrador-mvc/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	conn, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	conn.AutoMigrate(&model.Student{}, &model.AditionalInfo{}, &model.MotherInfo{})
	controller.StartHttpServer(conn)
}
