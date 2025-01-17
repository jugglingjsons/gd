//go:build !generate

package gd

import (
	"reflect"

	"grow.graphics/gd/internal/callframe"
	"grow.graphics/uc"
	"grow.graphics/xy"

	"runtime.link/mmm"
)

type Bool = bool

type (
	Float       = float64
	Int         = int64
	Vector2     = xy.Vector2
	Vector2i    = xy.Vector2i
	Rect2       = xy.Rect2
	Rect2i      = xy.Rect2i
	Vector3     = xy.Vector3
	Vector3i    = xy.Vector3i
	Transform2D = xy.Transform2D
	Vector4     = xy.Vector4
	Vector4i    = xy.Vector4i
	Plane       = xy.Plane
	Quaternion  = xy.Quaternion
	AABB        = xy.AABB
	Basis       = xy.Basis
	Transform3D = xy.Transform3D
	Projection  = xy.Projection
)

type Color = uc.Color

type (
	Side       = xy.Side
	EulerOrder = xy.EulerOrder
)

type RID uint64

type Callable mmm.Pointer[API, Callable, [2]uintptr]

func (c Callable) Free() {
	var frame = callframe.New()
	mmm.API(c).typeset.destruct.Callable(callframe.Arg(frame, mmm.End(c)).Uintptr())
	frame.Free()
}

type Variant mmm.Pointer[API, Variant, [3]uintptr]

func (s Variant) Free() {
	mmm.API(s).Variants.Destroy(s)
	mmm.End(s)
}

type Iterator struct {
	self Variant
	iter Variant
}

func (iter Iterator) Next() bool {
	return mmm.API(iter.self).Variants.IteratorNext(iter.self, iter.iter)
}

func (iter Iterator) Value(ctx Lifetime) Variant {
	val, ok := mmm.API(iter.self).Variants.IteratorGet(ctx, iter.self, iter.iter)
	if !ok {
		panic("failed to get iterator value")
	}
	return val
}

func variantTypeFromName(s string) (VariantType, reflect.Type) {
	switch s {
	case "Nil":
		return TypeNil, nil
	case "bool", "Bool":
		return TypeBool, reflect.TypeOf(false)
	case "int", "Int":
		return TypeInt, reflect.TypeOf(int64(0))
	case "float", "Float":
		return TypeFloat, reflect.TypeOf(Float(0))
	case "String":
		return TypeString, reflect.TypeOf(String{})
	case "Vector2":
		return TypeVector2, reflect.TypeOf(Vector2{})
	case "Vector2i":
		return TypeVector2i, reflect.TypeOf(Vector2i{})
	case "Rect2":
		return TypeRect2, reflect.TypeOf(Rect2{})
	case "Rect2i":
		return TypeRect2i, reflect.TypeOf(Rect2i{})
	case "Vector3":
		return TypeVector3, reflect.TypeOf(Vector3{})
	case "Vector3i":
		return TypeVector3i, reflect.TypeOf(Vector3i{})
	case "Transform2D":
		return TypeTransform2d, reflect.TypeOf(Transform2D{})
	case "Vector4":
		return TypeVector4, reflect.TypeOf(Vector4{})
	case "Vector4i":
		return TypeVector4i, reflect.TypeOf(Vector4i{})
	case "Plane":
		return TypePlane, reflect.TypeOf(Plane{})
	case "Quaternion":
		return TypeQuaternion, reflect.TypeOf(Quaternion{})
	case "AABB":
		return TypeAabb, reflect.TypeOf(AABB{})
	case "Basis":
		return TypeBasis, reflect.TypeOf(Basis{})
	case "Transform3D":
		return TypeTransform3d, reflect.TypeOf(Transform3D{})
	case "Projection":
		return TypeProjection, reflect.TypeOf(Projection{})
	case "Color":
		return TypeColor, reflect.TypeOf(Color{})
	case "StringName":
		return TypeStringName, reflect.TypeOf(StringName{})
	case "NodePath":
		return TypeNodePath, reflect.TypeOf(NodePath{})
	case "RID":
		return TypeRid, reflect.TypeOf(RID(0))
	case "Object":
		return TypeObject, reflect.TypeOf(uintptr(0))
	case "Callable":
		return TypeCallable, reflect.TypeOf(Callable{})
	case "Signal":
		return TypeSignal, reflect.TypeOf(Signal{})
	case "Dictionary":
		return TypeDictionary, reflect.TypeOf(Dictionary{})
	case "Array":
		return TypeArray, reflect.TypeOf(Array{})
	case "PackedByteArray":
		return TypePackedByteArray, reflect.TypeOf(PackedByteArray{})
	case "PackedInt32Array":
		return TypePackedInt32Array, reflect.TypeOf(PackedInt32Array{})
	case "PackedInt64Array":
		return TypePackedInt64Array, reflect.TypeOf(PackedInt64Array{})
	case "PackedFloat32Array":
		return TypePackedFloat32Array, reflect.TypeOf(PackedFloat32Array{})
	case "PackedFloat64Array":
		return TypePackedFloat64Array, reflect.TypeOf(PackedFloat64Array{})
	case "PackedStringArray":
		return TypePackedStringArray, reflect.TypeOf(PackedStringArray{})
	case "PackedVector2Array":
		return TypePackedVector2Array, reflect.TypeOf(PackedVector2Array{})
	case "PackedVector3Array":
		return TypePackedVector3Array, reflect.TypeOf(PackedVector3Array{})
	case "PackedColorArray":
		return TypePackedColorArray, reflect.TypeOf(PackedColorArray{})
	default:
		panic("gdextension.variantTypeFromName: unknown type " + s)
	}
}

func operatoTypeFromName(name string) Operator {
	switch name {
	case "Equals":
		return Equal
	case "NotEqual":
		return NotEqual
	case "Less":
		return Less
	case "LessEqual":
		return LessEqual
	case "Greater":
		return Greater
	case "GreaterEqual":
		return GreaterEqual
	case "Add":
		return Add
	case "Subtract":
		return Subtract
	case "Multiply":
		return Multiply
	case "Divide":
		return Divide
	case "Negate":
		return Negate
	case "Module":
		return Module
	case "Power":
		return Power
	case "ShiftLeft":
		return ShiftLeft
	case "ShiftRight":
		return ShiftRight
	case "BitAnd":
		return BitAnd
	case "BitOr":
		return BitOr
	case "BitXor":
		return BitXor
	case "BitNegate":
		return BitNegate
	case "And":
		return LogicalAnd
	case "Or":
		return LogicalOr
	case "Xor":
		return LogicalXor
	case "Not":
		return LogicalNegate
	case "In":
		return In
	default:
		panic("gdextension.operatoTypeFromName: unknown operator " + name)
	}
}
