package precode

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := http.Request{Method: http.MethodGet, URL: &url.URL{Path: "http://localhost:8080/cafe?count=5&city=moscow"}} // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, &req)

	// здесь нужно добавить необходимые проверки
	countStr := req.URL.Query().Get("count")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		fmt.Printf("wrong count value %s\n", err.Error())
		return
	}

	assert.Len(t, cafeList[req.URL.Query().Get("city")], count)
}
