package models

import (
	"gocoon/core/models/entity"
)

var ModelList = []interface{}{
	&entity.User{},
	&entity.Todo{},
}

var ModelMap = map[string]interface{}{
	"User": &entity.User{},
	"Todo": &entity.Todo{},
}

type Model interface{}

func GetModelByName(name string) (interface{}, bool) {
	model, ok := ModelMap[name]
	return model, ok
}
