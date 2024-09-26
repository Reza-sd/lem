package room

// ================================
const (
	_set_name    = "set.name."
	_set_name_10 = _set_name + "10:exceed max Name"
)

// ---------------------------------------------
func (set *setT) name(name rT) []errT {
	//Guard clause
	if name > maxName {
		return logger.Err.Rlog(set_name_10, nil, "name:", name, ">", "maxName:", maxName)
	}
	//check if name valid to set
	set.room.data.name = name
	return nil
}

//---------------------------------------------
