package main
import "fmt time math/rand"

func main() {
	rand.Seed(time.Now().UnixNano());//to get random number
	diceNumber := rand.Intn(6) + 1
	fmt.Printf("The dice number is %v \n", diceNumber)

	switch diceNumber {
	case 1 :
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
		//fallthrough -> to go for next case too
	case 5:
		fmt.Println("5")
	case 6:
		fmt.Println("6")
	default:
		fmt.Println("What was that again?\n")
	}
}