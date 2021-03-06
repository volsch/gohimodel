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
)

var uriTypeSpec = newElementTypeSpec("uri")

var uriRegexp = regexp.MustCompile("^\\S*$")

type uriType struct {
	PrimitiveType
	value string
}

type URIAccessor interface {
	PrimitiveAccessor
}

func IsURI(accessor Accessor) bool {
	dt := accessor.DataType()
	return dt == URIDataType
}

func IsURINoURL(accessor Accessor) bool {
	dt := accessor.DataType()
	return dt == URIDataType
}

func NewURINil() URIAccessor {
	return newURI(true, "")
}

func NewURI(value string) URIAccessor {
	if !uriRegexp.MatchString(value) {
		panic(fmt.Sprintf("not a valid URI: %s", value))
	}
	return newURI(false, value)
}

func ParseURI(value string) (URIAccessor, error) {
	if !uriRegexp.MatchString(value) {
		return nil, fmt.Errorf("not a valid URI: %s", value)
	}
	return newURI(false, value), nil
}

func newURI(nilValue bool, value string) URIAccessor {
	return &uriType{
		PrimitiveType: PrimitiveType{
			nilValue: nilValue,
		},
		value: value,
	}
}

func (t *uriType) String() string {
	return t.value
}

func (t *uriType) DataType() DataTypes {
	return URIDataType
}

func (e *uriType) TypeSpec() TypeSpecAccessor {
	return uriTypeSpec
}

func (t *uriType) Equal(accessor Accessor) bool {
	if o, ok := accessor.(URIAccessor); !ok || !IsURI(accessor) {
		return false
	} else {
		return t.Nil() == o.Nil() && t.String() == o.String()
	}
}

func (t *uriType) Equivalent(accessor Accessor) bool {
	return t.Equal(accessor)
}
