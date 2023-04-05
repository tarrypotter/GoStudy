// ListenAndServe 함수를 호출해 웹서버 시작 
//요청을 받으면 문자열을 반환 하는 웹서버

package main 

import (

	"fmt"
	"net/http"
)

func ListenAndServe() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"hi!")
   })
	
	http.ListenAndServe(":3000",nil)
}