# bytes - repeat

This returns a new slice repeating the input slice contents based on the input count. Since it is a new slice, the assumption is that it is also a newly allocated underlying byte array for the new slice.

Functionally, I don't know how much value this brings in real-world use repeating a sequence of input bytes a certain number of times. I don't know what the business cases would be.
