# ethereum_blockchain_services

This repository contains 2 ethereum services interact with ethereum RPC service

## 1 API Service
### 1 Get last n blocks [implemented with goroutine]
```json==
request method: GET
request uri: /blocks
request header: {
  'Content-Type': 'application/json'
}
query paramter: {
  limit: 1
}
response header: 200 OK
response body: {
  "blocks": [{
    "block_num": 1,
    "block_hash": "",
    "block_time": 12356789,
    "parent_hash": ""
  }]
}
```
### get block by id [implemented with just jsonrpc]
```json==
request method: GET
request uri: /blocks
request header: {
  'Content-Type': 'application/json'
}
path parameter: {
  id: $id_value
}
response header: 200 OK
response body: {
  {
    "block_num": 1,
    "block_hash": "",
    "block_time": 12356789,
    "parent_hash": "",
    "transactions": [
      "0x12345678"
      "0x87654321"
    ]
  }
}
```
### get transaction data with event logs [implemented with just jsonrpc]
```json==
request method: GET
request uri: /blocks
request header: {
  'Content-Type': 'application/json'
}
path paramter: {
  txHash: $txHash_value
}
response header: 200 OK
response body: {
  "tx_hash": "0x666",
  "from": "0x4321",
  "to": "0x1234",
  "nonce": 1,
  "data": "0xeb12",
  "value": "12345678",
  "logs": [
    {
      "index": 0,
      "data": "0x12345678"
    }
  ]
}
```
## 2 Ethereum block indexer service[doing]
   
  根據 web3 API 透過 RPC 將區塊內的資料掃進 db

## test
```shell=
go test -v ./tests/
```