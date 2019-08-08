package v1

import (
	"net/url"
	"simple-ims/utils"
	"testing"
)

func TestUserLogin(t *testing.T) {

	utils.TUrl.RawQuery = url.Values{
		"username": {"zing"},
		"password": {"123456"},
	}.Encode()
	utils.Get("/api/v1/user/login")

}

func TestUserLists(t *testing.T) {

	utils.Get("/api/v1/user/lists")

}
