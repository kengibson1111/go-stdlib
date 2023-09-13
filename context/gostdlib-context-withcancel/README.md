# context - withcancel

This example shows how to prevent go subroutine leaks using WithCancel().

Code comments explain the mechanics. There are many use cases where go subroutines are created to do work - each should check if the execution context is done so the subroutine can complete as expected.
