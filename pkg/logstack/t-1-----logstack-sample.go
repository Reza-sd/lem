package logstack

var (

	SampleLogger  = LogCollector{
		PackageName: "pkgName",
		LogToFile: false,
		LogToCli: true,
		
	}
)