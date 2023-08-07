package handler

import (
	"errors"
	"ferry-learn/middleware"
	"ferry-learn/models/system"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	// "github.com/mssola/user_agent"
)

func Authenticator(c *gin.Context) (any, error) {
	var (
		loginVal system.Login
	)
	// ua := user_agent.New(c.Request.UserAgent())
	//todo: loginlog
	if err := c.ShouldBind(&loginVal); err != nil {
		log.Println("parse login request fail: ", err)
	}
	user, role, err := loginVal.GetUser()
	if err == nil {
		if user.Status == 1 {
			return nil, errors.New("the user is banned")
		}
		return map[string]any{"user": user, "role": role}, nil
	} else {
		log.Println("[Authenicator]login fail: ", err.Error())
	}
	return nil, errors.New("incorrect Username or Password")
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(system.SysUser)
		r, _ := v["role"].(system.SysRole)
		return jwt.MapClaims{
			middleware.IdentityKey: u.UserId,
			middleware.RoleIdKey:   r.RoleId,
			middleware.RoleKey:     r.RoleKey,
			middleware.NiceKey:     u.Username,
			middleware.RoleNameKey: r.RoleName,
		}

	}
	return jwt.MapClaims{}
}
