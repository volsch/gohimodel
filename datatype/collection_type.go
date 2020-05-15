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

var collectionTypeInfo = NewTypeInfo(NewTypeName("Collection"), nil)

type CollectionType struct {
	itemTypeInfo TypeInfoAccessor
	items        []Accessor
}

type CollectionAccessor interface {
	ItemTypeInfo() TypeInfoAccessor
	Count() int
	Get(i int) Accessor
}

type CollectionModifier interface {
	CollectionAccessor
	Add(accessor Accessor)
}

func NewCollection(itemTypeInfo TypeInfoAccessor) *CollectionType {
	if itemTypeInfo == nil {
		panic("no item type has been specified")
	}
	return &CollectionType{
		itemTypeInfo: itemTypeInfo,
	}
}

func (c *CollectionType) DataType() DataTypes {
	return CollectionDataType
}

func (c *CollectionType) TypeInfo() TypeInfoAccessor {
	return collectionTypeInfo
}

func (c *CollectionType) ItemTypeInfo() TypeInfoAccessor {
	return c.itemTypeInfo
}

func (c *CollectionType) Empty() bool {
	return c.Count() == 0
}

func (c *CollectionType) Count() int {
	if c.items == nil {
		return 0
	}
	return len(c.items)
}

func (c *CollectionType) Get(i int) Accessor {
	if c.items == nil {
		panic("collection is empty")
	}
	return c.items[i]
}

func (c *CollectionType) Add(accessor Accessor) {
	if c.items == nil {
		c.items = make([]Accessor, 0)
	}
	c.items = append(c.items, accessor)
}