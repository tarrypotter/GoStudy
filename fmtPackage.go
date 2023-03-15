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

// Print 함수는 문자열과 변수의 값을 연속으로 출력함. 
// Println 함수는 Print 함수와 같지만 출력 후 자동으로 개행함. 
// Printf 함수는 형식을 지정하여 출력함.
// %s는 문자열을, %d는 10진수 정수를, %f는 부동 소수점 숫자를 출력함.
// \n은 줄바꿈

//부가설명 
 // Print: 매개변수로 전달된 값을 화면에 출력하는 함수. 줄바꿈이 되지 않음.
 // ex) fmt.Print("Hello ", "world") // "Hello world" 출력

// Println: 매개변수로 전달된 값을 화면에 출력하는 함수. 마지막에 줄바꿈을 추가함.
// ex) fmt.Println("Hello", "world") // "Hello world" 출력 후 한 줄 아래로 커서 이동

// Printf: 문자열 포맷팅 기능을 지원하는 함수. 포맷 문자열을 지정하고 해당 포맷 문자열에 매칭되는 값을 전달하여 문자열을 생성하고 출력함.
// ex) fmt.Printf("%d + %d = %d", 1, 2, 3) // "1 + 2 = 3" 출력