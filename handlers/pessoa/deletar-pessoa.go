/*
Package pessoa fornece manipuladores de rota para deletar uma pessoa existente.
*/

package pessoa

import (
	"github.com/VictorCCole/go-restohub-api/models"
	"net/http"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/gin-gonic/gin"
)

// DeletarPessoa deleta uma pessoa existente pelo ID.
func DeletarPessoa(c *gin.Context) {
	// Obter o ID da pessoa da URL
	id := c.Param("id")

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Deletar a pessoa pelo ID
	if err := db.Delete(&models.Pessoa{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar pessoa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pessoa deletada com sucesso"})
}
