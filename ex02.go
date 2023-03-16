//그냥 예제02

//패키지 선언
package main
//패키지 가져오기
import "fmt"    
//함수를 선언하고 중괄호로 시작을 알림 
func ex02() {
    var a int  
    var b int

   n , err := fmt.Scan(&a,&b)
  if err != nil {
   fmt.Println(n,err)
  }else {
    fmt.Println(n,a,b)
  }
    
    fmt.Scanln(a,b)
    fmt.Println(a,b)
 }
