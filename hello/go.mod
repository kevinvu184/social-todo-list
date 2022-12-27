module github.com/kevinvu184/go-tutorials/hello

go 1.19

require (
	github.com/kevinvu184/go-tutorials/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.5.0 // indirect
	rsc.io/sampler v1.99.99 // indirect
)

replace github.com/kevinvu184/go-tutorials/greetings => ../greetings
