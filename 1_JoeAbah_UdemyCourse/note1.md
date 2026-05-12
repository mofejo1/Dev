# Study Note

- `go build "name to build " `, this is will convert to an executable file
- then you can run with `./name to build`
- while `go run` will complie the program into a temporary directory and run the program.
- you can also use `go fmt name of program` this wil help you format
- `go test` for testing
- `go get & go install` for installing an entire package
- `http://golangci-lint.run/`
  `brew install golangci-lint` and
  `brew upgrade golangci-lint`
  golangci-lint is a fast, all-in-one linting tool for Go that helps catch bugs, style issues, and vulnerabilities in your code before they ship.
- now we go here `https://github.com/golang-migrate/migrate`
  copy `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
  golang-migrate is a tool that allows you to manage and automate your database schema changes through versioned "up" and "down" scripts, either as a standalone CLI or as a library within your Go applications.
  so after this , if you type `migrate` on your terminal you will see that migration is installed
- next we go here https://github.com/swaggo/swag where we run `go install github.com/swaggo/swag/cmd/swag@latest` from the 'getting started ' section
  it allows us to automatically generate and host interactive, production-ready API documentation directly from your code comments, ensuring your manual is always accurate and testable without extra effort.
- next is `gqlgen.com` run `go install github.com/99designs/gqlgen@latest`
  gqlgen is a library for building GraphQL servers in Go that uses a "schema-first" approach to automatically generate type-safe code, ensuring your API implementation always stays perfectly in sync with your data definitions.
- To see all the Go programs (binaries) you have installed in your Go Bin folder, you can run this command in your terminal: `ls $(go env GOPATH)/bin`
- next you download `Goland from jetbrain`instal , thne go to `settings - Go - Go modules - always make usre thr "enable go module is set"`

- next you go to pgAdmin.org, and download the app for your system type
