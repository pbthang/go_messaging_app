#!/bin/bash

URL_SEND="http://localhost:8080/api/send"
URL_PULL="http://localhost:8080/api/pull"
JSON_SEND='{"Chat": "john:jane", "Text": "Lorem ipsum dolor ...", "Sender": "john"}'
JSON_PULL='{"Chat": "john:jane", "Cursor": 0, "Limit": 20, "Reverse": false}'
METHOD_POST="POST"
METHOD_GET="GET"
CONCURRENCY=20
REQUESTS=50
LOG_FILE_SEND="stress_test_send.log"
LOG_FILE_PULL="stress_test_pull.log"

if [ ! -z "$1" ]; then
    CONCURRENCY=$1
fi
if [ ! -z "$2" ]; then
    REQUESTS=$2
fi


function send_request() {
    local url=$1
    local data=$2
    local method=$3
    local log_file=$4
    curl -s -X "$method" -H "Accept: application/json" -H "Content-Type: application/json" -d "$data" "$url" -o dev/null -w "Status %{http_code}, took %{time_total}s\n" >> "$log_file"
}

# Start time
start_time=$(date +%s)
# clear log files
date +"%F %T%Z" > "$LOG_FILE_SEND"
date +"%F %T%Z" > "$LOG_FILE_PULL"

# Run concurrent requests for POST
for ((i = 1; i <= CONCURRENCY; i++)); do
    {
        for ((j = 1; j <= REQUESTS; j++)); do
            send_request "$URL_SEND" "$JSON_SEND" "$METHOD_POST" "$LOG_FILE_SEND"
        done
    } &
done

# Run concurrent requests for GET
for ((i = 1; i <= CONCURRENCY; i++)); do
    {
        for ((j = 1; j <= REQUESTS; j++)); do
            send_request "$URL_PULL" "$JSON_PULL" "$METHOD_GET" "$LOG_FILE_PULL"
        done
    } &
done

# Wait for all child processes to finish
wait

# End time
end_time=$(date +%s)

# Calculate duration
duration=$((end_time - start_time))

echo "Test completed."
echo "Total requests: $((CONCURRENCY * REQUESTS))"
echo "Concurrency level: $CONCURRENCY"
echo "Total duration: $duration seconds"
