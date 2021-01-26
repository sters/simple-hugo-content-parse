package example

import (
	"bytes"
	"fmt"

	hugocontent "github.com/sters/simple-hugo-content-parse"
)

func ExampleParseMarkdownWithYaml() {
	raw := `---
foo: bar
baz: 1
---

foo
`

	c, err := hugocontent.ParseMarkdownWithYaml(bytes.NewBufferString(raw))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s, %d\n", c.FrontMatter["foo"], c.FrontMatter["baz"])
	fmt.Printf("%s", c.Body)
	// Output:
	// bar, 1
	//
	// foo
}
