package dao

import (
	"encoding/json"
	"errors"
	"os"
	"pili-apiserver/pkg/model"
)

type Dao interface {
	Init()
	Get(id int) (model.Character, error)
	List() []model.Character
	Save(character model.Character)
	Update(character model.Character) error
	Delete(id int) error
}

type SliceDao struct {
	characters []model.Character
}

func (sd *SliceDao) Init() {
	jsonFile, err := os.Open("docs/Homework/布袋戲資料.json")
	if err != nil {
		panic("Read characters data fails")
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	decoder.Decode(&sd.characters)
}

func (sd *SliceDao) Get(id int) (model.Character, error) {
	for _, character := range sd.characters {
		if character.ID == id {
			return character, nil
		}
	}
	return model.Character{}, errors.New("No such Character")
}

func (sd *SliceDao) List() []model.Character {
	return sd.characters
}

func (sd *SliceDao) Save(character model.Character) {
	sd.characters = append(sd.characters, character)
}

func (sd *SliceDao) Update(character model.Character) error {
	for index, oldCharacter := range sd.characters {
		if oldCharacter.ID == character.ID {
			sd.characters[index] = character
			return nil
		}
	}
	return errors.New("No such Character")
}

func (sd *SliceDao) Delete(id int) error {
	for index, character := range sd.characters {
		if character.ID == id {
			sd.characters = append(sd.characters[:index], sd.characters[index+1:]...)
		}
	}
	return errors.New("No such Character")
}
