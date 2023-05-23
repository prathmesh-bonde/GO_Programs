package main

import "fmt"

func main() {
	// var mymap map[string]string

	// mymap := map[string]string{
	// 	"key1": "value1",
	// 	"key2": "value2",
	// }

	mymap := make(map[int]string)
	// adding elements to map
	mymap[1] = "Hello"
	mymap[10] = "World"

	// deleting elements in map
	delete(mymap, 1)
	mymap[2] = "red"
	mymap[4] = "ww"
	mymap[8] = "kds"

	// fmt.Println(mymap)
	iterate(mymap)
}

func iterate(m map[int]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
