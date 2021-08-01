package repo

import (
	"errors"
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

type Tag struct {
	Model
	User
	Name string `json:"name" gorm:"unique"`
	Code uint64 `json:"code" gorm:"uniqueIndex"`
	View
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

type Resource interface {
	SelectFields(ref map[string]interface{}) []string

	SetId(id uint64)

	SetUserId(userId string)

	CheckAndSetViewPath()

	ValidateFields(fieldNames ...string) error
}

// ******************************
// Post

// SelectFields selects fields of Post which appear in ref.
// Returns these fields' name as string.
func (p *Post) SelectFields(ref map[string]interface{}) []string {
	return _s(ref, modelFields["Post"])
}

// SetId set Post ID.
func (p *Post) SetId(id uint64) {
	p.ID = id
}

// SetUserId set Post UserId.
func (p *Post) SetUserId(userId string) {
	p.UserId = userId
}

// CheckAndSetViewPath set a unique random ViewPath if ViewPath of Post is ""
func (p *Post) CheckAndSetViewPath() {
	if p.ViewPath == "" {
		p.ViewPath = util.GenerateUUID()
	}
}

// ValidateFields validate Post fields using predefined functions.
// If no field is specified, call all validation functions,
// otherwise call validation functions specified by fieldNames.
func (p *Post) ValidateFields(fieldNames ...string) error {
	validators := map[string]func() error{
		"Title": func() error {
			if p.Title == "" {
				return errors.New("empty Title")
			}
			return nil
		},

		"ViewPath": func() error {
			if p.ViewPath == "" {
				return errors.New("empty ViewPath")
			}
			return nil
		},

		"Tags": func() error {
			existingTags, err := allTags(p.UserId)
			if err != nil {
				return err
			}
			if existingTags|p.Tags != existingTags {
				return errors.New("undefined tags")
			}
			return nil
		},
	}
	return _runValidator(validators, fieldNames...)
}

// ******************************
// Tag

// SelectFields selects fields of Tag which appear in ref.
// Returns these fields' name as string.
func (t *Tag) SelectFields(ref map[string]interface{}) []string {
	return _s(ref, modelFields["Tag"])
}

// SetId set Tag ID.
func (t *Tag) SetId(id uint64) {
	t.ID = id
}

// SetUserId set Tag UserId.
func (t *Tag) SetUserId(userId string) {
	t.UserId = userId
}

// CheckAndSetViewPath set a unique random ViewPath if ViewPath of Tag is ""
func (t *Tag) CheckAndSetViewPath() {
	if t.ViewPath == "" {
		t.ViewPath = util.GenerateUUID()
	}
}

// ValidateFields validate Tag fields using predefined functions.
// If no field is specified, call all validation functions,
// otherwise call validation functions specified by fieldNames.
// *Note*: This function may modify input Tag model.
func (t *Tag) ValidateFields(fieldNames ...string) error {
	validators := map[string]func() error{
		"Name": func() error {
			if t.Name == "" {
				return errors.New("empty Tag Name")
			}
			return nil
		},

		"ViewPath": func() error {
			if t.ViewPath == "" {
				return errors.New("empty Tag ViewPath")
			}
			return nil
		},

		"Code": func() error {
			existingTags, err := allTags(t.UserId)
			if err != nil {
				return err
			}
			validBits := ^existingTags
			k, available := 0, false
			for validBits != 0 {
				if validBits&1 == 1 {
					available = true
					break
				} else {
					k++
					validBits >>= 1
				}
			}
			if !available {
				return errors.New("exceed maximum tags number")
			}
			validTag := uint64(1) << k
			// ** here modify the input Tag model
			t.Code = validTag
			return nil
		},
	}
	return _runValidator(validators, fieldNames...)
}

// ******************************

func _s(args map[string]interface{}, ref map[string]string) []string {
	var ret []string
	for fieldName := range args {
		if camelFieldName, present := ref[fieldName]; present {
			ret = append(ret, camelFieldName)
		}
	}
	return ret
}

func _runValidator(validators map[string]func() error, fieldNames ...string) error {
	if len(fieldNames) == 0 {
		for _, validator := range validators {
			if err := validator(); err != nil {
				return err
			}
		}
	} else {
		for _, fieldName := range fieldNames {
			if _, present := validators[fieldName]; present {
				if err := validators[fieldName](); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func allTags(userId string) (uint64, error) {
	if userId == "" {
		return 0, errors.New("empty userId")
	}

	var tags []Tag
	var cond = Condition{
		Query: "user_id = ?",
		Args:  []interface{}{userId},
	}
	if _, err := FindAll(&tags, cond); err != nil {
		return 0, err
	}

	if len(tags) == 0 {
		// no defined tags
		return 0, nil
	}

	tc := tags[0].Code
	for i := 1; i < len(tags); i++ {
		tc |= tags[i].Code
	}
	return tc, nil
}
