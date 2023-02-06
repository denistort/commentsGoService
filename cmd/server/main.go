package main

import "fmt"

// Run is going to responsible for
// Instantiation and startup our go application
func Run() error {
	return nil
}
func main() {
	fmt.Println("Go res clie")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
