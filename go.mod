require (
	gotrading/app/controllers v0.0.0-00010101000000-000000000000
	gotrading/app/models v1.2.3
	gotrading/config v1.2.3
	gotrading/utils v1.2.3
)

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gotrading/bitflyer v1.2.3 // indirect
)

replace (
	gotrading/app/controllers => ./app/controllers
	gotrading/app/models => ./app/models
	gotrading/bitflyer => ./bitflyer
	gotrading/config => ./config
	gotrading/utils => ./utils
)

module gotrading

go 1.19
