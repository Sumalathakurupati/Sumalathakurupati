package main

import "fmt"

/*type Student struct {
	RollNumber  int
	StudentName string
}

func main() {

	student1 := Student{
		RollNumber:  301,
		StudentName: "Krishna",
	}
	student2 := Student{
		RollNumber:  302,
		StudentName: "Ram",
	}

	student := map[string]Student{

		"sid1": student1,
		"sid2": student2,
	}
	for sid, data := range student {
		fmt.Println(sid, data)
	}

}*/

func main() {
	employee := map[string]int{
		"John":   60000,
		"Martin": 70000,
	}

	newemployee := "John"
	value, ok := employee[newemployee]
	if ok == true {
		fmt.Println(newemployee, value)
		return
	}
	fmt.Println(newemployee, "Not found")
}
