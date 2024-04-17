/*
Package database gerencia a conexão com o banco de dados e as operações relacionadas.

Este arquivo contém funções para inicializar e fechar a conexão com o banco de dados, bem como para obter a instância do banco de dados para uso em outras partes da aplicação.
*/

package database

import (
	"log"
	"sync"

	"github.com/VictorCCole/go-restohub-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// InitDB inicializa a conexão com o banco de dados e executa migrações automáticas.
// Retorna a instância do banco de dados.
func InitDB() *gorm.DB {
	dbOnce.Do(func() {
		// Configuração da conexão com o banco de dados PostgreSQL
		dsn := "host=localhost user=postgres password=postgres dbname=db_restohub port=5432 sslmode=disable"
		database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Falha ao conectar com o banco de dados: %v", err)
		}

		// Executar migrações automáticas para criar as tabelas no banco de dados, se necessário
		err = database.AutoMigrate(&models.Pessoa{}, &models.Restaurante{}, &models.Pedido{})
		if err != nil {
			log.Fatalf("Falha ao executar migrações automáticas: %v", err)
		}

		db = database
	})

	return db
}

// GetDB retorna a instância do banco de dados.
// Se a instância ainda não foi inicializada, chama InitDB para inicializá-la e a retorna.
func GetDB() *gorm.DB {
	return InitDB()
}

// CloseDB fecha a conexão com o banco de dados.
func CloseDB() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Erro ao obter conexão do banco de dados: %v", err)
		}
		sqlDB.Close()
	}
}
