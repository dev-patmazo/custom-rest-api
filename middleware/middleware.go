package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cristalhq/jwt"
	log "github.com/sirupsen/logrus"
)

var (
	id        = "4F67353C9762CBC7"
	secretKey = "7QHFEtZH6PlJpRyAgq5opm4cC1s9itQ7"
	issuer    = "patz.garcia"
	audience  = []string{"company-example"}
	expiry    = time.Now().Add(time.Hour + time.Duration(2)).Unix() // 2hrs
)

var claims = &jwt.StandardClaims{
	ID:        id,
	Issuer:    issuer,
	Audience:  audience,
	ExpiresAt: jwt.Timestamp(expiry),
}
var sampleToken *jwt.Token
var testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb21wYW55LWV4YW1wbGUiLCJleHAiOjE1ODY3OTg0MjYsImp0aSI6IjRGNjczNTNDOTc2MkNCQzciLCJpc3MiOiJwYXR6LmdhcmNpYSJ9.W4HIxAtASQ8a3Jjxs2R8lHMJIlgheLOhp3aTIhK4i8s"

func Tokenizer() {

	signer, _ := jwt.NewHS256([]byte(secretKey))
	builder := jwt.NewTokenBuilder(signer)

	token, _ := builder.Build(claims)
	//raw := token.Raw()
	fmt.Println(token)
	sampleToken = token
}

func Detokenizer() {

	newClaim := &jwt.StandardClaims{}
	_ = json.Unmarshal(sampleToken.RawClaims(), newClaim)

	validator := jwt.NewValidator(
		jwt.IDChecker(newClaim.ID),
		jwt.IssuerChecker(newClaim.Issuer),
		jwt.AudienceChecker(newClaim.Audience),
		jwt.ExpirationTimeChecker(time.Now()),
	)

	err := validator.Validate(newClaim)
	if err != nil {
		log.Error(err)
	}

}
