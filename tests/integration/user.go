package integration

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		expectedFields []string
	}{
		{
			name:           "GET /api/users возвращает список пользователей",
			method:         "GET",
			path:           "/api/users",
			expectedStatus: http.StatusOK,
			expectedFields: []string{"id", "name", "email"},
		},
		{
			name:           "GET /api/users/1 возвращает одного пользователя",
			method:         "GET",
			path:           "/api/users/1",
			expectedStatus: http.StatusOK,
			expectedFields: []string{"id", "name", "email"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if w.Code == http.StatusOK {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				for _, field := range tt.expectedFields {
					_, exists := response[field]
					assert.True(t, exists, "Поле %s должно присутствовать в ответе", field)
				}
			}
		})
	}
}
