// EX7-2

//패키지 선언
package main  
//패키지 가져오기 
import 	"fmt"	

func AAA(){
  fmt.Println("start AAA()")
    BBB()
  fmt.Println("end AAA()")
}

func BBB() {
  fmt.Println("BBB()")
  fmt.Println("BBB(B둘둘둘둘기기기기야야야야밥밥밥밥묵묵묵묵자자자자999999999999)")
}
  
//함수를 선언하고 중괄호로 시작을 알림 
func EX7-2() {
    AAA()
  }

