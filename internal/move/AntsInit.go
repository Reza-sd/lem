package move
import (
	"fmt"
)

func (myAnts *Ants)AntsInit (num int) {

	myAnts.Ants = make(map[int]Ant)

for i :=1; i <= num; i++ {
	name:=fmt.Sprintf("L%d", i)
	myAnts.Ants[i] = Ant{Name: name}

}


}