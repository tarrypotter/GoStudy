//변수 선언 하는 소스

//패키지 선언
package main
//패키지 가져오기
import "fmt"
//함수를 선언하고 중괄호로 시작을 알림 
func var01() {
  // go변수선언
  // var 변수명 타입 = 초기값 
  var a int =15
  var b int =9
  fmt.Println(a,b)
  fmt.Println(a+b)
  fmt.Println(a-b)
  fmt.Println(a*b)
  fmt.Println(a/b)
  fmt.Println(a%b)
}