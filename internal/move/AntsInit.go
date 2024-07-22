package move
import (
	"fmt"
)

func (myAnts *Ants)AntsInit (num int) {

	myAnts.Ants = make(map[string]Ant)

for i :=1; i <= num; i++ {
	name:=fmt.Sprintf("L%d", i)
	myAnts.Ants[name] = Ant{Name: name}

}


}