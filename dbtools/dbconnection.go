package dbtools

import (
	"fidelis.com/communication/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DBInitializer struct {
	DB *gorm.DB
}

func Connect(dataSoruceName string) (*DBInitializer, error) {
	conn, err := gorm.Open(mysql.Open(dataSoruceName), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	conn.AutoMigrate(&model.Student{})

	return &DBInitializer{
		DB: conn,
	}, err
}

func (dbInit DBInitializer) GetStudentById(id int64) model.Student {
	newStudent := model.Student{}
	dbInit.DB.First(&newStudent, id)

	return newStudent
}

func (dbInit DBInitializer) GetStudentsByName(name string) []model.Student {
	var students []model.Student

	dbInit.DB.Where("name = ?", name).Find(&students)

	return students
}
