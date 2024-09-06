
## 项目结构
`
|--cli 客户端 main
|--cmd 服务端 main
|--global 全局对象
|--internal 服务端、客户端中间件实现
|--pkg 错误处理、meta、swagger、tracer封装
|--proto grpc proto文件及其生成代码文件
|--server 业务代码
|--third_party 额外的引入
`

## 使用到功能
1. grpc proto工具和生成代码；
2. swagger 生成静态文件及如何作为接口文档引入grpc项目；
3. jaeger 链路追踪引入grpc项目；
4. grpc中间件实现，以及如何利用链式中间件，突破grpc只允许设置单个中间件的限制；
5. 引入 grpc-gateway 实现但端口提供 http 和 grpc 接口；
6. 