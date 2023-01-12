module github.com/joeyak/go-twitch-eventsub

go 1.19

require (
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.8.1
	nhooyr.io/websocket v1.8.7
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/klauspost/compress v1.10.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract v0.1.9 // Build error when refactoring type to be exposed
