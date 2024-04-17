/*
Package pedido fornece manipuladores de rota para atualizar um pedido existente.
*/

package pedido

import (
	"net/http"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/VictorCCole/go-restohub-api/models"
	"github.com/gin-gonic/gin"
)

// AtualizarPedido atualiza um pedido existente pelo ID.
func AtualizarPedido(c *gin.Context) {
	// Obter o ID do pedido da URL
	id := c.Param("id")

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Buscar o pedido pelo ID
	var pedido models.Pedido
	if err := db.First(&pedido, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido não encontrado"})
		return
	}

	// Atualizar o pedido com os dados fornecidos no corpo da solicitação
	if err := c.BindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Salvar as alterações no banco de dados
	if err := db.Save(&pedido).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar pedido"})
		return
	}

	c.JSON(http.StatusOK, pedido)
}
