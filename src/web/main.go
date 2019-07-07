package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func sayHelloController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数
	fmt.Println(r.Form) // 请求信息
	fmt.Println("path ", r.URL.Path)
	fmt.Println("schema ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Printf("key: %s, val: %s\n", k, strings.Join(v, ","))
	}
	fmt.Fprintf(w, "Hello world")
}

/**
* Form的提交内容都会转化为url.Values类型，是key/value形式
* r.FormValue解析Form的元素，不需要调用r.ParseForm()
* POST,GET的参数如果一样，变量的值会放到一个slice里面，用r.FormValue取值只会取到第一个值
 */
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		fmt.Println("username: ", r.FormValue("username"))
		fmt.Println("password: ", r.FormValue("password"))

		// 校验，正则表达式
		m, _ := regexp.MatchString("^[a-zA-Z]$", r.FormValue("username"))
		if !m {
			fmt.Println("error")
		}
		http.Error(w, "bad request", http.StatusBadRequest)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		curTime := time.Now().Unix() // 获取时间戳
		fmt.Println("curTime:", curTime)
		h := md5.New()                                    // hash对象
		io.WriteString(h, strconv.FormatInt(curTime, 10)) // 将字符串写入h
		fmt.Println("h:", h)
		token := fmt.Sprintf("%x", h.Sum(nil)) // 计算hash直接
		fmt.Println("token:", token)
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(1024)               // 定义允许上传的文件大小内存部分的最大容量，其余的存在临时文件
		file, handler, err := r.FormFile("file") //获取文件
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 创建文件
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file) //复制文件
	}
}

// go不需要web容器，本身的http包已经扮演了nginx的角色，打包为执行文件，执行就是一个web服务端程序
// go程序直接监听端口
/**
* 1. 首先通过调用，往DefaultServerMux的一个map添加路由规则
* 2. 然后监听端口，每次一个新的请求都会用一个gorountine处理
* 3. 如果自定义handle为空，就用DefaultServerMux的路由规则
* 4. 否则就使用handle的规则对请求进行处理
 */
func main() {
	http.HandleFunc("/", sayHelloController) // 设置路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil) // 监听端口
	if err != nil {
		log.Fatal("error ", err)
	}

}
