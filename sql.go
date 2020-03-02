package sql

import (
	"fmt"
	"strings"
	"sync"
)

const (
	Select   = "select"
	From     = "from"
	Join     = "join"
	LeftJoin = "left join"
	Where    = "where"
	GroupBy  = "group by"
)

type Sql struct {
	data map[string][]*Data
	*sync.RWMutex
}

func Create() *Sql {
	return &Sql{
		data:    make(map[string][]*Data),
		RWMutex: new(sync.RWMutex),
	}
}

func (sql *Sql) add(pass bool, category string, statement string, values []interface{}) *Sql {
	sql.Lock()
	defer sql.Unlock()

	sql.data[category] = append(sql.data[category], &Data{
		Statement: statement,
		Values:    values,
		Category:  category,
		RWMutex:   new(sync.RWMutex),
	})

	return sql
}

func (sql *Sql) GetData() map[string][]*Data {
	sql.RLock()
	defer sql.RUnlock()

	return sql.data
}

func (sql *Sql) groupOrder() []string {
	return []string{
		Select,
		From,
		Join,
		LeftJoin,
		Where,
		GroupBy,
	}
}

func (sql *Sql) GetSQL() string {
	var statements []string
	var sqlData = sql.GetData()

	for _, group := range sql.groupOrder() {
		if (sqlData[group]) == nil {
			continue
		}

		for index, data := range sqlData[group] {
			if data == nil {
				continue
			}

			if index == 0 {
				statements = append(statements, data.getCategory())
				statements = append(statements, data.getStatement())
				continue
			}

			switch data.getCategory() {
			case Where:
				statements = append(statements, fmt.Sprintf("AND %s", data.getStatement()))
			case Join, LeftJoin:
				statements = append(statements, data.getCategory())
				statements = append(statements, data.getStatement())
			case Select, From, GroupBy:
				statements = append(statements, fmt.Sprintf(", %s", data.getStatement()))
			}
		}
	}

	return strings.Join(statements, " ")
}

func (sql *Sql) GetValues() []interface{} {
	var values []interface{}
	var sqlData = sql.GetData()

	for _, group := range sql.groupOrder() {
		if (sqlData[group]) == nil {
			continue
		}

		for _, data := range sqlData[group] {
			if data == nil {
				continue
			}

			values = append(values, data.getValues()...)
		}
	}

	return values
}

func (sql *Sql) Select(pass bool, statement string, values ...interface{}) *Sql {
	if !pass {
		return sql
	}

	return sql.add(pass, Select, statement, values)
}

func (sql *Sql) From(pass bool, statement string, values ...interface{}) *Sql {
	if !pass {
		return sql
	}

	return sql.add(pass, From, statement, values)
}

func (sql *Sql) Join(pass bool, statement string, values ...interface{}) *Sql {
	if !pass {
		return sql
	}

	return sql.add(pass, Join, statement, values)
}

func (sql *Sql) LeftJoin(pass bool, statement string, values ...interface{}) *Sql {
	if !pass {
		return sql
	}

	return sql.add(pass, LeftJoin, statement, values)
}

func (sql *Sql) Where(pass bool, statement string, values ...interface{}) *Sql {
	if !pass {
		return sql
	}

	return sql.add(pass, Where, statement, values)
}

func (sql *Sql) GroupBy(pass bool, statement string, values ...interface{}) *Sql {
	if !pass {
		return sql
	}

	return sql.add(pass, GroupBy, statement, values)
}
