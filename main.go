//replit(사이트에서하는거랑,앱)에서 하는거라서 한파일만 run이 되는 바람에 이 파일은 항상 바뀔것

//패키지 선언
package main
//패키지 가져오기
import "fmt"
//함수를 선언하고 중괄호로 시작을 알림 
func main() {
  // 출력하는 함수
  fmt.Println("hi")
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