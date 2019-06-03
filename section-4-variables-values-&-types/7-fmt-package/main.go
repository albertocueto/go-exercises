package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("%#x\n", y)
	// format printing with multiple values passed
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)
	// Scape characters:
	// \a   U+0007 alert or bell
	// \b   U+0008 backspace
	// \f   U+000C form feed
	// \n   U+000A line feed or newline
	// \r   U+000D carriage return
	// \t   U+0009 horizontal tab
	// \v   U+000b vertical tab
	// \\   U+005c backslash
	// \'   U+0027 single quote  (valid escape only within rune literals)
	// \"   U+0022 double quote  (valid escape only within string literals)

	// string printing with fmt:
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)

	// There's also a File printer with "Fprint"
	fmt.Printf("%v\n", y)
}
