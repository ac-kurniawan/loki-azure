# Loki Azure

## How to run
1. Create `.env` file, you can refer to `env.example` to know the parameters needed
2. To migrate all model: `make migration`, please add `uuid` extension in your db using this query `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
3. To start apps: `make run` (server should be run in :3222 by default)

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