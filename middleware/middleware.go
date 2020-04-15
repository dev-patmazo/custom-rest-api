package middleware

import (
	"encoding/json"
	"net/http"
	"rest-api/helper"
	"strings"
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

var message = make(map[string]interface{})

// Configure this for session usage
func Tokenizer() {
	signer, _ := jwt.NewHS256([]byte(secretKey))
	builder := jwt.NewTokenBuilder(signer)

	token, _ := builder.Build(claims)
	log.Debug(string(token.Raw()))
}

func Interceptor(inner http.Handler) http.Handler {

	claimContainer := &jwt.StandardClaims{}

	mw := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		auth := r.Header.Get("Authorization")
		bearer := strings.Split(auth, "Bearer")
		token := strings.TrimSpace(bearer[1])

		rawToken, parseErr := jwt.Parse([]byte(token))
		if parseErr != nil {
			log.Debug(parseErr)
			message["code"] = http.StatusUnauthorized
			message["status"] = http.StatusText(http.StatusUnauthorized)
			helper.Response(message, w)
			return
		}

		_ = json.Unmarshal(rawToken.RawClaims(), claimContainer)

		log.Debug(claimContainer.IsExpired(time.Now()))

		if claimContainer.IsExpired(time.Now()) {
			message["code"] = http.StatusUnauthorized
			message["status"] = http.StatusText(http.StatusUnauthorized)
			helper.Response(message, w)
			return
		} else {
			validator := jwt.NewValidator(
				jwt.IDChecker(claimContainer.ID),
				jwt.IssuerChecker(claimContainer.Issuer),
				jwt.AudienceChecker(claimContainer.Audience),
			)

			validErr := validator.Validate(claims)
			if validErr != nil {
				log.Debug(validErr)
				message["code"] = http.StatusUnauthorized
				message["status"] = http.StatusText(http.StatusUnauthorized)
				helper.Response(message, w)
				return
			}
		}
	}

	return http.HandlerFunc(mw)

}
