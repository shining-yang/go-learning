package main
import "fmt"

func main() {
	var week_days [7] string = [7] string { "Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat" }
	var working_days [] string = week_days[1:6];

	fmt.Println("Working days are:");
	for _, day := range working_days {
		fmt.Print(day, "\n");
	} 

	slice1 := make([] int, 4, 8);
	fmt.Printf("About slice: len %d, cap %d\n", len(slice1), cap(slice1));
}
