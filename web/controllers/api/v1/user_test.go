package v1

import (
	"log"
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

func TestUserRegister(t *testing.T) {
	utils.Post("/api/v1/user/register", url.Values{
		"username": {"zing123"},
		"password": {"123456"},
		"role":     {"admin"},
		"mail":     {"599490911@qq.com"},
		"phone":    {"18800563600"},
	}.Encode())

}

func TestUserDelete(t *testing.T) {
	utils.TUrl.RawQuery = url.Values{
		"id": {"19"},
	}.Encode()
	utils.Get("/api/v1/user/delete")
}

func TestUserUpdate(t *testing.T) {
	utils.Post("/api/v1/user/register", url.Values{
		"id":       {"16"},
		"username": {"zing123"},
		"password": {"123456"},
		"role":     {"admin"},
		"mail":     {"599490911@qq.com"},
		"phone":    {"18800563600"},
	}.Encode())
}

func TestCheckMail(t *testing.T) {
	log.Println(checkMail("1@qq.com"))
	log.Println(checkMail("1@@.com"))
	log.Println(checkMail("1@.com"))
}

func TestCheckPhone(t *testing.T) {
	log.Println(checkPhone("18800000000"))
	log.Println(checkPhone("188000000"))
	log.Println(checkPhone("18000000000"))
}
