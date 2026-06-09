package utils

import (
	"errors"

	"github.com/thedevsaddam/govalidator"
	"gorm.io/gorm"
)

// Struct
type ValidateCollection struct{}

// Construct
func Validation() *ValidateCollection {
	validate := &ValidateCollection{}
	return validate
}

// Is Unique
func (v ValidateCollection) Unique(db *gorm.DB, model interface{}, field string, value string, ignores ...interface{}) bool {
	query := db.Where(field+"=?", value)

	switch len(ignores) {
	case 1:
		query.Where("id!=?", ignores[0])
	case 2:
		if ignore_field, ok := ignores[1].(string); ok {
			query.Where(ignore_field+"!=?", ignores[0])
		}
	}
	result := query.First(&model)
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// Is Exist
func (v ValidateCollection) Exist(db *gorm.DB, model interface{}, field string, value interface{}) bool {
	result := db.Where(field+"=?", value).First(&model)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// Validate
func (v *ValidateCollection) Validate(data interface{}, rules govalidator.MapData) interface{} {
	opts := govalidator.Options{
		Data:  data,
		Rules: rules,
	}

	validator := govalidator.New(opts)
	errors := validator.ValidateStruct()

	formatedError := make(map[string]string)
	for key, err := range errors {
		formatedError[key] = err[0] // assuming single value for simplicity
	}

	if len(formatedError) > 0 {
		return formatedError
	}

	return nil
}
