package main

import (
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}
// c.Increment is
// (*c).Increment()

func (c Counter) GetTotal() int {
	return c.total
}

func (c *Counter) WhatTimeIsIt() time.Time {

	return time.Now()
}



type Foo struct {
	Field1 string
	Field2 int
}

func MakeFooAntiPattern(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}
// メモ：json.Unmarshalとresty.SetResultがそのパターンを使っている
// 参考：
// https://pkg.go.dev/encoding/json#Unmarshal
// https://pkg.go.dev/github.com/go-resty/resty/v2#Request.SetResult


func MakeFooOK() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}