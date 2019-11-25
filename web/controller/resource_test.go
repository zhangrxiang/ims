package controller

import (
	"github.com/elliotchance/pie/pie"
	"log"
	"testing"
)

func TestName(t *testing.T) {

	ints := pie.Ints{1, 2, 3, 4, 5, 6}
	ints2 := pie.Ints{11, 12, 13, 14, 15, 16}

	added, removed := ints.Diff(ints2)
	log.Println(added, removed)
	if len(removed) != len(ints) || len(added) != len(ints2) {
		log.Println("存在")
	} else {
		log.Println("不存在")
	}
	ids := pie.Ints{33}
	rhids := pie.Strings{"1", "2", "3", "4", "5"}
	added, removed = ids.Diff(rhids.Ints())
	log.Println(added, removed)
	if len(removed) != len(ids) || len(added) != len(rhids) {
		log.Println("存在")
	} else {
		log.Println("不存在")
	}

}
