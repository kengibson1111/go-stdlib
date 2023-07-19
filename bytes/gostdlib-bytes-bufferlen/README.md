# bytes - buffer (len)

This shows how to get a buffer's underlying byte-slice length. This isn't necessarily the capacity. The buffer length represents the number of unread bytes in the buffer. Notice how the buffer length is zero after os.Stdout reads the buffer.
