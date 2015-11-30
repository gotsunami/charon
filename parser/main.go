package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type LIA struct {
	Models map[string]map[string]LIAField
	Admin  interface{}
}

type LIAType string

const (
	Number LIAType = "number"
	Text   LIAType = "text"
	Point  LIAType = "point"
	File   LIAType = "file"
)

type LIAConstraint string

const (
	Integer  LIAConstraint = "integer"
	Float    LIAConstraint = "float"
	Positive LIAConstraint = "positive"
	Negative LIAConstraint = "negative"
	NotNull  LIAConstraint = "not null"
	NotEmpty LIAConstraint = "not empty"
)

type LIASimpleRange string

type LIARange []LIASimpleRange

type LIAField struct {
	T           LIAType `yaml:"type"`
	Constraints *[]interface{}
	Quantity    *LIARange
}

func parseNumberConstraints(constraints *[]interface{}, arr []string) []string {
	if constraints != nil {
		for _, c := range *constraints {
			if val, ok := c.(string); ok {
				if LIAConstraint(val) == Integer {
					arr[0] = "\tint64\t"
				} else {
					arr[0] = "\tfloat64\t"
				}
			} else {
				fmt.Println(c.(map[interface{}]interface{}))
				fmt.Println(">>>>>", c)
			}
		}
	}
	return arr
}

func parseQuantity(quantity *LIARange, arr []string) []string {
	// default quantity == 1
	if quantity == nil {
		return arr
	}
	// quantity: n, create an array of fixed size 'n'
	if uval, err := strconv.ParseUint(string(*quantity), 10, 0); err == nil {
		if uval > 1 {
			arr = append(arr, fmt.Sprintf("[%d]", uval))
			return arr
		}
	}
	// an 'or'
	ranges := strings.Split(string(*quantity[0]), " or ")
	if len(ranges) > 1 {
		if ranges[0] == "0" && ranges[1] == "1" {
			arr = append(arr, "*")
		}
		if ranges[1] != "1" {
			return append(arr, "[]")
		}
	}
	// a 'to'
	ranges = strings.Split(string(*quantity[0]), " to ")
	if len(ranges) > 1 {
		if ranges[0] == "0" {
			arr = append(arr, "*")
		}
		if ranges[1] != "1" {
			return append(arr, "[]")
		}
	}
	// ?
	return arr
}

func main() {
	input, err := ioutil.ReadFile("../examples/galaxy.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	lia := new(LIA)

	if err := yaml.Unmarshal(input, lia); err != nil {
		fmt.Println(err)
		return
	}

	for key, model := range lia.Models {
		modelname := fmt.Sprintf("%s%s", strings.ToUpper(string(key[0])), key[1:])

		fmt.Println("type", modelname, "struct {")

		for key, field := range model {
			fieldname := fmt.Sprintf("%s%s", strings.ToUpper(string(key[0])), key[1:])

			arr := make([]string, 1)
			switch field.T {
			case Number:
				arr = parseNumberConstraints(field.Constraints, arr)
			case Text:
				arr[0] = "\tstring\t"
			case Point:
				arr[0] = "\tPoint\t"
			case File:
				arr[0] = "\tFile\t"
			default:
				if lia.Models[string(field.T)] != nil {
					modelname := fmt.Sprintf("%s%s", strings.ToUpper(string(field.T[0])), field.T[1:])
					arr[0] = fmt.Sprintf("\t%s\t", modelname)
				} else {
					fmt.Fprintln(os.Stderr, "Unknown type", field.T)
				}
			}

			arr = parseQuantity(field.Quantity, arr)

			arr = append(arr, fieldname)
			fmt.Println(strings.Join(arr, ""))

		}

		fmt.Println("}\n")
	}

}
