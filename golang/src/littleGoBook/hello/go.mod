module cmon.consulting/go/hello

go 1.16

require (
	cmon.consulting/go/greetings v0.0.0-00010101000000-000000000000
	golang.org/x/text v0.3.3 // indirect
	rsc.io/quote v1.5.2
)

replace cmon.consulting/go/greetings => ../greetings

replace example.com/greetings => ../greetings
