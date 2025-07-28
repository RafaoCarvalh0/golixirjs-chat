package router_test

import (
	"matchmaker-go/internal/router"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

func createMockToken(sub string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenStr, _ := token.SignedString(secretKey)
	return tokenStr
}

func TestRouteExists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.NewRouter()

	req := httptest.NewRequest(http.MethodPost, "/create-match", nil)
	token := createMockToken("usr0001")
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code == http.StatusNotFound {
		t.Errorf("Expected route to exist, got 404")
	}
}

func TestWithoutAuthorizationHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.NewRouter()

	req := httptest.NewRequest(http.MethodPost, "/create-match", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 Unauthorized, got %d", resp.Code)
	}
}

func TestWithInvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.NewRouter()

	req := httptest.NewRequest(http.MethodPost, "/create-match", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken123")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401 Unauthorized for invalid token, got %d", resp.Code)
	}
}

func TestWithValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.NewRouter()

	token := createMockToken("usr0001")
	req := httptest.NewRequest(http.MethodPost, "/create-match", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK && resp.Code != http.StatusInternalServerError {
		t.Errorf("Expected status 200 or 500 for valid token, got %d", resp.Code)
	}
}
