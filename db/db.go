package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yuanyu90221/ethereum_blockchain_services/configs"
)

type DBClient struct {
	client *gorm.DB
}

func (m *DBClient) Connect() {
	config := configs.GetEnvConfig()
	client, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			config.POSTGRES_HOST,
			config.POSTGRES_PORT,
			config.POSTGRES_USER,
			config.POSTGRES_DB,
			config.POSTGRES_PASSWORD,
		),
	)
	if err != nil {
		panic(err)
	}
	m.client = client
}

func (m *DBClient) Disconnect() {
	m.client.Close()
}
