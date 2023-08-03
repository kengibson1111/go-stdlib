# bytes - reader (len)

This returns the length of the unread portion of a Reader's byte slice.

This example highlights the difference of ASCII UTF-8 characters that required 1 byte for each character and other UTF-8 characters that require more than 1 byte for each character.

Remember that a Go string is a read-only slice of bytes. It does not have to be Unicode. It does not have to be UTF-8. How the slice of bytes is interpreted for output purposes depends on the format directive in print statements - which leads to the next point.

**Go source code is UTF-8**. So, if Go source code contains string literals, the string literals are UTF-8 text. The string literal could be a raw string (a string using back quotes instead of double quotes) which means that the string cannot contain escape sequences. If the string literal uses double quotes, escape sequences can be part of the string literal. In both cases, the string literal value is stored as a read-only slice of bytes.

This is actually very powerful compared to string implementations in other languages/runtimes. For example, REST APIs that process JSON inputs and outputs rely on efficient string parsing and validation. Being able to allocate one string for a REST API input and parsing that string without allocating additional memory for substrings - that is a game changer. Each substring is a slice which translates to a very small amount of memory allocated and the slice pointing to the same underlying byte array.
