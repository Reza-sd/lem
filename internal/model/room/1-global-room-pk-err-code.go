package room

// ========================================
const ( //func or method error code
	Null errT = iota

	NewRuledRoom_10
	NewRuledRoom_20
	NewRuledRoom_30
	NewRuledRoom_40
	NewRuledRoom_50
	NewRuledRoom_60

	//roomTsetT
	roomTsetT_name_10
	//roomTgetT
	roomTgetT_OneRandomNextRoom_10

	//roomTactT
	roomTactT_UpdateOneCome_10
	roomTactT_UpdateOneLeave_10
)

// -------------------------------
var ErrCodeDb = map[errT]string{ //for log purpose
	//room builder
	NewRuledRoom_10: "NewRuledRoom.10:",
	NewRuledRoom_20: "NewRuledRoom.20:",
	NewRuledRoom_30: "NewRuledRoom.30:",
	NewRuledRoom_40: "NewRuledRoom.40:",
	NewRuledRoom_50: "NewRuledRoom.50:",
	NewRuledRoom_60: "NewRuledRoom.60:",
	//Set
	roomTsetT_name_10: "roomTsetT_name_10:",
	//roomTgetT
	roomTgetT_OneRandomNextRoom_10: "roomTgetT_OneRandomNextRoom_10:",

	//roomTactT
	roomTactT_UpdateOneCome_10: "roomTactT_UpdateOneCome_10:",

	roomTactT_UpdateOneLeave_10: "roomTactT_UpdateOneLeave_10:",
}

//------------------------------------------
