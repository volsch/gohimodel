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
	"github.com/stretchr/testify/assert"
	"testing"
)

var testTypeInfo = newElementTypeInfo("test")

func TestCollectionDataType(t *testing.T) {
	c := NewCollection(testTypeInfo)
	assert.Equal(t, CollectionDataType, c.DataType())
}

func TestCollectionTypeInfo(t *testing.T) {
	c := NewCollection(testTypeInfo)
	assert.Equal(t, "Collection", c.TypeInfo().String())
}

func TestNewCollectionNoItemType(t *testing.T) {
	assert.Panics(t, func() { NewCollection(nil) })
}

func TestNewCollection(t *testing.T) {
	c := NewCollection(testTypeInfo)
	assert.Same(t, testTypeInfo, c.ItemTypeInfo())
	assert.True(t, c.Empty(), "new collection must be empty")
	assert.Equal(t, 0, c.Count())
}

func TestNewCollectionGetEmpty(t *testing.T) {
	c := NewCollection(testTypeInfo)
	assert.Panics(t, func() { c.Get(0) })
}

func TestCollectionAddGet(t *testing.T) {
	item1 := newAccessorMock()
	item2 := newAccessorMock()
	c := NewCollection(testTypeInfo)
	c.Add(item1)
	c.Add(item2)
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
}

type accessorMock struct {
}

func newAccessorMock() Accessor {
	return &accessorMock{}
}

func (a accessorMock) DataType() DataTypes {
	panic("implement me")
}

func (a accessorMock) TypeInfo() TypeInfoAccessor {
	panic("implement me")
}

func (a accessorMock) Empty() bool {
	panic("implement me")
}