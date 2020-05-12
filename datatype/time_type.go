// Copyright (c) 2020, Volker Schmidt (volker@volsch.eu)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package datatype

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

var timeRegexp = regexp.MustCompile("^([01]\\d|2[0-3]):([0-5]\\d):([0-5]\\d|60)(?:\\.(\\d+))?$")
var fluentTimeRegexp = regexp.MustCompile("^([01]\\d|2[0-3])(?::([0-5]\\d)(?::([0-5]\\d|60)(?:\\.(\\d+))?)?)?$")

type TimeType struct {
	hour       int
	minute     int
	second     int
	nanosecond int
	precision  DateTimePrecisions
}

type TimeAccessor interface {
	PrimitiveAccessor
	Value() time.Time
	Hour() int
	Minute() int
	Second() int
	Nanosecond() int
	Precision() DateTimePrecisions
}

func NewTimeType(value time.Time) *TimeType {
	return NewTimeTypeHMSN(value.Hour(), value.Minute(), value.Second(), value.Nanosecond())
}

func NewTimeTypeHMSN(hour int, minute int, second int, nanosecond int) *TimeType {
	return &TimeType{hour: hour, minute: minute, second: second, nanosecond: nanosecond, precision: NanoTimePrecision}
}

func ParseTimeValue(value string) (*TimeType, error) {
	parts := timeRegexp.FindStringSubmatch(value)
	if parts == nil {
		return nil, fmt.Errorf("not a valid time string: %s", value)
	}
	return newTimeType(parts), nil
}

func ParseFluentTimeValue(value string) (*TimeType, error) {
	parts := fluentTimeRegexp.FindStringSubmatch(value)
	if parts == nil {
		return nil, fmt.Errorf("not a valid fluent time string: %s", value)
	}
	return newTimeType(parts), nil
}

func newTimeType(parts []string) *TimeType {
	hour, _ := strconv.Atoi(parts[1])
	precision := HourTimePrecision

	minute := 0
	if parts[2] != "" {
		minute, _ = strconv.Atoi(parts[2])
		precision = MinuteTimePrecision
	}

	second := 0
	if parts[3] != "" {
		second, _ = strconv.Atoi(parts[3])
		precision = SecondTimePrecision
	}

	nanosecond := 0
	if parts[4] != "" {
		nanosecond = parseNanosecond(parts[4])
		precision = NanoTimePrecision
	}

	return &TimeType{hour: hour, minute: minute, second: second, nanosecond: nanosecond, precision: precision}
}

func parseNanosecond(value string) int {
	if value == "" {
		return 0
	}
	nanoValue := value
	if len(nanoValue) > 9 {
		nanoValue = nanoValue[0:9]
	}
	nano, _ := strconv.Atoi(nanoValue)
	nano = nano * int(math.Pow10(9-len(nanoValue)))
	return nano
}

func (t *TimeType) DataType() DataTypes {
	return TimeDataType
}

func (t *TimeType) Hour() int {
	return t.hour
}

func (t *TimeType) Minute() int {
	return t.minute
}

func (t *TimeType) Second() int {
	return t.second
}

func (t *TimeType) Nanosecond() int {
	return t.nanosecond
}

func (t *TimeType) Value() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), t.hour, t.minute, t.second, t.nanosecond, now.Location())
}

func (t *TimeType) Precision() DateTimePrecisions {
	return t.precision
}
