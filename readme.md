# Setup

run `go mod init myapp && go mod tidy` in project directory

run `task`

go to `http://127.0.0.1:8090/_/` and make an admin login.

create a `count` collection with a single `count` column. Add a record to this collection and copy the `id` for this row.

paste it as the `countId` in `main.go`

go to `http://127.0.0.1:8090` to see if it's working!
