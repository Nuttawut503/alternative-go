package main

import (
	bp "app/main/design_patterns/behavioral_patterns"
	"fmt"
	"time"
)

var a int = 15

func main() {
	bp.CommandExample()

	var employees []string
	employees = append(employees, "ABC")
	employees = append(employees, "DEF")
	for _, employee := range employees {
		fmt.Print(employee, " ")
	}

	ages := map[string]int{}
	ages["john"] = 3

	if _, found := ages["John"]; found {
		// a key name "John" exists
		println("found")
	} else {
		// no key name "John" there
		println("noo")
	}

	item := "Chocolate"
	price := 15.33333

	line := fmt.Sprintf("%s (%.2f)", item, price)

	println(line)

	fmt.Println(time.Now().Unix())
}
