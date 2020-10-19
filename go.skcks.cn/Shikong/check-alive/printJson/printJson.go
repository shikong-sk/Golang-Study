package printJson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func Json2Map(jsonString string) map[string]interface{} {
	var jsonMap map[string]interface{}
	_ = json.Unmarshal([]byte(jsonString), &jsonMap)
	return jsonMap
}

type printJsonOption struct {
	deep int
}

func Print(jsonMap map[string]interface{}, opts ...printJsonOption) {
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
			Print(v.(map[string]interface{}), printJsonOption{deep + 1})
		} else if valueIsArr {
			fmt.Printf("\t => \t \n")
			for key, val := range v.([]interface{}) {
				fmt.Printf(prefixBlank + "\t    \t\t\t")
				fmt.Printf("%v \t => \t", key)
				JsonArray(val, printJsonOption{deep + 6})
			}
			fmt.Printf(prefixBlank + "\n")

		} else {
			fmt.Printf("\t => \t %v\n", v)
		}
	}
}

func JsonArray(value interface{}, opts ...printJsonOption) {
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
		Print(value.(map[string]interface{}), printJsonOption{deep + 1})
	} else if valueIsArr {
		JsonArray(value.([]interface{}), printJsonOption{deep + 1})
	} else {
		fmt.Printf("%v\n", value)
	}
}
