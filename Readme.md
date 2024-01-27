# bochat

## 環境建置
change log 
https://github.com/git-chglog/git-chglog


## 資料夾擺放規則

```shell
├── cmd                    application
├── configuration          服務有哪些控制項
├── deployment             服務部署資訊包括 Config file 和 dockerfile 但目前沒有在使用
├── docker-compose.yml     本地需要的一些服務
├── docs
│   └── graph              GraphQL 如果要更改 schema 會到這個資料夾更改 然後 make gen.graphql
│       └── schema
│           ├── app
│           └── platform
├── internal               一些各種第三方元件的包裝
├── makefile               基本上所有操作指令都會使用 makefile
├── pkg                    服務有的模組
│   ├── converter          轉換模組 兩方的結構皆不存在 pkg 底下, 拆開後以後獨立出去會較簡單
│   │   └── auth.go
│   ├── delivery           傳輸層, 分為 GRPC, Restful, GraphQL, Worker
│   │   ├── graph
│   │   │   ├── platform
│   │   │   └── view       資料庫的資料轉換成 GraphQL 的資料
│   │   ├── redis_worker  redis cache 及 queue 的 worker
│   │   └── restful
│   │       └── backend
│   ├── hub                管理長連線的機制
│   ├── iface              抽象 handler, service, repository
│   ├── infra              一些基礎設施, 例如 logger, db, cache, config, redis, grpc, graphql
│   ├── middleware         中介層
│   ├── model              對應 database 的資料結構
│   │   ├── dto            database 的資料結構
│   │   ├── option         操作 database 各種 查詢, 更新的資料結構
│   │   ├── enums          各種 enum
│   │   ├── events         各種事件處理的觸發，使用 redis 的 __keyevent@%d__:expired 來進行事件觸發
│   │   └── vo             用來傳輸的資料結構
│   ├── repository         資料庫溝通層
│   │   ├── cache          針對 redis 包裝的 repo
│   └── service            各種商業邏輯會放在這下面, 後續資料夾可能會依照 domain 做區分
├── platform_gqlgen.yml    產生 graphql 的設定檔文件
└── test                   測試需要用到的文件可能都會擺在這底下
```

## 開發流程

### 新增接口

1. GraphQL
  a. docs/graph/{所屬的domain}/ 底下增加相對應的 service
  b. make gen.graphql
  c. pkg/delivery/graph/{所屬的domain}/ 實作該 service
  d. mode 的轉換在 pkg/delivery/graph/view 中實作
2. Restful
  a. pkg/delivery/restful/{所屬的domain}/ 參照目錄底下實作
  b. 在各自的 handler 中實作方法接口，並在 router.go 中註冊
  c. 方法上自行加入Summary, Description, Tags, Param, SuccessResponse, FailureResponse 等註解
  d. 執行 make gen.swagger
3. Worker
  a. pkg/delivery/redis_worker/{所屬的domain}/ 參照目錄底下實作

GraphQL 一些大致上的用法
query 無序 Read
mutation 有序 Write Update Delete
subscription websocket

### 新增中介層
- 在 middleware 下新增 middleware ，可參考：
  - pkg/middleware/host_deny.go

### 新增抽象接口層
- 在 iface 下新增 interface ，可參考：
  - pkg/iface/service.go
  
### 新增商務邏輯層
- 實作方法於物件下即可於 handler 層內物件呼叫到 可參考：
  - pkg/service/user.go

### 新增資料庫溝通層
- 實作方法於物件下即可於 service 層內物件呼叫到 可參考：
  - pkg/repository/user.go
  
### 資料庫同步方法
- 將物件定義好後並將新表物件添加至以下檔案
  - pkg/repository/BaseRep.go
  
### Swagger文件生成
- 於終端機下指令  make gen.swagger 即可

### 新增排程


### 新增工具類
- 在 pkg/infra 下新增工具類，可參考：
  - pkg/infra/


### 新增事件觸發
- 使用redis 的 `__keyevent@%d__:expired` 來進行事件觸發
- 並由 `expired.go` 的 `HandlerExpired` 來判別並觸發該事件處理，可參考：
  - pkg/delivery/redis_worker/expired.go


### 常用服務說明
1. auth_svc : 認證服務，提供使用者的註冊、登入、登出、修改資料等功能 => 實際的登入處理由bochat-iam 處理

### API Authentication
- 取得 token
  打開 GraphQL playground <http://127.0.0.1:8080/graph/playground>
  呼叫api : http://127.0.0.1:8080/graph/query
    Request:
    ```graphql
    mutation {
      Login(in: {username: "jimmy", password: "000000"}) {
        token,
        deviceID
      }
    }
    ```

    Response:

    ```json
    {
      "data": {
        "Login": {
          "token": "2d00e899-4662-4333-aa48-7f52deaf081b.fx0BTlR3ERqfubS3aTMxqTQRR-g",
          "deviceID": "23ec5ceb92edd62de9f99deeeec85821"
        }
      }
    }
    ```

- 使用 token

  - GraphQL Header:

    ```json
    {
      "Authorization": "Bearer <token>",
      "X-Device-ID": "<deviceID>"
    }
    ```

  - GRPC Header:

    ```
      token: <token>
      device_id: <deviceID>
    ```