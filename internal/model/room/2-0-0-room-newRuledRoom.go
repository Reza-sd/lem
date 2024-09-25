package room

// -------------------------------------------------------------
func newPlainRoom() *room { //Constructor=factory function=builder

	r := &room{}

	r.data.connectionSlice = []RT{}

	r.Get.room = r
	r.set.room = r
	r.Act.room = r

	return r
}

// ------------------------------------
// type rmBuildArg struct {
// 	name            RT
// 	endRoomName     RT
// 	connectionSlice []RT
// }

func NewRuledRoom(name RT, endRoomName RT, connectionSlice []RT) (*room, errArrT) { //Constructor=factory function=builder
	r := newPlainRoom()

	if err := r.set.name(name); err != nil {
		return nil, logger.Err.Rlog(1, err)
	}
	if err := r.set.connectionSlice(connectionSlice); err != nil {
		return nil, logger.Err.Rlog(2, err)
	}

	if name == startRoomName || name == endRoomName {

		if err := r.set.allSeats(MaxSeatsStartEnd); err != nil {
			return nil, logger.Err.Rlog(3, err)
		}
		if err := r.set.usedSeats(UsedSeatsStartEnd); err != nil {
			return nil, logger.Err.Rlog(4, err)
		}
	} else {
		if err := r.set.allSeats(AllSeatsNormalRoom); err != nil {
			return nil, logger.Err.Rlog(5, err)
		}
		if err := r.set.usedSeats(0); err != nil {
			return nil, logger.Err.Rlog(6, err)
		}
	}

	return r, nil
	// if its first then? if end then

}

//---------------------------------------
