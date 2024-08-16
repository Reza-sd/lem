package antgroup

type AntGroup struct {
	NumberOfAnts          Mtag
	CurrentSequenceNumber Mtag
	AntsDb                map[Mtag]Ant
	//UsedTunnelinThisSeq	map[string]string//
	UsedTunnel         map[Mtag][]Mtag
	NotArrivedAntsName map[Mtag]struct{}
}
