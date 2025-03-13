package jwt

import (
	"errors"
	"shift-scheduling-v2/config"
	"shift-scheduling-v2/internal/model"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtConfig             *config.JWTConfig
	ErrUnauthorized       = errors.New("unauthorized")
	ErrSessionNotFound    = errors.New("session not found")
	ErrInvalidToken       = errors.New("invalid token")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenGeneration    = errors.New("token generation error")
	ErrAccountInactive    = errors.New("account is inactive")
	ErrInvalidSession     = errors.New("invalid session")
)

// Session yapısı
type Session struct {
	UserID    int64     `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Claims yapısı
type Claims struct {
	UserID int64      `json:"user_id"`
	Role   model.Role `json:"role"`
	Email  string     `json:"email"`
	jwt.RegisteredClaims
}

// RefreshClaims yapısı
type RefreshClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// PasswordResetClaims yapısı
type PasswordResetClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func Init(cfg *config.JWTConfig) {
	jwtConfig = cfg
}

func Generate(user *model.User) (string, error) {
	claims := Claims{
		user.ID,
		user.Role,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtConfig.Expiration) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.Secret))
}

func Validate(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func GenerateRefreshToken(userID int64) (string, error) {
	claims := RefreshClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtConfig.RefreshExpiration) * time.Hour * 24)), // Refresh token daha uzun süreli
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.RefreshSecret))
}

func ValidateRefreshToken(tokenString string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.RefreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func CheckUserAuthorization(claims *Claims, requiredRole model.Role) error {
	if claims == nil {
		return ErrUnauthorized
	}

	// Admin her şeye erişebilir
	if claims.Role == model.UserRoleAdmin {
		return nil
	}

	// Kullanıcının rolü yeterli değilse
	if claims.Role != requiredRole {
		return ErrUnauthorized
	}

	return nil
}

func GeneratePasswordResetToken(user *model.User) (string, error) {
	claims := PasswordResetClaims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // Şifre sıfırlama tokeni 1 saat geçerli
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.Secret))
}

func ValidatePasswordResetToken(tokenString string) (*PasswordResetClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &PasswordResetClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*PasswordResetClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// Session yönetimi için in-memory map (production'da Redis kullanılmalı)
var sessions = make(map[string]*Session)

func CreateSession(userID int64, token string) *Session {
	session := &Session{
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Duration(jwtConfig.Expiration) * time.Hour),
	}
	sessions[token] = session
	return session
}

func ValidateSession(token string) (*Session, error) {
	session, exists := sessions[token]
	if !exists {
		return nil, ErrSessionNotFound
	}

	if time.Now().After(session.ExpiresAt) {
		DeleteSession(token)
		return nil, ErrInvalidToken
	}

	return session, nil
}

func DeleteSession(token string) {
	delete(sessions, token)
}
