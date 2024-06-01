package Api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-rabbitmq-docker/config"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestPublish(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		gin.SetMode(gin.TestMode)
		app := gin.Default()
		routes := config.NewRoutes()
		routes.SetUpRoutes(app)

		// Create form data
		formData := url.Values{}
		formData.Set("number", "9")

		// Create a new HTTP request with form data
		req, err := http.NewRequest("POST", "/api/publish", strings.NewReader(formData.Encode()))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Create a response recorder to record the response
		rec := httptest.NewRecorder()

		// Serve the HTTP request to the router
		app.ServeHTTP(rec, req)

		// Check the response status code
		if rec.Code != http.StatusOK {
			t.Errorf("expected status %d but got %d", http.StatusOK, rec.Code)
		}

		// Decode the JSON response
		var jsonResponse map[string]interface{}
		err = json.NewDecoder(rec.Body).Decode(&jsonResponse)
		if err != nil {
			t.Errorf("Error decoding response: %v\n", err)
		}

		// Define the expected response
		expectedResponse := map[string]interface{}{
			"status_code": "GO-ES-001",
			"http_code":   200,
			"content": map[string]interface{}{
				"message": "Event successfully added in the queue!",
			},
		}

		// Compare specific fields in the response
		if jsonResponse["status_code"] != expectedResponse["status_code"] ||
			int(jsonResponse["http_code"].(float64)) != expectedResponse["http_code"] ||
			jsonResponse["content"].(map[string]interface{})["message"] != expectedResponse["content"].(map[string]interface{})["message"] {
			t.Errorf("Expected response %v, but got %v\n", expectedResponse, jsonResponse)
		}

	})

}
