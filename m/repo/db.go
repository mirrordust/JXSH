package repo

import (
	"errors"
	"log"

	"gorm.io/gorm"

	"github.com/mirrordust/w/m/repo/sqlite"
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
	err = DB.AutoMigrate(&Post{}, &Tag{})
	if err != nil {
		log.Panicln("DB AutoMigrate error")
	}
}

// ******************************
// CRUD functions

func Create(model interface{}) (rows int64, err error) {
	result := DB.Create(model)
	if result.Error != nil {
		return result.RowsAffected, newDBError(result.Error.Error())
	}
	return result.RowsAffected, nil
}

func FindOne(model interface{}, conds ...interface{}) (rows int64, err error) {
	result := DB.First(model, conds...)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.RowsAffected, nil
		}
		return result.RowsAffected, newDBError(result.Error.Error())
	}
	return result.RowsAffected, nil
}

type Condition struct {
	Query         string
	Args          []interface{}
	Orders        []interface{}
	Offset, Limit int
}

func FindAll(models interface{}, c Condition) (rows int64, err error) {
	tx := DB.Where(c.Query, c.Args...)
	for _, order := range c.Orders {
		tx = tx.Order(order)
	}
	if c.Offset != 0 {
		tx = tx.Offset(c.Offset)
	}
	if c.Limit != 0 {
		tx = tx.Limit(c.Limit)
	}
	result := tx.Find(models)
	if result.Error != nil {
		return result.RowsAffected, newDBError(result.Error.Error())
	}
	return result.RowsAffected, nil
}

func UpdateOne(model interface{}, fields []string, newValue interface{}, conds ...interface{}) (rows int64, err error) {
	result := DB.Model(model).Where(conds[0], conds[1:]...).Select(fields).Updates(newValue)
	if result.Error != nil {
		return result.RowsAffected, newDBError(result.Error.Error())
	}
	return result.RowsAffected, nil
}

func Delete(model interface{}, conds ...interface{}) (rows int64, err error) {
	result := DB.Delete(model, conds...)
	if result.Error != nil {
		return result.RowsAffected, newDBError(result.Error.Error())
	}
	return result.RowsAffected, nil
}

// ******************************
// DB Error

type DBError struct {
	Msg string
}

func (e *DBError) Error() string {
	return "DB error: " + e.Msg
}

func newDBError(msg string) error {
	return &DBError{msg}
}
