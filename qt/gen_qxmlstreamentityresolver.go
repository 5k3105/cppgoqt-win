package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QXmlStreamEntityResolver_new uintptr
	addr_QXmlStreamEntityResolver_resolveEntity uintptr
	addr_QXmlStreamEntityResolver_resolveUndeclaredEntity uintptr
	addr_QXmlStreamEntityResolver_resolveEntity_default uintptr
	addr_QXmlStreamEntityResolver_resolveUndeclaredEntity_default uintptr
	pg_QXmlStreamEntityResolver__set_override func(uintptr, int32)
	pg_QXmlStreamEntityResolver_delete func(uintptr)
)

func init() {
	symaddr(&addr_QXmlStreamEntityResolver_new, "QXmlStreamEntityResolver_new")
	symaddr(&addr_QXmlStreamEntityResolver_resolveEntity, "QXmlStreamEntityResolver_resolveEntity")
	symaddr(&addr_QXmlStreamEntityResolver_resolveUndeclaredEntity, "QXmlStreamEntityResolver_resolveUndeclaredEntity")
	symaddr(&addr_QXmlStreamEntityResolver_resolveEntity_default, "QXmlStreamEntityResolver_resolveEntity_default")
	symaddr(&addr_QXmlStreamEntityResolver_resolveUndeclaredEntity_default, "QXmlStreamEntityResolver_resolveUndeclaredEntity_default")
	reg(&pg_QXmlStreamEntityResolver__set_override, "QXmlStreamEntityResolver__set_override")
	reg(&pg_QXmlStreamEntityResolver_delete, "QXmlStreamEntityResolver_delete")
}

type QXmlStreamEntityResolver struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QXmlStreamEntityResolver) UnsafeGoID() int64     { return this.vid }
func (this *QXmlStreamEntityResolver) UnsafeSetGoID(id int64) { this.vid = id }

func bindQXmlStreamEntityResolver(h uintptr) *QXmlStreamEntityResolver {
	if h == 0 {
		return nil
	}
	return &QXmlStreamEntityResolver{h: h}
}

func newownedQXmlStreamEntityResolver(h uintptr) *QXmlStreamEntityResolver {
	o := bindQXmlStreamEntityResolver(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QXmlStreamEntityResolver).Delete)
	}
	return o
}

func (this *QXmlStreamEntityResolver) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QXmlStreamEntityResolver) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQXmlStreamEntityResolver(h unsafe.Pointer) *QXmlStreamEntityResolver { return bindQXmlStreamEntityResolver(uintptr(h)) }
func UnsafeNewQXmlStreamEntityResolverOwned(h unsafe.Pointer) *QXmlStreamEntityResolver { return newownedQXmlStreamEntityResolver(uintptr(h)) }

func (this *QXmlStreamEntityResolver) UnsafeRelease() { this.h = 0 }

func NewQXmlStreamEntityResolver() *QXmlStreamEntityResolver {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityResolver_new, uintptr(__goid))
	__o := newownedQXmlStreamEntityResolver(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QXmlStreamEntityResolver_virtIdx, QXmlStreamEntityResolver_enable)
	return __o
}

func (this *QXmlStreamEntityResolver) ResolveEntity(publicId string, systemId string) string {
	var publicId_p uintptr
	if len(publicId) > 0 {
		publicId_p = uintptr(unsafe.Pointer(unsafe.StringData(publicId)))
	}
	var systemId_p uintptr
	if len(systemId) > 0 {
		systemId_p = uintptr(unsafe.Pointer(unsafe.StringData(systemId)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityResolver_resolveEntity, this.h, publicId_p, uintptr(len(publicId)), systemId_p, uintptr(len(systemId)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(publicId)
	runtime.KeepAlive(systemId)
	return pgStr(__r, __n)
}

func (this *QXmlStreamEntityResolver) ResolveUndeclaredEntity(name string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityResolver_resolveUndeclaredEntity, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return pgStr(__r, __n)
}

func (this *QXmlStreamEntityResolver) ResolveEntityDefault(publicId string, systemId string) string {
	var publicId_p uintptr
	if len(publicId) > 0 {
		publicId_p = uintptr(unsafe.Pointer(unsafe.StringData(publicId)))
	}
	var systemId_p uintptr
	if len(systemId) > 0 {
		systemId_p = uintptr(unsafe.Pointer(unsafe.StringData(systemId)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityResolver_resolveEntity_default, this.h, publicId_p, uintptr(len(publicId)), systemId_p, uintptr(len(systemId)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(publicId)
	runtime.KeepAlive(systemId)
	return pgStr(__r, __n)
}

func (this *QXmlStreamEntityResolver) ResolveUndeclaredEntityDefault(name string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityResolver_resolveUndeclaredEntity_default, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return pgStr(__r, __n)
}

func QXmlStreamEntityResolver_virtIdx(slot int64) int32 {
	switch slot {
	case -1345840861565463748:
		return 0
	case -5256568367070285975:
		return 1
	}
	return -1
}

func QXmlStreamEntityResolver_enable(h uintptr, idx int32) { pg_QXmlStreamEntityResolver__set_override(h, idx) }

func (this *QXmlStreamEntityResolver) OnResolveEntity(fn func(string, string) string) {
	registerVirt(this.UnsafeGoID(), -1345840861565463748, func(publicId uintptr, publicId_len int64, systemId uintptr, systemId_len int64) (uintptr, int64) {
		__s := fn(pgStrNoFree(publicId, publicId_len), pgStrNoFree(systemId, systemId_len))
		return pgCString(__s), int64(len(__s))
	}) // slot "resolveEntity"
}

func (this *QXmlStreamEntityResolver) OnResolveUndeclaredEntity(fn func(string) string) {
	registerVirt(this.UnsafeGoID(), -5256568367070285975, func(name uintptr, name_len int64) (uintptr, int64) {
		__s := fn(pgStrNoFree(name, name_len))
		return pgCString(__s), int64(len(__s))
	}) // slot "resolveUndeclaredEntity"
}

func (this *QXmlStreamEntityResolver) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QXmlStreamEntityResolver_delete(this.h)
	}
	this.UnsafeRelease()
}
