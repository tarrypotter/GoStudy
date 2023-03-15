//지역변수 선언하는 코드
//변수 범위

//패키지 선언
package main
//패키지 가져오기
import "fmt"
//함수를 선언하고 중괄호로 시작을 알림 

func float() { 

 var a float32 = 1234.567
 var b float32 = 5678.901
 var c float32 = a*b
 var d float32 = c*5
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}