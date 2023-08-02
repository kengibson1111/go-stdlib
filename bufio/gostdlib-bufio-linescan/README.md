# bufio - line scan

Buffered I/O can be very effective. A Scanner uses an underlying slice buffer of MaxScanTokenSize by default. That is what the NewScanner() function is setting up. Therefore, by default NewScanner() is allocating an underlying byte array for a slice with a capacity of MaxScanTokenSize. The slice buffer itself can be customized using the Scanner.Buffer() function, but for most cases the default buffer should suffice.

In this example, the default Scanner split function will be used - ScanLines(). Each carriage return (\n) from standard input will viewed as a token delimiter for the Scanner.Scan() function. A token length of zero will exit the example (carriage return from standard input with no other characters typed).

Remember that a Go string is a read-only slice of bytes. It does not have to be Unicode. It does not have to be UTF-8. How the slice of bytes is interpreted for output purposes depends on the format directive in print statements - which leads to the next point.

Go source code is UTF-8. So, if Go source code contains string literals, the string literals are UTF-8 text. The string literal could be a raw string (a string using back quotes instead of double quotes) which means that the string cannot contain escape sequences. If the string literal uses double quotes, escape sequences can be part of the string literal. In both cases, the string literal value is stored as a read-only slice of bytes.

This is actually very powerful compared to string implementations in other languages/runtimes. For example, REST APIs that process JSON inputs and outputs rely on efficient string parsing and validation. Being able to allocate one string for a REST API input and parsing that string without allocating additional memory for substrings - that is a game changer. Each substring is a slice which translates to a very small amount of memory allocated and the slice pointing to the same underlying byte array.
