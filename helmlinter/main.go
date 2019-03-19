package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var (
	globalUsage = `Helm plugin to find hardcoded passwords in values.yaml files

Examples:
  $ helm linter mychart/
`
)

func main() {
	if len(os.Args) > 1 {
		chartname := os.Args[1]
		if chartname != "" {
			fmt.Println("Checking chart " + chartname)
			fmt.Println("")
			chartname = chartname + "/values.yaml"
			if _, err := os.Stat(chartname); os.IsNotExist(err) {
				fmt.Println(chartname + " does not exist!")
				os.Exit(0)
			}
			var config interface{}
			source, err := ioutil.ReadFile(chartname)
			if err != nil {
				panic(err)
			}
			err = yaml.Unmarshal(source, &config)
			if err != nil {
				panic(err)
			}
			result := checkpasswds(config)
			if len(result) > 1 {
				// fmt.Println("")
				fmt.Println("Found some hard-coded password/s:")
				for k, v := range result {
					fmt.Println(" "+k+":", v)
				}
			} else {
				// fmt.Println("")
				fmt.Println("Hard-coded password/s have not been found :-)")
			}
		}
	} else {
		fmt.Println(globalUsage)
		os.Exit(0)
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
