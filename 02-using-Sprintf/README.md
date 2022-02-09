# My first time using Sprintf
To sart this readme I need explain what is a 'Sprintf' and why it is different of a normal Print or a Printf. A Sprintf allows that you give a many valors for 1 variable, while a Print or Printf just show the valors of the variables.

# Declaring variables
To start the program, we need declare the variables that we will use (x, y, z). In my code, x is a int, y is a string and, z is a bool (for a short summary of what we will be using here, bool is a variable with 'false' how your inicial valor, but you can change this for 'true'). After declare the type of the variables, we need give they a valor like: 'var x int = 42'.

# Func main
Until this moment we did not declare the variable that will go recibe the valors of x, y and z, but we will do this in func main using the Sprintf. In func main, I wrote 's :=' for 's' recibe the valor and the definition of the others variables, 's := fmt.Sprintf("%v\t%v\t%v", x, y, z)', 'Sprintf' allows take the valor of others variables, '%v' allows me control where the variable's valor will are in a phrase/sentence, while '\t' do a space among this valors, facilitating the understanding. To end the code, I used a normal Print to show the valor was assigned to 's'.

# Now just put it to work!