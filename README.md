# Go package for Converting Decimal to an Indonesian Wording format

***This package is still in early development and might contain a bug***

This package will help you to convert a decimal to an Indonesian wording format, for example `123` will be converted to `seratus dua puluh tiga`. This package was inspired from [develoka/angka-terbilang-js](https://github.com/develoka/angka-terbilang-js) by [Develoka Team](https://github.com/develoka).

## Getting Started

These instructions will get you this package can be used in your project.

### Prerequisites

Prequisites Packages:
* Go (Go Programming Language)

### Using This Package

Below is the instructions to use this package:
* Import this package to your code
```go
...
import (
  ...
  angka "github.com/dimaskiddo/angka-terbilang-go"
  ...
)
...
```
* Use ToTerbilang function to convert your decimal string
```go
...
  fmtTerbilang := angka.ToTerbilang("123");
  fmt.Println(fmtTerbilang);
...
```
* Download and make the package as vendor
```sh
GO111MODULE=on go mod vendor
```

### Example Usage of Package

Below is the simple example source code:
```go
package main

import (
  angka "github.com/dimaskiddo/angka-terbilang-go"
)

func main() {
  fmt.Printf("%v\n", angka.ToTerbilang("123"));
}
```

## Running The Tests

Currently the test is not ready yet :)

## Built With

* [Go](https://golang.org/) - Go Programming Languange

## Authors

* **Dimas Restu Hidayanto** - *Initial Work* - [DimasKiddo](https://github.com/dimaskiddo)

See also the list of [contributors](https://github.com/dimaskiddo/angka-terbilang-go/contributors) who participated in this project
