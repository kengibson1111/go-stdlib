# bytes - runes

Runes() treats the input slice as a set of UTF-8-encoded code points. Each UTF-8-encoded code point can be 1 to 4 bytes. The example shows ASCII characters which all translate to one byte each.

Runes() returns a slice of Unicode code points. Even though, technically, Unicode code points only require a maximum of 3 bytes, Unicode code points mostly require 2 bytes each. However, there are Unicode code points that fall into the 4-byte size.

Understanding this difference is important if you know that your applications potentially use Unicode characters that require 4 bytes each. Fortunately, Go provides a level of protection from different encoding strategies (UTF-8, UTF-16, UTF-32) and the Unicode character set. If your applications uses functions that are founded on runes, they should be safe from potential errors related to encoding differences.
