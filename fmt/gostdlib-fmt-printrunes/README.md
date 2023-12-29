# fmt - printrunes

Notice the term "char" is not used that much in golang. That is because the Unicode standard uses the term "code point". Those include Latin characters we are familiar with and much more. So a Unicode code point is a golang "rune". And most of the time it doesn't matter because we usually only are interested in Latin characters that can be contained within one byte. We use a for loop one byte at a time and everything is fine. We use a for range loop and still everything is fine with Latin characters.

But with runes, a for range loop will start each iteration at the beginning of the rune while a for loop will iterate one byte at a time. Something to think about when processing strings. The example code uses a for range loop. Notice the byte position.
