## 1. for
Go의 반복문음 for만 있다.  
타언어와 마찬가지로 초기, 조건식, 증감식을 사용하고 ()를 쓰면 에러가 난다.
```go
package main
 
func main() {
    //1. 기본
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    println(sum)

    //2. 조건식만 쓰는 for 루프
    n := 1
    for n < 100 {
        n *= 2      
        //if n > 90 {
        //   break 
        //}     
    }
    println(n)

    //3. for range
    names := []string{"홍길동", "이순신", "강감찬"} 
    for index, name := range names {
        println(index, name)
    }

    //4. break, continue, goto 문

    var a = 1
    for a < 15 {
        if a == 5 {
            a += a
            continue // for루프 시작으로
        }
        a++
        if a > 10 {
            break  //루프 빠져나옴
        }
    }
    if a == 11 {
        goto END //goto 사용예
    }
    println(a)
 
END:
    println("End")

    //5. 무한루프 - ctrl+C로 빠져나온다.
    for {
        println("Infinite loop")        
    }
}
```