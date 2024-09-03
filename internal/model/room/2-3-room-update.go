package room

import "fmt"

// =====================================================
func (u *updater) oneCome() answer[*room] {

	if u.room.get.usedSeats() == u.room.get.allSeats() || u.room.get.usedSeats()+1 > u.room.get.allSeats() {
		return answer[*room]{u.room, fmt.Errorf("full")}
	}

	u.room.set.usedSeats(u.room.get.usedSeats() + 1)
	return answer[*room]{u.room, nil}
}

// =====================================================
func (u *updater) oneLeave() answer[*room] {
	if u.room.get.usedSeats() == 0 {
		return answer[*room]{u.room, fmt.Errorf("full")}
	}
	u.room.set.usedSeats(u.room.get.usedSeats() - 1)
	return answer[*room]{u.room, nil}
}

// =====================================================
