package middleware

import (
	"ferry-learn/handler"
	jwta "ferry-learn/pkg/jwtauth"
	"time"
)

func AuthInit() (*jwta.AuthMiddleware, error) {
	return jwta.New(&jwta.AuthMiddleware{
		Realm: "test zone",
		Timeout: time.Hour,
		PayloadFunc: handler.PayloadFunc,
		Authenticator: handler.Authenticator,
		Authorizator: handler.Authorizator,
		Unauthorized: handler.Unauthorized,
		TimeFunc: time.Now,
		Key: []byte("secret"),
	})
}