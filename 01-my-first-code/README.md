# Go-Lang
In this my first code I start with the standard Golang procedure (package main ; import "fmt"). After that, I start the code itself.

In this program, I'm attaching a func and a struct.

# Type definition
I start by defining 'Customer' as a structure that will receive 'Name' as a string, 'Email' also as a string, and finally 'Country' as a string.

# Creating the function Show
Right after that, I start creating the inner function that will do what I want.

I declare the func named Show (since it will serve to show the indicated data) with a connection to the Client struct, like this: 'func (c Client) Show'.

Right after that, I declare a Printf, which will allow me to create a gap inside the sentence. Then I write the sentence 'The client %v [gap]', at the end of the sentence, I use a comma and declare the element of the struct that will replace the gap, 'c.Name' and do something similar for the Email and the country.

After that, I end the Show func.

# Func main
After having created everything necessary, I proceed to create the main function.

I declare that 'client' will have the same properties as 'Client' using: 'client := Customer'. After having done that, I declare the desired values ​​for the variables that exist in 'Client', then I ask for the Show function to be used, taking the information declared in the Main function: 'client.Show()'.

Now just put it to work!