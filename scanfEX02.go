//scan예제2

//패키지 선언
package main  
//패키지 가져오기 
import 	"fmt"	
//함수를 선언하고 중괄호로 시작을 알림 
func scanfEX02() {
	var a string
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	c := fmt.Sprintf("%s %d", a, b)
	fmt.Println(c)
}