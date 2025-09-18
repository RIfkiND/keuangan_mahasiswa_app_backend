package v1

import (
    "encoding/json"
    "net/http"
    "strconv"
    "keuangan/backend/internals/models"
    "keuangan/backend/internals/services"
    "keuangan/backend/internals/core"
)


// var validate = validator.New()

type PostHandler struct {
    Service *services.PostService
}


func NewPostHandler(service *services.PostService) *PostHandler {
    return &PostHandler{Service: service}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
    var post models.Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := core.Validate.Struct(post); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.Service.Create(&post); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    post, err := h.Service.GetByID(uint(id))
    if err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    posts, err := h.Service.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    var post models.Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    post.ID = uint(id)
    if err := h.Service.Update(&post); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    if err := h.Service.Delete(uint(id)); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Deleted"})
}