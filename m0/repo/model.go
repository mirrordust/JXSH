package repo

import (
	"log"

	"github.com/mirrordust/splendour/m0/util"
)

const PUBLISHED byte = 0b1 // otherwise `DRAFT`

// structFields {modelName: {camel fieldName or snake field_name: corresponding camel fieldName}}
var structFields map[string]map[string]string

func init() {
	log.Println("init model info...")

	var f = func(v interface{}) map[string]string {
		info, err := util.StructFields(v)
		if err != nil {
			log.Fatalln("get model fields info error")
		}
		fields := make(map[string]string, 2*len(info))
		for k := range info {
			fields[k] = k
			fields[util.ToSnakeCase(k)] = k
		}
		return fields
	}

	structFields = make(map[string]map[string]string, 3)
	structFields["Post"] = f(Post{})
	structFields["Tag"] = f(Tag{})
	structFields["User"] = f(User{})
}

// ********** entity models **********

type Post struct {
	Model
	Title       string `gorm:"unique;not null"`
	Abstract    string
	Content     string
	ContentType string
	TOC         string
	Status      byte
	Tags        uint64
	View
}

type Tag struct {
	Model
	Name string `gorm:"unique"`
	Code uint64 `gorm:"uniqueIndex"`
	View
}

type User struct {
	Model
	Name     string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique"`
}

// ********** auxiliary models **********

type Model struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt int64
	UpdatedAt int64
}

type View struct {
	ViewPath string `gorm:"unique;not null"`
}

// ******************************

//type selector interface {
//	Select(v map[string]interface{}) []string
//}

func (p *Post) Select(v map[string]interface{}) []string {
	return _select(v, structFields["Post"])
}

func (t *Tag) Select(v map[string]interface{}) []string {
	return _select(v, structFields["Tag"])
}

func (u *User) Select(v map[string]interface{}) []string {
	return _select(v, structFields["User"])
}

func _select(args map[string]interface{}, ref map[string]string) []string {
	var ret []string
	for fieldName := range args {
		if camelFieldName, present := ref[fieldName]; present {
			ret = append(ret, camelFieldName)
		}
	}
	return ret
}
