// Copyright (c) 2020-2021, Volker Schmidt (volker@volsch.eu)
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
	"strings"
	"time"
)

var dateTypeSpec = newElementTypeSpec("date")

var dateRegexp = regexp.MustCompile("^(\\d(?:\\d(?:\\d[1-9]|[1-9]0)|[1-9]00)|[1-9]000)(?:-(0[1-9]|1[0-2])(?:-(0[1-9]|[1-2]\\d|3[0-1]))?)?$")

type dateType struct {
	TemporalType
	year  int
	month int
	day   int
}

type DateAccessor interface {
	DateTemporalAccessor
	Year() int
	Month() int
	Day() int
}

func NewDateNil() DateAccessor {
	return newDate(true, 1970, 1, 1, DayDatePrecision)
}

func NewDate(value time.Time) DateAccessor {
	return NewDateYMD(value.Year(), int(value.Month()), value.Day())
}

func NewDateYMD(year int, month int, day int) DateAccessor {
	return newDate(false, year, month, day, DayDatePrecision)
}

func NewDateYMDWithPrecision(year int, month int, day int, precision DateTimePrecisions) DateAccessor {
	if precision <= YearDatePrecision {
		precision = YearDatePrecision
	} else if precision > DayDatePrecision {
		precision = DayDatePrecision
	}

	if precision < DayDatePrecision {
		day = 1
	}
	if precision < MonthDatePrecision {
		month = 1
	}

	return newDate(false, year, month, day, precision)
}

func ParseDate(value string) (DateAccessor, error) {
	parts := dateRegexp.FindStringSubmatch(value)
	if parts == nil {
		return nil, fmt.Errorf("not a valid date string: %s", value)
	}
	return newDateFromParts(parts), nil
}

func newDateFromParts(parts []string) DateAccessor {
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

	return newDate(false, year, month, day, precision)
}

func newDate(nilValue bool, year int, month int, day int, precision DateTimePrecisions) DateAccessor {
	return &dateType{
		TemporalType: TemporalType{
			PrimitiveType: PrimitiveType{
				nilValue: nilValue,
			},
			precision: precision,
		},
		year:  year,
		month: month,
		day:   day,
	}
}

func (t *dateType) DataType() DataTypes {
	return DateDataType
}

func (t *dateType) LowestPrecision() DateTimePrecisions {
	return YearDatePrecision
}

func (t *dateType) Year() int {
	return t.year
}

func (t *dateType) Month() int {
	return t.month
}

func (t *dateType) Day() int {
	return t.day
}

func (t *dateType) Time() time.Time {
	return time.Date(t.year, time.Month(t.month), t.day, 0, 0, 0, 0, time.Local)
}

func (e *dateType) TypeSpec() TypeSpecAccessor {
	return dateTypeSpec
}

func (t *dateType) Equal(accessor Accessor) bool {
	if o, ok := accessor.(DateTemporalAccessor); !ok {
		return false
	} else {
		return t.Precision() == o.Precision() && dateValueEqual(t, o)
	}
}

func (t *dateType) Equivalent(accessor Accessor) bool {
	if o, ok := accessor.(DateTemporalAccessor); !ok {
		return false
	} else {
		return dateValueEqual(t, o)
	}
}

func dateValueEqual(dt1 DateTemporalAccessor, dt2 DateTemporalAccessor) bool {
	return dt1.Nil() == dt2.Nil() && dt1.Year() == dt2.Year() && dt1.Month() == dt2.Month() && dt1.Day() == dt2.Day()
}

func (t *dateType) String() string {
	if t.nilValue {
		return ""
	}

	var b strings.Builder
	b.Grow(10)

	writeStringBuilderInt(&b, t.year, 4)
	if t.precision >= MonthDatePrecision {
		b.WriteByte('-')
		writeStringBuilderInt(&b, int(t.month), 2)
	}
	if t.precision >= DayDatePrecision {
		b.WriteByte('-')
		writeStringBuilderInt(&b, t.day, 2)
	}

	return b.String()
}
