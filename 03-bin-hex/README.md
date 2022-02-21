# Changing a number to binary or/and hexadecimal

You don't need another function beyond the main func. You don't need declare a variable. You only need, of course, the 'package main', 'import fmt' and the main func.

# Func main

Declare X with a ':=', as you didn't declare this variable before. Then write 'x := 10'. After that I wrote 'fmt.PrintF' to declare X twice. '%b' and '%#x' will change X value to binary and hexadecimal respectively.

# Why '#' before X in '%#x'?

You can change the X value to hex with out the '#' but it can be confusing. When you do this, the result is similar to decimal number, because it is a hex number with out the hex symbol, as '0x14' will just be '14'.