package main

import "fmt"

type Student struct {
	name  []string 
	score []int
}

func (s Student) Avarage() float64 {
}
func (s Student) Max() (max int, name string) {

}

func main() {
	var a = Student{}
	
for i := 0; i < 6; i++ {
	var name string 
	fmt.Print("Input"+ string(i) + " Student' s Name : ")
	fmt.Scan(&name)
	a.name = append(a.name, name)
	var score int 
	fmt.Print("Input " + name + " Score : ")
	fmt.Scan(&score)
	a.score = append(a.score, score)
}

fmt.Println("\n\nAvarage Score Student is", a.Avarage())
scoreMax, nameMax, nameMax := a.Max()
fmt.Println("Max score Students is : "+nameMax" (", scoreMax, ")")
scoreMin, nameMin := a.Min()
fmt.println("Min Score Students is : "+nameMin+" (" scoreMin, ")")
}