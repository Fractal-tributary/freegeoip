package utils

import (
	"fmt"
	"net/http"
	"path/filepath"
	"testing"
	"time"
)
var MaxMindDB = "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=R9bpXwUb0nfuavnC&suffix=tar.gz"


func TestDeCompress(t *testing.T) {
	err := DeCompress("GeoLite2-City_20191231.tar.gz","")
	if err != nil{
		t.Fatal(err)
	}
}

func TestNewDeCompress(t *testing.T) {
	err :=NewDeCompress("GeoLite2-City_20191231.tar.gz")
	if err != nil{
		t.Fatal(err)
	}
}

func TestUrlDC(t *testing.T) {
	d,err :=UrlDC(MaxMindDB)
	if err != nil{
		t.Fatal(err)
	}
	t.Log(d)
}

func TestUTarGz(t *testing.T) {
	res, err := http.Get(MaxMindDB)
	if err !=nil{
		return
	}
	err = UnTarGz("",res.Body)
	if err !=nil{
		return
	}
}

func TestTarGz(t *testing.T) {
	f,err := FindFile("./","GeoLite2-City.mmdb")
	if err != nil{
		t.Fatal(err)
	}
	t.Log(f)
	tmpfile := filepath.Join(
		fmt.Sprintf("_freegeoip.%d.db.gz", time.Now().UnixNano()))
	err = TarGz(f,tmpfile,0)
	if err != nil{
		t.Fatal(err)
	}
}