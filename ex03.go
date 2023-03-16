//그냥 예제03

//fmt.Printf와 서식문자를 이용한코드 출력결과의 최소너비 6으로 지정.
//출력결과
  //   123
  //004567
  //  3.14

//패키지 선언
package main
//패키지 가져오기
import "fmt"    
//함수를 선언하고 중괄호로 시작을 알림 
func ex03() {
    var a =123
    var b int = 4567
    f := 3.14159269

  // a를 이용해 출력
     fmt.Printf("%6d\n",a)
  // b를 이용해 출력
     fmt.Printf("%06d\n",b)
  // f를 이용해 출력
     fmt.Printf("%6.2f\n",f)
     fmt.Printf("%6.3g\n",f)
    
 }
