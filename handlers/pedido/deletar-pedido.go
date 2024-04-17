/*
Package pedido fornece manipuladores de rota para deletar um pedido existente.
*/

package pedido

import (
	"github.com/VictorCCole/go-restohub-api/models"
	"net/http"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/gin-gonic/gin"
)

// DeletarPedido deleta um pedido existente pelo ID.
func DeletarPedido(c *gin.Context) {
	// Obter o ID do pedido da URL
	id := c.Param("id")

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Deletar o pedido pelo ID
	if err := db.Delete(&models.Pedido{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar pedido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pedido deletado com sucesso"})
}
