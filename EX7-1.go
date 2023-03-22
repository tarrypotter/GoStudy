// EX7-1

//패키지 선언
package main  
//패키지 가져오기 
import 	"fmt"	

func Multiple(a,b int)int{ 
  return a*b
}

//함수를 선언하고 중괄호로 시작을 알림 
func EX7-1() {
     c := Multiple(1,9)
     d := Multiple(3,3)
     e := Multiple(9,1)
  
     fmt.Print("비둘기야 밥묵자",c)
     fmt.Print(c)
     fmt.Println(c)
     fmt.Println("---------------")
     fmt.Print("비둘기야 밥묵자",c,d,e)
  }