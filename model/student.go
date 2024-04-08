package model

import "gorm.io/gorm"

type Student struct {
	*gorm.Model
	Name string
	Age  int32
	ID   int64 `gorm:"primaryKey"`
}
