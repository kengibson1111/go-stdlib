# bufio - empty final token

Buffered I/O can be very effective. A Scanner uses an underlying slice buffer of MaxScanTokenSize by default. That is what the NewScanner() function is setting up. Therefore, by default NewScanner() is allocating an underlying byte array for a slice with a capacity of MaxScanTokenSize. The slice buffer itself can be customized using the Scanner.Buffer() function, but for most cases the default buffer should suffice.

This example shows how to plug in a custom 32-bit integer split function. The default split function is ScanLines(). The custom function does not use any other split function available in the Go Standard Library. It parses tokens based on a comma separator. The Scanner.Split() function replaces the current split function with a new split function.

The example will intentionally not fail to demonstrate a forgiving string parser when an empty string is encountered.
