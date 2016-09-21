package main

import (
	"fmt"
	pref "./preferences"
)

// desk 에서 사용하는 
func system_print_string(before string) string {
	return "\t\t\t>>>> "+before+" <<<<<\t\t\t"
}

// desk 프로그램이 종료할때의 작업을 구현합니다
func shutdown(){
	fmt.Println(system_print_string(pref.GOING_SHUTDOWN_MESSAGE))
}
//test
func main(){
	var path string
	// var _ n int
	for{
		fmt.Println(system_print_string(pref.GOING_WAIT_REQUEST_MESSAGE))
		fmt.Scanln(&path)
		if path == "exit"{
			shutdown()
			break
		}
		fmt.Println(path)
	}
	
}