package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест для проверки ответа, когда в параметре city больше кафе, чем в данном городе
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=5", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// если статус не ОК, завершаем данный тест
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	body := responseRecorder.Body.String()
	city := strings.Split(body, ",")

	// проверяем количество кафе в ответе сервера с нужным количеством
	assert.Len(t, city, totalCount)
}

// тест для проверки статуса и тела ответа сервера при запросе с неверным городом
func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	rightResponse := "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?city=omsk&count=4", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	res := responseRecorder.Body.String()

	// проверки через пакет assert, чтобы проверить два условия сразу
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, rightResponse, res)
}

// проверка статуса и тела ответа сервера, при корректном запросе
func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	res := responseRecorder.Body.String()

	// проверки через пакет assert, чтобы проверить два условия сразу
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, res)
}
