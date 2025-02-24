Note 1: no other additional third party packages are allowed except the ones which are already included in go.mod.

Note 2: please consider readability, reusability and testability.

Question 1.

Write a middleware that validates an attribute named "api-key" (see line 12 in main.go) in request Header for test.Test1 (line 18).

---

Question 2.

Complete test.Test1 (see line 16 in main.go). This api request handler takes an array named "Array" from data in request Body. Convert this array to a linked list. In each node, a property of interface{} type is used to store the value.

Example
input:
{
"Array": ["a", "b", "c", "d", "e"],
}

output:
head -> a
node1 -> b
node2 -> c
node3 -> d
tail -> e



## 啟動服務
```
cd playsee && go run main.go
```

## 測試 Array to Linked list
```
 curl -X POST http://localhost:8082/test-1 \
      -H "Content-Type: application/json" \
      -H "APIKey: qwerklj1230dsa350123l2k1j4kl1j24" \
      -d '{"Array": ["a", "b", "c", "d", "e"]}'
```

## 預期 Response
```
{
  "Value": "a",
  "Next": {
    "Value": "b",
    "Next": {
      "Value": "c",
      "Next": {
        "Value": "d",
        "Next": {
          "Value": "e",
          "Next": null
        }
      }
    }
  }
}
```