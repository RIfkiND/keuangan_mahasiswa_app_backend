package services

import (
    "os"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "keuangan/backend/internals/models"
    "keuangan/backend/internals/repositories"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type AuthService struct {
    UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
    return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Register(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return s.UserRepo.Create(user)
}

func (s *AuthService) Authenticate(email, password string) (*models.User, error) {
    user, err := s.UserRepo.FindByEmail(email)
    if err != nil {
        return nil, err
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, err
    }
    return user, nil
}

func (s *AuthService) GenerateToken(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })
    return token.SignedString(jwtSecret)
}