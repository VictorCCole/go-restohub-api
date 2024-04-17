/*
Package pedido fornece manipuladores de rota para listar pedidos existentes.
*/

package pedido

import (
	"net/http"
	"strconv"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/VictorCCole/go-restohub-api/models"
	"github.com/gin-gonic/gin"
)

// ListarPedido lista todos os pedidos existentes.
func ListarPedido(c *gin.Context) {
	// Obter todos os pedidos do banco de dados
	var pedidos []models.Pedido
	db := database.GetDB()
	if err := db.Find(&pedidos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar pedidos"})
		return
	}

	c.JSON(http.StatusOK, pedidos)
}

// ListarPedidoID lista um pedido específico pelo ID.
func ListarPedidoID(c *gin.Context) {
	// Obter o ID do pedido da URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Buscar o pedido pelo ID
	var pedido models.Pedido
	if err := db.First(&pedido, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido não encontrado"})
		return
	}

	c.JSON(http.StatusOK, pedido)
}
