// fmt패키지 표준입력 scan,scanf,scanln 

//패키지 선언
package main
//패키지 가져오기
import "fmt"
//함수를 선언하고 중괄호로 시작을 알림 
func fmtPackage02() {
  //scan 쓸때 a,b변수 쓰고 나머지는 주석으로 해놓으면 됨.
   var a int 
   var b int 

  //scanf 쓸때 c,d변수 쓰고 나머지는 주석으로 해놓으면 됨.
   //var c int 
   //var d int 

  //scanln 쓸때 e,f변수 쓰고 나머지는 주석으로 해놓으면 됨.
   //var e int 
   //var f int

  //scan
   //위에 변수선언 확인!! 주석풀었는지 제대로 적용했는지 
   // 한칸 띄어서 두개 입력 
   // ex)  5 3 
  n , err := fmt.Scan(&a,&b)
  if err != nil {
   fmt.Println(n,err)
  }else {
    fmt.Println(n,a,b)
  }
  
  
  //scanf 
    //위에 변수선언 확인!!  주석풀었는지 제대로 적용했는지 
    // 한칸 띄어서 두개 입력 
    // ex)  5 3 
   /*n,err := fmt.Scanf("%d %d\n", &c,&d)
    if err != nil {
      fmt.Println(n,err)
    }else {
      fmt.Println(n,c,d)
    }
  */
  
  //scanln
   //위에 변수선언 확인!! 주석풀었는지 제대로 적용했는지 
   // 한칸 띄어서 두개 입력 
   // ex)  5 3 
 /* n,err := fmt.Scanln(&e,&f)
    if err != nil {
      fmt.Println(n,err)
    }else {
      fmt.Println(n,e,f)
    }
 */

} // <-주석달때 조심 이거 걸리면 오류남 
  
 // scan은 표준 입력에서 값을 입력.
 // scanf은 표준 입력에서 서식형태로 값을 입력받음.
 // scanln은 표준입력에서 한줄을 읽어서 값을 입력받음.
 // %s는 문자열을, %d는 10진수 정수를, %f는 부동 소수점 숫자를 출력.
 // \n은 줄바꿈

 //부가설명
  // scan() 함수는 사용자로부터 입력 받은 내용을 공백과 개행 문자로 분리하여 입력 값을 받는 함수.
  // ex) "hello world"를 입력 받을 때, scan() 함수는 "hello"를 입력 받고, 공백 이후의 "world"는 받지 않음.
  // 또한 입력 값을 변수에 할당할 때 포인터 변수를 사용해야함.

  // scanf() 함수는 형식 지정자를 사용하여 입력 값을 받는 함수.
  // 형식 지정자는 C 언어와 유사하게 사용된다. scanf() 함수는 공백을 포함한 모든 문자열을 입력으로 받음.

  // scanln() 함수는 사용자로부터 입력 값을 받을 때 개행 문자를 사용해 한 줄 전체를 입력 받는 함수.
  // ex) "hello world"를 입력 받을 때, scanln() 함수는 "hello world"를 한 줄 전체로 입력 받음.