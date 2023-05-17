## GO fundamentals

- fundamentals
- standard libraries
- Go Routines
- fmt Package
- http Package
- Templating
- Web Servers & Services

### Run Go Code

- create go file `index.go`
- run command from terminal

```bash
go run <filename>.go
```

### Create Modules

- create module
```bash
go mod init <name>
```

- run the code
```bash
go run .
```

### Workspaces and CLI

```bash
go work init
```

### Variables

```go

//
// !variables are nil by default
//
var x int // syntax: var var_name type
var name string

// constants
const z int  = 2

//
// !string uses double quotes and backticks ``
//
var text string 
text = "hello!" // assigning values to a variable

// variable initialization shortcut
otherText := "bye!"

```

### DataTypes

- some common data types in GO
```go
string
int int8 int16 int32 int64 
uint uint8 uint16 uint32 uint64
float32 float64
bool
pointers
```

### Operators

```go
== != < > <= >= && || !
```


### Packages

```go
package main

import "fmt" // standard package | custom package

...
```

> In go, packages are like partial classes or an assembly  c#
> all files with same package name can access all functions and variables defined in that package.

> 

> index.go
```go
package main 
import "fmt" // Note! packages don't share imports globally. Other files using main package must import fmt package to access it

...
```

<!-- git_url -->

