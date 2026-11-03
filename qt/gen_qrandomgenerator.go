package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRandomGenerator_new uintptr
	addr_QRandomGenerator_new2 uintptr
	addr_QRandomGenerator_new3 uintptr
	addr_QRandomGenerator_OperatorAssign uintptr
	addr_QRandomGenerator_generate uintptr
	addr_QRandomGenerator_generate64 uintptr
	pg_QRandomGenerator_generateDouble func(uintptr) float64
	pg_QRandomGenerator_bounded func(uintptr, float64) float64
	addr_QRandomGenerator_bounded2 uintptr
	addr_QRandomGenerator_bounded3 uintptr
	addr_QRandomGenerator_bounded4 uintptr
	addr_QRandomGenerator_bounded5 uintptr
	addr_QRandomGenerator_bounded6 uintptr
	addr_QRandomGenerator_bounded7 uintptr
	addr_QRandomGenerator_bounded8 uintptr
	addr_QRandomGenerator_bounded9 uintptr
	addr_QRandomGenerator_bounded10 uintptr
	addr_QRandomGenerator_bounded11 uintptr
	addr_QRandomGenerator_bounded12 uintptr
	addr_QRandomGenerator_bounded13 uintptr
	addr_QRandomGenerator_generate2 uintptr
	addr_QRandomGenerator_OperatorCall uintptr
	addr_QRandomGenerator_seed uintptr
	addr_QRandomGenerator_discard uintptr
	addr_QRandomGenerator_min uintptr
	addr_QRandomGenerator_max uintptr
	addr_QRandomGenerator_system uintptr
	addr_QRandomGenerator_global uintptr
	addr_QRandomGenerator_securelySeeded uintptr
	addr_QRandomGenerator_seed2 uintptr
	addr_QRandomGenerator_OperatorEqual uintptr
	addr_QRandomGenerator_OperatorNotEqual uintptr
	pg_QRandomGenerator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRandomGenerator_new, "QRandomGenerator_new")
	symaddr(&addr_QRandomGenerator_new2, "QRandomGenerator_new2")
	symaddr(&addr_QRandomGenerator_new3, "QRandomGenerator_new3")
	symaddr(&addr_QRandomGenerator_OperatorAssign, "QRandomGenerator_OperatorAssign")
	symaddr(&addr_QRandomGenerator_generate, "QRandomGenerator_generate")
	symaddr(&addr_QRandomGenerator_generate64, "QRandomGenerator_generate64")
	reg(&pg_QRandomGenerator_generateDouble, "QRandomGenerator_generateDouble")
	reg(&pg_QRandomGenerator_bounded, "QRandomGenerator_bounded")
	symaddr(&addr_QRandomGenerator_bounded2, "QRandomGenerator_bounded2")
	symaddr(&addr_QRandomGenerator_bounded3, "QRandomGenerator_bounded3")
	symaddr(&addr_QRandomGenerator_bounded4, "QRandomGenerator_bounded4")
	symaddr(&addr_QRandomGenerator_bounded5, "QRandomGenerator_bounded5")
	symaddr(&addr_QRandomGenerator_bounded6, "QRandomGenerator_bounded6")
	symaddr(&addr_QRandomGenerator_bounded7, "QRandomGenerator_bounded7")
	symaddr(&addr_QRandomGenerator_bounded8, "QRandomGenerator_bounded8")
	symaddr(&addr_QRandomGenerator_bounded9, "QRandomGenerator_bounded9")
	symaddr(&addr_QRandomGenerator_bounded10, "QRandomGenerator_bounded10")
	symaddr(&addr_QRandomGenerator_bounded11, "QRandomGenerator_bounded11")
	symaddr(&addr_QRandomGenerator_bounded12, "QRandomGenerator_bounded12")
	symaddr(&addr_QRandomGenerator_bounded13, "QRandomGenerator_bounded13")
	symaddr(&addr_QRandomGenerator_generate2, "QRandomGenerator_generate2")
	symaddr(&addr_QRandomGenerator_OperatorCall, "QRandomGenerator_OperatorCall")
	symaddr(&addr_QRandomGenerator_seed, "QRandomGenerator_seed")
	symaddr(&addr_QRandomGenerator_discard, "QRandomGenerator_discard")
	symaddr(&addr_QRandomGenerator_min, "QRandomGenerator_min")
	symaddr(&addr_QRandomGenerator_max, "QRandomGenerator_max")
	symaddr(&addr_QRandomGenerator_system, "QRandomGenerator_system")
	symaddr(&addr_QRandomGenerator_global, "QRandomGenerator_global")
	symaddr(&addr_QRandomGenerator_securelySeeded, "QRandomGenerator_securelySeeded")
	symaddr(&addr_QRandomGenerator_seed2, "QRandomGenerator_seed2")
	symaddr(&addr_QRandomGenerator_OperatorEqual, "QRandomGenerator_OperatorEqual")
	symaddr(&addr_QRandomGenerator_OperatorNotEqual, "QRandomGenerator_OperatorNotEqual")
	reg(&pg_QRandomGenerator_delete, "QRandomGenerator_delete")
}

type QRandomGenerator struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QRandomGenerator) UnsafeGoID() int64     { return this.vid }
func (this *QRandomGenerator) UnsafeSetGoID(id int64) { this.vid = id }

func bindQRandomGenerator(h uintptr) *QRandomGenerator {
	if h == 0 {
		return nil
	}
	return &QRandomGenerator{h: h}
}

func newownedQRandomGenerator(h uintptr) *QRandomGenerator {
	o := bindQRandomGenerator(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRandomGenerator).Delete)
	}
	return o
}

func (this *QRandomGenerator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRandomGenerator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRandomGenerator(h unsafe.Pointer) *QRandomGenerator { return bindQRandomGenerator(uintptr(h)) }
func UnsafeNewQRandomGeneratorOwned(h unsafe.Pointer) *QRandomGenerator { return newownedQRandomGenerator(uintptr(h)) }

func (this *QRandomGenerator) UnsafeRelease() { this.h = 0 }

func NewQRandomGenerator() *QRandomGenerator {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator_new)
	__o := newownedQRandomGenerator(__r)
	return __o
}

func NewQRandomGenerator2(other *QRandomGenerator) *QRandomGenerator {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator_new2, uintptr(other.UnsafePointer()))
	__o := newownedQRandomGenerator(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQRandomGenerator3(seedValue uint) *QRandomGenerator {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator_new3, uintptr(seedValue))
	__o := newownedQRandomGenerator(__r)
	return __o
}

func (this *QRandomGenerator) OperatorAssign(other *QRandomGenerator) *QRandomGenerator {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQRandomGenerator(__r)
}

func (this *QRandomGenerator) Generate() uint {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_generate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QRandomGenerator) Generate64() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_generate64, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QRandomGenerator) GenerateDouble() float64 {
	__r := pg_QRandomGenerator_generateDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRandomGenerator) Bounded(highest float64) float64 {
	__r := pg_QRandomGenerator_bounded(this.h, highest)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRandomGenerator) BoundedWithHighest(highest uint) uint {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded2, this.h, uintptr(highest))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QRandomGenerator) Bounded2(lowest uint, highest uint) uint {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded3, this.h, uintptr(lowest), uintptr(highest))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QRandomGenerator) Bounded3(highest int) int {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded4, this.h, uintptr(highest))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRandomGenerator) Bounded4(lowest int, highest int) int {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded5, this.h, uintptr(lowest), uintptr(highest))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRandomGenerator) Bounded5(highest uint64) uint64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded6, this.h, uintptr(highest))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QRandomGenerator) Bounded6(lowest uint64, highest uint64) uint64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded7, this.h, uintptr(lowest), uintptr(highest))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QRandomGenerator) Bounded7(highest int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded8, this.h, uintptr(highest))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRandomGenerator) Bounded8(lowest int64, highest int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded9, this.h, uintptr(lowest), uintptr(highest))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRandomGenerator) Bounded9(lowest int, highest int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded10, this.h, uintptr(lowest), uintptr(highest))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRandomGenerator) Bounded10(lowest int64, highest int) int64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded11, this.h, uintptr(lowest), uintptr(highest))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRandomGenerator) Bounded11(lowest uint, highest uint64) uint64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded12, this.h, uintptr(lowest), uintptr(highest))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QRandomGenerator) Bounded12(lowest uint64, highest uint) uint64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_bounded13, this.h, uintptr(lowest), uintptr(highest))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QRandomGenerator) Generate2() (uint, uint) {
	var begin_out int32
	var end_out int32
	purego.SyscallN(addr_QRandomGenerator_generate2, this.h, uintptr(unsafe.Pointer(&begin_out)), uintptr(unsafe.Pointer(&end_out)))
	runtime.KeepAlive(this)
	return uint(begin_out), uint(end_out)
}

func (this *QRandomGenerator) OperatorCall() uint {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_OperatorCall, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QRandomGenerator) Seed() {
	purego.SyscallN(addr_QRandomGenerator_seed, this.h)
	runtime.KeepAlive(this)
}

func (this *QRandomGenerator) Discard(z uint64) {
	purego.SyscallN(addr_QRandomGenerator_discard, this.h, uintptr(z))
	runtime.KeepAlive(this)
}

func QRandomGenerator_Min() uint {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_min)
	__r := int32(uint32(__r0))
	return uint(__r)
}

func QRandomGenerator_Max() uint {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_max)
	__r := int32(uint32(__r0))
	return uint(__r)
}

func QRandomGenerator_System() *QRandomGenerator {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator_system)
	return bindQRandomGenerator(__r)
}

func QRandomGenerator_Global() *QRandomGenerator {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator_global)
	return bindQRandomGenerator(__r)
}

func QRandomGenerator_SecurelySeeded() *QRandomGenerator {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator_securelySeeded)
	return newownedQRandomGenerator(__r)
}

func (this *QRandomGenerator) SeedWith(s uint) {
	purego.SyscallN(addr_QRandomGenerator_seed2, this.h, uintptr(s))
	runtime.KeepAlive(this)
}

func (this *QRandomGenerator) OperatorEqual(rng2 *QRandomGenerator) bool {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_OperatorEqual, this.h, uintptr(rng2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rng2)
	return __r != 0
}

func (this *QRandomGenerator) OperatorNotEqual(rng2 *QRandomGenerator) bool {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator_OperatorNotEqual, this.h, uintptr(rng2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rng2)
	return __r != 0
}

func (this *QRandomGenerator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRandomGenerator_delete(this.h)
	}
	this.UnsafeRelease()
}
