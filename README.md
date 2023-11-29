# Booklet Printing

Helps to computes in which order to print in a booklet format.

1. Open the printing window
2. Paste in 'Pages' the result of the command
3. Select :
   * A4 format
   * 2 pages per sheet
   * Double-sided
   * Short edge (turned over)
4. Print the doc

## Usage

```sh
booklet-printing -p 12
```

Result :
```sh
2023/11/29 16:42:39 You asked for a 12 page document.
2023/11/29 16:42:39 Printing a booklet ordered page : 12,1,2,11,10,3,4,9,8,5,6,7
```
