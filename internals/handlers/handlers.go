package handlers

import (
    "keuangan/backend/internals/handlers/v1"
    "keuangan/backend/internals/services"
    "net/http"
    "strings"
)

func RegisterRoutes(
    prefix string,
    postService *services.PostService,
    categoryService *services.CategoryService,
    authService *services.AuthService,
    postImageService *services.PostImageService,
) {
    // Ensure prefix starts with "/" and does not end with "/"
    if !strings.HasPrefix(prefix, "/") {
        prefix = "/" + prefix
    }
    prefix = strings.TrimSuffix(prefix, "/")

    postHandler := v1.NewPostHandler(postService)
    categoryHandler := v1.NewCategoryHandler(categoryService)
    authHandler := v1.NewAuthHandler(authService)
    postImageHandler := v1.NewPostImageHandler(postImageService)

    // Auth
    http.HandleFunc(prefix+"/auth/register", authHandler.Register)
    http.HandleFunc(prefix+"/auth/login", authHandler.Login)

    // Category
    http.HandleFunc(prefix+"/categories/create", categoryHandler.Create)
    http.HandleFunc(prefix+"/categories/all", categoryHandler.GetAll)
    http.HandleFunc(prefix+"/categories/update", categoryHandler.Update)
    http.HandleFunc(prefix+"/categories/delete", categoryHandler.Delete)

    // Post
    http.HandleFunc(prefix+"/posts/create", postHandler.Create)
    http.HandleFunc(prefix+"/posts/get", postHandler.GetByID)
    http.HandleFunc(prefix+"/posts/all", postHandler.GetAll)
    http.HandleFunc(prefix+"/posts/update", postHandler.Update)

    // Post Image
    http.HandleFunc(prefix+"/posts/images/upload", postImageHandler.Upload)
    http.HandleFunc(prefix+"/posts/delete", postHandler.Delete)
}