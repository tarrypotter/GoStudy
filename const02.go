// 타입 없는 상수

// 패키지 선언
package main

//패키지 가져오기
import "fmt"	

const PI = 3.14 //타입 없는 상수
const FloatPI float64 = 3.14 //float64 타입상수


//함수를 선언하고 중괄호로 시작을 알림 
func const02() {  
    var a int = PI * 100
    // var b int = FloatPI * 100  -> FloatPI는 float64타입이기 때문에 오류남

  fmt.Println(a)
  // fmt.Println(b)
  }

