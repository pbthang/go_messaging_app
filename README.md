# assignment_demo_2023

![Tests](https://github.com/pbthang/go_messaging_app/actions/workflows/test.yml/badge.svg)

This is my implementation for backend assignment of 2023 TikTok Tech Immersion.

Requirements: [https://bytedance.sg.feishu.cn/docx/P9kQdDkh5oqG37xVm5slN1Mrgle](https://bytedance.sg.feishu.cn/docx/P9kQdDkh5oqG37xVm5slN1Mrgle)

## How to run with docker-compose

Make sure you have docker and docker-compose installed. Run the following command to start the app:
```bash
docker-compose up
```

### Delete all data

Execute the following command to connect to the running redis container:
```bash
docker exec -it $(docker ps | grep redis | awk '{print $1;}') redis-cli
```

Then execute the following command to delete all data:
```bash
FLUSHALL
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

Expected Response: status `200`

```json
{
  "messages": [
    {
      "chat": "user1:user2",
      "text": "Lorem ipsum dolor ...",
      "sender": "john",
      "send_time": 1684744610
    }, ...
  ]
}
```

## Testing

### Unit test

Execute the following command in `./rpc-server` folder to run unit tests and get test coverage:

```bash
 go test -race -cover -coverprofile=coverage.out $(go list ./... | grep -Ev "_gen") -coverpkg $(go list ./... | grep -Ev "_gen" | tr "\n" "," | sed 's/.$//')
```

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

