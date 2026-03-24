# Boyi 專案結構說明

本文件整理目前專案的主要目錄、模組責任、分層關係與維護建議，協助新成員快速理解系統。

## 1. 專案定位與整體架構

此專案為 Go 單體後端（Monolith with modular packages），使用分層與模組化混合設計。

- 入口層：`main.go` + `cmd/`（不同執行模式）
- 設定層：`configuration/` + `app*.properties`
- 交付層：`pkg/delivery/`（REST、Socket、Worker）
- 應用層：`pkg/service/`、`pkg/scheduler/`
- 領域模型層：`pkg/model/`
- 資料存取層：`pkg/repository/`
- 基礎設施層：`pkg/infra/`（DB、Redis、RabbitMQ、Mail、Storage 等）
- 介面抽象：`pkg/iface/`（handler/repository/service 介面）
- 測試與模擬：`internal/mock/`、`internal/test_fixture/`、`test/`

## 2. 目錄結構與職責

### 根目錄

- `main.go`：預設啟動入口。
- `go.mod`：Go module 與依賴管理。
- `makefile`：常用建置、測試與工具指令。
- `Dockerfile`：容器建置設定。
- `README.md`：專案說明與啟動指引。
- `CHANGELOG.md`：版本變更紀錄。
- `app*.properties`：不同環境與角色（admin/client/worker/schedule）配置。

### `cmd/`

提供不同角色啟動點，例如：

- `server.go`：主要服務進程。
- `worker.go`：背景任務進程。
- `scheduler.go`：排程進程。
- `migration.go`：資料或流程遷移入口。

此設計讓同一份程式碼可依用途切換執行模式。

### `configuration/`

- `config.go`：配置讀取與整合。

集中配置邏輯，避免設定值散落在各模組。

### `deployment/`

- `database/`：SQL migration 檔案。
- `file/`：初始化資料（CSV）。

部署與初始化資源集中管理，有利於環境重建與維運。

### `docs/`

- 專案文件與歷史資料（如 `old_project/`）。

### `internal/`

存放內部專用程式與測試輔助：

- `claims/`、`lock/`：基礎元件。
- `mock/`：外部服務與 repository mock。
- `test_fixture/`：測試資料與載入邏輯。

Go 的 `internal` 規則可限制外部引用，保護內部實作。

### `pkg/`

核心業務程式碼，主要依分層組織：

- `delivery/`：對外輸入輸出層（REST、Socket、MQ worker、Redis worker）。
- `service/`：應用服務層，承載業務流程與規則。
- `repository/`：資料存取邏輯。
- `model/`：DTO、enum、events、option 等模型。
- `infra/`：第三方與基礎設施整合（db、redis、rabbitmq、mail、storage、otel）。
- `middleware/`：橫切邏輯（JWT、權限、白名單、防重複請求、請求紀錄）。
- `iface/`：介面抽象，降低層間耦合。
- `hub/`：即時連線與訊息分發。
- `scheduler/`：排程任務實作。

### `test/`

整合或端對端測試（依專案實際內容擴充）。

## 3. 架構優點

- 清楚分層：delivery/service/repository/infra 角色明確，易於分工。
- 啟動模式彈性高：`cmd/` 支援 server、worker、scheduler 多進程運行。
- 介面導向：`pkg/iface/` + `internal/mock/` 提升可測試性與替換性。
- 基礎設施可集中治理：`pkg/infra/` 便於管理連線、套件與第三方整合。
- 配置分環境：多份 `app*.properties` 有利於不同場景部署。

## 4. 架構缺點

- 目錄規模大：模組多時，新人理解成本高。
- 邊界可能模糊：`pkg/` 與 `internal/` 若缺規範，責任容易重疊。
- 單體專案的耦合風險：跨模組直接引用過多時，修改影響面會快速擴大。
- 設定檔數量多：若缺一致命名與註解，容易誤用環境變數或配置。
- delivery 形式多樣（REST/Socket/Worker）時，若缺統一流程規範，行為可能不一致。

## 5. 開發與維護注意事項

- 嚴守分層依賴方向：
  - 建議方向：delivery -> service -> repository/infra。
  - 避免反向依賴（例如 repository 引用 delivery）。
- 介面放置原則：
  - 以「使用者端」定義介面（consumer side interface），避免過度抽象。
- 新功能落點建議：
  - API/事件入口放 `pkg/delivery/`。
  - 業務流程放 `pkg/service/`。
  - DB 存取放 `pkg/repository/`。
  - 外部系統整合放 `pkg/infra/`。
- 設定管理：
  - 新增設定時同步更新對應 `app*.properties` 與文件。
  - 敏感資訊改由環境變數或 secret 管理，不直接寫入版本庫。
- 測試策略：
  - 單元測試優先覆蓋 service/repository。
  - 使用 `internal/mock/` 與 `internal/test_fixture/` 控制測試穩定性。
  - 針對 worker/scheduler 增加排程與重試邏輯測試。
- migration 與部署：
  - SQL migration 檔名與執行順序需固定規範。
  - 部署初始化檔（CSV/SQL）需版本化並可重入（idempotent）。
- 可觀測性：
  - 透過 `pkg/infra/otel/` 與 `pkg/infra/zlog/` 統一追蹤、日誌與錯誤關聯 ID。

## 6. 建議後續優化方向

- 建立「模組依賴規範文件」與簡單檢查機制（CI lint 或架構測試）。
- 在 `docs/` 補充「請求流程圖」與「事件/排程流程圖」。
- 為關鍵模組建立 ADR（Architecture Decision Record），保留設計決策脈絡。
