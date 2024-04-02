# Application

## CIDR Calculator

The app's name stands for RDerik's Interactive  CIDR (RDICIDR), this is an CIDR calculator. This project uses golang and its library net/http to serve static files that are under web/public folder.


### Execute the application
```bash
PORT=3000 go run cmd/main/main.go
```

### Testing and Building the application
```bash
# Testing the application
go test ./...

# Download modules
go mod download

# Build
go build ./... # the executable main will be available in the root directory
```

### References
- Golang standard project layout https://github.com/golang-standards/project-layout
