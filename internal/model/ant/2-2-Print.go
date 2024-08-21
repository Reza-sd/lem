package ant

import (
	"fmt"
)

func (myAnt *Ant) Print() {

	//fmt.Println("\n----------AntGroup------------")
	str, _ := myAnt.ToString()
	fmt.Println(str)
	//fmt.Println("---------------------------")

}
