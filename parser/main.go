package main

import (
	"fmt"
	"io/ioutil"
	"os"

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

type LIARange string

type LIAField struct {
	T           LIAType `yaml:"type"`
	Constraints *[]interface{}
	Quantity    *LIARange
}

func main() {
	input, err := ioutil.ReadFile("examples/galaxy.yaml")
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
		fmt.Println("type", key, "struct {")

		for key, field := range model {
			switch field.T {
			case Number:
				if field.Constraints != nil {
					for _, c := range *field.Constraints {
						if val, ok := c.(string); ok {
							if LIAConstraint(val) == Integer {
								fmt.Println("\tint64", key)
							} else {
								fmt.Println("\tfloat64", key)
							}
						} else {
							fmt.Println(">>>>", c)
						}
					}
				}

			case Text:
				fmt.Println("\tstring", key)
			case Point:
				fmt.Println("\tPoint", key)
			case File:
				fmt.Println("\tFile", key)
			default:
				fmt.Fprintln(os.Stderr, "Unknown type", field.T)
			}
		}

		fmt.Println("}")
	}

}
