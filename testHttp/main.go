package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {

}

func (r *response) parseJson(){

}

func main() {

	res, err := http.Get("http://192.168.1.241:1080/ncda/v2/api-docs")
	if err != nil {
		fmt.Println(err)
		recover()
	}
	fmt.Println(res)

	baseBody , _ := ioutil.ReadAll(res.Body)
	body := string(baseBody)
	fmt.Println(body)
}
