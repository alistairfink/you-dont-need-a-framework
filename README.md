# you-dont-need-a-framework

Example project to show that you don't need an HTTP framework for go now.

This project also has examples of how to implement both pre- and post-request middlewares.

## Examples
### `200 OK` Response

#### With response code logging
```shell
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"word":"elephant"}' \
  http://localhost:4000/with_response_code
{"message":"Hello elephant"}
```
Logs:
```json
{
  "time": "2025-04-29T17:19:05.53114356+09:00",
  "level": "INFO",
  "msg": "without error",
  "context": {
    "Method": "POST",
    "URL": "/with_response_code",
    "Headers": {
      "Accept": [
        "*/*"
      ],
      "Content-Length": [
        "19"
      ],
      "Content-Type": [
        "application/json"
      ],
      "User-Agent": [
        "curl/7.81.0"
      ]
    },
    "User-Agent": "curl/7.81.0"
  }
}
{
  "time": "2025-04-29T17:19:05.53126029+09:00",
  "level": "INFO",
  "msg": "finished request",
  "context": {
    "Method": "POST",
    "URL": "/with_response_code",
    "Headers": {
      "Accept": [
        "*/*"
      ],
      "Content-Length": [
        "19"
      ],
      "Content-Type": [
        "application/json"
      ],
      "User-Agent": [
        "curl/7.81.0"
      ]
    },
    "User-Agent": "curl/7.81.0"
  },
  "status code": 200
}
```

#### Without response code logging
```shell
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"word":"elephant"}' \
  http://localhost:4000/without_response_code
{"message":"Hello elephant"}
```
Logs:
```json
{
  "time": "2025-04-29T17:20:06.39698949+09:00",
  "level": "INFO",
  "msg": "without error",
  "context": {
    "Method": "POST",
    "URL": "/without_response_code",
    "Headers": {
      "Accept": [
        "*/*"
      ],
      "Content-Length": [
        "19"
      ],
      "Content-Type": [
        "application/json"
      ],
      "User-Agent": [
        "curl/7.81.0"
      ]
    },
    "User-Agent": "curl/7.81.0"
  }
}
```

### `400 Bad Request` Response
#### With response code logging
```shell
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"word":"elephant"}' \
  http://localhost:4000/with_response_code_error
Invalid request
```
Logs:
```json
{
  "time": "2025-04-29T17:21:43.9001143+09:00",
  "level": "INFO",
  "msg": "with error",
  "context": {
    "Method": "POST",
    "URL": "/with_response_code_error",
    "Headers": {
      "Accept": [
        "*/*"
      ],
      "Content-Length": [
        "19"
      ],
      "Content-Type": [
        "application/json"
      ],
      "User-Agent": [
        "curl/7.81.0"
      ]
    },
    "User-Agent": "curl/7.81.0"
  }
}
{
  "time": "2025-04-29T17:21:43.90016047+09:00",
  "level": "INFO",
  "msg": "finished request",
  "context": {
    "Method": "POST",
    "URL": "/with_response_code_error",
    "Headers": {
      "Accept": [
        "*/*"
      ],
      "Content-Length": [
        "19"
      ],
      "Content-Type": [
        "application/json"
      ],
      "User-Agent": [
        "curl/7.81.0"
      ]
    },
    "User-Agent": "curl/7.81.0"
  },
  "status code": 400
}
```

#### Without response code logging
```shell
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"word":"elephant"}' \
  http://localhost:4000/without_response_code_error
Invalid request
```
Logs:
```json
{
  "time": "2025-04-29T17:22:31.396957542+09:00",
  "level": "INFO",
  "msg": "with error",
  "context": {
    "Method": "POST",
    "URL": "/without_response_code_error",
    "Headers": {
      "Accept": [
        "*/*"
      ],
      "Content-Length": [
        "19"
      ],
      "Content-Type": [
        "application/json"
      ],
      "User-Agent": [
        "curl/7.81.0"
      ]
    },
    "User-Agent": "curl/7.81.0"
  }
}
```

## Licence
As per the [LICENSE](LICENSE) file, you can do whatever you want with this.
I don't care, just don't bother me about it.