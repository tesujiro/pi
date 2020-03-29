go:
	go run pi_RamanujanBig.go

build:
	go build pi_RamanujanBig.go

prof:
	go tool pprof --svg ./pi_RamanujanBig ./cpu.pprof > ./cpu.svg

bench:
	go test -bench . -benchmem -test.timeout 60m

