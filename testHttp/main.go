package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func main() {

	res, err := http.Get("http://httpbin.org/json")
	if err != nil {
		fmt.Println(err)
		recover()
	}
	fmt.Println(res)

	baseBody, _ := ioutil.ReadAll(res.Body)
	body := string(baseBody)
	fmt.Println(body)

	var jsonMap map[string]interface{}
	_ = json.Unmarshal(baseBody, &jsonMap)

	fmt.Println(jsonMap)
	fmt.Printf("%T\n\n", jsonMap)
	forEachJson(jsonMap)
}

type forEachJsonOption struct {
	deep int
}

func forEachJson(jsonMap map[string]interface{}, opts ...forEachJsonOption) {
	var deep int
	for _, v := range opts {
		deep = v.deep
	}
	var prefixBlank string
	for i := 0; i < deep; i++ {
		prefixBlank += "\t"
	}
	for k, v := range jsonMap {
		fmt.Printf(prefixBlank)
		fmt.Printf(k)
		valueType := reflect.TypeOf(v)
		valueIsMap := strings.EqualFold("map[string]interface {}", valueType.String())
		valueIsArr := strings.EqualFold("[]interface {}", valueType.String())

		if valueIsMap {
			fmt.Println()
			forEachJson(v.(map[string]interface{}), forEachJsonOption{deep + 1})
		} else if valueIsArr {
			fmt.Printf("\t => \t [\n")
			for key, val := range v.([]interface{}) {
				fmt.Printf(prefixBlank + "\t    \t\t\t")
				fmt.Printf("%v \t => \t", key)
				parseArray2Json(val, forEachJsonOption{deep + 6})
			}
			fmt.Printf(prefixBlank + "\t    \t\t ]\n")

		} else {
			fmt.Printf("\t => \t %v\n", v)
		}
	}
}

func parseArray2Json(value interface{}, opts ...forEachJsonOption) {
	var deep int
	for _, v := range opts {
		deep = v.deep
	}
	var prefixBlank string
	for i := 0; i < deep; i++ {
		prefixBlank += "\t"
	}
	valueType := reflect.TypeOf(value)
	valueIsMap := strings.EqualFold("map[string]interface {}", valueType.String())
	valueIsArr := strings.EqualFold("[]interface {}", valueType.String())
	if valueIsMap {
		fmt.Println()
		forEachJson(value.(map[string]interface{}), forEachJsonOption{deep + 1})
	} else if valueIsArr {
		parseArray2Json(value.([]interface{}), forEachJsonOption{deep + 1})
	} else {
		fmt.Printf("%v\n", value)
	}
}
