package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	filename := os.Args[1]
	var config interface{}
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	result := checkpasswds(config)
	for k, v := range result {
		fmt.Println(k+":", v)
	}
}

func checkpasswds(in interface{}) map[string]string {
	out := make(map[string]string)
	inval := reflect.ValueOf(in)
	switch inval.Kind() {
	case reflect.Map:
		for k, v := range in.(map[interface{}]interface{}) {
			vval := reflect.ValueOf(v)
			switch vval.Kind() {
			case reflect.String:
				if v.(string) != "" && strings.Contains(strings.ToLower(k.(string)), "password") {
					out[k.(string)] = v.(string)
				}
			case reflect.Map:
				for k1, v1 := range checkpasswds(v) {
					out[k.(string)+"."+k1] = v1
				}
			}
		}
	}
	return out
}
