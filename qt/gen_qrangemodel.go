package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

// purego: unsupported shapes skipped (map returns, >15-arg / >9-arg-callback, exotic): SetRoleNames

var (
	addr_QRangeModel_chain uintptr
	addr_QRangeModel_tr uintptr
	addr_QRangeModel_index uintptr
	addr_QRangeModel_parent uintptr
	addr_QRangeModel_sibling uintptr
	addr_QRangeModel_rowCount uintptr
	addr_QRangeModel_columnCount uintptr
	addr_QRangeModel_insertColumns uintptr
	addr_QRangeModel_removeColumns uintptr
	addr_QRangeModel_moveColumns uintptr
	addr_QRangeModel_insertRows uintptr
	addr_QRangeModel_removeRows uintptr
	addr_QRangeModel_moveRows uintptr
	addr_QRangeModel_resetRoleNames uintptr
	addr_QRangeModel_hasChildren uintptr
	addr_QRangeModel_autoConnectPolicy uintptr
	addr_QRangeModel_setAutoConnectPolicy uintptr
	addr_QRangeModel_roleNamesChanged uintptr
	addr_QRangeModel_autoConnectPolicyChanged uintptr
	addr_QRangeModel_connect uintptr
	addr_QRangeModel_connect2 uintptr
	addr_QRangeModel_tr2 uintptr
	addr_QRangeModel_tr3 uintptr
	addr_QRangeModel_index2 uintptr
	addr_QRangeModel_rowCount2 uintptr
	addr_QRangeModel_columnCount2 uintptr
	addr_QRangeModel_insertColumns2 uintptr
	addr_QRangeModel_removeColumns2 uintptr
	addr_QRangeModel_insertRows2 uintptr
	addr_QRangeModel_removeRows2 uintptr
	addr_QRangeModel_hasChildren2 uintptr
	addr_QRangeModel_connect3 uintptr
	addr_QRangeModel_connect4 uintptr
	pg_QRangeModel_connect_roleNamesChanged func(uintptr, int64)
	pg_QRangeModel_connect_autoConnectPolicyChanged func(uintptr, int64)
	pg_QRangeModel_connect_headerDataChanged func(uintptr, int64)
	pg_QRangeModel_connect_layoutChanged func(uintptr, int64)
	pg_QRangeModel_connect_layoutChanged2 func(uintptr, int64)
	pg_QRangeModel_connect_layoutAboutToBeChanged func(uintptr, int64)
	pg_QRangeModel_connect_layoutAboutToBeChanged2 func(uintptr, int64)
	pg_QRangeModel_connect_rowsAboutToBeInserted func(uintptr, int64)
	pg_QRangeModel_connect_rowsInserted func(uintptr, int64)
	pg_QRangeModel_connect_rowsAboutToBeRemoved func(uintptr, int64)
	pg_QRangeModel_connect_rowsRemoved func(uintptr, int64)
	pg_QRangeModel_connect_columnsAboutToBeInserted func(uintptr, int64)
	pg_QRangeModel_connect_columnsInserted func(uintptr, int64)
	pg_QRangeModel_connect_columnsAboutToBeRemoved func(uintptr, int64)
	pg_QRangeModel_connect_columnsRemoved func(uintptr, int64)
	pg_QRangeModel_connect_modelAboutToBeReset func(uintptr, int64)
	pg_QRangeModel_connect_modelReset func(uintptr, int64)
	pg_QRangeModel_connect_rowsAboutToBeMoved func(uintptr, int64)
	pg_QRangeModel_connect_rowsMoved func(uintptr, int64)
	pg_QRangeModel_connect_columnsAboutToBeMoved func(uintptr, int64)
	pg_QRangeModel_connect_columnsMoved func(uintptr, int64)
	pg_QRangeModel_connect_destroyed func(uintptr, int64)
	pg_QRangeModel_connect_destroyed2 func(uintptr, int64)
	pg_QRangeModel_connect_objectNameChanged func(uintptr, int64)
	pg_QRangeModel__disconnect func(int64)
	pg_QRangeModel_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRangeModel_chain, "QRangeModel_chain")
	symaddr(&addr_QRangeModel_tr, "QRangeModel_tr")
	symaddr(&addr_QRangeModel_index, "QRangeModel_index")
	symaddr(&addr_QRangeModel_parent, "QRangeModel_parent")
	symaddr(&addr_QRangeModel_sibling, "QRangeModel_sibling")
	symaddr(&addr_QRangeModel_rowCount, "QRangeModel_rowCount")
	symaddr(&addr_QRangeModel_columnCount, "QRangeModel_columnCount")
	symaddr(&addr_QRangeModel_insertColumns, "QRangeModel_insertColumns")
	symaddr(&addr_QRangeModel_removeColumns, "QRangeModel_removeColumns")
	symaddr(&addr_QRangeModel_moveColumns, "QRangeModel_moveColumns")
	symaddr(&addr_QRangeModel_insertRows, "QRangeModel_insertRows")
	symaddr(&addr_QRangeModel_removeRows, "QRangeModel_removeRows")
	symaddr(&addr_QRangeModel_moveRows, "QRangeModel_moveRows")
	symaddr(&addr_QRangeModel_resetRoleNames, "QRangeModel_resetRoleNames")
	symaddr(&addr_QRangeModel_hasChildren, "QRangeModel_hasChildren")
	symaddr(&addr_QRangeModel_autoConnectPolicy, "QRangeModel_autoConnectPolicy")
	symaddr(&addr_QRangeModel_setAutoConnectPolicy, "QRangeModel_setAutoConnectPolicy")
	symaddr(&addr_QRangeModel_roleNamesChanged, "QRangeModel_roleNamesChanged")
	symaddr(&addr_QRangeModel_autoConnectPolicyChanged, "QRangeModel_autoConnectPolicyChanged")
	symaddr(&addr_QRangeModel_connect, "QRangeModel_connect")
	symaddr(&addr_QRangeModel_connect2, "QRangeModel_connect2")
	symaddr(&addr_QRangeModel_tr2, "QRangeModel_tr2")
	symaddr(&addr_QRangeModel_tr3, "QRangeModel_tr3")
	symaddr(&addr_QRangeModel_index2, "QRangeModel_index2")
	symaddr(&addr_QRangeModel_rowCount2, "QRangeModel_rowCount2")
	symaddr(&addr_QRangeModel_columnCount2, "QRangeModel_columnCount2")
	symaddr(&addr_QRangeModel_insertColumns2, "QRangeModel_insertColumns2")
	symaddr(&addr_QRangeModel_removeColumns2, "QRangeModel_removeColumns2")
	symaddr(&addr_QRangeModel_insertRows2, "QRangeModel_insertRows2")
	symaddr(&addr_QRangeModel_removeRows2, "QRangeModel_removeRows2")
	symaddr(&addr_QRangeModel_hasChildren2, "QRangeModel_hasChildren2")
	symaddr(&addr_QRangeModel_connect3, "QRangeModel_connect3")
	symaddr(&addr_QRangeModel_connect4, "QRangeModel_connect4")
	reg(&pg_QRangeModel_connect_roleNamesChanged, "QRangeModel_connect_roleNamesChanged")
	reg(&pg_QRangeModel_connect_autoConnectPolicyChanged, "QRangeModel_connect_autoConnectPolicyChanged")
	reg(&pg_QRangeModel_connect_headerDataChanged, "QRangeModel_connect_headerDataChanged")
	reg(&pg_QRangeModel_connect_layoutChanged, "QRangeModel_connect_layoutChanged")
	reg(&pg_QRangeModel_connect_layoutChanged2, "QRangeModel_connect_layoutChanged2")
	reg(&pg_QRangeModel_connect_layoutAboutToBeChanged, "QRangeModel_connect_layoutAboutToBeChanged")
	reg(&pg_QRangeModel_connect_layoutAboutToBeChanged2, "QRangeModel_connect_layoutAboutToBeChanged2")
	reg(&pg_QRangeModel_connect_rowsAboutToBeInserted, "QRangeModel_connect_rowsAboutToBeInserted")
	reg(&pg_QRangeModel_connect_rowsInserted, "QRangeModel_connect_rowsInserted")
	reg(&pg_QRangeModel_connect_rowsAboutToBeRemoved, "QRangeModel_connect_rowsAboutToBeRemoved")
	reg(&pg_QRangeModel_connect_rowsRemoved, "QRangeModel_connect_rowsRemoved")
	reg(&pg_QRangeModel_connect_columnsAboutToBeInserted, "QRangeModel_connect_columnsAboutToBeInserted")
	reg(&pg_QRangeModel_connect_columnsInserted, "QRangeModel_connect_columnsInserted")
	reg(&pg_QRangeModel_connect_columnsAboutToBeRemoved, "QRangeModel_connect_columnsAboutToBeRemoved")
	reg(&pg_QRangeModel_connect_columnsRemoved, "QRangeModel_connect_columnsRemoved")
	reg(&pg_QRangeModel_connect_modelAboutToBeReset, "QRangeModel_connect_modelAboutToBeReset")
	reg(&pg_QRangeModel_connect_modelReset, "QRangeModel_connect_modelReset")
	reg(&pg_QRangeModel_connect_rowsAboutToBeMoved, "QRangeModel_connect_rowsAboutToBeMoved")
	reg(&pg_QRangeModel_connect_rowsMoved, "QRangeModel_connect_rowsMoved")
	reg(&pg_QRangeModel_connect_columnsAboutToBeMoved, "QRangeModel_connect_columnsAboutToBeMoved")
	reg(&pg_QRangeModel_connect_columnsMoved, "QRangeModel_connect_columnsMoved")
	reg(&pg_QRangeModel_connect_destroyed, "QRangeModel_connect_destroyed")
	reg(&pg_QRangeModel_connect_destroyed2, "QRangeModel_connect_destroyed2")
	reg(&pg_QRangeModel_connect_objectNameChanged, "QRangeModel_connect_objectNameChanged")
	reg(&pg_QRangeModel__disconnect, "QRangeModel__disconnect")
	reg(&pg_QRangeModel_delete, "QRangeModel_delete")
}

type QRangeModel struct {
	h uintptr
	*QAbstractItemModel
}

func bindQRangeModel(h uintptr) *QRangeModel {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QRangeModel_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QRangeModel{h: h, QAbstractItemModel: &QAbstractItemModel{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQRangeModel(h uintptr) *QRangeModel { return bindQRangeModel(h) }

func (this *QRangeModel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRangeModel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRangeModel(h unsafe.Pointer) *QRangeModel { return bindQRangeModel(uintptr(h)) }
func UnsafeNewQRangeModelOwned(h unsafe.Pointer) *QRangeModel { return newownedQRangeModel(uintptr(h)) }

func (this *QRangeModel) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemModel != nil {
		this.QAbstractItemModel.UnsafeRelease()
	}
}

func QRangeModel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRangeModel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QRangeModel) Index(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QRangeModel_index, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QRangeModel) Parent(child *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QRangeModel_parent, this.h, uintptr(child.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return newownedQModelIndex(__r)
}

func (this *QRangeModel) Sibling(row int, column int, index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QRangeModel_sibling, this.h, uintptr(row), uintptr(column), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QRangeModel) RowCount() int {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_rowCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRangeModel) ColumnCount() int {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_columnCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRangeModel) InsertColumns(column int, count int) bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_insertColumns, this.h, uintptr(column), uintptr(count))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRangeModel) RemoveColumns(column int, count int) bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_removeColumns, this.h, uintptr(column), uintptr(count))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRangeModel) MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destParent *QModelIndex, destColumn int) bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_moveColumns, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(count), uintptr(destParent.UnsafePointer()), uintptr(destColumn))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destParent)
	return __r != 0
}

func (this *QRangeModel) InsertRows(row int, count int) bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_insertRows, this.h, uintptr(row), uintptr(count))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRangeModel) RemoveRows(row int, count int) bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_removeRows, this.h, uintptr(row), uintptr(count))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRangeModel) MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destParent *QModelIndex, destRow int) bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_moveRows, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(count), uintptr(destParent.UnsafePointer()), uintptr(destRow))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destParent)
	return __r != 0
}

func (this *QRangeModel) ResetRoleNames() {
	purego.SyscallN(addr_QRangeModel_resetRoleNames, this.h)
	runtime.KeepAlive(this)
}

func (this *QRangeModel) HasChildren() bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_hasChildren, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRangeModel) AutoConnectPolicy() QRangeModel__AutoConnectPolicy {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_autoConnectPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QRangeModel__AutoConnectPolicy(__r)
}

func (this *QRangeModel) SetAutoConnectPolicy(policy QRangeModel__AutoConnectPolicy) {
	purego.SyscallN(addr_QRangeModel_setAutoConnectPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QRangeModel) RoleNamesChanged() {
	purego.SyscallN(addr_QRangeModel_roleNamesChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QRangeModel) AutoConnectPolicyChanged(policy QRangeModel__AutoConnectPolicy) {
	purego.SyscallN(addr_QRangeModel_autoConnectPolicyChanged, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func QRangeModel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QRangeModel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QRangeModel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QRangeModel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QRangeModel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRangeModel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QRangeModel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRangeModel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QRangeModel) Index2(row int, column int, parent *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QRangeModel_index2, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return newownedQModelIndex(__r)
}

func (this *QRangeModel) RowCountWithParent(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_rowCount2, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QRangeModel) ColumnCountWithParent(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_columnCount2, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QRangeModel) InsertColumns2(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_insertColumns2, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QRangeModel) RemoveColumns2(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_removeColumns2, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QRangeModel) InsertRows2(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_insertRows2, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QRangeModel) RemoveRows2(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_removeRows2, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QRangeModel) HasChildrenWithParent(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QRangeModel_hasChildren2, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func QRangeModel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QRangeModel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QRangeModel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QRangeModel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QRangeModel) OnRoleNamesChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRangeModel_connect_roleNamesChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnAutoConnectPolicyChanged(fn func(QRangeModel__AutoConnectPolicy)) int64 {
	id := registerCallback(func(policy int32) { fn(QRangeModel__AutoConnectPolicy(policy)) })
	pg_QRangeModel_connect_autoConnectPolicyChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnHeaderDataChanged(fn func(Orientation, int, int)) int64 {
	id := registerCallback(func(orientation int32, first int32, last int32) { fn(Orientation(orientation), int(first), int(last)) })
	pg_QRangeModel_connect_headerDataChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRangeModel_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnLayoutChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QRangeModel_connect_layoutChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnLayoutAboutToBeChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRangeModel_connect_layoutAboutToBeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnLayoutAboutToBeChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QRangeModel_connect_layoutAboutToBeChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnRowsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QRangeModel_connect_rowsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnRowsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QRangeModel_connect_rowsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QRangeModel_connect_rowsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnRowsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QRangeModel_connect_rowsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnColumnsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QRangeModel_connect_columnsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnColumnsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QRangeModel_connect_columnsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnColumnsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QRangeModel_connect_columnsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnColumnsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QRangeModel_connect_columnsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnModelAboutToBeReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRangeModel_connect_modelAboutToBeReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnModelReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRangeModel_connect_modelReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnRowsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QRangeModel_connect_rowsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnRowsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QRangeModel_connect_rowsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnColumnsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QRangeModel_connect_columnsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnColumnsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QRangeModel_connect_columnsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRangeModel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QRangeModel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QRangeModel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRangeModel) Disconnect(handle int64) {
	pg_QRangeModel__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QRangeModel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRangeModel_delete(this.h)
	}
	this.UnsafeRelease()
}
