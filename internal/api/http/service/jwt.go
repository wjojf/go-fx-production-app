package service

import (
	types "github.com/wjojf/go-uber-fx/internal/api/http/types/auth"
	"github.com/wjojf/go-uber-fx/internal/pkg/config"
	"github.com/wjojf/go-uber-fx/pkg/jwt"
	"time"
)

type JwtService struct {
	cfg config.Config
}

func NewJwtService(cfg config.Config) JwtService {
	return JwtService{
		cfg: cfg,
	}
}

func (s JwtService) RefreshAccessToken(refreshToken string) (types.TokenRefreshed, error) {
	claims := types.JwtPayload{}
	err := jwt.DecodeToken(refreshToken, jwt.Hs256, s.cfg.JwtSigningKey, &claims)
	if err != nil {
		return types.TokenRefreshed{}, err
	}

	// Generate new token with the same user ID
	acessToken, err := s.generateToken(claims.UserId, time.Duration(s.cfg.JwtAccessTokenLifetimeHours)*time.Hour)
	if err != nil {
		return types.TokenRefreshed{}, err
	}

	return types.TokenRefreshed{AccessToken: acessToken, UserID: claims.UserId}, nil
}

func (s JwtService) GenerateTokens(userId string) (types.TokenPair, error) {

	access, err := s.generateToken(userId, time.Duration(s.cfg.JwtAccessTokenLifetimeHours)*time.Hour)
	if err != nil {
		return types.TokenPair{}, err
	}

	refresh, err := s.generateToken(userId, time.Duration(s.cfg.JwtRefreshTokenLifetimeHours)*time.Hour)
	if err != nil {
		return types.TokenPair{}, err
	}

	return types.TokenPair{AccessToken: access, RefreshToken: refresh}, nil
}

func (s JwtService) generateToken(userId string, lifetime time.Duration) (string, error) {
	claims := types.NewPayload(userId, lifetime)
	return jwt.GenerateToken(jwt.Hs256, claims, s.cfg.JwtSigningKey)
}
