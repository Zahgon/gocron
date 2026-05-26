package gocron

import (
	"errors"
	"reflect"
	"time"
)

var (
	ErrTimeFormat           = errors.New("time format error")
	ErrParamsNotAdapted     = errors.New("the number of params is not adapted")
	ErrNotAFunction         = errors.New("only functions can be schedule into the job queue")
	ErrPeriodNotSpecified   = errors.New("unspecified job period")
	ErrParameterCannotBeNil = errors.New("nil paramaters cannot be used with reflection")
)

// Job struct keeping information about job
type Job struct {
	interval uint64                   // pause interval * unit between runs
	jobFunc  string                   // the job jobFunc to run, func[jobFunc]
	unit     timeUnit                 // time units, ,e.g. 'minutes', 'hours'...
	atTime   time.Duration            // optional time at which this job runs
	err      error                    // error related to job
	loc      *time.Location           // optional timezone that the atTime is in
	lastRun  time.Time                // datetime of last run
	nextRun  time.Time                // datetime of next run
	startDay time.Weekday             // Specific day of the week to start on
	funcs    map[string]interface{}   // Map for the function task store
	fparams  map[string][]interface{} // Map for function and  params of function
	lock     bool                     // lock the job from running at same time form multiple instances
	tags     []string                 // allow the user to tag jobs with certain labels
}

// NewJob creates a new job with the time interval.
func NewJob(interval uint64) *Job { _ = "STUB: not implemented"; return nil }

// True if the job should be run now
func (j *Job) shouldRun() bool { _ = "STUB: not implemented"; return false }

// Run the job and immediately reschedule it
func (j *Job) run() ([]reflect.Value, error) { _ = "STUB: not implemented"; return nil, nil }

// Err should be checked to ensure an error didn't occur creating the job
func (j *Job) Err() error {
	_ = "STUB: not implemented"

	// Do specifies the jobFunc that should be called every time the job runs
	return nil
}

func (j *Job) Do(jobFun interface{}, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// DoSafely does the same thing as Do, but logs unexpected panics, instead of unwinding them up the chain
// Deprecated: DoSafely exists due to historical compatibility and will be removed soon. Use Do instead
func (j *Job) DoSafely(jobFun interface{}, params ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// At schedules job at specific time of day
//
//	s.Every(1).Day().At("10:30:01").Do(task)
//	s.Every(1).Monday().At("10:30:01").Do(task)
func (j *Job) At(t string) *Job { _ = "STUB: not implemented"; return nil }

// save atTime start as duration from midnight

// GetAt returns the specific time of day the job will run at
//
//	s.Every(1).Day().At("10:30").GetAt() == "10:30"
func (j *Job) GetAt() string { _ = "STUB: not implemented"; return "" }

// Loc sets the location for which to interpret "At"
//
//	s.Every(1).Day().At("10:30").Loc(time.UTC).Do(task)
func (j *Job) Loc(loc *time.Location) *Job { _ = "STUB: not implemented"; return nil }

// Tag allows you to add labels to a job
// they don't impact the functionality of the job.
func (j *Job) Tag(t string, others ...string) { _ = "STUB: not implemented"; return }

// Untag removes a tag from a job
func (j *Job) Untag(t string) { _ = "STUB: not implemented"; return }

// Tags returns the tags attached to the job
func (j *Job) Tags() []string { _ = "STUB: not implemented"; return nil }

func (j *Job) periodDuration() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// roundToMidnight truncate time to midnight
func (j *Job) roundToMidnight(t time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// scheduleNextRun Compute the instant when this job should run next
func (j *Job) scheduleNextRun() error { _ = "STUB: not implemented"; return nil }

// advance to next possible schedule

// NextScheduledTime returns the time of when this job is to run next
func (j *Job) NextScheduledTime() time.Time {
	_ = "STUB: not implemented"

	// set the job's unit with seconds,minutes,hours...
	return *new(time.Time)
}

func (j *Job) mustInterval(i uint64) error { _ = "STUB: not implemented"; return nil }

// From schedules the next run of the job
func (j *Job) From(t *time.Time) *Job { _ = "STUB: not implemented"; return nil }

// setUnit sets unit type
func (j *Job) setUnit(unit timeUnit) *Job { _ = "STUB: not implemented"; return nil }

// Seconds set the unit with seconds
func (j *Job) Seconds() *Job { _ = "STUB: not implemented"; return nil }

// Minutes set the unit with minute
func (j *Job) Minutes() *Job { _ = "STUB: not implemented"; return nil }

// Hours set the unit with hours
func (j *Job) Hours() *Job { _ = "STUB: not implemented"; return nil }

// Days set the job's unit with days
func (j *Job) Days() *Job { _ = "STUB: not implemented"; return nil }

// Weeks sets the units as weeks
func (j *Job) Weeks() *Job { _ = "STUB: not implemented"; return nil }

// Second sets the unit with second
func (j *Job) Second() *Job { _ = "STUB: not implemented"; return nil }

// Minute sets the unit  with minute, which interval is 1
func (j *Job) Minute() *Job { _ = "STUB: not implemented"; return nil }

// Hour sets the unit with hour, which interval is 1
func (j *Job) Hour() *Job { _ = "STUB: not implemented"; return nil }

// Day sets the job's unit with day, which interval is 1
func (j *Job) Day() *Job { _ = "STUB: not implemented"; return nil }

// Week sets the job's unit with week, which interval is 1
func (j *Job) Week() *Job { _ = "STUB: not implemented"; return nil }

// Weekday start job on specific Weekday
func (j *Job) Weekday(startDay time.Weekday) *Job { _ = "STUB: not implemented"; return nil }

// GetWeekday returns which day of the week the job will run on
// This should only be used when .Weekday(...) was called on the job.
func (j *Job) GetWeekday() time.Weekday {
	_ = "STUB: not implemented"

	// Monday set the start day with Monday
	// - s.Every(1).Monday().Do(task)
	return *new(time.Weekday)
}

func (j *Job) Monday() (job *Job) { _ = "STUB: not implemented"; return nil }

// Tuesday sets the job start day Tuesday
func (j *Job) Tuesday() *Job { _ = "STUB: not implemented"; return nil }

// Wednesday sets the job start day Wednesday
func (j *Job) Wednesday() *Job { _ = "STUB: not implemented"; return nil }

// Thursday sets the job start day Thursday
func (j *Job) Thursday() *Job { _ = "STUB: not implemented"; return nil }

// Friday sets the job start day Friday
func (j *Job) Friday() *Job { _ = "STUB: not implemented"; return nil }

// Saturday sets the job start day Saturday
func (j *Job) Saturday() *Job { _ = "STUB: not implemented"; return nil }

// Sunday sets the job start day Sunday
func (j *Job) Sunday() *Job { _ = "STUB: not implemented"; return nil }

// Lock prevents job to run from multiple instances of gocron
func (j *Job) Lock() *Job { _ = "STUB: not implemented"; return nil }
