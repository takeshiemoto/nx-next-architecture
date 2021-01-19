package infrastructure

import (
	"toh-echo/interfaces/database"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func (h *SqlHandler) Create(object interface{}) {
	h.db.Create(object)
}

func (h *SqlHandler) FindAll(object interface{}) {
	h.db.Find(object)
}

func (h *SqlHandler) DeleteById(object interface{}, id string) {
	h.db.Delete(object, id)
}

func NewSqlHandler() database.SqlHandler {
	dsn := "host=localhost user=toh dbname=toh password=toh sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}
