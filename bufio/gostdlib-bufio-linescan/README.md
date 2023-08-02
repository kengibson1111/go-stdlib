# bufio - line scan

Buffered I/O can be very effective. A Scanner uses an underlying slice buffer of MaxScanTokenSize by default. That is what the NewScanner() function is setting up. Therefore, by default NewScanner() is allocating an underlying byte array for a slice with a capacity of MaxScanTokenSize. The slice buffer itself can be customized using the Scanner.Buffer() function, but for most cases the default buffer should suffice.

In this example, the default Scanner split function will be used - ScanLines(). Each carriage return (\n) from standard input will viewed as a token delimiter for the Scanner.Scan() function. A token length of zero will exit the example (carriage return from standard input with no other characters typed).

Remember that a Go string is a read-only slice of bytes. It does not have to be Unicode. It does not have to be UTF-8. How the slice of bytes is interpreted for output depends on the format directive in print statements.
