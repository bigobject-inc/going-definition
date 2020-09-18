package sharedmemory

// Creator shared memory creator
type Creator func(dsn string, params interface{}) (SharedMemory, error)

// SharedMemoriesMap shared memories map
type SharedMemoriesMap map[string]SharedMemory
