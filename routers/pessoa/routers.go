package pessoa

import (
	"github.com/VictorCCole/go-restohub-api/handlers/pessoa"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/pessoa/criar", pessoa.CriarPessoa)
	router.PUT("/pessoa/atualizar/:id", pessoa.AtualizarPessoa)
	router.GET("/pessoa/listar/", pessoa.ListarPessoas)
	router.GET("/pessoa/listarid/:id", pessoa.ListarPessoasID)
	router.DELETE("/pessoa/deletar/:id", pessoa.DeletarPessoa)
}
