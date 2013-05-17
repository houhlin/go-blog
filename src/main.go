/**
 * User: Baoxu
 * Date: 13-5-16
 * Time: 下午4:07
 */
package main

import (
	"net/http"
	"log"
	"html/template"
	"fmt"
)

/**
 * 自定义类型，路由控制器
 * 根据URI中不同的地址，执行不同的函数
 */
type Router struct {
}

/**
 * 路由控制器指定路由之后的处理函数
 * 如果是请求根目录，执行此函数
 */
func serverDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	//解析模板
	t, err := template.ParseFiles("tmpl/html/index.html")
	if (err != nil) {
		log.Println(err)
	}
	p := Person{UserName: "我是内部成员变量"}
	//执行模板，合并变量并输出
	t.Execute(w, p)

}

/*******************************************************************/
/*                  主执行函数，主要作为端口监听服务器                   */
/*******************************************************************/
func main() {
	//设置访问路由
	http.Handle("/css/", http.FileServer(http.Dir("tmpl")))
	http.Handle("/js/", http.FileServer(http.Dir("tmpl")))

	//设置控制器路由
	http.HandleFunc("/", serverDefault)

	//设置监听的端口
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
