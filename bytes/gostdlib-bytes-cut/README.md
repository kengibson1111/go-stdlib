# bytes - cut

This shows how to use the Cut() function. Cut() tries to find the first instance of the separator slice in the input slice. The resulting slices are as if you tried to "cut" the separator out of the input slice and you're left with the pieces of the input slice.

All of this is done without creating new slices and copying data. That would be very expensive over time. All of the return slices are reslices of the original input slice. In other words, all reslices (and the original slice) have the same pointer to the underlying byte array.

I added length and capacity to the output of each reslice to emphasize. This matches behavior in [this Tour of Go example](https://github.com/kengibson1111/tour-of-go/tree/master/moretypes/tourofgo-moretypes-sliceslencap).
