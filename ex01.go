//그냥 예제01

//패키지 선언
package main
//패키지 가져오기
import "fmt"    
//함수를 선언하고 중괄호로 시작을 알림 
func ex01() {
    var a = 345
    var b = 3.1415
    
    fmt.Printf("%05d\n",a)
    fmt.Printf("%5.2f\n",b)
    }