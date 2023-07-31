package handler

import (
	"encoding/json"
	"fmt"

	"tokens-api/pkg/behaviour"

	"github.com/gin-gonic/gin"
)

type crudHandler struct {
	entity string
	reader behaviour.ReaderDomain
	writer behaviour.WriterDomain
}

func NewCrudHandler(entity string, reader behaviour.ReaderDomain, writer behaviour.WriterDomain) *crudHandler {
	return &crudHandler{entity, reader, writer}
}

func (h *crudHandler) Setup(r *gin.Engine) {
	baseEntityURI := fmt.Sprintf("/%s", h.entity)
	r.GET(baseEntityURI, h.getAll)
	r.GET(baseEntityURI+"/:id", h.getByID)
	r.POST(baseEntityURI, h.create)
	r.PUT(baseEntityURI+"/:id", h.update)
	r.DELETE(baseEntityURI+"/:id", h.delete)
}

// getAll returns all the entities based on paramters
func (h *crudHandler) getAll(c *gin.Context) {
	q := c.Request.URL.Query()
	entities, err := h.reader.All(c, convertQuery((q)))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, entities)
}

// getByID returns the entity with the given id
func (h *crudHandler) getByID(c *gin.Context) {
	id := c.Param("id")
	entity, err := h.reader.ByID(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, entity)
}

// create creates a new entity
func (h *crudHandler) create(c *gin.Context) {
	var entity any
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	entity, err := json.Marshal(&entity)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	entity, err = h.writer.Create(c, entity)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, entity)
}

// update updates the entity with the given id
func (h *crudHandler) update(c *gin.Context) {
	id := c.Param("id")
	var entity interface{}
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	entity, err := h.writer.Update(c, id, entity)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, entity)
}

// delete deletes the entity with the given id
func (h *crudHandler) delete(c *gin.Context) {
	id := c.Param("id")
	err := h.writer.Delete(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, gin.H{})
}

// convert URL.Values to map[string]interface{}
func convertQuery(q map[string][]string) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range q {
		m[k] = v[0]
	}
	return m
}
