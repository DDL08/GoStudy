# GoStudy
学go的时候开的

[![CodeSize](https://img.shields.io/github/languages/code-size/DDL08/GoStudy)](https://github.com/DDL08/GoStudy)

[![LICENSE](https://img.shields.io/badge/license-Apache--2.0-blue)](https://www.apache.org/licenses/LICENSE-2.0)


<img src="https://github.com/DDL08/images/blob/main/hr2.png?raw=true" width="600px" />

## 一
### 1
net/http 标准库来实现一个简单的 HTTP 服务器。这个代码定义了两个 HTTP 处理函数 Handler，然后使用 http.HandleFunc 将它们分别绑定到路径 / 和 /hello 
### 2 
定义空结构体engine 通过 ServeHTTP 方法来实现net/http库的http.Handler接口的一个方法。（ServeHTTP是net/http的一个接口http.Handler的一个核心方法）

通过实现 http.Handler 接口，Engine 可以直接被 http.ListenAndServe 使用
### 3
gee框架雏形,显示更改一下架构，然后区分了get和post,用map来装【通过engine封装后的路由】

gee/

  |--gee.go
  
  |--go.mod
  
main.go

go.mod

<img src="https://github.com/DDL08/images/blob/main/hr.gif?raw=true" width="600px" />

## 二
新建router.go、Context.go

结构体不再是路由map，而是声明了一个指针函数的结构体，然后单开一个router.go里面装路由map，Context是把请求个响应的请求方法+文本类型进行一个封装

<img src="https://github.com/DDL08/images/blob/main/hr.gif?raw=true" width="600px" />

## 三
新建tire.go

实现动态路由的注册，新建tire.go这个node结构，来实现路由节点的插入以及查找

修改👇

Context结构小改--->新加属性	Params     map[string]string

router结构大改--->改添加路由对应tire，新加getRoute+getRoutes

<img src="https://github.com/DDL08/images/blob/main/hr.gif?raw=true" width="600px" />

## 四
实现路由分组控制(Route Group Control)

新建结构体RouterGroup struct

原因:某一组路由需要相似的处理

测试👇

curl.exe "http://localhost:9999/v2/login/" -X POST -d "username=1" -d "password=2"

curl.exe "http://localhost:9999/v2/hello/121"

curl.exe "http://localhost:9999/v1/hello?name=121"


## 五

框架需要有一个插口，允许用户自己定义功能，嵌入到框架中，就是中间件(middlewares)

