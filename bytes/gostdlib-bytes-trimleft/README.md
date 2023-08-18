# bytes - trim (left)

TrimLeft() is UTF-8-based. It will treat the inputs as UTF-8 encoded code points trimming the cutset slice from the beginning of the input slice. It does not allocate byte-array memory for the return slice since the return is a subslice of the input slice.
