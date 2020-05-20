package routing

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
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyNbZz3Ig6VWUTxsBt5d4
Co+9VKIHm4BvQjG4ynh2v3a5an+gE7V6wY5ExBvIPNqOJnJWnvvEk22wYPB3to1T
6KMlpTmWmuO9aqBaLBwDY42UctS30B18bcOpz8wZy5gL1BkheTExfg09yOj0igW1
gMNyVCVYuhh5ye8NAinMCNxc9QgLz6ODxGXIfVlNN96C0iGhxAto7x9cMYTaT2FS
9GN6ZuOlbV4RnlmaiI+avbga6sy4m0WEiRFcx5Je7GZhsmtuQ65PaeUiOM/MpWNB
doBgwAWghhHc4WSTqyGbsVgl82qHvV+7Z9MmGq1k9fUk5zNtnP7Ou+gv2FBEMu9p
QQIDAQAB
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
	v0.GET("/board/hash/:hash", getBoardFromHash)
	v0AuthReq.DELETE("/board/:tid/:bid", deleteBoard)

	r.Start(":5002")
}
