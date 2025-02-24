# 給定一組資料舉例如下

|  post_id  |  user_id  |   lat    |    lon    | created_at |
------------|-----------|----------|-----------|------------|
| post_id_0 | user_id_3 |23.6468392|120.5358431|1616479608  |
| post_id_1 | user_id_1 |22.7344496|120.2845859|1616479408  |
| post_id_2 | user_id_3 |21.6468376|121.6538431|1616589608  |
| ...       | ...       |...       |...        |...         |


### NoSQL DB 優勢在於海量資料存取，速度快且成本低，雖然不像SQL DB可以下語法去拉出資料，但合理的rowkey設計可以做到預先準備好類SQL的statement效果，也能發揮NoSQL DB的最大效能

### (例如rowkey設計為post_id#user_id，則可以快速找出特定post_id的user_id是什麼)

## 問題A
設計一個NoSQL DB的rowkey，並說明設計原因，滿足
   - 找出某個user的post
   - 可由新到舊且由舊到新查找
   - 依照NoSQL DB特性，避免hotspot產生


## 問題B
設計一個NoSQL DB的rowkey，並說明設計原因，滿足
   - 在某個latlngbounds時，能快速找出結果
   - 依照NoSQL DB特性，避免hotspot產生
   
## 問題 A 解題
rowkey 設計
```
user_id + created_at
```
- user_id：作為 shard key，可快速查找用戶
- created_at：使用創建時間來確保按時間排序
- 使用複合片鍵 user_id + created_at 來避免 hotspot 

### 測試流程 
#### step 1 創建 table
```
db.createCollection("posts")

db.posts.insertMany([
  { post_id: "post_id_0", user_id: "user_id_3", lat: 23.6468392, lon: 120.5358431, created_at: 1616479608 },
  { post_id: "post_id_1", user_id: "user_id_1", lat: 22.7344496, lon: 120.2845859, created_at: 1616479408 },
  { post_id: "post_id_2", user_id: "user_id_3", lat: 21.6468376, lon: 121.6538431, created_at: 1616589608 }
])
```
#### step 2 創建 Index , created_at:-1 使用新到舊排序
```
db.posts.createIndex({ user_id: 1, created_at: -1})
```
#### step 3 使用 rowkey 查詢某個 UserPost & 由新到舊查資料
```
db.posts.find({ user_id: "user_id_3" }).sort({ created_at: -1 })
```









## 問題 B 解題
- rowkey 設計
```
2dsphere + created_at
```
- 建立 2dsphere + created_at 的複合索引
- 使用 2dsphere + created_at 作為 shard key 避免 hotspot


### Step 1 創建 table
```
db.createCollection("posts")

db.posts.insertMany([
  { post_id: "post_id_6", user_id: "user_id_3", created_at: 1616479608, location: { type: "Point", coordinates: [120.5358431, 23.6468392] } },
  { post_id: "post_id_7", user_id: "user_id_1", created_at: 1616479408, location: { type: "Point", coordinates: [120.2845859, 22.7344496] } },
  { post_id: "post_id_8", user_id: "user_id_3", created_at: 1616589608, location: { type: "Point", coordinates: [121.6538431, 21.6468376] } }
])
```

### Step 2 創建 Index 
```
db.posts.createIndex([("location", "2dsphere"), ("created_at", -1)])
```

### Step 3 查詢資料
```
db.posts.find({
  location: {
    $geoWithin: {
      $box: [[105.0, 21.0], [122.0, 24.0]]
    }
  }
}).sort({ created_at: -1 })
```