// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/we7coreteam/gorm-gen-yaml/output/entity"
)

func newTag(db *gorm.DB, opts ...gen.DOOption) tag {
	_tag := tag{}

	_tag.tagDo.UseDB(db, opts...)
	_tag.tagDo.UseModel(&entity.Tag{})

	tableName := _tag.tagDo.TableName()
	_tag.ALL = field.NewAsterisk(tableName)
	_tag.ID = field.NewInt32(tableName, "id")
	_tag.Name = field.NewString(tableName, "name")

	_tag.fillFieldMap()

	return _tag
}

type tag struct {
	tagDo

	ALL  field.Asterisk
	ID   field.Int32
	Name field.String

	fieldMap map[string]field.Expr
}

func (t tag) Table(newTableName string) *tag {
	t.tagDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tag) As(alias string) *tag {
	t.tagDo.DO = *(t.tagDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tag) updateTableName(table string) *tag {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.Name = field.NewString(table, "name")

	t.fillFieldMap()

	return t
}

func (t *tag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tag) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 2)
	t.fieldMap["id"] = t.ID
	t.fieldMap["name"] = t.Name
}

func (t tag) clone(db *gorm.DB) tag {
	t.tagDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tag) replaceDB(db *gorm.DB) tag {
	t.tagDo.ReplaceDB(db)
	return t
}

type tagDo struct{ gen.DO }

type ITagDo interface {
	gen.SubQuery
	Debug() ITagDo
	WithContext(ctx context.Context) ITagDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITagDo
	WriteDB() ITagDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITagDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITagDo
	Not(conds ...gen.Condition) ITagDo
	Or(conds ...gen.Condition) ITagDo
	Select(conds ...field.Expr) ITagDo
	Where(conds ...gen.Condition) ITagDo
	Order(conds ...field.Expr) ITagDo
	Distinct(cols ...field.Expr) ITagDo
	Omit(cols ...field.Expr) ITagDo
	Join(table schema.Tabler, on ...field.Expr) ITagDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITagDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITagDo
	Group(cols ...field.Expr) ITagDo
	Having(conds ...gen.Condition) ITagDo
	Limit(limit int) ITagDo
	Offset(offset int) ITagDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITagDo
	Unscoped() ITagDo
	Create(values ...*entity.Tag) error
	CreateInBatches(values []*entity.Tag, batchSize int) error
	Save(values ...*entity.Tag) error
	First() (*entity.Tag, error)
	Take() (*entity.Tag, error)
	Last() (*entity.Tag, error)
	Find() ([]*entity.Tag, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Tag, err error)
	FindInBatches(result *[]*entity.Tag, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Tag) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITagDo
	Assign(attrs ...field.AssignExpr) ITagDo
	Joins(fields ...field.RelationField) ITagDo
	Preload(fields ...field.RelationField) ITagDo
	FirstOrInit() (*entity.Tag, error)
	FirstOrCreate() (*entity.Tag, error)
	FindByPage(offset int, limit int) (result []*entity.Tag, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITagDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tagDo) Debug() ITagDo {
	return t.withDO(t.DO.Debug())
}

func (t tagDo) WithContext(ctx context.Context) ITagDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tagDo) ReadDB() ITagDo {
	return t.Clauses(dbresolver.Read)
}

func (t tagDo) WriteDB() ITagDo {
	return t.Clauses(dbresolver.Write)
}

func (t tagDo) Session(config *gorm.Session) ITagDo {
	return t.withDO(t.DO.Session(config))
}

func (t tagDo) Clauses(conds ...clause.Expression) ITagDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tagDo) Returning(value interface{}, columns ...string) ITagDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tagDo) Not(conds ...gen.Condition) ITagDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tagDo) Or(conds ...gen.Condition) ITagDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tagDo) Select(conds ...field.Expr) ITagDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tagDo) Where(conds ...gen.Condition) ITagDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tagDo) Order(conds ...field.Expr) ITagDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tagDo) Distinct(cols ...field.Expr) ITagDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tagDo) Omit(cols ...field.Expr) ITagDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tagDo) Join(table schema.Tabler, on ...field.Expr) ITagDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tagDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITagDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tagDo) RightJoin(table schema.Tabler, on ...field.Expr) ITagDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tagDo) Group(cols ...field.Expr) ITagDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tagDo) Having(conds ...gen.Condition) ITagDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tagDo) Limit(limit int) ITagDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tagDo) Offset(offset int) ITagDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITagDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tagDo) Unscoped() ITagDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tagDo) Create(values ...*entity.Tag) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tagDo) CreateInBatches(values []*entity.Tag, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tagDo) Save(values ...*entity.Tag) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tagDo) First() (*entity.Tag, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Tag), nil
	}
}

func (t tagDo) Take() (*entity.Tag, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Tag), nil
	}
}

func (t tagDo) Last() (*entity.Tag, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Tag), nil
	}
}

func (t tagDo) Find() ([]*entity.Tag, error) {
	result, err := t.DO.Find()
	return result.([]*entity.Tag), err
}

func (t tagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Tag, err error) {
	buf := make([]*entity.Tag, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tagDo) FindInBatches(result *[]*entity.Tag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tagDo) Attrs(attrs ...field.AssignExpr) ITagDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tagDo) Assign(attrs ...field.AssignExpr) ITagDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tagDo) Joins(fields ...field.RelationField) ITagDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tagDo) Preload(fields ...field.RelationField) ITagDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tagDo) FirstOrInit() (*entity.Tag, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Tag), nil
	}
}

func (t tagDo) FirstOrCreate() (*entity.Tag, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Tag), nil
	}
}

func (t tagDo) FindByPage(offset int, limit int) (result []*entity.Tag, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tagDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tagDo) Delete(models ...*entity.Tag) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tagDo) withDO(do gen.Dao) *tagDo {
	t.DO = *do.(*gen.DO)
	return t
}