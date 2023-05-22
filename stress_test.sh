#!/bin/bash

URL="http://localhost:8080/api/send"
JSON_DATA='{"Chat": "john:jane", "Text": "Lorem ipsum dolor ...", "Sender": "john"}'
CONCURRENCY=$1
REQUESTS=$2

if [ -z "$CONCURRENCY" ]; then
    CONCURRENCY=20
fi

if [ -z "$REQUESTS" ]; then
    REQUESTS=50
fi

function send_request() {
    local url=$1
    local data=$2
    curl -s -X POST -H "Content-Type: application/json" -d "$data" "$url" -o /dev/null -w "%{http_code}\n"
}

# Start time
start_time=$(date +%s)

# Run concurrent requests
for ((i = 1; i <= CONCURRENCY; i++)); do
    {
        for ((j = 1; j <= REQUESTS; j++)); do
            send_request "$URL" "$JSON_DATA"
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
