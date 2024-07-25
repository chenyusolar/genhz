// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"towin/biz/dal/model"
)

func newAddress(db *gorm.DB, opts ...gen.DOOption) address {
	_address := address{}

	_address.addressDo.UseDB(db, opts...)
	_address.addressDo.UseModel(&model.Address{})

	tableName := _address.addressDo.TableName()
	_address.ALL = field.NewAsterisk(tableName)
	_address.ID = field.NewInt64(tableName, "id")
	_address.Address = field.NewString(tableName, "address")
	_address.AddressType = field.NewString(tableName, "address_type")
	_address.AddressOwner = field.NewInt64(tableName, "address_owner")
	_address.Private = field.NewString(tableName, "private")
	_address.CreatedAt = field.NewTime(tableName, "created_at")
	_address.UpdatedAt = field.NewTime(tableName, "updated_at")
	_address.DeletedAt = field.NewField(tableName, "deleted_at")

	_address.fillFieldMap()

	return _address
}

// address Address information table
type address struct {
	addressDo addressDo

	ALL          field.Asterisk
	ID           field.Int64  // PK
	Address      field.String // Address
	AddressType  field.String // Address type
	AddressOwner field.Int64  // Address belongs to
	Private      field.String // Address private key
	CreatedAt    field.Time   // Address create time
	UpdatedAt    field.Time   // Address update time
	DeletedAt    field.Field  // Address delete time

	fieldMap map[string]field.Expr
}

func (a address) Table(newTableName string) *address {
	a.addressDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a address) As(alias string) *address {
	a.addressDo.DO = *(a.addressDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *address) updateTableName(table string) *address {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Address = field.NewString(table, "address")
	a.AddressType = field.NewString(table, "address_type")
	a.AddressOwner = field.NewInt64(table, "address_owner")
	a.Private = field.NewString(table, "private")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *address) WithContext(ctx context.Context) *addressDo { return a.addressDo.WithContext(ctx) }

func (a address) TableName() string { return a.addressDo.TableName() }

func (a address) Alias() string { return a.addressDo.Alias() }

func (a address) Columns(cols ...field.Expr) gen.Columns { return a.addressDo.Columns(cols...) }

func (a *address) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *address) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["address"] = a.Address
	a.fieldMap["address_type"] = a.AddressType
	a.fieldMap["address_owner"] = a.AddressOwner
	a.fieldMap["private"] = a.Private
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
}

func (a address) clone(db *gorm.DB) address {
	a.addressDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a address) replaceDB(db *gorm.DB) address {
	a.addressDo.ReplaceDB(db)
	return a
}

type addressDo struct{ gen.DO }

func (a addressDo) Debug() *addressDo {
	return a.withDO(a.DO.Debug())
}

func (a addressDo) WithContext(ctx context.Context) *addressDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a addressDo) ReadDB() *addressDo {
	return a.Clauses(dbresolver.Read)
}

func (a addressDo) WriteDB() *addressDo {
	return a.Clauses(dbresolver.Write)
}

func (a addressDo) Session(config *gorm.Session) *addressDo {
	return a.withDO(a.DO.Session(config))
}

func (a addressDo) Clauses(conds ...clause.Expression) *addressDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a addressDo) Returning(value interface{}, columns ...string) *addressDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a addressDo) Not(conds ...gen.Condition) *addressDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a addressDo) Or(conds ...gen.Condition) *addressDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a addressDo) Select(conds ...field.Expr) *addressDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a addressDo) Where(conds ...gen.Condition) *addressDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a addressDo) Order(conds ...field.Expr) *addressDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a addressDo) Distinct(cols ...field.Expr) *addressDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a addressDo) Omit(cols ...field.Expr) *addressDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a addressDo) Join(table schema.Tabler, on ...field.Expr) *addressDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a addressDo) LeftJoin(table schema.Tabler, on ...field.Expr) *addressDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a addressDo) RightJoin(table schema.Tabler, on ...field.Expr) *addressDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a addressDo) Group(cols ...field.Expr) *addressDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a addressDo) Having(conds ...gen.Condition) *addressDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a addressDo) Limit(limit int) *addressDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a addressDo) Offset(offset int) *addressDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a addressDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *addressDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a addressDo) Unscoped() *addressDo {
	return a.withDO(a.DO.Unscoped())
}

func (a addressDo) Create(values ...*model.Address) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a addressDo) CreateInBatches(values []*model.Address, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a addressDo) Save(values ...*model.Address) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a addressDo) First() (*model.Address, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) Take() (*model.Address, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) Last() (*model.Address, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) Find() ([]*model.Address, error) {
	result, err := a.DO.Find()
	return result.([]*model.Address), err
}

func (a addressDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Address, err error) {
	buf := make([]*model.Address, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a addressDo) FindInBatches(result *[]*model.Address, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a addressDo) Attrs(attrs ...field.AssignExpr) *addressDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a addressDo) Assign(attrs ...field.AssignExpr) *addressDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a addressDo) Joins(fields ...field.RelationField) *addressDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a addressDo) Preload(fields ...field.RelationField) *addressDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a addressDo) FirstOrInit() (*model.Address, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) FirstOrCreate() (*model.Address, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Address), nil
	}
}

func (a addressDo) FindByPage(offset int, limit int) (result []*model.Address, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a addressDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a addressDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a addressDo) Delete(models ...*model.Address) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *addressDo) withDO(do gen.Dao) *addressDo {
	a.DO = *do.(*gen.DO)
	return a
}
