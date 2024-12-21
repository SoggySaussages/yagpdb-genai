// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// GenaiConfig is an object representing the database table.
type GenaiConfig struct {
	GuildID        int64  `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	Enabled        bool   `boil:"enabled" json:"enabled" toml:"enabled" yaml:"enabled"`
	Provider       int    `boil:"provider" json:"provider" toml:"provider" yaml:"provider"`
	Model          string `boil:"model" json:"model" toml:"model" yaml:"model"`
	Key            []byte `boil:"key" json:"key" toml:"key" yaml:"key"`
	BaseCMDEnabled bool   `boil:"base_cmd_enabled" json:"base_cmd_enabled" toml:"base_cmd_enabled" yaml:"base_cmd_enabled"`
	MaxTokens      int64  `boil:"max_tokens" json:"max_tokens" toml:"max_tokens" yaml:"max_tokens"`

	R *genaiConfigR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L genaiConfigL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GenaiConfigColumns = struct {
	GuildID        string
	Enabled        string
	Provider       string
	Model          string
	Key            string
	BaseCMDEnabled string
	MaxTokens      string
}{
	GuildID:        "guild_id",
	Enabled:        "enabled",
	Provider:       "provider",
	Model:          "model",
	Key:            "key",
	BaseCMDEnabled: "base_cmd_enabled",
	MaxTokens:      "max_tokens",
}

var GenaiConfigTableColumns = struct {
	GuildID        string
	Enabled        string
	Provider       string
	Model          string
	Key            string
	BaseCMDEnabled string
	MaxTokens      string
}{
	GuildID:        "genai_configs.guild_id",
	Enabled:        "genai_configs.enabled",
	Provider:       "genai_configs.provider",
	Model:          "genai_configs.model",
	Key:            "genai_configs.key",
	BaseCMDEnabled: "genai_configs.base_cmd_enabled",
	MaxTokens:      "genai_configs.max_tokens",
}

// Generated where

type whereHelper__byte struct{ field string }

func (w whereHelper__byte) EQ(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelper__byte) NEQ(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelper__byte) LT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelper__byte) LTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelper__byte) GT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelper__byte) GTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var GenaiConfigWhere = struct {
	GuildID        whereHelperint64
	Enabled        whereHelperbool
	Provider       whereHelperint
	Model          whereHelperstring
	Key            whereHelper__byte
	BaseCMDEnabled whereHelperbool
	MaxTokens      whereHelperint64
}{
	GuildID:        whereHelperint64{field: "\"genai_configs\".\"guild_id\""},
	Enabled:        whereHelperbool{field: "\"genai_configs\".\"enabled\""},
	Provider:       whereHelperint{field: "\"genai_configs\".\"provider\""},
	Model:          whereHelperstring{field: "\"genai_configs\".\"model\""},
	Key:            whereHelper__byte{field: "\"genai_configs\".\"key\""},
	BaseCMDEnabled: whereHelperbool{field: "\"genai_configs\".\"base_cmd_enabled\""},
	MaxTokens:      whereHelperint64{field: "\"genai_configs\".\"max_tokens\""},
}

// GenaiConfigRels is where relationship names are stored.
var GenaiConfigRels = struct {
}{}

// genaiConfigR is where relationships are stored.
type genaiConfigR struct {
}

// NewStruct creates a new relationship struct
func (*genaiConfigR) NewStruct() *genaiConfigR {
	return &genaiConfigR{}
}

// genaiConfigL is where Load methods for each relationship are stored.
type genaiConfigL struct{}

var (
	genaiConfigAllColumns            = []string{"guild_id", "enabled", "provider", "model", "key", "base_cmd_enabled", "max_tokens"}
	genaiConfigColumnsWithoutDefault = []string{"guild_id", "enabled", "provider", "model", "key", "base_cmd_enabled"}
	genaiConfigColumnsWithDefault    = []string{"max_tokens"}
	genaiConfigPrimaryKeyColumns     = []string{"guild_id"}
	genaiConfigGeneratedColumns      = []string{}
)

type (
	// GenaiConfigSlice is an alias for a slice of pointers to GenaiConfig.
	// This should almost always be used instead of []GenaiConfig.
	GenaiConfigSlice []*GenaiConfig

	genaiConfigQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	genaiConfigType                 = reflect.TypeOf(&GenaiConfig{})
	genaiConfigMapping              = queries.MakeStructMapping(genaiConfigType)
	genaiConfigPrimaryKeyMapping, _ = queries.BindMapping(genaiConfigType, genaiConfigMapping, genaiConfigPrimaryKeyColumns)
	genaiConfigInsertCacheMut       sync.RWMutex
	genaiConfigInsertCache          = make(map[string]insertCache)
	genaiConfigUpdateCacheMut       sync.RWMutex
	genaiConfigUpdateCache          = make(map[string]updateCache)
	genaiConfigUpsertCacheMut       sync.RWMutex
	genaiConfigUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single genaiConfig record from the query using the global executor.
func (q genaiConfigQuery) OneG(ctx context.Context) (*GenaiConfig, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single genaiConfig record from the query.
func (q genaiConfigQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GenaiConfig, error) {
	o := &GenaiConfig{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for genai_configs")
	}

	return o, nil
}

// AllG returns all GenaiConfig records from the query using the global executor.
func (q genaiConfigQuery) AllG(ctx context.Context) (GenaiConfigSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all GenaiConfig records from the query.
func (q genaiConfigQuery) All(ctx context.Context, exec boil.ContextExecutor) (GenaiConfigSlice, error) {
	var o []*GenaiConfig

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to GenaiConfig slice")
	}

	return o, nil
}

// CountG returns the count of all GenaiConfig records in the query using the global executor
func (q genaiConfigQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all GenaiConfig records in the query.
func (q genaiConfigQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count genai_configs rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q genaiConfigQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q genaiConfigQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if genai_configs exists")
	}

	return count > 0, nil
}

// GenaiConfigs retrieves all the records using an executor.
func GenaiConfigs(mods ...qm.QueryMod) genaiConfigQuery {
	mods = append(mods, qm.From("\"genai_configs\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"genai_configs\".*"})
	}

	return genaiConfigQuery{q}
}

// FindGenaiConfigG retrieves a single record by ID.
func FindGenaiConfigG(ctx context.Context, guildID int64, selectCols ...string) (*GenaiConfig, error) {
	return FindGenaiConfig(ctx, boil.GetContextDB(), guildID, selectCols...)
}

// FindGenaiConfig retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGenaiConfig(ctx context.Context, exec boil.ContextExecutor, guildID int64, selectCols ...string) (*GenaiConfig, error) {
	genaiConfigObj := &GenaiConfig{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"genai_configs\" where \"guild_id\"=$1", sel,
	)

	q := queries.Raw(query, guildID)

	err := q.Bind(ctx, exec, genaiConfigObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from genai_configs")
	}

	return genaiConfigObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *GenaiConfig) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GenaiConfig) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no genai_configs provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(genaiConfigColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	genaiConfigInsertCacheMut.RLock()
	cache, cached := genaiConfigInsertCache[key]
	genaiConfigInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			genaiConfigAllColumns,
			genaiConfigColumnsWithDefault,
			genaiConfigColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(genaiConfigType, genaiConfigMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(genaiConfigType, genaiConfigMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"genai_configs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"genai_configs\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into genai_configs")
	}

	if !cached {
		genaiConfigInsertCacheMut.Lock()
		genaiConfigInsertCache[key] = cache
		genaiConfigInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single GenaiConfig record using the global executor.
// See Update for more documentation.
func (o *GenaiConfig) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the GenaiConfig.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GenaiConfig) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	genaiConfigUpdateCacheMut.RLock()
	cache, cached := genaiConfigUpdateCache[key]
	genaiConfigUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			genaiConfigAllColumns,
			genaiConfigPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update genai_configs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"genai_configs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, genaiConfigPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(genaiConfigType, genaiConfigMapping, append(wl, genaiConfigPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update genai_configs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for genai_configs")
	}

	if !cached {
		genaiConfigUpdateCacheMut.Lock()
		genaiConfigUpdateCache[key] = cache
		genaiConfigUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q genaiConfigQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q genaiConfigQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for genai_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for genai_configs")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o GenaiConfigSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GenaiConfigSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genaiConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"genai_configs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, genaiConfigPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in genaiConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all genaiConfig")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *GenaiConfig) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GenaiConfig) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no genai_configs provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(genaiConfigColumnsWithDefault, o)

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

	genaiConfigUpsertCacheMut.RLock()
	cache, cached := genaiConfigUpsertCache[key]
	genaiConfigUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			genaiConfigAllColumns,
			genaiConfigColumnsWithDefault,
			genaiConfigColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			genaiConfigAllColumns,
			genaiConfigPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert genai_configs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(genaiConfigPrimaryKeyColumns))
			copy(conflict, genaiConfigPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"genai_configs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(genaiConfigType, genaiConfigMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(genaiConfigType, genaiConfigMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert genai_configs")
	}

	if !cached {
		genaiConfigUpsertCacheMut.Lock()
		genaiConfigUpsertCache[key] = cache
		genaiConfigUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single GenaiConfig record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *GenaiConfig) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single GenaiConfig record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GenaiConfig) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no GenaiConfig provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), genaiConfigPrimaryKeyMapping)
	sql := "DELETE FROM \"genai_configs\" WHERE \"guild_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from genai_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for genai_configs")
	}

	return rowsAff, nil
}

func (q genaiConfigQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q genaiConfigQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no genaiConfigQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from genai_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for genai_configs")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o GenaiConfigSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GenaiConfigSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genaiConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"genai_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, genaiConfigPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from genaiConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for genai_configs")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *GenaiConfig) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no GenaiConfig provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GenaiConfig) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGenaiConfig(ctx, exec, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GenaiConfigSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty GenaiConfigSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GenaiConfigSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GenaiConfigSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genaiConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"genai_configs\".* FROM \"genai_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, genaiConfigPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GenaiConfigSlice")
	}

	*o = slice

	return nil
}

// GenaiConfigExistsG checks if the GenaiConfig row exists.
func GenaiConfigExistsG(ctx context.Context, guildID int64) (bool, error) {
	return GenaiConfigExists(ctx, boil.GetContextDB(), guildID)
}

// GenaiConfigExists checks if the GenaiConfig row exists.
func GenaiConfigExists(ctx context.Context, exec boil.ContextExecutor, guildID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"genai_configs\" where \"guild_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, guildID)
	}
	row := exec.QueryRowContext(ctx, sql, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if genai_configs exists")
	}

	return exists, nil
}

// Exists checks if the GenaiConfig row exists.
func (o *GenaiConfig) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GenaiConfigExists(ctx, exec, o.GuildID)
}
