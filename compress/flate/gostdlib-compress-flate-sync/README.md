# compress/flate - syncronization

This example shows how to use compressed data between go subroutines.

`defer wg.Wait()` at the top of the main function isn't executed until the end of the main function. That's how `defer` works. `wg.Add()` is what makes the `wg.Wait()` stop inline. Each subroutine also has a deferred `wg.Done()`. `wg.Done()` is what reduces the WaitGroup count.

When the WaitGroup count is zero, `wg.Wait()` doesn't have to wait anymore. The two subroutines are coordinated by compressed data passed from the PipeWriter to the PipeReader. The PipeReader and PipeWriter are created by `io.Pipe()`. One subroutine uses the PipeWriter, the other uses the PipeReader.
