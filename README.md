# Remote Management API

A Go-based REST API for remote power management of servers and computers using WS-Man protocol. This service allows you to control power states (on, off, cycle) and query power status of remote hosts over the network.

## Features

- **Power Management**: Turn servers/computers on, off, or cycle (reboot)
- **Status Query**: Check current power state of remote hosts
- **RESTful API**: Simple HTTP endpoints for integration
- **Dual Interface**: Support for both GET (URL parameters) and POST (JSON body) requests
- **WS-Man Integration**: Uses `wsman` command-line tool for WS-Management protocol communication

## Prerequisites

1. **Go 1.25.4 or later**: [Install Go](https://golang.org/dl/)
2. **wsman CLI tool**: Required for WS-Management protocol communication
   - On Ubuntu/Debian: `sudo apt-get install wsmancli`
   - On RHEL/CentOS: `sudo yum install wsmancli`
   - On macOS: `brew install wsmancli`
3. **Network Access**: The target hosts must be accessible on the network and support WS-Management protocol (typically enabled on servers with IPMI/BMC)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/isubhampadhi56/remote-management.git
   cd remote-management
   ```

2. Build the application:
   ```bash
   go build ./cmd/main.go
   ```

3. Run the server:
   ```bash
   ./main
   ```
   Or run directly:
   ```bash
   go run ./cmd/main.go
   ```

The server will start on port 8080 by default.

## API Endpoints

All endpoints are prefixed with `/power/`.

### Power On
Turn on a remote host.

**POST `/power/on`**
```json
{
  "host": "192.168.10.100"
}
```

**GET `/power/on/{host}`**
```
GET /power/on/192.168.10.100
```

**Response**: HTTP 200 OK on success

### Power Off
Turn off a remote host.

**POST `/power/off`**
```json
{
  "host": "192.168.10.100"
}
```

**GET `/power/off/{host}`**
```
GET /power/off/192.168.10.100
```

**Response**: HTTP 200 OK on success

### Power Cycle
Reboot a remote host (off then on).

**POST `/power/cycle`**
```json
{
  "host": "192.168.10.100"
}
```

**GET `/power/cycle/{host}`**
```
GET /power/cycle/192.168.10.100
```

**Response**: HTTP 200 OK on success

### Power Status
Query the current power state of a remote host.

**POST `/power/status`**
```json
{
  "host": "192.168.10.100"
}
```

**GET `/power/status/{host}`**
```
GET /power/status/192.168.10.100
```

**Response**:
```json
{
  "power_state": "on"
}
```

Possible `power_state` values:
- `"on"` - Power is on
- `"off"` - Power is off  
- `"sleep"` - In sleep mode
- `"hibernate"` - In hibernate mode
- `"reset"` - In reset state
- `"unknown"` - Unknown state or error querying

## Configuration

The API uses default credentials and port for WS-Management:
- **Port**: 623 (standard IPMI/BMC port)
- **Username**: Administrator
- **Password**: Realtek

To modify these defaults, edit the `manager()` function in `router/power_router.go`:

```go
func manager(host string) *power.Manager {
	return &power.Manager{
		Host:     host,
		Port:     "623",           // Change port here
		Username: "Administrator", // Change username here
		Password: "Realtek",       // Change password here
	}
}
```

## Usage Examples

### Using cURL

**Check power status:**
```bash
# GET request
curl "http://localhost:8080/power/status/192.168.10.100"

# POST request  
curl -X POST "http://localhost:8080/power/status" \
  -H "Content-Type: application/json" \
  -d '{"host":"192.168.10.100"}'
```

**Power on a server:**
```bash
# GET request
curl "http://localhost:8080/power/on/192.168.10.100"

# POST request
curl -X POST "http://localhost:8080/power/on" \
  -H "Content-Type: application/json" \
  -d '{"host":"192.168.10.100"}'
```

**Power off a server:**
```bash
# GET request
curl "http://localhost:8080/power/off/192.168.10.100"

# POST request
curl -X POST "http://localhost:8080/power/off" \
  -H "Content-Type: application/json" \
  -d '{"host":"192.168.10.100"}'
```

### Using Python

```python
import requests

base_url = "http://localhost:8080/power"

# Check status
response = requests.get(f"{base_url}/status/192.168.10.100")
print(response.json())

# Power on
response = requests.post(f"{base_url}/on", json={"host": "192.168.10.100"})
print(response.status_code)
```

## Project Structure

```
remote-management/
├── cmd/
│   └── main.go              # Application entry point
├── pkg/
│   └── power/
│       ├── manager.go       # Manager struct definition
│       ├── power_action.go  # Power control logic
│       ├── power_state.go   # Power state constants
│       └── soap.go          # WS-Man command execution
├── router/
│   ├── power_router.go      # Power management routes
│   └── router.go           # Main router configuration
├── go.mod                   # Go module dependencies
└── README.md               # This file
```

## Dependencies

- [github.com/go-chi/chi/v5](https://github.com/go-chi/chi) - Lightweight, idiomatic and composable router for building Go HTTP services

## Error Handling

The API returns appropriate HTTP status codes:

- **200 OK**: Request successful
- **400 Bad Request**: Missing or invalid parameters (e.g., "host required")
- **404 Not Found**: Invalid endpoint
- **500 Internal Server Error**: WS-Man command execution failed

## Troubleshooting

1. **"wsman: command not found"**
   - Install the wsmancli package for your operating system
   - Verify installation with `which wsman`

2. **Connection refused or timeout**
   - Ensure the target host is powered on and accessible
   - Verify network connectivity to the target host
   - Check if WS-Management service is running on the target (port 623 typically)

3. **Authentication errors**
   - Verify the username and password for the target host
   - Check if the credentials have power management privileges

4. **"host required" error**
   - Ensure you're providing the host parameter correctly
   - For GET requests: include host in URL path (`/power/status/{host}`)
   - For POST requests: include host in JSON body (`{"host": "192.168.10.100"}`)

## Development

### Building from Source
```bash
go build -o remote-management ./cmd/main.go
```

### Running Tests
(Add test instructions when tests are implemented)

### Code Formatting
```bash
go fmt ./...
```

## License

[Add license information here]

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## Support

For issues, questions, or feature requests, please open an issue on the GitHub repository.