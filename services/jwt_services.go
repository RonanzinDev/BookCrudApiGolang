package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtServices struct {
	secretKey string
	issure    string
}

func NewJwtServices() *jwtServices {
	return &jwtServices{
		secretKey: "secret-key", //FIMS ACADEMICOS
		issure:    "book-api",
	}
}

type Claim struct {
	Sub uint `json:"sub"` // referencia ao id do usuario, caso vc for no jwt.io

	jwt.StandardClaims
}

func (s *jwtServices) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    s.issure,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil

}

func (s *jwtServices) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}
