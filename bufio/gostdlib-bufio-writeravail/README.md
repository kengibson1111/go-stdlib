# bufio - writer (available buffer)

This shows how to use the available buffer on a writer in order to write to a stream. It is meant to be appended to before an immediate write. The buffer is only valid until the write operation. After the write operation, the buffer is not valid and can be reused.
