package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// LIA stands for Langage Intermediaire d'Abstraction or Intermediate Abstraction Language

// File structure
type LIA struct {
	Models map[string]map[string]LIAField
	Admin  interface{}
}

// possible types for fields
type LIAType string

const (
	Number LIAType = "number"
	Text   LIAType = "text"
	Point  LIAType = "point"
	File   LIAType = "file"
)

// possible values for constraints
type LIAConstraint string

const (
	Integer  LIAConstraint = "integer"
	Float    LIAConstraint = "float"
	Positive LIAConstraint = "positive"
	Negative LIAConstraint = "negative"
	NotNull  LIAConstraint = "not null"
	NotEmpty LIAConstraint = "not empty"
)

// Unique range
type LIASimpleRange string

// Liste of range
type LIARange LIASimpleRange

// Typical LIA field
type LIAField struct {
	T           LIAType `yaml:"type"`
	Constraints *[]interface{}
	Quantity    *LIARange
}

// Parse a LIA range and appends it to an array of strings to
// to be used to build the DatabaseScheme
func parseQuantity(field *LIAField, f *Field) error {
	quantity := field.Quantity

	// default quantity == 1
	if quantity == nil {
		f.Quantity = []Range{Range{Min: 1, Max: 1}}
		return nil
	}
	// quantity: n, create an array of fixed size 'n'
	if uval, err := strconv.ParseUint(string(*quantity), 10, 0); err == nil {
		f.Quantity = []Range{Range{Min: float64(uval), Max: float64(uval)}}
		return nil
	}
	// 	// an 'or'
	ranges := strings.Split(string(*quantity), " or ")
	if len(ranges) > 1 {
		minval, err := strconv.ParseUint(ranges[0], 10, 0)
		if err != nil {
			return err
		}
		maxval, err := strconv.ParseUint(ranges[1], 10, 0)
		if err != nil {
			return err
		}
		f.Quantity = []Range{Range{Min: float64(minval), Max: float64(maxval)}}
	}
	// 	// a 'to'
	// 	ranges = strings.Split(string(*quantity), " to ")
	// 	if len(ranges) > 1 {
	// 		if ranges[0] == "0" {
	// 			arr = append(arr, "*")
	// 		}
	// 		if ranges[1] != "1" {
	// 			return append(arr, "[]")
	// 		}
	// 	}
	// 	// ?
	return nil
}

func parseNumberConstraints(field *LIAField, f *Field) error {
	if field.Constraints != nil {
		for _, c := range *field.Constraints {
			if val, ok := c.(string); ok {
				if LIAConstraint(val) == Integer {
					f.Type = "int64"
				} else {
					f.Type = "float64"
				}
			} else {
				return errors.New(fmt.Sprintln("Invalid field constraint: ", c.(map[interface{}]interface{})))
			}
		}
	}
	return nil
}

func parseTextConstraints(field *LIAField, f *Field) error {
	f.Type = "string"
	return nil
}

func parsePointConstraints(field *LIAField, f *Field) error {
	f.Type = "Point"
	return nil
}

func parseFileConstraints(field *LIAField, f *Field) error {
	f.Type = "File"
	return nil
}

func parse() (*DatabaseScheme, error) {
	input, err := ioutil.ReadFile("../examples/galaxy.yaml")
	if err != nil {
		return nil, err
	}

	lia := new(LIA)
	db := &DatabaseScheme{Models: make([]Model, 0)}

	if err := yaml.Unmarshal(input, lia); err != nil {
		return nil, err
	}

	for key, model := range lia.Models {
		m := Model{
			Name:   capitalize(key),
			Fields: make([]Field, 0)}

		for key, field := range model {
			f := &Field{
				Name:        capitalize(key),
				Quantity:    make([]Range, 0),
				Constraints: make([]Constraint, 0)}

			switch field.T {
			case Number:
				parseNumberConstraints(&field, f)
			case Text:
				parseTextConstraints(&field, f)
			case Point:
				parsePointConstraints(&field, f)
			case File:
				parseFileConstraints(&field, f)
			default:
				if lia.Models[string(field.T)] != nil {
					f.Type = capitalize(string(field.T))
				} else {
					return nil, errors.New(fmt.Sprintf("Unknown type %v", field.T))
				}
			}

			parseQuantity(&field, f)

			m.Fields = append(m.Fields, *f)
		}
		db.Models = append(db.Models, m)
	}

	return db, nil
}
