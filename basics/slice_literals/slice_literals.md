# Slice literals

A slice literal is like an array literal without the length.

This is an array literal:
```text
[3]bool{true, true, false}
```
And this creates the same array as above, then builds a slice that references it:

```text
[]bool{true, true, false}
```
