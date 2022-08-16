invoke_url='https://integrate.api.nvidia.com/v1/chat/completions'

authorization_header='Authorization: Bearer $API_KEY_REQUIRED_IF_EXECUTING_OUTSIDE_NGC'
accept_header='Accept: application/json'
content_type_header='Content-Type: application/json'

data=$'{
  "messages": [
    {
      "role": "user",
      "content": "Write a limerick about the wonders of GPU computing."
    }
  ],
  "stream": true,
  "model": "meta/llama-3.1-405b-instruct",
  "max_tokens": 1024,
  "presence_penalty": 0,
  "frequency_penalty": 0,
  "top_p": 0.7,
  "temperature": 0.2
}'

response=$(curl --silent -i -w "\n%{http_code}" --request POST \
  --url "$invoke_url" \
  --header "$authorization_header" \
  --header "$accept_header" \
  --header "$content_type_header" \
  --data "$data"
)

echo "$response"