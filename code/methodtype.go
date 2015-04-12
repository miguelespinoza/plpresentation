// START OMIT
package main

import "fmt"

type Int int

func (i Int) add(a Int) Int {
	return a + i
}

func main() {
	i := Int(5)
	fmt.Println(i.add(5))
}

// END OMIT
