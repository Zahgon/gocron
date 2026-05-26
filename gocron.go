// Package gocron : A Golang Job Scheduling Package.
//
// Note from current maintainers:
//
// A currently maintained fork of this project has been migrated to https://github.com/go-co-op/gocron
//
// Disclaimer: we (the maintainers) tried, with no luck, to get in contact with Jason (the repository owner) in order to add new maintainers or leave the project within an organization. Unfortunately, he hasn't replied for months now (March, 2020).
//
// So, we decided to move the project to a new repository (as stated above), in order to keep the evolution of the project coming from as many people as possible. Feel free to reach over!
//
// An in-process scheduler for periodic jobs that uses the builder pattern
// for configuration. Schedule lets you run Golang functions periodically
// at pre-determined intervals using a simple, human-friendly syntax.
//
// Inspired by the Ruby module clockwork <https://github.com/tomykaira/clockwork>
// and
// Python package schedule <https://github.com/dbader/schedule>
//
// See also
// http://adam.heroku.com/past/2010/4/13/rethinking_cron/
// http://adam.heroku.com/past/2010/6/30/replace_cron_with_clockwork/
//
// Copyright 2014 Jason Lyu. jasonlvhit@gmail.com .
// All rights reserved.
// Use of this source code is governed by a BSD-style .
// license that can be found in the LICENSE file.
package gocron

import (
	"reflect"
	"time"
)

// Locker provides a method to lock jobs from running
// at the same time on multiple instances of gocron.
// You can provide any locker implementation you wish.
type Locker interface {
	Lock(key string) (bool, error)
	Unlock(key string) error
}

type timeUnit int

// MAXJOBNUM max number of jobs, hack it if you need.
const MAXJOBNUM = 10000

//go:generate stringer -type=timeUnit
const (
	seconds timeUnit = iota + 1
	minutes
	hours
	days
	weeks
)

var (
	loc    = time.Local // Time location, default set by the time.Local (*time.Location)
	locker Locker
)

// ChangeLoc change default the time location
func ChangeLoc(newLocation *time.Location) { _ = "STUB: not implemented"; return }

// SetLocker sets a locker implementation
func SetLocker(l Locker) { _ = "STUB: not implemented"; return }

func callJobFuncWithParams(jobFunc interface{}, params []interface{}) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for given function fn, get the name of function.
func getFunctionName(fn interface{}) string { _ = "STUB: not implemented"; return "" }

func getFunctionKey(funcName string) string { _ = "STUB: not implemented"; return "" }

// Jobs returns the list of Jobs from the defaultScheduler
func Jobs() []*Job { _ = "STUB: not implemented"; return nil }

func formatTime(t string) (hour, min, sec int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

// NextTick returns a pointer to a time that will run at the next tick
func NextTick() *time.Time { _ = "STUB: not implemented"; return nil }
