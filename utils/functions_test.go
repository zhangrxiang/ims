package utils

import (
	"fmt"
	"github.com/elliotchance/pie/pie"
	"log"
	"sort"
	"strings"
	"testing"
)

func TestName4(t *testing.T) {
	var a pie.Strings
	a = []string{"1,2", "2,3", "3,4", "11,12"}
	log.Println(a)
	a = strings.Split(a.Join(","), ",")
	log.Println(a.Unique().Ints().Sort())
	var b pie.Ints
	b = []int{1, 2, 3, 4, 5}
	b = []int{13, 14}
	log.Println(b.Diff(a.Unique().Ints().Sort()))
}

func TestName3(t *testing.T) {
	name := pie.Strings{"Bob", "Sally", "John", "Jane"}.
		FilterNot(func(name string) bool {
			return strings.HasPrefix(name, "J")
		}).
		Map(strings.ToUpper).
		Last()

	fmt.Println(name) // "SALLY"
}
func TestVersionCompare(t *testing.T) {
	log.Println(VersionCompare("0.0.1", "0.0.2"))
	log.Println(VersionCompare("0.0.2", "0.0.1"))
	log.Println(VersionCompare("0.1.2", "0.2.1"))
	log.Println(VersionCompare("0.2.2", "0.1.1"))
	log.Println(VersionCompare("2.2.2", "2.1.1"))
	log.Println(VersionCompare("2.2.2", "3.1.1"))
	log.Println(VersionCompare("2.0.2", "2.1.1"))
	log.Println(VersionCompare("2.0.2", "2.0.2"))
}

func TestStrToIntAlice(t *testing.T) {
	log.Println(StrToIntSlice(",1,1,", ","))
	log.Println(StrToIntSlice("1,1", ","))
}

func TestName(t *testing.T) {
	str := []string{"1,2", "3,4", "4,5"}
	log.Println(strings.Join(str, ","))
	log.Println(strings.Split(strings.Join(str, ","), ","))

}
func TestName2(t *testing.T) {
	b := []string{"a", "b", "c", "c", "e", "f", "a", "g", "b", "b", "c"}
	sort.Strings(b)
	fmt.Println(Duplicate(b))

	c := []int{1, 1, 2, 4, 6, 7, 8, 4, 3, 2, 5, 6, 6, 8}
	sort.Ints(c)
	fmt.Println(Duplicate(c))
	log.Println(ElementExists(Duplicate(c), []interface{}{1}))
	log.Println(ElementExists(Duplicate(c), []interface{}{8, 6}))
	log.Println(ElementExists(Duplicate(c), []interface{}{16}))
	log.Println(ElementExists(Duplicate(c), []interface{}{"1"}))
	log.Println(ElementExists(Duplicate(c), Duplicate(c)))
	log.Println(ElementExists(Duplicate(c), Duplicate([]string{"1"})))
}
