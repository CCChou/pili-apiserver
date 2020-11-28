package handler

import (
	"net/http"
	"pili-apiserver/pkg/dao"
	"pili-apiserver/pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	dao dao.Dao
}

func (h *Handler) Init() {
	d := dao.MapDao{}
	d.Init()
	h.dao = &d
}

func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	role, err := h.dao.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "")
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *Handler) List(c *gin.Context) {
	roles := h.dao.List()
	c.JSON(http.StatusOK, roles)
}

func (h *Handler) Save(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	h.dao.Save(role)
	c.JSON(http.StatusOK, role)
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	if _, err := h.dao.Get(id); err != nil {
		c.JSON(http.StatusNotFound, "")
		return
	}
	role.ID = id
	h.dao.Update(role)
	c.JSON(http.StatusOK, role)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	h.dao.Delete(id)
	c.JSON(http.StatusNoContent, "")
}
