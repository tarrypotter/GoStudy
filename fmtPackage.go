//fmt패키지 3가지 출력용 함수 print,println,printf

//패키지 선언
package main
//패키지 가져오기
import "fmt"
//함수를 선언하고 중괄호로 시작을 알림 
func fmtPackage() {
  var a int = 999
  var b string = "비둘기야 밥묵자"
  var f float32 = 99.9
  
  fmt.Print(b, a, "\n")
  fmt.Println(b, a)
  fmt.Printf("%s %d\n", b, a)
  
  fmt.Print(b, f, "\n")
  fmt.Println(b, f)
  fmt.Printf("%s %f\n", b, f)
}

// Print 함수는 문자열과 변수의 값을 연속으로 출력합니다. 
// Println 함수는 Print 함수와 같지만 출력 후 자동으로 개행합니다. 
// Printf 함수는 형식을 지정하여 출력합니다.
// %s는 문자열을, %d는 10진수 정수를, %f는 부동 소수점 숫자를 출력합니다.
// \n은 줄바꿈