module github.com/joeyak/go-twitch-eventsub

go 1.19

require (
	github.com/google/uuid v1.3.0
	nhooyr.io/websocket v1.8.7
)

require github.com/klauspost/compress v1.10.3 // indirect

retract v0.1.9 // Build error when refactoring type to be exposed
