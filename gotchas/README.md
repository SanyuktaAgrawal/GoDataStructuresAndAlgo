# Gotcha while appending to slices in Golang
 When you pass a slice to a function and try to change the element of slice, it works. But when you try to add a new element to slice inside a function, added elements are not visible to caller function.

 To understand above issue, lets see internal representation of slice in detail.

Internally a slice is represented by three things.

- Pointer to the underlying array
- Length of underlying array
- Total Capacity which is the maximum capacity to which the underlying array can expand.

```
type SliceHeader struct {
        Pointer uintptr
        Len  int
        Cap  int
}
```

It’s important to understand that even if slice contains a pointer to underlying array, it is itself a value. So When we are passing slice to modifySlice, we are passing a struct value which is holding a pointer to underlying array, length and capacity of array. A point to note here is slice is not a pointer to a struct.

Even though the slice header is passed by value, the header includes a pointer to elements of an array, so when we modified 0th index of slice s[0] = 11 in our modifySlice function, the underlying array element got changed. Therefore, when the function returns, the modified index was visible to the caller function.

Why newly added element inside modifySlice and changes made after that are not visible to caller function?

Whenever any append, copy or assign operations are made on the slice, a new array is created and the pointer to underlying array is overwritten. Now, the elements present in the original array are copied into a new array and returns a new slice pointing to the new array.
So when the our slice got reallocated, a new location of the memory is used. Even if the values in slice are the same, the slice points to a new memory location and therefore changes post slice re-allocation are not visible to the caller function.

Solution:
Return the value after appending to the slice and then assign it to the original slice.
package main

Go Playground [link](https://go.dev/play/p/VR6PvZf7CD_5).

One thing to be considered is even if you are returning the updated slice and assigning to the same value, its original length and capacity will change, which will lead to a new underlying array of different length. Check the length and capacity before and after changing the slice to see the difference.

`fmt.Println(len(slice), cap(slice))`
