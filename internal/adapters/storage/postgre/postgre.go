package postgre

import (
	"context"
	"fmt"

	"github.com/mrfade/case-sss/internal/adapters/configs"
	"github.com/mrfade/case-sss/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DSN string
	DB  *gorm.DB
	Ctx context.Context
}

func NewDB(cnf *configs.DB) *DB {
	//dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cnf.Host, cnf.User, cnf.Password, cnf.DbName, cnf.Port)
	return &DB{
		DSN: dsn,
		Ctx: context.Background(),
	}
}

func Open(db *DB) (err error) {
	db.DB, err = gorm.Open(postgres.Open(db.DSN))
	if err != nil {
		return errors.ErrUnableToConnectDB
	}

	return nil
}

func Migrate(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models...)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Close() {
	sqlDB, _ := db.DB.DB()
	sqlDB.Close()
}
