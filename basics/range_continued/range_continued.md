# Range continued

You can skip the index or value by assigning to `_`.

```text
for i, _ := range pow
for _, value := range pow
```

If you only want the index, you can omit the second variable.

`for i := range pow`