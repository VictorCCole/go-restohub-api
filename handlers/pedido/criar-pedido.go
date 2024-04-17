/*
Package pedido fornece manipuladores de rota para criar um novo pedido.
*/

package pedido

import (
	"net/http"

	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/VictorCCole/go-restohub-api/models"
	"github.com/gin-gonic/gin"
)

// CriarPedido cria um novo pedido.
func CriarPedido(c *gin.Context) {
	// BindJSON para decodificar o JSON da solicitação em uma struct de pedido
	var pedido models.Pedido
	if err := c.BindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obter a instância do banco de dados
	db := database.GetDB()

	// Verificar se a pessoa associada ao pedido existe
	var pessoa models.Pessoa
	if err := db.First(&pessoa, pedido.PessoaID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pessoa não encontrada"})
		return
	}

	// Verificar se o restaurante associado ao pedido existe
	var restaurante models.Restaurante
	if err := db.First(&restaurante, pedido.RestauranteID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Restaurante não encontrado"})
		return
	}

	// Criar o pedido no banco de dados
	if err := db.Create(&pedido).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pedido"})
		return
	}

	c.JSON(http.StatusCreated, pedido)
}
