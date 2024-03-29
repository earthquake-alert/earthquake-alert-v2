// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TsunamiConnect is an object representing the database table.
type TsunamiConnect struct {
	ID        uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	TsunamiID int       `boil:"tsunami_id" json:"tsunami_id" toml:"tsunami_id" yaml:"tsunami_id"`
	EventID   uint64    `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	Created   time.Time `boil:"created" json:"created" toml:"created" yaml:"created"`

	R *tsunamiConnectR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tsunamiConnectL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TsunamiConnectColumns = struct {
	ID        string
	TsunamiID string
	EventID   string
	Created   string
}{
	ID:        "id",
	TsunamiID: "tsunami_id",
	EventID:   "event_id",
	Created:   "created",
}

var TsunamiConnectTableColumns = struct {
	ID        string
	TsunamiID string
	EventID   string
	Created   string
}{
	ID:        "TsunamiConnects.id",
	TsunamiID: "TsunamiConnects.tsunami_id",
	EventID:   "TsunamiConnects.event_id",
	Created:   "TsunamiConnects.created",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var TsunamiConnectWhere = struct {
	ID        whereHelperuint
	TsunamiID whereHelperint
	EventID   whereHelperuint64
	Created   whereHelpertime_Time
}{
	ID:        whereHelperuint{field: "`TsunamiConnects`.`id`"},
	TsunamiID: whereHelperint{field: "`TsunamiConnects`.`tsunami_id`"},
	EventID:   whereHelperuint64{field: "`TsunamiConnects`.`event_id`"},
	Created:   whereHelpertime_Time{field: "`TsunamiConnects`.`created`"},
}

// TsunamiConnectRels is where relationship names are stored.
var TsunamiConnectRels = struct {
}{}

// tsunamiConnectR is where relationships are stored.
type tsunamiConnectR struct {
}

// NewStruct creates a new relationship struct
func (*tsunamiConnectR) NewStruct() *tsunamiConnectR {
	return &tsunamiConnectR{}
}

// tsunamiConnectL is where Load methods for each relationship are stored.
type tsunamiConnectL struct{}

var (
	tsunamiConnectAllColumns            = []string{"id", "tsunami_id", "event_id", "created"}
	tsunamiConnectColumnsWithoutDefault = []string{"tsunami_id", "event_id"}
	tsunamiConnectColumnsWithDefault    = []string{"id", "created"}
	tsunamiConnectPrimaryKeyColumns     = []string{"id"}
	tsunamiConnectGeneratedColumns      = []string{}
)

type (
	// TsunamiConnectSlice is an alias for a slice of pointers to TsunamiConnect.
	// This should almost always be used instead of []TsunamiConnect.
	TsunamiConnectSlice []*TsunamiConnect
	// TsunamiConnectHook is the signature for custom TsunamiConnect hook methods
	TsunamiConnectHook func(context.Context, boil.ContextExecutor, *TsunamiConnect) error

	tsunamiConnectQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tsunamiConnectType                 = reflect.TypeOf(&TsunamiConnect{})
	tsunamiConnectMapping              = queries.MakeStructMapping(tsunamiConnectType)
	tsunamiConnectPrimaryKeyMapping, _ = queries.BindMapping(tsunamiConnectType, tsunamiConnectMapping, tsunamiConnectPrimaryKeyColumns)
	tsunamiConnectInsertCacheMut       sync.RWMutex
	tsunamiConnectInsertCache          = make(map[string]insertCache)
	tsunamiConnectUpdateCacheMut       sync.RWMutex
	tsunamiConnectUpdateCache          = make(map[string]updateCache)
	tsunamiConnectUpsertCacheMut       sync.RWMutex
	tsunamiConnectUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tsunamiConnectAfterSelectHooks []TsunamiConnectHook

var tsunamiConnectBeforeInsertHooks []TsunamiConnectHook
var tsunamiConnectAfterInsertHooks []TsunamiConnectHook

var tsunamiConnectBeforeUpdateHooks []TsunamiConnectHook
var tsunamiConnectAfterUpdateHooks []TsunamiConnectHook

var tsunamiConnectBeforeDeleteHooks []TsunamiConnectHook
var tsunamiConnectAfterDeleteHooks []TsunamiConnectHook

var tsunamiConnectBeforeUpsertHooks []TsunamiConnectHook
var tsunamiConnectAfterUpsertHooks []TsunamiConnectHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TsunamiConnect) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tsunamiConnectAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TsunamiConnect) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tsunamiConnectBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TsunamiConnect) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tsunamiConnectAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TsunamiConnect) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tsunamiConnectBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TsunamiConnect) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tsunamiConnectAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TsunamiConnect) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tsunamiConnectBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TsunamiConnect) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tsunamiConnectAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TsunamiConnect) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tsunamiConnectBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TsunamiConnect) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tsunamiConnectAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTsunamiConnectHook registers your hook function for all future operations.
func AddTsunamiConnectHook(hookPoint boil.HookPoint, tsunamiConnectHook TsunamiConnectHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tsunamiConnectAfterSelectHooks = append(tsunamiConnectAfterSelectHooks, tsunamiConnectHook)
	case boil.BeforeInsertHook:
		tsunamiConnectBeforeInsertHooks = append(tsunamiConnectBeforeInsertHooks, tsunamiConnectHook)
	case boil.AfterInsertHook:
		tsunamiConnectAfterInsertHooks = append(tsunamiConnectAfterInsertHooks, tsunamiConnectHook)
	case boil.BeforeUpdateHook:
		tsunamiConnectBeforeUpdateHooks = append(tsunamiConnectBeforeUpdateHooks, tsunamiConnectHook)
	case boil.AfterUpdateHook:
		tsunamiConnectAfterUpdateHooks = append(tsunamiConnectAfterUpdateHooks, tsunamiConnectHook)
	case boil.BeforeDeleteHook:
		tsunamiConnectBeforeDeleteHooks = append(tsunamiConnectBeforeDeleteHooks, tsunamiConnectHook)
	case boil.AfterDeleteHook:
		tsunamiConnectAfterDeleteHooks = append(tsunamiConnectAfterDeleteHooks, tsunamiConnectHook)
	case boil.BeforeUpsertHook:
		tsunamiConnectBeforeUpsertHooks = append(tsunamiConnectBeforeUpsertHooks, tsunamiConnectHook)
	case boil.AfterUpsertHook:
		tsunamiConnectAfterUpsertHooks = append(tsunamiConnectAfterUpsertHooks, tsunamiConnectHook)
	}
}

// One returns a single tsunamiConnect record from the query.
func (q tsunamiConnectQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TsunamiConnect, error) {
	o := &TsunamiConnect{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for TsunamiConnects")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TsunamiConnect records from the query.
func (q tsunamiConnectQuery) All(ctx context.Context, exec boil.ContextExecutor) (TsunamiConnectSlice, error) {
	var o []*TsunamiConnect

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TsunamiConnect slice")
	}

	if len(tsunamiConnectAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TsunamiConnect records in the query.
func (q tsunamiConnectQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count TsunamiConnects rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tsunamiConnectQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if TsunamiConnects exists")
	}

	return count > 0, nil
}

// TsunamiConnects retrieves all the records using an executor.
func TsunamiConnects(mods ...qm.QueryMod) tsunamiConnectQuery {
	mods = append(mods, qm.From("`TsunamiConnects`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`TsunamiConnects`.*"})
	}

	return tsunamiConnectQuery{q}
}

// FindTsunamiConnect retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTsunamiConnect(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*TsunamiConnect, error) {
	tsunamiConnectObj := &TsunamiConnect{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `TsunamiConnects` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tsunamiConnectObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from TsunamiConnects")
	}

	if err = tsunamiConnectObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tsunamiConnectObj, err
	}

	return tsunamiConnectObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TsunamiConnect) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TsunamiConnects provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tsunamiConnectColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tsunamiConnectInsertCacheMut.RLock()
	cache, cached := tsunamiConnectInsertCache[key]
	tsunamiConnectInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tsunamiConnectAllColumns,
			tsunamiConnectColumnsWithDefault,
			tsunamiConnectColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tsunamiConnectType, tsunamiConnectMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tsunamiConnectType, tsunamiConnectMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `TsunamiConnects` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `TsunamiConnects` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `TsunamiConnects` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tsunamiConnectPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into TsunamiConnects")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tsunamiConnectMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TsunamiConnects")
	}

CacheNoHooks:
	if !cached {
		tsunamiConnectInsertCacheMut.Lock()
		tsunamiConnectInsertCache[key] = cache
		tsunamiConnectInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TsunamiConnect.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TsunamiConnect) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tsunamiConnectUpdateCacheMut.RLock()
	cache, cached := tsunamiConnectUpdateCache[key]
	tsunamiConnectUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tsunamiConnectAllColumns,
			tsunamiConnectPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update TsunamiConnects, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `TsunamiConnects` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tsunamiConnectPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tsunamiConnectType, tsunamiConnectMapping, append(wl, tsunamiConnectPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update TsunamiConnects row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for TsunamiConnects")
	}

	if !cached {
		tsunamiConnectUpdateCacheMut.Lock()
		tsunamiConnectUpdateCache[key] = cache
		tsunamiConnectUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tsunamiConnectQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for TsunamiConnects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for TsunamiConnects")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TsunamiConnectSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tsunamiConnectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `TsunamiConnects` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tsunamiConnectPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tsunamiConnect slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tsunamiConnect")
	}
	return rowsAff, nil
}

var mySQLTsunamiConnectUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TsunamiConnect) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TsunamiConnects provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tsunamiConnectColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTsunamiConnectUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	tsunamiConnectUpsertCacheMut.RLock()
	cache, cached := tsunamiConnectUpsertCache[key]
	tsunamiConnectUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tsunamiConnectAllColumns,
			tsunamiConnectColumnsWithDefault,
			tsunamiConnectColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tsunamiConnectAllColumns,
			tsunamiConnectPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert TsunamiConnects, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`TsunamiConnects`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `TsunamiConnects` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tsunamiConnectType, tsunamiConnectMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tsunamiConnectType, tsunamiConnectMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for TsunamiConnects")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tsunamiConnectMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tsunamiConnectType, tsunamiConnectMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for TsunamiConnects")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TsunamiConnects")
	}

CacheNoHooks:
	if !cached {
		tsunamiConnectUpsertCacheMut.Lock()
		tsunamiConnectUpsertCache[key] = cache
		tsunamiConnectUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TsunamiConnect record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TsunamiConnect) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TsunamiConnect provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tsunamiConnectPrimaryKeyMapping)
	sql := "DELETE FROM `TsunamiConnects` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from TsunamiConnects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for TsunamiConnects")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tsunamiConnectQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tsunamiConnectQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from TsunamiConnects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TsunamiConnects")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TsunamiConnectSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tsunamiConnectBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tsunamiConnectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `TsunamiConnects` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tsunamiConnectPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tsunamiConnect slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TsunamiConnects")
	}

	if len(tsunamiConnectAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TsunamiConnect) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTsunamiConnect(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TsunamiConnectSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TsunamiConnectSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tsunamiConnectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `TsunamiConnects`.* FROM `TsunamiConnects` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tsunamiConnectPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TsunamiConnectSlice")
	}

	*o = slice

	return nil
}

// TsunamiConnectExists checks if the TsunamiConnect row exists.
func TsunamiConnectExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `TsunamiConnects` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if TsunamiConnects exists")
	}

	return exists, nil
}

// Exists checks if the TsunamiConnect row exists.
func (o *TsunamiConnect) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TsunamiConnectExists(ctx, exec, o.ID)
}
