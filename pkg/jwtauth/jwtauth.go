package jwtauth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	Authenticator func(c *gin.Context) (any, error)
	Timeout time.Duration
	SigningAlgorithm string

	PayloadFunc func(data any) jwt.MapClaims

	Authorizator func(interface{}, *gin.Context) bool

	TimeFunc func() time.Time

	Realm string 

	Key []byte

	// SendCookie bool

	LoginResponse func(*gin.Context, int, string, time.Time)

	Unauthorized func(c *gin.Context, code int, message string) 
}

var (
	IdentityKey = "indentity"

	NiceKey = "nice"

	RoleIdKey = "roleid"

	RoleKey = "rolekey"

	RoleNameKey = "rolename"
)

func New(m *AuthMiddleware) (*AuthMiddleware, error) {
	if err := m.Init(); err != nil {
		return nil, err
	}
	return m, nil
}

func (mw *AuthMiddleware) Init() error {
	mw.Timeout = time.Hour
	if mw.TimeFunc == nil {
		mw.TimeFunc = time.Now
	}
	if mw.SigningAlgorithm == "" {
		mw.SigningAlgorithm = "HS256"
	}
	if mw.LoginResponse == nil {
		mw.LoginResponse = func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"token": token, 
				"expire": expire.Format(time.RFC1123),
			})
		}
	}
	return nil
}

func (mw *AuthMiddleware) LoginHandler(c *gin.Context) {
	var (
		data interface{}
		err error
	)
	if mw.Authenticator == nil {
		mw.unauthorized(c, 400, "auth middleware lack of authenticator")
	}
	data, err = mw.Authenticator(c)
	if err != nil {
		mw.unauthorized(c, 401, err.Error())
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims) 
	if mw.PayloadFunc != nil {
		for key, value := range mw.PayloadFunc(data) { //登陆并不关系data layer的具体model，只关心如何生成token
			claims[key] = value
		}
	}
	expire := mw.TimeFunc().Add(mw.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = mw.TimeFunc().Unix()
	tokenString, err := token.SignedString(mw.Key)
	if err != nil {
		mw.unauthorized(c, http.StatusOK, err.Error())
		return
	}

	//todo: add cookie

	mw.LoginResponse(c, http.StatusOK, tokenString, expire)
}


func (mw *AuthMiddleware) unauthorized(c *gin.Context, code int, message string) {
	c.Header("WWW-Authenticate", "JWT realm=" + mw.Realm)
	//todo: abort context
	mw.Unauthorized(c, code, message)

}


