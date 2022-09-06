module real-time-forum

go 1.19

require (
	github.com/mattn/go-sqlite3 v1.14.15
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90
)

require (
	github.com/gorilla/websocket v1.5.0
	github.com/satori/go.uuid v1.2.0
)

require gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect

replace github.com/satori/go.uuid v1.2.0 => github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
