package json_filter_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	sq "github.com/Masterminds/squirrel"
	"github.com/haarika-sugarcrm/json_filter"
)

func TestEq(t *testing.T) {

	filter := []byte(`{"MovieName":{"$eq": "Godzilla"}}`)
	q := sq.Select("*").From("db")

	query, err := json_filter.ApplyFilter(q, filter)
	assert.NoError(t, err)

	sql, args, err := query.ToSql()
	expectedSql := "SELECT * FROM db WHERE MovieName = ?"
	assert.Equal(t, expectedSql, sql)

	assert.Equal(t, []interface{}{"Godzilla"}, args)
}

func TestNe(t *testing.T) {

	filter := []byte(`{"ActressName":{"$ne": "Johny"}}`)
	q := sq.Select("*").From("db")

	query, err := json_filter.ApplyFilter(q, filter)
	assert.NoError(t, err)

	sql, args, err := query.ToSql()
	expectedSql := "SELECT * FROM db WHERE ActressName <> ?"
	assert.Equal(t, expectedSql, sql)

	assert.Equal(t, []interface{}{"Johny"}, args)
}

func TestGt(t *testing.T) {

	filter := []byte(`{"ReleaseDate":{"$gt": "2018-10-18"}}`)
	q := sq.Select("*").From("db")

	query, err := json_filter.ApplyFilter(q, filter)
	assert.NoError(t, err)

	sql, args, err := query.ToSql()
	expectedSql := "SELECT * FROM db WHERE ReleaseDate > ?"
	assert.Equal(t, expectedSql, sql)

	assert.Equal(t, []interface{}{"2018-10-18"}, args)
}

func TestGte(t *testing.T) {

	filter := []byte(`{"ReleaseDate":{"$gte": "2018-10-18"}}`)
	q := sq.Select("*").From("db")

	query, err := json_filter.ApplyFilter(q, filter)
	assert.NoError(t, err)

	sql, args, err := query.ToSql()
	expectedSql := "SELECT * FROM db WHERE ReleaseDate >= ?"
	assert.Equal(t, expectedSql, sql)

	assert.Equal(t, []interface{}{"2018-10-18"}, args)
}

func TestLt(t *testing.T) {

	filter := []byte(`{"ReleaseDate":{"$lt": "2018-10-18"}}`)
	q := sq.Select("*").From("db")

	query, err := json_filter.ApplyFilter(q, filter)
	assert.NoError(t, err)

	sql, args, err := query.ToSql()
	expectedSql := "SELECT * FROM db WHERE ReleaseDate < ?"
	assert.Equal(t, expectedSql, sql)

	assert.Equal(t, []interface{}{"2018-10-18"}, args)
}

func TestLte(t *testing.T) {

	filter := []byte(`{"ReleaseDate":{"$lte": "2018-10-18"}}`)
	q := sq.Select("*").From("db")

	query, err := json_filter.ApplyFilter(q, filter)
	assert.NoError(t, err)

	sql, args, err := query.ToSql()
	expectedSql := "SELECT * FROM db WHERE ReleaseDate <= ?"
	assert.Equal(t, expectedSql, sql)

	assert.Equal(t, []interface{}{"2018-10-18"}, args)
}

func TestIsNull(t *testing.T) {

	filter := []byte(`{"ReleaseDate":{"$isnull": true}}`)
	q := sq.Select("*").From("db")

	query, err := json_filter.ApplyFilter(q, filter)
	assert.NoError(t, err)

	sql, _, err := query.ToSql()
	expectedSql := "SELECT * FROM db WHERE ReleaseDate IS NULL"
	assert.Equal(t, expectedSql, sql)
}

func TestIsNotNull(t *testing.T) {

	filter := []byte(`{"ReleaseDate":{"$isnotnull": true}}`)
	q := sq.Select("*").From("db")

	query, err := json_filter.ApplyFilter(q, filter)
	assert.NoError(t, err)

	sql, _, err := query.ToSql()
	expectedSql := "SELECT * FROM db WHERE ReleaseDate IS NOT NULL"
	assert.Equal(t, expectedSql, sql)
}

func TestIn(t *testing.T) {

	filter := []byte(`{"ActressName":{"$in": ["Jamie", "Johnny"]}}`)
	q := sq.Select("*").From("db")

	query, err := json_filter.ApplyFilter(q, filter)
	assert.NoError(t, err)

	sql, args, err := query.ToSql()
	expectedSql := "SELECT * FROM db WHERE ActressName IN (?,?)"
	assert.Equal(t, expectedSql, sql)

	assert.Equal(t, []interface{}{"Jamie", "Johnny"}, args)
}

func TestNotIn(t *testing.T) {

	filter := []byte(`{"ActressName":{"$notin": ["Jamie", "Johnny"]}}`)
	q := sq.Select("*").From("db")

	query, err := json_filter.ApplyFilter(q, filter)
	assert.NoError(t, err)

	sql, args, err := query.ToSql()
	expectedSql := "SELECT * FROM db WHERE ActressName NOT IN (?,?)"
	assert.Equal(t, expectedSql, sql)

	assert.Equal(t, []interface{}{"Jamie", "Johnny"}, args)
}