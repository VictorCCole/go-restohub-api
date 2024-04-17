/*
Package restaurante fornece manipuladores de rota para listar restaurantes existentes.
*/

package restaurante

import (
	"net/http"
	"strconv"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/VictorCCole/go-restohub-api/models"
	"github.com/gin-gonic/gin"
)

// ListarRestaurante lista todos os restaurantes existentes.
func ListarRestaurante(c *gin.Context) {
	// Obter todos os restaurantes do banco de dados
	var restaurantes []models.Restaurante
	db := database.GetDB()
	if err := db.Find(&restaurantes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar restaurantes"})
		return
	}

	c.JSON(http.StatusOK, restaurantes)
}

// ListarRestauranteID lista um restaurante específico pelo ID.
func ListarRestauranteID(c *gin.Context) {
	// Obter o ID do restaurante da URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Buscar o restaurante pelo ID
	var restaurante models.Restaurante
	if err := db.First(&restaurante, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurante não encontrado"})
		return
	}

	c.JSON(http.StatusOK, restaurante)
}
