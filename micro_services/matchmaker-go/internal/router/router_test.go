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

func TestCreateMatch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.NewRouter()

	token1 := createMockToken("usr0001")
	req1 := httptest.NewRequest(http.MethodPost, "/create-match", nil)
	req1.Header.Set("Authorization", "Bearer "+token1)

	resp1 := httptest.NewRecorder()
	r.ServeHTTP(resp1, req1)

	expectedResponse := `{"data":{"error":false,"match":{"User":{"ID":""},"UserPair":{"ID":""}},"message":"waiting for a pair..."}}`
	if resp1.Body.String() != expectedResponse {
		t.Fatal("Expected body:", expectedResponse, "Got:", resp1.Body.String())
	}

	token2 := createMockToken("usr0002")
	req2 := httptest.NewRequest(http.MethodPost, "/create-match", nil)
	req2.Header.Set("Authorization", "Bearer "+token2)

	resp2 := httptest.NewRecorder()
	r.ServeHTTP(resp2, req2)

	expectedResponse = `{"data":{"error":false,"match":{"User":{"ID":"usr0002"},"UserPair":{"ID":"usr0001"}},"message":"match created"}}`
	if resp2.Body.String() != expectedResponse {
		t.Fatal("Expected body:", expectedResponse, "Got:", resp2.Body.String())
	}

	token3 := createMockToken("usr0003")
	req3 := httptest.NewRequest(http.MethodPost, "/create-match", nil)
	req3.Header.Set("Authorization", "Bearer "+token3)

	resp3 := httptest.NewRecorder()
	r.ServeHTTP(resp3, req3)

	expectedResponse = `{"data":{"error":false,"match":{"User":{"ID":""},"UserPair":{"ID":""}},"message":"waiting for a pair..."}}`
	if resp3.Body.String() != expectedResponse {
		t.Fatal("Expected body:", expectedResponse, "Got:", resp3.Body.String())
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
