package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"note-system/internal/domain"
	"note-system/internal/storage"
	"note-system/pkg/logging"
	"time"

	"note-system/pkg/jwt"
)

const (
	saltSize        = 16
	accessTokenTTL  = time.Hour * 24
	refreshTokenTTL = time.Hour * 24 * 7
)

type AuthService struct {
	storage    storage.Auth
	logger     *logging.Logger
	jwtManager *jwt.JWTManager
}

func NewAuthService(storage storage.Auth, logger *logging.Logger, manager *jwt.JWTManager) *AuthService {
	return &AuthService{
		storage:    storage,
		logger:     logger,
		jwtManager: manager,
	}
}

func (s *AuthService) SignUp(ctx context.Context, accountDTO domain.CreateAccountDTO) (int, error) {
	salt := generateSalt(saltSize)
	hashedPassword := generatePasswordHash(accountDTO.Password, salt)

	account := domain.Account{
		Username:     accountDTO.Username,
		PasswordHash: hashedPassword,
		HashSalt:     hex.EncodeToString(salt),
	}

	return s.storage.SignUp(ctx, account)
}

func (s *AuthService) SignIn(ctx context.Context, accountDTO domain.LoginAccountDTO) (string, error) {
	s.logger.Debugf("singing in: %s", accountDTO.Username)
	var account domain.Account

	account, err := s.storage.SignIn(ctx, accountDTO)
	if err != nil {
		return "", err
	}
	s.logger.Debugf("comparing passwords: %s", accountDTO.Username)
	if compare := comparePassword(accountDTO.Password, account.PasswordHash,
		account.HashSalt); !compare {
		return "", errors.New("invalid password")
	}

	s.logger.Debugf("generating token: %s", accountDTO.Username)
	token := s.jwtManager.NewJWT(account.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func comparePassword(password, passwordHash, salt string) bool {
	saltBytes, err := hex.DecodeString(salt)
	if err != nil {
		return false
	}

	providedPassword := generatePasswordHash(password, saltBytes)
	if passwordHash != providedPassword {
		return false
	}

	return true
}

func generatePasswordHash(password string, salt []byte) string {
	var passwordBytes = []byte(password)
	var sha256Hasher = sha256.New()
	passwordBytes = append(passwordBytes, salt...)

	sha256Hasher.Write(passwordBytes)

	hashedPasswordBytes := sha256Hasher.Sum(nil)

	return hex.EncodeToString(hashedPasswordBytes)
}

func generateSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	rand.Read(salt[:])

	return salt
}
