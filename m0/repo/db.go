package repo

import (
	"log"

	"gorm.io/gorm"

	"github.com/mirrordust/splendour/m0/repo/sqlite"
)

var DB *gorm.DB

func init() {
	// get database connection
	db, err := sqlite.InitDB()
	if err != nil {
		log.Panicln("InitDB error")
	}
	DB = db

	// migrate
	err = DB.AutoMigrate(&Post{}, &Tag{}, &User{})
	if err != nil {
		log.Panicln("DB AutoMigrate error")
	}
}

// ********** CRUD functions **********

type Condition struct {
	Query         string
	Args          []interface{}
	Orders        []interface{}
	Offset, Limit int
}

func Create(model interface{}) error {
	result := DB.Create(model)
	return result.Error
}

func FindOne(model interface{}, conditions ...interface{}) error {
	result := DB.First(model, conditions...)
	return result.Error
}

func FindAll(models interface{}, condition Condition) error {
	tx := DB.Where(condition.Query, condition.Args...)
	for _, order := range condition.Orders {
		tx = tx.Order(order)
	}

	if condition.Offset != 0 {
		tx = tx.Offset(condition.Offset)
	}
	if condition.Limit != 0 {
		tx = tx.Limit(condition.Limit)
	}
	result := tx.Find(models)
	return result.Error
}

func UpdateOne(model interface{}, fields []string, newValue interface{}) error {
	result := DB.Model(model).Select(fields).Updates(newValue)
	return result.Error
}

func Delete(model interface{}) error {
	result := DB.Delete(model)
	return result.Error
}
