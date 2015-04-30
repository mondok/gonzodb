package util

import (
	"errors"
	"reflect"
)

type T interface{}

type Query struct {
	values []T
	err    error
}

type queryable interface {
	Results() (T, error)
}

func From(input T) Query {
	var e error
	if input == nil {
		e = errors.New("")
	}

	out, ok := takeSliceArg(input)
	if !ok {
		e = errors.New("")
		out = nil
	}

	return Query{values: out, err: e}
}

func (q Query) First() (elem T, found bool, err error) {
	if q.err != nil {
		err = q.err
		return
	}
	if len(q.values) > 0 {
		found = true
		elem = q.values[0]
	}
	return
}

func takeSliceArg(arg T) (out []T, ok bool) {
	slice, success := takeArg(arg, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]T, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return out, true
}

func takeArg(arg T, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

func (q Query) Results() ([]T, error) {
	// we need to copy results for isolating user modification on returned
	// slice from the current Query instance.
	res := make([]T, len(q.values))
	_ = copy(res, q.values)
	return res, q.err
}

func (q Query) Where(f func(T) (bool, error)) (r Query) {
	if q.err != nil {
		r.err = q.err
		return r
	}
	if f == nil {
		r.err = errors.New("Function is not defined")
		return
	}

	for _, i := range q.values {
		ok, err := f(i)
		if err != nil {
			r.err = err
			return r
		}
		if ok {
			r.values = append(r.values, i)
		}
	}
	return
}
