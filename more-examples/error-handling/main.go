package main

import "fmt"

func handlePanic() {
	if a := recover(); a != nil {
		fmt.Println("3. RECOVER from - ", a)

	}
	fmt.Println("4. I am stopping here by closing all connections")
}
func entry(language *string, name *string) {
	      //2
	fmt.Println("1-Entry Regular") //1
	defer handlePanic() 
	defer doCleanUp()
	if language == nil {
		panic("Error: Language Cannot be nil")
	}
	if name == nil {
		panic("panic- because of name") //4
	}
	fmt.Println("language", language)
	fmt.Println("name", name)
}
func doCleanUp() {
	fmt.Println("2.defer-entry-cleanup")
}
func main() {
	// defer fmt.Println("i am the first Defer")
	// defer fmt.Println("6.defer- in main")
	// lang := "go lang"
	// entry(&lang, nil)
	// fmt.Println("5.main- normal", lang)
	// x:= [...]int{1,2,3}
	// y:=x[0:3]
	// y[0]= 10
	
	// fmt.Println(y)
	// fmt.Println(x);
	counter:=100
	P :=&counter
	fmt.Println(P)
	fmt.Println(*P)



}
