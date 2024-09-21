package gorm

import (
	"fmt"
	"github.com/askaroe/hexago/internal/domain/port/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type appDB struct {
	db             *gorm.DB
	defaultTimeout time.Duration
}

func NewDB() (database.IDB, error) {
	dsn := os.Getenv("DB_URL")

	if dsn == "" {
		return nil, fmt.Errorf("dsn is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	defaultTimeout := time.Second * 60

	return &appDB{db, defaultTimeout}, err
}

func (c *appDB) Close() (err error) {
	sqlDB, err := c.db.DB()

	if err != nil {
		return
	}

	err = sqlDB.Close()

	return
}
