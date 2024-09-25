package room

const (
	_NewRuledRoom_0  = "NewRuledRoom."
	_NewRuledRoom_10 = _NewRuledRoom_0 + ""
	_NewRuledRoom_20 = _NewRuledRoom_0 + ""
	_NewRuledRoom_30 = _NewRuledRoom_0 + ""
	_NewRuledRoom_40 = _NewRuledRoom_0 + ""
	_NewRuledRoom_50 = _NewRuledRoom_0 + ""
	_NewRuledRoom_60 = _NewRuledRoom_0 + ""
)

// ---------------------------------------
func NewRuledRoom(name RT, endRoomName RT, connectionSlice []RT) (*room, errArrT) { //Constructor=factory function=builder
	r := newPlainRoom()

	if err := r.set.name(name); err != nil {
		return nil, logger.Err.Rlog(NewRuledRoom_10, err)
	}
	if err := r.set.connectionSlice(connectionSlice); err != nil {
		return nil, logger.Err.Rlog(NewRuledRoom_20, err)
	}

	//if the room is start or end room
	if name == startRoomName || name == endRoomName {

		if err := r.set.allSeats(MaxSeatsStartEnd); err != nil {
			return nil, logger.Err.Rlog(NewRuledRoom_30, err)
		}
		if err := r.set.usedSeats(UsedSeatsStartEnd); err != nil {
			return nil, logger.Err.Rlog(NewRuledRoom_40, err)
		}
	} else {
		if err := r.set.allSeats(AllSeatsNormalRoom); err != nil {
			return nil, logger.Err.Rlog(NewRuledRoom_50, err)
		}
		if err := r.set.usedSeats(0); err != nil {
			return nil, logger.Err.Rlog(NewRuledRoom_60, err)
		}
	}

	return r, nil

}

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
