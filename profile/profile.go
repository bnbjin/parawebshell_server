package profile

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

const (
	EnvEnableProfiling = "PWS_PROF"
	CpuProfile         = "pws.cpuprof"
	HeapProfile        = "pws.memprof"
)

/**/
func ProfileIfEnabled() (func(), error) {
	// FIXME this is a temporary hack so profiling of asynchronous operations
	// works as intended.
	if os.Getenv(EnvEnableProfiling) != "" {
		StopProfilingFunc, err := StartProfiling() // TODO maybe change this to its own option... profiling makes it slower.
		if err != nil {
			return nil, err
		}
		return StopProfilingFunc, nil
	}
	return func() {}, nil
}

// startProfiling begins CPU profiling and returns a `stop` function to be
// executed as late as possible. The stop function captures the memprofile.
func StartProfiling() (func(), error) {
	// start CPU profiling as early as possible
	ofi, err := os.Create(CpuProfile)
	if err != nil {
		return nil, err
	}
	pprof.StartCPUProfile(ofi)

	go func() {
		for range time.NewTicker(time.Second * 30).C {
			err := WriteHeapProfileToFile()
			if err != nil {
				log.Panic(err)
			}
		}
	}()

	stopProfiling := func() {
		pprof.StopCPUProfile()
		ofi.Close() // captured by the closure
	}

	return stopProfiling, nil
}

/**/
func WriteHeapProfileToFile() error {
	mprof, err := os.Create(HeapProfile)
	if err != nil {
		return err
	}
	defer mprof.Close() // _after_ writing the heap profile

	return pprof.WriteHeapProfile(mprof)
}
