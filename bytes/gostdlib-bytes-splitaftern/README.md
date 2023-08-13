# bytes - split (after n)

This example shows SplitAfterN() behavior with a variety of input slices. Since each input slice is a string literal, it may be possible for each output slice in the slice array to point to the same underlying byte array as the input slice. The count (n) determines the number of maximum slices in the output slice array. A count (n) of -1 will result in the same behavior as SplitAfter().

This example is very similar to the SplitN() example. Notice the differences and similarities in the output of each.
