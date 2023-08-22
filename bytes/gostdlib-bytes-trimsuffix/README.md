# bytes - trim (suffix)

TrimSuffix() is UTF-8-based. It will treat the inputs as UTF-8 encoded code points trimming the suffix slice from the end of the input slice. It does not allocate byte-array memory for the return slice since the return is a subslice of the input slice.
