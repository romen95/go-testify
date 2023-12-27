package precode

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func GetResponse(url string) *httptest.ResponseRecorder {
	request := httptest.NewRequest(http.MethodGet, url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, request)

	return responseRecorder
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	expectedCount := len(cafeList["moscow"])

	responseRecorder := GetResponse(fmt.Sprintf("/cafe?count=%d&city=moscow", expectedCount+5))

	require.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, expectedCount)
}

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
	expectedCount := len(cafeList["moscow"])

	responseRecorder := GetResponse(fmt.Sprintf("/cafe?count=%d&city=moscow", expectedCount+5))

	require.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, expectedCount)
}

func TestMainHandlerWhenIncorrectCity(t *testing.T) {
	expectedCode := http.StatusBadRequest
	expectedBody := `wrong city value`

	responseRecorder := GetResponse("/cafe?count=5&city=vladyvostok")

	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := responseRecorder.Body.String()
	assert.Equal(t, expectedBody, body)
}
