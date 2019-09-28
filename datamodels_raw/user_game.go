// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package datamodels_raw

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// UserGame is an object representing the database table.
type UserGame struct {
	UserID string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GameID string `boil:"game_id" json:"game_id" toml:"game_id" yaml:"game_id"`

	R *userGameR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userGameL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserGameColumns = struct {
	UserID string
	GameID string
}{
	UserID: "user_id",
	GameID: "game_id",
}

// Generated where

var UserGameWhere = struct {
	UserID whereHelperstring
	GameID whereHelperstring
}{
	UserID: whereHelperstring{field: "\"user_game\".\"user_id\""},
	GameID: whereHelperstring{field: "\"user_game\".\"game_id\""},
}

// UserGameRels is where relationship names are stored.
var UserGameRels = struct {
}{}

// userGameR is where relationships are stored.
type userGameR struct {
}

// NewStruct creates a new relationship struct
func (*userGameR) NewStruct() *userGameR {
	return &userGameR{}
}

// userGameL is where Load methods for each relationship are stored.
type userGameL struct{}

var (
	userGameAllColumns            = []string{"user_id", "game_id"}
	userGameColumnsWithoutDefault = []string{"user_id", "game_id"}
	userGameColumnsWithDefault    = []string{}
	userGamePrimaryKeyColumns     = []string{"user_id", "game_id"}
)

type (
	// UserGameSlice is an alias for a slice of pointers to UserGame.
	// This should generally be used opposed to []UserGame.
	UserGameSlice []*UserGame
	// UserGameHook is the signature for custom UserGame hook methods
	UserGameHook func(context.Context, boil.ContextExecutor, *UserGame) error

	userGameQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userGameType                 = reflect.TypeOf(&UserGame{})
	userGameMapping              = queries.MakeStructMapping(userGameType)
	userGamePrimaryKeyMapping, _ = queries.BindMapping(userGameType, userGameMapping, userGamePrimaryKeyColumns)
	userGameInsertCacheMut       sync.RWMutex
	userGameInsertCache          = make(map[string]insertCache)
	userGameUpdateCacheMut       sync.RWMutex
	userGameUpdateCache          = make(map[string]updateCache)
	userGameUpsertCacheMut       sync.RWMutex
	userGameUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userGameBeforeInsertHooks []UserGameHook
var userGameBeforeUpdateHooks []UserGameHook
var userGameBeforeDeleteHooks []UserGameHook
var userGameBeforeUpsertHooks []UserGameHook

var userGameAfterInsertHooks []UserGameHook
var userGameAfterSelectHooks []UserGameHook
var userGameAfterUpdateHooks []UserGameHook
var userGameAfterDeleteHooks []UserGameHook
var userGameAfterUpsertHooks []UserGameHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserGame) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGameBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserGame) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGameBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserGame) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGameBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserGame) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGameBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserGame) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGameAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserGame) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGameAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserGame) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGameAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserGame) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGameAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserGame) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGameAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserGameHook registers your hook function for all future operations.
func AddUserGameHook(hookPoint boil.HookPoint, userGameHook UserGameHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userGameBeforeInsertHooks = append(userGameBeforeInsertHooks, userGameHook)
	case boil.BeforeUpdateHook:
		userGameBeforeUpdateHooks = append(userGameBeforeUpdateHooks, userGameHook)
	case boil.BeforeDeleteHook:
		userGameBeforeDeleteHooks = append(userGameBeforeDeleteHooks, userGameHook)
	case boil.BeforeUpsertHook:
		userGameBeforeUpsertHooks = append(userGameBeforeUpsertHooks, userGameHook)
	case boil.AfterInsertHook:
		userGameAfterInsertHooks = append(userGameAfterInsertHooks, userGameHook)
	case boil.AfterSelectHook:
		userGameAfterSelectHooks = append(userGameAfterSelectHooks, userGameHook)
	case boil.AfterUpdateHook:
		userGameAfterUpdateHooks = append(userGameAfterUpdateHooks, userGameHook)
	case boil.AfterDeleteHook:
		userGameAfterDeleteHooks = append(userGameAfterDeleteHooks, userGameHook)
	case boil.AfterUpsertHook:
		userGameAfterUpsertHooks = append(userGameAfterUpsertHooks, userGameHook)
	}
}

// One returns a single userGame record from the query.
func (q userGameQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserGame, error) {
	o := &UserGame{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datamodels_raw: failed to execute a one query for user_game")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserGame records from the query.
func (q userGameQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserGameSlice, error) {
	var o []*UserGame

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "datamodels_raw: failed to assign all query results to UserGame slice")
	}

	if len(userGameAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserGame records in the query.
func (q userGameQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to count user_game rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userGameQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "datamodels_raw: failed to check if user_game exists")
	}

	return count > 0, nil
}

// UserGames retrieves all the records using an executor.
func UserGames(mods ...qm.QueryMod) userGameQuery {
	mods = append(mods, qm.From("\"user_game\""))
	return userGameQuery{NewQuery(mods...)}
}

// FindUserGame retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserGame(ctx context.Context, exec boil.ContextExecutor, userID string, gameID string, selectCols ...string) (*UserGame, error) {
	userGameObj := &UserGame{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_game\" where \"user_id\"=$1 AND \"game_id\"=$2", sel,
	)

	q := queries.Raw(query, userID, gameID)

	err := q.Bind(ctx, exec, userGameObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datamodels_raw: unable to select from user_game")
	}

	return userGameObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserGame) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("datamodels_raw: no user_game provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userGameColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userGameInsertCacheMut.RLock()
	cache, cached := userGameInsertCache[key]
	userGameInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userGameAllColumns,
			userGameColumnsWithDefault,
			userGameColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userGameType, userGameMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userGameType, userGameMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_game\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_game\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "datamodels_raw: unable to insert into user_game")
	}

	if !cached {
		userGameInsertCacheMut.Lock()
		userGameInsertCache[key] = cache
		userGameInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserGame.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserGame) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userGameUpdateCacheMut.RLock()
	cache, cached := userGameUpdateCache[key]
	userGameUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userGameAllColumns,
			userGamePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("datamodels_raw: unable to update user_game, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_game\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userGamePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userGameType, userGameMapping, append(wl, userGamePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to update user_game row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by update for user_game")
	}

	if !cached {
		userGameUpdateCacheMut.Lock()
		userGameUpdateCache[key] = cache
		userGameUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userGameQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to update all for user_game")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to retrieve rows affected for user_game")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserGameSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("datamodels_raw: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userGamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_game\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userGamePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to update all in userGame slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to retrieve rows affected all in update all userGame")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserGame) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("datamodels_raw: no user_game provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userGameColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	userGameUpsertCacheMut.RLock()
	cache, cached := userGameUpsertCache[key]
	userGameUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userGameAllColumns,
			userGameColumnsWithDefault,
			userGameColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userGameAllColumns,
			userGamePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("datamodels_raw: unable to upsert user_game, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userGamePrimaryKeyColumns))
			copy(conflict, userGamePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_game\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userGameType, userGameMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userGameType, userGameMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "datamodels_raw: unable to upsert user_game")
	}

	if !cached {
		userGameUpsertCacheMut.Lock()
		userGameUpsertCache[key] = cache
		userGameUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserGame record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserGame) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("datamodels_raw: no UserGame provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userGamePrimaryKeyMapping)
	sql := "DELETE FROM \"user_game\" WHERE \"user_id\"=$1 AND \"game_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to delete from user_game")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by delete for user_game")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userGameQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("datamodels_raw: no userGameQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to delete all from user_game")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by deleteall for user_game")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserGameSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userGameBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userGamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_game\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userGamePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to delete all from userGame slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by deleteall for user_game")
	}

	if len(userGameAfterDeleteHooks) != 0 {
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
func (o *UserGame) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserGame(ctx, exec, o.UserID, o.GameID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserGameSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserGameSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userGamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_game\".* FROM \"user_game\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userGamePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "datamodels_raw: unable to reload all in UserGameSlice")
	}

	*o = slice

	return nil
}

// UserGameExists checks if the UserGame row exists.
func UserGameExists(ctx context.Context, exec boil.ContextExecutor, userID string, gameID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_game\" where \"user_id\"=$1 AND \"game_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userID, gameID)
	}

	row := exec.QueryRowContext(ctx, sql, userID, gameID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "datamodels_raw: unable to check if user_game exists")
	}

	return exists, nil
}
