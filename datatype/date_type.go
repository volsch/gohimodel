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
	"regexp"
	"strconv"
	"time"
)

var dateRegexp = regexp.MustCompile("^(\\d(?:\\d(?:\\d[1-9]|[1-9]0)|[1-9]00)|[1-9]000)(?:-(0[1-9]|1[0-2])(?:-(0[1-9]|[1-2]\\d|3[0-1]))?)?$")

type DateType struct {
	year      int
	month     int
	day       int
	precision DateTimePrecisions
}

type DateAccessor interface {
	PrimitiveAccessor
	Value() time.Time
	Year() int
	Month() int
	Day() int
	Precision() DateTimePrecisions
}

func NewDateType(value time.Time) *DateType {
	return NewDateTypeYMD(value.Year(), int(value.Month()), value.Day())
}

func NewDateTypeYMD(year int, month int, day int) *DateType {
	return &DateType{year: year, month: month, day: day, precision: DayDatePrecision}
}

func ParseDateValue(value string) (*DateType, error) {
	parts := dateRegexp.FindStringSubmatch(value)
	if parts == nil {
		return nil, fmt.Errorf("not a valid date string: %s", value)
	}
	return newDateType(parts), nil
}

func newDateType(parts []string) *DateType {
	year, _ := strconv.Atoi(parts[1])
	precision := YearDatePrecision

	month := 1
	if parts[2] != "" {
		month, _ = strconv.Atoi(parts[2])
		precision = MonthDatePrecision
	}

	day := 1
	if parts[3] != "" {
		day, _ = strconv.Atoi(parts[3])
		precision = DayDatePrecision
	}

	return &DateType{year: year, month: month, day: day, precision: precision}
}

func (t *DateType) DataType() DataTypes {
	return DateDataType
}

func (t *DateType) Year() int {
	return t.year
}

func (t *DateType) Month() int {
	return t.month
}

func (t *DateType) Day() int {
	return t.day
}

func (t *DateType) Value() time.Time {
	return time.Date(t.year, time.Month(t.month), t.day, 0, 0, 0, 0, time.Local)
}

func (t *DateType) Precision() DateTimePrecisions {
	return t.precision
}
