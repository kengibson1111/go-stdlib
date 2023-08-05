# bufio - word scan

Buffered I/O can be very effective. A Scanner uses an underlying slice buffer of MaxScanTokenSize by default. That is what the NewScanner() function is setting up. Therefore, by default NewScanner() is allocating an underlying byte array for a slice with a capacity of MaxScanTokenSize. The slice buffer itself can be customized using the Scanner.Buffer() function, but for most cases the default buffer should suffice.

So, why even use buffered I/O? The main reasons are to **help reduce runtime executable I/O** and **speed up executable processing** by reducing the number of I/O-related function calls. I/O is one of the metrics to watch carefully as part of a comprehensive Site Reliability Engineering strategy.

This example shows how to plug in another split function from the Go Standard Library. The default split function is ScanLines(). The Scanner.Split() function replaces the current split function with a new split function - in this case, ScanWords().

ScanWords() treats whitespace characters as a separator. The example input string to be scanned includes tab, space, and newline characters.
