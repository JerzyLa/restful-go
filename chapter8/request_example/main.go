package main

import (
	"github.com/levigross/grequests"
	"log"
)

func main() {
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	// You can modify the request by passing an optional
	// RequestOptions struct
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	log.Println(resp.String())

	//	resp, err := grequests.Get("http://httpbin.org/get", nil)
	//  // You can modify the request by passing an optional
	//  // RequestOptions struct
	//  if err != nil {
	//    log.Fatalln("Unable to make request: ", err)
	//  }
	//  var returnData map[string]interface{}
	//  resp.JSON(&returnData)
	//  log.Println(returnData)
}
