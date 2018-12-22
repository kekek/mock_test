package cmd

import (
	"fmt"

	"github.com/sgoertzen/html2text"
)

var originContent = `
	<html >
		<head> test </head>
		<body>
			啦啦啦啦啦啦啦啦啦啦
		</body>
	</html>
`

func main() {
	x, err := html2text.Textify(originContent)
	if err != nil {
		fmt.Println("err")
	}

	fmt.Println(x)

	testDef()

	fmt.Println("test")
}

func testDef() {
	fmt.Println("def")
}
