package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"net/http"
	"strings"
)

var Auth = NewAuth().Serve

type Authenticate struct {
	AdminRoute    map[string]interface{}
	OrdinaryRoute map[string]interface{}
}

func NewAuth() *Authenticate {
	return &Authenticate{
		AdminRoute: map[string]interface{}{
			"user":             "*",
			"resource":         "*",
			"resource-type":    "*",
			"resource-history": "*",
		},
		OrdinaryRoute: map[string]interface{}{
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
	case "admin":
		ctx.Next()
		return
	default:
		for _, v := range a.OrdinaryRoute {
			for _, v2 := range v.([]string) {
				if v2 == currentRoute {
					ctx.Next()
					return
				}
			}
		}
	}
	ctx.StatusCode(http.StatusForbidden)
	_, _ = ctx.JSON(iris.Map{
		"success": false,
		"err_msg": "当前操作没有权限",
		"data":    []int{},
	})
	return
}
