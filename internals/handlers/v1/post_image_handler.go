package v1

import (
    "net/http"
    "keuangan/backend/internals/services"
    // "mime/multipart"
    "strings"
)

type PostImageHandler struct {
    Service *services.PostImageService
}

func NewPostImageHandler(service *services.PostImageService) *PostImageHandler {
    return &PostImageHandler{Service: service}
}

func (h *PostImageHandler) Upload(w http.ResponseWriter, r *http.Request) {
    // Parse multipart form
    err := r.ParseMultipartForm(10 << 20) // 10 MB
    if err != nil {
        http.Error(w, "Could not parse multipart form", http.StatusBadRequest)
        return
    }

    file, handler, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Could not get image file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Optionally sanitize filename
    fileName := handler.Filename
    fileName = strings.ReplaceAll(fileName, " ", "_")

    url, err := h.Service.UploadImage(file, fileName)
    if err != nil {
        http.Error(w, "Failed to upload image: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"url":"` + url + `"}`))
}