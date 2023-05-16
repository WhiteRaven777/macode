# macode

`macode` is a Go package that provides functionality to encode (flatten) and decode (reconstruct) maps into ordered pairs of slices. It uses Go 1.18 generics, making it versatile and compatible with a wide range of types.

The `macode` package is particularly useful when you need to convert maps into a flat structure, for example, for serialization or for compatibility with certain data formats. Moreover, it ensures idempotent conversions, which is crucial for reproducibility of results in many applications.

## Installation
To install the macode package, you can use go get:

```bash
go get github.com/WhiteRaven777/macode
```

## Usage
Here is a basic example of how to use the `macode` package:

```go
package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/WhiteRaven777/macode"
)

func main() {
	m := map[int]string{
		3: "three",
		1: "one",
		2: "two",
	}

	p := macode.New(m)
	enc := p.Encode(m)
	fmt.Println(enc) // Output: [{1 one} {2 two} {3 three}]

	dec := p.Decode(enc)
	fmt.Println(dec) // Output: map[1:one 2:two 3:three]

	// Check that the original map and the decoded map are deeply equal
	if !reflect.DeepEqual(m, dec) {
		fmt.Println("The original map and the decoded map are not equal")
		os.Exit(1)
	} else {
		fmt.Println("The original map and the decoded map are equal")
	}
}
```

In this example, a map is encoded into an ordered slice of pairs using the `Encode` method, and then it is decoded back into a map using the `Decode` method. The original map and the decoded map are then compared using `reflect.DeepEqual` to ensure they are deeply equal.

## Testing
You can run the tests for the macode package using go test:

```bash
go test github.com/WhiteRaven777/macode
```

## License
This project is licensed under the terms of the MIT license.