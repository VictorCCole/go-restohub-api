package main

import (
	"github.com/VictorCCole/go-restohub-api/database"
	"github.com/VictorCCole/go-restohub-api/routers"
)

func main() {
	// Inicializar o banco de dados e executar migrações automáticas
	database.InitDB()
	defer database.CloseDB()

	// Configurar as rotas
	router := routers.SetupRoutes()

	// Iniciar o servidor
	router.Run(":8080")
}
