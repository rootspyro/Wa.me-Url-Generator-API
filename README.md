# Wa.me Url Generator API

Restful API to dynamically generate whatsapp api urls with a custom message.

## Technologies
This API is builded with the [gin gonic](https://gin-gonic.com/) framework.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![WhatsApp](https://img.shields.io/badge/WhatsApp-25D366?style=for-the-badge&logo=whatsapp&logoColor=white)

## Install and Run the Project
### 1 - Install and configuration
```shell
# Clone the repo
git clone https://github.com/rootspyro/Wa.me-Url-Generator-API.git
cd Wa.me-Url-Generator-API

# Install modules
go mod tidy
```
### 2 - Run the project
```shell
# Run the project
go run cmd/main.go
```

### 3 - Ping Test
```shell
curl http://localhost:{port}/api/ping
# Response: {"status":"success", "data": "pong"}
```

### 4 - First Url Generation
```shell
curl -XPOST -H "Content-type: application/json" -d '{
"phone":"+12225556666",
"msg":"Hello World!"
}' 'http://localhost:3000/ws/url'

# Response: {"status":"success", "data":{"contact-phone":"+12225556666","msg":"Hello World!","url":"https://wa.me/+12225556666?text=Hello%20World!"}}
# Generated Url: https://wa.me/+12225556666?text=Hello%20World!
```
