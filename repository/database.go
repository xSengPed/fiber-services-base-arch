package repository

import (
	"fmt"
	"go-fiber-basic/config"
	"go-fiber-basic/datamodels"

	"golang.org/x/exp/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *config.DbConfig) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DbName,
		cfg.Port,
		cfg.Ssl,
		cfg.Timezone,
	)

	slog.Info(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&datamodels.User{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	slog.Info("Database connected")

	DB = db

}
