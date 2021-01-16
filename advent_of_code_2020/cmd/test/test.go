package main

import "fmt"

func main()  {
	testMapStringToStrings := map[string][]string{"a": []string{"1", "2", "3"}, "b": []string{"4", "5", "6"}}
	testMapIntToStrings := map[int][]string{1: []string{"1", "2", "3"}, 2: []string{"4", "5", "6"}}
	testMapIntToInts := map[int][]int{1: []int{1, 2, 3}, 2: []int{4, 5, 6}}
	testMapIntToInt := map[int]int{1: 2, 2: 3}

	fmt.Printf("%#v\n%#v\n%#v\n%#v\n", 
	testMapStringToStrings, testMapIntToStrings, testMapIntToInts, testMapIntToInt)
}