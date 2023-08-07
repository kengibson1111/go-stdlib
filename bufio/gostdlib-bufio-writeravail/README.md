# bufio - writer (available buffer)

Buffered I/O can be very effective. A Writer uses an underlying slice buffer of default size. That is what the NewWriter() function is setting up. The example prints the default buffer size. Therefore, by default NewWriter() is allocating an underlying byte array for a slice with a capacity of Go's default buffer size.

So, why even use buffered I/O? The main reasons are to **help reduce runtime executable I/O** and **speed up executable processing** by reducing the number of I/O-related function calls. I/O is one of the metrics to watch carefully as part of a comprehensive Site Reliability Engineering strategy.

This example shows how to wrap a bufio.Writer around an existing io.Writer implementation. Then the bufio.Writer can be used in functions that normally expect an io.Writer interface since the bufio.Writer implements the io.Writer function signatures as well. Unique to this example is the *direct use of the available buffer associated with the bufio.Writer instance*. The available buffer is only valid until the bufio.Writer.Write() operation. So, the intent is: assign the available buffer slice to a variable, append to the buffer, write the buffer, don't do anything else with the buffer because it is considered invalid.

The example prints buffer contents, length, and capacity in the loop to clarify bufio.Writer available-buffer behavior. Notice how the buffer slice looks the same after the append functions and after bufio.Writer.Write(). But, as soon as the buffer is reassigned at the top of the loop before the append functions, the length is zero and the capacity is reduced.

Emphasizing available-buffer validity, *don't do any operations on an available buffer after bufio.Writer.Write() without first reassigning a variable to the bufio.Writer's available buffer*. Notice in the example, an additional append is attempted after bufio.Writer.Write(). Printing the buffer shows it. But, it is lost at the top of the loop with the buffer slice variable is reassigned. The loss is shown when flushing the writer. There aren't two spaces between each number, just one.
