# SigmatechTechnicalTest-Golang
1. Please import the example data and table structure at directory sql
2. Setup .env file with your local database configuration
3. Run command ```go run main.go``` or ```docker-compose up -d --build```
4. Request & Response

```
METHOD  GET 
URL     127.0.0.1:8008/v1/auth/login
BODY
        {
            "nik": "1111"
        }
RESPONSE
        {
            "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzkzMjQ0ODcsInVzZXJJRCI6Ijk2NDY1ZjNmLWM2YmMtMTFlZC05MTM4LTAwZmZjMmVmMjY0NSJ9.sjkeTvWzMyTuUC6JfvCU7lVB3CvUN7yXjKAe0K_qXFs",
            "err_response": null,
            "message": "success",
            "meta": null,
            "status": true
        }
```

```
METHOD  POST
URL     127.0.0.1:8008/v1/transaction/buy
HEADER  Authorization   Bearer {token}
BODY
        {
            "tenor": 2,
            "company_asset_id": "a23c00c8-c6ca-11ed-9138-00ffc2ef2645"
        }
RESPONSE
        {
            "data": "",
            "err_response": null,
            "message": "success",
            "meta": null,
            "status": true
        }
```

```
METHOD  GET
URL     127.0.0.1:8008/v1/transaction/list-order
HEADER  Authorization   Bearer {token}
RESPONSE
        {
            "data": [
                {
                    "transaction_id": "c0f78366-c6d6-11ed-9138-00ffc2ef2645",
                    "otr_price": 1000000,
                    "admin_fee": 5000,
                    "tenor": 2,
                    "interest": 100000,
                    "asset_name": "Mobil",
                    "total_price": 1105000
                }
            ],
            "err_response": null,
            "message": "success",
            "meta": null,
            "status": true
        }
```