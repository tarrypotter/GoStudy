//constTest

package main

//패키지 가져오기
import "fmt"	

const c1 int = 0
const c2 int = 1
const c3 int = 2
//어거지로 굳이 만들어 보리기
const c4 int = 3
const c5 int = 4
const c6 int = 5

func PrintCar(car1 int , car2 int ) { 
  //어거지로 굳이 만들어 보리기
  //그냥 해본것!
  if car1 == c1 {
    fmt.Println("R8")
  }
  if car1 == c2 {
    fmt.Println("M8")
  }
  if car1 == c3 {
    fmt.Println("AMG-GT R")
  }
  if car2 == c4 {
    fmt.Println("Rs-e트론 GT")
  }
  if car2 == c5 {
    fmt.Println("i7")
  }
  if car2 == c6 {
    fmt.Println("AMG-EQS")
  }
}

//함수를 선언하고 중괄호로 시작을 알림 
func constTest() {  
    //어거지로 굳이 만들어 보리기 
    //그냥 해본것!
    fmt.Println("--------------")
    fmt.Println("Audi")
    PrintCar(c1,c4)
    fmt.Println("--------------")
    fmt.Println("Bmw")
    PrintCar(c2,c5)
    fmt.Println("--------------")
    fmt.Println("Benz")
    PrintCar(c3,c6)
    fmt.Println("--------------")
  fmt.Println("1.50%확률로",c1+c2*c3*c4*c5*c6,"억원 받기 ")
  fmt.Println("2.지금당장 ",c1+c2+c3+c4+c5+c6,"억원 받기")
    fmt.Println("--------------")
}