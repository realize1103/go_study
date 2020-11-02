## 1. if
조건에 맞으면 {} 블럭안의 내용을 실행.  
조건식을 ()로 둘러 싸지 않아도 된다.  
`반드시 조건 블럭 시작 브레이스를  if 문과 같은 라인에 두어야 한다.`  
else if, else 와 사용 가능.
``` go
if k == 1 {
    println("One")
} else if k == 2 {  //같은 라인
    println("Two")
} else {   //같은 라인
    println("Other")
}
```
조건식 사용하기 이전에 간단한 문장(optional Statement)를 사용할 수 있음.
``` go
if val := i * 2; val < max {
    println(val)
}
 
// 아래 처럼 사용하면 Scope 벗어나 에러
val++

```
## 2. Switch
case 3,4처럼 콤마를 써서 나열할 수 있음.
``` go
package main
 
func main() {
    var name string
    var category = 1
 
    switch category {
    case 1:
        name = "Paper Book"
    case 2:
        name = "eBook"
    case 3, 4:
        name = "Blog"
    default:
        name = "Other"
    }
    println(name)
     
    // Expression을 사용한 경우
    switch x := category << 2; x - 1 {
        //...
    }   
}
```
타 언어와의 차이점.
| 구분 | 의미 |
|---|:---:|
| switch 뒤에 expression이 없을 수 있음 | 다른 언어는 switch 키워드 뒤에 변수나 expression 반드시 두지만, Go는 이를 쓰지 않아도 된다. 이 경우 Go는 switch expression을 true로 생각하고 첫번째 case문으로 이동하여 검사한다 |	
|case문에 expression을 쓸 수 있음|다른 언어의 case문은 일반적으로 리터럴 값만을 갖지만, Go는 case문에 복잡한 expression을 가질 수 있다|
|No default fall through|다른 언어의 case문은 break를 쓰지 않는 한 다음 case로 이동하지만, Go는 다음 case로 가지 않는다|
|Type switch|다른 언어의 switch는 일반적으로 변수의 값을 기준으로 case로 분기하지만, Go는 그 변수의 Type에 따라 case로 분기할 수 있다|


``` go
//switch 뒤에 조건변수 혹은 Expression을 적지 않는 용법
func grade(score int) {
    switch {
    case score >= 90:
        println("A")
    case score >= 80:
        println("B")
    case score >= 70:
        println("C")
    case score >= 60:
        println("D")
    default:
        println("No Hope")
    }
}

//타입을 검사하는 타입 switch
switch v.(type) {
case int:
    println("int")
case bool:
    println("bool")
case string:
    println("string")
default:
    println("unknown")
}
```
break이 없어도 다음 조건을 실행하지 않는다.  
`fallthrough를 사용하면 타 언어처럼 다음 조건까지 진행시킨다.`
``` go
package main
 
import "fmt"
 
func main() {
    check(2)
}
 
func check(val int) {
    switch val {
    case 1:
        fmt.Println("1 이하")
        fallthrough
    case 2:
        fmt.Println("2 이하")
        fallthrough
    case 3:
        fmt.Println("3 이하")
        fallthrough
    default:
        fmt.Println("default 도달")
    }
}
```