# Notes on primitive types

## Architecture-independent

These are portable across hardware architectures.

* bool - true or false. golang does not specify bit-length.

* uint8 - the set of all unsigned 8-bit integers (0 to 255)
* uint16 - the set of all unsigned 16-bit integers (0 to 65535)
* uint32 - the set of all unsigned 32-bit integers (0 to 4294967295)
* uint64 - the set of all unsigned 64-bit integers (0 to 18446744073709551615)

* int8 - the set of all signed 8-bit integers (-128 to 127)
* int16 - the set of all signed 16-bit integers (-32768 to 32767)
* int32 - the set of all signed 32-bit integers (-2147483648 to 2147483647)
* int64 - the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

* float32 - the set of all IEEE-754 32-bit floating-point numbers
* float64 - the set of all IEEE-754 64-bit floating-point numbers

* complex64 - the set of all complex numbers with float32 real and imaginary parts
* complex128 - the set of all complex numbers with float64 real and imaginary parts

* byte - alias for uint8
* rune - alias for int32

* string - a read-only slice of bytes. Look at my notes in tour-of-go concerning slices. Since
  it is read-only, there is no use for a capacity, but it is there technically. When you convert
  from a byte array to string or string to byte array, a copy is made. When you make a substring
  the substring can be another slice pointing to the same byte array.

## Architecture-specific

Better to stay away from to avoid portability issues.

* uint - either 32 or 64 bits

* int - same size as uint

* uintptr - an unsigned integer large enough to store the uninterpreted bits of a pointer value
