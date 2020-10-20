package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"net/http"
	"strings"
)

var Auth = NewAuth().Serve

const (
	admin      = "admin"
	uploader   = "uploader"
	downloader = "downloader"
)

type Authenticate struct {
	Admin      map[string]interface{}
	Downloader map[string]interface{}
	Uploader   map[string]interface{}
}

func NewAuth() *Authenticate {
	return &Authenticate{
		Admin:    map[string]interface{}{},
		Uploader: map[string]interface{}{},
		Downloader: map[string]interface{}{
			"resource": []string{
				"/resource-history/lists",
				"/resource/group-lists",
				"/resource/download",
			},
			"project": []string{
				"/project/lists",
				"/project/download",
			},
			"project-history": []string{
				"/project-history/lists",
			},
		},
	}
}

func (a *Authenticate) Serve(ctx context.Context) {
	currentRoute := strings.TrimPrefix(ctx.GetCurrentRoute().Path(), "/api/v1")
	token := ctx.Values().Get("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	ctx.Values().Set("user", claims["user"])
	switch claims["role"] {
	case admin:
		ctx.Next()
		return
	case uploader:
		if !strings.Contains(currentRoute, "user") {
			ctx.Next()
			return
		}
	case downloader:
		for _, v := range a.Downloader {
			for _, v2 := range v.([]string) {
				if v2 == currentRoute {
					ctx.Next()
					return
				}
			}
		}
	default:
	}
	ctx.StatusCode(http.StatusForbidden)
	_, _ = ctx.JSON(iris.Map{
		"success": false,
		"err_msg": "当前操作没有权限",
		"data":    []int{},
	})
	return
}
