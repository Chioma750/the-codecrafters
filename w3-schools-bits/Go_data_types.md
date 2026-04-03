# Data Types
    In Go, data types are foundational because they determine a variable's memory size and the kind of information it can hold. As a statically typed language, Go requires a variable's type to be fixed at the time of its definition, meaning it can only store values of that specific type.

Go has three basic data types:

    bool: represents a boolean value and is either true or false
    Numeric: represents integer types, floating point values, and complex types
    string: represents a string value
## Boolean Data Type   
    A boolean data type is declared with the bool keyword that can only take the values true or false.
    The default value of a boolean data type is false.
**NB** Boolean values are primarily used to drive conditional logic, which is covered in detail in the Go Conditions chapter.

## Integer Data Types
Integer data types are used to store a whole number without decimals, like 76, -54, or 234000.
The integer data type has two categories:

    Signed integers - can store both positive and negative values
    Unsigned integers - can only store non-negative values
**NB** The default type for integer is int. If you do not specify a type, the type will be int.

Go has five keywords/types of signed integers:

    1. int

    2. int8

    3. int16

    4. int32

    5. int64

### Unsigned Integers
    Unsigned integers, declared with one of the uint keywords, can only store non-negative values:

Go has five keywords/types of unsigned integers:

    1. unit

    2. unit8

    3. unit16

    4. unit32

    5. unit64

## Go Float Data Types
    The float data types are used to store positive and negative numbers with a decimal point, like 22.2, -34.2, or 34987.3597

The float data type has two keywords:

    1. float32

    2. float64

## String Data Type  
    The string data type is used to store sequence of characters. Remember string values must be surrounded by double quotation.



