package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Lovodia/ProxyAPI/internal/client"
)

type Handler struct {
	Client client.PostGetter
}

func New(c *client.HTTPClient) *Handler {
	return &Handler{Client: c}
}

// GetPostHandler godoc
// @Summary Получить пост по ID
// @Description Проксирует запрос к внешнему API (jsonplaceholder)
// @Tags posts
// @Accept json
// @Produce json
// @Param id query int true "ID поста"
// @Success 200 {object} models.Post
// @Failure 400 {string} string "Неверный запрос (например, отсутствует параметр id)"
// @Failure 500 {string} string "Ошибка при запросе к внешнему API"
// @Router /post [get]
func (h *Handler) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id paramter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}
	ctx := context.Background()

	post, err := h.Client.GetPost(ctx, id)
	if err != nil {
		http.Error(w, "failed to fetch post:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
