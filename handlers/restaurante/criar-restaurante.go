/*
Package restaurante fornece manipuladores de rota para criar um novo restaurante.
*/

package restaurante

import (
	"net/http"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/VictorCCole/go-restohub-api/models"
	"github.com/gin-gonic/gin"
)

// CriarRestaurante cria um novo restaurante.
func CriarRestaurante(c *gin.Context) {
	// BindJSON para decodificar o JSON da solicitação em uma struct de restaurante
	var restaurante models.Restaurante
	if err := c.BindJSON(&restaurante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Criar o restaurante no banco de dados
	if err := db.Create(&restaurante).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar restaurante"})
		return
	}

	c.JSON(http.StatusCreated, restaurante)
}
