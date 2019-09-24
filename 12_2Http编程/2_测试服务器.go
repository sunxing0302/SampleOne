//访问网址是：http://127.0.0.1:8000/test
package main

import (
	"fmt"
	"net/http"
)

//服务器编写的业务逻辑处理程序
func myHandler(w http.ResponseWriter, r *http.Request) {
	//s := "\n Host:" + string(r.Host)
	//fmt.Fprintln(w, s)
	fmt.Fprintln(w, "this is a test server")
}


func main() {

	http.HandleFunc("/", myHandler)

	//在指定的地址进行监听，开启一个HTTP
	http.ListenAndServe("127.0.0.1:8000", nil)
}
