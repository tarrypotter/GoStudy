// EX7-3

// 패키지 선언
package main

//패키지 가져오기
import "fmt"	

func F(n int)int{
 if n < 2 {
   return n
 }
 return F(n-2) + F(n-1) 
}
  
//함수를 선언하고 중괄호로 시작을 알림 
func EX7-3() {
  fmt.Println(F(9))
  }

