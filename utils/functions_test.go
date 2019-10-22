package utils

import (
	"log"
	"testing"
)

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
