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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Game is an object representing the database table.
type Game struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Secret    []byte    `boil:"secret" json:"secret" toml:"secret" yaml:"secret"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *gameR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gameL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GameColumns = struct {
	ID        string
	Name      string
	Secret    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	Secret:    "secret",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

type whereHelper__byte struct{ field string }

func (w whereHelper__byte) EQ(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelper__byte) NEQ(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelper__byte) LT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelper__byte) LTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelper__byte) GT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelper__byte) GTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var GameWhere = struct {
	ID        whereHelperstring
	Name      whereHelperstring
	Secret    whereHelper__byte
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "\"games\".\"id\""},
	Name:      whereHelperstring{field: "\"games\".\"name\""},
	Secret:    whereHelper__byte{field: "\"games\".\"secret\""},
	CreatedAt: whereHelpernull_Time{field: "\"games\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"games\".\"updated_at\""},
}

// GameRels is where relationship names are stored.
var GameRels = struct {
	GameStats string
	UserGames string
}{
	GameStats: "GameStats",
	UserGames: "UserGames",
}

// gameR is where relationships are stored.
type gameR struct {
	GameStats GameStatSlice
	UserGames UserGameSlice
}

// NewStruct creates a new relationship struct
func (*gameR) NewStruct() *gameR {
	return &gameR{}
}

// gameL is where Load methods for each relationship are stored.
type gameL struct{}

var (
	gameAllColumns            = []string{"id", "name", "secret", "created_at", "updated_at"}
	gameColumnsWithoutDefault = []string{"id", "name", "secret", "created_at", "updated_at"}
	gameColumnsWithDefault    = []string{}
	gamePrimaryKeyColumns     = []string{"id"}
)

type (
	// GameSlice is an alias for a slice of pointers to Game.
	// This should generally be used opposed to []Game.
	GameSlice []*Game
	// GameHook is the signature for custom Game hook methods
	GameHook func(context.Context, boil.ContextExecutor, *Game) error

	gameQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gameType                 = reflect.TypeOf(&Game{})
	gameMapping              = queries.MakeStructMapping(gameType)
	gamePrimaryKeyMapping, _ = queries.BindMapping(gameType, gameMapping, gamePrimaryKeyColumns)
	gameInsertCacheMut       sync.RWMutex
	gameInsertCache          = make(map[string]insertCache)
	gameUpdateCacheMut       sync.RWMutex
	gameUpdateCache          = make(map[string]updateCache)
	gameUpsertCacheMut       sync.RWMutex
	gameUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var gameBeforeInsertHooks []GameHook
var gameBeforeUpdateHooks []GameHook
var gameBeforeDeleteHooks []GameHook
var gameBeforeUpsertHooks []GameHook

var gameAfterInsertHooks []GameHook
var gameAfterSelectHooks []GameHook
var gameAfterUpdateHooks []GameHook
var gameAfterDeleteHooks []GameHook
var gameAfterUpsertHooks []GameHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Game) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Game) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Game) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Game) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Game) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Game) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Game) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Game) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Game) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGameHook registers your hook function for all future operations.
func AddGameHook(hookPoint boil.HookPoint, gameHook GameHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		gameBeforeInsertHooks = append(gameBeforeInsertHooks, gameHook)
	case boil.BeforeUpdateHook:
		gameBeforeUpdateHooks = append(gameBeforeUpdateHooks, gameHook)
	case boil.BeforeDeleteHook:
		gameBeforeDeleteHooks = append(gameBeforeDeleteHooks, gameHook)
	case boil.BeforeUpsertHook:
		gameBeforeUpsertHooks = append(gameBeforeUpsertHooks, gameHook)
	case boil.AfterInsertHook:
		gameAfterInsertHooks = append(gameAfterInsertHooks, gameHook)
	case boil.AfterSelectHook:
		gameAfterSelectHooks = append(gameAfterSelectHooks, gameHook)
	case boil.AfterUpdateHook:
		gameAfterUpdateHooks = append(gameAfterUpdateHooks, gameHook)
	case boil.AfterDeleteHook:
		gameAfterDeleteHooks = append(gameAfterDeleteHooks, gameHook)
	case boil.AfterUpsertHook:
		gameAfterUpsertHooks = append(gameAfterUpsertHooks, gameHook)
	}
}

// One returns a single game record from the query.
func (q gameQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Game, error) {
	o := &Game{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datamodels_raw: failed to execute a one query for games")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Game records from the query.
func (q gameQuery) All(ctx context.Context, exec boil.ContextExecutor) (GameSlice, error) {
	var o []*Game

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "datamodels_raw: failed to assign all query results to Game slice")
	}

	if len(gameAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Game records in the query.
func (q gameQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to count games rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q gameQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "datamodels_raw: failed to check if games exists")
	}

	return count > 0, nil
}

// GameStats retrieves all the game_stat's GameStats with an executor.
func (o *Game) GameStats(mods ...qm.QueryMod) gameStatQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"game_stats\".\"game_id\"=?", o.ID),
	)

	query := GameStats(queryMods...)
	queries.SetFrom(query.Query, "\"game_stats\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"game_stats\".*"})
	}

	return query
}

// UserGames retrieves all the user_game's UserGames with an executor.
func (o *Game) UserGames(mods ...qm.QueryMod) userGameQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_game\".\"game_id\"=?", o.ID),
	)

	query := UserGames(queryMods...)
	queries.SetFrom(query.Query, "\"user_game\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"user_game\".*"})
	}

	return query
}

// LoadGameStats allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (gameL) LoadGameStats(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGame interface{}, mods queries.Applicator) error {
	var slice []*Game
	var object *Game

	if singular {
		object = maybeGame.(*Game)
	} else {
		slice = *maybeGame.(*[]*Game)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &gameR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gameR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`game_stats`), qm.WhereIn(`game_stats.game_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load game_stats")
	}

	var resultSlice []*GameStat
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice game_stats")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on game_stats")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for game_stats")
	}

	if len(gameStatAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.GameStats = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &gameStatR{}
			}
			foreign.R.Game = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GameID {
				local.R.GameStats = append(local.R.GameStats, foreign)
				if foreign.R == nil {
					foreign.R = &gameStatR{}
				}
				foreign.R.Game = local
				break
			}
		}
	}

	return nil
}

// LoadUserGames allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (gameL) LoadUserGames(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGame interface{}, mods queries.Applicator) error {
	var slice []*Game
	var object *Game

	if singular {
		object = maybeGame.(*Game)
	} else {
		slice = *maybeGame.(*[]*Game)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &gameR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gameR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`user_game`), qm.WhereIn(`user_game.game_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_game")
	}

	var resultSlice []*UserGame
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_game")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_game")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_game")
	}

	if len(userGameAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserGames = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userGameR{}
			}
			foreign.R.Game = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GameID {
				local.R.UserGames = append(local.R.UserGames, foreign)
				if foreign.R == nil {
					foreign.R = &userGameR{}
				}
				foreign.R.Game = local
				break
			}
		}
	}

	return nil
}

// AddGameStats adds the given related objects to the existing relationships
// of the game, optionally inserting them as new records.
// Appends related to o.R.GameStats.
// Sets related.R.Game appropriately.
func (o *Game) AddGameStats(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*GameStat) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GameID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"game_stats\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"game_id"}),
				strmangle.WhereClause("\"", "\"", 2, gameStatPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GameID = o.ID
		}
	}

	if o.R == nil {
		o.R = &gameR{
			GameStats: related,
		}
	} else {
		o.R.GameStats = append(o.R.GameStats, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &gameStatR{
				Game: o,
			}
		} else {
			rel.R.Game = o
		}
	}
	return nil
}

// AddUserGames adds the given related objects to the existing relationships
// of the game, optionally inserting them as new records.
// Appends related to o.R.UserGames.
// Sets related.R.Game appropriately.
func (o *Game) AddUserGames(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserGame) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GameID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_game\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"game_id"}),
				strmangle.WhereClause("\"", "\"", 2, userGamePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.UserID, rel.GameID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GameID = o.ID
		}
	}

	if o.R == nil {
		o.R = &gameR{
			UserGames: related,
		}
	} else {
		o.R.UserGames = append(o.R.UserGames, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userGameR{
				Game: o,
			}
		} else {
			rel.R.Game = o
		}
	}
	return nil
}

// Games retrieves all the records using an executor.
func Games(mods ...qm.QueryMod) gameQuery {
	mods = append(mods, qm.From("\"games\""))
	return gameQuery{NewQuery(mods...)}
}

// FindGame retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGame(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Game, error) {
	gameObj := &Game{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"games\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, gameObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datamodels_raw: unable to select from games")
	}

	return gameObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Game) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("datamodels_raw: no games provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gameColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gameInsertCacheMut.RLock()
	cache, cached := gameInsertCache[key]
	gameInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gameAllColumns,
			gameColumnsWithDefault,
			gameColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gameType, gameMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gameType, gameMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"games\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"games\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "datamodels_raw: unable to insert into games")
	}

	if !cached {
		gameInsertCacheMut.Lock()
		gameInsertCache[key] = cache
		gameInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Game.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Game) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	gameUpdateCacheMut.RLock()
	cache, cached := gameUpdateCache[key]
	gameUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gameAllColumns,
			gamePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("datamodels_raw: unable to update games, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"games\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, gamePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gameType, gameMapping, append(wl, gamePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "datamodels_raw: unable to update games row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by update for games")
	}

	if !cached {
		gameUpdateCacheMut.Lock()
		gameUpdateCache[key] = cache
		gameUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q gameQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to update all for games")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to retrieve rows affected for games")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GameSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"games\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, gamePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to update all in game slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to retrieve rows affected all in update all game")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Game) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("datamodels_raw: no games provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gameColumnsWithDefault, o)

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

	gameUpsertCacheMut.RLock()
	cache, cached := gameUpsertCache[key]
	gameUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			gameAllColumns,
			gameColumnsWithDefault,
			gameColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			gameAllColumns,
			gamePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("datamodels_raw: unable to upsert games, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(gamePrimaryKeyColumns))
			copy(conflict, gamePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"games\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(gameType, gameMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gameType, gameMapping, ret)
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
		return errors.Wrap(err, "datamodels_raw: unable to upsert games")
	}

	if !cached {
		gameUpsertCacheMut.Lock()
		gameUpsertCache[key] = cache
		gameUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Game record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Game) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("datamodels_raw: no Game provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gamePrimaryKeyMapping)
	sql := "DELETE FROM \"games\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to delete from games")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by delete for games")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q gameQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("datamodels_raw: no gameQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to delete all from games")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by deleteall for games")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GameSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(gameBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"games\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gamePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: unable to delete all from game slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodels_raw: failed to get rows affected by deleteall for games")
	}

	if len(gameAfterDeleteHooks) != 0 {
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
func (o *Game) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGame(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GameSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GameSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"games\".* FROM \"games\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gamePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "datamodels_raw: unable to reload all in GameSlice")
	}

	*o = slice

	return nil
}

// GameExists checks if the Game row exists.
func GameExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"games\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "datamodels_raw: unable to check if games exists")
	}

	return exists, nil
}
