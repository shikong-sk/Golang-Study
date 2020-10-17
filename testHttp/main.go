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
	//url := "http://httpbin.org/json"
	url := "https://m.weibo.cn/api/config/list"
	res, err := http.Get(url)
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
	printJson(jsonMap)
}

type printJsonOption struct {
	deep int
}

func printJson(jsonMap map[string]interface{}, opts ...printJsonOption) {
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
			printJson(v.(map[string]interface{}), printJsonOption{deep + 1})
		} else if valueIsArr {
			fmt.Printf("\t => \t [\n")
			for key, val := range v.([]interface{}) {
				fmt.Printf(prefixBlank + "\t    \t\t\t")
				fmt.Printf("%v \t => \t", key)
				printJsonArray(val, printJsonOption{deep + 6})
			}
			fmt.Printf(prefixBlank + "\t    \t\t ]\n")

		} else {
			fmt.Printf("\t => \t %v\n", v)
		}
	}
}

func printJsonArray(value interface{}, opts ...printJsonOption) {
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
		printJson(value.(map[string]interface{}), printJsonOption{deep + 1})
	} else if valueIsArr {
		printJsonArray(value.([]interface{}), printJsonOption{deep + 1})
	} else {
		fmt.Printf("%v\n", value)
	}
}
