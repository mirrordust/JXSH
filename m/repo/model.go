package repo

import (
	"log"

	"github.com/mirrordust/w/m/util"
)

// PUBLISHED status for published post
const PUBLISHED byte = 0b0000_0001 // otherwise `DRAFT`

// {modelName: {camel fieldName or snake field_name: corresponding camel fieldName}}
var modelFields map[string]map[string]string

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
			fields[util.SnakeCase(k)] = k
		}
		return fields
	}

	modelFields = make(map[string]map[string]string, 2)
	modelFields["Post"] = f(Post{})
	modelFields["Tag"] = f(Tag{})
}

// ******************************
// model definition and related methods

type Post struct {
	Model
	User
	Title       string `json:"title" gorm:"unique;not null"`
	Abstract    string `json:"abstract"`
	Content     string `json:"content"`
	ContentType string `json:"contentType"`
	TOC         string `json:"toc"`
	Status      byte   `json:"status"`
	Tags        uint64 `json:"tags"`
	View
}

// SelectFields selects fields of Post which appear in ref.
// Returns these fields' name as string.
func (p *Post) SelectFields(ref map[string]interface{}) []string {
	return _s(ref, modelFields["Post"])
}

type Tag struct {
	Model
	User
	Name string `json:"name" gorm:"unique"`
	Code uint64 `json:"code" gorm:"uniqueIndex"`
	View
}

// SelectFields selects fields of Tag which appear in ref.
// Returns these fields' name as string.
func (t *Tag) SelectFields(ref map[string]interface{}) []string {
	return _s(ref, modelFields["Tag"])
}

type Model struct {
	ID        uint64 `json:"id" gorm:"primarykey"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type User struct {
	UserId string `json:"userId" gorm:"not null;index"`
}

type View struct {
	ViewPath string `json:"viewPath" gorm:"unique;not null"`
}

func _s(args map[string]interface{}, ref map[string]string) []string {
	var ret []string
	for fieldName := range args {
		if camelFieldName, present := ref[fieldName]; present {
			ret = append(ret, camelFieldName)
		}
	}
	return ret
}
