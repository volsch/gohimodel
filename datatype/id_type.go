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

var idTypeSpec = newElementTypeSpecWithBase("id", stringTypeSpec)

var idRegexp = regexp.MustCompile("^[A-Za-z0-9\\-.]{1,64}$")

type idType struct {
	stringType
}

type IDAccessor interface {
	StringAccessor
}

func NewIDNil() IDAccessor {
	return newID(true, "")
}

func NewID(value string) IDAccessor {
	if !idRegexp.MatchString(value) {
		panic(fmt.Sprintf("not a valid ID: %s", value))
	}
	return newID(false, value)
}

func ParseID(value string) (IDAccessor, error) {
	if !idRegexp.MatchString(value) {
		return nil, fmt.Errorf("not a valid ID: %s", value)
	}
	return newID(false, value), nil
}

func newID(nilValue bool, value string) IDAccessor {
	return &idType{
		stringType{
			PrimitiveType: PrimitiveType{
				nilValue: nilValue,
			},
			value: value,
		},
	}
}

func (t *idType) DataType() DataTypes {
	return IDDataType
}

func (e *idType) TypeSpec() TypeSpecAccessor {
	return idTypeSpec
}
