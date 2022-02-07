package main
import "fmt"

type Client struct{
	Name string
	Email string
	Country string
}

func (c Client) Show(){
	fmt.Printf("The client %v ", c.Name)
	fmt.Printf("uses %v email ", c.Email)
	fmt.Printf("and lives in %v\n", c.Country)
}

func main(){
	client := Client{
		Name: "Peter",
		Email: "p@p.com",
		Country: "Canada",
	}

	client2 := Client{
		Name: "John",
		Email: "j@j.com",
		Country: "Brazil",
	}

	client.Show()
	client2.Show()
}
