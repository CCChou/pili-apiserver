package dao

import (
	"encoding/json"
	"errors"
	"os"
	"pili-apiserver/pkg/model"
)

type MapDao struct {
	roles map[int]model.Role
}

func (md *MapDao) Init() {
	jsonFile, err := os.Open("docs/Homework/布袋戲資料.json")
	if err != nil {
		panic("Read roles data fails")
	}
	defer jsonFile.Close()

	var roles []model.Role
	decoder := json.NewDecoder(jsonFile)
	decoder.Decode(&roles)

	md.roles = make(map[int]model.Role)
	for _, role := range roles {
		md.roles[role.ID] = role
	}
}

func (md *MapDao) Get(id int) (model.Role, error) {
	role, exists := md.roles[id]
	var err error
	if !exists {
		err = errors.New("No such role")
	}
	return role, err
}

func (md *MapDao) List() []model.Role {
	values := make([]model.Role, 0, len(md.roles))
	for _, value := range md.roles {
		values = append(values, value)
	}
	return values
}

func (md *MapDao) Save(role model.Role) {
	md.roles[role.ID] = role
}

func (md *MapDao) Update(role model.Role) error {
	_, exists := md.roles[role.ID]
	if !exists {
		return errors.New("No such role")
	}
	md.roles[role.ID] = role
	return nil
}

func (md *MapDao) Delete(id int) error {
	_, exists := md.roles[id]
	if !exists {
		return errors.New("No such role")
	}
	delete(md.roles, id)
	return nil
}
