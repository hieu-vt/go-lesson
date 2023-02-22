package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"lesson-5-goland/plugin/jwtprovider"
)

type jwtProvider struct {
	secret string
}

func NewTokenJwt(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload jwtprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

//func (j *jwtProvider) Generate(data jwtprovider.TokenPayload, expiry int) (*jwtprovider.Token, error) {
//	// generate the JWT
//	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
//		data,
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
//			IssuedAt:  time.Now().Local().Unix(),
//		},
//	})
//
//	myToken, err := t.SignedString([]byte(j.secret))
//	if err != nil {
//		return nil, err
//	}
//
//	// return the token
//	return &jwtprovider.Token{
//		Token:   myToken,
//		Expiry:  expiry,
//		Created: time.Now(),
//	}, nil
//}

func (j *jwtProvider) Validate(myToken string) (*jwtprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, jwtprovider.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, jwtprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, jwtprovider.ErrInvalidToken
	}

	// return the token
	return &claims.Payload, nil
}

func (j *jwtProvider) String() string {
	return "JWT implement Provider"
}
