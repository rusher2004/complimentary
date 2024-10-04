# Complimentary

Complimentary is a high performance (O(n)) utility aimed towards increasing developer productivity by providing thoughtful and encouraging feedback on startup. It is especially effective in projects with live reload.

## Installation and Usage

Add to your Go project

```sh
go get github.com/rusher2004/complimentary
```

Import wherever it is needed.

```go
package main

import (
	"fmt"

	_ "github.com/rusher2004/complimentary"
)

func main() {
	fmt.Println("Hello, World!")
}
```
