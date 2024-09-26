package room

// ========================================
const ( //func or method error code
	Null errT = iota
	PkgName

	NewRuledRoom_0
	NewRuledRoom_10
	NewRuledRoom_20
	NewRuledRoom_30
	NewRuledRoom_40
	NewRuledRoom_50
	NewRuledRoom_60

	//set
	roomT_setT_name
	roomT_setT_name_10
	//Get
	roomT_getT_OneRandomNextRoom    //func name
	roomT_getT_OneRandomNextRoom_10 //EmptyConnectionSlice

	//Act
	roomT_actT_UpdateOneCome
	roomT_actT_UpdateOneCome_10 //OverCap

	roomT_actT_UpdateOneLeave
	roomT_actT_UpdateOneLeave_10 //OverCap
)

var ErrCodeDes = map[errT]string{ //for log purpose
	Null:    "Null",
	PkgName: pkgName_,
	//room builder
	NewRuledRoom_0:  _NewRuledRoom_0,
	NewRuledRoom_10: _NewRuledRoom_10,
	NewRuledRoom_20: _NewRuledRoom_20,
	NewRuledRoom_30: _NewRuledRoom_30,
	NewRuledRoom_40: _NewRuledRoom_40,
	NewRuledRoom_50: _NewRuledRoom_50,
	NewRuledRoom_60: _NewRuledRoom_60,
	//Set
	roomT_setT_name:    _set_name,
	roomT_setT_name_10: _set_name_10,
	//Get
	roomT_getT_OneRandomNextRoom:    _roomT_getT_OneRandomNextRoom,
	roomT_getT_OneRandomNextRoom_10: _roomT_getT_OneRandomNextRoom_10,

	//Act
	roomT_actT_UpdateOneCome:    _Act_UpdateOneCome,
	roomT_actT_UpdateOneCome_10: _Act_UpdateOneCome10,

	roomT_actT_UpdateOneLeave:    _Act_UpdateOneLeave,
	roomT_actT_UpdateOneLeave_10: _Act_UpdateOneLeave10,
}
