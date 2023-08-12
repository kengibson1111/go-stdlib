# bytes - split (n)

This example shows SplitN() behavior with a variety of input slices. Since each input slice is a string literal, it may be possible for each output slice in the slice array to point to the same underlying byte array as the input slice. The count (n) determines the number of maximum slices in the output slice array. A count (n) of -1 will result in the same behavior as Split().
