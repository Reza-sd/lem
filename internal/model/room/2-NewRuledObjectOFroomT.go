package room

// ---------------------------------------
func Return_NewRuledObjectOFroomT(name_param rT, connectionSlice_param []rT, isEndroom_param bool) (*roomT, []errT) { //Constructor=factory function=builder
	myRoom := return_NewPlainRoom()

	if err_set_name := myRoom.set.name(name_param); err_set_name != nil {
		return nil, logger.Act.Err.Rlog(_ERR_NewRuledRoom_10, err_set_name, "set.name(name):", "name:", name_param)
	}

	if err_set_connectionSlice := myRoom.set.connectionSlice(connectionSlice_param); err_set_connectionSlice != nil {
		return nil, logger.Act.Err.Rlog(_ERR_NewRuledRoom_20, err_set_connectionSlice, "r.set.connectionSlice")
	}

	//if the room is start or end room
	if name_param == _START_ROOM_NAME || isEndroom_param {

		if err_set_allSeats := myRoom.set.allSeats(_MAX_SEATS_START_END); err_set_allSeats != nil {
			return nil, logger.Act.Err.Rlog(_ERR_NewRuledRoom_30, err_set_allSeats, "r.set.allSeats")
		}

		if err_set_usedSeats := myRoom.set.usedSeats(_USED_SEATS_START_END); err_set_usedSeats != nil {
			return nil, logger.Act.Err.Rlog(_ERR_NewRuledRoom_40, err_set_usedSeats, "r.set.usedSeats")
		}

	} else {

		if err_set_allSeats := myRoom.set.allSeats(_ALL_SEATS_NORMAL_ROOM); err_set_allSeats != nil {
			return nil, logger.Act.Err.Rlog(_ERR_NewRuledRoom_50, err_set_allSeats, "r.set.allSeats")
		}

		if err_set_usedSeats := myRoom.set.usedSeats(_EMPTY_USED_SEATS); err_set_usedSeats != nil {
			return nil, logger.Act.Err.Rlog(_ERR_NewRuledRoom_60, err_set_usedSeats, "r.set.usedSeats(_EMPTY_USED_SEATS)", "_EMPTY_USED_SEATS=", _EMPTY_USED_SEATS)
		}

	}

	return myRoom, nil

}

// ---------------------------------------------------
func return_NewPlainRoom() *roomT { //Constructor=factory function=builder

	r := &roomT{}

	r.data.connectionSlice = []rT{}

	r.Get.room = r
	r.set.room = r
	r.Act.room = r

	return r
}

// ------------------------------------
