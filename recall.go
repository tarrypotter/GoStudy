//재귀함수
//recall

//패키지 선언
package main  
//패키지 가져오기 
import 	"fmt"	

func printNo(n int){
  if n ==0 {
    return
  }  
   fmt.Println(n)
   printNo(n-1)
   fmt.Println("After",n)
}
//함수를 선언하고 중괄호로 시작을 알림 
func recall() {
    printNo(3)
  }

