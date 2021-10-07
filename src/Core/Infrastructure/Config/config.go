package config_core

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type Config struct {
	dbPost *pg.DB
	route *gin.Engine
}

/**
Only bootstrap initiate
 */
func newConfig(dbPost *pg.DB, route *gin.Engine) *Config {
	return &Config{
		dbPost,
		route,
	}
}

func (c Config) GetDB() *pg.DB {
	return c.dbPost
}

func (c Config) GetEngine() *gin.Engine {
	return c.route
}

func (c Config) AddRoute(addr ...string) error {

	return c.route.Run(addr...)
}
