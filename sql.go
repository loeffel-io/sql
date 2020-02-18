package sql

import (
	"sync"
)

type Sql struct {
	data []*Data
	*sync.RWMutex
}

func Create() *Sql {
	return &Sql{
		data:    make([]*Data, 0),
		RWMutex: new(sync.RWMutex),
	}
}

func (sql *Sql) Add(statement string, pass bool, values ...interface{}) {
	sql.Lock()
	defer sql.Unlock()

	if !pass {
		return
	}

	sql.AddData(&Data{
		Statement: statement,
		Values:    values,
		RWMutex:   new(sync.RWMutex),
	})
}

func (sql *Sql) GetData() []*Data {
	sql.RLock()
	defer sql.RUnlock()

	return sql.data
}

func (sql *Sql) AddData(data *Data) {
	sql.Lock()
	defer sql.Unlock()

	sql.data = append(sql.data, data)
}
