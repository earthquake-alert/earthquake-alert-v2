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

// EarthquakeReport is an object representing the database table.
type EarthquakeReport struct {
	ID      uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	EventID uint64    `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	MaxInt  string    `boil:"max_int" json:"max_int" toml:"max_int" yaml:"max_int"`
	Date    time.Time `boil:"date" json:"date" toml:"date" yaml:"date"`
	Created time.Time `boil:"created" json:"created" toml:"created" yaml:"created"`
	Row     string    `boil:"row" json:"row" toml:"row" yaml:"row"`

	R *earthquakeReportR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L earthquakeReportL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EarthquakeReportColumns = struct {
	ID      string
	EventID string
	MaxInt  string
	Date    string
	Created string
	Row     string
}{
	ID:      "id",
	EventID: "event_id",
	MaxInt:  "max_int",
	Date:    "date",
	Created: "created",
	Row:     "row",
}

var EarthquakeReportTableColumns = struct {
	ID      string
	EventID string
	MaxInt  string
	Date    string
	Created string
	Row     string
}{
	ID:      "EarthquakeReports.id",
	EventID: "EarthquakeReports.event_id",
	MaxInt:  "EarthquakeReports.max_int",
	Date:    "EarthquakeReports.date",
	Created: "EarthquakeReports.created",
	Row:     "EarthquakeReports.row",
}

// Generated where

var EarthquakeReportWhere = struct {
	ID      whereHelperuint
	EventID whereHelperuint64
	MaxInt  whereHelperstring
	Date    whereHelpertime_Time
	Created whereHelpertime_Time
	Row     whereHelperstring
}{
	ID:      whereHelperuint{field: "`EarthquakeReports`.`id`"},
	EventID: whereHelperuint64{field: "`EarthquakeReports`.`event_id`"},
	MaxInt:  whereHelperstring{field: "`EarthquakeReports`.`max_int`"},
	Date:    whereHelpertime_Time{field: "`EarthquakeReports`.`date`"},
	Created: whereHelpertime_Time{field: "`EarthquakeReports`.`created`"},
	Row:     whereHelperstring{field: "`EarthquakeReports`.`row`"},
}

// EarthquakeReportRels is where relationship names are stored.
var EarthquakeReportRels = struct {
}{}

// earthquakeReportR is where relationships are stored.
type earthquakeReportR struct {
}

// NewStruct creates a new relationship struct
func (*earthquakeReportR) NewStruct() *earthquakeReportR {
	return &earthquakeReportR{}
}

// earthquakeReportL is where Load methods for each relationship are stored.
type earthquakeReportL struct{}

var (
	earthquakeReportAllColumns            = []string{"id", "event_id", "max_int", "date", "created", "row"}
	earthquakeReportColumnsWithoutDefault = []string{"event_id", "max_int", "date", "row"}
	earthquakeReportColumnsWithDefault    = []string{"id", "created"}
	earthquakeReportPrimaryKeyColumns     = []string{"id"}
	earthquakeReportGeneratedColumns      = []string{}
)

type (
	// EarthquakeReportSlice is an alias for a slice of pointers to EarthquakeReport.
	// This should almost always be used instead of []EarthquakeReport.
	EarthquakeReportSlice []*EarthquakeReport
	// EarthquakeReportHook is the signature for custom EarthquakeReport hook methods
	EarthquakeReportHook func(context.Context, boil.ContextExecutor, *EarthquakeReport) error

	earthquakeReportQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	earthquakeReportType                 = reflect.TypeOf(&EarthquakeReport{})
	earthquakeReportMapping              = queries.MakeStructMapping(earthquakeReportType)
	earthquakeReportPrimaryKeyMapping, _ = queries.BindMapping(earthquakeReportType, earthquakeReportMapping, earthquakeReportPrimaryKeyColumns)
	earthquakeReportInsertCacheMut       sync.RWMutex
	earthquakeReportInsertCache          = make(map[string]insertCache)
	earthquakeReportUpdateCacheMut       sync.RWMutex
	earthquakeReportUpdateCache          = make(map[string]updateCache)
	earthquakeReportUpsertCacheMut       sync.RWMutex
	earthquakeReportUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var earthquakeReportAfterSelectHooks []EarthquakeReportHook

var earthquakeReportBeforeInsertHooks []EarthquakeReportHook
var earthquakeReportAfterInsertHooks []EarthquakeReportHook

var earthquakeReportBeforeUpdateHooks []EarthquakeReportHook
var earthquakeReportAfterUpdateHooks []EarthquakeReportHook

var earthquakeReportBeforeDeleteHooks []EarthquakeReportHook
var earthquakeReportAfterDeleteHooks []EarthquakeReportHook

var earthquakeReportBeforeUpsertHooks []EarthquakeReportHook
var earthquakeReportAfterUpsertHooks []EarthquakeReportHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EarthquakeReport) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeReportAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EarthquakeReport) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeReportBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EarthquakeReport) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeReportAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EarthquakeReport) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeReportBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EarthquakeReport) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeReportAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EarthquakeReport) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeReportBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EarthquakeReport) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeReportAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EarthquakeReport) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeReportBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EarthquakeReport) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeReportAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEarthquakeReportHook registers your hook function for all future operations.
func AddEarthquakeReportHook(hookPoint boil.HookPoint, earthquakeReportHook EarthquakeReportHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		earthquakeReportAfterSelectHooks = append(earthquakeReportAfterSelectHooks, earthquakeReportHook)
	case boil.BeforeInsertHook:
		earthquakeReportBeforeInsertHooks = append(earthquakeReportBeforeInsertHooks, earthquakeReportHook)
	case boil.AfterInsertHook:
		earthquakeReportAfterInsertHooks = append(earthquakeReportAfterInsertHooks, earthquakeReportHook)
	case boil.BeforeUpdateHook:
		earthquakeReportBeforeUpdateHooks = append(earthquakeReportBeforeUpdateHooks, earthquakeReportHook)
	case boil.AfterUpdateHook:
		earthquakeReportAfterUpdateHooks = append(earthquakeReportAfterUpdateHooks, earthquakeReportHook)
	case boil.BeforeDeleteHook:
		earthquakeReportBeforeDeleteHooks = append(earthquakeReportBeforeDeleteHooks, earthquakeReportHook)
	case boil.AfterDeleteHook:
		earthquakeReportAfterDeleteHooks = append(earthquakeReportAfterDeleteHooks, earthquakeReportHook)
	case boil.BeforeUpsertHook:
		earthquakeReportBeforeUpsertHooks = append(earthquakeReportBeforeUpsertHooks, earthquakeReportHook)
	case boil.AfterUpsertHook:
		earthquakeReportAfterUpsertHooks = append(earthquakeReportAfterUpsertHooks, earthquakeReportHook)
	}
}

// One returns a single earthquakeReport record from the query.
func (q earthquakeReportQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EarthquakeReport, error) {
	o := &EarthquakeReport{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for EarthquakeReports")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EarthquakeReport records from the query.
func (q earthquakeReportQuery) All(ctx context.Context, exec boil.ContextExecutor) (EarthquakeReportSlice, error) {
	var o []*EarthquakeReport

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EarthquakeReport slice")
	}

	if len(earthquakeReportAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EarthquakeReport records in the query.
func (q earthquakeReportQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count EarthquakeReports rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q earthquakeReportQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if EarthquakeReports exists")
	}

	return count > 0, nil
}

// EarthquakeReports retrieves all the records using an executor.
func EarthquakeReports(mods ...qm.QueryMod) earthquakeReportQuery {
	mods = append(mods, qm.From("`EarthquakeReports`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`EarthquakeReports`.*"})
	}

	return earthquakeReportQuery{q}
}

// FindEarthquakeReport retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEarthquakeReport(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*EarthquakeReport, error) {
	earthquakeReportObj := &EarthquakeReport{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `EarthquakeReports` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, earthquakeReportObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from EarthquakeReports")
	}

	if err = earthquakeReportObj.doAfterSelectHooks(ctx, exec); err != nil {
		return earthquakeReportObj, err
	}

	return earthquakeReportObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EarthquakeReport) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no EarthquakeReports provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(earthquakeReportColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	earthquakeReportInsertCacheMut.RLock()
	cache, cached := earthquakeReportInsertCache[key]
	earthquakeReportInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			earthquakeReportAllColumns,
			earthquakeReportColumnsWithDefault,
			earthquakeReportColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(earthquakeReportType, earthquakeReportMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(earthquakeReportType, earthquakeReportMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `EarthquakeReports` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `EarthquakeReports` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `EarthquakeReports` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, earthquakeReportPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into EarthquakeReports")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == earthquakeReportMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for EarthquakeReports")
	}

CacheNoHooks:
	if !cached {
		earthquakeReportInsertCacheMut.Lock()
		earthquakeReportInsertCache[key] = cache
		earthquakeReportInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EarthquakeReport.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EarthquakeReport) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	earthquakeReportUpdateCacheMut.RLock()
	cache, cached := earthquakeReportUpdateCache[key]
	earthquakeReportUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			earthquakeReportAllColumns,
			earthquakeReportPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update EarthquakeReports, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `EarthquakeReports` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, earthquakeReportPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(earthquakeReportType, earthquakeReportMapping, append(wl, earthquakeReportPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update EarthquakeReports row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for EarthquakeReports")
	}

	if !cached {
		earthquakeReportUpdateCacheMut.Lock()
		earthquakeReportUpdateCache[key] = cache
		earthquakeReportUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q earthquakeReportQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for EarthquakeReports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for EarthquakeReports")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EarthquakeReportSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earthquakeReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `EarthquakeReports` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earthquakeReportPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in earthquakeReport slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all earthquakeReport")
	}
	return rowsAff, nil
}

var mySQLEarthquakeReportUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EarthquakeReport) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no EarthquakeReports provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(earthquakeReportColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEarthquakeReportUniqueColumns, o)

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

	earthquakeReportUpsertCacheMut.RLock()
	cache, cached := earthquakeReportUpsertCache[key]
	earthquakeReportUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			earthquakeReportAllColumns,
			earthquakeReportColumnsWithDefault,
			earthquakeReportColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			earthquakeReportAllColumns,
			earthquakeReportPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert EarthquakeReports, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`EarthquakeReports`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `EarthquakeReports` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(earthquakeReportType, earthquakeReportMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(earthquakeReportType, earthquakeReportMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for EarthquakeReports")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == earthquakeReportMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(earthquakeReportType, earthquakeReportMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for EarthquakeReports")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for EarthquakeReports")
	}

CacheNoHooks:
	if !cached {
		earthquakeReportUpsertCacheMut.Lock()
		earthquakeReportUpsertCache[key] = cache
		earthquakeReportUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EarthquakeReport record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EarthquakeReport) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EarthquakeReport provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), earthquakeReportPrimaryKeyMapping)
	sql := "DELETE FROM `EarthquakeReports` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from EarthquakeReports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for EarthquakeReports")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q earthquakeReportQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no earthquakeReportQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from EarthquakeReports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for EarthquakeReports")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EarthquakeReportSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(earthquakeReportBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earthquakeReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `EarthquakeReports` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earthquakeReportPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from earthquakeReport slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for EarthquakeReports")
	}

	if len(earthquakeReportAfterDeleteHooks) != 0 {
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
func (o *EarthquakeReport) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEarthquakeReport(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EarthquakeReportSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EarthquakeReportSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earthquakeReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `EarthquakeReports`.* FROM `EarthquakeReports` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earthquakeReportPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EarthquakeReportSlice")
	}

	*o = slice

	return nil
}

// EarthquakeReportExists checks if the EarthquakeReport row exists.
func EarthquakeReportExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `EarthquakeReports` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if EarthquakeReports exists")
	}

	return exists, nil
}

// Exists checks if the EarthquakeReport row exists.
func (o *EarthquakeReport) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EarthquakeReportExists(ctx, exec, o.ID)
}
