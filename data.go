package sql

import "sync"

type Data struct {
	Statement string
	Values    []interface{}
	*sync.RWMutex
}

func (data *Data) getStatement() string {
	data.RLock()
	defer data.RUnlock()

	return data.Statement
}

func (data *Data) getValues() []interface{} {
	data.RLock()
	defer data.RUnlock()

	return data.Values
}
