package watcher

type Hook struct {
	CreatedDir       func(path string)
	CreatedFile      func(path string)
	RemovedOrRenamed func(path string)
}
