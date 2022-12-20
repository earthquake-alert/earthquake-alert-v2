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

func testTwitterThreads(t *testing.T) {
	t.Parallel()

	query := TwitterThreads()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTwitterThreadsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
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

	count, err := TwitterThreads().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTwitterThreadsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TwitterThreads().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TwitterThreads().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTwitterThreadsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TwitterThreadSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TwitterThreads().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTwitterThreadsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TwitterThreadExists(ctx, tx, o.EventID)
	if err != nil {
		t.Errorf("Unable to check if TwitterThread exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TwitterThreadExists to return true, but got false.")
	}
}

func testTwitterThreadsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	twitterThreadFound, err := FindTwitterThread(ctx, tx, o.EventID)
	if err != nil {
		t.Error(err)
	}

	if twitterThreadFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTwitterThreadsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TwitterThreads().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTwitterThreadsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TwitterThreads().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTwitterThreadsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	twitterThreadOne := &TwitterThread{}
	twitterThreadTwo := &TwitterThread{}
	if err = randomize.Struct(seed, twitterThreadOne, twitterThreadDBTypes, false, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}
	if err = randomize.Struct(seed, twitterThreadTwo, twitterThreadDBTypes, false, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = twitterThreadOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = twitterThreadTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TwitterThreads().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTwitterThreadsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	twitterThreadOne := &TwitterThread{}
	twitterThreadTwo := &TwitterThread{}
	if err = randomize.Struct(seed, twitterThreadOne, twitterThreadDBTypes, false, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}
	if err = randomize.Struct(seed, twitterThreadTwo, twitterThreadDBTypes, false, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = twitterThreadOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = twitterThreadTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TwitterThreads().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func twitterThreadBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TwitterThread) error {
	*o = TwitterThread{}
	return nil
}

func twitterThreadAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TwitterThread) error {
	*o = TwitterThread{}
	return nil
}

func twitterThreadAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TwitterThread) error {
	*o = TwitterThread{}
	return nil
}

func twitterThreadBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TwitterThread) error {
	*o = TwitterThread{}
	return nil
}

func twitterThreadAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TwitterThread) error {
	*o = TwitterThread{}
	return nil
}

func twitterThreadBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TwitterThread) error {
	*o = TwitterThread{}
	return nil
}

func twitterThreadAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TwitterThread) error {
	*o = TwitterThread{}
	return nil
}

func twitterThreadBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TwitterThread) error {
	*o = TwitterThread{}
	return nil
}

func twitterThreadAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TwitterThread) error {
	*o = TwitterThread{}
	return nil
}

func testTwitterThreadsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TwitterThread{}
	o := &TwitterThread{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TwitterThread object: %s", err)
	}

	AddTwitterThreadHook(boil.BeforeInsertHook, twitterThreadBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	twitterThreadBeforeInsertHooks = []TwitterThreadHook{}

	AddTwitterThreadHook(boil.AfterInsertHook, twitterThreadAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	twitterThreadAfterInsertHooks = []TwitterThreadHook{}

	AddTwitterThreadHook(boil.AfterSelectHook, twitterThreadAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	twitterThreadAfterSelectHooks = []TwitterThreadHook{}

	AddTwitterThreadHook(boil.BeforeUpdateHook, twitterThreadBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	twitterThreadBeforeUpdateHooks = []TwitterThreadHook{}

	AddTwitterThreadHook(boil.AfterUpdateHook, twitterThreadAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	twitterThreadAfterUpdateHooks = []TwitterThreadHook{}

	AddTwitterThreadHook(boil.BeforeDeleteHook, twitterThreadBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	twitterThreadBeforeDeleteHooks = []TwitterThreadHook{}

	AddTwitterThreadHook(boil.AfterDeleteHook, twitterThreadAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	twitterThreadAfterDeleteHooks = []TwitterThreadHook{}

	AddTwitterThreadHook(boil.BeforeUpsertHook, twitterThreadBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	twitterThreadBeforeUpsertHooks = []TwitterThreadHook{}

	AddTwitterThreadHook(boil.AfterUpsertHook, twitterThreadAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	twitterThreadAfterUpsertHooks = []TwitterThreadHook{}
}

func testTwitterThreadsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TwitterThreads().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTwitterThreadsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(twitterThreadColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TwitterThreads().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTwitterThreadsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
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

func testTwitterThreadsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TwitterThreadSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTwitterThreadsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TwitterThreads().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	twitterThreadDBTypes = map[string]string{`EventID`: `bigint`, `TweetID`: `varchar`, `Created`: `datetime`}
	_                    = bytes.MinRead
)

func testTwitterThreadsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(twitterThreadPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(twitterThreadAllColumns) == len(twitterThreadPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TwitterThreads().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTwitterThreadsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(twitterThreadAllColumns) == len(twitterThreadPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TwitterThread{}
	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TwitterThreads().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, twitterThreadDBTypes, true, twitterThreadPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(twitterThreadAllColumns, twitterThreadPrimaryKeyColumns) {
		fields = twitterThreadAllColumns
	} else {
		fields = strmangle.SetComplement(
			twitterThreadAllColumns,
			twitterThreadPrimaryKeyColumns,
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

	slice := TwitterThreadSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTwitterThreadsUpsert(t *testing.T) {
	t.Parallel()

	if len(twitterThreadAllColumns) == len(twitterThreadPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLTwitterThreadUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TwitterThread{}
	if err = randomize.Struct(seed, &o, twitterThreadDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TwitterThread: %s", err)
	}

	count, err := TwitterThreads().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, twitterThreadDBTypes, false, twitterThreadPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TwitterThread struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TwitterThread: %s", err)
	}

	count, err = TwitterThreads().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
