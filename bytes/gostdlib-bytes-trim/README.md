# bytes - trim

Trim() is UTF-8-based. It will treat the inputs as UTF-8 encoded code point trimming the cutset slice from the input slice. It does not allocate byte-array memory for the return slice since the return is a subslice of the input slice.
