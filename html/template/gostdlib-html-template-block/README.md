# html/template - block

This is showing how to customize a template block. block is shorthand for defining a template and executing it in place. That is what allows the template block to be redefined.

In this sample, overlay is redefining the master template block. Notice how the master template instance assign the join function pointer and not the overlay. But the overlay template redefinition of the master template block uses join. Starting to get interesting when considering dynamic execution functionality. This is the same as the text/template block sample.
