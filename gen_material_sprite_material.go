package three
// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// material_method_generator -materialName SpriteMaterial -materialSlug sprite_material

import (
	"github.com/fatih/structs"
	"github.com/gopherjs/gopherjs/js"
)
	
func (m SpriteMaterial) OnBeforeCompile() {
	m.Call("onBeforeCompile")
}

func (m SpriteMaterial) SetValues(values MaterialParameters) {
	m.Call("setValues", structs.Map(values))
}

func (m SpriteMaterial) ToJSON(meta interface{}) interface{} {
	return m.Call("toJSON", meta)
}

func (m SpriteMaterial) Clone() {
	m.Call("clone")
}

func (m SpriteMaterial) Copy(source Object3D) {
	m.Call("copy", source)
}

func (m SpriteMaterial) Dispose() {
	m.Call("dispose")
}

func (m SpriteMaterial) getInternalObject() *js.Object {
	return m.Object
}

