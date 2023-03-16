//최소출력너비지정 
//SpecifyMinimumOutputWidth

//패키지 선언
package main
//패키지 가져오기
import "fmt"    
//함수를 선언하고 중괄호로 시작을 알림 
func SpecifyMinimumOutputWidth() {
    var a = 123
    var b = 4.56

    //최소너비5 총숫자 2자리  
    fmt.Printf("%5d\n",a)
    //최소너비5 총숫자 2자리 0을채움
    fmt.Printf("%05d\n",a)
    //최소너비5 소숫점이하 2자리 0을채움
    fmt.Printf("%05.2f\n",b)
    //최소너비5 소숫점이하 2자리 
    fmt.Printf("%5.2f\n",b)
    //최소너비5 총숫자 2자리 0을채움
    fmt.Printf("%05.2g\n",b)
    //최소너비5 총숫자 2자리
    fmt.Printf("%5.2g\n",b)
     //소수점 이하 6자리까지 출력
    fmt.Printf("%f\n",b)
 }
