// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MultiInt8ValueColumn struct {
	_tab flatbuffers.Table
}

func GetRootAsMultiInt8ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *MultiInt8ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MultiInt8ValueColumn{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMultiInt8ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *MultiInt8ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MultiInt8ValueColumn{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MultiInt8ValueColumn) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MultiInt8ValueColumn) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MultiInt8ValueColumn) Value(obj *MultiInt8Value, j int) bool {
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

func (rcv *MultiInt8ValueColumn) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MultiInt8ValueColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MultiInt8ValueColumnAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func MultiInt8ValueColumnStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MultiInt8ValueColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}