package testdata

import (
	"context"
	"errors"
	"testing"
)

func TestSingleCache(t *testing.T) {
	d := New()
	meta := &Demo{ID: 1}
	getFromCache := func(c context.Context, id int64) (*Demo, error) { return meta, nil }
	notGetFromCache := func(c context.Context, id int64) (*Demo, error) { return nil, errors.New("err") }
	getFromSource := func(c context.Context, id int64) (*Demo, error) { return meta, nil }
	notGetFromSource := func(c context.Context, id int64) (*Demo, error) { return meta, errors.New("err") }
	addToCache := func(c context.Context, id int64, values *Demo) error { return nil }
	// get from cache
	_singleCacheFunc = getFromCache
	_singleRawFunc = notGetFromSource
	_singleAddCacheFunc = addToCache
	res, err := d.Demo(context.TODO(), 1)
	if err != nil {
		t.Fatalf("err should be nil, get: %v", err)
	}
	if res.ID != 1 {
		t.Fatalf("id should be 1")
	}
	// get from source
	_singleCacheFunc = notGetFromCache
	_singleRawFunc = getFromSource
	res, err = d.Demo(context.TODO(), 1)
	if err != nil {
		t.Fatalf("err should be nil, get: %v", err)
	}
	if res.ID != 1 {
		t.Fatalf("id should be 1")
	}
	// with null cache
	nullCache := &Demo{ID: -1}
	getNullFromCache := func(c context.Context, id int64) (*Demo, error) { return nullCache, nil }
	_singleCacheFunc = getNullFromCache
	_singleRawFunc = notGetFromSource
	res, err = d.Demo(context.TODO(), 1)
	if err != nil {
		t.Fatalf("err should be nil, get: %v", err)
	}
	if res != nil {
		t.Fatalf("res should be nil")
	}
}
