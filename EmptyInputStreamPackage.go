//scanln함수가 실패한경우 표준입력스트림을 비워주는 예제
// EmptyInputStreamPackage

//패키지 선언
package main
//패키지 가져오기
import (
   "bufio"  // io를 감당하는 객체
   "fmt"    // 표준 입출력등을 가지고 있는 객체
   "os"
)
//함수를 선언하고 중괄호로 시작을 알림 
func EmptyInputStreamPackage() {
    stdin := bufio.NewReader(os.Stdin) //표준 입출력을 읽는 객체

  var a int
  var b int

  n , err := fmt.Scanln(&a,&b)

  if err != nil {              //에러 발생시
      fmt.Println(err)         //에러출력
      stdin.ReadString('\n')   //표준 입출력 스트림 지우기
    }else {
      fmt.Println(n,a,b)
    }

  n,err = fmt.Scanln(&a,&b)     //다시 입력받기
    if err != nil{
      fmt.Println(err)
    }else {
      fmt.Println(n,a,b)
    }
  }

//표준 입력 스트림에서 한 줄을 읽어오는데 bufio,os등의 패키지를 사용.
//bufio는 입력스트림으로 한줄을 읽는 Reader객체를 제공
//func NewReader(rd io.Reader) *Reader

//ner Reader()함수는 인수로 입력되는 입력스트림을 가지고 Reader 객체를 생성해줌.
//이 코드에선 표준 입력 스트림을 나태내는 os.stdin을 사용해서 Reader객체를 만듬.

//stdin.ReadString('\n') 은 줄바꿈 문자가 나올 때까지 읽고 이렇게 하면 표준 입출력스트림이 비워짐.

