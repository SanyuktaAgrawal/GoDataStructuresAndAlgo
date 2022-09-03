# Description
 write a program to print the sum of the squares and cubes of the individual digits of a number using channels.

For example, if 123 is the input, then this program will calculate the output as

```
squares = (1 * 1) + (2 * 2) + (3 * 3)
cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3)
output = squares + cubes = 50
```

 Structure the program such that the squares are calculated in a separate Goroutine, cubes in another Goroutine and the final summation happens in the main Goroutine.
