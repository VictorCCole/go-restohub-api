package routers

import (
	"github.com/VictorCCole/go-restohub-api/routers/pedido"
	"github.com/VictorCCole/go-restohub-api/routers/pessoa"
	"github.com/VictorCCole/go-restohub-api/routers/restaurante"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	pessoa.SetupRoutes(router)

	restaurante.SetupRoutes(router)

	pedido.SetupRoutes(router)

	return router
}
