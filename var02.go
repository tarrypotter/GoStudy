//지역변수 선언하는 코드
//변수 범위

//패키지 선언
package main
//패키지 가져오기
import "fmt"
//함수를 선언하고 중괄호로 시작을 알림 

// 패키지 전역 변수 선언 
var a int =3

func var02() {
  //지역 변수 선언
  var b int = 4
  {
    //지역 변수 선언
    var c int = 5
    fmt.Println(c)
  } // 변수 c 사라짐
  fmt.Println(a,b)
}