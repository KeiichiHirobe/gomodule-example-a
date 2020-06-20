module github.com/KeiichiHirobe/gomodule-example-a

go 1.14

require (
	github.com/KeiichiHirobe/gomodule-example-b v1.0.0
	github.com/KeiichiHirobe/gomodule-example-c v1.1.0
)

replace github.com/KeiichiHirobe/gomodule-example-c => github.com/KeiichiHirobe/gomodule-example-c v1.0.0