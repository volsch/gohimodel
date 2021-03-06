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
	"strings"
)

type QuantityComparator CodeAccessor

var (
	LessThanQuantityComparator           QuantityComparator = NewCode("<")
	LessOrEqualThanQuantityComparator    QuantityComparator = NewCode("<=")
	GreaterThanQuantityComparator        QuantityComparator = NewCode(">")
	GreaterOrEqualThanQuantityComparator QuantityComparator = NewCode(">=")
)

var UCUMSystemURI = NewURI("http://unitsofmeasure.org")

var quantityTypeSpec = newElementTypeSpec("Quantity")

type quantityType struct {
	value      DecimalAccessor
	comparator QuantityComparator
	unit       StringAccessor
	system     URIAccessor
	code       CodeAccessor
}

type QuantityAccessor interface {
	ElementAccessor
	Stringifier

	Value() DecimalAccessor
	Comparator() QuantityComparator
	Unit() StringAccessor
	System() URIAccessor
	Code() CodeAccessor
}

type QuantityModifier interface {
	QuantityAccessor

	SetValue(value DecimalAccessor) QuantityModifier
	SetComparator(value QuantityComparator) QuantityModifier
	SetUnit(value StringAccessor) QuantityModifier
	SetSystem(value URIAccessor) QuantityModifier
	SetCode(value CodeAccessor) QuantityModifier
}

func NewQuantityEmpty() QuantityModifier {
	return &quantityType{}
}

func NewQuantity(value DecimalAccessor, comparator QuantityComparator,
	unit StringAccessor, system URIAccessor, code CodeAccessor) QuantityModifier {
	return &quantityType{
		value:      value,
		comparator: comparator,
		unit:       unit,
		system:     system,
		code:       code,
	}
}

func (t *quantityType) DataType() DataTypes {
	return QuantityDataType
}

func (t *quantityType) Empty() bool {
	return t.value == nil &&
		t.comparator == nil &&
		t.unit == nil &&
		t.system == nil &&
		t.code == nil
}

func (t *quantityType) Value() DecimalAccessor {
	return t.value
}

func (t *quantityType) Comparator() QuantityComparator {
	return t.comparator
}

func (t *quantityType) Unit() StringAccessor {
	return t.unit
}

func (t *quantityType) System() URIAccessor {
	return t.system
}

func (t *quantityType) Code() CodeAccessor {
	return t.code
}

func (t *quantityType) SetValue(value DecimalAccessor) QuantityModifier {
	t.value = value
	return t
}

func (t *quantityType) SetComparator(value QuantityComparator) QuantityModifier {
	t.comparator = value
	return t
}

func (t *quantityType) SetUnit(value StringAccessor) QuantityModifier {
	t.unit = value
	return t
}

func (t *quantityType) SetSystem(value URIAccessor) QuantityModifier {
	t.system = value
	return t
}

func (t *quantityType) SetCode(value CodeAccessor) QuantityModifier {
	t.code = value
	return t
}

func (e *quantityType) TypeSpec() TypeSpecAccessor {
	return quantityTypeSpec
}

func (t *quantityType) Equal(accessor Accessor) bool {
	if accessor == nil {
		return false
	}
	if o, ok := accessor.(QuantityAccessor); !ok {
		return false
	} else {
		return Equal(t.value, o.Value()) &&
			Equal(t.comparator, o.Comparator()) &&
			Equal(t.unit, o.Unit()) &&
			Equal(t.system, o.System()) &&
			Equal(t.code, o.Code())
	}
}

func (t *quantityType) Equivalent(accessor Accessor) bool {
	return quantityEqual(t, accessor)
}

func quantityEqual(t QuantityAccessor, accessor Accessor) bool {
	if o, ok := accessor.(QuantityAccessor); !ok {
		return false
	} else {
		return Equal(t.Value(), o.Value()) &&
			Equal(t.System(), o.System()) &&
			Equal(t.Code(), o.Code())
	}
}

func (t *quantityType) String() string {
	var b strings.Builder
	b.Grow(32)
	if t.value != nil {
		b.WriteString(t.value.String())
	}
	if t.code != nil {
		if b.Len() > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(t.code.String())
	}
	return b.String()
}
