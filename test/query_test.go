package test

import (
	"awdb-golang/awdb-golang"
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

func TestAvg(t *testing.T) {

	//db, err := awdb.Open("C:/Users/admin/Desktop/allawdb/IP_city_single_WGS84_awdb.awdb")
	db, err := awdb.Open("/Users/pandy/Downloads/IP_basic_single_WGS84_awdb/IP_basic_single_WGS84.awdb")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	dura := time.Duration(0)
	for i := 0; i < 1000000; i++ {
		currentNSecond := time.Now()
		ip := net.ParseIP("166.111.4.100")

		var record interface{}
		err = db.Lookup(ip, &record)
		if err != nil {
			log.Fatal(err)
		}
		//var resMap = record.(map[string]interface{})
		//fmt.Printf("accuracy:%s", resMap["accuracy"])
		//fmt.Println()
		//fmt.Printf("%s\n", record)
		finishedNTime := time.Now()
		dura += finishedNTime.Sub(currentNSecond)
	}

	fmt.Printf((dura / 1000000).String())

}
