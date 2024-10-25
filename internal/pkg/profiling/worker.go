package profiling

import (
	"os"
	"runtime/pprof"
)

func Start(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		return err
	}

	return nil
}
