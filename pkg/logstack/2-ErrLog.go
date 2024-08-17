package logstack

import (
//"fmt"
//"log/slog"
)

// ---------------------------------
func (l *LogCollector) ErrLog(FuncName string, OperationName string, err interface{}, operationDescription string) {
	msg, agrs := LogMsgGenerator(l.PackageName, FuncName, OperationName, operationDescription, err)
	//msg=fmt.Sprintf("%v-%v", "ErrLog",msg)

	if l.LogToCli {
		loggerToCli.Error(msg, agrs)

	}

	if l.LogToFile {
		loggerToFile.Error(msg, agrs)
	}

}

// ---------------------------------
