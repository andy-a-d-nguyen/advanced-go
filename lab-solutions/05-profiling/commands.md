* Run bench tests
    * `go test -bench .`

* Run bench test with memory stats
    * `go test -bench . -benchmem`

* Memory profile
    * `go test -bench . -memprofile mem.profile -memprofilerate=1`
    * `go tool pprof -web mem.profile`
    * largest consumer - CSV reader

* CPU profile
    * `go test -bench . -cpuprofile cpu.profile`
    * `go tool pprof -web cpu.profile`
    * largest consumer - createRecords (sending messages to channels)

* Blocking profile
    * `go test -bench . -blockprofile block.profile -blockprofilerate=1`
    * `go tool pprof -web block.profile`
    * largest block - receiving messages in `processDay`
        * why? likely due to need to accumulate all records before processing to ensure that each days' data is fully present
