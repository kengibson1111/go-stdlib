# Notes on types

* string -  a read-only slice of bytes. Look at my notes in tour-of-go concerning slices. Since
  it is read-only, there is no use for a capacity, but it is there technically. When you convert
  from a byte array to string or string to byte array, a copy is made. When you make a substring
  the substring can be another slice pointing to the same byte array.
