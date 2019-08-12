package v1

import (
	"net/url"
	"simple-ims/utils"
	"testing"
)

func TestResourceTypeAdd(t *testing.T) {

	utils.Post("/api/v1/resource-type/add", url.Values{
		"name": {"周界三"},
		"desc": {"周界三周界三"},
	}.Encode())

}

func TestResourceTypeDelete(t *testing.T) {

	utils.TUrl.RawQuery = url.Values{
		"id": {"17"},
	}.Encode()
	utils.Get("/api/v1/resource-type/delete")

}

func TestResourceTypeUpdate(t *testing.T) {

	utils.Post("/api/v1/resource-type/update", url.Values{
		"id":   {"14"},
		"name": {"周界1二"},
		"desc": {"周界一周界一周界一周界一"},
	}.Encode())

}

func TestResourceTypeLists(t *testing.T) {

	utils.Get("/api/v1/resource-type/lists")

}
