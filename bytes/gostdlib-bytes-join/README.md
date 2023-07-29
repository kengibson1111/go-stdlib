# bytes - join

This concatenates all of the slices in the input slice of slices adding the seperator slice between each slice in the input slice of slices. The result is a new slice and *heap memory being allocated for an underlying byte array*.

This is different than several of the functions in **bytes** that return a slice pointing to an existing underlying byte array (like Next()). Whenever a slice is being returned from a function, the assumption is that the slice struct itself is newly allocated on the heap. Most of the time when this happens, the returned slice can be pointing to an existing underlying byte array. Join() does not reuse an existing underlying byte array for its returned slice.
