package sql

import (
	"strings"
	"sync"
)

type Sql struct {
	data []string
	*sync.RWMutex
}

func Create() *Sql {
	return &Sql{
		data:    nil,
		RWMutex: new(sync.RWMutex),
	}
}

func (sql *Sql) Add(statement string, pass bool) {
	sql.Lock()
	defer sql.Unlock()

	if !pass {
		return
	}

	sql.data = append(sql.data, statement)
}

func (sql *Sql) GetData() []string {
	sql.RLock()
	defer sql.RUnlock()

	return sql.data
}

func (sql *Sql) ToString() string {
	return strings.Join(sql.GetData(), " ")
}
