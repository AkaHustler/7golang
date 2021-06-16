# 7golang
## day1-http
- 实现了web框架的基础功能，提供了创建实例、添加路由、启动web服务。实现了路由映射表，支持用户添加静态路由，包装了启动服务的函数。
- go mod 的使用，在包的目录下使用go mod init example，其他包引用相对路径的package需要使用
  
  require gee v0.0.0
  
  replace gee => ./gee

  在go.mod中使用 **replace** 将gee指向 **./gee** 指向其包目录
- 通过查看http包的源码发现，ListenAndServe启动web服务的第二个参数需要实现 **ServeHTTP** 方法，换而言之，所有实现ServeHTTP方法的接口实例，所有的HTTP请求，都会交给该实例处理。
  
  即，接收一个（或多个接口）作为参数的函数，其 **实参** 可以是任何实现了该接口的类型的变量。**实现了某个接口的类型可以被传给任何以此接口为参数的函数**
- Go语言中，实现了接口方法的struct都可以强制转换为接口类型

## day2-context 上下文
- 必要性：对于web服务来说，无非是根据 *http Request，构造响应 http.ResponseWriter。但是这两个对象提供的接口粒度太细，如果我们要构造一个完整的响应，需要考虑消息头(Header)和消息体(Body)，而Header又包括了状态码(StatusCode)，消息类型(ContentType)等几乎每次请求都需要设置的信息。

  因此如果不进行有效的封装，那么用户将需要做大量重复、繁杂的工作，而且容易出错。
- 针对使用场景，封装 *http.Request 和 http.ResponseWriter 的方法，简化相关接口的调用，只是设计Context的原因之一。对于框架而言，还需要支撑额外的功能。

  例如，在解析动态路由 **/hello/:name** ，参数:name放在哪里？框架支持中间件，中间件产生的信息放在哪里？Context随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由Context承载。设计Context结构，扩展性和复杂性留在了内部，对外简化了接口。

## day3-router 前缀树路由
