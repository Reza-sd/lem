package room

// ---------------------------------------
func NewRuledObjectOFroomT(name rT, connectionSlice []rT, isEndroom bool) (*roomT, []errT) { //Constructor=factory function=builder
	r := newPlainRoom()

	if err := r.set.name(name); err != nil {
		return nil, logger.Err.Rlog(NewRuledRoom_10, err, "name:", name)
	}
	if err := r.set.connectionSlice(connectionSlice); err != nil {
		return nil, logger.Err.Rlog(NewRuledRoom_20, err)
	}

	//if the room is start or end room
	if name == startRoomName || isEndroom {

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

// ---------------------------------------------------
func newPlainRoom() *roomT { //Constructor=factory function=builder

	r := &roomT{}

	r.data.connectionSlice = []rT{}

	r.Get.room = r
	r.set.room = r
	r.Act.room = r

	return r
}

// ------------------------------------
