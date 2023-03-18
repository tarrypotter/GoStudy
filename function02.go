//함수2
//function

//패키지 선언
package main  
//패키지 가져오기 
import 	"fmt"	

func add02(math int , eng int, history int)int {
  return math+eng+history
}
//함수를 선언하고 중괄호로 시작을 알림 
func function02() {
  //대입연신(복사)
	a:=add02(30,60,90)
  b:=add02(90,87,100)
  c:=add02(30,20,100)
  
  fmt.Println("a의 평균 점수는 " , a/3 , " 점 ")
  fmt.Println("b의 평균 점수는 " , b/3 , " 점 ")
  fmt.Println("c의 평균 점수는 " , c/3 , " 점 ")
}

