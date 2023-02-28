// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MultiInt8Value struct {
	_tab flatbuffers.Table
}

func GetRootAsMultiInt8Value(buf []byte, offset flatbuffers.UOffsetT) *MultiInt8Value {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MultiInt8Value{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMultiInt8Value(buf []byte, offset flatbuffers.UOffsetT) *MultiInt8Value {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MultiInt8Value{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MultiInt8Value) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MultiInt8Value) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MultiInt8Value) Value(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *MultiInt8Value) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MultiInt8Value) MutateValue(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func MultiInt8ValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MultiInt8ValueAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func MultiInt8ValueStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MultiInt8ValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}