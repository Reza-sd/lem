package logstack

var (

	SampleLogger  = LogCollector{
		PackageName: pkgName,
		LogToFile: true,
		LogToCli: true,
		
	}
)