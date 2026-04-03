## Go Slices
Go slices are flexible and powerful data structures used to manage collections of the same data type. While they resemble arrays, slices are dynamic and can change size as needed. 
## Creating Slices
You can initialize a slice in Go using several methods:

**Slice Literals**: Define a slice with values directly using the []type{values} syntax. For example, myslice := []int{1, 2, 3} creates a slice with a length and capacity of 3.

**Empty Slices**: Use myslice := []int{} to create an empty slice with zero length and capacity.

**Using make()**: The built-in make() function allows you to specify a type, length, and optional capacity, such as make([]int, 5, 10).

**From Arrays**: You can create a slice from an existing array by specifying a range: array_name[start:end]. 

### Measuring Slices
Two key functions help manage a slice's size:

    len(): Returns the current number of elements stored in the slice.
    cap(): Returns the slice's capacity, which is the total number of elements it can hold before the underlying array must be resized. 

## Access Elements of a Slice

You can access a specific slice element by referring to the index number.

In Go, indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

## Memory Efficiency
    In Go, slicing a large array keeps the entire underlying array in memory, which can lead to inefficient memory usage if only a small portion is needed. To optimize, use the copy() function to create a new, smaller slice containing only the necessary elements. This method ensures the original large array can be garbage collected, significantly reducing the program's memory footprint. 
