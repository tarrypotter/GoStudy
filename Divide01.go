//멀티반환함수
//Divide

//패키지 선언
package main  
//패키지 가져오기 
import 	"fmt"	

func Divide(a,b int)(int,bool) {
  if b ==0 {
    return 0, false 
  }
  return a/b , true 
}
//함수를 선언하고 중괄호로 시작을 알림 
func Divide01() {
    c,success := Divide(9,3)
    fmt.Println(c,success)
    d,success := Divide(9,0)
    fmt.Println(d,success)
  }

