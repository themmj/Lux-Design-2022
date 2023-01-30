package lux

import (
	"encoding/json"
	"fmt"
	"os"
)

func Log(formatString string, args ...any) {
	fmt.Fprintf(os.Stderr, formatString, args...)
	fmt.Fprintf(os.Stderr, "\n")
}

func DumpJsonToFile(filePath string, obj any) {
	b, err := json.Marshal(obj) // TODO: check if raw obj or pointer should be passed
	if err != nil {
		Log("could not marshal object to dump into json file %s", filePath)
		return
	}
	var outFile *os.File
	outFile, err = os.Create(filePath)
	if err != nil {
		Log("could not open file %s to dump json into", filePath)
		return
	}
	fmt.Fprint(outFile, string(b))
}
