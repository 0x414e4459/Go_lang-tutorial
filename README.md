# Go lang-Tutuorial

Go lang is a compiled statically typed language. Boiler template includes a package main, iport fmt;func main.  
eg:  
package main;  
import "fmt";  
func main(){  
fmt.Println(("hello world!"));  
}

## Variables and Data types

variables are denoted by **_"var"_** keyword followed by **name of variable** followed by **type** of variable.  
**Implicit** declaration: using "var" keyword before var_name, type is optional. We can declare or the compiler decides on what type we assigned it.  
eg:var x=5; var y uint8=3;  
**Explicit** declaration: No "var" keyword is used. Instead " **:=** " is used and type is optional.  
eg: x:=5;

#### 1. Boolean

They are boolean types and consists of the two predefined constants: (a) true (b) false

#### 2. String/char

A string type represents the set of string values. Its value is a sequence of bytes. Strings are immutable types that is once created, it is not possible to change the contents of a string. The predeclared string type is string.

#### 3,Numeric

They are again arithmetic types and they represents a) integer types or b) floating point values throughout the program.

##### i. Integer

The predefined architecture-independent integer types are:

1. uint8: Unsigned 8-bit integers (0 to 255)
2. uint16: Unsigned 16-bit integers (0 to 65535)
3. uint32: Unsigned 32-bit integers (0 to 4294967295)
4. uint64: Unsigned 64-bit integers (0 to 18446744073709551615)
5. int8: Signed 8-bit integers (-128 to 127)
6. int16: Signed 16-bit integers (-32768 to 32767)
7. int32: Signed 32-bit integers (-2147483648 to 2147483647)
8. int64: Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

##### ii.Float

The predefined architecture-independent float types are:

1. float32: IEEE-754 32-bit floating-point numbers
2. float64: IEEE-754 64-bit floating-point numbers
3. complex64: Complex numbers with float32 real and imaginary parts
4. complex128: Complex numbers with float64 real and imaginary parts

### fmt module

**printf()** -- logs the formatted output to console. takes two args:formatter, variable. Printf formats according to a format specifier and writes to standard output. It returns the number of bytes written and any write error encountered.  
**sprintf()** -- same as printf but, used to assign formatted output to variables.

#### formaters:

- _General_  
  %v --value of the variable  
  %T -- type of variable  
  %% --literal %

- _Boolean_  
  %t -- true or false

- _Integer_  
  %b -- base2  
  %o -- base8  
  %d -- bsae10  
  %x -- base16

* _Flaot_  
  %e --scentific notattion  
  %f/%F --decimal no exponent  
  %g -- for larger exponents

* _String_  
  %s --default  
  %q -- double quoted string

* _Width & precsion_  
  %f --default precision  
  %(n).(m)f -- width of size "n" and precision upto "m" digits  
  %4.2f -- width of 4, precision upto 2 decimals

#### Console Input

required packages/modules: 'bufio'  
Syntax:  
scanner=bufio.NewScanner();  
scanner.scan() //reads the line from console  
by default input treated as **string**

### Arthametic Operators

- Basic operators +,-,/,\*,%
- operation's are allowed between same data types.
- Conversion can be done by using data_type(variable)

#### Bianry operators

and -- &&  
or -- ||  
not --!

### Conditional Statements

**_if-else_**:
syntax:  
if(condotion){  
//logic} else if(  
condition){  
//logic}else{  
//logic}  
else,else if keywords shoild start at the end of parenthesis.

### Iteration statements
