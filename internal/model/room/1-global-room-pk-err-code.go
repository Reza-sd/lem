package room

// ========================================
const ( //func or method error code
	Null errT = iota

	_NewRuledRoom_10
	_NewRuledRoom_20
	_NewRuledRoom_30
	_NewRuledRoom_40
	_NewRuledRoom_50
	_NewRuledRoom_60

	//roomTsetT
	_roomTsetT_name_10
	//roomTgetT
	_roomTgetT_OneRandomNextRoom_10

	//roomTactT
	_roomTactT_UpdateOneCome_10
	_roomTactT_UpdateOneLeave_10
)

// -------------------------------
var ErrCodeDb = map[errT]string{ //for log purpose
	//room builder
	_NewRuledRoom_10: "NewRuledRoom.10:",
	_NewRuledRoom_20: "NewRuledRoom.20:",
	_NewRuledRoom_30: "NewRuledRoom.30:",
	_NewRuledRoom_40: "NewRuledRoom.40:",
	_NewRuledRoom_50: "NewRuledRoom.50:",
	_NewRuledRoom_60: "NewRuledRoom.60:",
	//Set
	_roomTsetT_name_10: "roomTsetT_name_10:",
	//roomTgetT
	_roomTgetT_OneRandomNextRoom_10: "roomTgetT_OneRandomNextRoom_10:",

	//roomTactT
	_roomTactT_UpdateOneCome_10: "roomTactT_UpdateOneCome_10:",

	_roomTactT_UpdateOneLeave_10: "roomTactT_UpdateOneLeave_10:",
}

//------------------------------------------
