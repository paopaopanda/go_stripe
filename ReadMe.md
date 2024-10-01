stripe 流程:
fronted 發送 request ->
backend(server.go) 發送 前端提供的資訊至 stripe.go ->
stripe server 將 client_key 回傳至 server.go ->
server.go 將 client_key 傳至 frontend
