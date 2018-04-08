package specs

import (
	"testing"
	"time"

	"github.com/Fs02/grimoire"
	"github.com/Fs02/grimoire/adapter/sqlutil"
	"github.com/Fs02/grimoire/changeset"
	"github.com/stretchr/testify/assert"
)

// Insert tests insert specifications.
func Insert(t *testing.T, repo grimoire.Repo) {
	user := User{}
	assert.Nil(t, repo.From(users).Put(&user))

	tests := []struct {
		query  grimoire.Query
		record interface{}
		params map[string]interface{}
	}{
		{repo.From(users), &User{}, map[string]interface{}{}},
		{repo.From(users), &User{}, map[string]interface{}{"name": "insert", "age": 100}},
		{repo.From(users), &User{}, map[string]interface{}{"name": "insert", "age": 100, "note": "note"}},
		{repo.From(users), &User{}, map[string]interface{}{"note": "note"}},
		{repo.From(addresses), &Address{}, map[string]interface{}{}},
		{repo.From(addresses), &Address{}, map[string]interface{}{"address": "address"}},
		{repo.From(addresses), &Address{}, map[string]interface{}{"user_id": user.ID}},
		{repo.From(addresses), &Address{}, map[string]interface{}{"address": "address", "user_id": user.ID}},
	}

	for _, test := range tests {
		ch := changeset.Cast(test.record, test.params, []string{"name", "age", "note", "address", "user_id"})
		statement, _ := sqlutil.NewBuilder("?", false).Insert(test.query.Collection, ch.Changes())

		t.Run("Insert|"+statement, func(t *testing.T) {
			assert.Nil(t, ch.Error())

			assert.Nil(t, test.query.Insert(nil, ch))
			assert.Nil(t, test.query.Insert(test.record, ch))

			// multiple insert
			assert.Nil(t, test.query.Insert(nil, ch, ch, ch))
		})

		t.Run("InsertAll|"+statement, func(t *testing.T) {
			assert.Nil(t, ch.Error())

			// multiple insert
			assert.Nil(t, test.query.Insert(nil, ch, ch, ch))
		})
	}
}

// InsertSet tests insert specifications only using Set query.
func InsertSet(t *testing.T, repo grimoire.Repo) {
	user := User{}
	assert.Nil(t, repo.From(users).Put(&user))
	now := time.Now()

	tests := []struct {
		query  grimoire.Query
		record interface{}
	}{
		{repo.From(users).Set("created_at", now).Set("updated_at", now).Set("name", "insert set"), &User{}},
		{repo.From(users).Set("created_at", now).Set("updated_at", now).Set("name", "insert set").Set("age", 100), &User{}},
		{repo.From(users).Set("created_at", now).Set("updated_at", now).Set("name", "insert set").Set("age", 100).Set("note", "note"), &User{}},
		{repo.From(users).Set("created_at", now).Set("updated_at", now).Set("note", "note"), &User{}},
		{repo.From(addresses).Set("created_at", now).Set("updated_at", now).Set("address", "address"), &Address{}},
		{repo.From(addresses).Set("created_at", now).Set("updated_at", now).Set("address", "address").Set("user_id", user.ID), &Address{}},
		{repo.From(addresses).Set("created_at", now).Set("updated_at", now).Set("user_id", user.ID), &Address{}},
	}

	for _, test := range tests {
		statement, _ := sqlutil.NewBuilder("?", false).Insert(test.query.Collection, test.query.Changes)

		t.Run("InsertSet|"+statement, func(t *testing.T) {
			assert.Nil(t, test.query.Insert(nil))
			assert.Nil(t, test.query.Insert(test.record))
		})
	}
}