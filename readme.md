This Golang project provides access to the Canvas API

# Run all Tests:
NOTE!!!!!! This will run against the Canvas Instance you use to generate the token. Create a test account and use that
account id in the .env file.

create a .env file and add a valid Canvas API Token
`CANVAS_API_TOKEN=a valid token`
`CANVAS_ACCOUNT_ID=578`

gotest ./... -v

# Update packages
go get -u

# Build lib:
go mod tidy             // Clean up the mod file
go build ./...
