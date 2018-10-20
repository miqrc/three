package three
// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// object3d_method_generator -typeName Mesh -typeSlug mesh

import "github.com/gopherjs/gopherjs/js"

func (obj *Mesh) ApplyMatrix(matrix *Matrix4) {
	obj.Call("applyMatrix", matrix)
}

func (obj *Mesh) Add(m Object3D) {
	obj.Object.Call("add", m)
}

func (obj *Mesh) Remove(m *js.Object) {
	obj.Object.Call("remove", m)
}

func (obj *Mesh) GetObjectById(id int) *js.Object {
	return obj.Call("getObjectById", id)
}

// func (obj *Mesh) Copy() *Mesh {
// 	return &Mesh{Object: obj.getInternalObject().Call("copy")}
// }

func (obj *Mesh) ToJSON() interface{} {
	return obj.Object.Call("toJSON").Interface()
}

func (obj *Mesh) getInternalObject() *js.Object {
	return obj.Object
}

func (obj *Mesh) UpdateMatrix() {
	obj.Call("updateMatrix")
}

