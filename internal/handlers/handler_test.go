package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Lovodia/ProxyAPI/internal/models"
	"golang.org/x/net/context"
)

type mockClient struct {
	GetPostFunc func(ctx context.Context, id int) (*models.Post, error)
}

func (m *mockClient) GetPost(ctx context.Context, id int) (*models.Post, error) {
	return m.GetPostFunc(ctx, id)
}

func TestGetPostHandler(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		mockClient     *mockClient
		expectedStatus int
		expectedBody   string
	}{
		{
			name:  "successful request",
			query: "id=1",
			mockClient: &mockClient{
				GetPostFunc: func(ctx context.Context, id int) (*models.Post, error) {
					return &models.Post{
						ID:     id,
						UserID: 1,
						Title:  "Test Title",
						Body:   "Test Body",
					}, nil
				},
			},

			expectedStatus: http.StatusOK,
			expectedBody:   `"title":"Test Title"`,
		},
		{
			name:           "missing id parameter",
			query:          "",
			mockClient:     nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "missing id",
		},
		{
			name:           "invalid id",
			query:          "id=abc",
			mockClient:     nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid id",
		},
		{
			name:  "client error",
			query: "id=123",
			mockClient: &mockClient{
				GetPostFunc: func(ctx context.Context, id int) (*models.Post, error) {
					return nil, errors.New("ferch error")
				},
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "failed to fetch post",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &Handler{Client: tt.mockClient}

			req := httptest.NewRequest(http.MethodGet, "/post?"+tt.query, nil)
			w := httptest.NewRecorder()

			handler.GetPostHandler(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}
			var body bytes.Buffer
			_, _ = io.Copy(&body, res.Body)

			if !strings.Contains(body.String(), tt.expectedBody) {
				t.Errorf("response body mismatch. expected to contain %q, got %q", tt.expectedBody, body.String())
			}
		})
	}
}
