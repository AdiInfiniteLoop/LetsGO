package main

import ("fmt", "bufio", "strconv", "strings")

func main() {
	/* Lexer in G0: UNderstand the grammar and make sure correct grammar is used */
	/* Types in GO:
		function starting with capital is public	(write as this)
		:= walrus operator
	*/
	welcome :=	'Welcome to the user input'
	//Input using buffer
	reader := bufio.NewReader(os.Stdin) //using bufio take input using New Reader
	
	
	fmt.Println("Enter the rating for our Pizza")
	//comma-ok syntax as there is no try catch in GO
	input, _ := reader.ReadString('\n') //(an ender: i.e., keep reading till we reach \n) // _ is used when we don't care about it(like here we do not care for err)
	fmt.Println("Thanks for reading ", input)

	//Conversions in Go
	fmt.Println('Please rate our pizza between 1 to 5')
	data , _ := reader.ReadString('\n')

	numrating, err = strconv.ParseFloat(strings.TrimSpace(data), 64) //cannot add 1 to string //64 means bit size
	if err != nil {
		fmt.Println(err)
		// panic(err) //to end the program
	}
	else {
	fmt.Println("The rating is ", numrating % 5)
	}


}