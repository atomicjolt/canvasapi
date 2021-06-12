This go project provides access to the Canvas API

# Build:
go build

go build --race
# Run all Tests:
create a .env file and add a valid Canvas API Token
`CANVAS_API_TOKEN=a valid token`

gotest ./... -v

# Find bad stuff
go vet

# Run benchmarks:
go test -bench=. -benchmem

# Use the profiler
Install newer pprof
  `go get github.com/google/pprof`

and GraphViz
  `brew install graphviz`

Generate the profile
  `go test -bench=. -benchmem -cpuprofile=cpu.pb.gz`

Open the profiler
  `go tool pprof cpu.pb.gz`

Once you are in the profiler run `top` to show the lines of code run the most often
`web` show a graph with all the functions called and the time we spent on each. It should load in graphviz

# Update packages
go get -u

# Build lib:
go mod tidy             // Clean up the mod file
go mod vendor           // Download all modules into the vendor directory
go build
git push
