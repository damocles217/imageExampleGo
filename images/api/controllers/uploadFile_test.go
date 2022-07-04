package controllers_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/damocles217/images_service/images/api/database"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/mock"
// )

// // TODO
// // ! Complete this test file
// func TestUploadFile(t *testing.T) {
// 	client := &database.MockTodoClient{}
// 	tests := map[string]struct {
// 		payload      string
// 		expectedCode int
// 	}{
// 		"should return 200": {
// 			payload:      `{"userId":1,"title":"learning golang","completed":false}`,
// 			expectedCode: 200,
// 		},
// 		"should return 400": {
// 			payload:      "invalid json string",
// 			expectedCode: 400,
// 		},
// 	}

// 	for name, test := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			client.On("Insert", mock.Anything).Return(models.Todo{}, nil)
// 			req, _ := http.NewRequest("POST", "/todos", strings.NewReader(test.payload))
// 			rec := httptest.NewRecorder()

// 			r := gin.Default()
// 			r.POST("/todos", handlers.InsertTodo(client))
// 			r.ServeHTTP(rec, req)

// 			if test.expectedCode == 200 {
// 				client.AssertExpectations(t)
// 			} else {
// 				client.AssertNotCalled(t, "Insert")
// 			}
// 		})
// 	}

// }
