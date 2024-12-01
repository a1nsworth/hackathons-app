package hackathon

import (
	"net/http"
	"strconv"

	"hackathons-app/internal/models"
	"hackathons-app/internal/services"

	"github.com/gin-gonic/gin"
)

// HackathonHandler содержит сервис для работы с хакатонами
type HackathonHandler struct {
	service services.HackathonService
}

// NewHackathonHandler создает новый хэндлер для работы с хакатонами
func NewHackathonHandler(service services.HackathonService) *HackathonHandler {
	return &HackathonHandler{service: service}
}

// GetAll - хэндлер для получения всех хакатонов
// @Summary Get all hackathons
// @Description Получение списка всех хакатонов
// @Tags hackathons
// @Accept json
// @Produce json
// @Success 200 {array} models.Hackathon
// @Router /hackathons [get]
func (h *HackathonHandler) GetAll(c *gin.Context) {
	hackathons, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hackathons)
}

// GetById - хэндлер для получения хакатона по ID
// @Summary Get hackathon by ID
// @Description Получение хакатона по ID
// @Tags hackathons
// @Accept json
// @Produce json
// @Param id path int true "Hackathon ID"
// @Success 200 {object} models.Hackathon
// @Router /hackathons/{id} [get]
func (h *HackathonHandler) GetById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	hackathon, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}
	c.JSON(http.StatusOK, hackathon)
}

// Create - хэндлер для создания нового хакатона
// @Summary Create a new hackathon
// @Description Создание нового хакатона
// @Tags hackathons
// @Accept json
// @Produce json
// @Param hackathon body models.Hackathon true "Hackathon data"
// @Success 201 {object} models.Hackathon
// @Router /hackathons [post]
func (h *HackathonHandler) Create(c *gin.Context) {
	var hackathon models.Hackathon
	if err := c.ShouldBindJSON(&hackathon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.service.Create(&hackathon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, hackathon)
}

// Update - хэндлер для обновления хакатона
// @Summary Update hackathon data
// @Description Обновление данных хакатона
// @Tags hackathons
// @Accept json
// @Produce json
// @Param id path int true "Hackathon ID"
// @Param hackathon body models.Hackathon true "Hackathon data"
// @Success 200 {object} models.Hackathon
// @Router /hackathons/{id} [put]
func (h *HackathonHandler) Update(c *gin.Context) {
	var hackathon models.Hackathon
	if err := c.ShouldBindJSON(&hackathon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.service.Update(&hackathon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hackathon)
}

// DeleteById - хэндлер для удаления хакатона по ID
// @Summary Delete hackathon by ID
// @Description Удаление хакатона по ID
// @Tags hackathons
// @Accept json
// @Produce json
// @Param id path int true "Hackathon ID"
// @Router /hackathons/{id} [delete]
func (h *HackathonHandler) DeleteById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.service.DeleteById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
