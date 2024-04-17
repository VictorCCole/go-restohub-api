package pedido

import (
	"github.com/VictorCCole/go-restohub-api/handlers/pedido"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/pedido/criar", pedido.CriarPedido)
	router.PUT("/pedido/atualizar/:id", pedido.AtualizarPedido)
	router.GET("/pedido/listar/", pedido.ListarPedido)
	router.GET("/pedido/listarid/:id", pedido.ListarPedidoID)
	router.DELETE("/pedido/deletar/:id", pedido.DeletarPedido)
}
