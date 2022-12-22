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

func testEarthquakeInfos(t *testing.T) {
	t.Parallel()

	query := EarthquakeInfos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEarthquakeInfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
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

	count, err := EarthquakeInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEarthquakeInfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EarthquakeInfos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EarthquakeInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEarthquakeInfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EarthquakeInfoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EarthquakeInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEarthquakeInfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EarthquakeInfoExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if EarthquakeInfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EarthquakeInfoExists to return true, but got false.")
	}
}

func testEarthquakeInfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	earthquakeInfoFound, err := FindEarthquakeInfo(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if earthquakeInfoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEarthquakeInfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EarthquakeInfos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEarthquakeInfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EarthquakeInfos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEarthquakeInfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	earthquakeInfoOne := &EarthquakeInfo{}
	earthquakeInfoTwo := &EarthquakeInfo{}
	if err = randomize.Struct(seed, earthquakeInfoOne, earthquakeInfoDBTypes, false, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, earthquakeInfoTwo, earthquakeInfoDBTypes, false, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = earthquakeInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = earthquakeInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EarthquakeInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEarthquakeInfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	earthquakeInfoOne := &EarthquakeInfo{}
	earthquakeInfoTwo := &EarthquakeInfo{}
	if err = randomize.Struct(seed, earthquakeInfoOne, earthquakeInfoDBTypes, false, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, earthquakeInfoTwo, earthquakeInfoDBTypes, false, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = earthquakeInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = earthquakeInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarthquakeInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func earthquakeInfoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeInfo) error {
	*o = EarthquakeInfo{}
	return nil
}

func earthquakeInfoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeInfo) error {
	*o = EarthquakeInfo{}
	return nil
}

func earthquakeInfoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeInfo) error {
	*o = EarthquakeInfo{}
	return nil
}

func earthquakeInfoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeInfo) error {
	*o = EarthquakeInfo{}
	return nil
}

func earthquakeInfoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeInfo) error {
	*o = EarthquakeInfo{}
	return nil
}

func earthquakeInfoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeInfo) error {
	*o = EarthquakeInfo{}
	return nil
}

func earthquakeInfoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeInfo) error {
	*o = EarthquakeInfo{}
	return nil
}

func earthquakeInfoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeInfo) error {
	*o = EarthquakeInfo{}
	return nil
}

func earthquakeInfoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EarthquakeInfo) error {
	*o = EarthquakeInfo{}
	return nil
}

func testEarthquakeInfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &EarthquakeInfo{}
	o := &EarthquakeInfo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo object: %s", err)
	}

	AddEarthquakeInfoHook(boil.BeforeInsertHook, earthquakeInfoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	earthquakeInfoBeforeInsertHooks = []EarthquakeInfoHook{}

	AddEarthquakeInfoHook(boil.AfterInsertHook, earthquakeInfoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	earthquakeInfoAfterInsertHooks = []EarthquakeInfoHook{}

	AddEarthquakeInfoHook(boil.AfterSelectHook, earthquakeInfoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	earthquakeInfoAfterSelectHooks = []EarthquakeInfoHook{}

	AddEarthquakeInfoHook(boil.BeforeUpdateHook, earthquakeInfoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	earthquakeInfoBeforeUpdateHooks = []EarthquakeInfoHook{}

	AddEarthquakeInfoHook(boil.AfterUpdateHook, earthquakeInfoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	earthquakeInfoAfterUpdateHooks = []EarthquakeInfoHook{}

	AddEarthquakeInfoHook(boil.BeforeDeleteHook, earthquakeInfoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	earthquakeInfoBeforeDeleteHooks = []EarthquakeInfoHook{}

	AddEarthquakeInfoHook(boil.AfterDeleteHook, earthquakeInfoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	earthquakeInfoAfterDeleteHooks = []EarthquakeInfoHook{}

	AddEarthquakeInfoHook(boil.BeforeUpsertHook, earthquakeInfoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	earthquakeInfoBeforeUpsertHooks = []EarthquakeInfoHook{}

	AddEarthquakeInfoHook(boil.AfterUpsertHook, earthquakeInfoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	earthquakeInfoAfterUpsertHooks = []EarthquakeInfoHook{}
}

func testEarthquakeInfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarthquakeInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEarthquakeInfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(earthquakeInfoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EarthquakeInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEarthquakeInfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
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

func testEarthquakeInfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EarthquakeInfoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEarthquakeInfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EarthquakeInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	earthquakeInfoDBTypes = map[string]string{`ID`: `int`, `EventID`: `bigint`, `Lat`: `double`, `Lon`: `double`, `Depth`: `int`, `EpicenterName`: `text`, `MaxInt`: `varchar`, `Magnitude`: `text`, `Date`: `datetime`, `Created`: `datetime`, `Row`: `text`}
	_                     = bytes.MinRead
)

func testEarthquakeInfosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(earthquakeInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(earthquakeInfoAllColumns) == len(earthquakeInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarthquakeInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEarthquakeInfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(earthquakeInfoAllColumns) == len(earthquakeInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EarthquakeInfo{}
	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarthquakeInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, earthquakeInfoDBTypes, true, earthquakeInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(earthquakeInfoAllColumns, earthquakeInfoPrimaryKeyColumns) {
		fields = earthquakeInfoAllColumns
	} else {
		fields = strmangle.SetComplement(
			earthquakeInfoAllColumns,
			earthquakeInfoPrimaryKeyColumns,
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

	slice := EarthquakeInfoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEarthquakeInfosUpsert(t *testing.T) {
	t.Parallel()

	if len(earthquakeInfoAllColumns) == len(earthquakeInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLEarthquakeInfoUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EarthquakeInfo{}
	if err = randomize.Struct(seed, &o, earthquakeInfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EarthquakeInfo: %s", err)
	}

	count, err := EarthquakeInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, earthquakeInfoDBTypes, false, earthquakeInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EarthquakeInfo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EarthquakeInfo: %s", err)
	}

	count, err = EarthquakeInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
