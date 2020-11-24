package handler

import (
	"pili-apiserver/pkg/dao"
	"pili-apiserver/pkg/model"
)

type Handler struct {
	dao dao.Dao
}

func (h *Handler) Init() {
	sliceDao := dao.SliceDao{}
	sliceDao.Init()
	h.dao = &sliceDao
}

func (h *Handler) Get(id int) (model.Character, error) {
	return h.dao.Get(id)
}

func (h *Handler) List() []model.Character {
	return h.dao.List()
}

func (h *Handler) Save(characeter model.Character) {
	h.dao.Save(characeter)
}

func (h *Handler) Update(character model.Character) error {
	return h.dao.Update(character)
}

func (h *Handler) Delete(id int) error {
	return h.dao.Delete(id)
}
