package main

import (
    "gorestapi/config"
    "gorestapi/internals/repositories"
    "gorestapi/internals/services"
    "gorestapi/internals/handlers"
    "net/http"
)

func initApp() (
    postService *services.PostService,
    categoryService *services.CategoryService,
    authService *services.AuthService,
    postImageService *services.PostImageService,
) {
    config.ConnectGorm()

    postRepo := repositories.NewPostRepository(config.DB)
    categoryRepo := repositories.NewCategoryRepository(config.DB)
    userRepo := repositories.NewUserRepository(config.DB)
    postImageRepo := repositories.NewPostImageRepository()

    postService = services.NewPostService(postRepo)
    categoryService = services.NewCategoryService(categoryRepo)
    authService = services.NewAuthService(userRepo)
    postImageService = services.NewPostImageService(postImageRepo)

    return
}

func main() {
    postService, categoryService, authService, postImageService := initApp()
    handlers.RegisterRoutes("/v1", postService, categoryService, authService, postImageService)
    http.ListenAndServe(":8080", nil)
}