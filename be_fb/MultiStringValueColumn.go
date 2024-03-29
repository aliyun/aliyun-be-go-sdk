// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MultiStringValueColumn struct {
	_tab flatbuffers.Table
}

func GetRootAsMultiStringValueColumn(buf []byte, offset flatbuffers.UOffsetT) *MultiStringValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MultiStringValueColumn{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMultiStringValueColumn(buf []byte, offset flatbuffers.UOffsetT) *MultiStringValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MultiStringValueColumn{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MultiStringValueColumn) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MultiStringValueColumn) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MultiStringValueColumn) Value(obj *MultiStringValue, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *MultiStringValueColumn) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MultiStringValueColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MultiStringValueColumnAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func MultiStringValueColumnStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MultiStringValueColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
