package main
import "fmt"
func main() {
	map01 := make(map[string] int);
	map01["Shining"] = 35;
	map01["Kite"] = 33;

	for k, v := range map01 {
		fmt.Println(k, v);
	}

	age, ok := map01["Torres"];
	if ok {
		fmt.Print("Torres's age is: \n", age);
	} else {
		fmt.Print("Not found\n");
	}

	map01["Torres"] = 2;
	age = map01["Torres"];
	fmt.Printf("Torres's age is: %d\n", age);
}
