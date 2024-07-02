# 数据同步API定义(v1版本)

中大型企业通常都使用自研的IAM(Identity and Access Management)系统,来统一纳管所有的员工账号、组织架构的信息, 包括员工生命周期的管理(入离职、转岗等). 

而业务系统往往也需要这部分数据, 惯用的一种做法是业务系统通过调用IAM的API, 周期性的从IAM中同步数据. 但因为这里缺少统一的API协议规范, 经常需要定制化开发. 基于此背景, 本文尝试定义一个协议, 来规范化业务系统和IAM系统的数据同步API.

该协议为"客户端+服务端"模型, 数据提供方(通常是IAM)作为服务端来实现这个协议, 数据使用方(通常是业务系统)作为客户端来按照协议来发起API调用.

## 基本概念介绍 

1. 凭证信息: client_id + client_secret
    > IAM系统给业务系统颁发的凭证信息, 业务系统通过凭证信息来请求接口调用必须的鉴权信息.
2. 鉴权access_token
   > 业务系统调用同步接口的鉴权信息. 参考下文"如何获取access_token"
3. 公开配置信息: 又称 .well-known
   >.well-known相当于一个公开的注册表, 里面包含了各个接口的请求地址, 主要包括:
    - token_endpoint: 请求access_token的接口地址  
    - list_department_endpoint: 获取部门列表的接口地址
    - list_deptartment_users_endpoint: 获取部门下用户成员详情的接口地址
    - search_department_endpoint: 部门搜索接口
    - search_user_endpoint: 用户搜索接口
4. 其他: 如无特殊说明, 所有API都遵循RESTFUL风格的定义, 包括,
    - 接口鉴权方式: Bearer Token的方式, 即Authorization: Bearer xxxxx-token, 若请求token过期或无效,接口需要返回http status为**401**
    - content-type为application/json
    - 数据结构一律为snake_case风格
    - 分页接口一律采用游标风格(cursor + size)的分页, 不支持offset方式的分页.
    - 业务侧返回的error均遵循统一的数据结构(如下)

## 数据使用方

1. 前置条件: 从数据提供方获取
    - 凭证信息(client_id + client_secret)
    - .well-known接口地址
2. 解析.well-known配置: 调用.well-known接口来获取"公开配置信息"
3. 获取鉴权access_token: 调用token_endpoint接口来获取(为了提升性能,建议缓存)
4. 数据同步:
    - 获取部门列表: 分页循环获取
    - 获取用户列表: 基于上面获取的部门id, 分页循环获取部门下用户数据 

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
    | list_deptartment_users_endpoint| url| 获取部门直属成员详情的接口地址|
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

1. 接口鉴权方式: 不需要
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
    | http status|code| 说明|
    | ---- | ---    | ---     |
    | 400 | invalid_request| 比如入参缺少client_id或client_secret|
    | 401| invalid_client| 比如client_id/client_secret校验失败|

### 获取部门列表

1. 接口鉴权方式: Bearer Token的方式, 即Authorization: Bearer xxxxx-token, 若请求token过期或无效,接口需要返回http status为**401**
2. 请求方式: GET
3. 参数说明: 以query的形式, 传递分页参数 
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | cursor| string | 分页请求的游标, 初始请求为""|
    | size| int| 分页大小, 最大支持100|
4. 返回字段说明:
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | has_next| bool| 是否还有数据未返回|
    | cursor| string |分页标记,当has_next为true时,同时返回下一次分页请求的标记. 当has_next为false时,不需要返回|
    | data| []department| 返回的部门数据, 部门的数据结构参考下面定义|
5. 部门数据结构说明
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | id| string| 部门不可变的唯一标识, 长度<=64|
    | name| string |部门名称,长度<=128 |
    | parent| string |父部门唯一标识. 若为根部门则返回""|
    | order| int| 部门在其同级部门的展示顺序 |
6. 成功返回示例:
    ```json
    {
        "has_next": true,
        "cursor": "xxxx-cursor",
        "data": [
            {"id": "1", "parent": "", "name": "中国", "order": 0},
            {"id": "1.1", "parent": "1", "name": "北京", "order": 0},
            {"id": "1.2", "parent": "1", "name": "上海", "order": 0},
            {"id": "1.3", "parent": "1", "name": "辽宁", "order": 0},
            {"id": "1.1.1", "parent": "1.1", "name": "朝阳", "order": 0}
        ]
    }
    ```
7. 错误返回示例: 参见通用的错误返回

### 获取部门直属成员详情

1. 接口鉴权方式: Bearer Token的方式, 即Authorization: Bearer xxxxx-token, 若请求token过期或无效,接口需要返回http status为**401**
2. 请求方式: GET
3. 参数说明: 以query的形式, 传递参数 
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | deptid| string | 部门唯一标识|
    | cursor| string | 分页请求的游标, 初始请求为""|
    | size| int| 分页大小, 最大支持100|
4. 返回字段说明:
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | has_next| bool| 是否还有数据未返回|
    | cursor| string |分页标记,当has_next为true时,同时返回下一次分页请求的标记. 当has_next为false时,不需要返回|
    | data| []user| 返回的部门直属用户详情数据, 用户的数据结构参考下面定义|
5. 用户详情数据结构说明
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | id| string| 用户不可变的唯一标识, 长度<=64,必须返回|
    | name| string |显示名,长度<=64,必须返回|
    | username| string |登录名,唯一,长度<=64,可不返回 |
    | email| string |邮箱,唯一,长度<=128,可不返回 |
    | mobile| string |手机号,唯一,可不返回,需遵循[E.164格式](https://en.wikipedia.org/wiki/E.164)),比如+8613411112222|
    | position| string |职务,长度<=64,可不返回 |
    | employee_number| string|工号,长度<=64,可不返回 |
    | join_time| timestamp|入职时间戳(unix timestamp),可不返回 |
    | status| int|用户状态,0:禁用, 1:待激活, 2:启用|
    | avatar| url|头像url,可不返回 |
    | main_department| string|用户所属主部门唯一标识,必须返回 |
    | other_departments| []string|用户所属副主部门唯一标识,可不返回 |
    | order| int| 部门在其主部门下的展示顺序,可不返回 |
    | extattrs| map| 其他属性,以key-value的形式存在|
6. 成功返回示例:
    ```json
    {
        "has_next": false,
        "cursor": "",
        "data": [
            {
                "id": "uid-2",
                "name": "user 2",
                "username": "user2",
                "email": "user2@example.com",
                "mobile": "+8613411112222",
                "position": "developer",
                "employee_number": "22345",
                "join_time": 1719935216,
                "avatar": "https://example.com/avatar/uid-2.png",
                "status": 1,
                "main_department": "1.1",
                "other_departments": null,
                "order": 8,
                "extattrs": {"age": 20}
            },
            {
                "id": "uid-2.1",
                "name": "user 2.1",
                "username": "user2.1",
                "email": "user2.1@example.com",
                "mobile": "+8613411113333",
                "position": "qa",
                "employee_number": "12345",
                "join_time": 1719935216,
                "avatar": "https://example.com/avatar/uid-2.1.png",
                "status": 1,
                "main_department": "1.1",
                "other_departments": ["1.2"],
                "order": 5,
                "extattrs": {"age": 30} 
            }
        ]
    }
    ```
7. 错误返回示例: 参见通用的错误返回

### 搜索部门

1. 接口鉴权方式: Bearer Token的方式, 即Authorization: Bearer xxxxx-token, 若请求token过期或无效,接口需要返回http status为**401**
2. 请求方式: GET
3. 参数说明: 以query的形式, 传递参数. 
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | keyword| string | 搜索关键字|
4. 返回说明:
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | data| []department| 返回的部门数据, 部门的数据结构同上述部门定义. 若没有匹配的部门,则接口应该返回200, 返回data为空|
5. 成功返回示例:
    ```json
    {
        "data": [
            {"id": "1.1.1", "parent": "1.1", "name": "朝阳", "order": 0},
            {"id": "1.3.1", "parent": "1.3", "name": "朝阳", "order": 0}
        ]
    }
    ```
6. 接口实现建议:
    - 最多返回10个匹配部门
    - 搜索关键字支持根据部门名称做模糊匹配, 或支持根据id进行过滤
7. 错误返回示例: 参见通用的错误返回

### 搜索用户

1. 接口鉴权方式: Bearer Token的方式, 即Authorization: Bearer xxxxx-token, 若请求token过期或无效,接口需要返回http status为**401**
2. 请求方式: GET
3. 参数说明: 以query的形式, 传递参数. 
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | keyword| string | 搜索关键字|
4. 返回说明:
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | data| []user| 返回的用户详情数据, 用户的数据结构同上述部门定义. 若没有匹配的用户,则接口应该返回200, 返回data为空|
5. 成功返回示例:
    ```json
    {
        "data": [
            {
                "id": "uid-2.1",
                "name": "user 2.1",
                "username": "user2.1",
                "email": "user2.1@example.com",
                "mobile": "+8613411113333",
                "position": "qa",
                "employee_number": "12345",
                "join_time": 1719935216,
                "avatar": "https://example.com/avatar/uid-2.1.png",
                "status": 1,
                "main_department": "1.1",
                "other_departments": ["1.2"],
                "order": 5,
                "extattrs": {"age": 30} 
            }
        ]
    }
    ```
6. 接口实现建议:
    - 最多返回10个匹配用户
    - 搜索关键字支持根据用户名称做模糊匹配, 或支持根据id、登录名、邮箱、手机号进行过滤
7. 错误返回示例: 参见通用的错误返回

## 附录

### 分页请求

1. 分页请求参数一律采用游标风格(cursor + size)的分页, 不支持offset方式的分页.
2. 分页请求的参数为:
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | cursor| string | 分页请求的游标, 初始请求为""|
    | size| int| 分页大小, 最大支持100|
3. 分页请求的返回为:
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | has_next| bool| 是否还有数据未返回|
    | cursor| string |分页标记,当has_next为true时,同时返回下一次分页请求的标记. 当has_next为false时,不需要返回|
    | data| []object| 返回的数据|

### 通用错误返回

1. 错误返回复用http status作为接口返回的状态码
2. 接口业务侧使用通用的数据结构来返回错误
    | 字段名| 类型| 说明|
    | ---    | ---   | ---     |
    | code| string | 错误码, 当有错误时,必须返回. 为""时表示请求成功|
    | msg| string| 错误的描述信息, 建议返回|
    | request_id| string|请求的唯一标识, 建议返回|
3. 返回示例: http 401
    ```json
    {
        "code": "invalid_client",
        "msg": "invalid client id or client secret",
        "request_id": "zBXaFZpIYrsllcrEjAEBoqIUpuuQFgzq"
    }

    ```
4. 常见错误码
    | http status|code| 说明|
    | ---- | ---    | ---     |
    | 400 | invalid_request| 比如非法请求,入参缺少client_id或client_secret|
    | 401| invalid_client| 无效的client_id或client_secret|
    | 401| invalid_token| 接口调用时access_token校验失败|
    | 429| too_many_requests| 接口调用超出频率限制|

### 访问频率限制

1. 单接口的最大QPS限制在50次/秒
2. 若数据提供方拦截到请求超出频率限制时, 需要返回http status为429, 业务侧的错误返回码为:
```json
    {
        "code": "too_many_requests",
        "msg": "too many requests",
        "request_id": "zBXaFZpIYrsllcrEjAEBoqIUpuuQFgzq"
    }
```
4. 建议接口调用方主动限制频率, 若被限频, 则sleep 1s后进行重试

## 参考实现

见[go的参考实现](https://github.com/idaaser/syncdemov1)