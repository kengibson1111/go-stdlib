# flag - textvar

This shows how to use the TextVar function unmarshaling a string into another type (`net.IP`). This specific example works because `net.IP` implements `encoding.TextUnmarshaler`. This means that any custom type can also implement `encoding.TextUnmarshaler` and use the TextVar function.

Another thing to watch: the type of the default value in TextVar has to match the type of the underlying variable which the first parameter points to. Both parameters are implementations of `encoding.TextUnmarshaler`.
