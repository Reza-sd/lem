package room

// ---------------------------------------------
func (set *roomTsetT) name(name rT) []errT {
	//Guard clause
	if name > _MAX_NAME {
		return logger.Act.Err.Rlog(_roomTsetT_name_10, nil, "if name > maxName ", "name:", name, ">", "maxName:", _MAX_NAME)
	}
	//check if name valid to set
	set.room.data.name = name
	return nil

}

//---------------------------------------------
