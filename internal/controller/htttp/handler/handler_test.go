package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ZhdanovichVlad/service-podof/internal/controller/htttp/handler/mocks"
	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
)

func setupTest(t *testing.T) (*gin.Engine, *mocks.Mockapi, *Handler) {
    ctrl := gomock.NewController(t)
    mockAPI := mocks.NewMockapi(ctrl)
    logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
    handler := NewHandler(mockAPI, logger)

    gin.SetMode(gin.TestMode)
    router := gin.New()
    
    return router, mockAPI, handler
}

func TestHandler_DummyLogin(t *testing.T) {
    router, mockAPI, handler := setupTest(t)
    router.POST("/dummy-login", handler.DummyLogin)

    tests := []struct {
        name         string
        input        *entity.DummyLogin
        setupMock    func(*mocks.Mockapi)
        expectedCode int
        expectedBody map[string]interface{}
    }{
        {
            name: "success",
            input: &entity.DummyLogin{
                Role: "user",
            },
            setupMock: func(m *mocks.Mockapi) {
                m.EXPECT().
                    DummyLogin(gomock.Any(), gomock.Any()).
                    Return(&entity.JwtToken{Token: "test_token"}, nil)
            },
            expectedCode: http.StatusOK,
            expectedBody: map[string]interface{}{
                "message": "test_token", // изменено с "token" на "message" согласно JSON тегу
            },
        },
        {
            name: "empty_role",
            input: &entity.DummyLogin{},
            setupMock: func(m *mocks.Mockapi) {},
            expectedCode: http.StatusBadRequest,
            expectedBody: map[string]interface{}{
                "error": "empty field",
            },
        },
        {
            name: "internal_error",
            input: &entity.DummyLogin{
                Role: "user",
            },
            setupMock: func(m *mocks.Mockapi) {
                m.EXPECT().
                    DummyLogin(gomock.Any(), gomock.Any()).
                    Return(nil, errorsx.ErrInternal)
            },
            expectedCode: http.StatusInternalServerError,
            expectedBody: map[string]interface{}{
                "error": "internal server error",
            },
        },
    }


    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.setupMock(mockAPI)

            body, err := json.Marshal(tt.input)
            require.NoError(t, err)

            req := httptest.NewRequest(http.MethodPost, "/dummy-login", bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")
            w := httptest.NewRecorder()

            router.ServeHTTP(w, req)

            assert.Equal(t, tt.expectedCode, w.Code)

        })
    }
}




