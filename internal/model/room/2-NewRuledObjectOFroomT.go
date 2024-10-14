package room

// ---------------------------------------
func NewRuledObjectOFroomT(name rT, connectionSlice []rT, isEndroom bool) (*roomT, []errT) { //Constructor=factory function=builder
	r := newPlainRoom()

	if err := r.set.name(name); err != nil {
		return nil, logger.Act.Err.Rlog(_NewRuledRoom_10, err, "set.name(name):", "name:", name)
	}
	if err := r.set.connectionSlice(connectionSlice); err != nil {
		return nil, logger.Act.Err.Rlog(_NewRuledRoom_20, err, "r.set.connectionSlice")
	}

	//if the room is start or end room
	if name == _START_ROOM_NAME || isEndroom {

		if err := r.set.allSeats(_MAX_SEATS_START_END); err != nil {
			return nil, logger.Act.Err.Rlog(_NewRuledRoom_30, err, "r.set.allSeats")
		}
		if err := r.set.usedSeats(_USED_SEATS_START_END); err != nil {
			return nil, logger.Act.Err.Rlog(_NewRuledRoom_40, err, "r.set.usedSeats")
		}
	} else {
		if err := r.set.allSeats(_ALL_SEATS_NORMAL_ROOM); err != nil {
			return nil, logger.Act.Err.Rlog(_NewRuledRoom_50, err, "r.set.allSeats")
		}
		if err := r.set.usedSeats(0); err != nil {
			return nil, logger.Act.Err.Rlog(_NewRuledRoom_60, err, "r.set.usedSeats(0)")
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
