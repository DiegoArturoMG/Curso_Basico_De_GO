package main

import "fmt"

type pc struct {
	ram int
	disk int
	brand string
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	
	myPc := pc{ram:16, disk: 1000, brand: "msi"}
	fmt.Println(myPc)
}

