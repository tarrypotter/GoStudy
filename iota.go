//iota 열거값 

// 패키지 선언
package main

//패키지 가져오기
import "fmt"	

const (
 Blue int = iota  // 0
 Red int = iota   // 1
 Green int = iota // 2
)
const (
  C1 uint = iota + 1  // 1 = 0 + 1
  C2                  // 2 = 1 + 1
  C3                  // 3 = 2 + 1
)
const (
  BitFlag1 uint = 1 << iota  //1 = 1 << 0
  BitFlag2   // 2 = 1 << 1
  BitFlag3   // 4 = 1 << 2
  BitFlag4   // 8 = 1 << 3
)


//함수를 선언하고 중괄호로 시작을 알림 
func iota() {  
    fmt.Println(Blue)
    fmt.Println(Red)
    fmt.Println(Green)
    fmt.Println("ㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡ")
    fmt.Println(C1)
    fmt.Println(C2)
    fmt.Println(C3)
    fmt.Println("ㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡ")
    fmt.Println(BitFlag1)
    fmt.Println(BitFlag2)
    fmt.Println(BitFlag3)
    fmt.Println(BitFlag4)
  }

