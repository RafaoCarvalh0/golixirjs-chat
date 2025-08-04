package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func Test_JWTAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	token, err := CreateMockToken("myuser")
	if err != nil {
		t.Fatalf("error creating mock token: %v", err)
	}

	router := gin.New()
	router.Use(JWTAuthMiddleware())
	router.GET("/protected", func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "userID not found in context"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"userID": userID})
	})

	t.Log("Given a valid http request with a valid jwt.")
	{
		req := httptest.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		testDescription := "\t\tReturns http status ok."
		if resp.Code != http.StatusOK {
			t.Fatal(testDescription, resp, req, ballotX)
		}
		t.Log(testDescription, checkMark)

		var body map[string]any
		err = json.Unmarshal(resp.Body.Bytes(), &body)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}

		userID, ok := body["userID"].(string)
		if !ok {
			t.Fatal("userID not found in response body")
		}

		testDescription = "\t\tExtract user ID from JWT claims."
		expectedUserID := "usr0001"
		if userID != expectedUserID {
			t.Fatalf(testDescription, expectedUserID, userID)
		}
		t.Log(testDescription, checkMark)
	}

	t.Log("Given a valid http request with an expired jwt.")
	{
		expiredToken, err := CreateMockTokenExpired("myuser")
		if err != nil {
			t.Fatalf("error creating mock token: %v", err)
		}

		req := httptest.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+expiredToken)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		var body map[string]any
		err = json.Unmarshal(resp.Body.Bytes(), &body)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}

		testDescription := "\t\tReturns body with error 'invalid token'."
		if body["error"] != "invalid token" {
			t.Fatal(testDescription, resp.Body, ballotX)
		}
		t.Log(testDescription, checkMark)
	}

}

func CreateMockToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":      "usr0001",
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateMockTokenExpired(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":      "usr0001",
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24 * -1).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
