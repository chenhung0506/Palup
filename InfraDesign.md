假設公司有個訂單系統, 所有使用者的資料都存在MySQL資料庫底下, 請設計一個報表系統提供分析數據, 
並滿足「在資料量大的情況下系統仍可以正常運作」
試著描述出此系統的大致架構與思維

## 思維
- Mysql 做分區 Partition ，如按月存儲，減少單次查詢的數據量，提高性能
- 依使用場景可創建 Index 如 userId + created_date 或 status + created_date 
- 如果是不會常態性修改資料如「使用者資料」放入 Redis 做 cache 減少 Query 的資源消耗 
- 如果要做報表資料，使用 logstash 工具將資料同步至 ELK，ELK 支持橫向擴展，且有多元的查詢語法可供報表系統查詢
- 如要新增訂單與修改訂單等操作，放入 RabbitMQ 做隊列管理，避免瞬間大量 新、刪、修 訂單造成 SQL Lock

## 架構
### 由前到後依次:
- 使用 google recaptcha 與 cloudflare 來防惡意 DDOS
- 使用 Nginx 防同個 IP 頻繁撈資料，並用 Nginx Loadbalance 機制分散 Request 至多台服務
- 使用 React 或 Vue 呈現報表給使用者
- 使用 Gateway server 做服務權限統一控管
- 使用 API（如 REST 或 GraphQL）提供彈性的查詢服務
- 使用 Redis 做 cache 
- 查報表直接對 ELK 查詢，新刪修先排隊列 (RabbitMQ)
