package main

import (
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/gin-gonic/gin"

	"net/http"
)

const sharedSecret = "186163c9826c3a0762319a81a3889dd9"

type AuthContext int

const (
	Publishing  AuthContext = 0
	Subscribing AuthContext = 1
)

func Authenticate(ctx AuthContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		sig, err := jws.ParseFromHeader(c.Request, jws.Compact)

		if err != nil {
			respondWithError(http.StatusUnauthorized, err.Error(), c)
			return
		}

		if err = sig.Verify([]byte(sharedSecret), crypto.SigningMethodHS512); err != nil {
			respondWithError(http.StatusUnauthorized, err.Error(), c)
			return
		}

		j, err := jws.ParseJWTFromRequest(c.Request)

		if err != nil {
			respondWithError(http.StatusUnauthorized, err.Error(), c)
			return
		}

		sub, ok := j.Claims().Subject()

		// We only need subject for auth as subscriber
		if !ok && ctx != Subscribing {
			respondWithError(http.StatusUnauthorized, err.Error(), c)
			return
		}

		aud, ok := j.Claims().Audience()


		if !ok {
			respondWithError(http.StatusUnauthorized, err.Error(), c)
			return
		}

		if len(aud) > 0 && aud[0] == "subscriber" && ctx == Subscribing {
			c.Set("username", sub)
			c.Next()
			return
		}


		if len(aud) == 0 && aud[0] == "publisher" && ctx == Publishing {
			c.Next()
			return
		}


		respondWithError(http.StatusUnauthorized, "Audience miss-match", c)
		return
	}
}


func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.AbortWithStatusJSON(code, resp)
}
