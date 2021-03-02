package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	username := request.FormValue("username")
	//	gender := request.FormValue("gender")
	//	fmt.Fprintln(writer, fmt.Sprintf("用户名：%s,性别:%s", username, gender))
	//})

	http.HandleFunc("/main.html", func(writer http.ResponseWriter, request *http.Request) {
		GetPageBooksByPrice(writer, request)

	})

	//request.HeaderExtractor{}.Get("Content-Type") //返回的是string
	//request.Header["Content-Type"]     //返回的是[]string
	fmt.Println(http.ListenAndServe(":8080", nil))

}

func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//获取价格范围
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	fmt.Fprintln(w, fmt.Sprintf("用户名：%s,性别:%s", minPrice, maxPrice))
}
