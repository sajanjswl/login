package v1

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/sajanjswl/auth/config"
	"go.uber.org/zap"
)

type RestServer struct {
	Db     *gorm.DB
	Mux    *http.ServeMux
	cfg    *config.Config
	logger *zap.Logger
}

func NewRestServer(db *gorm.DB, cfg *config.Config, logger *zap.Logger) *RestServer {
	return &RestServer{
		Db:     db,
		cfg:    cfg,
		logger: logger,
	}
}
