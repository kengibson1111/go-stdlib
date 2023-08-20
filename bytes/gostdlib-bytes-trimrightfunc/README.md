# bytes - trim (right func)

TrimRightFunc() is UTF-8-based for the input slice. It will use the input function to determine the bytes to trim from the end of the input slice. It most likely does not allocate byte-array memory for the return slice since the return should be a subslice of the input slice.

The function input is a rune, so there must be an underlying cast of bytes to rune walking backward from the end of the input slice. Most of the time, there isn't going to be an issue between UTF-8 encoded code points (runes) and Unicode code points (runes).
