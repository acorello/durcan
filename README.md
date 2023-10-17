# durcan

Convert durations to canonical format.

Optionally:

- truncate the output to give duration amount
- divide each value by given integer

```text
$ durcan -truncate-to 5s 3600s 6s 1s
3600s       =    1h0m0s
6s          =        5s
1s          =        0s
```

```text
$ durcan -truncate-to 5s -div-by 2 3600s 6s 1s
3600s       =    1h0m0s
  div-by  2 =     30m0s
6s          =        5s
  div-by  2 =        0s
1s          =        0s
  div-by  2 =        0s
```
