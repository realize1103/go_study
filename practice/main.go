package main

import (
	"fmt"
	"strings"
)

//func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

func lenUpperCase(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
func lenUpperCase2(name string) (length int, upperName string) { //반환값을 변수로 생성할 수 있음. - Naked return
	defer fmt.Println("I'm done.") //함수가 끝났을때 실행되는 아이.
	length = len(name)
	upperName = strings.ToUpper(name)
	return
}

func supperAdd(numbers ...int) int {
	total := 0

	for number := range numbers {
		fmt.Println(number)
		total += number
	}
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return total
}

func canIDrink(age int) bool {
	//go에서는 스위치나 이프문에서 변수생성가능
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
	/*
		switch {
		case age < 18:
			return false
		case age >= 18:
			return true
		}
		return false
	*/

}

func repeateMe(words ...string) {
	fmt.Println(words)
}

//구조체
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	const name string = "kenneth"
	fmt.Println(multiply(2, 3))
	totalLength, upperName := lenUpperCase("kenneth") //한개이상 반환하는 함수
	fmt.Println(totalLength, upperName)
	//repeateMe("aaa", "bbbb", "cccc", "dddd")         //무제한 입력받기
	//total := supperAdd(2, 4, 3, 2, 3, 4, 5, 2, 3, 9) // for문
	//fmt.Println(total)
	//fmt.Println(canIDrink(16)) // if문

	//pointers!
	a := 2
	b := &a
	//&a 는 메모리의 주소를 나타냄
	//*a 는 살펴보거나 훑어본다.

	fmt.Println(&a, &b, a, b, *b)
	*b = 20 //b를 수정했으나 a까지 수정됨. 이게 포인터임. 값을 복사한게 아니고 a의 메모리 주소를 복사했기 때문.
	fmt.Println(a)

	//Array
	//Way 1
	names := [5]string{"kwanwoo", "damin", "cy"}
	names[3] = "fff"
	names[4] = "sssss"
	fmt.Println(names)
	//Way 2 - slice
	names2 := []string{"kwanwoo", "damin", "cy"}
	names2 = append(names2, "bbb")
	fmt.Println(names2)

	//Map
	//nico := map[string]string{"name": "nico", "age": "12"}

	favFood := []string{"kimchi", "ramen"}
	nico := person{"kwanwoo", 18, favFood}
	nico2 := person{name: "kenneth", age: 18, favFood: favFood} // 둘다 같은 방법이지만 가독성은 이게 훨씬 좋다.
	fmt.Println(nico.name)
	fmt.Println(nico2.name)
}
