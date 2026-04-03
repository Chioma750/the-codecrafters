## Go Arrays
 

1. With the var keyword

var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred

2. With the := sign:

array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred 

**NB**  In Go, arrays possess a constant size that determines how many elements they can hold. This length can be established either by explicitly providing a numeric value or by allowing the compiler to automatically calculate it based on the number of elements provided.

### Access Elements of an Array

You can access a specific array element by referring to the index number.

In Go, array indexes start at 0. That means that [0] is the first element, [1] is the second element, [2] is the third element, and so on.

 If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.

**Tip**: The default value for int is 0, and the default value for string is "".
