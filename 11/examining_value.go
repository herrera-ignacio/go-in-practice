package main

/**
Types to examine
*/
type MyInt int

type Person struct {
	Name    *Name
	Address *Address
}

type Name struct {
	Title, First, Last string
}

type Address struct {
	Street, Region string
}

func main() {
	fmt.Println("Walking a simple integer")
	var one MyInt = 1
	walk(one, 0)

	fmt.Println("Walking a simple struct")
	two := struct{ Name string }{"foo"}
	walk(two, 0)

	fmt.Println("Walking a struct with struct fields")
	p := &Person{
		Name: &Name{"Count", "Tyrone", "Rugen"},
		Address: &Address{"Humperdink Castle", "Florian"}
	}
	walk(p, 0)
}

func walk(u interface{}, depth int) {
	// For your unknown value u, you get the reflect.Value.
	// If it's a pointer, you dereference the pointer.
	val := reflect.Indirect(reflect.ValueOf(u))

	// Depth helps you do some tab identing for prettier output.
	tabs := strings.Repeat("\t", depth+1)

	t := val.Type()
	fmt.Printf("%sValue is type %q (%s)\n", tabs, t, val.Kind())

	// If the kind is struct, you examine its fields
	if val.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldVal := reflect.Indirect(val.Field(i))
			tabs := strings.Repeat("\t", depth+2)
			fmt.Printf("%sField %q is a type %q (%s)\n", tabs, field.Name, field.Type, fieldVal.Kind())

			if fieldVal.Kind == reflect.Struct {
				walk(fieldVal.Interface(), depth+1)
			}
		}
	}
}
