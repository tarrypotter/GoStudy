// if test

package main

//패키지 가져오기
import "fmt"

// 함수를 선언하고 중괄호로 시작을 알림
func if_test01() {

	Light := "red"

	if Light == "green" {
		fmt.Println("길을 건넌다")
	} else {
		fmt.Println("대기")
	}
}
