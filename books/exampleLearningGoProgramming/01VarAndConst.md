## 1. 변수
a라는 정수(int)변수를 선언
``` go
var a int 
```

변수 선언시 초기값 할당. float32 변수 f에 11.0 할당.
``` go
var f float32 = 11.
```
선언이 후 값 변경
``` go
a = 10
f = 12.0
```

`* 선언된 변수가 사용되지 않으면 Go는 에러를 발생시킨다.`

동일한 타입의 변수가 복수개일 때
``` go
var i, j, k int
```
동일한 타입의 변수가 복수개이고 값을 할당할때.
``` go
var i, j, k int = 1, 2, 3
```

초기값을 지정하지 않으면 Zero Value가 기본으로 할당된다.  
` 숫자형은 0, bool은 false, String은 ""`

변수를 선언하는 다른 방식으로 Short Assignment Statement(:=)가 있다.  
var i = 1 대신 i := 1로 var를 생략하고 사용할수 있다.  
`하지만 함수(func) 내에서만 사용할 수 있으며, 함수 밖에서는 var를 사용해야한다.`

## 2. 상수

상수는 const로 선언.
``` go
const c int = 0
const s string = "Hi"
```

Go에서는 할당되는 값을 보고 그 타입을 추론하는 기능이 자주 사용됨. 아래처럼 생략가능
``` go
const c = 0
const s = "Hi"
```
여러 개의 상수들 묶어서 지정할 수 있는데, 아래와 같이 괄호 안에 상수들을 나열하여 사용할 수 있다.
``` go
const (
    Visa = "Visa"
    Master = "MasterCard"
    Amex = "American Express"
)
```

한가지 유용한 팁으로 상수값을 0부터 순차적으로 부여하기 위해 iota 라는 identifier를 사용할 수 있다. 이 경우 iota가 지정된 Apple에는 0이 할당되고, 나머지 상수들을 순서대로 1씩 증가된 값을 부여받는다.
``` go
const (
    Apple = iota // 0
    Grape        // 1
    Orange       // 2
``` 

## 3. Go 키워드
25개의 예약 키워드가 있다.
``` go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```