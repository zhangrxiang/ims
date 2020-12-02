package models

import (
	"log"
	"simple-ims/utils"
	"testing"
)

func init() {
	GetDBInstance()
}
func TestResourceAll(t *testing.T) {

	model := ResourceModel{}
	all, err := model.All()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(all)
}

func TestResourceFind(t *testing.T) {
	alice := utils.StrToIntSlice("1,2,3,4,5", ",")
	log.Println(alice)
	//go func(alice []int) {
	//	for _, id := range alice {
	//		model := &ResourceModel{}
	//		model.Id = id
	//		log.Println(id)
	//		resourceModel, err := model.Find()
	//		log.Println(resourceModel)
	//		if err != nil {
	//			log.Fatalln(err)
	//		}
	//		log.Println(resourceModel.Path)
	//	}
	//}(alice)
	//time.Sleep(3*time.Second)

	//<-time.AfterFunc(time.Second, func() {
	//	alice := utils.StrToIntSlice("1,2,3,4,5", ",")
	//	for _, id := range alice {
	//		model := &ResourceModel{}
	//		model.Id = id
	//		log.Println("id",id)
	//		resourceModel, err := model.Find()
	//		log.Println(resourceModel)
	//		if err != nil {
	//			log.Fatalln(err)
	//		}
	//		log.Println(resourceModel.Path)
	//	}
	//}).C
	//time.Sleep(time.Second)

	go find(alice)
	//time.Sleep(10*time.Second)
}
func find(alice []int) {
	for _, id := range alice {
		model := &ResourceModel{}
		model.ID = id
		resourceModel, err := model.FindBy()
		log.Println(resourceModel)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
