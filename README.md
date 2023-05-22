# assignment_demo_2023

![Tests](https://github.com/pbthang/go_messaging_app/actions/workflows/test.yml/badge.svg)

This is my implementation for backend assignment of 2023 TikTok Tech Immersion.

Requirements: [https://bytedance.sg.feishu.cn/docx/P9kQdDkh5oqG37xVm5slN1Mrgle](https://bytedance.sg.feishu.cn/docx/P9kQdDkh5oqG37xVm5slN1Mrgle)

## How to run with docker-compose

Make sure you have docker and docker-compose installed. Run the following command to start the app:
```bash
docker-compose up
```
## API Documentation

### Ping

Check if the server is running

```bash
curl -X GET http://localhost:8080/ping
```

Expected response: status `200`
```json
{
    "message": "pong"
}
```

### Send Message

Send a message to a chat

```bash
curl -X POST \
  http://localhost:8080/api/send \
  -H 'Content-Type: application/json' \
  -d '{
    "Chat": "user1:user2",
    "Text": "Hello World",
    "Sender": "user1"
}'
```

Expected response: status `201`

```text
ok
```

### Get Messages

Get messages from a chat from `Cursor` with `Limit` messages and sorting order `Reverse` (default: `false`)

```bash
curl -X GET \
  http://localhost:8080/api/pull \
  -H 'Content-Type: application/json' \
  -d '{
    "Chat": "user1:user2",
    "Cursor": 0,
    "Limit": 20,
    "Reverse": false
}'
```

Expected Response:

```json
{
  "messages": [
    {
      "chat": "user1:user2",
      "text": "Lorem ipsum dolor ...",
      "sender": "john",
      "send_time": 1684744610
    }
  ]
}
```

## Testing

### Stress test

Use the script `stress_test.sh` to run a stress test with `<concurrent>` concurrent requests and `<requests>` request batches.

```bash
./stress_test.sh <concurrent> <requests>
```

## Tech Stack

- go
- kitex
- redis
- docker
- github actions

