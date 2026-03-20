package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	handlers "test/internal/api/handlers/v2"
	"test/internal/config"
	service "test/internal/service/v2"
	"testing"

	"github.com/gin-gonic/gin"
)

type FakeDb struct {
	GetEntityByNameFunc func(ctx context.Context, name, language string) (json.RawMessage, error)
}

func (r FakeDb) GetEntityByName(ctx context.Context, name, language string) (json.RawMessage, error) {
	return r.GetEntityByNameFunc(ctx, name, language)
}

var (
	env = map[string]string{
		"APP_PORT":        "8080",
		"ALLOWED_ORIGINS": "*",
		"DB_USER":         "",
		"DB_PASS":         "",
		"DB_HOST":         "",
		"DB_PORT":         "",
		"DB_NAME":         "",
	}
)

func setupRouter() *gin.Engine {
	h := handlers.NewCvHandler(service.NewCvService(FakeDb{
		GetEntityByNameFunc: func(ctx context.Context, name, language string) (json.RawMessage, error) {
			resp := fmt.Sprintf(`{"name": "%s", "language": "%s"}`, name, language)
			return json.RawMessage(resp), nil
		},
	}))

	conf, _ := config.Load()

	return Register(h, conf)
}

func TestTheEndpoint(t *testing.T) {
	for name, value := range env {
		t.Setenv(name, value)
	}

	r := setupRouter()

	tests := []struct {
		language string
		result   string
	}{
		{"ru", `"language":"ru"`},
		{"en", `"language":"en"`},
	}

	for _, tt := range tests {
		t.Run(tt.language, func(t *testing.T) {
			target := fmt.Sprintf("/api/v2/cv?lang=%s", tt.language)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, target, nil)

			r.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("%s: %d", tt.language, rec.Code)
			}

			if !strings.Contains(rec.Body.String(), tt.result) {
				t.Fatalf("result does not match: %s", rec.Body.String())
			}
		})
	}
}
