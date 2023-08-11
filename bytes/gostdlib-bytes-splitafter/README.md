# bytes - split (after)

This example shows SplitAfter() behavior which includes the separator in each output slice except for the last output slice in the returned slice array. Since each input slice is a string literal, it may be possible for each output slice in the slice array to point to the same underlying byte array as the input slice.
