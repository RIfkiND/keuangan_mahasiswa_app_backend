package middleware

import (
    "context"
    "net/http"
    "os"
    "strings"

    "github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        jwtSecret := []byte(os.Getenv("JWT_SECRET"))
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Optionally, set user_id in context
        if claims, ok := token.Claims.(jwt.MapClaims); ok {
            if userID, ok := claims["user_id"].(float64); ok {
                ctx := context.WithValue(r.Context(), "user_id", uint(userID))
                r = r.WithContext(ctx)
            }
        }

        next.ServeHTTP(w, r)
    })
}