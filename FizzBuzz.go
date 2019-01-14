package main

func main(){
	for idx := 0; idx < 101; idx++ {
		if idx % 3 == 0 {
			println("fizz");
		} else if idx % 5 == 0 {
			println("buzz")
		} else if idx % 15 == 0 {
			println("fizz buzz")
		}
	}
}