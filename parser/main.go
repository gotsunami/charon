package main

import (
	"fmt"
	"strings"
)

type DatabaseScheme struct {
	Models []Model

	Result []string
}

type Model struct {
	Name   string
	Fields []Field

	Result []string
}

type Field struct {
	Name        string
	Type        string
	Quantity    []Range
	Constraints []Constraint

	Result []string
}

type Range struct {
	Min string
	Max string
}

type Constraint interface {
	Parse(constraint interface{}) error
	ToString() (string, error)
}

func (db *DatabaseScheme) ToString() (string, error) {
	db.Result = make([]string, 0)
	for _, model := range db.Models {
		str, err := model.ToString()
		if err != nil {
			return "", err
		}
		db.Result = append(db.Result, str)
	}
	return strings.Join(db.Result, "\n"), nil
}

func (m *Model) ToString() (string, error) {
	m.Result = []string{fmt.Sprintf("type %s struct {", m.Name)}
	for _, field := range m.Fields {
		str, err := field.ToString()
		if err != nil {
			return "", err
		}
		m.Result = append(m.Result, str)
	}
	return strings.Join(m.Result, "\n"), nil
}

func (f *Field) ToString() (string, error) {
	f.Result = []string{fmt.Sprintf("%s %s %s", f.Name, f.Quantity, f.Type)}
	for _, c := range f.Constraints {
		str, err := c.ToString()
		if err != nil {
			return "", err
		}
		f.Result = append(f.Result, str)
	}
	return strings.Join(f.Result, "\n"), nil
}

// Actually parse the file
func main() {
	db, err := parse()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(db.ToString())
	}
}
