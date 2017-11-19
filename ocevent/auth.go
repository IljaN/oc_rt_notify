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
			respondWithError(err.Error(), c)
			return
		}

		if err = sig.Verify([]byte(sharedSecret), crypto.SigningMethodHS512); err != nil {
			respondWithError(err.Error(), c)
			return
		}

		j, err := jws.ParseJWTFromRequest(c.Request)

		if err != nil {
			respondWithError(err.Error(), c)
			return
		}

		if !j.Claims().Has("aud")  {
			respondWithError("Missing claims", c)
			return
		}

		auds, ok := j.Claims().Audience()

		if !ok || len(auds) < 1 {
			respondWithError("Missing claims", c)
			return
		}

		aud := auds[0]

		if aud == "subscriber" && ctx == Subscribing {
			sub, ok := j.Claims().Subject()

			if !ok {
				respondWithError("Missing claims", c)
				return
			}

			c.Set("username", sub)
			c.Next()
			return
		}


		if aud == "publisher" && ctx == Publishing {
			c.Next()
			return
		}

		respondWithError("Wrong audience", c)
		return
	}
}


func respondWithError(message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
}
