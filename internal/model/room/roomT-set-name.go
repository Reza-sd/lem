package room

// ---------------------------------------------
func (set *roomTsetT) name(name rT) []errT {
	//Guard clause
	if name > maxName {
		return logger.Act.Err.Rlog(roomTsetT_name_10, nil, "name:", name, ">", "maxName:", maxName)
	}
	//check if name valid to set
	set.room.data.name = name
	return nil

}

//---------------------------------------------
