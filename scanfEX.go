//scanf 
//패키지 선언
package main  
// 이거 왜 안되지 ㅜㅜexport LANG=ko_KR.UTF-8
//패키지 가져오기 
import 	"fmt"	
//함수를 선언하고 중괄호로 시작을 알림 
func scanfEX() {
	var a string
	var b int
	fmt.Scanf("%s\n",&a)
  fmt.Scanf("%d\n",&b)
	fmt.Println(a,b)
}

