package tests

import (
	"github.com/sergeysergeevru/go-blog-service/internal/config"
	"github.com/sergeysergeevru/go-blog-service/internal/controllers"
	"os"
	"testing"
)

var basicCfg = &config.Cfg{Server: config.Server{Prefix: "/api/v1"}}

func TestMain(m *testing.M) {
	root := controllers.CreateHandler(basicCfg)
	go func() {
		root.Run(":8080")
	}()
	code := m.Run()
	os.Exit(code)
}
