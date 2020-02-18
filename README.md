# SQL Query Builder - ORM Wrapper

- Full flexibility
- Zero third-party dependencies
- Useable by any ORM
- Support for Select, From, Join, Where, OrderBy

### [Gorm](https://github.com/jinzhu/gorm) Usage

```go
subquery := sql.Create().
    Select(true, "purchases.*").
    From(true, "purchases").
    Join(true, "transactions ON transactions.purchase_id=purchases.id")

query := sql.Create().Select(true, "*").
    From(true, "(?) purchases", gorm.Expr(subquery.GetSQL(), subquery.GetValues()...))

db.
    Raw(query.GetSQL(), query.GetValues()...).
    Offset(...).
    Limit(...).
    Order(...).
    Unscoped().
    Find(&purchases).
    Error
```
