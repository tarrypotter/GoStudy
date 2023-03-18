//함수 
//function

//패키지 선언
package main  
//패키지 가져오기 
import 	"fmt"	

func add (a int , b int)int {
  return a+b
}
//함수를 선언하고 중괄호로 시작을 알림 
func function() {
  //대입연신(복사)
	c:=add(3,6)

  fmt.Println(c)
  fmt.Println("a+b=" , c)

}

