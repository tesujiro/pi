go:
	go run pi_Chudnovsky.go

build:
	go build pi_Chudnovsky.go

prof:
	go tool pprof --svg ./pi_Chudnovsky ./cpu.pprof > ./cpu.svg

bench:
	go test -bench . -benchmem -test.timeout 60m

