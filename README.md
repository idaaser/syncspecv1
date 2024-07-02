# 数据同步API定义(v1版本)

企业内部通常都有自己的IAM系统, 来统一管理所有的员工账号、组织架构的信息, 包括员工生命周期的管理(入离职、转岗等). 

而业务系统通常又需要这部分数据, 惯用的一种做法是业务系统通过调用IAM系统的API, 周期性的进行数据同步. 但这里的API缺少一个统一的协议规范, 所以本文档尝试定义一个协议, 业务系统通过这个协议规范来进行数据同步, IAM系统基于这个协议规范来返回数据. 

该协议为客户端、服务端模型, 数据提供方(通常是IAM)作为服务端来实现这个协议, 数据适用方(通常是业务系统)作为客户端来按照协议来发起API调用.

## 基本概念介绍 

1. 凭证信息: client_id + client_secret

   IAM系统给业务系统颁发的凭证信息, 业务系统通过凭证信息来请求接口调用必须的鉴权信息.
2. 鉴权access_token
   
   业务系统调用同步接口的鉴权信息. 参考下文"如何获取access_token"
3. 公开配置信息: 又称 .well-known

   .well-known相当于一个公开的注册表, 里面包含了各个接口的请求地址, 主要包括:
    - token_endpoint: 请求access_token的接口地址  
    - list_department_endpoint: 获取部门列表的接口地址
    - list_deptartment_users_endpoint: 获取部门下用户成员详情的接口地址
    - search_department_endpoint: 部门搜索接口
    - search_user_endpoint: 用户搜索接口
4. 其他: 如无特殊说明, 所有API都遵循RESTFUL风格的定义, 包括:
    - content-type为application/json
    - 数据结构一律为snake_case风格
    - 分页接口一律采用游标风格(cursor + size)的分页, 不支持offset方式的分页.
    - 业务侧返回的error均遵循统一的数据结构(如下)

## 数据使用方

1. 配置获取: 从数据提供方获取
    - 凭证信息(client_id + client_secret)
    - .well-known接口地址
2. 获取.well-known配置: 调用.well-known接口来获取"公开配置信息"
3. 获取鉴权access_token: 调用token_endpoint接口来获取(为了提升性能,建议缓存)
4. 数据同步:
    - 获取部门列表: 分页循环调用
    - 基于部门id, 分页循环获取部门下用户数据 

## 数据提供方

作为服务端, 需要实现下述的几个API.

### 公开配置信息接口(.well-known)

1. 接口鉴权方式: 不需要鉴权, 公开接口
2. 请求方式: GET
3. 参数说明: 不需要
4. 返回字段说明:
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | spec| string | 协议版本号, 固定为v1|
    | token_endpoint| url| 获取access_token的接口地址|
    | list_department_endpoint| url| 获取部门列表的接口地址|
    | list_deptartment_users_endpoint| url| 获取部门成员详情的接口地址|
    | search_department_endpoint| url| 搜索部门的接口地址|
    | search_user_endpoint| url| 搜索用户的接口地址|
5. 返回示例:
    ```json
    {
        "spec": "v1",
        "token_endpoint": "https://example.com/v1/token",
        "list_department_endpoint": "http://example.com/v1/depts",
        "list_deptartment_users_endpoint": "https://example.com/v1/users",
        "search_department_endpoint": "http://example.com/v1/depts:search",
        "search_user_endpoint": "http://example.com/v1/users:search"
    }
    ```

### 请求access_token 

1. 接口鉴权方式: 通过调用参数, 如下
2. 请求方式: POST 
3. 参数说明: body的方式传入 
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | grant_type| string | 固定为client_credentials|
    | client_id| string| 凭证信息中的client_id|
    | client_secret| string| 凭证信息中的client_secret|
4. 返回字段说明:
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | token_type| string | token类型,固定为Bearer|
    | access_token| string| 实际返回的access_token的值|
    | expires_in| int| token的有效期, 单位为秒, 比如7200表示2小时. 调用者可以根据有效期来缓存token|
5. 返回示例:
    ```json
    {
        "token_type": "Bearer",
        "access_token": "xxxxxxxxxxxxx-access-token",
        "expires_in": 1800
    }
    ```
6. 错误返回: 当颁发access_token失败时, http status返回400或401, 常见错误返回包括 
    | status|code| 说明|
    | ---- | ---    | ---     |
    | 400 | invalid_request| 比如入参缺少client_id或client_secret|
    | 401| invalid_client| 比如client_id/client_secret校验失败|
    

### 获取部门列表

### 搜索部门

### 获取部门成员详情

### 搜索用户

## 附录

### 分页请求

### 错误码

### 访问频率限制

## 参考实现

见[go的参考实现](https://github.com/idaaser/syncdemov1)