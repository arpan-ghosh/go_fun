package main

import "fmt"

func var_dump(expression interface{} ) {
	fmt.Printf("(%v, %T)\n", expression, expression)
}

func var_dump_empty_interface(expression ...interface{} ) {
	fmt.Printf("(%v, %T)\n", expression, expression)
}

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
	
	var z, l int = 4, 8
	fmt.Println(z)
	fmt.Println(l)
	
	tasty := "ice cream"
	horrible := "marmite"

	_, _ = tasty, horrible
	fmt.Println(tasty)
	fmt.Println(horrible)
	
	var_dump(tasty)
	var_dump(k)

	var_dump_empty_interface(tasty);
	var_dump_empty_interface(k);
}