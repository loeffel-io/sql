package sql

import (
	"strings"
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

func (sql *Sql) Add(pass bool, statement string, values ...interface{}) {
	sql.Lock()
	defer sql.Unlock()

	if !pass {
		return
	}

	sql.addData(&Data{
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

func (sql *Sql) addData(data *Data) {
	sql.Lock()
	defer sql.Unlock()

	sql.data = append(sql.data, data)
}

func (sql *Sql) GetSQL() string {
	var statements []string

	for _, item := range sql.GetData() {
		statements = append(statements, item.getStatement())
	}

	return strings.Join(statements, " ")
}

func (sql *Sql) GetValues() []interface{} {
	var values []interface{}

	for _, item := range sql.GetData() {
		values = append(values, item.getValues()...)
	}

	return values
}
