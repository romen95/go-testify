package precode

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=5&city=moscow", nil)

	expectedCount := len(cafeList["moscow"])

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, expectedCount)
}

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=5&city=moscow", nil)

	expectedCode := http.StatusOK
	expectedCount := len(cafeList["moscow"])

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, expectedCount)
}

func TestMainHandlerWhenIncorrectCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=5&city=vladyvostok", nil)

	expectedCode := http.StatusBadRequest
	expectedBody := `wrong city value`

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := responseRecorder.Body.String()
	assert.Equal(t, expectedBody, body)
}
