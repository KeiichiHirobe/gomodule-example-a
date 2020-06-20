# gomodule-example-a
This repository just checks the behavior of Minimal Version Selection of GoModules.

see also:
https://research.swtch.com/vgo-mvs

# explanation
Module A (this repository) is top-level module, A depends on B and C. And B depends on C too.

```
        +--> B(v1.0.0) ---+
        |                 |
 A -----+                 |      C(v1.0.0       <--+
        |                 +-------> or v1.1.0   <--+
        |                           or v2.0.0)  <--+
        |                                          |
        +------------------------------------------+
```

You can follow flow below at master commit.

## A requires C1.0.0
```
$ go run .
INFO[0000] Version v1.0.0 load                          
v1.0.0
```

## A requires C1.0.0, B requires C1.1.0
Minimal Version Selection select C1.1.0. 

Automatically C records of go.mod change from v1.0.0 to v1.1.0 when you `go build`.  
If you run  `go mod tidy` , that deletes C1.0.0 records of `go.sum`.

```
$ go run .
INFO[0000] Version v1.1.0 load                          
v1.1.0
v1.1.0
v1.1.0!!!!!!
update by module b
update by module b  // Here, both of A and B references the same C global variable
```

## Add replace manually
Add to go.mod as below.
```
replace github.com/KeiichiHirobe/gomodule-example-c => github.com/KeiichiHirobe/gomodule-example-c v1.0.0
```
Because A is is top-level module, `replace` is valid.
Since B needs function `GetDecolatedVersionConfirm` which is not defined in C1.0.0, `go build` failed.

## A requires C2.0.0, B requires C1.1.0
Since different major versions are just different modules, C2.0.0 and C1.1.0 exists simultaneouly.

```
$ go run .
INFO[0000] Version v1.1.0 load                          
INFO[0000] Version v2.0.0 load                          
v2.0.0
v1.1.0
v1.1.0!!!!!!
update by module b
v2.0.0
```