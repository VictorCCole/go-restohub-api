/*
Package restaurante fornece manipuladores de rota para deletar um restaurante existente.
*/

package restaurante

import (
	"github.com/VictorCCole/go-restohub-api/models"
	"net/http"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/gin-gonic/gin"
)

// DeletarRestaurante deleta um restaurante existente pelo ID.
func DeletarRestaurante(c *gin.Context) {
	// Obter o ID do restaurante da URL
	id := c.Param("id")

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Deletar o restaurante pelo ID
	if err := db.Delete(&models.Restaurante{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar restaurante"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Restaurante deletado com sucesso"})
}
