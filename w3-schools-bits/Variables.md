## Variables

Variables are containers for storing data values.
We have different types of variables in Go, here are some examples :

1. **Int**: Which stores integers i.e numbers only e.g 123, +123 or -123
2. **String**: Which stores only words that are surrounded with double quotation e.g "Chioma"
3. **Float32 and Float64**: These stores only numbers with decimals e.g 99.9
4. **Bool** These stores the true and false value.

## Declaring Variables
We have two ways of declaring variable in Go which are :
1. With the var keyword
2. With the := sign

**Explaining with examples**
1. With var keyword

Use the var keyword, followed by the variable name and type.

E.g  var variablename type = value

i.e var name string = Chioma

Note That: You always have to specify either your type or your value

2. With the := keyword

Use the := sign , followed by the variable value.

E.g variablename := value

i.e name := "Chioma"

In this case the compiler decides the type of the variable based on the value.

## Difference Between var and :=
1. var can be used inside and outside of functions **While** := can only be used inside functions.  

2. In var, variable declaration and value assignment can be done seperatly **While** in := variable declaration and value assignment must be done in the same line.

