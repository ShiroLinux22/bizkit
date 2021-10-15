module github.com/chakernet/ryuko/info

go 1.17

require (
	github.com/chakernet/ryuko/common v0.0.0-00010101000000-000000000000
	github.com/diamondburned/arikawa/v3 v3.0.0-rc.2
	github.com/joho/godotenv v1.4.0
	github.com/streadway/amqp v1.0.0
)

require (
	github.com/gorilla/schema v1.2.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
)

replace github.com/chakernet/ryuko/common => ../common
