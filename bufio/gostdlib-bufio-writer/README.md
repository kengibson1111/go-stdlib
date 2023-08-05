# bufio - writer

Buffered I/O can be very effective. A Writer uses an underlying slice buffer of default size. That is what the NewWriter() function is setting up. The example prints the default buffer size. Therefore, by default NewWriter() is allocating an underlying byte array for a slice with a capacity of Go's default buffer size.

So, why even use buffered I/O? The main reasons are to **help reduce runtime executable I/O** and **speed up executable processing** by reducing the number of I/O-related function calls. I/O is one of the metrics to watch carefully as part of a comprehensive Site Reliability Engineering strategy.

This example shows how to wrap a bufio.Writer around an existing io.Writer implementation. Then the bufio.Writer can be used in functions that normally expect an io.Writer interface since the bufio.Writer implements the io.Writer function signatures as well.
