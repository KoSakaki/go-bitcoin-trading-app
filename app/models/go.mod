module gotrading/app/models

require (
	github.com/mattn/go-sqlite3 v1.14.15
	gotrading/bitflyer v1.2.3
	gotrading/config v1.2.3
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	gotrading/bitflyer => ../../bitflyer
	gotrading/config => ../../config
)

go 1.19
