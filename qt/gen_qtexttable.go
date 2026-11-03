package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextTable_chain uintptr
	addr_QTextTable_new uintptr
	addr_QTextTable_tr uintptr
	addr_QTextTable_resize uintptr
	addr_QTextTable_insertRows uintptr
	addr_QTextTable_insertColumns uintptr
	addr_QTextTable_appendRows uintptr
	addr_QTextTable_appendColumns uintptr
	addr_QTextTable_removeRows uintptr
	addr_QTextTable_removeColumns uintptr
	addr_QTextTable_mergeCells uintptr
	addr_QTextTable_mergeCells2 uintptr
	addr_QTextTable_splitCell uintptr
	addr_QTextTable_rows uintptr
	addr_QTextTable_columns uintptr
	addr_QTextTable_cellAt uintptr
	addr_QTextTable_cellAt2 uintptr
	addr_QTextTable_cellAt3 uintptr
	addr_QTextTable_rowStart uintptr
	addr_QTextTable_rowEnd uintptr
	addr_QTextTable_setFormat uintptr
	addr_QTextTable_format uintptr
	addr_QTextTable_connect uintptr
	addr_QTextTable_connect2 uintptr
	addr_QTextTable_tr2 uintptr
	addr_QTextTable_tr3 uintptr
	addr_QTextTable_connect3 uintptr
	addr_QTextTable_connect4 uintptr
	pg_QTextTable_connect_destroyed func(uintptr, int64)
	pg_QTextTable_connect_destroyed2 func(uintptr, int64)
	pg_QTextTable_connect_objectNameChanged func(uintptr, int64)
	pg_QTextTable__disconnect func(int64)
	pg_QTextTable_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextTable_chain, "QTextTable_chain")
	symaddr(&addr_QTextTable_new, "QTextTable_new")
	symaddr(&addr_QTextTable_tr, "QTextTable_tr")
	symaddr(&addr_QTextTable_resize, "QTextTable_resize")
	symaddr(&addr_QTextTable_insertRows, "QTextTable_insertRows")
	symaddr(&addr_QTextTable_insertColumns, "QTextTable_insertColumns")
	symaddr(&addr_QTextTable_appendRows, "QTextTable_appendRows")
	symaddr(&addr_QTextTable_appendColumns, "QTextTable_appendColumns")
	symaddr(&addr_QTextTable_removeRows, "QTextTable_removeRows")
	symaddr(&addr_QTextTable_removeColumns, "QTextTable_removeColumns")
	symaddr(&addr_QTextTable_mergeCells, "QTextTable_mergeCells")
	symaddr(&addr_QTextTable_mergeCells2, "QTextTable_mergeCells2")
	symaddr(&addr_QTextTable_splitCell, "QTextTable_splitCell")
	symaddr(&addr_QTextTable_rows, "QTextTable_rows")
	symaddr(&addr_QTextTable_columns, "QTextTable_columns")
	symaddr(&addr_QTextTable_cellAt, "QTextTable_cellAt")
	symaddr(&addr_QTextTable_cellAt2, "QTextTable_cellAt2")
	symaddr(&addr_QTextTable_cellAt3, "QTextTable_cellAt3")
	symaddr(&addr_QTextTable_rowStart, "QTextTable_rowStart")
	symaddr(&addr_QTextTable_rowEnd, "QTextTable_rowEnd")
	symaddr(&addr_QTextTable_setFormat, "QTextTable_setFormat")
	symaddr(&addr_QTextTable_format, "QTextTable_format")
	symaddr(&addr_QTextTable_connect, "QTextTable_connect")
	symaddr(&addr_QTextTable_connect2, "QTextTable_connect2")
	symaddr(&addr_QTextTable_tr2, "QTextTable_tr2")
	symaddr(&addr_QTextTable_tr3, "QTextTable_tr3")
	symaddr(&addr_QTextTable_connect3, "QTextTable_connect3")
	symaddr(&addr_QTextTable_connect4, "QTextTable_connect4")
	reg(&pg_QTextTable_connect_destroyed, "QTextTable_connect_destroyed")
	reg(&pg_QTextTable_connect_destroyed2, "QTextTable_connect_destroyed2")
	reg(&pg_QTextTable_connect_objectNameChanged, "QTextTable_connect_objectNameChanged")
	reg(&pg_QTextTable__disconnect, "QTextTable__disconnect")
	reg(&pg_QTextTable_delete, "QTextTable_delete")
}

type QTextTable struct {
	h uintptr
	*QTextFrame
}

func bindQTextTable(h uintptr) *QTextTable {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QTextTable_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTextTable{h: h, QTextFrame: &QTextFrame{h: __c[0], QTextObject: &QTextObject{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQTextTable(h uintptr) *QTextTable { return bindQTextTable(h) }

func (this *QTextTable) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextTable) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextTable(h unsafe.Pointer) *QTextTable { return bindQTextTable(uintptr(h)) }
func UnsafeNewQTextTableOwned(h unsafe.Pointer) *QTextTable { return newownedQTextTable(uintptr(h)) }

func (this *QTextTable) UnsafeRelease() {
	this.h = 0
	if this.QTextFrame != nil {
		this.QTextFrame.UnsafeRelease()
	}
}

func NewQTextTable(doc *QTextDocument) *QTextTable {
	__r, _, _ := purego.SyscallN(addr_QTextTable_new, uintptr(doc.UnsafePointer()))
	__o := newownedQTextTable(__r)
	runtime.KeepAlive(doc)
	return __o
}

func QTextTable_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextTable_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTextTable) Resize(rows int, cols int) {
	purego.SyscallN(addr_QTextTable_resize, this.h, uintptr(rows), uintptr(cols))
	runtime.KeepAlive(this)
}

func (this *QTextTable) InsertRows(pos int, num int) {
	purego.SyscallN(addr_QTextTable_insertRows, this.h, uintptr(pos), uintptr(num))
	runtime.KeepAlive(this)
}

func (this *QTextTable) InsertColumns(pos int, num int) {
	purego.SyscallN(addr_QTextTable_insertColumns, this.h, uintptr(pos), uintptr(num))
	runtime.KeepAlive(this)
}

func (this *QTextTable) AppendRows(count int) {
	purego.SyscallN(addr_QTextTable_appendRows, this.h, uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QTextTable) AppendColumns(count int) {
	purego.SyscallN(addr_QTextTable_appendColumns, this.h, uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QTextTable) RemoveRows(pos int, num int) {
	purego.SyscallN(addr_QTextTable_removeRows, this.h, uintptr(pos), uintptr(num))
	runtime.KeepAlive(this)
}

func (this *QTextTable) RemoveColumns(pos int, num int) {
	purego.SyscallN(addr_QTextTable_removeColumns, this.h, uintptr(pos), uintptr(num))
	runtime.KeepAlive(this)
}

func (this *QTextTable) MergeCells(row int, col int, numRows int, numCols int) {
	purego.SyscallN(addr_QTextTable_mergeCells, this.h, uintptr(row), uintptr(col), uintptr(numRows), uintptr(numCols))
	runtime.KeepAlive(this)
}

func (this *QTextTable) MergeCellsWithCursor(cursor *QTextCursor) {
	purego.SyscallN(addr_QTextTable_mergeCells2, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QTextTable) SplitCell(row int, col int, numRows int, numCols int) {
	purego.SyscallN(addr_QTextTable_splitCell, this.h, uintptr(row), uintptr(col), uintptr(numRows), uintptr(numCols))
	runtime.KeepAlive(this)
}

func (this *QTextTable) Rows() int {
	__r0, _, _ := purego.SyscallN(addr_QTextTable_rows, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextTable) Columns() int {
	__r0, _, _ := purego.SyscallN(addr_QTextTable_columns, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextTable) CellAt(row int, col int) *QTextTableCell {
	__r, _, _ := purego.SyscallN(addr_QTextTable_cellAt, this.h, uintptr(row), uintptr(col))
	runtime.KeepAlive(this)
	return newownedQTextTableCell(__r)
}

func (this *QTextTable) CellAtWithPosition(position int) *QTextTableCell {
	__r, _, _ := purego.SyscallN(addr_QTextTable_cellAt2, this.h, uintptr(position))
	runtime.KeepAlive(this)
	return newownedQTextTableCell(__r)
}

func (this *QTextTable) CellAtWithQTextCursor(c *QTextCursor) *QTextTableCell {
	__r, _, _ := purego.SyscallN(addr_QTextTable_cellAt3, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return newownedQTextTableCell(__r)
}

func (this *QTextTable) RowStart(c *QTextCursor) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextTable_rowStart, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return newownedQTextCursor(__r)
}

func (this *QTextTable) RowEnd(c *QTextCursor) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextTable_rowEnd, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return newownedQTextCursor(__r)
}

func (this *QTextTable) SetFormat(format *QTextTableFormat) {
	purego.SyscallN(addr_QTextTable_setFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextTable) Format() *QTextTableFormat {
	__r, _, _ := purego.SyscallN(addr_QTextTable_format, this.h)
	runtime.KeepAlive(this)
	return newownedQTextTableFormat(__r)
}

func QTextTable_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextTable_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextTable_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextTable_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QTextTable_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextTable_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextTable_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextTable_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextTable_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextTable_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextTable_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextTable_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTextTable) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextTable_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextTable) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTextTable_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextTable) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTextTable_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextTable) Disconnect(handle int64) {
	pg_QTextTable__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QTextTable) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextTable_delete(this.h)
	}
	this.UnsafeRelease()
}
