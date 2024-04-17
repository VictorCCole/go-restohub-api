/*
Package restaurante fornece manipuladores de rota para atualizar um restaurante existente.
*/

package restaurante

import (
	"net/http"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/VictorCCole/go-restohub-api/models"
	"github.com/gin-gonic/gin"
)

// AtualizarRestaurante atualiza um restaurante existente pelo ID.
func AtualizarRestaurante(c *gin.Context) {
	// Obter o ID do restaurante da URL
	id := c.Param("id")

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Verificar se o restaurante existe
	var restaurante models.Restaurante
	if err := db.First(&restaurante, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Restaurante não encontrado"})
		return
	}

	// Atualizar o restaurante com os dados fornecidos no corpo da solicitação
	if err := c.BindJSON(&restaurante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Salvar as alterações no banco de dados
	if err := db.Save(&restaurante).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar restaurante"})
		return
	}

	c.JSON(http.StatusOK, restaurante)
}
