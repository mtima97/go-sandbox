package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	handlers "test/internal/api/handlers/v2"
	"test/internal/config"
	service "test/internal/service/v2"
	"testing"

	"github.com/gin-gonic/gin"
)

type FakeDb struct{}

func (r FakeDb) GetEntityByName(ctx context.Context, name, language string) (json.RawMessage, error) {
	return json.RawMessage("[]"), nil
}

func setupRouter(c func(k, v string)) *gin.Engine {
	h := handlers.NewCvHandler(service.NewCvService(FakeDb{}))

	envMock := map[string]string{
		"APP_PORT":        "8080",
		"ALLOWED_ORIGINS": "*",
		"DB_USER":         "",
		"DB_PASS":         "",
		"DB_HOST":         "",
		"DB_PORT":         "",
		"DB_NAME":         "",
	}

	for k, v := range envMock {
		c(k, v)
	}

	conf, _ := config.Load()

	return Register(h, conf)
}

func TestTheEndpoint(t *testing.T) {
	r := setupRouter(func(k, v string) {
		t.Setenv(k, v)
	})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v2/cv?lang=ru", nil)

	if rec.Code != http.StatusOK {
		t.Fatal(rec.Code)
	}

	r.ServeHTTP(rec, req)
}
