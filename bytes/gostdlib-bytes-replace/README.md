# bytes - replace

Replace() finds byte-slice matches in the input slice and replaces the byte-slice matches with the new byte slice.

If the count is greater than zero, count determines the number of byte-slice matches from the beginning of the input slice to replace. If the count is less than zero, all byte-slice matches are replaced. If count is greater than zero and the number of byte-slice matches is less than or equal to the count, all byte-slice matches are replaced.

The returned slice is newly allocated including the underlying byte array for the slice.
