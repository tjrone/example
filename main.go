package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func main(){
	http.HandleFunc("/abc", index)
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	num:=os.Getenv("Num")
	if num==""{
		Fibonacci(10)
	}else{
		numInt,_:=strconv.Atoi(num)
		Fibonacci(numInt)
	}
	w.Write([]byte("success"))
	log.Println("success")
}

func Fibonacci(n int)int{
	if n<=2{
		return 1
	}else{
		return Fibonacci(n-1)+Fibonacci(n-2)
	}
}
