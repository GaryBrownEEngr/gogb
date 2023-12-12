package gogb

import (
	"fmt"
	"net/http"
	"time"

	"github.com/GaryBrownEEngr/gogb/stacktrs"
	"github.com/GaryBrownEEngr/gogb/uerr"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type TokenMaker interface {
	Create(username string, duration time.Duration) (*Token, error)
	Verify(token *Token) (*Payload, error)
}

type Token string

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	if username == "" {
		return nil, fmt.Errorf("A username is required")
	}

	if duration < 0 {
		return nil, fmt.Errorf("The duration must be positive: %v", duration)
	}

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	ret := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return ret, nil
}

////
////

type pasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

var _ TokenMaker = &pasetoMaker{}

func NewPasetoMaker(symmetricKey string) (*pasetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("Invalid key size: must be %d bytes", chacha20poly1305.KeySize)
	}
	ret := &pasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return ret, nil
}

func (s *pasetoMaker) Create(username string, duration time.Duration) (*Token, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return nil, err
	}

	ret, err := s.paseto.Encrypt(s.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	token := Token(ret)

	return &token, nil
}

func (s *pasetoMaker) Verify(token *Token) (*Payload, error) {
	if token == nil {
		return nil, uerr.UErrLogHash("Token Verify Error", http.StatusInternalServerError, stacktrs.Errorf("Token is nil"))
	}

	payload := &Payload{}

	err := s.paseto.Decrypt(string(*token), s.symmetricKey, payload, nil)
	if err != nil {
		return nil, uerr.UErrLogHash("Token format invalid", http.StatusInternalServerError, fmt.Errorf("%#v", *token))
	}

	now := time.Now()
	if now.After(payload.ExpiredAt) {
		return nil, uerr.UErrLog("Token expired", http.StatusUnauthorized, fmt.Errorf(payload.Username))
	}

	// Make sure the token was created in the past, with 5 seconds of wiggle room.
	if now.Add(time.Second * 5).Before(payload.IssuedAt) {
		return nil, uerr.UErrLogHash("Token being used before created time", http.StatusUnauthorized, fmt.Errorf(payload.Username))
	}

	return payload, nil
}
