package restaurante

import (
	"github.com/VictorCCole/go-restohub-api/handlers/restaurante"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/restaurante/criar", restaurante.CriarRestaurante)
	router.PUT("/restaurante/atualizar/:id", restaurante.AtualizarRestaurante)
	router.GET("/restaurante/listar/", restaurante.ListarRestaurante)
	router.GET("/restaurante/listarid/:id", restaurante.ListarRestauranteID)
	router.DELETE("/restaurante/deletar/:id", restaurante.DeletarRestaurante)
}
