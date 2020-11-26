package dao

import (
	"encoding/json"
	"errors"
	"os"
	"pili-apiserver/pkg/model"
)

type Dao interface {
	Init()
	Get(id int) (model.Role, error)
	List() []model.Role
	Save(role model.Role)
	Update(role model.Role) error
	Delete(id int) error
}

type SliceDao struct {
	roles []model.Role
}

func (sd *SliceDao) Init() {
	jsonFile, err := os.Open("docs/Homework/布袋戲資料.json")
	if err != nil {
		panic("Read roles data fails")
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	decoder.Decode(&sd.roles)
}

func (sd *SliceDao) Get(id int) (model.Role, error) {
	for _, role := range sd.roles {
		if role.ID == id {
			return role, nil
		}
	}
	return model.Role{}, errors.New("No such role")
}

func (sd *SliceDao) List() []model.Role {
	return sd.roles
}

func (sd *SliceDao) Save(role model.Role) {
	sd.roles = append(sd.roles, role)
}

func (sd *SliceDao) Update(role model.Role) error {
	for index, oldrole := range sd.roles {
		if oldrole.ID == role.ID {
			sd.roles[index] = role
			return nil
		}
	}
	return errors.New("No such role")
}

func (sd *SliceDao) Delete(id int) error {
	for index, role := range sd.roles {
		if role.ID == id {
			sd.roles = append(sd.roles[:index], sd.roles[index+1:]...)
		}
	}
	return errors.New("No such role")
}
