![discoverus](https://user-images.githubusercontent.com/42062381/146393207-2a596522-0808-437c-97ca-33a500a86520.png)


# Docs

## REST Operations

| **Operation**                       | **HTTP Action**                   | **Return Value** |
|-------------------------------------|-----------------------------------|------------------|
| Register new Instance               | POST /apps/**appId** | 200 OK  |
| De-register application instance    | DELETE /apps/**appId**/**instanceId** | 200 OK |
| Send application instance heartbeat | PUT /apps/**appId**/**instanceId** | 200 OK |
| Query for all appId instances       | GET /apps/**appId** | 200 OK |
| Query for a specific instanceId     | GET /apps/**appId**/**instanceId** | 200 OK |


## How-To Register
Send a POST request to /apps/_appId_ with the following object in json format:
```json
{
  "id": "user"
  "serviceType": "service"
  "ip": "192_144.3.5"
  "port": 87
  "healthCheckUrl": "/hc"
}
```
