# SQL Query Builder - ORM Wrapper

[![Build Status](http://ci.loeffel.io/api/badges/loeffel-io/sql/status.svg)](http://ci.loeffel.io/loeffel-io/sql)
[![Go Report Card](https://goreportcard.com/badge/github.com/loeffel-io/sql)](https://goreportcard.com/report/github.com/loeffel-io/sql)

- Full flexibility
- Zero third-party dependencies
- Useable by any ORM
- Support for Select, From, Join, Where, OrderBy

### [Gorm](https://github.com/jinzhu/gorm) Usage

```go
subquery := sql.Create().
    Select(true, "purchases.*").
    Select(true, "...").
    From(true, "purchases").
    Join(true, "transactions ON transactions.purchase_id=purchases.id")

query := sql.Create().Select(true, "*").
    From(true, "(?) purchases", gorm.Expr(subquery.GetSQL(), subquery.GetValues()...)).
    Join(true, "transactions ON transactions.id=purchases.last_transaction_id")

db.
    Raw(query.GetSQL(), query.GetValues()...).
    Offset(...).
    Limit(...).
    Order(...).
    Unscoped().
    Find(&purchases).
    Error
```
