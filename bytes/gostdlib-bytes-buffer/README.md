# bytes - buffer

This shows how to use a Buffer in various ways (direct write, as a Writer, transferring to a File).

An empty Buffer is ready to use without any memory preallocation. A Buffer is a Writer, so the Write() function can be used directly or indirectly by a function that uses a Writer - like fmt.Fprintf(). Buffer contents can be transferred to a File - like os.Stdout.
