package routers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
	"github.com/spidernest-go/mux/middleware"
)

var r *echo.Echo

const (
	ErrGeneric   = `{"errno": "404", "message": "Bad Request"}`
	rsaPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAl3Mw0lnzWr+KrhhP1/jnKHblCM/DqIhvUHgsOYZWrE3+fEvHjc6wrUT9RtC3eRZfRxtdyxa9CPuSnPEt/Jmu2YPVRWxOVUJfUxgZQg0OPXurMy0h6O1Yal4s9yNq0+OmCSIE3DFVNTs5hlYNI7TNkjPp/UJx8Xc+J+g/gUPrIVQo+XWNGoKv+udiQhi9LrYZuQOy9MZPKgUKSfJwmwWRBb7CZmvWSwprQ3/619+2vf1gS/K3vqenlZfCRFadPuxebmQ595LKAn0tgnw2R0c4aAU/G1LsJsBFfY0kvhE/asFvNSoAoJA3jnQMYmMekqgVdVNV2FLrLWve5520EjTeMQIDAQAB
-----END PUBLIC KEY-----`
)

func ListenAndServe() {
	// decode pem block into rsa public key
	block, _ := pem.Decode([]byte(rsaPublicKey))
	if block == nil {
		logger.Fatal().
			Msg("PEM RSA public key block was invalid and failed to decode.")
	}
	pkey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Decoded PEM block failed to parse.")
	}
	rsaKey, ok := pkey.(*rsa.PublicKey)
	if !ok {
		logger.Fatal().
			Msgf("got unexpected key type: %T", pkey)
	}

	// route apis and start http multiplexer
	r = echo.New()

	v0 := r.Group("/api/v0")
	v0AuthReq := v0.Group("", middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "RS256",
		SigningKey:    rsaKey,
	}))
	v0AuthReq.POST("/board", createBoard)
	v0.GET("/board/:id", getBoard)
	v0AuthReq.PUT("/board/:id", updateBoard)
	v0AuthReq.PATCH("/board/:id", editBoard)
	v0.GET("/board/track/:id", getBoardsFromTrack)
	v0.GET("/board/sha3/:hash", getBoardFromHash)
	v0AuthReq.DELETE("/board/:tid/:bid", deleteBoard)

	r.Start(":5002")
}
