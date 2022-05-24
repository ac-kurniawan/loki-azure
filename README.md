# Loki Azure

## How to run
1. Make sure that you use go version >= 1.18
2. Create `.env` file, you can refer to `env.example` to know the parameters needed
3. To migrate all model: `make migration`, please add `uuid` extension in your db using this query `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
4. To start apps: `make run` (server should be run in :3222 by default)

## HTTP REST API

### 1. Create event
```curl
curl --location --request POST 'localhost:3222/event' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "lorem",
    "description": "lorem ipsum dolor",
    "isPublished": false
}' 
```
response:
```json
{
    "status": 201,
    "data": {
        "eventId": "13e6ceae-ad95-4413-a34f-1d5850f286c1",
        "name": "lorem",
        "description": "lorem ipsum dolor",
        "isPublished": true,
        "createdAt": "2022-05-18T16:15:45.093973+07:00",
        "updatedAt": "2022-05-18T16:15:45.093973+07:00"
    }
}
```
### 2. Get event by event id
```curl
curl --location --request GET 'localhost:3222/event/13e6ceae-ad95-4413-a34f-1d5850f286c1'
```
response:
```json
{
    "status": 200,
    "data": {
        "eventId": "13e6ceae-ad95-4413-a34f-1d5850f286c1",
        "name": "lorem",
        "description": "lorem ipsum dolor",
        "isPublished": true,
        "createdAt": "2022-05-18T16:15:45.093973+07:00",
        "updatedAt": "2022-05-18T16:15:45.093973+07:00"
    }
}
```
### 3. Create schedule within event
```curl
curl --location --request POST 'localhost:3222/event/13e6ceae-ad95-4413-a34f-1d5850f286c1/schedule' \
--header 'Content-Type: application/json' \
--data-raw '{
    "startTime": "2021-02-18T21:54:42.123Z",
    "location": "here",
    "basePrice": 100000,
    "quota": 100
}'
```
response:
```json
{
    "status": 201,
    "data": {
        "scheduleId": "86cbb019-b69f-4c6c-be4e-34d307822a2e",
        "startTime": "2021-02-18T21:54:42.123Z",
        "endTime": null,
        "location": "here",
        "basePrice": 100000,
        "quota": 100,
        "booked": 0
    }
}
```
### 4. Get schedules by event id
```curl
curl --location --request GET 'localhost:3222/event/13e6ceae-ad95-4413-a34f-1d5850f286c1/schedule'
```
response:
```json
{
    "status": 200,
    "data": [
        {
            "scheduleId": "d239e068-ed98-42cf-b37a-928b65914edb",
            "startTime": "2021-02-19T04:54:42.123+07:00",
            "endTime": null,
            "location": "here",
            "basePrice": 100000,
            "quota": 100,
            "booked": 0
        },
        {
            "scheduleId": "b4aa4ee7-3235-45fa-9d35-12f4a9820f21",
            "startTime": "2021-02-19T04:54:42.123+07:00",
            "endTime": null,
            "location": "here",
            "basePrice": 100000,
            "quota": 100,
            "booked": 0
        }
    ]
}
```

### 5. Create order
```curl
curl --location --request POST 'localhost:3222/order' \
--header 'Content-Type: application/json' \
--data-raw '{
    "phoneNumber": "628000000001",
    "email": "asd@asd.id",
    "scheduleId": "855c50a3-7483-4783-b526-ecb0064d9880",
    "qty": 2
}'
```
response:
```json
{
    "status": 201,
    "data": {
        "orderId": "bbf3c675-8508-4cd1-942c-88c40bdd3633",
        "phoneNumber": "628000000001",
        "email": "asd@asd.id",
        "status": "WAITING_FOR_PAYMENT",
        "scheduleId": "855c50a3-7483-4783-b526-ecb0064d9880",
        "qty": 2,
        "createdAt": "2022-05-24T20:25:31.473821+07:00",
        "updateAt": "2022-05-24T20:25:31.473821+07:00"
    }
}
```

### 6. Get order by id
```curl
curl --location --request GET 'localhost:3222/order/7713da7b-ea3e-43e0-8941-e5cd68a45329'
```

response:
```json
{
    "status": 200,
    "data": {
        "orderId": "7713da7b-ea3e-43e0-8941-e5cd68a45329",
        "phoneNumber": "628000000001",
        "email": "asd@asd.id",
        "status": "WAITING_FOR_PAYMENT",
        "scheduleId": "363ff703-4b3b-4ec0-bd45-32e56df4e663",
        "qty": 2,
        "createdAt": "2022-05-24T20:08:29.917114+07:00",
        "updateAt": "2022-05-24T20:08:29.917114+07:00"
    }
}
```

### 7. Checkout order
```curl
curl --location --request POST 'localhost:3222/order/checkout' \
--header 'Content-Type: application/json' \
--data-raw '{
    "orderId": "bbf3c675-8508-4cd1-942c-88c40bdd3633"
}'
```

response:
```json
{
    "status": 201,
    "data": {
        "orderId": "bbf3c675-8508-4cd1-942c-88c40bdd3633",
        "phoneNumber": "628000000001",
        "email": "asd@asd.id",
        "status": "SUCCESS",
        "scheduleId": "855c50a3-7483-4783-b526-ecb0064d9880",
        "qty": 2,
        "createdAt": "2022-05-24T20:25:31.473821+07:00",
        "updateAt": "2022-05-24T20:25:31.473821+07:00"
    }
}
```