require (
	example.com/bitflyer v1.2.3
	example.com/config v1.2.3
	example.com/utils v1.2.3
)

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)

replace (
	example.com/bitflyer => ./bitflyer
	example.com/config => ./config
	example.com/utils => ./utils
)

module gotrading

go 1.19
