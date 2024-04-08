package main

import (
	"fmt"

	"github.com/Moji00f/celeritas"
)

func main() {

	result := celeritas.TestFunc(1, 2)
	result2 := celeritas.TestFunc2(2, 1)
	result3 := celeritas.TestFunc3(2, 2)
	result4 := celeritas.TestFunc4(20, 2)
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
