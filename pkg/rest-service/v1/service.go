package v1

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

type RestServer struct {
	Db  *gorm.DB
	Mux *http.ServeMux
}

func NewRestServer(db *gorm.DB) *RestServer {
	return &RestServer{Db: db}
}
