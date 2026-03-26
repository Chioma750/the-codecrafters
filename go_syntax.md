## Go syntax
> A go file consists of the following parts which are listed below:

    * package declaration
    * import packages
    * functions
    * statements

**for example**
``` go
package main
import ("fmt")

func main () {
    fmt.Println ("Hello, World!")
}
```
> Explaining the code example line by line,

**Line 1:** In go every program has a package, that is why we defined this using the package keyword. And in the above example the program belongs to the main package. 

**Line 2:** import ("fmt") Here we are importing files included in the fmt package.

**Line 3:** A blank line, Yes, because having white spaces in code makes it more readable.

**Line 4:** func main () {} is a function, and any code inside it's curly brackets will be executed.

**Line 5:** fmt.Println () is a function made available from the "fmt" package, it is used to print(output) text. In the above example it will print "Hello, World!".  

