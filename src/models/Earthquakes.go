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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Earthquake is an object representing the database table.
type Earthquake struct {
	EventID       uint64       `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	Lat           null.Float64 `boil:"lat" json:"lat,omitempty" toml:"lat" yaml:"lat,omitempty"`
	Lon           null.Float64 `boil:"lon" json:"lon,omitempty" toml:"lon" yaml:"lon,omitempty"`
	Depth         null.Int     `boil:"depth" json:"depth,omitempty" toml:"depth" yaml:"depth,omitempty"`
	EpicenterName null.String  `boil:"epicenter_name" json:"epicenter_name,omitempty" toml:"epicenter_name" yaml:"epicenter_name,omitempty"`
	MaxInt        string       `boil:"max_int" json:"max_int" toml:"max_int" yaml:"max_int"`
	Magnitude     null.String  `boil:"magnitude" json:"magnitude,omitempty" toml:"magnitude" yaml:"magnitude,omitempty"`
	Date          time.Time    `boil:"date" json:"date" toml:"date" yaml:"date"`
	Created       time.Time    `boil:"created" json:"created" toml:"created" yaml:"created"`
	Modified      time.Time    `boil:"modified" json:"modified" toml:"modified" yaml:"modified"`

	R *earthquakeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L earthquakeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EarthquakeColumns = struct {
	EventID       string
	Lat           string
	Lon           string
	Depth         string
	EpicenterName string
	MaxInt        string
	Magnitude     string
	Date          string
	Created       string
	Modified      string
}{
	EventID:       "event_id",
	Lat:           "lat",
	Lon:           "lon",
	Depth:         "depth",
	EpicenterName: "epicenter_name",
	MaxInt:        "max_int",
	Magnitude:     "magnitude",
	Date:          "date",
	Created:       "created",
	Modified:      "modified",
}

var EarthquakeTableColumns = struct {
	EventID       string
	Lat           string
	Lon           string
	Depth         string
	EpicenterName string
	MaxInt        string
	Magnitude     string
	Date          string
	Created       string
	Modified      string
}{
	EventID:       "Earthquakes.event_id",
	Lat:           "Earthquakes.lat",
	Lon:           "Earthquakes.lon",
	Depth:         "Earthquakes.depth",
	EpicenterName: "Earthquakes.epicenter_name",
	MaxInt:        "Earthquakes.max_int",
	Magnitude:     "Earthquakes.magnitude",
	Date:          "Earthquakes.date",
	Created:       "Earthquakes.created",
	Modified:      "Earthquakes.modified",
}

// Generated where

var EarthquakeWhere = struct {
	EventID       whereHelperuint64
	Lat           whereHelpernull_Float64
	Lon           whereHelpernull_Float64
	Depth         whereHelpernull_Int
	EpicenterName whereHelpernull_String
	MaxInt        whereHelperstring
	Magnitude     whereHelpernull_String
	Date          whereHelpertime_Time
	Created       whereHelpertime_Time
	Modified      whereHelpertime_Time
}{
	EventID:       whereHelperuint64{field: "`Earthquakes`.`event_id`"},
	Lat:           whereHelpernull_Float64{field: "`Earthquakes`.`lat`"},
	Lon:           whereHelpernull_Float64{field: "`Earthquakes`.`lon`"},
	Depth:         whereHelpernull_Int{field: "`Earthquakes`.`depth`"},
	EpicenterName: whereHelpernull_String{field: "`Earthquakes`.`epicenter_name`"},
	MaxInt:        whereHelperstring{field: "`Earthquakes`.`max_int`"},
	Magnitude:     whereHelpernull_String{field: "`Earthquakes`.`magnitude`"},
	Date:          whereHelpertime_Time{field: "`Earthquakes`.`date`"},
	Created:       whereHelpertime_Time{field: "`Earthquakes`.`created`"},
	Modified:      whereHelpertime_Time{field: "`Earthquakes`.`modified`"},
}

// EarthquakeRels is where relationship names are stored.
var EarthquakeRels = struct {
}{}

// earthquakeR is where relationships are stored.
type earthquakeR struct {
}

// NewStruct creates a new relationship struct
func (*earthquakeR) NewStruct() *earthquakeR {
	return &earthquakeR{}
}

// earthquakeL is where Load methods for each relationship are stored.
type earthquakeL struct{}

var (
	earthquakeAllColumns            = []string{"event_id", "lat", "lon", "depth", "epicenter_name", "max_int", "magnitude", "date", "created", "modified"}
	earthquakeColumnsWithoutDefault = []string{"event_id", "lat", "lon", "depth", "epicenter_name", "max_int", "magnitude", "date"}
	earthquakeColumnsWithDefault    = []string{"created", "modified"}
	earthquakePrimaryKeyColumns     = []string{"event_id"}
	earthquakeGeneratedColumns      = []string{}
)

type (
	// EarthquakeSlice is an alias for a slice of pointers to Earthquake.
	// This should almost always be used instead of []Earthquake.
	EarthquakeSlice []*Earthquake
	// EarthquakeHook is the signature for custom Earthquake hook methods
	EarthquakeHook func(context.Context, boil.ContextExecutor, *Earthquake) error

	earthquakeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	earthquakeType                 = reflect.TypeOf(&Earthquake{})
	earthquakeMapping              = queries.MakeStructMapping(earthquakeType)
	earthquakePrimaryKeyMapping, _ = queries.BindMapping(earthquakeType, earthquakeMapping, earthquakePrimaryKeyColumns)
	earthquakeInsertCacheMut       sync.RWMutex
	earthquakeInsertCache          = make(map[string]insertCache)
	earthquakeUpdateCacheMut       sync.RWMutex
	earthquakeUpdateCache          = make(map[string]updateCache)
	earthquakeUpsertCacheMut       sync.RWMutex
	earthquakeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var earthquakeAfterSelectHooks []EarthquakeHook

var earthquakeBeforeInsertHooks []EarthquakeHook
var earthquakeAfterInsertHooks []EarthquakeHook

var earthquakeBeforeUpdateHooks []EarthquakeHook
var earthquakeAfterUpdateHooks []EarthquakeHook

var earthquakeBeforeDeleteHooks []EarthquakeHook
var earthquakeAfterDeleteHooks []EarthquakeHook

var earthquakeBeforeUpsertHooks []EarthquakeHook
var earthquakeAfterUpsertHooks []EarthquakeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Earthquake) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Earthquake) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Earthquake) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Earthquake) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Earthquake) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Earthquake) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Earthquake) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Earthquake) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Earthquake) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earthquakeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEarthquakeHook registers your hook function for all future operations.
func AddEarthquakeHook(hookPoint boil.HookPoint, earthquakeHook EarthquakeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		earthquakeAfterSelectHooks = append(earthquakeAfterSelectHooks, earthquakeHook)
	case boil.BeforeInsertHook:
		earthquakeBeforeInsertHooks = append(earthquakeBeforeInsertHooks, earthquakeHook)
	case boil.AfterInsertHook:
		earthquakeAfterInsertHooks = append(earthquakeAfterInsertHooks, earthquakeHook)
	case boil.BeforeUpdateHook:
		earthquakeBeforeUpdateHooks = append(earthquakeBeforeUpdateHooks, earthquakeHook)
	case boil.AfterUpdateHook:
		earthquakeAfterUpdateHooks = append(earthquakeAfterUpdateHooks, earthquakeHook)
	case boil.BeforeDeleteHook:
		earthquakeBeforeDeleteHooks = append(earthquakeBeforeDeleteHooks, earthquakeHook)
	case boil.AfterDeleteHook:
		earthquakeAfterDeleteHooks = append(earthquakeAfterDeleteHooks, earthquakeHook)
	case boil.BeforeUpsertHook:
		earthquakeBeforeUpsertHooks = append(earthquakeBeforeUpsertHooks, earthquakeHook)
	case boil.AfterUpsertHook:
		earthquakeAfterUpsertHooks = append(earthquakeAfterUpsertHooks, earthquakeHook)
	}
}

// One returns a single earthquake record from the query.
func (q earthquakeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Earthquake, error) {
	o := &Earthquake{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Earthquakes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Earthquake records from the query.
func (q earthquakeQuery) All(ctx context.Context, exec boil.ContextExecutor) (EarthquakeSlice, error) {
	var o []*Earthquake

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Earthquake slice")
	}

	if len(earthquakeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Earthquake records in the query.
func (q earthquakeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Earthquakes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q earthquakeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Earthquakes exists")
	}

	return count > 0, nil
}

// Earthquakes retrieves all the records using an executor.
func Earthquakes(mods ...qm.QueryMod) earthquakeQuery {
	mods = append(mods, qm.From("`Earthquakes`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`Earthquakes`.*"})
	}

	return earthquakeQuery{q}
}

// FindEarthquake retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEarthquake(ctx context.Context, exec boil.ContextExecutor, eventID uint64, selectCols ...string) (*Earthquake, error) {
	earthquakeObj := &Earthquake{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `Earthquakes` where `event_id`=?", sel,
	)

	q := queries.Raw(query, eventID)

	err := q.Bind(ctx, exec, earthquakeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Earthquakes")
	}

	if err = earthquakeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return earthquakeObj, err
	}

	return earthquakeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Earthquake) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Earthquakes provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(earthquakeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	earthquakeInsertCacheMut.RLock()
	cache, cached := earthquakeInsertCache[key]
	earthquakeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			earthquakeAllColumns,
			earthquakeColumnsWithDefault,
			earthquakeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(earthquakeType, earthquakeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(earthquakeType, earthquakeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `Earthquakes` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `Earthquakes` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `Earthquakes` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, earthquakePrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into Earthquakes")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.EventID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for Earthquakes")
	}

CacheNoHooks:
	if !cached {
		earthquakeInsertCacheMut.Lock()
		earthquakeInsertCache[key] = cache
		earthquakeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Earthquake.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Earthquake) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	earthquakeUpdateCacheMut.RLock()
	cache, cached := earthquakeUpdateCache[key]
	earthquakeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			earthquakeAllColumns,
			earthquakePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Earthquakes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `Earthquakes` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, earthquakePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(earthquakeType, earthquakeMapping, append(wl, earthquakePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update Earthquakes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Earthquakes")
	}

	if !cached {
		earthquakeUpdateCacheMut.Lock()
		earthquakeUpdateCache[key] = cache
		earthquakeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q earthquakeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Earthquakes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Earthquakes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EarthquakeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earthquakePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `Earthquakes` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earthquakePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in earthquake slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all earthquake")
	}
	return rowsAff, nil
}

var mySQLEarthquakeUniqueColumns = []string{
	"event_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Earthquake) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Earthquakes provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(earthquakeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEarthquakeUniqueColumns, o)

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

	earthquakeUpsertCacheMut.RLock()
	cache, cached := earthquakeUpsertCache[key]
	earthquakeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			earthquakeAllColumns,
			earthquakeColumnsWithDefault,
			earthquakeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			earthquakeAllColumns,
			earthquakePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert Earthquakes, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`Earthquakes`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `Earthquakes` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(earthquakeType, earthquakeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(earthquakeType, earthquakeMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for Earthquakes")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(earthquakeType, earthquakeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for Earthquakes")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for Earthquakes")
	}

CacheNoHooks:
	if !cached {
		earthquakeUpsertCacheMut.Lock()
		earthquakeUpsertCache[key] = cache
		earthquakeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Earthquake record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Earthquake) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Earthquake provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), earthquakePrimaryKeyMapping)
	sql := "DELETE FROM `Earthquakes` WHERE `event_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Earthquakes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Earthquakes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q earthquakeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no earthquakeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Earthquakes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Earthquakes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EarthquakeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(earthquakeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earthquakePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `Earthquakes` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earthquakePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from earthquake slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Earthquakes")
	}

	if len(earthquakeAfterDeleteHooks) != 0 {
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
func (o *Earthquake) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEarthquake(ctx, exec, o.EventID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EarthquakeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EarthquakeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earthquakePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `Earthquakes`.* FROM `Earthquakes` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earthquakePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EarthquakeSlice")
	}

	*o = slice

	return nil
}

// EarthquakeExists checks if the Earthquake row exists.
func EarthquakeExists(ctx context.Context, exec boil.ContextExecutor, eventID uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `Earthquakes` where `event_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, eventID)
	}
	row := exec.QueryRowContext(ctx, sql, eventID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Earthquakes exists")
	}

	return exists, nil
}

// Exists checks if the Earthquake row exists.
func (o *Earthquake) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EarthquakeExists(ctx, exec, o.EventID)
}
