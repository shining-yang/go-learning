package main
import "fmt"

func MyVarFunc(args ... interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, " is an integer variable");
		case string:
			fmt.Println(arg, " is a string var");
		default:
			fmt.Println("Unidentified: ", arg);
		}
	}
}

func main() {
	MyVarFunc(1, 3.14, "Hello", [3] int { 1, 2, 3});
}
	
