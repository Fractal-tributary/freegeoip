package utils

import (
	"testing"
)


func TestFindFile(t *testing.T) {
	f,err := FindFile("./","GeoLite2-City.mmdb")
	if err != nil{
		t.Fatal(err)
	}
	t.Log(f)
}