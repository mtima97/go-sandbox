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

type FakeDb struct {
	//
}

func (r FakeDb) GetEntityByName(ctx context.Context, name, language string) (json.RawMessage, error) {
	return json.RawMessage("[]"), nil
}

func setupRouter() *gin.Engine {
	h := handlers.NewCvHandler(service.NewCvService(FakeDb{}))

	c := config.Cfg{
		AllowedOrigins: []string{"*"},
	}

	return Register(h, c)
}

func TestTheEndpoint(t *testing.T) {
	r := setupRouter()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v2/cv?lang=ru", nil)

	if rec.Code != http.StatusOK {
		t.Fatal(rec.Code)
	}

	r.ServeHTTP(rec, req)
}
