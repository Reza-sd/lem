package room

// ========================================
const ( //func or method error code
	_CODE_Null errT = iota

	_CODE_NewRuledRoom_10
	_CODE_NewRuledRoom_20
	_CODE_NewRuledRoom_30
	_CODE_NewRuledRoom_40
	_CODE_NewRuledRoom_50
	_CODE_NewRuledRoom_60

	//roomTsetT
	_CODE_roomTsetT_name_10
	//roomTgetT
	_CODE_roomTgetT_OneRandomNextRoom_10

	//roomTactT
	_CODE_roomTactT_UpdateOneCome_10
	_CODE_roomTactT_UpdateOneLeave_10
)

// -------------------------------
var ErrCodeDb = map[errT]string{ //for log purpose
	//room builder
	_CODE_NewRuledRoom_10: "NewRuledRoom.10:",
	_CODE_NewRuledRoom_20: "NewRuledRoom.20:",
	_CODE_NewRuledRoom_30: "NewRuledRoom.30:",
	_CODE_NewRuledRoom_40: "NewRuledRoom.40:",
	_CODE_NewRuledRoom_50: "NewRuledRoom.50:",
	_CODE_NewRuledRoom_60: "NewRuledRoom.60:",
	//Set
	_CODE_roomTsetT_name_10: "roomTsetT_name_10:",
	//roomTgetT
	_CODE_roomTgetT_OneRandomNextRoom_10: "roomTgetT_OneRandomNextRoom_10:",

	//roomTactT
	_CODE_roomTactT_UpdateOneCome_10: "roomTactT_UpdateOneCome_10:",

	_CODE_roomTactT_UpdateOneLeave_10: "roomTactT_UpdateOneLeave_10:",
}

//------------------------------------------
