package room

// ========================================
const ( //func or method error code
	_ERR_Null errT = iota

	_ERR_NewRuledRoom_10
	_ERR_NewRuledRoom_20
	_ERR_NewRuledRoom_30
	_ERR_NewRuledRoom_40
	_ERR_NewRuledRoom_50
	_ERR_NewRuledRoom_60

	//roomTsetT
	_ERR_roomT_setT_name_10
	//roomTgetT
	_ERR_roomT_getT_OneRandomNextRoom_10

	//roomTactT
	_ERR_roomT_actT_UpdateOneCome_10
	_ERR_roomT_actT_UpdateOneLeave_10
)

// -------------------------------
var ErrCodeDb = map[errT]string{ //for log purpose
	//room builder
	_ERR_NewRuledRoom_10: "NewRuledRoom.10:",
	_ERR_NewRuledRoom_20: "NewRuledRoom.20:",
	_ERR_NewRuledRoom_30: "NewRuledRoom.30:",
	_ERR_NewRuledRoom_40: "NewRuledRoom.40:",
	_ERR_NewRuledRoom_50: "NewRuledRoom.50:",
	_ERR_NewRuledRoom_60: "NewRuledRoom.60:",
	//Set
	_ERR_roomT_setT_name_10: "roomTsetT_name_10:",
	//roomTgetT
	_ERR_roomT_getT_OneRandomNextRoom_10: "roomTgetT_OneRandomNextRoom_10:",

	//roomTactT
	_ERR_roomT_actT_UpdateOneCome_10: "roomTactT_UpdateOneCome_10:",

	_ERR_roomT_actT_UpdateOneLeave_10: "roomTactT_UpdateOneLeave_10:",
}

//------------------------------------------
