/*
Package pessoa fornece manipuladores de rota para atualizar uma pessoa existente.
*/

package pessoa

import (
	"net/http"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/VictorCCole/go-restohub-api/models"
	"github.com/gin-gonic/gin"
)

// AtualizarPessoa atualiza uma pessoa existente pelo ID.
func AtualizarPessoa(c *gin.Context) {
	// Obter o ID da pessoa da URL
	id := c.Param("id")

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Verificar se a pessoa existe
	var pessoa models.Pessoa
	if err := db.First(&pessoa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pessoa não encontrada"})
		return
	}

	// Atualizar a pessoa com os dados fornecidos no corpo da solicitação
	if err := c.BindJSON(&pessoa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Salvar as alterações no banco de dados
	if err := db.Save(&pessoa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar pessoa"})
		return
	}

	c.JSON(http.StatusOK, pessoa)
}
