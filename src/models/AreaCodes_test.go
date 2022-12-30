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

func testAreaCodes(t *testing.T) {
	t.Parallel()

	query := AreaCodes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAreaCodesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
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

	count, err := AreaCodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAreaCodesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AreaCodes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AreaCodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAreaCodesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AreaCodeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AreaCodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAreaCodesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AreaCodeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AreaCode exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AreaCodeExists to return true, but got false.")
	}
}

func testAreaCodesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	areaCodeFound, err := FindAreaCode(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if areaCodeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAreaCodesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AreaCodes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAreaCodesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AreaCodes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAreaCodesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	areaCodeOne := &AreaCode{}
	areaCodeTwo := &AreaCode{}
	if err = randomize.Struct(seed, areaCodeOne, areaCodeDBTypes, false, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}
	if err = randomize.Struct(seed, areaCodeTwo, areaCodeDBTypes, false, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = areaCodeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = areaCodeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AreaCodes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAreaCodesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	areaCodeOne := &AreaCode{}
	areaCodeTwo := &AreaCode{}
	if err = randomize.Struct(seed, areaCodeOne, areaCodeDBTypes, false, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}
	if err = randomize.Struct(seed, areaCodeTwo, areaCodeDBTypes, false, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = areaCodeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = areaCodeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AreaCodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func areaCodeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AreaCode) error {
	*o = AreaCode{}
	return nil
}

func areaCodeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AreaCode) error {
	*o = AreaCode{}
	return nil
}

func areaCodeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AreaCode) error {
	*o = AreaCode{}
	return nil
}

func areaCodeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AreaCode) error {
	*o = AreaCode{}
	return nil
}

func areaCodeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AreaCode) error {
	*o = AreaCode{}
	return nil
}

func areaCodeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AreaCode) error {
	*o = AreaCode{}
	return nil
}

func areaCodeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AreaCode) error {
	*o = AreaCode{}
	return nil
}

func areaCodeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AreaCode) error {
	*o = AreaCode{}
	return nil
}

func areaCodeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AreaCode) error {
	*o = AreaCode{}
	return nil
}

func testAreaCodesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AreaCode{}
	o := &AreaCode{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, areaCodeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AreaCode object: %s", err)
	}

	AddAreaCodeHook(boil.BeforeInsertHook, areaCodeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	areaCodeBeforeInsertHooks = []AreaCodeHook{}

	AddAreaCodeHook(boil.AfterInsertHook, areaCodeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	areaCodeAfterInsertHooks = []AreaCodeHook{}

	AddAreaCodeHook(boil.AfterSelectHook, areaCodeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	areaCodeAfterSelectHooks = []AreaCodeHook{}

	AddAreaCodeHook(boil.BeforeUpdateHook, areaCodeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	areaCodeBeforeUpdateHooks = []AreaCodeHook{}

	AddAreaCodeHook(boil.AfterUpdateHook, areaCodeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	areaCodeAfterUpdateHooks = []AreaCodeHook{}

	AddAreaCodeHook(boil.BeforeDeleteHook, areaCodeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	areaCodeBeforeDeleteHooks = []AreaCodeHook{}

	AddAreaCodeHook(boil.AfterDeleteHook, areaCodeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	areaCodeAfterDeleteHooks = []AreaCodeHook{}

	AddAreaCodeHook(boil.BeforeUpsertHook, areaCodeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	areaCodeBeforeUpsertHooks = []AreaCodeHook{}

	AddAreaCodeHook(boil.AfterUpsertHook, areaCodeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	areaCodeAfterUpsertHooks = []AreaCodeHook{}
}

func testAreaCodesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AreaCodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAreaCodesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(areaCodeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AreaCodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAreaCodesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
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

func testAreaCodesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AreaCodeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAreaCodesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AreaCodes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	areaCodeDBTypes = map[string]string{`ID`: `int`, `Name`: `text`, `PrefectureID`: `int`, `Created`: `datetime`, `Updated`: `datetime`}
	_               = bytes.MinRead
)

func testAreaCodesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(areaCodePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(areaCodeAllColumns) == len(areaCodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AreaCodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAreaCodesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(areaCodeAllColumns) == len(areaCodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AreaCode{}
	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AreaCodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, areaCodeDBTypes, true, areaCodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(areaCodeAllColumns, areaCodePrimaryKeyColumns) {
		fields = areaCodeAllColumns
	} else {
		fields = strmangle.SetComplement(
			areaCodeAllColumns,
			areaCodePrimaryKeyColumns,
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

	slice := AreaCodeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAreaCodesUpsert(t *testing.T) {
	t.Parallel()

	if len(areaCodeAllColumns) == len(areaCodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLAreaCodeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AreaCode{}
	if err = randomize.Struct(seed, &o, areaCodeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AreaCode: %s", err)
	}

	count, err := AreaCodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, areaCodeDBTypes, false, areaCodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AreaCode struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AreaCode: %s", err)
	}

	count, err = AreaCodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}