### 创建用户
> curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"kong","password":"kong123"}' | fx

### 删除用户
> curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/2

### 更新用户
> curl -XPUT -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/2 -d'{"username":"kong","password":"kongmodify"}' | fx

### 查看用户
> curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/kong | fx
> 为什么查看用户要通过用户名？？？

### 获取用户列表
> curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8082/v1/user -d'{"offset": 0, "limit": 20}' | fx

### 登陆接口
> curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/login -d'{"username":"admin","password":"admin"}'
#### 请求不带token:
> curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"user1","password":"user1234"}'
#### 请求携带token
> curl -XPOST -H "Authorization: Bearer eyJhbGcixxxxxxxCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ" -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"user1","password":"user1234"}'

### 验证https可用性
#### 请求不携带证书
> curl -XGET -H "Authorization: Bearer testtesttesttesttesttesttesttesttesttest" -H "Content-Type: application/json" https://127.0.0.1:8081/v1/user
#### 请求携带证书
> curl -XGET -H "Authorization: Bearer teesstesttesttesttesttesttesttesttest" -H "Content-Type: application/json" https://127.0.0.1:8081/v1/user --cacert conf/server.crt --cert conf/server.crt --key conf/server.key