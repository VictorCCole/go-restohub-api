/*
Package pessoa fornece manipuladores de rota para criar uma nova pessoa.
*/

package pessoa

import (
	"net/http"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/VictorCCole/go-restohub-api/models"
	"github.com/gin-gonic/gin"
)

// CriarPessoa cria uma nova pessoa.
func CriarPessoa(c *gin.Context) {
	// BindJSON para decodificar o JSON da solicitação em uma struct de pessoa
	var pessoa models.Pessoa
	if err := c.BindJSON(&pessoa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Criar a pessoa no banco de dados
	if err := db.Create(&pessoa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pessoa"})
		return
	}

	c.JSON(http.StatusCreated, pessoa)
}
