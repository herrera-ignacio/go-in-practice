package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Processes struct {
	Total    int     `ini:"total"`
	Running  int     `ini:"running"`
	Sleeping int     `ini:"sleeping"`
	Threads  int     `ini:"threads"`
	Load     float32 `ini:"load"`
}

func main() {
	fmt.Println("Write a struct to output:")
	proc := &Processes{
		Total: 23,
		Running: 3,
		Sleeping: 20,
		Threads: 34,
		Load: 1.8,
	}

	// Marshals the struct into a []byte
	data, err := Marshal(proc)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	fmt.Println("Read the data back into a struct")
	// Creates a new Processes struct and unmarshals the data into it
	proc2 := &Processes{}
	if err := Unmarshal(data, proc2; err != nil {
		panic(err)
	})
	fmt.Printf("Struct: %#v", proc2)
}

// Utility function to read tags off of struct fields
func fieldName(field reflect.StructField) string {
	if t := field.Tag.Get("ini"); t != "" {
		return t
	}
	return field.Name
}

func Marshal(v interface{}) ([]byte, error) {
	var b bytes.Buffer
	// Gets a reflect.Value of the current interface. Dereferences pointers.
	val := reflect.Indirect(reflect.ValueOf(v))
	if val.Kind() != reflect.Struct {
		return []byte{}, errors.New("unmarshal can only take structs")
	}

	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// Gets the name from tagName
		name := fieldName(f)
		raw := val.Field(i).Interface()
		fmt.Fprintf(&b, "%s=%v\n", name, raw)
	}
}

func Unmarshal(data []byte, v interface{}) error {
	val := reflect.Indirect(reflect.ValueOf(v))
	t := val.Type()

	b := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(b)

	// From data, you use a scanner to read one line of INI data at a time
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.SplitN(line, "=", 2)
		if len(pair) < 2 {
			// Skip any malformed lines
			continue
		}
		// Passes the task of setting the value to setField
		setField(pair[0], pair[1], t, val)
	}
}

func setField(name, value string, t reflect.Type, v reflect.Value) {
	// Iterates through each field on the struct
	// looking for one whose name matches the INI field's name
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if name == fieldName(field) {
			var dest reflect.Value
			switch field.Type.Kind() {
			default:
				fmt.Printf("Kind %s not supported.\n", field.Type.Kind())
				continue // if you don't know about the kind, just skip the field.
			}
		case reflect.Int:
			ival, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("Could not convert %q to int: %s\n", value , err)
				continue
			}
			dest = reflect.ValueOf(ival)
		case reflect.Float64:
			fval, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Printf("Could not convert %q to float64: %s\n", value, err)
				continue
			}
			dest = reflect.ValueOf(fval)
		case reflect.String:
			dest = reflect.ValueOf(value)
		case reflect.Bool:
			bval, err := strconv.ParseBool(value)
			if err != nil {
				fmt.Printf("Could not convert %q to bool: %s\n", value, err)
				continue
			}
			dest = reflect.ValueOf(bval)
		}
		v.Field(i).Set(dest)
	}
}
