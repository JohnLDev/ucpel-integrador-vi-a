package main

import (
	"github.com/johnldev/integrador-mvc/controller"
	"github.com/johnldev/integrador-mvc/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// Realiza conexão com banco de dados
	conn, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Cria tabelas no banco de dados
	conn.AutoMigrate(&model.Student{}, &model.AditionalInfo{}, &model.MotherInfo{})
	// Inicia servidor HTTP
	controller.StartHttpServer(conn)
}
