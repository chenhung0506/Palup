## 設計一張資料表 並撰寫 sql 找出第一次登入後7天內還有登入的使用者 ##
例如：3/10第一次登入，3/12有再登入，滿足第一次登入後7天內還有登入

   - 任何 sql 語言回答皆可 
   - 簡單描述語法邏輯
   - 答案請提供 schema (column, type) 與 sql 



## Table schema
```
CREATE TABLE user_logins (
    user_id INT NOT NULL,
    login_time DATETIME NOT NULL,
    PRIMARY KEY (user_id, login_time), -- 避免同一使用者在同一時間紀錄重複數據
    INDEX idx_user_id (user_id)      -- 為 user_id 建索引
);

```

## Query SQL
```
SELECT 
    user_id
FROM 
    user_logins
GROUP BY 
    user_id
HAVING 
    DATEDIFF(MAX(login_time), MIN(login_time)) > 7;
```

## Explain
- 使用 GROUP BY user_id，將每位使用者的紀錄分組
- 使用 MIN 與 MAX 函數找到第一次與最後一次登入時間
- 使用 HAVING 條件篩選出大於 7 天的使用者