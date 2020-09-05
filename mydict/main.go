package main

import (
	"fmt"
	"learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}

	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dictionary)
	err1 := dictionary.Update(word, "bbbb")
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Println(dictionary)
	dictionary.Delete(word)
	fmt.Println(dictionary)

	/*

		searchedDefinition, err := dictionary.Search("second")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(searchedDefinition)
		}*/

}
