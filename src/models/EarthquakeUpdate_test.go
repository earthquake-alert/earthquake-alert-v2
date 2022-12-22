// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testEarthquakeUpdates(t *testing.T) {
	t.Parallel()

	query := EarthquakeUpdates()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEarthquakeUpdatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EarthquakeUpdates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEarthquakeUpdatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EarthquakeUpdates().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EarthquakeUpdates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEarthquakeUpdatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EarthquakeUpdateSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EarthquakeUpdates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEarthquakeUpdatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EarthquakeUpdateExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if EarthquakeUpdate exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EarthquakeUpdateExists to return true, but got false.")
	}
}

func testEarthquakeUpdatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	earthquakeUpdateFound, err := FindEarthquakeUpdate(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if earthquakeUpdateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEarthquakeUpdatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EarthquakeUpdates().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEarthquakeUpdatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EarthquakeUpdates().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEarthquakeUpdatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	earthquakeUpdateOne := &EarthquakeUpdate{}
	earthquakeUpdateTwo := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, earthquakeUpdateOne, earthquakeUpdateDBTypes, false, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}
	if err = randomize.Struct(seed, earthquakeUpdateTwo, earthquakeUpdateDBTypes, false, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = earthquakeUpdateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = earthquakeUpdateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EarthquakeUpdates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEarthquakeUpdatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	earthquakeUpdateOne := &EarthquakeUpdate{}
	earthquakeUpdateTwo := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, earthquakeUpdateOne, earthquakeUpdateDBTypes, false, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}
	if err = randomize.Struct(seed, earthquakeUpdateTwo, earthquakeUpdateDBTypes, false, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = earthquakeUpdateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = earthquakeUpdateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarthquakeUpdates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func earthquakeUpdateBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeUpdate) error {
	*o = EarthquakeUpdate{}
	return nil
}

func earthquakeUpdateAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeUpdate) error {
	*o = EarthquakeUpdate{}
	return nil
}

func earthquakeUpdateAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeUpdate) error {
	*o = EarthquakeUpdate{}
	return nil
}

func earthquakeUpdateBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeUpdate) error {
	*o = EarthquakeUpdate{}
	return nil
}

func earthquakeUpdateAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeUpdate) error {
	*o = EarthquakeUpdate{}
	return nil
}

func earthquakeUpdateBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeUpdate) error {
	*o = EarthquakeUpdate{}
	return nil
}

func earthquakeUpdateAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeUpdate) error {
	*o = EarthquakeUpdate{}
	return nil
}

func earthquakeUpdateBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeUpdate) error {
	*o = EarthquakeUpdate{}
	return nil
}

func earthquakeUpdateAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeUpdate) error {
	*o = EarthquakeUpdate{}
	return nil
}

func testEarthquakeUpdatesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &EarthquakeUpdate{}
	o := &EarthquakeUpdate{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate object: %s", err)
	}

	AddEarthquakeUpdateHook(boil.BeforeInsertHook, earthquakeUpdateBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	earthquakeUpdateBeforeInsertHooks = []EarthquakeUpdateHook{}

	AddEarthquakeUpdateHook(boil.AfterInsertHook, earthquakeUpdateAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	earthquakeUpdateAfterInsertHooks = []EarthquakeUpdateHook{}

	AddEarthquakeUpdateHook(boil.AfterSelectHook, earthquakeUpdateAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	earthquakeUpdateAfterSelectHooks = []EarthquakeUpdateHook{}

	AddEarthquakeUpdateHook(boil.BeforeUpdateHook, earthquakeUpdateBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	earthquakeUpdateBeforeUpdateHooks = []EarthquakeUpdateHook{}

	AddEarthquakeUpdateHook(boil.AfterUpdateHook, earthquakeUpdateAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	earthquakeUpdateAfterUpdateHooks = []EarthquakeUpdateHook{}

	AddEarthquakeUpdateHook(boil.BeforeDeleteHook, earthquakeUpdateBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	earthquakeUpdateBeforeDeleteHooks = []EarthquakeUpdateHook{}

	AddEarthquakeUpdateHook(boil.AfterDeleteHook, earthquakeUpdateAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	earthquakeUpdateAfterDeleteHooks = []EarthquakeUpdateHook{}

	AddEarthquakeUpdateHook(boil.BeforeUpsertHook, earthquakeUpdateBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	earthquakeUpdateBeforeUpsertHooks = []EarthquakeUpdateHook{}

	AddEarthquakeUpdateHook(boil.AfterUpsertHook, earthquakeUpdateAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	earthquakeUpdateAfterUpsertHooks = []EarthquakeUpdateHook{}
}

func testEarthquakeUpdatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarthquakeUpdates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEarthquakeUpdatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(earthquakeUpdateColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EarthquakeUpdates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEarthquakeUpdatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEarthquakeUpdatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EarthquakeUpdateSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEarthquakeUpdatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EarthquakeUpdates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	earthquakeUpdateDBTypes = map[string]string{`ID`: `int`, `EventID`: `bigint`, `Lat`: `double`, `Lon`: `double`, `Depth`: `int`, `Magnitude`: `text`, `Date`: `datetime`, `Created`: `datetime`, `Row`: `text`}
	_                       = bytes.MinRead
)

func testEarthquakeUpdatesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(earthquakeUpdatePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(earthquakeUpdateAllColumns) == len(earthquakeUpdatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarthquakeUpdates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEarthquakeUpdatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(earthquakeUpdateAllColumns) == len(earthquakeUpdatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeUpdate{}
	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarthquakeUpdates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, earthquakeUpdateDBTypes, true, earthquakeUpdatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(earthquakeUpdateAllColumns, earthquakeUpdatePrimaryKeyColumns) {
		fields = earthquakeUpdateAllColumns
	} else {
		fields = strmangle.SetComplement(
			earthquakeUpdateAllColumns,
			earthquakeUpdatePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := EarthquakeUpdateSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEarthquakeUpdatesUpsert(t *testing.T) {
	t.Parallel()

	if len(earthquakeUpdateAllColumns) == len(earthquakeUpdatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLEarthquakeUpdateUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EarthquakeUpdate{}
	if err = randomize.Struct(seed, &o, earthquakeUpdateDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EarthquakeUpdate: %s", err)
	}

	count, err := EarthquakeUpdates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, earthquakeUpdateDBTypes, false, earthquakeUpdatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EarthquakeUpdate struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EarthquakeUpdate: %s", err)
	}

	count, err = EarthquakeUpdates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
