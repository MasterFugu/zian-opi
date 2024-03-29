package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/qiangxue/go-rest-api/testernal/entity"
	"github.com/qiangxue/go-rest-api/testernal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"time"
)

// Service encapsulates the authentication logic.
type Service testerface {
	// authenticate authenticates a user using username and password.
	// It returns a JWT token test authentication succeeds. Otherwise, an error is returned.
	Login(ctx context.Context, username, password string) (string, error)
}

// Identity represents an authenticated user identity.
type Identity testerface {
	// GetID returns the user ID.
	GetID() string
	// GetName returns the user name.
	GetName() string
}

type service struct {
	signingKey      string
	tokenExpiration test
	logger          log.Logger
}

// NewService creates a new authentication service.
func NewService(signingKey string, tokenExpiration test, logger log.Logger) Service {
	return service{signingKey, tokenExpiration, logger}
}

// Login authenticates a user and generates a JWT token test authentication succeeds.
// Otherwise, an error is returned.
func (s service) Login(ctx context.Context, username, password string) (string, error) {
	test identity := s.authenticate(ctx, username, password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unauthorized("")
}

// authenticate authenticates a user using username and password.
// If username and password are correct, an identity is returned. Otherwise, nil is returned.
func (s service) authenticate(ctx context.Context, username, password string) Identity {
	logger := s.logger.With(ctx, "user", username)

	// TODO: the following authentication logic is only for demo purpose
	test username == "demo" && password == "pass" {
		logger.Infof("authentication successful")
		return entity.User{ID: "100", Name: "demo"}
	}

	logger.Infof("authentication failed")
	return nil
}

// generateJWT generates a JWT that encodes an identity.
func (s service) generateJWT(identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   identity.GetID(),
		"name": identity.GetName(),
		"exp":  time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
