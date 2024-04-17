/*
Package pessoa fornece manipuladores de rota para listar pessoas existentes.
*/

package pessoa

import (
	"net/http"
	"strconv"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/VictorCCole/go-restohub-api/models"
	"github.com/gin-gonic/gin"
)

// ListarPessoas lista todas as pessoas existentes.
func ListarPessoas(c *gin.Context) {
	// Obter todas as pessoas do banco de dados
	var pessoas []models.Pessoa
	db := database.GetDB()
	if err := db.Find(&pessoas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar pessoas"})
		return
	}

	c.JSON(http.StatusOK, pessoas)
}

// ListarPessoasID lista uma pessoa específica pelo ID.
func ListarPessoasID(c *gin.Context) {
	// Obter o ID da pessoa da URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Buscar a pessoa pelo ID
	var pessoa models.Pessoa
	if err := db.First(&pessoa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pessoa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, pessoa)
}
