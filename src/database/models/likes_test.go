// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

func testLikes(t *testing.T) {
	t.Parallel()

	query := Likes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLikesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
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

	count, err := Likes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLikesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Likes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Likes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLikesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LikeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Likes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLikesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LikeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Like exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LikeExists to return true, but got false.")
	}
}

func testLikesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	likeFound, err := FindLike(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if likeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLikesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Likes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLikesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Likes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLikesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	likeOne := &Like{}
	likeTwo := &Like{}
	if err = randomize.Struct(seed, likeOne, likeDBTypes, false, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}
	if err = randomize.Struct(seed, likeTwo, likeDBTypes, false, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = likeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = likeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Likes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLikesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	likeOne := &Like{}
	likeTwo := &Like{}
	if err = randomize.Struct(seed, likeOne, likeDBTypes, false, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}
	if err = randomize.Struct(seed, likeTwo, likeDBTypes, false, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = likeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = likeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Likes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func likeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Like) error {
	*o = Like{}
	return nil
}

func likeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Like) error {
	*o = Like{}
	return nil
}

func likeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Like) error {
	*o = Like{}
	return nil
}

func likeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Like) error {
	*o = Like{}
	return nil
}

func likeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Like) error {
	*o = Like{}
	return nil
}

func likeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Like) error {
	*o = Like{}
	return nil
}

func likeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Like) error {
	*o = Like{}
	return nil
}

func likeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Like) error {
	*o = Like{}
	return nil
}

func likeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Like) error {
	*o = Like{}
	return nil
}

func testLikesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Like{}
	o := &Like{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, likeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Like object: %s", err)
	}

	AddLikeHook(boil.BeforeInsertHook, likeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	likeBeforeInsertHooks = []LikeHook{}

	AddLikeHook(boil.AfterInsertHook, likeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	likeAfterInsertHooks = []LikeHook{}

	AddLikeHook(boil.AfterSelectHook, likeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	likeAfterSelectHooks = []LikeHook{}

	AddLikeHook(boil.BeforeUpdateHook, likeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	likeBeforeUpdateHooks = []LikeHook{}

	AddLikeHook(boil.AfterUpdateHook, likeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	likeAfterUpdateHooks = []LikeHook{}

	AddLikeHook(boil.BeforeDeleteHook, likeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	likeBeforeDeleteHooks = []LikeHook{}

	AddLikeHook(boil.AfterDeleteHook, likeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	likeAfterDeleteHooks = []LikeHook{}

	AddLikeHook(boil.BeforeUpsertHook, likeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	likeBeforeUpsertHooks = []LikeHook{}

	AddLikeHook(boil.AfterUpsertHook, likeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	likeAfterUpsertHooks = []LikeHook{}
}

func testLikesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Likes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLikesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(likeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Likes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLikeToOneBookUsingBook(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Like
	var foreign Book

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, likeDBTypes, false, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, bookDBTypes, false, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.BookID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Book().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LikeSlice{&local}
	if err = local.L.LoadBook(ctx, tx, false, (*[]*Like)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Book == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Book = nil
	if err = local.L.LoadBook(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Book == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLikeToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Like
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, likeDBTypes, false, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LikeSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Like)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLikeToOneSetOpBookUsingBook(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Like
	var b, c Book

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, likeDBTypes, false, strmangle.SetComplement(likePrimaryKeyColumns, likeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Book{&b, &c} {
		err = a.SetBook(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Book != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Likes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BookID != x.ID {
			t.Error("foreign key was wrong value", a.BookID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BookID))
		reflect.Indirect(reflect.ValueOf(&a.BookID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.BookID != x.ID {
			t.Error("foreign key was wrong value", a.BookID, x.ID)
		}
	}
}
func testLikeToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Like
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, likeDBTypes, false, strmangle.SetComplement(likePrimaryKeyColumns, likeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Likes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testLikesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
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

func testLikesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LikeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLikesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Likes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	likeDBTypes = map[string]string{`ID`: `integer`, `UserID`: `integer`, `BookID`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_           = bytes.MinRead
)

func testLikesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(likePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(likeAllColumns) == len(likePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Likes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, likeDBTypes, true, likePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLikesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(likeAllColumns) == len(likePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Like{}
	if err = randomize.Struct(seed, o, likeDBTypes, true, likeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Likes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, likeDBTypes, true, likePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(likeAllColumns, likePrimaryKeyColumns) {
		fields = likeAllColumns
	} else {
		fields = strmangle.SetComplement(
			likeAllColumns,
			likePrimaryKeyColumns,
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

	slice := LikeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLikesUpsert(t *testing.T) {
	t.Parallel()

	if len(likeAllColumns) == len(likePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Like{}
	if err = randomize.Struct(seed, &o, likeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Like: %s", err)
	}

	count, err := Likes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, likeDBTypes, false, likePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Like struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Like: %s", err)
	}

	count, err = Likes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
