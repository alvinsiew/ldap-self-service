# ldap-self-service

## Self-compile

```bash
# MacOS
env GOOS=darwin GOARCH=amd64 go build -o bin/ldapss cmd/ldapss/main.go

# Linux
env GOOS=linux GOARCH=amd64 go build -o bin/ldapss cmd/ldapss/main.go
```
