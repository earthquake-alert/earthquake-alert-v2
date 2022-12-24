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

// EarthquakeActivity is an object representing the database table.
type EarthquakeActivity struct {
	ID      uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	EventID uint64    `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	Date    time.Time `boil:"date" json:"date" toml:"date" yaml:"date"`
	Created time.Time `boil:"created" json:"created" toml:"created" yaml:"created"`
	Row     string    `boil:"row" json:"row" toml:"row" yaml:"row"`

	R *earthquakeActivityR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L earthquakeActivityL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EarthquakeActivityColumns = struct {
	ID      string
	EventID string
	Date    string
	Created string
	Row     string
}{
	ID:      "id",
	EventID: "event_id",
	Date:    "date",
	Created: "created",
	Row:     "row",
}

var EarthquakeActivityTableColumns = struct {
	ID      string
	EventID string
	Date    string
	Created string
	Row     string
}{
	ID:      "EarthquakeActivity.id",
	EventID: "EarthquakeActivity.event_id",
	Date:    "EarthquakeActivity.date",
	Created: "EarthquakeActivity.created",
	Row:     "EarthquakeActivity.row",
}

// Generated where

type whereHelperuint struct{ field string }

func (w whereHelperuint) EQ(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint) NEQ(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint) LT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint) LTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint) GT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint) GTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint) IN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint) NIN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperuint64 struct{ field string }

func (w whereHelperuint64) EQ(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint64) NEQ(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint64) LT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint64) LTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint64) GT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint64) GTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint64) IN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint64) NIN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var EarthquakeActivityWhere = struct {
	ID      whereHelperuint
	EventID whereHelperuint64
	Date    whereHelpertime_Time
	Created whereHelpertime_Time
	Row     whereHelperstring
}{
	ID:      whereHelperuint{field: "`EarthquakeActivity`.`id`"},
	EventID: whereHelperuint64{field: "`EarthquakeActivity`.`event_id`"},
	Date:    whereHelpertime_Time{field: "`EarthquakeActivity`.`date`"},
	Created: whereHelpertime_Time{field: "`EarthquakeActivity`.`created`"},
	Row:     whereHelperstring{field: "`EarthquakeActivity`.`row`"},
}

// EarthquakeActivityRels is where relationship names are stored.
var EarthquakeActivityRels = struct {
}{}

// earthquakeActivityR is where relationships are stored.
type earthquakeActivityR struct {
}

// NewStruct creates a new relationship struct
func (*earthquakeActivityR) NewStruct() *earthquakeActivityR {
	return &earthquakeActivityR{}
}

// earthquakeActivityL is where Load methods for each relationship are stored.
type earthquakeActivityL struct{}

var (
	earthquakeActivityAllColumns            = []string{"id", "event_id", "date", "created", "row"}
	earthquakeActivityColumnsWithoutDefault = []string{"event_id", "date", "row"}
	earthquakeActivityColumnsWithDefault    = []string{"id", "created"}
	earthquakeActivityPrimaryKeyColumns     = []string{"id"}
	earthquakeActivityGeneratedColumns      = []string{}
)

type (
	// EarthquakeActivitySlice is an alias for a slice of pointers to EarthquakeActivity.
	// This should almost always be used instead of []EarthquakeActivity.
	EarthquakeActivitySlice []*EarthquakeActivity
	// EarthquakeActivityHook is the signature for custom EarthquakeActivity hook methods
	EarthquakeActivityHook func(context.Context, boil.ContextExecutor, *EarthquakeActivity) error

	earthquakeActivityQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	earthquakeActivityType                 = reflect.TypeOf(&EarthquakeActivity{})
	earthquakeActivityMapping              = queries.MakeStructMapping(earthquakeActivityType)
	earthquakeActivityPrimaryKeyMapping, _ = queries.BindMapping(earthquakeActivityType, earthquakeActivityMapping, earthquakeActivityPrimaryKeyColumns)
	earthquakeActivityInsertCacheMut       sync.RWMutex
	earthquakeActivityInsertCache          = make(map[string]insertCache)
	earthquakeActivityUpdateCacheMut       sync.RWMutex
	earthquakeActivityUpdateCache          = make(map[string]updateCache)
	earthquakeActivityUpsertCacheMut       sync.RWMutex
	earthquakeActivityUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var earthquakeActivityAfterSelectHooks []EarthquakeActivityHook

var earthquakeActivityBeforeInsertHooks []EarthquakeActivityHook
var earthquakeActivityAfterInsertHooks []EarthquakeActivityHook

var earthquakeActivityBeforeUpdateHooks []EarthquakeActivityHook
var earthquakeActivityAfterUpdateHooks []EarthquakeActivityHook

var earthquakeActivityBeforeDeleteHooks []EarthquakeActivityHook
var earthquakeActivityAfterDeleteHooks []EarthquakeActivityHook

var earthquakeActivityBeforeUpsertHooks []EarthquakeActivityHook
var earthquakeActivityAfterUpsertHooks []EarthquakeActivityHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EarthquakeActivity) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeActivityAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EarthquakeActivity) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeActivityBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EarthquakeActivity) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeActivityAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EarthquakeActivity) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeActivityBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EarthquakeActivity) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeActivityAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EarthquakeActivity) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeActivityBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EarthquakeActivity) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeActivityAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EarthquakeActivity) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeActivityBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EarthquakeActivity) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeActivityAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEarthquakeActivityHook registers your hook function for all future operations.
func AddEarthquakeActivityHook(hookPoint boil.HookPoint, earthquakeActivityHook EarthquakeActivityHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		earthquakeActivityAfterSelectHooks = append(earthquakeActivityAfterSelectHooks, earthquakeActivityHook)
	case boil.BeforeInsertHook:
		earthquakeActivityBeforeInsertHooks = append(earthquakeActivityBeforeInsertHooks, earthquakeActivityHook)
	case boil.AfterInsertHook:
		earthquakeActivityAfterInsertHooks = append(earthquakeActivityAfterInsertHooks, earthquakeActivityHook)
	case boil.BeforeUpdateHook:
		earthquakeActivityBeforeUpdateHooks = append(earthquakeActivityBeforeUpdateHooks, earthquakeActivityHook)
	case boil.AfterUpdateHook:
		earthquakeActivityAfterUpdateHooks = append(earthquakeActivityAfterUpdateHooks, earthquakeActivityHook)
	case boil.BeforeDeleteHook:
		earthquakeActivityBeforeDeleteHooks = append(earthquakeActivityBeforeDeleteHooks, earthquakeActivityHook)
	case boil.AfterDeleteHook:
		earthquakeActivityAfterDeleteHooks = append(earthquakeActivityAfterDeleteHooks, earthquakeActivityHook)
	case boil.BeforeUpsertHook:
		earthquakeActivityBeforeUpsertHooks = append(earthquakeActivityBeforeUpsertHooks, earthquakeActivityHook)
	case boil.AfterUpsertHook:
		earthquakeActivityAfterUpsertHooks = append(earthquakeActivityAfterUpsertHooks, earthquakeActivityHook)
	}
}

// One returns a single earthquakeActivity record from the query.
func (q earthquakeActivityQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EarthquakeActivity, error) {
	o := &EarthquakeActivity{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for EarthquakeActivity")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EarthquakeActivity records from the query.
func (q earthquakeActivityQuery) All(ctx context.Context, exec boil.ContextExecutor) (EarthquakeActivitySlice, error) {
	var o []*EarthquakeActivity

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EarthquakeActivity slice")
	}

	if len(earthquakeActivityAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EarthquakeActivity records in the query.
func (q earthquakeActivityQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count EarthquakeActivity rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q earthquakeActivityQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if EarthquakeActivity exists")
	}

	return count > 0, nil
}

// EarthquakeActivities retrieves all the records using an executor.
func EarthquakeActivities(mods ...qm.QueryMod) earthquakeActivityQuery {
	mods = append(mods, qm.From("`EarthquakeActivity`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`EarthquakeActivity`.*"})
	}

	return earthquakeActivityQuery{q}
}

// FindEarthquakeActivity retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEarthquakeActivity(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*EarthquakeActivity, error) {
	earthquakeActivityObj := &EarthquakeActivity{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `EarthquakeActivity` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, earthquakeActivityObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from EarthquakeActivity")
	}

	if err = earthquakeActivityObj.doAfterSelectHooks(ctx, exec); err != nil {
		return earthquakeActivityObj, err
	}

	return earthquakeActivityObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EarthquakeActivity) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no EarthquakeActivity provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(earthquakeActivityColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	earthquakeActivityInsertCacheMut.RLock()
	cache, cached := earthquakeActivityInsertCache[key]
	earthquakeActivityInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			earthquakeActivityAllColumns,
			earthquakeActivityColumnsWithDefault,
			earthquakeActivityColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(earthquakeActivityType, earthquakeActivityMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(earthquakeActivityType, earthquakeActivityMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `EarthquakeActivity` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `EarthquakeActivity` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `EarthquakeActivity` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, earthquakeActivityPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into EarthquakeActivity")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == earthquakeActivityMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for EarthquakeActivity")
	}

CacheNoHooks:
	if !cached {
		earthquakeActivityInsertCacheMut.Lock()
		earthquakeActivityInsertCache[key] = cache
		earthquakeActivityInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EarthquakeActivity.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EarthquakeActivity) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	earthquakeActivityUpdateCacheMut.RLock()
	cache, cached := earthquakeActivityUpdateCache[key]
	earthquakeActivityUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			earthquakeActivityAllColumns,
			earthquakeActivityPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update EarthquakeActivity, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `EarthquakeActivity` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, earthquakeActivityPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(earthquakeActivityType, earthquakeActivityMapping, append(wl, earthquakeActivityPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update EarthquakeActivity row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for EarthquakeActivity")
	}

	if !cached {
		earthquakeActivityUpdateCacheMut.Lock()
		earthquakeActivityUpdateCache[key] = cache
		earthquakeActivityUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q earthquakeActivityQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for EarthquakeActivity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for EarthquakeActivity")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EarthquakeActivitySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earthquakeActivityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `EarthquakeActivity` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earthquakeActivityPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in earthquakeActivity slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all earthquakeActivity")
	}
	return rowsAff, nil
}

var mySQLEarthquakeActivityUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EarthquakeActivity) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no EarthquakeActivity provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(earthquakeActivityColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEarthquakeActivityUniqueColumns, o)

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

	earthquakeActivityUpsertCacheMut.RLock()
	cache, cached := earthquakeActivityUpsertCache[key]
	earthquakeActivityUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			earthquakeActivityAllColumns,
			earthquakeActivityColumnsWithDefault,
			earthquakeActivityColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			earthquakeActivityAllColumns,
			earthquakeActivityPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert EarthquakeActivity, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`EarthquakeActivity`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `EarthquakeActivity` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(earthquakeActivityType, earthquakeActivityMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(earthquakeActivityType, earthquakeActivityMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for EarthquakeActivity")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == earthquakeActivityMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(earthquakeActivityType, earthquakeActivityMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for EarthquakeActivity")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for EarthquakeActivity")
	}

CacheNoHooks:
	if !cached {
		earthquakeActivityUpsertCacheMut.Lock()
		earthquakeActivityUpsertCache[key] = cache
		earthquakeActivityUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EarthquakeActivity record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EarthquakeActivity) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EarthquakeActivity provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), earthquakeActivityPrimaryKeyMapping)
	sql := "DELETE FROM `EarthquakeActivity` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from EarthquakeActivity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for EarthquakeActivity")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q earthquakeActivityQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no earthquakeActivityQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from EarthquakeActivity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for EarthquakeActivity")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EarthquakeActivitySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(earthquakeActivityBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earthquakeActivityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `EarthquakeActivity` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earthquakeActivityPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from earthquakeActivity slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for EarthquakeActivity")
	}

	if len(earthquakeActivityAfterDeleteHooks) != 0 {
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
func (o *EarthquakeActivity) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEarthquakeActivity(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EarthquakeActivitySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EarthquakeActivitySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earthquakeActivityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `EarthquakeActivity`.* FROM `EarthquakeActivity` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earthquakeActivityPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EarthquakeActivitySlice")
	}

	*o = slice

	return nil
}

// EarthquakeActivityExists checks if the EarthquakeActivity row exists.
func EarthquakeActivityExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `EarthquakeActivity` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if EarthquakeActivity exists")
	}

	return exists, nil
}

// Exists checks if the EarthquakeActivity row exists.
func (o *EarthquakeActivity) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EarthquakeActivityExists(ctx, exec, o.ID)
}
