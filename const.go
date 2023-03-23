// 상수
//const

// 패키지 선언
package main

//패키지 가져오기
import "fmt"	

const pig int = 0
const cow int = 1
const chicken int =2


func PrintAnimal(animal int) {
  if animal == pig {
    fmt.Println("돼지고기")
  }else if animal == cow {
    fmt.Println("소고기")
  }else if animal == chicken {
    fmt.Println("치킨")
  }else{
    fmt.Println("고기 아님")
  }
}

//함수를 선언하고 중괄호로 시작을 알림 
func const() {  
  fmt.Print("사줘")
  PrintAnimal(pig)
  fmt.Print("사줘")
  PrintAnimal(cow)
  fmt.Print("사줘")
  PrintAnimal(chicken)
  }

