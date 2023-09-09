# compress/flate - reset

This example shows how to reset current compressor or decompressor state. The reset takes advantage of previously allocated memory.

This is a nice optimization that won't be required for most applications. Based on Google's requirements for massive scale, it does not surprise me that this level of capability detail is included in low-level packages.
