module gotrading/app/controllers

require (
	gotrading/app/models v1.2.3
	gotrading/bitflyer v1.2.3
	gotrading/config v1.2.3
)

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)

replace (
	gotrading/app/models => ../models
	gotrading/bitflyer => ../../bitflyer
	gotrading/config => ../../config
)

go 1.19
