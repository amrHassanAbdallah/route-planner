package api

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouteEndpoint(t *testing.T) {
	// Define the payload
	payload := map[string]interface{}{
		"route": [][]string{
			{"SFO", "EWR"},
		},
	}

	// Encode the payload as JSON
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new request with the encoded JSON payload
	req, err := http.NewRequest("POST", "/route", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rec := httptest.NewRecorder()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	server := NewServer(logger)
	// Call the endpoint with the request and recorder
	handler := http.HandlerFunc(server.PostRoute)
	handler.ServeHTTP(rec, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rec.Code, "unexpected status code")

	// Decode the response body
	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response body
	assert.Equal(t, "SFO", response["source"], "unexpected source")
	assert.Equal(t, "EWR", response["destination"], "unexpected destination")
}
