package helper

// BatchExecutor is use for execute by batch
type BatchExecutor struct {
	Size     int
	executor func(batch []interface{})
	batch    []interface{}
}

func NewBatchExecutor(size int, exec func([]interface{})) *BatchExecutor {
	return &BatchExecutor{
		Size:     size,
		executor: exec,
		batch:    make([]interface{}, 0, size),
	}
}

func (b *BatchExecutor) Process(i interface{}) (isExecuted bool) {
	b.batch = append(b.batch, i)

	if len(b.batch) == b.Size {
		b.executor(b.batch)
		b.batch = b.batch[:0]
		return true
	}

	return false
}

func (b *BatchExecutor) Finish() (isExecuted bool) {
	if len(b.batch) > 0 {
		b.executor(b.batch)
		return true
	}

	return false
}
