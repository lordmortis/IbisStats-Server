// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package datamodels_raw

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testGameStats(t *testing.T) {
	t.Parallel()

	query := GameStats()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGameStatsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
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

	count, err := GameStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGameStatsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := GameStats().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GameStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGameStatsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GameStatSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GameStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGameStatsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GameStatExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if GameStat exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GameStatExists to return true, but got false.")
	}
}

func testGameStatsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	gameStatFound, err := FindGameStat(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if gameStatFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGameStatsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = GameStats().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGameStatsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := GameStats().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGameStatsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gameStatOne := &GameStat{}
	gameStatTwo := &GameStat{}
	if err = randomize.Struct(seed, gameStatOne, gameStatDBTypes, false, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}
	if err = randomize.Struct(seed, gameStatTwo, gameStatDBTypes, false, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = gameStatOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = gameStatTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GameStats().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGameStatsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	gameStatOne := &GameStat{}
	gameStatTwo := &GameStat{}
	if err = randomize.Struct(seed, gameStatOne, gameStatDBTypes, false, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}
	if err = randomize.Struct(seed, gameStatTwo, gameStatDBTypes, false, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = gameStatOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = gameStatTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GameStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func gameStatBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *GameStat) error {
	*o = GameStat{}
	return nil
}

func gameStatAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *GameStat) error {
	*o = GameStat{}
	return nil
}

func gameStatAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *GameStat) error {
	*o = GameStat{}
	return nil
}

func gameStatBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *GameStat) error {
	*o = GameStat{}
	return nil
}

func gameStatAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *GameStat) error {
	*o = GameStat{}
	return nil
}

func gameStatBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *GameStat) error {
	*o = GameStat{}
	return nil
}

func gameStatAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *GameStat) error {
	*o = GameStat{}
	return nil
}

func gameStatBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *GameStat) error {
	*o = GameStat{}
	return nil
}

func gameStatAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *GameStat) error {
	*o = GameStat{}
	return nil
}

func testGameStatsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &GameStat{}
	o := &GameStat{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, gameStatDBTypes, false); err != nil {
		t.Errorf("Unable to randomize GameStat object: %s", err)
	}

	AddGameStatHook(boil.BeforeInsertHook, gameStatBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	gameStatBeforeInsertHooks = []GameStatHook{}

	AddGameStatHook(boil.AfterInsertHook, gameStatAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	gameStatAfterInsertHooks = []GameStatHook{}

	AddGameStatHook(boil.AfterSelectHook, gameStatAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	gameStatAfterSelectHooks = []GameStatHook{}

	AddGameStatHook(boil.BeforeUpdateHook, gameStatBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	gameStatBeforeUpdateHooks = []GameStatHook{}

	AddGameStatHook(boil.AfterUpdateHook, gameStatAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	gameStatAfterUpdateHooks = []GameStatHook{}

	AddGameStatHook(boil.BeforeDeleteHook, gameStatBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	gameStatBeforeDeleteHooks = []GameStatHook{}

	AddGameStatHook(boil.AfterDeleteHook, gameStatAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	gameStatAfterDeleteHooks = []GameStatHook{}

	AddGameStatHook(boil.BeforeUpsertHook, gameStatBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	gameStatBeforeUpsertHooks = []GameStatHook{}

	AddGameStatHook(boil.AfterUpsertHook, gameStatAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	gameStatAfterUpsertHooks = []GameStatHook{}
}

func testGameStatsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GameStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGameStatsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(gameStatColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := GameStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGameStatToOneGameUsingGame(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local GameStat
	var foreign Game

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, gameStatDBTypes, false, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, gameDBTypes, false, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.GameID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Game().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := GameStatSlice{&local}
	if err = local.L.LoadGame(ctx, tx, false, (*[]*GameStat)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Game == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Game = nil
	if err = local.L.LoadGame(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Game == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testGameStatToOneSetOpGameUsingGame(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a GameStat
	var b, c Game

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, gameStatDBTypes, false, strmangle.SetComplement(gameStatPrimaryKeyColumns, gameStatColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, gameDBTypes, false, strmangle.SetComplement(gamePrimaryKeyColumns, gameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, gameDBTypes, false, strmangle.SetComplement(gamePrimaryKeyColumns, gameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Game{&b, &c} {
		err = a.SetGame(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Game != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.GameStats[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GameID != x.ID {
			t.Error("foreign key was wrong value", a.GameID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GameID))
		reflect.Indirect(reflect.ValueOf(&a.GameID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GameID != x.ID {
			t.Error("foreign key was wrong value", a.GameID, x.ID)
		}
	}
}

func testGameStatsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
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

func testGameStatsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GameStatSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGameStatsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GameStats().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	gameStatDBTypes = map[string]string{`ID`: `uuid`, `GameID`: `uuid`, `Name`: `text`, `Type`: `enum.game_stat_type('score','time')`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_               = bytes.MinRead
)

func testGameStatsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(gameStatPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(gameStatAllColumns) == len(gameStatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GameStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGameStatsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(gameStatAllColumns) == len(gameStatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GameStat{}
	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GameStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, gameStatDBTypes, true, gameStatPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(gameStatAllColumns, gameStatPrimaryKeyColumns) {
		fields = gameStatAllColumns
	} else {
		fields = strmangle.SetComplement(
			gameStatAllColumns,
			gameStatPrimaryKeyColumns,
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

	slice := GameStatSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGameStatsUpsert(t *testing.T) {
	t.Parallel()

	if len(gameStatAllColumns) == len(gameStatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := GameStat{}
	if err = randomize.Struct(seed, &o, gameStatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert GameStat: %s", err)
	}

	count, err := GameStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, gameStatDBTypes, false, gameStatPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GameStat struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert GameStat: %s", err)
	}

	count, err = GameStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
