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
)

/**
 * 自定义类型，路由控制器
 * 根据URI中不同的地址，执行不同的函数
 */
type Router struct {
}

/**
 * 实现路由控制器
 * param w 输出到用户客户端的信息
 * param r 客户端的请求信息
 */
func (p *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		serverDefault(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

/**
 * 路由控制器指定路由之后的处理函数
 * 如果是请求根目录，执行此函数
 */
func serverDefault(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("serverDefault.gtpl")
	t.Execute(w, nil)
	//fmt.Fprintf(w, "Based On GO Language")
}

/*******************************************************************/
/*                  主执行函数，主要作为端口监听服务器                   */
/*******************************************************************/
func main() {
	//设置控制器路径
	mux := &Router{}
	//设置监听的端口
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
