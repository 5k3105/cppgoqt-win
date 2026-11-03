#include <QtCore/qbytearrayview.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qchar.h>
#include <QtCore/qstringview.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qstring.h>
#include "gen_qbytearrayview.h"

extern "C" {

void* QByteArrayView_new() {
	return new QByteArrayView();
}
void* QByteArrayView_new2(void* arg1) {
	return new QByteArrayView(*(QByteArrayView*)arg1);
}
void* QByteArrayView_toByteArray(void* self) {
	auto&& __r = ((QByteArrayView*)self)->toByteArray();
	return new QByteArray(__r);
}
long long QByteArrayView_size(void* self) {
	auto&& __r = ((QByteArrayView*)self)->size();
	return __r;
}
const char* QByteArrayView_data(void* self) {
	auto&& __r = ((QByteArrayView*)self)->data();
	return __r;
}
const char* QByteArrayView_constData(void* self) {
	auto&& __r = ((QByteArrayView*)self)->constData();
	return __r;
}
char QByteArrayView_OperatorSubscript(void* self, long long n) {
	auto&& __r = ((QByteArrayView*)self)->operator[](n);
	return __r;
}
char QByteArrayView_at(void* self, long long n) {
	auto&& __r = ((QByteArrayView*)self)->at(n);
	return __r;
}
void* QByteArrayView_first(void* self, long long n) {
	auto&& __r = ((QByteArrayView*)self)->first(n);
	return new QByteArrayView(__r);
}
void* QByteArrayView_last(void* self, long long n) {
	auto&& __r = ((QByteArrayView*)self)->last(n);
	return new QByteArrayView(__r);
}
void* QByteArrayView_sliced(void* self, long long pos) {
	auto&& __r = ((QByteArrayView*)self)->sliced(pos);
	return new QByteArrayView(__r);
}
void* QByteArrayView_sliced2(void* self, long long pos, long long n) {
	auto&& __r = ((QByteArrayView*)self)->sliced(pos, n);
	return new QByteArrayView(__r);
}
void* QByteArrayView_slice(void* self, long long pos) {
	auto&& __r = ((QByteArrayView*)self)->slice(pos);
	return (void*)&(__r);
}
void* QByteArrayView_slice2(void* self, long long pos, long long n) {
	auto&& __r = ((QByteArrayView*)self)->slice(pos, n);
	return (void*)&(__r);
}
void* QByteArrayView_chopped(void* self, long long len_) {
	auto&& __r = ((QByteArrayView*)self)->chopped(len_);
	return new QByteArrayView(__r);
}
void* QByteArrayView_left(void* self, long long n) {
	auto&& __r = ((QByteArrayView*)self)->left(n);
	return new QByteArrayView(__r);
}
void* QByteArrayView_right(void* self, long long n) {
	auto&& __r = ((QByteArrayView*)self)->right(n);
	return new QByteArrayView(__r);
}
void* QByteArrayView_mid(void* self, long long pos) {
	auto&& __r = ((QByteArrayView*)self)->mid(pos);
	return new QByteArrayView(__r);
}
void QByteArrayView_truncate(void* self, long long n) {
	((QByteArrayView*)self)->truncate(n);
}
void QByteArrayView_chop(void* self, long long n) {
	((QByteArrayView*)self)->chop(n);
}
void* QByteArrayView_trimmed(void* self) {
	auto&& __r = ((QByteArrayView*)self)->trimmed();
	return new QByteArrayView(__r);
}
short QByteArrayView_toShort(void* self) {
	auto&& __r = ((QByteArrayView*)self)->toShort();
	return __r;
}
unsigned short QByteArrayView_toUShort(void* self) {
	auto&& __r = ((QByteArrayView*)self)->toUShort();
	return __r;
}
int QByteArrayView_toInt(void* self) {
	auto&& __r = ((QByteArrayView*)self)->toInt();
	return __r;
}
unsigned int QByteArrayView_toUInt(void* self) {
	auto&& __r = ((QByteArrayView*)self)->toUInt();
	return __r;
}
long QByteArrayView_toLong(void* self) {
	auto&& __r = ((QByteArrayView*)self)->toLong();
	return __r;
}
unsigned long QByteArrayView_toULong(void* self) {
	auto&& __r = ((QByteArrayView*)self)->toULong();
	return __r;
}
long long QByteArrayView_toLongLong(void* self) {
	auto&& __r = ((QByteArrayView*)self)->toLongLong();
	return __r;
}
unsigned long long QByteArrayView_toULongLong(void* self) {
	auto&& __r = ((QByteArrayView*)self)->toULongLong();
	return __r;
}
float QByteArrayView_toFloat(void* self) {
	auto&& __r = ((QByteArrayView*)self)->toFloat();
	return __r;
}
double QByteArrayView_toDouble(void* self) {
	auto&& __r = ((QByteArrayView*)self)->toDouble();
	return __r;
}
int QByteArrayView_startsWith(void* self, void* other) {
	auto&& __r = ((QByteArrayView*)self)->startsWith(*(QByteArrayView*)other);
	return __r;
}
int QByteArrayView_startsWith2(void* self, char c) {
	auto&& __r = ((QByteArrayView*)self)->startsWith(c);
	return __r;
}
int QByteArrayView_endsWith(void* self, void* other) {
	auto&& __r = ((QByteArrayView*)self)->endsWith(*(QByteArrayView*)other);
	return __r;
}
int QByteArrayView_endsWith2(void* self, char c) {
	auto&& __r = ((QByteArrayView*)self)->endsWith(c);
	return __r;
}
long long QByteArrayView_indexOf(void* self, void* a) {
	auto&& __r = ((QByteArrayView*)self)->indexOf(*(QByteArrayView*)a);
	return __r;
}
long long QByteArrayView_indexOf2(void* self, char ch) {
	auto&& __r = ((QByteArrayView*)self)->indexOf(ch);
	return __r;
}
int QByteArrayView_contains(void* self, void* a) {
	auto&& __r = ((QByteArrayView*)self)->contains(*(QByteArrayView*)a);
	return __r;
}
int QByteArrayView_contains2(void* self, char c) {
	auto&& __r = ((QByteArrayView*)self)->contains(c);
	return __r;
}
long long QByteArrayView_lastIndexOf(void* self, void* a) {
	auto&& __r = ((QByteArrayView*)self)->lastIndexOf(*(QByteArrayView*)a);
	return __r;
}
long long QByteArrayView_lastIndexOf2(void* self, void* a, long long from) {
	auto&& __r = ((QByteArrayView*)self)->lastIndexOf(*(QByteArrayView*)a, from);
	return __r;
}
long long QByteArrayView_lastIndexOf3(void* self, char ch) {
	auto&& __r = ((QByteArrayView*)self)->lastIndexOf(ch);
	return __r;
}
long long QByteArrayView_count(void* self, void* a) {
	auto&& __r = ((QByteArrayView*)self)->count(*(QByteArrayView*)a);
	return __r;
}
long long QByteArrayView_count2(void* self, char ch) {
	auto&& __r = ((QByteArrayView*)self)->count(ch);
	return __r;
}
int QByteArrayView_compare(void* self, void* a) {
	auto&& __r = ((QByteArrayView*)self)->compare(*(QByteArrayView*)a);
	return __r;
}
int QByteArrayView_isValidUtf8(void* self) {
	auto&& __r = ((QByteArrayView*)self)->isValidUtf8();
	return __r;
}
const char* QByteArrayView_begin(void* self) {
	auto&& __r = ((QByteArrayView*)self)->begin();
	return __r;
}
const char* QByteArrayView_end(void* self) {
	auto&& __r = ((QByteArrayView*)self)->end();
	return __r;
}
const char* QByteArrayView_cbegin(void* self) {
	auto&& __r = ((QByteArrayView*)self)->cbegin();
	return __r;
}
const char* QByteArrayView_cend(void* self) {
	auto&& __r = ((QByteArrayView*)self)->cend();
	return __r;
}
int QByteArrayView_empty(void* self) {
	auto&& __r = ((QByteArrayView*)self)->empty();
	return __r;
}
char QByteArrayView_front(void* self) {
	auto&& __r = ((QByteArrayView*)self)->front();
	return __r;
}
char QByteArrayView_back(void* self) {
	auto&& __r = ((QByteArrayView*)self)->back();
	return __r;
}
long long QByteArrayView_max_size(void* self) {
	auto&& __r = ((QByteArrayView*)self)->max_size();
	return __r;
}
int QByteArrayView_isNull(void* self) {
	auto&& __r = ((QByteArrayView*)self)->isNull();
	return __r;
}
int QByteArrayView_isEmpty(void* self) {
	auto&& __r = ((QByteArrayView*)self)->isEmpty();
	return __r;
}
long long QByteArrayView_length(void* self) {
	auto&& __r = ((QByteArrayView*)self)->length();
	return __r;
}
char QByteArrayView_first2(void* self) {
	auto&& __r = ((QByteArrayView*)self)->first();
	return __r;
}
char QByteArrayView_last2(void* self) {
	auto&& __r = ((QByteArrayView*)self)->last();
	return __r;
}
long long QByteArrayView_maxSize() {
	auto&& __r = QByteArrayView::maxSize();
	return __r;
}
void* QByteArrayView_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QByteArrayView*)self)->operator=(*(QByteArrayView*)arg1);
	return (void*)&(__r);
}
void* QByteArrayView_mid2(void* self, long long pos, long long n) {
	auto&& __r = ((QByteArrayView*)self)->mid(pos, n);
	return new QByteArrayView(__r);
}
short QByteArrayView_toShort2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toShort(&ok_local);
	*ok = ok_local;
	return __r;
}
short QByteArrayView_toShort3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toShort(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned short QByteArrayView_toUShort2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toUShort(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned short QByteArrayView_toUShort3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toUShort(&ok_local, base);
	*ok = ok_local;
	return __r;
}
int QByteArrayView_toInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toInt(&ok_local);
	*ok = ok_local;
	return __r;
}
int QByteArrayView_toInt3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toInt(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned int QByteArrayView_toUInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toUInt(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned int QByteArrayView_toUInt3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toUInt(&ok_local, base);
	*ok = ok_local;
	return __r;
}
long QByteArrayView_toLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toLong(&ok_local);
	*ok = ok_local;
	return __r;
}
long QByteArrayView_toLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned long QByteArrayView_toULong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toULong(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long QByteArrayView_toULong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toULong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
long long QByteArrayView_toLongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toLongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
long long QByteArrayView_toLongLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toLongLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned long long QByteArrayView_toULongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toULongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long long QByteArrayView_toULongLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toULongLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
float QByteArrayView_toFloat2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toFloat(&ok_local);
	*ok = ok_local;
	return __r;
}
double QByteArrayView_toDouble2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArrayView*)self)->toDouble(&ok_local);
	*ok = ok_local;
	return __r;
}
long long QByteArrayView_indexOf3(void* self, void* a, long long from) {
	auto&& __r = ((QByteArrayView*)self)->indexOf(*(QByteArrayView*)a, from);
	return __r;
}
long long QByteArrayView_indexOf4(void* self, char ch, long long from) {
	auto&& __r = ((QByteArrayView*)self)->indexOf(ch, from);
	return __r;
}
long long QByteArrayView_lastIndexOf4(void* self, char ch, long long from) {
	auto&& __r = ((QByteArrayView*)self)->lastIndexOf(ch, from);
	return __r;
}
int QByteArrayView_compare2(void* self, void* a, int cs) {
	auto&& __r = ((QByteArrayView*)self)->compare(*(QByteArrayView*)a, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QByteArrayView_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) == (*(QByteArrayView*)rhs);
	return __r;
}
int QByteArrayView_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) != (*(QByteArrayView*)rhs);
	return __r;
}
int QByteArrayView_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) < (*(QByteArrayView*)rhs);
	return __r;
}
int QByteArrayView_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) > (*(QByteArrayView*)rhs);
	return __r;
}
int QByteArrayView_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) <= (*(QByteArrayView*)rhs);
	return __r;
}
int QByteArrayView_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) >= (*(QByteArrayView*)rhs);
	return __r;
}
int QByteArrayView_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) == (*(QChar*)rhs);
	return __r;
}
int QByteArrayView_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) != (*(QChar*)rhs);
	return __r;
}
int QByteArrayView_OperatorLesser2(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) < (*(QChar*)rhs);
	return __r;
}
int QByteArrayView_OperatorGreater2(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) > (*(QChar*)rhs);
	return __r;
}
int QByteArrayView_OperatorLesserOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) <= (*(QChar*)rhs);
	return __r;
}
int QByteArrayView_OperatorGreaterOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) >= (*(QChar*)rhs);
	return __r;
}
int QByteArrayView_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) == (*(QStringView*)rhs);
	return __r;
}
int QByteArrayView_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) != (*(QStringView*)rhs);
	return __r;
}
int QByteArrayView_OperatorLesser3(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) < (*(QStringView*)rhs);
	return __r;
}
int QByteArrayView_OperatorGreater3(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) > (*(QStringView*)rhs);
	return __r;
}
int QByteArrayView_OperatorLesserOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) <= (*(QStringView*)rhs);
	return __r;
}
int QByteArrayView_OperatorGreaterOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QByteArrayView*)self) >= (*(QStringView*)rhs);
	return __r;
}
int QByteArrayView_OperatorEqual4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArrayView*)self) == (rhs_qstr);
	return __r;
}
int QByteArrayView_OperatorNotEqual4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArrayView*)self) != (rhs_qstr);
	return __r;
}
int QByteArrayView_OperatorLesser4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArrayView*)self) < (rhs_qstr);
	return __r;
}
int QByteArrayView_OperatorGreater4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArrayView*)self) > (rhs_qstr);
	return __r;
}
int QByteArrayView_OperatorLesserOrEqual4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArrayView*)self) <= (rhs_qstr);
	return __r;
}
int QByteArrayView_OperatorGreaterOrEqual4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArrayView*)self) >= (rhs_qstr);
	return __r;
}
void QByteArrayView_delete(void* self) { delete (QByteArrayView*)self; }

}

#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qatomic.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qobjectdata.h"

extern "C" {

void* QObjectData_dynamicMetaObject(void* self) {
	auto&& __r = ((QObjectData*)self)->dynamicMetaObject();
	return (void*)(__r);
}
void* QObjectData_q_ptr_get(void* self) {
	auto __r = ((QObjectData*)self)->q_ptr;
	return (void*)(__r);
}
void QObjectData_q_ptr_set(void* self, void* q_ptr) {
	((QObjectData*)self)->q_ptr = (QObject*)q_ptr;
}
void* QObjectData_parent_get(void* self) {
	auto __r = ((QObjectData*)self)->parent;
	return (void*)(__r);
}
void QObjectData_parent_set(void* self, void* parent) {
	((QObjectData*)self)->parent = (QObject*)parent;
}
void** QObjectData_children_get(void* self, long long* __out_len) {
	auto __r = ((QObjectData*)self)->children;
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QObjectData_children_set(void* self, void** children, long long children_len) {
	QList<QObject*> children_list; children_list.reserve(children_len);
	for (long long __i = 0; __i < children_len; __i++) children_list.append((QObject*)children[__i]);
	((QObjectData*)self)->children = children_list;
}
unsigned int QObjectData_isWidget_get(void* self) {
	auto __r = ((QObjectData*)self)->isWidget;
	return __r;
}
void QObjectData_isWidget_set(void* self, unsigned int isWidget) {
	((QObjectData*)self)->isWidget = isWidget;
}
unsigned int QObjectData_blockSig_get(void* self) {
	auto __r = ((QObjectData*)self)->blockSig;
	return __r;
}
void QObjectData_blockSig_set(void* self, unsigned int blockSig) {
	((QObjectData*)self)->blockSig = blockSig;
}
unsigned int QObjectData_wasDeleted_get(void* self) {
	auto __r = ((QObjectData*)self)->wasDeleted;
	return __r;
}
void QObjectData_wasDeleted_set(void* self, unsigned int wasDeleted) {
	((QObjectData*)self)->wasDeleted = wasDeleted;
}
unsigned int QObjectData_isDeletingChildren_get(void* self) {
	auto __r = ((QObjectData*)self)->isDeletingChildren;
	return __r;
}
void QObjectData_isDeletingChildren_set(void* self, unsigned int isDeletingChildren) {
	((QObjectData*)self)->isDeletingChildren = isDeletingChildren;
}
unsigned int QObjectData_sendChildEvents_get(void* self) {
	auto __r = ((QObjectData*)self)->sendChildEvents;
	return __r;
}
void QObjectData_sendChildEvents_set(void* self, unsigned int sendChildEvents) {
	((QObjectData*)self)->sendChildEvents = sendChildEvents;
}
unsigned int QObjectData_receiveChildEvents_get(void* self) {
	auto __r = ((QObjectData*)self)->receiveChildEvents;
	return __r;
}
void QObjectData_receiveChildEvents_set(void* self, unsigned int receiveChildEvents) {
	((QObjectData*)self)->receiveChildEvents = receiveChildEvents;
}
unsigned int QObjectData_isWindow_get(void* self) {
	auto __r = ((QObjectData*)self)->isWindow;
	return __r;
}
void QObjectData_isWindow_set(void* self, unsigned int isWindow) {
	((QObjectData*)self)->isWindow = isWindow;
}
unsigned int QObjectData_deleteLaterCalled_get(void* self) {
	auto __r = ((QObjectData*)self)->deleteLaterCalled;
	return __r;
}
void QObjectData_deleteLaterCalled_set(void* self, unsigned int deleteLaterCalled) {
	((QObjectData*)self)->deleteLaterCalled = deleteLaterCalled;
}
unsigned int QObjectData_isQuickItem_get(void* self) {
	auto __r = ((QObjectData*)self)->isQuickItem;
	return __r;
}
void QObjectData_isQuickItem_set(void* self, unsigned int isQuickItem) {
	((QObjectData*)self)->isQuickItem = isQuickItem;
}
unsigned int QObjectData_willBeWidget_get(void* self) {
	auto __r = ((QObjectData*)self)->willBeWidget;
	return __r;
}
void QObjectData_willBeWidget_set(void* self, unsigned int willBeWidget) {
	((QObjectData*)self)->willBeWidget = willBeWidget;
}
unsigned int QObjectData_wasWidget_get(void* self) {
	auto __r = ((QObjectData*)self)->wasWidget;
	return __r;
}
void QObjectData_wasWidget_set(void* self, unsigned int wasWidget) {
	((QObjectData*)self)->wasWidget = wasWidget;
}
unsigned int QObjectData_receiveParentEvents_get(void* self) {
	auto __r = ((QObjectData*)self)->receiveParentEvents;
	return __r;
}
void QObjectData_receiveParentEvents_set(void* self, unsigned int receiveParentEvents) {
	((QObjectData*)self)->receiveParentEvents = receiveParentEvents;
}
unsigned int QObjectData_unused_get(void* self) {
	auto __r = ((QObjectData*)self)->unused;
	return __r;
}
void QObjectData_unused_set(void* self, unsigned int unused) {
	((QObjectData*)self)->unused = unused;
}
void* QObjectData_postedEvents_get(void* self) {
	auto __r = ((QObjectData*)self)->postedEvents;
	return new QAtomicInt(__r);
}
void QObjectData_postedEvents_set(void* self, void* postedEvents) {
	((QObjectData*)self)->postedEvents = *(QAtomicInt*)postedEvents;
}
void* QObjectData_bindingStorage_get(void* self) {
	auto __r = ((QObjectData*)self)->bindingStorage;
	return new QBindingStorage(__r);
}
void QObjectData_bindingStorage_set(void* self, void* bindingStorage) {
	((QObjectData*)self)->bindingStorage = *(QBindingStorage*)bindingStorage;
}
void QObjectData_delete(void* self) { delete (QObjectData*)self; }

}

#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qvariant.h>
#include "gen_qmodelroledata.h"

extern "C" {

void* QModelRoleData_new(int role) {
	return new QModelRoleData(role);
}
void* QModelRoleData_new2(void* arg1) {
	return new QModelRoleData(*(QModelRoleData*)arg1);
}
int QModelRoleData_role(void* self) {
	auto&& __r = ((QModelRoleData*)self)->role();
	return __r;
}
void* QModelRoleData_data(void* self) {
	auto&& __r = ((QModelRoleData*)self)->data();
	return (void*)&(__r);
}
void* QModelRoleData_data2(void* self) {
	auto&& __r = ((QModelRoleData*)self)->data();
	return (void*)&(__r);
}
void QModelRoleData_clearData(void* self) {
	((QModelRoleData*)self)->clearData();
}
void* QModelRoleData_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QModelRoleData*)self)->operator=(*(QModelRoleData*)arg1);
	return (void*)&(__r);
}
void QModelRoleData_delete(void* self) { delete (QModelRoleData*)self; }

}

#include <QtCore/qiodevice.h>
#include <QtCore/qobject.h>
#include <QtCore/qiodevicebase.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qiodevice.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QIODevice;
static std::mutex __connsMu_QIODevice;

extern "C" {

void* QIODevice_base(void* self) { return (void*)static_cast<QObject*>((QIODevice*)self); }
void* QIODevice_fromQObject(void* o) { return (void*)qobject_cast<QIODevice*>((QObject*)o); }
char* QIODevice_tr(const char* s, long long* __out_len) {
	auto&& __r = QIODevice::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QIODevice_openMode(void* self) {
	auto&& __r = ((QIODevice*)self)->openMode();
	return (int)(__r);
}
void QIODevice_setTextModeEnabled(void* self, int enabled) {
	((QIODevice*)self)->setTextModeEnabled(enabled);
}
int QIODevice_isTextModeEnabled(void* self) {
	auto&& __r = ((QIODevice*)self)->isTextModeEnabled();
	return __r;
}
int QIODevice_isOpen(void* self) {
	auto&& __r = ((QIODevice*)self)->isOpen();
	return __r;
}
int QIODevice_isReadable(void* self) {
	auto&& __r = ((QIODevice*)self)->isReadable();
	return __r;
}
int QIODevice_isWritable(void* self) {
	auto&& __r = ((QIODevice*)self)->isWritable();
	return __r;
}
int QIODevice_isSequential(void* self) {
	auto&& __r = ((QIODevice*)self)->isSequential();
	return __r;
}
int QIODevice_readChannelCount(void* self) {
	auto&& __r = ((QIODevice*)self)->readChannelCount();
	return __r;
}
int QIODevice_writeChannelCount(void* self) {
	auto&& __r = ((QIODevice*)self)->writeChannelCount();
	return __r;
}
int QIODevice_currentReadChannel(void* self) {
	auto&& __r = ((QIODevice*)self)->currentReadChannel();
	return __r;
}
void QIODevice_setCurrentReadChannel(void* self, int channel) {
	((QIODevice*)self)->setCurrentReadChannel(channel);
}
int QIODevice_currentWriteChannel(void* self) {
	auto&& __r = ((QIODevice*)self)->currentWriteChannel();
	return __r;
}
void QIODevice_setCurrentWriteChannel(void* self, int channel) {
	((QIODevice*)self)->setCurrentWriteChannel(channel);
}
int QIODevice_open(void* self, int mode) {
	auto&& __r = ((QIODevice*)self)->open(QFlags<QIODeviceBase::OpenModeFlag>(mode));
	return __r;
}
void QIODevice_close(void* self) {
	((QIODevice*)self)->close();
}
long long QIODevice_pos(void* self) {
	auto&& __r = ((QIODevice*)self)->pos();
	return __r;
}
long long QIODevice_size(void* self) {
	auto&& __r = ((QIODevice*)self)->size();
	return __r;
}
int QIODevice_seek(void* self, long long pos) {
	auto&& __r = ((QIODevice*)self)->seek(pos);
	return __r;
}
int QIODevice_atEnd(void* self) {
	auto&& __r = ((QIODevice*)self)->atEnd();
	return __r;
}
int QIODevice_reset(void* self) {
	auto&& __r = ((QIODevice*)self)->reset();
	return __r;
}
long long QIODevice_bytesAvailable(void* self) {
	auto&& __r = ((QIODevice*)self)->bytesAvailable();
	return __r;
}
long long QIODevice_bytesToWrite(void* self) {
	auto&& __r = ((QIODevice*)self)->bytesToWrite();
	return __r;
}
long long QIODevice_read(void* self, char* data, long long maxlen) {
	auto&& __r = ((QIODevice*)self)->read((char*)data, maxlen);
	return __r;
}
void* QIODevice_read2(void* self, long long maxlen) {
	auto&& __r = ((QIODevice*)self)->read(maxlen);
	return new QByteArray(__r);
}
void* QIODevice_readAll(void* self) {
	auto&& __r = ((QIODevice*)self)->readAll();
	return new QByteArray(__r);
}
long long QIODevice_readLine(void* self, char* data, long long maxlen) {
	auto&& __r = ((QIODevice*)self)->readLine((char*)data, maxlen);
	return __r;
}
void* QIODevice_readLine2(void* self) {
	auto&& __r = ((QIODevice*)self)->readLine();
	return new QByteArray(__r);
}
int QIODevice_readLineInto(void* self, void* result) {
	auto&& __r = ((QIODevice*)self)->readLineInto((QByteArray*)result);
	return __r;
}
int QIODevice_canReadLine(void* self) {
	auto&& __r = ((QIODevice*)self)->canReadLine();
	return __r;
}
void QIODevice_startTransaction(void* self) {
	((QIODevice*)self)->startTransaction();
}
void QIODevice_commitTransaction(void* self) {
	((QIODevice*)self)->commitTransaction();
}
void QIODevice_rollbackTransaction(void* self) {
	((QIODevice*)self)->rollbackTransaction();
}
int QIODevice_isTransactionStarted(void* self) {
	auto&& __r = ((QIODevice*)self)->isTransactionStarted();
	return __r;
}
long long QIODevice_write(void* self, const char* data, long long len_) {
	auto&& __r = ((QIODevice*)self)->write(data, len_);
	return __r;
}
long long QIODevice_write2(void* self, const char* data) {
	auto&& __r = ((QIODevice*)self)->write(data);
	return __r;
}
long long QIODevice_write3(void* self, void* data) {
	auto&& __r = ((QIODevice*)self)->write(*(QByteArray*)data);
	return __r;
}
long long QIODevice_peek(void* self, char* data, long long maxlen) {
	auto&& __r = ((QIODevice*)self)->peek((char*)data, maxlen);
	return __r;
}
void* QIODevice_peek2(void* self, long long maxlen) {
	auto&& __r = ((QIODevice*)self)->peek(maxlen);
	return new QByteArray(__r);
}
long long QIODevice_skip(void* self, long long maxSize) {
	auto&& __r = ((QIODevice*)self)->skip(maxSize);
	return __r;
}
int QIODevice_waitForReadyRead(void* self, int msecs) {
	auto&& __r = ((QIODevice*)self)->waitForReadyRead(msecs);
	return __r;
}
int QIODevice_waitForBytesWritten(void* self, int msecs) {
	auto&& __r = ((QIODevice*)self)->waitForBytesWritten(msecs);
	return __r;
}
void QIODevice_ungetChar(void* self, char c) {
	((QIODevice*)self)->ungetChar(c);
}
int QIODevice_putChar(void* self, char c) {
	auto&& __r = ((QIODevice*)self)->putChar(c);
	return __r;
}
int QIODevice_getChar(void* self, char* c) {
	auto&& __r = ((QIODevice*)self)->getChar((char*)c);
	return __r;
}
char* QIODevice_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QIODevice*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QIODevice_readyRead(void* self) {
	((QIODevice*)self)->readyRead();
}
void QIODevice_channelReadyRead(void* self, int channel) {
	((QIODevice*)self)->channelReadyRead(channel);
}
void QIODevice_bytesWritten(void* self, long long bytes) {
	((QIODevice*)self)->bytesWritten(bytes);
}
void QIODevice_channelBytesWritten(void* self, int channel, long long bytes) {
	((QIODevice*)self)->channelBytesWritten(channel, bytes);
}
void QIODevice_aboutToClose(void* self) {
	((QIODevice*)self)->aboutToClose();
}
void QIODevice_readChannelFinished(void* self) {
	((QIODevice*)self)->readChannelFinished();
}
void* QIODevice_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QIODevice::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QIODevice_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QIODevice::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QIODevice_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QIODevice::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QIODevice_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QIODevice::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QIODevice_readLine3(void* self, long long maxlen) {
	auto&& __r = ((QIODevice*)self)->readLine(maxlen);
	return new QByteArray(__r);
}
int QIODevice_readLineInto2(void* self, void* result, long long maxlen) {
	auto&& __r = ((QIODevice*)self)->readLineInto((QByteArray*)result, maxlen);
	return __r;
}
void* QIODevice_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QIODevice::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QIODevice_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QIODevice::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QIODevice_asQIODeviceBase(void* self) {
	auto&& __r = static_cast<QIODeviceBase*>((QIODevice*)self);
	return (void*)(__r);
}
void QIODevice_connect_readyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIODevice*)self, &QIODevice::readyRead, (QIODevice*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIODevice);
	__conns_QIODevice[__id] = __c;
}
void QIODevice_connect_channelReadyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIODevice*)self, &QIODevice::channelReadyRead, (QIODevice*)self, [__id](int channel){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, channel);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIODevice);
	__conns_QIODevice[__id] = __c;
}
void QIODevice_connect_bytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIODevice*)self, &QIODevice::bytesWritten, (QIODevice*)self, [__id](qint64 bytes){
		if (cppgo_sig_void_longlong) ((void(*)(long long, long long))cppgo_sig_void_longlong)(__id, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIODevice);
	__conns_QIODevice[__id] = __c;
}
void QIODevice_connect_channelBytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIODevice*)self, &QIODevice::channelBytesWritten, (QIODevice*)self, [__id](int channel, qint64 bytes){
		if (cppgo_sig_void_int_longlong) ((void(*)(long long, int, long long))cppgo_sig_void_int_longlong)(__id, channel, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIODevice);
	__conns_QIODevice[__id] = __c;
}
void QIODevice_connect_aboutToClose(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIODevice*)self, &QIODevice::aboutToClose, (QIODevice*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIODevice);
	__conns_QIODevice[__id] = __c;
}
void QIODevice_connect_readChannelFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIODevice*)self, &QIODevice::readChannelFinished, (QIODevice*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIODevice);
	__conns_QIODevice[__id] = __c;
}
void QIODevice_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIODevice*)self, &QIODevice::destroyed, (QIODevice*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIODevice);
	__conns_QIODevice[__id] = __c;
}
void QIODevice_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIODevice*)self, &QIODevice::destroyed, (QIODevice*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIODevice);
	__conns_QIODevice[__id] = __c;
}
void QIODevice_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIODevice*)self, &QIODevice::objectNameChanged, (QIODevice*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIODevice);
	__conns_QIODevice[__id] = __c;
}
void QIODevice__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QIODevice);
		auto __it = __conns_QIODevice.find(id);
		if (__it == __conns_QIODevice.end()) return;
		__c = __it->second;
		__conns_QIODevice.erase(__it);
	}
	QObject::disconnect(__c);
}
void QIODevice_delete(void* self) { delete (QIODevice*)self; }

}

#include <QtCore/qcborstreamreader.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qcborcommon.h>
#include <QtCore/qstring.h>
#include <QtCore/qfloat16.h>
#include "gen_qcborstreamreader.h"

extern "C" {

void* QCborStreamReader_new() {
	return new QCborStreamReader();
}
void* QCborStreamReader_new2(const char* data, long long len_) {
	return new QCborStreamReader(data, len_);
}
void* QCborStreamReader_new3(const unsigned char* data, long long len_) {
	return new QCborStreamReader(data, len_);
}
void* QCborStreamReader_new4(void* data) {
	return new QCborStreamReader(*(QByteArray*)data);
}
void* QCborStreamReader_new5(void* device) {
	return new QCborStreamReader((QIODevice*)device);
}
void QCborStreamReader_setDevice(void* self, void* device) {
	((QCborStreamReader*)self)->setDevice((QIODevice*)device);
}
void* QCborStreamReader_device(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->device();
	return (void*)(__r);
}
void QCborStreamReader_addData(void* self, void* data) {
	((QCborStreamReader*)self)->addData(*(QByteArray*)data);
}
void QCborStreamReader_addData2(void* self, const char* data, long long len_) {
	((QCborStreamReader*)self)->addData(data, len_);
}
void QCborStreamReader_addData3(void* self, const unsigned char* data, long long len_) {
	((QCborStreamReader*)self)->addData(data, len_);
}
void QCborStreamReader_reparse(void* self) {
	((QCborStreamReader*)self)->reparse();
}
void QCborStreamReader_clear(void* self) {
	((QCborStreamReader*)self)->clear();
}
void QCborStreamReader_reset(void* self) {
	((QCborStreamReader*)self)->reset();
}
void* QCborStreamReader_lastError(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->lastError();
	return new QCborError(__r);
}
long long QCborStreamReader_currentOffset(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->currentOffset();
	return __r;
}
int QCborStreamReader_isValid(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isValid();
	return __r;
}
int QCborStreamReader_containerDepth(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->containerDepth();
	return __r;
}
int QCborStreamReader_parentContainerType(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->parentContainerType();
	return (int)(__r);
}
int QCborStreamReader_hasNext(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->hasNext();
	return __r;
}
int QCborStreamReader_next(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->next();
	return __r;
}
int QCborStreamReader_type(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->type();
	return (int)(__r);
}
int QCborStreamReader_isUnsignedInteger(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isUnsignedInteger();
	return __r;
}
int QCborStreamReader_isNegativeInteger(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isNegativeInteger();
	return __r;
}
int QCborStreamReader_isInteger(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isInteger();
	return __r;
}
int QCborStreamReader_isByteArray(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isByteArray();
	return __r;
}
int QCborStreamReader_isString(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isString();
	return __r;
}
int QCborStreamReader_isArray(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isArray();
	return __r;
}
int QCborStreamReader_isMap(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isMap();
	return __r;
}
int QCborStreamReader_isTag(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isTag();
	return __r;
}
int QCborStreamReader_isSimpleType(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isSimpleType();
	return __r;
}
int QCborStreamReader_isFloat16(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isFloat16();
	return __r;
}
int QCborStreamReader_isFloat(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isFloat();
	return __r;
}
int QCborStreamReader_isDouble(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isDouble();
	return __r;
}
int QCborStreamReader_isInvalid(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isInvalid();
	return __r;
}
int QCborStreamReader_isSimpleType2(void* self, int st) {
	auto&& __r = ((QCborStreamReader*)self)->isSimpleType((QCborSimpleType)(st));
	return __r;
}
int QCborStreamReader_isFalse(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isFalse();
	return __r;
}
int QCborStreamReader_isTrue(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isTrue();
	return __r;
}
int QCborStreamReader_isBool(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isBool();
	return __r;
}
int QCborStreamReader_isNull(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isNull();
	return __r;
}
int QCborStreamReader_isUndefined(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isUndefined();
	return __r;
}
int QCborStreamReader_isLengthKnown(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isLengthKnown();
	return __r;
}
unsigned long long QCborStreamReader_length(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->length();
	return __r;
}
int QCborStreamReader_isContainer(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->isContainer();
	return __r;
}
int QCborStreamReader_enterContainer(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->enterContainer();
	return __r;
}
int QCborStreamReader_leaveContainer(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->leaveContainer();
	return __r;
}
int QCborStreamReader_readAndAppendToString(void* self, void* dst) {
	auto&& __r = ((QCborStreamReader*)self)->readAndAppendToString(*(QString*)dst);
	return __r;
}
int QCborStreamReader_readAndAppendToUtf8String(void* self, void* dst) {
	auto&& __r = ((QCborStreamReader*)self)->readAndAppendToUtf8String(*(QByteArray*)dst);
	return __r;
}
int QCborStreamReader_readAndAppendToByteArray(void* self, void* dst) {
	auto&& __r = ((QCborStreamReader*)self)->readAndAppendToByteArray(*(QByteArray*)dst);
	return __r;
}
long long QCborStreamReader_currentStringChunkSize(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->currentStringChunkSize();
	return __r;
}
int QCborStreamReader_toBool(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->toBool();
	return __r;
}
int QCborStreamReader_toTag(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->toTag();
	return (int)(__r);
}
unsigned long long QCborStreamReader_toUnsignedInteger(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->toUnsignedInteger();
	return __r;
}
int QCborStreamReader_toNegativeInteger(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->toNegativeInteger();
	return (int)(__r);
}
int QCborStreamReader_toSimpleType(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->toSimpleType();
	return (int)(__r);
}
void* QCborStreamReader_toFloat16(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->toFloat16();
	return new qfloat16(__r);
}
float QCborStreamReader_toFloat(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->toFloat();
	return __r;
}
double QCborStreamReader_toDouble(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->toDouble();
	return __r;
}
long long QCborStreamReader_toInteger(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->toInteger();
	return __r;
}
char* QCborStreamReader_readAllString(void* self, long long* __out_len) {
	auto&& __r = ((QCborStreamReader*)self)->readAllString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCborStreamReader_readAllUtf8String(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->readAllUtf8String();
	return new QByteArray(__r);
}
void* QCborStreamReader_readAllByteArray(void* self) {
	auto&& __r = ((QCborStreamReader*)self)->readAllByteArray();
	return new QByteArray(__r);
}
int QCborStreamReader_next2(void* self, int maxRecursion) {
	auto&& __r = ((QCborStreamReader*)self)->next(maxRecursion);
	return __r;
}
void QCborStreamReader_delete(void* self) { delete (QCborStreamReader*)self; }

}

#include <QtCore/qfile.h>
#include "gen_qntfspermissioncheckguard.h"

extern "C" {

void* QNtfsPermissionCheckGuard_new() {
	return new QNtfsPermissionCheckGuard();
}
void QNtfsPermissionCheckGuard_delete(void* self) { delete (QNtfsPermissionCheckGuard*)self; }

}

#include <QtCore/qjsondocument.h>
#include <QtCore/qjsonobject.h>
#include <QtCore/qjsonarray.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qjsonparseerror.h>
#include <QtCore/qjsonvalue.h>
#include <QtCore/qstring.h>
#include <QtCore/qstringview.h>
#include <QtCore/qlatin1stringview.h>
#include "gen_qjsondocument.h"

extern "C" {

void* QJsonDocument_new() {
	return new QJsonDocument();
}
void* QJsonDocument_new2(void* object) {
	return new QJsonDocument(*(QJsonObject*)object);
}
void* QJsonDocument_new3(void* array) {
	return new QJsonDocument(*(QJsonArray*)array);
}
void* QJsonDocument_new4(void* other) {
	return new QJsonDocument(*(QJsonDocument*)other);
}
void* QJsonDocument_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QJsonDocument*)self)->operator=(*(QJsonDocument*)other);
	return (void*)&(__r);
}
void QJsonDocument_swap(void* self, void* other) {
	((QJsonDocument*)self)->swap(*(QJsonDocument*)other);
}
void* QJsonDocument_fromVariant(void* variant) {
	auto&& __r = QJsonDocument::fromVariant(*(QVariant*)variant);
	return new QJsonDocument(__r);
}
void* QJsonDocument_toVariant(void* self) {
	auto&& __r = ((QJsonDocument*)self)->toVariant();
	return new QVariant(__r);
}
void* QJsonDocument_fromJson(void* json) {
	auto&& __r = QJsonDocument::fromJson(*(QByteArray*)json);
	return new QJsonDocument(__r);
}
void* QJsonDocument_toJson(void* self) {
	auto&& __r = ((QJsonDocument*)self)->toJson();
	return new QByteArray(__r);
}
int QJsonDocument_isEmpty(void* self) {
	auto&& __r = ((QJsonDocument*)self)->isEmpty();
	return __r;
}
int QJsonDocument_isArray(void* self) {
	auto&& __r = ((QJsonDocument*)self)->isArray();
	return __r;
}
int QJsonDocument_isObject(void* self) {
	auto&& __r = ((QJsonDocument*)self)->isObject();
	return __r;
}
void* QJsonDocument_object(void* self) {
	auto&& __r = ((QJsonDocument*)self)->object();
	return new QJsonObject(__r);
}
void* QJsonDocument_array(void* self) {
	auto&& __r = ((QJsonDocument*)self)->array();
	return new QJsonArray(__r);
}
void QJsonDocument_setObject(void* self, void* object) {
	((QJsonDocument*)self)->setObject(*(QJsonObject*)object);
}
void QJsonDocument_setArray(void* self, void* array) {
	((QJsonDocument*)self)->setArray(*(QJsonArray*)array);
}
void* QJsonDocument_OperatorSubscript(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QJsonDocument*)self)->operator[](key_qstr);
	return new QJsonValue(__r);
}
void* QJsonDocument_OperatorSubscript2(void* self, void* key) {
	auto&& __r = ((QJsonDocument*)self)->operator[](*(QStringView*)key);
	return new QJsonValue(__r);
}
void* QJsonDocument_OperatorSubscript3(void* self, void* key) {
	auto&& __r = ((QJsonDocument*)self)->operator[](*(QLatin1String*)key);
	return new QJsonValue(__r);
}
void* QJsonDocument_OperatorSubscript4(void* self, long long i) {
	auto&& __r = ((QJsonDocument*)self)->operator[](i);
	return new QJsonValue(__r);
}
int QJsonDocument_isNull(void* self) {
	auto&& __r = ((QJsonDocument*)self)->isNull();
	return __r;
}
void* QJsonDocument_fromJson2(void* json, void* error_) {
	auto&& __r = QJsonDocument::fromJson(*(QByteArray*)json, (QJsonParseError*)error_);
	return new QJsonDocument(__r);
}
void* QJsonDocument_toJson2(void* self, int format) {
	auto&& __r = ((QJsonDocument*)self)->toJson((QJsonDocument::JsonFormat)(format));
	return new QByteArray(__r);
}
int QJsonDocument_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonDocument*)self) == (*(QJsonDocument*)rhs);
	return __r;
}
int QJsonDocument_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonDocument*)self) != (*(QJsonDocument*)rhs);
	return __r;
}
void QJsonDocument_delete(void* self) { delete (QJsonDocument*)self; }

}

#include <QtCore/qmimedata.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qurl.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qmimedata.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QMimeData__Go : public QMimeData {
public:
	long long __goId;
	unsigned char __goOn[11] = {};
	template<typename... A> QMimeData__Go(long long id, A&&... a) : QMimeData(std::forward<A>(a)...), __goId(id) {}
	using QMimeData::sender;
	using QMimeData::senderSignalIndex;
	using QMimeData::receivers;
	using QMimeData::isSignalConnected;
	QVariant cppgoBase_retrieveData(const QString& __a0, QMetaType __a1) { return QMimeData::retrieveData(__a0, __a1); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QMimeData::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QMimeData::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QMimeData::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QMimeData::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QMimeData::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QMimeData::metaObject();
	}
	bool hasFormat(const QString& mimetype) const override {
		int __h = 0;
		QByteArray mimetype_ba = mimetype.toUtf8();
		if (__goOn[1] && cppgo_virt_int_constcharp_longlong) { int __r = ((int(*)(long long, long long, const char*, long long, int*))cppgo_virt_int_constcharp_longlong)(__goId, 7523410874251807628LL, (char*)mimetype_ba.constData(), mimetype_ba.size(), &__h); if (__h) return __r; }
		return QMimeData::hasFormat(mimetype);
	}
	QList<QString> formats() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 6926839562432501171LL, &__len, &__h);
			if (__h) {
				QList<QString> __r; __r.reserve(__len);
				char** __aa = (char**)__a;
				for (long long __i = 0; __i < __len; __i++) { __r.append(QString::fromUtf8(__aa[__i])); free(__aa[__i]); }
				if (__a) free(__a);
				return __r;
			}
		}
		return QMimeData::formats();
	}
	QVariant retrieveData(const QString& mimetype, QMetaType preferredType) const override {
		int __h = 0;
		QByteArray mimetype_ba = mimetype.toUtf8();
		if (__goOn[3] && cppgo_virt_voidp_constcharp_longlong_voidp) { void* __r = ((void*(*)(long long, long long, const char*, long long, void*, int*))cppgo_virt_voidp_constcharp_longlong_voidp)(__goId, -2679356254762874311LL, (char*)mimetype_ba.constData(), mimetype_ba.size(), (void*)&preferredType, &__h); if (__h) return *(QVariant*)__r; }
		return QMimeData::retrieveData(mimetype, preferredType);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QMimeData::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QMimeData::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QMimeData::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QMimeData::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QMimeData::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QMimeData::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QMimeData::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QMimeData;
static std::mutex __connsMu_QMimeData;

extern "C" {

void* QMimeData_base(void* self) { return (void*)static_cast<QObject*>((QMimeData*)self); }
void QMimeData__set_override(void* self, int idx) { ((QMimeData__Go*)self)->__goOn[idx] = 1; }
void* QMimeData_fromQObject(void* o) { return (void*)qobject_cast<QMimeData*>((QObject*)o); }
void* QMimeData_new(long long __goid) {
	return new QMimeData__Go(__goid);
}
char* QMimeData_tr(const char* s, long long* __out_len) {
	auto&& __r = QMimeData::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void** QMimeData_urls(void* self, long long* __out_len) {
	auto&& __r = ((QMimeData*)self)->urls();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void QMimeData_setUrls(void* self, void** urls, long long urls_len) {
	QList<QUrl> urls_list; urls_list.reserve(urls_len);
	for (long long __i = 0; __i < urls_len; __i++) urls_list.append(*(QUrl*)urls[__i]);
	((QMimeData*)self)->setUrls(urls_list);
}
int QMimeData_hasUrls(void* self) {
	auto&& __r = ((QMimeData*)self)->hasUrls();
	return __r;
}
char* QMimeData_text(void* self, long long* __out_len) {
	auto&& __r = ((QMimeData*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMimeData_setText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QMimeData*)self)->setText(text_qstr);
}
int QMimeData_hasText(void* self) {
	auto&& __r = ((QMimeData*)self)->hasText();
	return __r;
}
char* QMimeData_html(void* self, long long* __out_len) {
	auto&& __r = ((QMimeData*)self)->html();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMimeData_setHtml(void* self, const char* html, long long html_len) {
	QString html_qstr = QString::fromUtf8(html, html_len);
	((QMimeData*)self)->setHtml(html_qstr);
}
int QMimeData_hasHtml(void* self) {
	auto&& __r = ((QMimeData*)self)->hasHtml();
	return __r;
}
void* QMimeData_imageData(void* self) {
	auto&& __r = ((QMimeData*)self)->imageData();
	return new QVariant(__r);
}
void QMimeData_setImageData(void* self, void* image) {
	((QMimeData*)self)->setImageData(*(QVariant*)image);
}
int QMimeData_hasImage(void* self) {
	auto&& __r = ((QMimeData*)self)->hasImage();
	return __r;
}
void* QMimeData_colorData(void* self) {
	auto&& __r = ((QMimeData*)self)->colorData();
	return new QVariant(__r);
}
void QMimeData_setColorData(void* self, void* color) {
	((QMimeData*)self)->setColorData(*(QVariant*)color);
}
int QMimeData_hasColor(void* self) {
	auto&& __r = ((QMimeData*)self)->hasColor();
	return __r;
}
void* QMimeData_data(void* self, const char* mimetype, long long mimetype_len) {
	QString mimetype_qstr = QString::fromUtf8(mimetype, mimetype_len);
	auto&& __r = ((QMimeData*)self)->data(mimetype_qstr);
	return new QByteArray(__r);
}
void QMimeData_setData(void* self, const char* mimetype, long long mimetype_len, void* data) {
	QString mimetype_qstr = QString::fromUtf8(mimetype, mimetype_len);
	((QMimeData*)self)->setData(mimetype_qstr, *(QByteArray*)data);
}
void QMimeData_removeFormat(void* self, const char* mimetype, long long mimetype_len) {
	QString mimetype_qstr = QString::fromUtf8(mimetype, mimetype_len);
	((QMimeData*)self)->removeFormat(mimetype_qstr);
}
int QMimeData_hasFormat(void* self, const char* mimetype, long long mimetype_len) {
	QString mimetype_qstr = QString::fromUtf8(mimetype, mimetype_len);
	auto&& __r = ((QMimeData*)self)->hasFormat(mimetype_qstr);
	return __r;
}
char** QMimeData_formats(void* self, long long* __out_len) {
	auto&& __r = ((QMimeData*)self)->formats();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QMimeData_clear(void* self) {
	((QMimeData*)self)->clear();
}
void* QMimeData_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QMimeData::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QMimeData_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QMimeData::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QMimeData_sender(void* self) {
	auto&& __r = ((QMimeData__Go*)self)->sender();
	return (void*)(__r);
}
int QMimeData_senderSignalIndex(void* self) {
	auto&& __r = ((QMimeData__Go*)self)->senderSignalIndex();
	return __r;
}
int QMimeData_receivers(void* self, const char* signal) {
	auto&& __r = ((QMimeData__Go*)self)->receivers(signal);
	return __r;
}
int QMimeData_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QMimeData__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QMimeData_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QMimeData::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QMimeData_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QMimeData::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QMimeData_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QMimeData::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QMimeData_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QMimeData::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
int QMimeData_hasFormat_default(void* self, const char* mimetype, long long mimetype_len) {
	QString mimetype_qstr = QString::fromUtf8(mimetype, mimetype_len);
	auto&& __r = ((QMimeData*)self)->QMimeData::hasFormat(mimetype_qstr);
	return __r;
}
char** QMimeData_formats_default(void* self, long long* __out_len) {
	auto&& __r = ((QMimeData*)self)->QMimeData::formats();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QMimeData_metaObject_default(void* self) {
	auto&& __r = ((QMimeData*)self)->QMimeData::metaObject();
	return (void*)(__r);
}
void* QMimeData_retrieveData_default(void* self, const char* mimetype, long long mimetype_len, void* preferredType) {
	QString mimetype_qstr = QString::fromUtf8(mimetype, mimetype_len);
	auto&& __r = ((QMimeData__Go*)self)->cppgoBase_retrieveData(mimetype_qstr, *(QMetaType*)preferredType);
	return new QVariant(__r);
}
int QMimeData_event_default(void* self, void* event) {
	auto&& __r = ((QMimeData*)self)->QMimeData::event((QEvent*)event);
	return __r;
}
int QMimeData_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QMimeData*)self)->QMimeData::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QMimeData_timerEvent_default(void* self, void* event) {
	((QMimeData__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QMimeData_childEvent_default(void* self, void* event) {
	((QMimeData__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QMimeData_customEvent_default(void* self, void* event) {
	((QMimeData__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QMimeData_connectNotify_default(void* self, void* signal) {
	((QMimeData__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QMimeData_disconnectNotify_default(void* self, void* signal) {
	((QMimeData__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QMimeData_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMimeData*)self, &QMimeData::destroyed, (QMimeData*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMimeData);
	__conns_QMimeData[__id] = __c;
}
void QMimeData_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMimeData*)self, &QMimeData::destroyed, (QMimeData*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMimeData);
	__conns_QMimeData[__id] = __c;
}
void QMimeData_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMimeData*)self, &QMimeData::objectNameChanged, (QMimeData*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMimeData);
	__conns_QMimeData[__id] = __c;
}
void QMimeData__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QMimeData);
		auto __it = __conns_QMimeData.find(id);
		if (__it == __conns_QMimeData.end()) return;
		__c = __it->second;
		__conns_QMimeData.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QMimeData__Go_new(long long id) { return new QMimeData__Go(id); }
void QMimeData_delete(void* self) { delete (QMimeData*)self; }

}

#include <QtCore/qprocess.h>
#include "gen_qprocess__createprocessarguments.h"

extern "C" {

int QProcess__CreateProcessArguments_inheritHandles_get(void* self) {
	auto __r = ((QProcess::CreateProcessArguments*)self)->inheritHandles;
	return __r;
}
void QProcess__CreateProcessArguments_inheritHandles_set(void* self, int inheritHandles) {
	((QProcess::CreateProcessArguments*)self)->inheritHandles = inheritHandles;
}
unsigned long QProcess__CreateProcessArguments_flags_get(void* self) {
	auto __r = ((QProcess::CreateProcessArguments*)self)->flags;
	return __r;
}
void QProcess__CreateProcessArguments_flags_set(void* self, unsigned long flags) {
	((QProcess::CreateProcessArguments*)self)->flags = flags;
}
void QProcess__CreateProcessArguments_delete(void* self) { delete (QProcess::CreateProcessArguments*)self; }

}

#include <QtCore/qsharedmemory.h>
#include <QtCore/qobject.h>
#include <QtCore/qtipccommon.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qsharedmemory.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSharedMemory__Go : public QSharedMemory {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QSharedMemory__Go(long long id, A&&... a) : QSharedMemory(std::forward<A>(a)...), __goId(id) {}
	using QSharedMemory::sender;
	using QSharedMemory::senderSignalIndex;
	using QSharedMemory::receivers;
	using QSharedMemory::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSharedMemory::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSharedMemory::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSharedMemory::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSharedMemory::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSharedMemory::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSharedMemory::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QSharedMemory::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSharedMemory::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSharedMemory::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSharedMemory::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSharedMemory::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSharedMemory::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSharedMemory::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSharedMemory;
static std::mutex __connsMu_QSharedMemory;

extern "C" {

void* QSharedMemory_base(void* self) { return (void*)static_cast<QObject*>((QSharedMemory*)self); }
void QSharedMemory__set_override(void* self, int idx) { ((QSharedMemory__Go*)self)->__goOn[idx] = 1; }
void* QSharedMemory_fromQObject(void* o) { return (void*)qobject_cast<QSharedMemory*>((QObject*)o); }
void* QSharedMemory_new(long long __goid) {
	return new QSharedMemory__Go(__goid);
}
void* QSharedMemory_new2(long long __goid, void* key) {
	return new QSharedMemory__Go(__goid, *(QNativeIpcKey*)key);
}
void* QSharedMemory_new3(long long __goid, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	return new QSharedMemory__Go(__goid, key_qstr);
}
void* QSharedMemory_new4(long long __goid, void* parent) {
	return new QSharedMemory__Go(__goid, (QObject*)parent);
}
void* QSharedMemory_new5(long long __goid, void* key, void* parent) {
	return new QSharedMemory__Go(__goid, *(QNativeIpcKey*)key, (QObject*)parent);
}
void* QSharedMemory_new6(long long __goid, const char* key, long long key_len, void* parent) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	return new QSharedMemory__Go(__goid, key_qstr, (QObject*)parent);
}
char* QSharedMemory_tr(const char* s, long long* __out_len) {
	auto&& __r = QSharedMemory::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSharedMemory_setKey(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QSharedMemory*)self)->setKey(key_qstr);
}
char* QSharedMemory_key(void* self, long long* __out_len) {
	auto&& __r = ((QSharedMemory*)self)->key();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSharedMemory_setNativeKey(void* self, void* key) {
	((QSharedMemory*)self)->setNativeKey(*(QNativeIpcKey*)key);
}
void QSharedMemory_setNativeKey2(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QSharedMemory*)self)->setNativeKey(key_qstr);
}
char* QSharedMemory_nativeKey(void* self, long long* __out_len) {
	auto&& __r = ((QSharedMemory*)self)->nativeKey();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSharedMemory_nativeIpcKey(void* self) {
	auto&& __r = ((QSharedMemory*)self)->nativeIpcKey();
	return new QNativeIpcKey(__r);
}
int QSharedMemory_create(void* self, long long size) {
	auto&& __r = ((QSharedMemory*)self)->create(size);
	return __r;
}
long long QSharedMemory_size(void* self) {
	auto&& __r = ((QSharedMemory*)self)->size();
	return __r;
}
int QSharedMemory_attach(void* self) {
	auto&& __r = ((QSharedMemory*)self)->attach();
	return __r;
}
int QSharedMemory_isAttached(void* self) {
	auto&& __r = ((QSharedMemory*)self)->isAttached();
	return __r;
}
int QSharedMemory_detach(void* self) {
	auto&& __r = ((QSharedMemory*)self)->detach();
	return __r;
}
int QSharedMemory_lock(void* self) {
	auto&& __r = ((QSharedMemory*)self)->lock();
	return __r;
}
int QSharedMemory_unlock(void* self) {
	auto&& __r = ((QSharedMemory*)self)->unlock();
	return __r;
}
int QSharedMemory_error(void* self) {
	auto&& __r = ((QSharedMemory*)self)->error();
	return (int)(__r);
}
char* QSharedMemory_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QSharedMemory*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QSharedMemory_isKeyTypeSupported(int type_) {
	auto&& __r = QSharedMemory::isKeyTypeSupported((QNativeIpcKey::Type)(type_));
	return __r;
}
void* QSharedMemory_platformSafeKey(const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = QSharedMemory::platformSafeKey(key_qstr);
	return new QNativeIpcKey(__r);
}
void* QSharedMemory_legacyNativeKey(const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = QSharedMemory::legacyNativeKey(key_qstr);
	return new QNativeIpcKey(__r);
}
void* QSharedMemory_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSharedMemory::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSharedMemory_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSharedMemory::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSharedMemory_sender(void* self) {
	auto&& __r = ((QSharedMemory__Go*)self)->sender();
	return (void*)(__r);
}
int QSharedMemory_senderSignalIndex(void* self) {
	auto&& __r = ((QSharedMemory__Go*)self)->senderSignalIndex();
	return __r;
}
int QSharedMemory_receivers(void* self, const char* signal) {
	auto&& __r = ((QSharedMemory__Go*)self)->receivers(signal);
	return __r;
}
int QSharedMemory_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSharedMemory__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QSharedMemory_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSharedMemory::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSharedMemory_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSharedMemory::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSharedMemory_setNativeKey3(void* self, const char* key, long long key_len, int type_) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QSharedMemory*)self)->setNativeKey(key_qstr, (QNativeIpcKey::Type)(type_));
}
int QSharedMemory_create2(void* self, long long size, int mode) {
	auto&& __r = ((QSharedMemory*)self)->create(size, (QSharedMemory::AccessMode)(mode));
	return __r;
}
int QSharedMemory_attach2(void* self, int mode) {
	auto&& __r = ((QSharedMemory*)self)->attach((QSharedMemory::AccessMode)(mode));
	return __r;
}
void* QSharedMemory_platformSafeKey2(const char* key, long long key_len, int type_) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = QSharedMemory::platformSafeKey(key_qstr, (QNativeIpcKey::Type)(type_));
	return new QNativeIpcKey(__r);
}
void* QSharedMemory_legacyNativeKey2(const char* key, long long key_len, int type_) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = QSharedMemory::legacyNativeKey(key_qstr, (QNativeIpcKey::Type)(type_));
	return new QNativeIpcKey(__r);
}
void* QSharedMemory_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSharedMemory::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSharedMemory_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSharedMemory::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSharedMemory_metaObject_default(void* self) {
	auto&& __r = ((QSharedMemory*)self)->QSharedMemory::metaObject();
	return (void*)(__r);
}
int QSharedMemory_event_default(void* self, void* event) {
	auto&& __r = ((QSharedMemory*)self)->QSharedMemory::event((QEvent*)event);
	return __r;
}
int QSharedMemory_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSharedMemory*)self)->QSharedMemory::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSharedMemory_timerEvent_default(void* self, void* event) {
	((QSharedMemory__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSharedMemory_childEvent_default(void* self, void* event) {
	((QSharedMemory__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSharedMemory_customEvent_default(void* self, void* event) {
	((QSharedMemory__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSharedMemory_connectNotify_default(void* self, void* signal) {
	((QSharedMemory__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSharedMemory_disconnectNotify_default(void* self, void* signal) {
	((QSharedMemory__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSharedMemory_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSharedMemory*)self, &QSharedMemory::destroyed, (QSharedMemory*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSharedMemory);
	__conns_QSharedMemory[__id] = __c;
}
void QSharedMemory_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSharedMemory*)self, &QSharedMemory::destroyed, (QSharedMemory*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSharedMemory);
	__conns_QSharedMemory[__id] = __c;
}
void QSharedMemory_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSharedMemory*)self, &QSharedMemory::objectNameChanged, (QSharedMemory*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSharedMemory);
	__conns_QSharedMemory[__id] = __c;
}
void QSharedMemory__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSharedMemory);
		auto __it = __conns_QSharedMemory.find(id);
		if (__it == __conns_QSharedMemory.end()) return;
		__c = __it->second;
		__conns_QSharedMemory.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSharedMemory__Go_new(long long id, void* parent) { return new QSharedMemory__Go(id, (QObject*)parent); }
void QSharedMemory_delete(void* self) { delete (QSharedMemory*)self; }

}

#include <QtCore/qxmlstream.h>
#include <QtCore/qstring.h>
#include "gen_qxmlstreamentityresolver.h"
#include "gen_purego_cb.h"

class QXmlStreamEntityResolver__Go : public QXmlStreamEntityResolver {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QXmlStreamEntityResolver__Go(long long id, A&&... a) : QXmlStreamEntityResolver(std::forward<A>(a)...), __goId(id) {}
	QString resolveEntity(const QString& publicId, const QString& systemId) override {
		int __h = 0;
		QByteArray publicId_ba = publicId.toUtf8();
		QByteArray systemId_ba = systemId.toUtf8();
		if (__goOn[0] && cppgo_virtp_charp_constcharp_longlong_constcharp_longlong) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, const char*, long long, const char*, long long, long long*, int*))cppgo_virtp_charp_constcharp_longlong_constcharp_longlong)(__goId, -1345840861565463748LL, (char*)publicId_ba.constData(), publicId_ba.size(), (char*)systemId_ba.constData(), systemId_ba.size(), &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QXmlStreamEntityResolver::resolveEntity(publicId, systemId);
	}
	QString resolveUndeclaredEntity(const QString& name) override {
		int __h = 0;
		QByteArray name_ba = name.toUtf8();
		if (__goOn[1] && cppgo_virtp_charp_constcharp_longlong) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, const char*, long long, long long*, int*))cppgo_virtp_charp_constcharp_longlong)(__goId, -5256568367070285975LL, (char*)name_ba.constData(), name_ba.size(), &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QXmlStreamEntityResolver::resolveUndeclaredEntity(name);
	}
};

extern "C" {

void QXmlStreamEntityResolver__set_override(void* self, int idx) { ((QXmlStreamEntityResolver__Go*)self)->__goOn[idx] = 1; }
void* QXmlStreamEntityResolver_new(long long __goid) {
	return new QXmlStreamEntityResolver__Go(__goid);
}
char* QXmlStreamEntityResolver_resolveEntity(void* self, const char* publicId, long long publicId_len, const char* systemId, long long systemId_len, long long* __out_len) {
	QString publicId_qstr = QString::fromUtf8(publicId, publicId_len);
	QString systemId_qstr = QString::fromUtf8(systemId, systemId_len);
	auto&& __r = ((QXmlStreamEntityResolver*)self)->resolveEntity(publicId_qstr, systemId_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QXmlStreamEntityResolver_resolveUndeclaredEntity(void* self, const char* name, long long name_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QXmlStreamEntityResolver*)self)->resolveUndeclaredEntity(name_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QXmlStreamEntityResolver_resolveEntity_default(void* self, const char* publicId, long long publicId_len, const char* systemId, long long systemId_len, long long* __out_len) {
	QString publicId_qstr = QString::fromUtf8(publicId, publicId_len);
	QString systemId_qstr = QString::fromUtf8(systemId, systemId_len);
	auto&& __r = ((QXmlStreamEntityResolver*)self)->QXmlStreamEntityResolver::resolveEntity(publicId_qstr, systemId_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QXmlStreamEntityResolver_resolveUndeclaredEntity_default(void* self, const char* name, long long name_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QXmlStreamEntityResolver*)self)->QXmlStreamEntityResolver::resolveUndeclaredEntity(name_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QXmlStreamEntityResolver__Go_new(long long id) { return new QXmlStreamEntityResolver__Go(id); }
void QXmlStreamEntityResolver_delete(void* self) { delete (QXmlStreamEntityResolver*)self; }

}

#include <QtGui/qeventpoint.h>
#include <QtGui/qpointingdevice.h>
#include <QtCore/qpoint.h>
#include <QtGui/qvectornd.h>
#include <QtCore/qsize.h>
#include "gen_qeventpoint.h"

extern "C" {

void* QEventPoint_new() {
	return new QEventPoint();
}
void* QEventPoint_new2(int pointId, int state, void* scenePosition, void* globalPosition) {
	return new QEventPoint(pointId, (QEventPoint::State)(state), *(QPointF*)scenePosition, *(QPointF*)globalPosition);
}
void* QEventPoint_new3(void* other) {
	return new QEventPoint(*(QEventPoint*)other);
}
void* QEventPoint_new4(int id) {
	return new QEventPoint(id);
}
void* QEventPoint_new5(int id, void* device) {
	return new QEventPoint(id, (QPointingDevice*)device);
}
void* QEventPoint_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QEventPoint*)self)->operator=(*(QEventPoint*)other);
	return (void*)&(__r);
}
int QEventPoint_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QEventPoint*)self)->operator==(*(QEventPoint*)other);
	return __r;
}
int QEventPoint_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QEventPoint*)self)->operator!=(*(QEventPoint*)other);
	return __r;
}
void QEventPoint_swap(void* self, void* other) {
	((QEventPoint*)self)->swap(*(QEventPoint*)other);
}
void* QEventPoint_position(void* self) {
	auto&& __r = ((QEventPoint*)self)->position();
	return new QPointF(__r);
}
void* QEventPoint_pressPosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->pressPosition();
	return new QPointF(__r);
}
void* QEventPoint_grabPosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->grabPosition();
	return new QPointF(__r);
}
void* QEventPoint_lastPosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->lastPosition();
	return new QPointF(__r);
}
void* QEventPoint_scenePosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->scenePosition();
	return new QPointF(__r);
}
void* QEventPoint_scenePressPosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->scenePressPosition();
	return new QPointF(__r);
}
void* QEventPoint_sceneGrabPosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->sceneGrabPosition();
	return new QPointF(__r);
}
void* QEventPoint_sceneLastPosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->sceneLastPosition();
	return new QPointF(__r);
}
void* QEventPoint_globalPosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->globalPosition();
	return new QPointF(__r);
}
void* QEventPoint_globalPressPosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->globalPressPosition();
	return new QPointF(__r);
}
void* QEventPoint_globalGrabPosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->globalGrabPosition();
	return new QPointF(__r);
}
void* QEventPoint_globalLastPosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->globalLastPosition();
	return new QPointF(__r);
}
void* QEventPoint_normalizedPosition(void* self) {
	auto&& __r = ((QEventPoint*)self)->normalizedPosition();
	return new QPointF(__r);
}
void* QEventPoint_pos(void* self) {
	auto&& __r = ((QEventPoint*)self)->pos();
	return new QPointF(__r);
}
void* QEventPoint_startPos(void* self) {
	auto&& __r = ((QEventPoint*)self)->startPos();
	return new QPointF(__r);
}
void* QEventPoint_scenePos(void* self) {
	auto&& __r = ((QEventPoint*)self)->scenePos();
	return new QPointF(__r);
}
void* QEventPoint_startScenePos(void* self) {
	auto&& __r = ((QEventPoint*)self)->startScenePos();
	return new QPointF(__r);
}
void* QEventPoint_screenPos(void* self) {
	auto&& __r = ((QEventPoint*)self)->screenPos();
	return new QPointF(__r);
}
void* QEventPoint_startScreenPos(void* self) {
	auto&& __r = ((QEventPoint*)self)->startScreenPos();
	return new QPointF(__r);
}
void* QEventPoint_startNormalizedPos(void* self) {
	auto&& __r = ((QEventPoint*)self)->startNormalizedPos();
	return new QPointF(__r);
}
void* QEventPoint_normalizedPos(void* self) {
	auto&& __r = ((QEventPoint*)self)->normalizedPos();
	return new QPointF(__r);
}
void* QEventPoint_lastPos(void* self) {
	auto&& __r = ((QEventPoint*)self)->lastPos();
	return new QPointF(__r);
}
void* QEventPoint_lastScenePos(void* self) {
	auto&& __r = ((QEventPoint*)self)->lastScenePos();
	return new QPointF(__r);
}
void* QEventPoint_lastScreenPos(void* self) {
	auto&& __r = ((QEventPoint*)self)->lastScreenPos();
	return new QPointF(__r);
}
void* QEventPoint_lastNormalizedPos(void* self) {
	auto&& __r = ((QEventPoint*)self)->lastNormalizedPos();
	return new QPointF(__r);
}
void* QEventPoint_velocity(void* self) {
	auto&& __r = ((QEventPoint*)self)->velocity();
	return new QVector2D(__r);
}
int QEventPoint_state(void* self) {
	auto&& __r = ((QEventPoint*)self)->state();
	return (int)(__r);
}
void* QEventPoint_device(void* self) {
	auto&& __r = ((QEventPoint*)self)->device();
	return (void*)(__r);
}
int QEventPoint_id(void* self) {
	auto&& __r = ((QEventPoint*)self)->id();
	return __r;
}
void* QEventPoint_uniqueId(void* self) {
	auto&& __r = ((QEventPoint*)self)->uniqueId();
	return new QPointingDeviceUniqueId(__r);
}
unsigned long QEventPoint_timestamp(void* self) {
	auto&& __r = ((QEventPoint*)self)->timestamp();
	return __r;
}
unsigned long QEventPoint_lastTimestamp(void* self) {
	auto&& __r = ((QEventPoint*)self)->lastTimestamp();
	return __r;
}
unsigned long QEventPoint_pressTimestamp(void* self) {
	auto&& __r = ((QEventPoint*)self)->pressTimestamp();
	return __r;
}
double QEventPoint_timeHeld(void* self) {
	auto&& __r = ((QEventPoint*)self)->timeHeld();
	return __r;
}
double QEventPoint_pressure(void* self) {
	auto&& __r = ((QEventPoint*)self)->pressure();
	return __r;
}
double QEventPoint_rotation(void* self) {
	auto&& __r = ((QEventPoint*)self)->rotation();
	return __r;
}
void* QEventPoint_ellipseDiameters(void* self) {
	auto&& __r = ((QEventPoint*)self)->ellipseDiameters();
	return new QSizeF(__r);
}
int QEventPoint_isAccepted(void* self) {
	auto&& __r = ((QEventPoint*)self)->isAccepted();
	return __r;
}
void QEventPoint_setAccepted(void* self) {
	((QEventPoint*)self)->setAccepted();
}
void QEventPoint_setAccepted2(void* self, int accepted) {
	((QEventPoint*)self)->setAccepted(accepted);
}
void QEventPoint_delete(void* self) { delete (QEventPoint*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qvariant.h>
#include "gen_qinputmethodevent__attribute.h"

extern "C" {

void* QInputMethodEvent__Attribute_new(int typ, int s, int l, void* val) {
	return new QInputMethodEvent::Attribute((QInputMethodEvent::AttributeType)(typ), s, l, *(QVariant*)val);
}
void* QInputMethodEvent__Attribute_new2(int typ, int s, int l) {
	return new QInputMethodEvent::Attribute((QInputMethodEvent::AttributeType)(typ), s, l);
}
void* QInputMethodEvent__Attribute_new3(void* arg1) {
	return new QInputMethodEvent::Attribute(*(QInputMethodEvent::Attribute*)arg1);
}
void* QInputMethodEvent__Attribute_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QInputMethodEvent::Attribute*)self)->operator=(*(QInputMethodEvent::Attribute*)arg1);
	return (void*)&(__r);
}
int QInputMethodEvent__Attribute_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QInputMethodEvent::Attribute*)self) == (*(QInputMethodEvent::Attribute*)rhs);
	return __r;
}
int QInputMethodEvent__Attribute_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QInputMethodEvent::Attribute*)self) != (*(QInputMethodEvent::Attribute*)rhs);
	return __r;
}
int QInputMethodEvent__Attribute_type_get(void* self) {
	auto __r = ((QInputMethodEvent::Attribute*)self)->type;
	return (int)(__r);
}
void QInputMethodEvent__Attribute_type_set(void* self, int type_) {
	((QInputMethodEvent::Attribute*)self)->type = (QInputMethodEvent::AttributeType)(type_);
}
int QInputMethodEvent__Attribute_start_get(void* self) {
	auto __r = ((QInputMethodEvent::Attribute*)self)->start;
	return __r;
}
void QInputMethodEvent__Attribute_start_set(void* self, int start) {
	((QInputMethodEvent::Attribute*)self)->start = start;
}
int QInputMethodEvent__Attribute_length_get(void* self) {
	auto __r = ((QInputMethodEvent::Attribute*)self)->length;
	return __r;
}
void QInputMethodEvent__Attribute_length_set(void* self, int length) {
	((QInputMethodEvent::Attribute*)self)->length = length;
}
void* QInputMethodEvent__Attribute_value_get(void* self) {
	auto __r = ((QInputMethodEvent::Attribute*)self)->value;
	return new QVariant(__r);
}
void QInputMethodEvent__Attribute_value_set(void* self, void* value) {
	((QInputMethodEvent::Attribute*)self)->value = *(QVariant*)value;
}
void QInputMethodEvent__Attribute_delete(void* self) { delete (QInputMethodEvent::Attribute*)self; }

}

#include <QtGui/qglyphrun.h>
#include <QtGui/qrawfont.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtCore/qstring.h>
#include "gen_qglyphrun.h"

extern "C" {

void* QGlyphRun_new() {
	return new QGlyphRun();
}
void* QGlyphRun_new2(void* other) {
	return new QGlyphRun(*(QGlyphRun*)other);
}
void* QGlyphRun_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QGlyphRun*)self)->operator=(*(QGlyphRun*)other);
	return (void*)&(__r);
}
void QGlyphRun_swap(void* self, void* other) {
	((QGlyphRun*)self)->swap(*(QGlyphRun*)other);
}
void* QGlyphRun_rawFont(void* self) {
	auto&& __r = ((QGlyphRun*)self)->rawFont();
	return new QRawFont(__r);
}
void QGlyphRun_setRawFont(void* self, void* rawFont) {
	((QGlyphRun*)self)->setRawFont(*(QRawFont*)rawFont);
}
unsigned int* QGlyphRun_glyphIndexes(void* self, long long* __out_len) {
	auto&& __r = ((QGlyphRun*)self)->glyphIndexes();
	*__out_len = __r.size();
	unsigned int* __arr = (unsigned int*)malloc(sizeof(unsigned int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void QGlyphRun_setGlyphIndexes(void* self, unsigned int* glyphIndexes, long long glyphIndexes_len) {
	QList<unsigned int> glyphIndexes_list; glyphIndexes_list.reserve(glyphIndexes_len);
	for (long long __i = 0; __i < glyphIndexes_len; __i++) glyphIndexes_list.append(glyphIndexes[__i]);
	((QGlyphRun*)self)->setGlyphIndexes(glyphIndexes_list);
}
void** QGlyphRun_positions(void* self, long long* __out_len) {
	auto&& __r = ((QGlyphRun*)self)->positions();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void QGlyphRun_setPositions(void* self, void** positions, long long positions_len) {
	QList<QPointF> positions_list; positions_list.reserve(positions_len);
	for (long long __i = 0; __i < positions_len; __i++) positions_list.append(*(QPointF*)positions[__i]);
	((QGlyphRun*)self)->setPositions(positions_list);
}
void QGlyphRun_clear(void* self) {
	((QGlyphRun*)self)->clear();
}
int QGlyphRun_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QGlyphRun*)self)->operator==(*(QGlyphRun*)other);
	return __r;
}
int QGlyphRun_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QGlyphRun*)self)->operator!=(*(QGlyphRun*)other);
	return __r;
}
void QGlyphRun_setOverline(void* self, int overline) {
	((QGlyphRun*)self)->setOverline(overline);
}
int QGlyphRun_overline(void* self) {
	auto&& __r = ((QGlyphRun*)self)->overline();
	return __r;
}
void QGlyphRun_setUnderline(void* self, int underline) {
	((QGlyphRun*)self)->setUnderline(underline);
}
int QGlyphRun_underline(void* self) {
	auto&& __r = ((QGlyphRun*)self)->underline();
	return __r;
}
void QGlyphRun_setStrikeOut(void* self, int strikeOut) {
	((QGlyphRun*)self)->setStrikeOut(strikeOut);
}
int QGlyphRun_strikeOut(void* self) {
	auto&& __r = ((QGlyphRun*)self)->strikeOut();
	return __r;
}
void QGlyphRun_setRightToLeft(void* self, int on) {
	((QGlyphRun*)self)->setRightToLeft(on);
}
int QGlyphRun_isRightToLeft(void* self) {
	auto&& __r = ((QGlyphRun*)self)->isRightToLeft();
	return __r;
}
void QGlyphRun_setFlag(void* self, int flag) {
	((QGlyphRun*)self)->setFlag((QGlyphRun::GlyphRunFlag)(flag));
}
void QGlyphRun_setFlags(void* self, int flags) {
	((QGlyphRun*)self)->setFlags(QFlags<QGlyphRun::GlyphRunFlag>(flags));
}
int QGlyphRun_flags(void* self) {
	auto&& __r = ((QGlyphRun*)self)->flags();
	return (int)(__r);
}
void QGlyphRun_setBoundingRect(void* self, void* boundingRect) {
	((QGlyphRun*)self)->setBoundingRect(*(QRectF*)boundingRect);
}
void* QGlyphRun_boundingRect(void* self) {
	auto&& __r = ((QGlyphRun*)self)->boundingRect();
	return new QRectF(__r);
}
long long* QGlyphRun_stringIndexes(void* self, long long* __out_len) {
	auto&& __r = ((QGlyphRun*)self)->stringIndexes();
	*__out_len = __r.size();
	long long* __arr = (long long*)malloc(sizeof(long long) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void QGlyphRun_setStringIndexes(void* self, long long* stringIndexes, long long stringIndexes_len) {
	QList<long long> stringIndexes_list; stringIndexes_list.reserve(stringIndexes_len);
	for (long long __i = 0; __i < stringIndexes_len; __i++) stringIndexes_list.append(stringIndexes[__i]);
	((QGlyphRun*)self)->setStringIndexes(stringIndexes_list);
}
void QGlyphRun_setSourceString(void* self, const char* sourceString, long long sourceString_len) {
	QString sourceString_qstr = QString::fromUtf8(sourceString, sourceString_len);
	((QGlyphRun*)self)->setSourceString(sourceString_qstr);
}
char* QGlyphRun_sourceString(void* self, long long* __out_len) {
	auto&& __r = ((QGlyphRun*)self)->sourceString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QGlyphRun_isEmpty(void* self) {
	auto&& __r = ((QGlyphRun*)self)->isEmpty();
	return __r;
}
void QGlyphRun_setFlag2(void* self, int flag, int enabled) {
	((QGlyphRun*)self)->setFlag((QGlyphRun::GlyphRunFlag)(flag), enabled);
}
void QGlyphRun_delete(void* self) { delete (QGlyphRun*)self; }

}

#include <QtGui/qtextlayout.h>
#include <QtCore/qstring.h>
#include <QtGui/qfont.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qtextobject.h>
#include <QtGui/qrawfont.h>
#include <QtGui/qtextoption.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qpainter.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtGui/qglyphrun.h>
#include "gen_qtextlayout.h"

extern "C" {

void* QTextLayout_new() {
	return new QTextLayout();
}
void* QTextLayout_new2(const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QTextLayout(text_qstr);
}
void* QTextLayout_new3(const char* text, long long text_len, void* font) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QTextLayout(text_qstr, *(QFont*)font);
}
void* QTextLayout_new4(void* b) {
	return new QTextLayout(*(QTextBlock*)b);
}
void* QTextLayout_new5(const char* text, long long text_len, void* font, void* paintdevice) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QTextLayout(text_qstr, *(QFont*)font, (QPaintDevice*)paintdevice);
}
void QTextLayout_setFont(void* self, void* f) {
	((QTextLayout*)self)->setFont(*(QFont*)f);
}
void* QTextLayout_font(void* self) {
	auto&& __r = ((QTextLayout*)self)->font();
	return new QFont(__r);
}
void QTextLayout_setRawFont(void* self, void* rawFont) {
	((QTextLayout*)self)->setRawFont(*(QRawFont*)rawFont);
}
void QTextLayout_setText(void* self, const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	((QTextLayout*)self)->setText(string__qstr);
}
char* QTextLayout_text(void* self, long long* __out_len) {
	auto&& __r = ((QTextLayout*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextLayout_setTextOption(void* self, void* option) {
	((QTextLayout*)self)->setTextOption(*(QTextOption*)option);
}
void* QTextLayout_textOption(void* self) {
	auto&& __r = ((QTextLayout*)self)->textOption();
	return (void*)&(__r);
}
void QTextLayout_setPreeditArea(void* self, int position, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTextLayout*)self)->setPreeditArea(position, text_qstr);
}
int QTextLayout_preeditAreaPosition(void* self) {
	auto&& __r = ((QTextLayout*)self)->preeditAreaPosition();
	return __r;
}
char* QTextLayout_preeditAreaText(void* self, long long* __out_len) {
	auto&& __r = ((QTextLayout*)self)->preeditAreaText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextLayout_setFormats(void* self, void** overrides, long long overrides_len) {
	QList<QTextLayout::FormatRange> overrides_list; overrides_list.reserve(overrides_len);
	for (long long __i = 0; __i < overrides_len; __i++) overrides_list.append(*(QTextLayout::FormatRange*)overrides[__i]);
	((QTextLayout*)self)->setFormats(overrides_list);
}
void** QTextLayout_formats(void* self, long long* __out_len) {
	auto&& __r = ((QTextLayout*)self)->formats();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QTextLayout::FormatRange(__e);
	return __arr;
}
void QTextLayout_clearFormats(void* self) {
	((QTextLayout*)self)->clearFormats();
}
void QTextLayout_setCacheEnabled(void* self, int enable) {
	((QTextLayout*)self)->setCacheEnabled(enable);
}
int QTextLayout_cacheEnabled(void* self) {
	auto&& __r = ((QTextLayout*)self)->cacheEnabled();
	return __r;
}
void QTextLayout_setCursorMoveStyle(void* self, int style) {
	((QTextLayout*)self)->setCursorMoveStyle((Qt::CursorMoveStyle)(style));
}
int QTextLayout_cursorMoveStyle(void* self) {
	auto&& __r = ((QTextLayout*)self)->cursorMoveStyle();
	return (int)(__r);
}
void QTextLayout_beginLayout(void* self) {
	((QTextLayout*)self)->beginLayout();
}
void QTextLayout_endLayout(void* self) {
	((QTextLayout*)self)->endLayout();
}
void QTextLayout_clearLayout(void* self) {
	((QTextLayout*)self)->clearLayout();
}
void* QTextLayout_createLine(void* self) {
	auto&& __r = ((QTextLayout*)self)->createLine();
	return new QTextLine(__r);
}
int QTextLayout_lineCount(void* self) {
	auto&& __r = ((QTextLayout*)self)->lineCount();
	return __r;
}
void* QTextLayout_lineAt(void* self, int i) {
	auto&& __r = ((QTextLayout*)self)->lineAt(i);
	return new QTextLine(__r);
}
void* QTextLayout_lineForTextPosition(void* self, int pos) {
	auto&& __r = ((QTextLayout*)self)->lineForTextPosition(pos);
	return new QTextLine(__r);
}
int QTextLayout_isValidCursorPosition(void* self, int pos) {
	auto&& __r = ((QTextLayout*)self)->isValidCursorPosition(pos);
	return __r;
}
int QTextLayout_nextCursorPosition(void* self, int oldPos) {
	auto&& __r = ((QTextLayout*)self)->nextCursorPosition(oldPos);
	return __r;
}
int QTextLayout_previousCursorPosition(void* self, int oldPos) {
	auto&& __r = ((QTextLayout*)self)->previousCursorPosition(oldPos);
	return __r;
}
int QTextLayout_leftCursorPosition(void* self, int oldPos) {
	auto&& __r = ((QTextLayout*)self)->leftCursorPosition(oldPos);
	return __r;
}
int QTextLayout_rightCursorPosition(void* self, int oldPos) {
	auto&& __r = ((QTextLayout*)self)->rightCursorPosition(oldPos);
	return __r;
}
void QTextLayout_draw(void* self, void* p, void* pos) {
	((QTextLayout*)self)->draw((QPainter*)p, *(QPointF*)pos);
}
void QTextLayout_drawCursor(void* self, void* p, void* pos, int cursorPosition) {
	((QTextLayout*)self)->drawCursor((QPainter*)p, *(QPointF*)pos, cursorPosition);
}
void QTextLayout_drawCursor2(void* self, void* p, void* pos, int cursorPosition, int width) {
	((QTextLayout*)self)->drawCursor((QPainter*)p, *(QPointF*)pos, cursorPosition, width);
}
void* QTextLayout_position(void* self) {
	auto&& __r = ((QTextLayout*)self)->position();
	return new QPointF(__r);
}
void QTextLayout_setPosition(void* self, void* p) {
	((QTextLayout*)self)->setPosition(*(QPointF*)p);
}
void* QTextLayout_boundingRect(void* self) {
	auto&& __r = ((QTextLayout*)self)->boundingRect();
	return new QRectF(__r);
}
double QTextLayout_minimumWidth(void* self) {
	auto&& __r = ((QTextLayout*)self)->minimumWidth();
	return __r;
}
double QTextLayout_maximumWidth(void* self) {
	auto&& __r = ((QTextLayout*)self)->maximumWidth();
	return __r;
}
void** QTextLayout_glyphRuns(void* self, int from, int length, int flags, long long* __out_len) {
	auto&& __r = ((QTextLayout*)self)->glyphRuns(from, length, QFlags<QTextLayout::GlyphRunRetrievalFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QGlyphRun(__e);
	return __arr;
}
void** QTextLayout_glyphRuns2(void* self, long long* __out_len) {
	auto&& __r = ((QTextLayout*)self)->glyphRuns();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QGlyphRun(__e);
	return __arr;
}
void QTextLayout_setFlags(void* self, int flags) {
	((QTextLayout*)self)->setFlags(flags);
}
int QTextLayout_nextCursorPosition2(void* self, int oldPos, int mode) {
	auto&& __r = ((QTextLayout*)self)->nextCursorPosition(oldPos, (QTextLayout::CursorMode)(mode));
	return __r;
}
int QTextLayout_previousCursorPosition2(void* self, int oldPos, int mode) {
	auto&& __r = ((QTextLayout*)self)->previousCursorPosition(oldPos, (QTextLayout::CursorMode)(mode));
	return __r;
}
void QTextLayout_draw2(void* self, void* p, void* pos, void** selections, long long selections_len) {
	QList<QTextLayout::FormatRange> selections_list; selections_list.reserve(selections_len);
	for (long long __i = 0; __i < selections_len; __i++) selections_list.append(*(QTextLayout::FormatRange*)selections[__i]);
	((QTextLayout*)self)->draw((QPainter*)p, *(QPointF*)pos, selections_list);
}
void QTextLayout_draw3(void* self, void* p, void* pos, void** selections, long long selections_len, void* clip) {
	QList<QTextLayout::FormatRange> selections_list; selections_list.reserve(selections_len);
	for (long long __i = 0; __i < selections_len; __i++) selections_list.append(*(QTextLayout::FormatRange*)selections[__i]);
	((QTextLayout*)self)->draw((QPainter*)p, *(QPointF*)pos, selections_list, *(QRectF*)clip);
}
void** QTextLayout_glyphRuns3(void* self, int from, long long* __out_len) {
	auto&& __r = ((QTextLayout*)self)->glyphRuns(from);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QGlyphRun(__e);
	return __arr;
}
void** QTextLayout_glyphRuns4(void* self, int from, int length, long long* __out_len) {
	auto&& __r = ((QTextLayout*)self)->glyphRuns(from, length);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QGlyphRun(__e);
	return __arr;
}
void QTextLayout_delete(void* self) { delete (QTextLayout*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtGui/qaccessible_base.h>
#include "gen_qaccessiblestatechangeevent.h"

extern "C" {

void* QAccessibleStateChangeEvent_base(void* self) { return (void*)static_cast<QAccessibleEvent*>((QAccessibleStateChangeEvent*)self); }
void* QAccessibleStateChangeEvent_new(void* obj, void* state) {
	return new QAccessibleStateChangeEvent((QObject*)obj, *(QAccessible::State*)state);
}
void* QAccessibleStateChangeEvent_new2(void* iface, void* state) {
	return new QAccessibleStateChangeEvent((QAccessibleInterface*)iface, *(QAccessible::State*)state);
}
void* QAccessibleStateChangeEvent_changedStates(void* self) {
	auto&& __r = ((QAccessibleStateChangeEvent*)self)->changedStates();
	return new QAccessible::State(__r);
}
void QAccessibleStateChangeEvent_delete(void* self) { delete (QAccessibleStateChangeEvent*)self; }

}

#include <QtGui/qcolorspace.h>
#include <QtCore/qpoint.h>
#include <QtCore/qstring.h>
#include <QtCore/qbytearray.h>
#include <QtGui/qcolortransform.h>
#include <QtCore/qvariant.h>
#include "gen_qcolorspace.h"

extern "C" {

void* QColorSpace_new() {
	return new QColorSpace();
}
void* QColorSpace_new2(int namedColorSpace) {
	return new QColorSpace((QColorSpace::NamedColorSpace)(namedColorSpace));
}
void* QColorSpace_new3(void* whitePoint, int transferFunction) {
	return new QColorSpace(*(QPointF*)whitePoint, (QColorSpace::TransferFunction)(transferFunction));
}
void* QColorSpace_new4(void* whitePoint, unsigned short* transferFunctionTable, long long transferFunctionTable_len) {
	QList<unsigned short> transferFunctionTable_list; transferFunctionTable_list.reserve(transferFunctionTable_len);
	for (long long __i = 0; __i < transferFunctionTable_len; __i++) transferFunctionTable_list.append(transferFunctionTable[__i]);
	return new QColorSpace(*(QPointF*)whitePoint, transferFunctionTable_list);
}
void* QColorSpace_new5(int primaries, int transferFunction) {
	return new QColorSpace((QColorSpace::Primaries)(primaries), (QColorSpace::TransferFunction)(transferFunction));
}
void* QColorSpace_new6(int primaries, float gamma) {
	return new QColorSpace((QColorSpace::Primaries)(primaries), gamma);
}
void* QColorSpace_new7(int primaries, unsigned short* transferFunctionTable, long long transferFunctionTable_len) {
	QList<unsigned short> transferFunctionTable_list; transferFunctionTable_list.reserve(transferFunctionTable_len);
	for (long long __i = 0; __i < transferFunctionTable_len; __i++) transferFunctionTable_list.append(transferFunctionTable[__i]);
	return new QColorSpace((QColorSpace::Primaries)(primaries), transferFunctionTable_list);
}
void* QColorSpace_new8(void* whitePoint, void* redPoint, void* greenPoint, void* bluePoint, int transferFunction) {
	return new QColorSpace(*(QPointF*)whitePoint, *(QPointF*)redPoint, *(QPointF*)greenPoint, *(QPointF*)bluePoint, (QColorSpace::TransferFunction)(transferFunction));
}
void* QColorSpace_new9(void* primaryPoints, int transferFunction) {
	return new QColorSpace(*(QColorSpace::PrimaryPoints*)primaryPoints, (QColorSpace::TransferFunction)(transferFunction));
}
void* QColorSpace_new10(void* whitePoint, void* redPoint, void* greenPoint, void* bluePoint, unsigned short* transferFunctionTable, long long transferFunctionTable_len) {
	QList<unsigned short> transferFunctionTable_list; transferFunctionTable_list.reserve(transferFunctionTable_len);
	for (long long __i = 0; __i < transferFunctionTable_len; __i++) transferFunctionTable_list.append(transferFunctionTable[__i]);
	return new QColorSpace(*(QPointF*)whitePoint, *(QPointF*)redPoint, *(QPointF*)greenPoint, *(QPointF*)bluePoint, transferFunctionTable_list);
}
void* QColorSpace_new11(void* whitePoint, void* redPoint, void* greenPoint, void* bluePoint, unsigned short* redTransferFunctionTable, long long redTransferFunctionTable_len, unsigned short* greenTransferFunctionTable, long long greenTransferFunctionTable_len, unsigned short* blueTransferFunctionTable, long long blueTransferFunctionTable_len) {
	QList<unsigned short> redTransferFunctionTable_list; redTransferFunctionTable_list.reserve(redTransferFunctionTable_len);
	for (long long __i = 0; __i < redTransferFunctionTable_len; __i++) redTransferFunctionTable_list.append(redTransferFunctionTable[__i]);
	QList<unsigned short> greenTransferFunctionTable_list; greenTransferFunctionTable_list.reserve(greenTransferFunctionTable_len);
	for (long long __i = 0; __i < greenTransferFunctionTable_len; __i++) greenTransferFunctionTable_list.append(greenTransferFunctionTable[__i]);
	QList<unsigned short> blueTransferFunctionTable_list; blueTransferFunctionTable_list.reserve(blueTransferFunctionTable_len);
	for (long long __i = 0; __i < blueTransferFunctionTable_len; __i++) blueTransferFunctionTable_list.append(blueTransferFunctionTable[__i]);
	return new QColorSpace(*(QPointF*)whitePoint, *(QPointF*)redPoint, *(QPointF*)greenPoint, *(QPointF*)bluePoint, redTransferFunctionTable_list, greenTransferFunctionTable_list, blueTransferFunctionTable_list);
}
void* QColorSpace_new12(void* colorSpace) {
	return new QColorSpace(*(QColorSpace*)colorSpace);
}
void* QColorSpace_new13(void* whitePoint, int transferFunction, float gamma) {
	return new QColorSpace(*(QPointF*)whitePoint, (QColorSpace::TransferFunction)(transferFunction), gamma);
}
void* QColorSpace_new14(int primaries, int transferFunction, float gamma) {
	return new QColorSpace((QColorSpace::Primaries)(primaries), (QColorSpace::TransferFunction)(transferFunction), gamma);
}
void* QColorSpace_new15(void* whitePoint, void* redPoint, void* greenPoint, void* bluePoint, int transferFunction, float gamma) {
	return new QColorSpace(*(QPointF*)whitePoint, *(QPointF*)redPoint, *(QPointF*)greenPoint, *(QPointF*)bluePoint, (QColorSpace::TransferFunction)(transferFunction), gamma);
}
void* QColorSpace_new16(void* primaryPoints, int transferFunction, float gamma) {
	return new QColorSpace(*(QColorSpace::PrimaryPoints*)primaryPoints, (QColorSpace::TransferFunction)(transferFunction), gamma);
}
void* QColorSpace_OperatorAssign(void* self, void* colorSpace) {
	auto&& __r = ((QColorSpace*)self)->operator=(*(QColorSpace*)colorSpace);
	return (void*)&(__r);
}
void QColorSpace_swap(void* self, void* colorSpace) {
	((QColorSpace*)self)->swap(*(QColorSpace*)colorSpace);
}
int QColorSpace_primaries(void* self) {
	auto&& __r = ((QColorSpace*)self)->primaries();
	return (int)(__r);
}
int QColorSpace_transferFunction(void* self) {
	auto&& __r = ((QColorSpace*)self)->transferFunction();
	return (int)(__r);
}
float QColorSpace_gamma(void* self) {
	auto&& __r = ((QColorSpace*)self)->gamma();
	return __r;
}
char* QColorSpace_description(void* self, long long* __out_len) {
	auto&& __r = ((QColorSpace*)self)->description();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QColorSpace_setDescription(void* self, const char* description, long long description_len) {
	QString description_qstr = QString::fromUtf8(description, description_len);
	((QColorSpace*)self)->setDescription(description_qstr);
}
void QColorSpace_setTransferFunction(void* self, int transferFunction) {
	((QColorSpace*)self)->setTransferFunction((QColorSpace::TransferFunction)(transferFunction));
}
void QColorSpace_setTransferFunction2(void* self, unsigned short* transferFunctionTable, long long transferFunctionTable_len) {
	QList<unsigned short> transferFunctionTable_list; transferFunctionTable_list.reserve(transferFunctionTable_len);
	for (long long __i = 0; __i < transferFunctionTable_len; __i++) transferFunctionTable_list.append(transferFunctionTable[__i]);
	((QColorSpace*)self)->setTransferFunction(transferFunctionTable_list);
}
void QColorSpace_setTransferFunctions(void* self, unsigned short* redTransferFunctionTable, long long redTransferFunctionTable_len, unsigned short* greenTransferFunctionTable, long long greenTransferFunctionTable_len, unsigned short* blueTransferFunctionTable, long long blueTransferFunctionTable_len) {
	QList<unsigned short> redTransferFunctionTable_list; redTransferFunctionTable_list.reserve(redTransferFunctionTable_len);
	for (long long __i = 0; __i < redTransferFunctionTable_len; __i++) redTransferFunctionTable_list.append(redTransferFunctionTable[__i]);
	QList<unsigned short> greenTransferFunctionTable_list; greenTransferFunctionTable_list.reserve(greenTransferFunctionTable_len);
	for (long long __i = 0; __i < greenTransferFunctionTable_len; __i++) greenTransferFunctionTable_list.append(greenTransferFunctionTable[__i]);
	QList<unsigned short> blueTransferFunctionTable_list; blueTransferFunctionTable_list.reserve(blueTransferFunctionTable_len);
	for (long long __i = 0; __i < blueTransferFunctionTable_len; __i++) blueTransferFunctionTable_list.append(blueTransferFunctionTable[__i]);
	((QColorSpace*)self)->setTransferFunctions(redTransferFunctionTable_list, greenTransferFunctionTable_list, blueTransferFunctionTable_list);
}
void* QColorSpace_withTransferFunction(void* self, int transferFunction) {
	auto&& __r = ((QColorSpace*)self)->withTransferFunction((QColorSpace::TransferFunction)(transferFunction));
	return new QColorSpace(__r);
}
void* QColorSpace_withTransferFunction2(void* self, unsigned short* transferFunctionTable, long long transferFunctionTable_len) {
	QList<unsigned short> transferFunctionTable_list; transferFunctionTable_list.reserve(transferFunctionTable_len);
	for (long long __i = 0; __i < transferFunctionTable_len; __i++) transferFunctionTable_list.append(transferFunctionTable[__i]);
	auto&& __r = ((QColorSpace*)self)->withTransferFunction(transferFunctionTable_list);
	return new QColorSpace(__r);
}
void* QColorSpace_withTransferFunctions(void* self, unsigned short* redTransferFunctionTable, long long redTransferFunctionTable_len, unsigned short* greenTransferFunctionTable, long long greenTransferFunctionTable_len, unsigned short* blueTransferFunctionTable, long long blueTransferFunctionTable_len) {
	QList<unsigned short> redTransferFunctionTable_list; redTransferFunctionTable_list.reserve(redTransferFunctionTable_len);
	for (long long __i = 0; __i < redTransferFunctionTable_len; __i++) redTransferFunctionTable_list.append(redTransferFunctionTable[__i]);
	QList<unsigned short> greenTransferFunctionTable_list; greenTransferFunctionTable_list.reserve(greenTransferFunctionTable_len);
	for (long long __i = 0; __i < greenTransferFunctionTable_len; __i++) greenTransferFunctionTable_list.append(greenTransferFunctionTable[__i]);
	QList<unsigned short> blueTransferFunctionTable_list; blueTransferFunctionTable_list.reserve(blueTransferFunctionTable_len);
	for (long long __i = 0; __i < blueTransferFunctionTable_len; __i++) blueTransferFunctionTable_list.append(blueTransferFunctionTable[__i]);
	auto&& __r = ((QColorSpace*)self)->withTransferFunctions(redTransferFunctionTable_list, greenTransferFunctionTable_list, blueTransferFunctionTable_list);
	return new QColorSpace(__r);
}
void QColorSpace_setPrimaries(void* self, int primariesId) {
	((QColorSpace*)self)->setPrimaries((QColorSpace::Primaries)(primariesId));
}
void QColorSpace_setPrimaries2(void* self, void* whitePoint, void* redPoint, void* greenPoint, void* bluePoint) {
	((QColorSpace*)self)->setPrimaries(*(QPointF*)whitePoint, *(QPointF*)redPoint, *(QPointF*)greenPoint, *(QPointF*)bluePoint);
}
void QColorSpace_setWhitePoint(void* self, void* whitePoint) {
	((QColorSpace*)self)->setWhitePoint(*(QPointF*)whitePoint);
}
void* QColorSpace_whitePoint(void* self) {
	auto&& __r = ((QColorSpace*)self)->whitePoint();
	return new QPointF(__r);
}
void QColorSpace_setPrimaryPoints(void* self, void* primaryPoints) {
	((QColorSpace*)self)->setPrimaryPoints(*(QColorSpace::PrimaryPoints*)primaryPoints);
}
void* QColorSpace_primaryPoints(void* self) {
	auto&& __r = ((QColorSpace*)self)->primaryPoints();
	return new QColorSpace::PrimaryPoints(__r);
}
int QColorSpace_transformModel(void* self) {
	auto&& __r = ((QColorSpace*)self)->transformModel();
	return (int)(__r);
}
int QColorSpace_colorModel(void* self) {
	auto&& __r = ((QColorSpace*)self)->colorModel();
	return (int)(__r);
}
void QColorSpace_detach(void* self) {
	((QColorSpace*)self)->detach();
}
int QColorSpace_isValid(void* self) {
	auto&& __r = ((QColorSpace*)self)->isValid();
	return __r;
}
int QColorSpace_isValidTarget(void* self) {
	auto&& __r = ((QColorSpace*)self)->isValidTarget();
	return __r;
}
void* QColorSpace_fromIccProfile(void* iccProfile) {
	auto&& __r = QColorSpace::fromIccProfile(*(QByteArray*)iccProfile);
	return new QColorSpace(__r);
}
void* QColorSpace_iccProfile(void* self) {
	auto&& __r = ((QColorSpace*)self)->iccProfile();
	return new QByteArray(__r);
}
void* QColorSpace_transformationToColorSpace(void* self, void* colorspace) {
	auto&& __r = ((QColorSpace*)self)->transformationToColorSpace(*(QColorSpace*)colorspace);
	return new QColorTransform(__r);
}
void* QColorSpace_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QColorSpace*)self);
	return new QVariant(__r);
}
void QColorSpace_setTransferFunction3(void* self, int transferFunction, float gamma) {
	((QColorSpace*)self)->setTransferFunction((QColorSpace::TransferFunction)(transferFunction), gamma);
}
void* QColorSpace_withTransferFunction3(void* self, int transferFunction, float gamma) {
	auto&& __r = ((QColorSpace*)self)->withTransferFunction((QColorSpace::TransferFunction)(transferFunction), gamma);
	return new QColorSpace(__r);
}
int QColorSpace_OperatorEqual(void* self, void* colorSpace2) {
	auto&& __r = (*(QColorSpace*)self) == (*(QColorSpace*)colorSpace2);
	return __r;
}
int QColorSpace_OperatorNotEqual(void* self, void* colorSpace2) {
	auto&& __r = (*(QColorSpace*)self) != (*(QColorSpace*)colorSpace2);
	return __r;
}
void QColorSpace_delete(void* self) { delete (QColorSpace*)self; }

}

#include <QtGui/qopenglextrafunctions.h>
#include <QtGui/qopenglfunctions.h>
#include <QtGui/qopenglcontext.h>
#include "gen_qopenglextrafunctions.h"

extern "C" {

void* QOpenGLExtraFunctions_base(void* self) { return (void*)static_cast<QOpenGLFunctions*>((QOpenGLExtraFunctions*)self); }
void* QOpenGLExtraFunctions_new() {
	return new QOpenGLExtraFunctions();
}
void* QOpenGLExtraFunctions_new2(void* context) {
	return new QOpenGLExtraFunctions((QOpenGLContext*)context);
}
void* QOpenGLExtraFunctions_new3(void* arg1) {
	return new QOpenGLExtraFunctions(*(QOpenGLExtraFunctions*)arg1);
}
void QOpenGLExtraFunctions_glReadBuffer(void* self, unsigned int mode) {
	((QOpenGLExtraFunctions*)self)->glReadBuffer(mode);
}
void QOpenGLExtraFunctions_glCopyTexSubImage3D(void* self, unsigned int target, int level, int xoffset, int yoffset, int zoffset, int x, int y, int width, int height) {
	((QOpenGLExtraFunctions*)self)->glCopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}
void QOpenGLExtraFunctions_glGenQueries(void* self, int n, unsigned int* ids) {
	GLuint ids_local{};
	((QOpenGLExtraFunctions*)self)->glGenQueries(n, &ids_local);
	*ids = ids_local;
}
unsigned char QOpenGLExtraFunctions_glIsQuery(void* self, unsigned int id) {
	auto&& __r = ((QOpenGLExtraFunctions*)self)->glIsQuery(id);
	return __r;
}
void QOpenGLExtraFunctions_glBeginQuery(void* self, unsigned int target, unsigned int id) {
	((QOpenGLExtraFunctions*)self)->glBeginQuery(target, id);
}
void QOpenGLExtraFunctions_glEndQuery(void* self, unsigned int target) {
	((QOpenGLExtraFunctions*)self)->glEndQuery(target);
}
void QOpenGLExtraFunctions_glGetQueryiv(void* self, unsigned int target, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetQueryiv(target, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetQueryObjectuiv(void* self, unsigned int id, unsigned int pname, unsigned int* params) {
	GLuint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetQueryObjectuiv(id, pname, &params_local);
	*params = params_local;
}
unsigned char QOpenGLExtraFunctions_glUnmapBuffer(void* self, unsigned int target) {
	auto&& __r = ((QOpenGLExtraFunctions*)self)->glUnmapBuffer(target);
	return __r;
}
void QOpenGLExtraFunctions_glBlitFramebuffer(void* self, int srcX0, int srcY0, int srcX1, int srcY1, int dstX0, int dstY0, int dstX1, int dstY1, unsigned int mask, unsigned int filter) {
	((QOpenGLExtraFunctions*)self)->glBlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
}
void QOpenGLExtraFunctions_glRenderbufferStorageMultisample(void* self, unsigned int target, int samples, unsigned int internalformat, int width, int height) {
	((QOpenGLExtraFunctions*)self)->glRenderbufferStorageMultisample(target, samples, internalformat, width, height);
}
void QOpenGLExtraFunctions_glFramebufferTextureLayer(void* self, unsigned int target, unsigned int attachment, unsigned int texture, int level, int layer) {
	((QOpenGLExtraFunctions*)self)->glFramebufferTextureLayer(target, attachment, texture, level, layer);
}
void QOpenGLExtraFunctions_glFlushMappedBufferRange(void* self, unsigned int target, long long offset, long long length) {
	((QOpenGLExtraFunctions*)self)->glFlushMappedBufferRange(target, offset, length);
}
void QOpenGLExtraFunctions_glBindVertexArray(void* self, unsigned int array) {
	((QOpenGLExtraFunctions*)self)->glBindVertexArray(array);
}
void QOpenGLExtraFunctions_glGenVertexArrays(void* self, int n, unsigned int* arrays) {
	GLuint arrays_local{};
	((QOpenGLExtraFunctions*)self)->glGenVertexArrays(n, &arrays_local);
	*arrays = arrays_local;
}
unsigned char QOpenGLExtraFunctions_glIsVertexArray(void* self, unsigned int array) {
	auto&& __r = ((QOpenGLExtraFunctions*)self)->glIsVertexArray(array);
	return __r;
}
void QOpenGLExtraFunctions_glGetIntegeri_v(void* self, unsigned int target, unsigned int index, int* data) {
	GLint data_local{};
	((QOpenGLExtraFunctions*)self)->glGetIntegeri_v(target, index, &data_local);
	*data = data_local;
}
void QOpenGLExtraFunctions_glBeginTransformFeedback(void* self, unsigned int primitiveMode) {
	((QOpenGLExtraFunctions*)self)->glBeginTransformFeedback(primitiveMode);
}
void QOpenGLExtraFunctions_glEndTransformFeedback(void* self) {
	((QOpenGLExtraFunctions*)self)->glEndTransformFeedback();
}
void QOpenGLExtraFunctions_glBindBufferRange(void* self, unsigned int target, unsigned int index, unsigned int buffer, long long offset, long long size) {
	((QOpenGLExtraFunctions*)self)->glBindBufferRange(target, index, buffer, offset, size);
}
void QOpenGLExtraFunctions_glBindBufferBase(void* self, unsigned int target, unsigned int index, unsigned int buffer) {
	((QOpenGLExtraFunctions*)self)->glBindBufferBase(target, index, buffer);
}
void QOpenGLExtraFunctions_glGetTransformFeedbackVarying(void* self, unsigned int program, unsigned int index, int bufSize, int* length, int* size, unsigned int* type_, char* name) {
	GLsizei length_local{};
	GLsizei size_local{};
	GLenum type__local{};
	GLchar name_local{};
	((QOpenGLExtraFunctions*)self)->glGetTransformFeedbackVarying(program, index, bufSize, &length_local, &size_local, &type__local, &name_local);
	*length = length_local;
	*size = size_local;
	*type_ = type__local;
	*name = name_local;
}
void QOpenGLExtraFunctions_glGetVertexAttribIiv(void* self, unsigned int index, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetVertexAttribIiv(index, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetVertexAttribIuiv(void* self, unsigned int index, unsigned int pname, unsigned int* params) {
	GLuint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetVertexAttribIuiv(index, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glVertexAttribI4i(void* self, unsigned int index, int x, int y, int z, int w) {
	((QOpenGLExtraFunctions*)self)->glVertexAttribI4i(index, x, y, z, w);
}
void QOpenGLExtraFunctions_glVertexAttribI4ui(void* self, unsigned int index, unsigned int x, unsigned int y, unsigned int z, unsigned int w) {
	((QOpenGLExtraFunctions*)self)->glVertexAttribI4ui(index, x, y, z, w);
}
void QOpenGLExtraFunctions_glGetUniformuiv(void* self, unsigned int program, int location, unsigned int* params) {
	GLuint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetUniformuiv(program, location, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glUniform1ui(void* self, int location, unsigned int v0) {
	((QOpenGLExtraFunctions*)self)->glUniform1ui(location, v0);
}
void QOpenGLExtraFunctions_glUniform2ui(void* self, int location, unsigned int v0, unsigned int v1) {
	((QOpenGLExtraFunctions*)self)->glUniform2ui(location, v0, v1);
}
void QOpenGLExtraFunctions_glUniform3ui(void* self, int location, unsigned int v0, unsigned int v1, unsigned int v2) {
	((QOpenGLExtraFunctions*)self)->glUniform3ui(location, v0, v1, v2);
}
void QOpenGLExtraFunctions_glUniform4ui(void* self, int location, unsigned int v0, unsigned int v1, unsigned int v2, unsigned int v3) {
	((QOpenGLExtraFunctions*)self)->glUniform4ui(location, v0, v1, v2, v3);
}
void QOpenGLExtraFunctions_glClearBufferfi(void* self, unsigned int buffer, int drawbuffer, float depth, int stencil) {
	((QOpenGLExtraFunctions*)self)->glClearBufferfi(buffer, drawbuffer, depth, stencil);
}
void QOpenGLExtraFunctions_glCopyBufferSubData(void* self, unsigned int readTarget, unsigned int writeTarget, long long readOffset, long long writeOffset, long long size) {
	((QOpenGLExtraFunctions*)self)->glCopyBufferSubData(readTarget, writeTarget, readOffset, writeOffset, size);
}
void QOpenGLExtraFunctions_glGetActiveUniformBlockiv(void* self, unsigned int program, unsigned int uniformBlockIndex, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetActiveUniformBlockiv(program, uniformBlockIndex, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetActiveUniformBlockName(void* self, unsigned int program, unsigned int uniformBlockIndex, int bufSize, int* length, char* uniformBlockName) {
	GLsizei length_local{};
	GLchar uniformBlockName_local{};
	((QOpenGLExtraFunctions*)self)->glGetActiveUniformBlockName(program, uniformBlockIndex, bufSize, &length_local, &uniformBlockName_local);
	*length = length_local;
	*uniformBlockName = uniformBlockName_local;
}
void QOpenGLExtraFunctions_glUniformBlockBinding(void* self, unsigned int program, unsigned int uniformBlockIndex, unsigned int uniformBlockBinding) {
	((QOpenGLExtraFunctions*)self)->glUniformBlockBinding(program, uniformBlockIndex, uniformBlockBinding);
}
void QOpenGLExtraFunctions_glDrawArraysInstanced(void* self, unsigned int mode, int first, int count, int instancecount) {
	((QOpenGLExtraFunctions*)self)->glDrawArraysInstanced(mode, first, count, instancecount);
}
void QOpenGLExtraFunctions_glGetInteger64v(void* self, unsigned int pname, long long* data) {
	GLint64 data_local{};
	((QOpenGLExtraFunctions*)self)->glGetInteger64v(pname, &data_local);
	*data = data_local;
}
void QOpenGLExtraFunctions_glGetInteger64i_v(void* self, unsigned int target, unsigned int index, long long* data) {
	GLint64 data_local{};
	((QOpenGLExtraFunctions*)self)->glGetInteger64i_v(target, index, &data_local);
	*data = data_local;
}
void QOpenGLExtraFunctions_glGetBufferParameteri64v(void* self, unsigned int target, unsigned int pname, long long* params) {
	GLint64 params_local{};
	((QOpenGLExtraFunctions*)self)->glGetBufferParameteri64v(target, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGenSamplers(void* self, int count, unsigned int* samplers) {
	GLuint samplers_local{};
	((QOpenGLExtraFunctions*)self)->glGenSamplers(count, &samplers_local);
	*samplers = samplers_local;
}
unsigned char QOpenGLExtraFunctions_glIsSampler(void* self, unsigned int sampler) {
	auto&& __r = ((QOpenGLExtraFunctions*)self)->glIsSampler(sampler);
	return __r;
}
void QOpenGLExtraFunctions_glBindSampler(void* self, unsigned int unit, unsigned int sampler) {
	((QOpenGLExtraFunctions*)self)->glBindSampler(unit, sampler);
}
void QOpenGLExtraFunctions_glSamplerParameteri(void* self, unsigned int sampler, unsigned int pname, int param) {
	((QOpenGLExtraFunctions*)self)->glSamplerParameteri(sampler, pname, param);
}
void QOpenGLExtraFunctions_glSamplerParameterf(void* self, unsigned int sampler, unsigned int pname, float param) {
	((QOpenGLExtraFunctions*)self)->glSamplerParameterf(sampler, pname, param);
}
void QOpenGLExtraFunctions_glGetSamplerParameteriv(void* self, unsigned int sampler, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetSamplerParameteriv(sampler, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetSamplerParameterfv(void* self, unsigned int sampler, unsigned int pname, float* params) {
	GLfloat params_local{};
	((QOpenGLExtraFunctions*)self)->glGetSamplerParameterfv(sampler, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glVertexAttribDivisor(void* self, unsigned int index, unsigned int divisor) {
	((QOpenGLExtraFunctions*)self)->glVertexAttribDivisor(index, divisor);
}
void QOpenGLExtraFunctions_glBindTransformFeedback(void* self, unsigned int target, unsigned int id) {
	((QOpenGLExtraFunctions*)self)->glBindTransformFeedback(target, id);
}
void QOpenGLExtraFunctions_glGenTransformFeedbacks(void* self, int n, unsigned int* ids) {
	GLuint ids_local{};
	((QOpenGLExtraFunctions*)self)->glGenTransformFeedbacks(n, &ids_local);
	*ids = ids_local;
}
unsigned char QOpenGLExtraFunctions_glIsTransformFeedback(void* self, unsigned int id) {
	auto&& __r = ((QOpenGLExtraFunctions*)self)->glIsTransformFeedback(id);
	return __r;
}
void QOpenGLExtraFunctions_glPauseTransformFeedback(void* self) {
	((QOpenGLExtraFunctions*)self)->glPauseTransformFeedback();
}
void QOpenGLExtraFunctions_glResumeTransformFeedback(void* self) {
	((QOpenGLExtraFunctions*)self)->glResumeTransformFeedback();
}
void QOpenGLExtraFunctions_glProgramParameteri(void* self, unsigned int program, unsigned int pname, int value) {
	((QOpenGLExtraFunctions*)self)->glProgramParameteri(program, pname, value);
}
void QOpenGLExtraFunctions_glTexStorage2D(void* self, unsigned int target, int levels, unsigned int internalformat, int width, int height) {
	((QOpenGLExtraFunctions*)self)->glTexStorage2D(target, levels, internalformat, width, height);
}
void QOpenGLExtraFunctions_glTexStorage3D(void* self, unsigned int target, int levels, unsigned int internalformat, int width, int height, int depth) {
	((QOpenGLExtraFunctions*)self)->glTexStorage3D(target, levels, internalformat, width, height, depth);
}
void QOpenGLExtraFunctions_glGetInternalformativ(void* self, unsigned int target, unsigned int internalformat, unsigned int pname, int bufSize, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetInternalformativ(target, internalformat, pname, bufSize, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glDispatchCompute(void* self, unsigned int num_groups_x, unsigned int num_groups_y, unsigned int num_groups_z) {
	((QOpenGLExtraFunctions*)self)->glDispatchCompute(num_groups_x, num_groups_y, num_groups_z);
}
void QOpenGLExtraFunctions_glDispatchComputeIndirect(void* self, long long indirect) {
	((QOpenGLExtraFunctions*)self)->glDispatchComputeIndirect(indirect);
}
void QOpenGLExtraFunctions_glFramebufferParameteri(void* self, unsigned int target, unsigned int pname, int param) {
	((QOpenGLExtraFunctions*)self)->glFramebufferParameteri(target, pname, param);
}
void QOpenGLExtraFunctions_glGetFramebufferParameteriv(void* self, unsigned int target, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetFramebufferParameteriv(target, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetProgramInterfaceiv(void* self, unsigned int program, unsigned int programInterface, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetProgramInterfaceiv(program, programInterface, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetProgramResourceName(void* self, unsigned int program, unsigned int programInterface, unsigned int index, int bufSize, int* length, char* name) {
	GLsizei length_local{};
	GLchar name_local{};
	((QOpenGLExtraFunctions*)self)->glGetProgramResourceName(program, programInterface, index, bufSize, &length_local, &name_local);
	*length = length_local;
	*name = name_local;
}
void QOpenGLExtraFunctions_glUseProgramStages(void* self, unsigned int pipeline, unsigned int stages, unsigned int program) {
	((QOpenGLExtraFunctions*)self)->glUseProgramStages(pipeline, stages, program);
}
void QOpenGLExtraFunctions_glActiveShaderProgram(void* self, unsigned int pipeline, unsigned int program) {
	((QOpenGLExtraFunctions*)self)->glActiveShaderProgram(pipeline, program);
}
void QOpenGLExtraFunctions_glBindProgramPipeline(void* self, unsigned int pipeline) {
	((QOpenGLExtraFunctions*)self)->glBindProgramPipeline(pipeline);
}
void QOpenGLExtraFunctions_glGenProgramPipelines(void* self, int n, unsigned int* pipelines) {
	GLuint pipelines_local{};
	((QOpenGLExtraFunctions*)self)->glGenProgramPipelines(n, &pipelines_local);
	*pipelines = pipelines_local;
}
unsigned char QOpenGLExtraFunctions_glIsProgramPipeline(void* self, unsigned int pipeline) {
	auto&& __r = ((QOpenGLExtraFunctions*)self)->glIsProgramPipeline(pipeline);
	return __r;
}
void QOpenGLExtraFunctions_glGetProgramPipelineiv(void* self, unsigned int pipeline, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetProgramPipelineiv(pipeline, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glProgramUniform1i(void* self, unsigned int program, int location, int v0) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform1i(program, location, v0);
}
void QOpenGLExtraFunctions_glProgramUniform2i(void* self, unsigned int program, int location, int v0, int v1) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform2i(program, location, v0, v1);
}
void QOpenGLExtraFunctions_glProgramUniform3i(void* self, unsigned int program, int location, int v0, int v1, int v2) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform3i(program, location, v0, v1, v2);
}
void QOpenGLExtraFunctions_glProgramUniform4i(void* self, unsigned int program, int location, int v0, int v1, int v2, int v3) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform4i(program, location, v0, v1, v2, v3);
}
void QOpenGLExtraFunctions_glProgramUniform1ui(void* self, unsigned int program, int location, unsigned int v0) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform1ui(program, location, v0);
}
void QOpenGLExtraFunctions_glProgramUniform2ui(void* self, unsigned int program, int location, unsigned int v0, unsigned int v1) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform2ui(program, location, v0, v1);
}
void QOpenGLExtraFunctions_glProgramUniform3ui(void* self, unsigned int program, int location, unsigned int v0, unsigned int v1, unsigned int v2) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform3ui(program, location, v0, v1, v2);
}
void QOpenGLExtraFunctions_glProgramUniform4ui(void* self, unsigned int program, int location, unsigned int v0, unsigned int v1, unsigned int v2, unsigned int v3) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform4ui(program, location, v0, v1, v2, v3);
}
void QOpenGLExtraFunctions_glProgramUniform1f(void* self, unsigned int program, int location, float v0) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform1f(program, location, v0);
}
void QOpenGLExtraFunctions_glProgramUniform2f(void* self, unsigned int program, int location, float v0, float v1) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform2f(program, location, v0, v1);
}
void QOpenGLExtraFunctions_glProgramUniform3f(void* self, unsigned int program, int location, float v0, float v1, float v2) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform3f(program, location, v0, v1, v2);
}
void QOpenGLExtraFunctions_glProgramUniform4f(void* self, unsigned int program, int location, float v0, float v1, float v2, float v3) {
	((QOpenGLExtraFunctions*)self)->glProgramUniform4f(program, location, v0, v1, v2, v3);
}
void QOpenGLExtraFunctions_glValidateProgramPipeline(void* self, unsigned int pipeline) {
	((QOpenGLExtraFunctions*)self)->glValidateProgramPipeline(pipeline);
}
void QOpenGLExtraFunctions_glGetProgramPipelineInfoLog(void* self, unsigned int pipeline, int bufSize, int* length, char* infoLog) {
	GLsizei length_local{};
	GLchar infoLog_local{};
	((QOpenGLExtraFunctions*)self)->glGetProgramPipelineInfoLog(pipeline, bufSize, &length_local, &infoLog_local);
	*length = length_local;
	*infoLog = infoLog_local;
}
void QOpenGLExtraFunctions_glBindImageTexture(void* self, unsigned int unit, unsigned int texture, int level, unsigned char layered, int layer, unsigned int access, unsigned int format) {
	((QOpenGLExtraFunctions*)self)->glBindImageTexture(unit, texture, level, layered, layer, access, format);
}
void QOpenGLExtraFunctions_glGetBooleani_v(void* self, unsigned int target, unsigned int index, unsigned char* data) {
	GLboolean data_local{};
	((QOpenGLExtraFunctions*)self)->glGetBooleani_v(target, index, &data_local);
	*data = data_local;
}
void QOpenGLExtraFunctions_glMemoryBarrier(void* self, unsigned int barriers) {
	((QOpenGLExtraFunctions*)self)->glMemoryBarrier(barriers);
}
void QOpenGLExtraFunctions_glMemoryBarrierByRegion(void* self, unsigned int barriers) {
	((QOpenGLExtraFunctions*)self)->glMemoryBarrierByRegion(barriers);
}
void QOpenGLExtraFunctions_glTexStorage2DMultisample(void* self, unsigned int target, int samples, unsigned int internalformat, int width, int height, unsigned char fixedsamplelocations) {
	((QOpenGLExtraFunctions*)self)->glTexStorage2DMultisample(target, samples, internalformat, width, height, fixedsamplelocations);
}
void QOpenGLExtraFunctions_glGetMultisamplefv(void* self, unsigned int pname, unsigned int index, float* val) {
	GLfloat val_local{};
	((QOpenGLExtraFunctions*)self)->glGetMultisamplefv(pname, index, &val_local);
	*val = val_local;
}
void QOpenGLExtraFunctions_glSampleMaski(void* self, unsigned int maskNumber, unsigned int mask) {
	((QOpenGLExtraFunctions*)self)->glSampleMaski(maskNumber, mask);
}
void QOpenGLExtraFunctions_glGetTexLevelParameteriv(void* self, unsigned int target, int level, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetTexLevelParameteriv(target, level, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetTexLevelParameterfv(void* self, unsigned int target, int level, unsigned int pname, float* params) {
	GLfloat params_local{};
	((QOpenGLExtraFunctions*)self)->glGetTexLevelParameterfv(target, level, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glBindVertexBuffer(void* self, unsigned int bindingindex, unsigned int buffer, long long offset, int stride) {
	((QOpenGLExtraFunctions*)self)->glBindVertexBuffer(bindingindex, buffer, offset, stride);
}
void QOpenGLExtraFunctions_glVertexAttribFormat(void* self, unsigned int attribindex, int size, unsigned int type_, unsigned char normalized, unsigned int relativeoffset) {
	((QOpenGLExtraFunctions*)self)->glVertexAttribFormat(attribindex, size, type_, normalized, relativeoffset);
}
void QOpenGLExtraFunctions_glVertexAttribIFormat(void* self, unsigned int attribindex, int size, unsigned int type_, unsigned int relativeoffset) {
	((QOpenGLExtraFunctions*)self)->glVertexAttribIFormat(attribindex, size, type_, relativeoffset);
}
void QOpenGLExtraFunctions_glVertexAttribBinding(void* self, unsigned int attribindex, unsigned int bindingindex) {
	((QOpenGLExtraFunctions*)self)->glVertexAttribBinding(attribindex, bindingindex);
}
void QOpenGLExtraFunctions_glVertexBindingDivisor(void* self, unsigned int bindingindex, unsigned int divisor) {
	((QOpenGLExtraFunctions*)self)->glVertexBindingDivisor(bindingindex, divisor);
}
void QOpenGLExtraFunctions_glBlendBarrier(void* self) {
	((QOpenGLExtraFunctions*)self)->glBlendBarrier();
}
void QOpenGLExtraFunctions_glCopyImageSubData(void* self, void* __ap) {
	size_t* __a = (size_t*)__ap;
	unsigned int srcName = (unsigned int)__a[0];
	unsigned int srcTarget = (unsigned int)__a[1];
	int srcLevel = (int)__a[2];
	int srcX = (int)__a[3];
	int srcY = (int)__a[4];
	int srcZ = (int)__a[5];
	unsigned int dstName = (unsigned int)__a[6];
	unsigned int dstTarget = (unsigned int)__a[7];
	int dstLevel = (int)__a[8];
	int dstX = (int)__a[9];
	int dstY = (int)__a[10];
	int dstZ = (int)__a[11];
	int srcWidth = (int)__a[12];
	int srcHeight = (int)__a[13];
	int srcDepth = (int)__a[14];
	((QOpenGLExtraFunctions*)self)->glCopyImageSubData(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);
}
unsigned int QOpenGLExtraFunctions_glGetDebugMessageLog(void* self, unsigned int count, int bufSize, unsigned int* sources, unsigned int* types, unsigned int* ids, unsigned int* severities, int* lengths, char* messageLog) {
	GLenum sources_local{};
	GLenum types_local{};
	GLuint ids_local{};
	GLenum severities_local{};
	GLsizei lengths_local{};
	GLchar messageLog_local{};
	auto&& __r = ((QOpenGLExtraFunctions*)self)->glGetDebugMessageLog(count, bufSize, &sources_local, &types_local, &ids_local, &severities_local, &lengths_local, &messageLog_local);
	*sources = sources_local;
	*types = types_local;
	*ids = ids_local;
	*severities = severities_local;
	*lengths = lengths_local;
	*messageLog = messageLog_local;
	return __r;
}
void QOpenGLExtraFunctions_glPopDebugGroup(void* self) {
	((QOpenGLExtraFunctions*)self)->glPopDebugGroup();
}
void QOpenGLExtraFunctions_glGetObjectLabel(void* self, unsigned int identifier, unsigned int name, int bufSize, int* length, char* label) {
	GLsizei length_local{};
	GLchar label_local{};
	((QOpenGLExtraFunctions*)self)->glGetObjectLabel(identifier, name, bufSize, &length_local, &label_local);
	*length = length_local;
	*label = label_local;
}
void QOpenGLExtraFunctions_glEnablei(void* self, unsigned int target, unsigned int index) {
	((QOpenGLExtraFunctions*)self)->glEnablei(target, index);
}
void QOpenGLExtraFunctions_glDisablei(void* self, unsigned int target, unsigned int index) {
	((QOpenGLExtraFunctions*)self)->glDisablei(target, index);
}
void QOpenGLExtraFunctions_glBlendEquationi(void* self, unsigned int buf, unsigned int mode) {
	((QOpenGLExtraFunctions*)self)->glBlendEquationi(buf, mode);
}
void QOpenGLExtraFunctions_glBlendEquationSeparatei(void* self, unsigned int buf, unsigned int modeRGB, unsigned int modeAlpha) {
	((QOpenGLExtraFunctions*)self)->glBlendEquationSeparatei(buf, modeRGB, modeAlpha);
}
void QOpenGLExtraFunctions_glBlendFunci(void* self, unsigned int buf, unsigned int src, unsigned int dst) {
	((QOpenGLExtraFunctions*)self)->glBlendFunci(buf, src, dst);
}
void QOpenGLExtraFunctions_glBlendFuncSeparatei(void* self, unsigned int buf, unsigned int srcRGB, unsigned int dstRGB, unsigned int srcAlpha, unsigned int dstAlpha) {
	((QOpenGLExtraFunctions*)self)->glBlendFuncSeparatei(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
}
void QOpenGLExtraFunctions_glColorMaski(void* self, unsigned int index, unsigned char r, unsigned char g, unsigned char b, unsigned char a) {
	((QOpenGLExtraFunctions*)self)->glColorMaski(index, r, g, b, a);
}
unsigned char QOpenGLExtraFunctions_glIsEnabledi(void* self, unsigned int target, unsigned int index) {
	auto&& __r = ((QOpenGLExtraFunctions*)self)->glIsEnabledi(target, index);
	return __r;
}
void QOpenGLExtraFunctions_glFramebufferTexture(void* self, unsigned int target, unsigned int attachment, unsigned int texture, int level) {
	((QOpenGLExtraFunctions*)self)->glFramebufferTexture(target, attachment, texture, level);
}
void QOpenGLExtraFunctions_glPrimitiveBoundingBox(void* self, float minX, float minY, float minZ, float minW, float maxX, float maxY, float maxZ, float maxW) {
	((QOpenGLExtraFunctions*)self)->glPrimitiveBoundingBox(minX, minY, minZ, minW, maxX, maxY, maxZ, maxW);
}
unsigned int QOpenGLExtraFunctions_glGetGraphicsResetStatus(void* self) {
	auto&& __r = ((QOpenGLExtraFunctions*)self)->glGetGraphicsResetStatus();
	return __r;
}
void QOpenGLExtraFunctions_glGetnUniformfv(void* self, unsigned int program, int location, int bufSize, float* params) {
	GLfloat params_local{};
	((QOpenGLExtraFunctions*)self)->glGetnUniformfv(program, location, bufSize, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetnUniformiv(void* self, unsigned int program, int location, int bufSize, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetnUniformiv(program, location, bufSize, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetnUniformuiv(void* self, unsigned int program, int location, int bufSize, unsigned int* params) {
	GLuint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetnUniformuiv(program, location, bufSize, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glMinSampleShading(void* self, float value) {
	((QOpenGLExtraFunctions*)self)->glMinSampleShading(value);
}
void QOpenGLExtraFunctions_glPatchParameteri(void* self, unsigned int pname, int value) {
	((QOpenGLExtraFunctions*)self)->glPatchParameteri(pname, value);
}
void QOpenGLExtraFunctions_glGetTexParameterIiv(void* self, unsigned int target, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetTexParameterIiv(target, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetTexParameterIuiv(void* self, unsigned int target, unsigned int pname, unsigned int* params) {
	GLuint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetTexParameterIuiv(target, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetSamplerParameterIiv(void* self, unsigned int sampler, unsigned int pname, int* params) {
	GLint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetSamplerParameterIiv(sampler, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glGetSamplerParameterIuiv(void* self, unsigned int sampler, unsigned int pname, unsigned int* params) {
	GLuint params_local{};
	((QOpenGLExtraFunctions*)self)->glGetSamplerParameterIuiv(sampler, pname, &params_local);
	*params = params_local;
}
void QOpenGLExtraFunctions_glTexBuffer(void* self, unsigned int target, unsigned int internalformat, unsigned int buffer) {
	((QOpenGLExtraFunctions*)self)->glTexBuffer(target, internalformat, buffer);
}
void QOpenGLExtraFunctions_glTexBufferRange(void* self, unsigned int target, unsigned int internalformat, unsigned int buffer, long long offset, long long size) {
	((QOpenGLExtraFunctions*)self)->glTexBufferRange(target, internalformat, buffer, offset, size);
}
void QOpenGLExtraFunctions_glTexStorage3DMultisample(void* self, unsigned int target, int samples, unsigned int internalformat, int width, int height, int depth, unsigned char fixedsamplelocations) {
	((QOpenGLExtraFunctions*)self)->glTexStorage3DMultisample(target, samples, internalformat, width, height, depth, fixedsamplelocations);
}
void QOpenGLExtraFunctions_delete(void* self) { delete (QOpenGLExtraFunctions*)self; }

}

#include <QtGui/qrasterwindow.h>
#include <QtGui/qpaintdevicewindow.h>
#include <QtGui/qwindow.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qpaintdevice.h>
#include <QtCore/qpoint.h>
#include <QtGui/qevent.h>
#include <QtCore/qrect.h>
#include <QtGui/qregion.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qsurface.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qsurfaceformat.h>
#include <QtCore/qsize.h>
#include <QtCore/qmargins.h>
#include <QtGui/qicon.h>
#include <QtGui/qscreen.h>
#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtGui/qcursor.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include <QtGui/qpainter.h>
#include "gen_qrasterwindow.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QRasterWindow__Go : public QRasterWindow {
public:
	long long __goId;
	unsigned char __goOn[37] = {};
	template<typename... A> QRasterWindow__Go(long long id, A&&... a) : QRasterWindow(std::forward<A>(a)...), __goId(id) {}
	using QRasterWindow::sender;
	using QRasterWindow::senderSignalIndex;
	using QRasterWindow::receivers;
	using QRasterWindow::isSignalConnected;
	using QRasterWindow::getDecodedMetricF;
	int cppgoBase_metric(int __a0) { return QRasterWindow::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QRasterWindow::redirected(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QRasterWindow::resizeEvent(__a0); }
	void cppgoBase_exposeEvent(QExposeEvent* __a0) { QRasterWindow::exposeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QRasterWindow::paintEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QRasterWindow::event(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QRasterWindow::moveEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QRasterWindow::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QRasterWindow::focusOutEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QRasterWindow::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QRasterWindow::hideEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QRasterWindow::closeEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QRasterWindow::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QRasterWindow::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QRasterWindow::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QRasterWindow::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QRasterWindow::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QRasterWindow::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QRasterWindow::wheelEvent(__a0); }
	void cppgoBase_touchEvent(QTouchEvent* __a0) { QRasterWindow::touchEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QRasterWindow::tabletEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QRasterWindow::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QRasterWindow::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QRasterWindow::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QRasterWindow::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QRasterWindow::disconnectNotify(__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QRasterWindow::initPainter(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QRasterWindow::sharedPainter(); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QRasterWindow::metaObject();
	}
	int metric(QPaintDevice::PaintDeviceMetric metric) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(metric), &__h); if (__h) return __r; }
		return QRasterWindow::metric(metric);
	}
	QPaintDevice* redirected(QPoint* arg1) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)arg1, &__h); if (__h) return (QPaintDevice*)__r; }
		return QRasterWindow::redirected(arg1);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QRasterWindow::resizeEvent(event);
	}
	void exposeEvent(QExposeEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3239118852587300383LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::exposeEvent(arg1);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QRasterWindow::paintEvent(event);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QRasterWindow::event(event);
	}
	QSurface::SurfaceType surfaceType() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 7339588379958472532LL, &__h); if (__h) return (QSurface::SurfaceType)(__r); }
		return QRasterWindow::surfaceType();
	}
	QSurfaceFormat format() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7267508590720539762LL, &__h); if (__h) return *(QSurfaceFormat*)__r; }
		return QRasterWindow::format();
	}
	QSize size() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5614464919154503228LL, &__h); if (__h) return *(QSize*)__r; }
		return QRasterWindow::size();
	}
	QAccessibleInterface* accessibleRoot() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -2255663269888600317LL, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return QRasterWindow::accessibleRoot();
	}
	QObject* focusObject() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 8913927178009815944LL, &__h); if (__h) return (QObject*)__r; }
		return QRasterWindow::focusObject();
	}
	void moveEvent(QMoveEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::moveEvent(arg1);
	}
	void focusInEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::focusInEvent(arg1);
	}
	void focusOutEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::focusOutEvent(arg1);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::showEvent(arg1);
	}
	void hideEvent(QHideEvent* arg1) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::hideEvent(arg1);
	}
	void closeEvent(QCloseEvent* arg1) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::closeEvent(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::keyPressEvent(arg1);
	}
	void keyReleaseEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::keyReleaseEvent(arg1);
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::mousePressEvent(arg1);
	}
	void mouseReleaseEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::mouseReleaseEvent(arg1);
	}
	void mouseDoubleClickEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::mouseDoubleClickEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::mouseMoveEvent(arg1);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::wheelEvent(arg1);
	}
	void touchEvent(QTouchEvent* arg1) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2204038732330958816LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::touchEvent(arg1);
	}
	void tabletEvent(QTabletEvent* arg1) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)arg1, &__h); if (__h) return; }
		QRasterWindow::tabletEvent(arg1);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QRasterWindow::nativeEvent(eventType, message, result);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QRasterWindow::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QRasterWindow::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QRasterWindow::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QRasterWindow::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QRasterWindow::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QRasterWindow::disconnectNotify(signal);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QRasterWindow::devType();
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QRasterWindow::initPainter(painter);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QRasterWindow::sharedPainter();
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QRasterWindow;
static std::mutex __connsMu_QRasterWindow;

extern "C" {

void* QRasterWindow_base(void* self) { return (void*)static_cast<QPaintDeviceWindow*>((QRasterWindow*)self); }
void QRasterWindow_chain(void* self, void** out) {
	QRasterWindow* __p = (QRasterWindow*)self;
	QPaintDeviceWindow* __b0 = static_cast<QPaintDeviceWindow*>(__p); out[0] = (void*)__b0;
	QWindow* __b1 = static_cast<QWindow*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QRasterWindow__set_override(void* self, int idx) { ((QRasterWindow__Go*)self)->__goOn[idx] = 1; }
void* QRasterWindow_fromQObject(void* o) { return (void*)qobject_cast<QRasterWindow*>((QObject*)o); }
void* QRasterWindow_new(long long __goid) {
	return new QRasterWindow__Go(__goid);
}
void* QRasterWindow_new2(long long __goid, void* parent) {
	return new QRasterWindow__Go(__goid, (QWindow*)parent);
}
char* QRasterWindow_tr(const char* s, long long* __out_len) {
	auto&& __r = QRasterWindow::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRasterWindow_fromWinId(unsigned long long id) {
	auto&& __r = QRasterWindow::fromWinId(id);
	return (void*)(__r);
}
void* QRasterWindow_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QRasterWindow::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QRasterWindow_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QRasterWindow::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QRasterWindow_sender(void* self) {
	auto&& __r = ((QRasterWindow__Go*)self)->sender();
	return (void*)(__r);
}
int QRasterWindow_senderSignalIndex(void* self) {
	auto&& __r = ((QRasterWindow__Go*)self)->senderSignalIndex();
	return __r;
}
int QRasterWindow_receivers(void* self, const char* signal) {
	auto&& __r = ((QRasterWindow__Go*)self)->receivers(signal);
	return __r;
}
int QRasterWindow_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QRasterWindow__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QRasterWindow_devicePixelRatioFScale() {
	auto&& __r = QRasterWindow::devicePixelRatioFScale();
	return __r;
}
int QRasterWindow_encodeMetricF(int metric, double value) {
	auto&& __r = QRasterWindow::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QRasterWindow_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QRasterWindow__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QRasterWindow_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QRasterWindow::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRasterWindow_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QRasterWindow::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRasterWindow_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QRasterWindow::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QRasterWindow_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QRasterWindow::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QRasterWindow_metaObject_default(void* self) {
	auto&& __r = ((QRasterWindow*)self)->QRasterWindow::metaObject();
	return (void*)(__r);
}
int QRasterWindow_metric_default(void* self, int metric) {
	auto&& __r = ((QRasterWindow__Go*)self)->cppgoBase_metric(metric);
	return __r;
}
void* QRasterWindow_redirected_default(void* self, void* arg1) {
	auto&& __r = ((QRasterWindow__Go*)self)->cppgoBase_redirected((QPoint*)arg1);
	return (void*)(__r);
}
void QRasterWindow_resizeEvent_default(void* self, void* event) {
	((QRasterWindow__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QRasterWindow_exposeEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_exposeEvent((QExposeEvent*)arg1);
}
void QRasterWindow_paintEvent_default(void* self, void* event) {
	((QRasterWindow__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
int QRasterWindow_event_default(void* self, void* event) {
	auto&& __r = ((QRasterWindow__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QRasterWindow_surfaceType_default(void* self) {
	auto&& __r = ((QRasterWindow*)self)->QRasterWindow::surfaceType();
	return (int)(__r);
}
void* QRasterWindow_format_default(void* self) {
	auto&& __r = ((QRasterWindow*)self)->QRasterWindow::format();
	return new QSurfaceFormat(__r);
}
void* QRasterWindow_size_default(void* self) {
	auto&& __r = ((QRasterWindow*)self)->QRasterWindow::size();
	return new QSize(__r);
}
void* QRasterWindow_accessibleRoot_default(void* self) {
	auto&& __r = ((QRasterWindow*)self)->QRasterWindow::accessibleRoot();
	return (void*)(__r);
}
void* QRasterWindow_focusObject_default(void* self) {
	auto&& __r = ((QRasterWindow*)self)->QRasterWindow::focusObject();
	return (void*)(__r);
}
void QRasterWindow_moveEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)arg1);
}
void QRasterWindow_focusInEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)arg1);
}
void QRasterWindow_focusOutEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)arg1);
}
void QRasterWindow_showEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QRasterWindow_hideEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_hideEvent((QHideEvent*)arg1);
}
void QRasterWindow_closeEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)arg1);
}
void QRasterWindow_keyPressEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QRasterWindow_keyReleaseEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)arg1);
}
void QRasterWindow_mousePressEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
void QRasterWindow_mouseReleaseEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)arg1);
}
void QRasterWindow_mouseDoubleClickEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)arg1);
}
void QRasterWindow_mouseMoveEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QRasterWindow_wheelEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QRasterWindow_touchEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_touchEvent((QTouchEvent*)arg1);
}
void QRasterWindow_tabletEvent_default(void* self, void* arg1) {
	((QRasterWindow__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)arg1);
}
int QRasterWindow_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QRasterWindow*)self)->QRasterWindow::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QRasterWindow_timerEvent_default(void* self, void* event) {
	((QRasterWindow__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QRasterWindow_childEvent_default(void* self, void* event) {
	((QRasterWindow__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QRasterWindow_customEvent_default(void* self, void* event) {
	((QRasterWindow__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QRasterWindow_connectNotify_default(void* self, void* signal) {
	((QRasterWindow__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QRasterWindow_disconnectNotify_default(void* self, void* signal) {
	((QRasterWindow__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
int QRasterWindow_devType_default(void* self) {
	auto&& __r = ((QRasterWindow*)self)->QRasterWindow::devType();
	return __r;
}
void QRasterWindow_initPainter_default(void* self, void* painter) {
	((QRasterWindow__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QRasterWindow_sharedPainter_default(void* self) {
	auto&& __r = ((QRasterWindow__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QRasterWindow_connect_screenChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::screenChanged, (QRasterWindow*)self, [__id](QScreen* screen){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)screen);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_modalityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::modalityChanged, (QRasterWindow*)self, [__id](Qt::WindowModality modality){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(modality));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_flagsChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::flagsChanged, (QRasterWindow*)self, [__id](QFlags<Qt::WindowType> flags){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(flags));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_windowStateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::windowStateChanged, (QRasterWindow*)self, [__id](Qt::WindowState windowState){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(windowState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::windowTitleChanged, (QRasterWindow*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_xChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::xChanged, (QRasterWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_yChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::yChanged, (QRasterWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_widthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::widthChanged, (QRasterWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_heightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::heightChanged, (QRasterWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_minimumWidthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::minimumWidthChanged, (QRasterWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_minimumHeightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::minimumHeightChanged, (QRasterWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_maximumWidthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::maximumWidthChanged, (QRasterWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_maximumHeightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::maximumHeightChanged, (QRasterWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_safeAreaMarginsChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::safeAreaMarginsChanged, (QRasterWindow*)self, [__id](QMargins arg){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_visibleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::visibleChanged, (QRasterWindow*)self, [__id](bool arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_visibilityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::visibilityChanged, (QRasterWindow*)self, [__id](QWindow::Visibility visibility){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(visibility));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_activeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::activeChanged, (QRasterWindow*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_contentOrientationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::contentOrientationChanged, (QRasterWindow*)self, [__id](Qt::ScreenOrientation orientation){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(orientation));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_focusObjectChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::focusObjectChanged, (QRasterWindow*)self, [__id](QObject* object){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)object);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_opacityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::opacityChanged, (QRasterWindow*)self, [__id](qreal opacity){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(opacity)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_transientParentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::transientParentChanged, (QRasterWindow*)self, [__id](QWindow* transientParent){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)transientParent);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::destroyed, (QRasterWindow*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::destroyed, (QRasterWindow*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRasterWindow*)self, &QRasterWindow::objectNameChanged, (QRasterWindow*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
	__conns_QRasterWindow[__id] = __c;
}
void QRasterWindow__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QRasterWindow);
		auto __it = __conns_QRasterWindow.find(id);
		if (__it == __conns_QRasterWindow.end()) return;
		__c = __it->second;
		__conns_QRasterWindow.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QRasterWindow__Go_new(long long id) { return new QRasterWindow__Go(id); }
void QRasterWindow_delete(void* self) { delete (QRasterWindow*)self; }

}

#include <QtGui/qvalidator.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qlocale.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qintvalidator.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QIntValidator__Go : public QIntValidator {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QIntValidator__Go(long long id, A&&... a) : QIntValidator(std::forward<A>(a)...), __goId(id) {}
	using QIntValidator::sender;
	using QIntValidator::senderSignalIndex;
	using QIntValidator::receivers;
	using QIntValidator::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QIntValidator::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QIntValidator::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QIntValidator::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QIntValidator::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QIntValidator::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QIntValidator::metaObject();
	}
	QValidator::State validate(QString& arg1, int& arg2) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp_intp) { int __r = ((int(*)(long long, long long, void*, int*, int*))cppgo_virt_int_voidp_intp)(__goId, 6784179835071232555LL, (void*)&arg1, (int*)&arg2, &__h); if (__h) return (QValidator::State)(__r); }
		return QIntValidator::validate(arg1, arg2);
	}
	void fixup(QString& input) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4855206404746966715LL, (void*)&input, &__h); if (__h) return; }
		QIntValidator::fixup(input);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QIntValidator::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QIntValidator::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QIntValidator::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QIntValidator::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QIntValidator::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QIntValidator::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QIntValidator::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QIntValidator;
static std::mutex __connsMu_QIntValidator;

extern "C" {

void* QIntValidator_base(void* self) { return (void*)static_cast<QValidator*>((QIntValidator*)self); }
void QIntValidator_chain(void* self, void** out) {
	QIntValidator* __p = (QIntValidator*)self;
	QValidator* __b0 = static_cast<QValidator*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QIntValidator__set_override(void* self, int idx) { ((QIntValidator__Go*)self)->__goOn[idx] = 1; }
void* QIntValidator_fromQObject(void* o) { return (void*)qobject_cast<QIntValidator*>((QObject*)o); }
void* QIntValidator_new(long long __goid) {
	return new QIntValidator__Go(__goid);
}
void* QIntValidator_new2(long long __goid, int bottom, int top) {
	return new QIntValidator__Go(__goid, bottom, top);
}
void* QIntValidator_new3(long long __goid, void* parent) {
	return new QIntValidator__Go(__goid, (QObject*)parent);
}
void* QIntValidator_new4(long long __goid, int bottom, int top, void* parent) {
	return new QIntValidator__Go(__goid, bottom, top, (QObject*)parent);
}
char* QIntValidator_tr(const char* s, long long* __out_len) {
	auto&& __r = QIntValidator::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QIntValidator_setBottom(void* self, int arg1) {
	((QIntValidator*)self)->setBottom(arg1);
}
void QIntValidator_setTop(void* self, int arg1) {
	((QIntValidator*)self)->setTop(arg1);
}
void QIntValidator_setRange(void* self, int bottom, int top) {
	((QIntValidator*)self)->setRange(bottom, top);
}
int QIntValidator_bottom(void* self) {
	auto&& __r = ((QIntValidator*)self)->bottom();
	return __r;
}
int QIntValidator_top(void* self) {
	auto&& __r = ((QIntValidator*)self)->top();
	return __r;
}
void QIntValidator_bottomChanged(void* self, int bottom) {
	((QIntValidator*)self)->bottomChanged(bottom);
}
void QIntValidator_topChanged(void* self, int top) {
	((QIntValidator*)self)->topChanged(top);
}
void* QIntValidator_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QIntValidator::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QIntValidator_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QIntValidator::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QIntValidator_sender(void* self) {
	auto&& __r = ((QIntValidator__Go*)self)->sender();
	return (void*)(__r);
}
int QIntValidator_senderSignalIndex(void* self) {
	auto&& __r = ((QIntValidator__Go*)self)->senderSignalIndex();
	return __r;
}
int QIntValidator_receivers(void* self, const char* signal) {
	auto&& __r = ((QIntValidator__Go*)self)->receivers(signal);
	return __r;
}
int QIntValidator_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QIntValidator__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QIntValidator_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QIntValidator::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QIntValidator_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QIntValidator::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QIntValidator_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QIntValidator::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QIntValidator_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QIntValidator::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QIntValidator_metaObject_default(void* self) {
	auto&& __r = ((QIntValidator*)self)->QIntValidator::metaObject();
	return (void*)(__r);
}
int QIntValidator_validate_default(void* self, void* arg1, int* arg2) {
	int arg2_ref = *arg2;
	auto&& __r = ((QIntValidator*)self)->QIntValidator::validate(*(QString*)arg1, arg2_ref);
	*arg2 = arg2_ref;
	return (int)(__r);
}
void QIntValidator_fixup_default(void* self, void* input) {
	((QIntValidator*)self)->QIntValidator::fixup(*(QString*)input);
}
int QIntValidator_event_default(void* self, void* event) {
	auto&& __r = ((QIntValidator*)self)->QIntValidator::event((QEvent*)event);
	return __r;
}
int QIntValidator_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QIntValidator*)self)->QIntValidator::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QIntValidator_timerEvent_default(void* self, void* event) {
	((QIntValidator__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QIntValidator_childEvent_default(void* self, void* event) {
	((QIntValidator__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QIntValidator_customEvent_default(void* self, void* event) {
	((QIntValidator__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QIntValidator_connectNotify_default(void* self, void* signal) {
	((QIntValidator__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QIntValidator_disconnectNotify_default(void* self, void* signal) {
	((QIntValidator__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QIntValidator_connect_bottomChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIntValidator*)self, &QIntValidator::bottomChanged, (QIntValidator*)self, [__id](int bottom){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, bottom);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIntValidator);
	__conns_QIntValidator[__id] = __c;
}
void QIntValidator_connect_topChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIntValidator*)self, &QIntValidator::topChanged, (QIntValidator*)self, [__id](int top){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, top);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIntValidator);
	__conns_QIntValidator[__id] = __c;
}
void QIntValidator_connect_changed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIntValidator*)self, &QIntValidator::changed, (QIntValidator*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIntValidator);
	__conns_QIntValidator[__id] = __c;
}
void QIntValidator_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIntValidator*)self, &QIntValidator::destroyed, (QIntValidator*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIntValidator);
	__conns_QIntValidator[__id] = __c;
}
void QIntValidator_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIntValidator*)self, &QIntValidator::destroyed, (QIntValidator*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIntValidator);
	__conns_QIntValidator[__id] = __c;
}
void QIntValidator_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIntValidator*)self, &QIntValidator::objectNameChanged, (QIntValidator*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIntValidator);
	__conns_QIntValidator[__id] = __c;
}
void QIntValidator__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QIntValidator);
		auto __it = __conns_QIntValidator.find(id);
		if (__it == __conns_QIntValidator.end()) return;
		__c = __it->second;
		__conns_QIntValidator.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QIntValidator__Go_new(long long id, void* parent) { return new QIntValidator__Go(id, (QObject*)parent); }
void QIntValidator_delete(void* self) { delete (QIntValidator*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qtabbar.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qsize.h>
#include "gen_qstyleoptiontab.h"

extern "C" {

void* QStyleOptionTab_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionTab*)self); }
void* QStyleOptionTab_new() {
	return new QStyleOptionTab();
}
void* QStyleOptionTab_new2(void* other) {
	return new QStyleOptionTab(*(QStyleOptionTab*)other);
}
void* QStyleOptionTab_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionTab*)self)->operator=(*(QStyleOptionTab*)arg1);
	return (void*)&(__r);
}
int QStyleOptionTab_shape_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->shape;
	return (int)(__r);
}
void QStyleOptionTab_shape_set(void* self, int shape) {
	((QStyleOptionTab*)self)->shape = (QTabBar::Shape)(shape);
}
char* QStyleOptionTab_text_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionTab*)self)->text;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionTab_text_set(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyleOptionTab*)self)->text = text_qstr;
}
void* QStyleOptionTab_icon_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->icon;
	return new QIcon(__r);
}
void QStyleOptionTab_icon_set(void* self, void* icon) {
	((QStyleOptionTab*)self)->icon = *(QIcon*)icon;
}
int QStyleOptionTab_row_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->row;
	return __r;
}
void QStyleOptionTab_row_set(void* self, int row) {
	((QStyleOptionTab*)self)->row = row;
}
int QStyleOptionTab_position_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->position;
	return (int)(__r);
}
void QStyleOptionTab_position_set(void* self, int position) {
	((QStyleOptionTab*)self)->position = (QStyleOptionTab::TabPosition)(position);
}
int QStyleOptionTab_selectedPosition_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->selectedPosition;
	return (int)(__r);
}
void QStyleOptionTab_selectedPosition_set(void* self, int selectedPosition) {
	((QStyleOptionTab*)self)->selectedPosition = (QStyleOptionTab::SelectedPosition)(selectedPosition);
}
int QStyleOptionTab_cornerWidgets_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->cornerWidgets;
	return (int)(__r);
}
void QStyleOptionTab_cornerWidgets_set(void* self, int cornerWidgets) {
	((QStyleOptionTab*)self)->cornerWidgets = QFlags<QStyleOptionTab::CornerWidget>(cornerWidgets);
}
void* QStyleOptionTab_iconSize_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->iconSize;
	return new QSize(__r);
}
void QStyleOptionTab_iconSize_set(void* self, void* iconSize) {
	((QStyleOptionTab*)self)->iconSize = *(QSize*)iconSize;
}
int QStyleOptionTab_documentMode_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->documentMode;
	return __r;
}
void QStyleOptionTab_documentMode_set(void* self, int documentMode) {
	((QStyleOptionTab*)self)->documentMode = documentMode;
}
void* QStyleOptionTab_leftButtonSize_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->leftButtonSize;
	return new QSize(__r);
}
void QStyleOptionTab_leftButtonSize_set(void* self, void* leftButtonSize) {
	((QStyleOptionTab*)self)->leftButtonSize = *(QSize*)leftButtonSize;
}
void* QStyleOptionTab_rightButtonSize_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->rightButtonSize;
	return new QSize(__r);
}
void QStyleOptionTab_rightButtonSize_set(void* self, void* rightButtonSize) {
	((QStyleOptionTab*)self)->rightButtonSize = *(QSize*)rightButtonSize;
}
int QStyleOptionTab_features_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->features;
	return (int)(__r);
}
void QStyleOptionTab_features_set(void* self, int features) {
	((QStyleOptionTab*)self)->features = QFlags<QStyleOptionTab::TabFeature>(features);
}
int QStyleOptionTab_tabIndex_get(void* self) {
	auto __r = ((QStyleOptionTab*)self)->tabIndex;
	return __r;
}
void QStyleOptionTab_tabIndex_set(void* self, int tabIndex) {
	((QStyleOptionTab*)self)->tabIndex = tabIndex;
}
void QStyleOptionTab_delete(void* self) { delete (QStyleOptionTab*)self; }

}

#include <QtWidgets/qaccessiblewidget.h>
#include <QtGui/qaccessibleobject.h>
#include <QtGui/qaccessible.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qaccessible_base.h>
#include <QtCore/qstring.h>
#include <QtGui/qwindow.h>
#include <QtCore/qrect.h>
#include <QtGui/qcolor.h>
#include <QtCore/qobject.h>
#include "gen_qaccessiblewidget.h"

extern "C" {

void* QAccessibleWidget_base(void* self) { return (void*)static_cast<QAccessibleObject*>((QAccessibleWidget*)self); }
void QAccessibleWidget_chain(void* self, void** out) {
	QAccessibleWidget* __p = (QAccessibleWidget*)self;
	QAccessibleObject* __b0 = static_cast<QAccessibleObject*>(__p); out[0] = (void*)__b0;
	QAccessibleInterface* __b1 = static_cast<QAccessibleInterface*>(__b0); out[1] = (void*)__b1;
}
void* QAccessibleWidget_new(void* o) {
	return new QAccessibleWidget((QWidget*)o);
}
void* QAccessibleWidget_new2(void* o, int r, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QAccessibleWidget((QWidget*)o, (QAccessible::Role)(r), name_qstr);
}
void* QAccessibleWidget_new3(void* o, int r) {
	return new QAccessibleWidget((QWidget*)o, (QAccessible::Role)(r));
}
char** QAccessibleWidget_actionNames(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleWidget*)self)->actionNames();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QAccessibleWidget_doAction(void* self, const char* actionName, long long actionName_len) {
	QString actionName_qstr = QString::fromUtf8(actionName, actionName_len);
	((QAccessibleWidget*)self)->doAction(actionName_qstr);
}
char** QAccessibleWidget_keyBindingsForAction(void* self, const char* actionName, long long actionName_len, long long* __out_len) {
	QString actionName_qstr = QString::fromUtf8(actionName, actionName_len);
	auto&& __r = ((QAccessibleWidget*)self)->keyBindingsForAction(actionName_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QAccessibleWidget_tr(const char* sourceText, long long* __out_len) {
	auto&& __r = QAccessibleWidget::tr(sourceText);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_localizedActionName(void* self, const char* name, long long name_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QAccessibleWidget*)self)->localizedActionName(name_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_localizedActionDescription(void* self, const char* name, long long name_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QAccessibleWidget*)self)->localizedActionDescription(name_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_pressAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::pressAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_increaseAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::increaseAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_decreaseAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::decreaseAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_showMenuAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::showMenuAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_setFocusAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::setFocusAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_toggleAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::toggleAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_scrollLeftAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::scrollLeftAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_scrollRightAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::scrollRightAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_scrollUpAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::scrollUpAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_scrollDownAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::scrollDownAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_nextPageAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::nextPageAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_previousPageAction(long long* __out_len) {
	auto&& __r = QAccessibleWidget::previousPageAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_tr2(const char* sourceText, const char* disambiguation, long long* __out_len) {
	auto&& __r = QAccessibleWidget::tr(sourceText, disambiguation);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidget_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len) {
	auto&& __r = QAccessibleWidget::tr(sourceText, disambiguation, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessibleWidget_asQAccessibleActionInterface(void* self) {
	auto&& __r = static_cast<QAccessibleActionInterface*>((QAccessibleWidget*)self);
	return (void*)(__r);
}

}

#include <QtWidgets/qdatawidgetmapper.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qdatawidgetmapper.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QDataWidgetMapper__Go : public QDataWidgetMapper {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QDataWidgetMapper__Go(long long id, A&&... a) : QDataWidgetMapper(std::forward<A>(a)...), __goId(id) {}
	using QDataWidgetMapper::sender;
	using QDataWidgetMapper::senderSignalIndex;
	using QDataWidgetMapper::receivers;
	using QDataWidgetMapper::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QDataWidgetMapper::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QDataWidgetMapper::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QDataWidgetMapper::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QDataWidgetMapper::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QDataWidgetMapper::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QDataWidgetMapper::metaObject();
	}
	void setCurrentIndex(int index) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -6491574503686613496LL, index, &__h); if (__h) return; }
		QDataWidgetMapper::setCurrentIndex(index);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QDataWidgetMapper::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QDataWidgetMapper::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QDataWidgetMapper::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QDataWidgetMapper::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QDataWidgetMapper::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QDataWidgetMapper::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QDataWidgetMapper::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QDataWidgetMapper;
static std::mutex __connsMu_QDataWidgetMapper;

extern "C" {

void* QDataWidgetMapper_base(void* self) { return (void*)static_cast<QObject*>((QDataWidgetMapper*)self); }
void QDataWidgetMapper__set_override(void* self, int idx) { ((QDataWidgetMapper__Go*)self)->__goOn[idx] = 1; }
void* QDataWidgetMapper_fromQObject(void* o) { return (void*)qobject_cast<QDataWidgetMapper*>((QObject*)o); }
void* QDataWidgetMapper_new(long long __goid) {
	return new QDataWidgetMapper__Go(__goid);
}
void* QDataWidgetMapper_new2(long long __goid, void* parent) {
	return new QDataWidgetMapper__Go(__goid, (QObject*)parent);
}
char* QDataWidgetMapper_tr(const char* s, long long* __out_len) {
	auto&& __r = QDataWidgetMapper::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDataWidgetMapper_setModel(void* self, void* model) {
	((QDataWidgetMapper*)self)->setModel((QAbstractItemModel*)model);
}
void* QDataWidgetMapper_model(void* self) {
	auto&& __r = ((QDataWidgetMapper*)self)->model();
	return (void*)(__r);
}
void QDataWidgetMapper_setItemDelegate(void* self, void* delegate) {
	((QDataWidgetMapper*)self)->setItemDelegate((QAbstractItemDelegate*)delegate);
}
void* QDataWidgetMapper_itemDelegate(void* self) {
	auto&& __r = ((QDataWidgetMapper*)self)->itemDelegate();
	return (void*)(__r);
}
void QDataWidgetMapper_setRootIndex(void* self, void* index) {
	((QDataWidgetMapper*)self)->setRootIndex(*(QModelIndex*)index);
}
void* QDataWidgetMapper_rootIndex(void* self) {
	auto&& __r = ((QDataWidgetMapper*)self)->rootIndex();
	return new QModelIndex(__r);
}
void QDataWidgetMapper_setOrientation(void* self, int aOrientation) {
	((QDataWidgetMapper*)self)->setOrientation((Qt::Orientation)(aOrientation));
}
int QDataWidgetMapper_orientation(void* self) {
	auto&& __r = ((QDataWidgetMapper*)self)->orientation();
	return (int)(__r);
}
void QDataWidgetMapper_setSubmitPolicy(void* self, int policy) {
	((QDataWidgetMapper*)self)->setSubmitPolicy((QDataWidgetMapper::SubmitPolicy)(policy));
}
int QDataWidgetMapper_submitPolicy(void* self) {
	auto&& __r = ((QDataWidgetMapper*)self)->submitPolicy();
	return (int)(__r);
}
void QDataWidgetMapper_addMapping(void* self, void* widget, int section) {
	((QDataWidgetMapper*)self)->addMapping((QWidget*)widget, section);
}
void QDataWidgetMapper_addMapping2(void* self, void* widget, int section, void* propertyName) {
	((QDataWidgetMapper*)self)->addMapping((QWidget*)widget, section, *(QByteArray*)propertyName);
}
void QDataWidgetMapper_removeMapping(void* self, void* widget) {
	((QDataWidgetMapper*)self)->removeMapping((QWidget*)widget);
}
int QDataWidgetMapper_mappedSection(void* self, void* widget) {
	auto&& __r = ((QDataWidgetMapper*)self)->mappedSection((QWidget*)widget);
	return __r;
}
void* QDataWidgetMapper_mappedPropertyName(void* self, void* widget) {
	auto&& __r = ((QDataWidgetMapper*)self)->mappedPropertyName((QWidget*)widget);
	return new QByteArray(__r);
}
void* QDataWidgetMapper_mappedWidgetAt(void* self, int section) {
	auto&& __r = ((QDataWidgetMapper*)self)->mappedWidgetAt(section);
	return (void*)(__r);
}
void QDataWidgetMapper_clearMapping(void* self) {
	((QDataWidgetMapper*)self)->clearMapping();
}
int QDataWidgetMapper_currentIndex(void* self) {
	auto&& __r = ((QDataWidgetMapper*)self)->currentIndex();
	return __r;
}
void QDataWidgetMapper_revert(void* self) {
	((QDataWidgetMapper*)self)->revert();
}
int QDataWidgetMapper_submit(void* self) {
	auto&& __r = ((QDataWidgetMapper*)self)->submit();
	return __r;
}
void QDataWidgetMapper_toFirst(void* self) {
	((QDataWidgetMapper*)self)->toFirst();
}
void QDataWidgetMapper_toLast(void* self) {
	((QDataWidgetMapper*)self)->toLast();
}
void QDataWidgetMapper_toNext(void* self) {
	((QDataWidgetMapper*)self)->toNext();
}
void QDataWidgetMapper_toPrevious(void* self) {
	((QDataWidgetMapper*)self)->toPrevious();
}
void QDataWidgetMapper_setCurrentIndex(void* self, int index) {
	((QDataWidgetMapper*)self)->setCurrentIndex(index);
}
void QDataWidgetMapper_setCurrentModelIndex(void* self, void* index) {
	((QDataWidgetMapper*)self)->setCurrentModelIndex(*(QModelIndex*)index);
}
void QDataWidgetMapper_currentIndexChanged(void* self, int index) {
	((QDataWidgetMapper*)self)->currentIndexChanged(index);
}
void* QDataWidgetMapper_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QDataWidgetMapper::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QDataWidgetMapper_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QDataWidgetMapper::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QDataWidgetMapper_sender(void* self) {
	auto&& __r = ((QDataWidgetMapper__Go*)self)->sender();
	return (void*)(__r);
}
int QDataWidgetMapper_senderSignalIndex(void* self) {
	auto&& __r = ((QDataWidgetMapper__Go*)self)->senderSignalIndex();
	return __r;
}
int QDataWidgetMapper_receivers(void* self, const char* signal) {
	auto&& __r = ((QDataWidgetMapper__Go*)self)->receivers(signal);
	return __r;
}
int QDataWidgetMapper_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QDataWidgetMapper__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QDataWidgetMapper_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QDataWidgetMapper::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDataWidgetMapper_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QDataWidgetMapper::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDataWidgetMapper_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QDataWidgetMapper::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QDataWidgetMapper_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QDataWidgetMapper::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QDataWidgetMapper_setCurrentIndex_default(void* self, int index) {
	((QDataWidgetMapper*)self)->QDataWidgetMapper::setCurrentIndex(index);
}
void* QDataWidgetMapper_metaObject_default(void* self) {
	auto&& __r = ((QDataWidgetMapper*)self)->QDataWidgetMapper::metaObject();
	return (void*)(__r);
}
int QDataWidgetMapper_event_default(void* self, void* event) {
	auto&& __r = ((QDataWidgetMapper*)self)->QDataWidgetMapper::event((QEvent*)event);
	return __r;
}
int QDataWidgetMapper_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QDataWidgetMapper*)self)->QDataWidgetMapper::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QDataWidgetMapper_timerEvent_default(void* self, void* event) {
	((QDataWidgetMapper__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QDataWidgetMapper_childEvent_default(void* self, void* event) {
	((QDataWidgetMapper__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QDataWidgetMapper_customEvent_default(void* self, void* event) {
	((QDataWidgetMapper__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QDataWidgetMapper_connectNotify_default(void* self, void* signal) {
	((QDataWidgetMapper__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QDataWidgetMapper_disconnectNotify_default(void* self, void* signal) {
	((QDataWidgetMapper__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QDataWidgetMapper_connect_currentIndexChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDataWidgetMapper*)self, &QDataWidgetMapper::currentIndexChanged, (QDataWidgetMapper*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDataWidgetMapper);
	__conns_QDataWidgetMapper[__id] = __c;
}
void QDataWidgetMapper_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDataWidgetMapper*)self, &QDataWidgetMapper::destroyed, (QDataWidgetMapper*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDataWidgetMapper);
	__conns_QDataWidgetMapper[__id] = __c;
}
void QDataWidgetMapper_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDataWidgetMapper*)self, &QDataWidgetMapper::destroyed, (QDataWidgetMapper*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDataWidgetMapper);
	__conns_QDataWidgetMapper[__id] = __c;
}
void QDataWidgetMapper_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDataWidgetMapper*)self, &QDataWidgetMapper::objectNameChanged, (QDataWidgetMapper*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDataWidgetMapper);
	__conns_QDataWidgetMapper[__id] = __c;
}
void QDataWidgetMapper__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QDataWidgetMapper);
		auto __it = __conns_QDataWidgetMapper.find(id);
		if (__it == __conns_QDataWidgetMapper.end()) return;
		__c = __it->second;
		__conns_QDataWidgetMapper.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QDataWidgetMapper__Go_new(long long id, void* parent) { return new QDataWidgetMapper__Go(id, (QObject*)parent); }
void QDataWidgetMapper_delete(void* self) { delete (QDataWidgetMapper*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
#include <QtWidgets/qgraphicsscene.h>
#include <QtWidgets/qgraphicswidget.h>
#include <QtCore/qsize.h>
#include <QtCore/qstring.h>
#include <QtGui/qcursor.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtGui/qtransform.h>
#include <QtWidgets/qgraphicstransform.h>
#include <QtGui/qpainterpath.h>
#include <QtGui/qregion.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qpolygon.h>
#include <QtCore/qvariant.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtGui/qevent.h>
#include "gen_qgraphicsitem.h"
#include "gen_purego_cb.h"

class QGraphicsItem__Go : public QGraphicsItem {
public:
	long long __goId;
	unsigned char __goOn[35] = {};
	template<typename... A> QGraphicsItem__Go(long long id, A&&... a) : QGraphicsItem(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsItem::updateMicroFocus;
	using QGraphicsItem::addToIndex;
	using QGraphicsItem::removeFromIndex;
	using QGraphicsItem::prepareGeometryChange;
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsItem::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsItem::sceneEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsItem::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsItem::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsItem::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsItem::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsItem::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsItem::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsItem::focusOutEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsItem::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsItem::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsItem::hoverLeaveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsItem::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsItem::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsItem::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsItem::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsItem::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsItem::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsItem::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsItem::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsItem::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsItem::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsItem::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsItem::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsItem::extension(__a0); }
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QRectF{};
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		return;
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsItem::advance(phase);
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsItem::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsItem::contains(point);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsItem::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsItem::collidesWithPath(path, mode);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsItem::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsItem::opaqueArea();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsItem::type();
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsItem::sceneEventFilter(watched, event);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsItem::sceneEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::focusOutEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::hoverLeaveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItem::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsItem::inputMethodQuery(query);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsItem::itemChange(change, value);
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsItem::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsItem::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsItem::extension(variant);
	}
};

extern "C" {

void QGraphicsItem__set_override(void* self, int idx) { ((QGraphicsItem__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsItem_scene(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->scene();
	return (void*)(__r);
}
void* QGraphicsItem_parentItem(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->parentItem();
	return (void*)(__r);
}
void* QGraphicsItem_topLevelItem(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->topLevelItem();
	return (void*)(__r);
}
void* QGraphicsItem_parentObject(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->parentObject();
	return (void*)(__r);
}
void* QGraphicsItem_parentWidget(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->parentWidget();
	return (void*)(__r);
}
void* QGraphicsItem_topLevelWidget(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->topLevelWidget();
	return (void*)(__r);
}
void* QGraphicsItem_window(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->window();
	return (void*)(__r);
}
void* QGraphicsItem_panel(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->panel();
	return (void*)(__r);
}
void QGraphicsItem_setParentItem(void* self, void* parent) {
	((QGraphicsItem*)self)->setParentItem((QGraphicsItem*)parent);
}
void** QGraphicsItem_childItems(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsItem*)self)->childItems();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QGraphicsItem_isWidget(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->isWidget();
	return __r;
}
int QGraphicsItem_isWindow(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->isWindow();
	return __r;
}
int QGraphicsItem_isPanel(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->isPanel();
	return __r;
}
void* QGraphicsItem_toGraphicsObject(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->toGraphicsObject();
	return (void*)(__r);
}
void* QGraphicsItem_toGraphicsObject2(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->toGraphicsObject();
	return (void*)(__r);
}
void* QGraphicsItem_group(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->group();
	return (void*)(__r);
}
void QGraphicsItem_setGroup(void* self, void* group) {
	((QGraphicsItem*)self)->setGroup((QGraphicsItemGroup*)group);
}
int QGraphicsItem_flags(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->flags();
	return (int)(__r);
}
void QGraphicsItem_setFlag(void* self, int flag) {
	((QGraphicsItem*)self)->setFlag((QGraphicsItem::GraphicsItemFlag)(flag));
}
void QGraphicsItem_setFlags(void* self, int flags) {
	((QGraphicsItem*)self)->setFlags(QFlags<QGraphicsItem::GraphicsItemFlag>(flags));
}
int QGraphicsItem_cacheMode(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->cacheMode();
	return (int)(__r);
}
void QGraphicsItem_setCacheMode(void* self, int mode) {
	((QGraphicsItem*)self)->setCacheMode((QGraphicsItem::CacheMode)(mode));
}
int QGraphicsItem_panelModality(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->panelModality();
	return (int)(__r);
}
void QGraphicsItem_setPanelModality(void* self, int panelModality) {
	((QGraphicsItem*)self)->setPanelModality((QGraphicsItem::PanelModality)(panelModality));
}
int QGraphicsItem_isBlockedByModalPanel(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->isBlockedByModalPanel();
	return __r;
}
char* QGraphicsItem_toolTip(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsItem*)self)->toolTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsItem_setToolTip(void* self, const char* toolTip, long long toolTip_len) {
	QString toolTip_qstr = QString::fromUtf8(toolTip, toolTip_len);
	((QGraphicsItem*)self)->setToolTip(toolTip_qstr);
}
void* QGraphicsItem_cursor(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->cursor();
	return new QCursor(__r);
}
void QGraphicsItem_setCursor(void* self, void* cursor) {
	((QGraphicsItem*)self)->setCursor(*(QCursor*)cursor);
}
int QGraphicsItem_hasCursor(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->hasCursor();
	return __r;
}
void QGraphicsItem_unsetCursor(void* self) {
	((QGraphicsItem*)self)->unsetCursor();
}
int QGraphicsItem_isVisible(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->isVisible();
	return __r;
}
int QGraphicsItem_isVisibleTo(void* self, void* parent) {
	auto&& __r = ((QGraphicsItem*)self)->isVisibleTo((QGraphicsItem*)parent);
	return __r;
}
void QGraphicsItem_setVisible(void* self, int visible) {
	((QGraphicsItem*)self)->setVisible(visible);
}
void QGraphicsItem_hide(void* self) {
	((QGraphicsItem*)self)->hide();
}
void QGraphicsItem_show(void* self) {
	((QGraphicsItem*)self)->show();
}
int QGraphicsItem_isEnabled(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->isEnabled();
	return __r;
}
void QGraphicsItem_setEnabled(void* self, int enabled) {
	((QGraphicsItem*)self)->setEnabled(enabled);
}
int QGraphicsItem_isSelected(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->isSelected();
	return __r;
}
void QGraphicsItem_setSelected(void* self, int selected) {
	((QGraphicsItem*)self)->setSelected(selected);
}
int QGraphicsItem_acceptDrops(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->acceptDrops();
	return __r;
}
void QGraphicsItem_setAcceptDrops(void* self, int on) {
	((QGraphicsItem*)self)->setAcceptDrops(on);
}
double QGraphicsItem_opacity(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->opacity();
	return __r;
}
double QGraphicsItem_effectiveOpacity(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->effectiveOpacity();
	return __r;
}
void QGraphicsItem_setOpacity(void* self, double opacity) {
	((QGraphicsItem*)self)->setOpacity(opacity);
}
void* QGraphicsItem_graphicsEffect(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->graphicsEffect();
	return (void*)(__r);
}
void QGraphicsItem_setGraphicsEffect(void* self, void* effect) {
	((QGraphicsItem*)self)->setGraphicsEffect((QGraphicsEffect*)effect);
}
int QGraphicsItem_acceptedMouseButtons(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->acceptedMouseButtons();
	return (int)(__r);
}
void QGraphicsItem_setAcceptedMouseButtons(void* self, int buttons) {
	((QGraphicsItem*)self)->setAcceptedMouseButtons(QFlags<Qt::MouseButton>(buttons));
}
int QGraphicsItem_acceptHoverEvents(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->acceptHoverEvents();
	return __r;
}
void QGraphicsItem_setAcceptHoverEvents(void* self, int enabled) {
	((QGraphicsItem*)self)->setAcceptHoverEvents(enabled);
}
int QGraphicsItem_acceptTouchEvents(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->acceptTouchEvents();
	return __r;
}
void QGraphicsItem_setAcceptTouchEvents(void* self, int enabled) {
	((QGraphicsItem*)self)->setAcceptTouchEvents(enabled);
}
int QGraphicsItem_filtersChildEvents(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->filtersChildEvents();
	return __r;
}
void QGraphicsItem_setFiltersChildEvents(void* self, int enabled) {
	((QGraphicsItem*)self)->setFiltersChildEvents(enabled);
}
int QGraphicsItem_handlesChildEvents(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->handlesChildEvents();
	return __r;
}
void QGraphicsItem_setHandlesChildEvents(void* self, int enabled) {
	((QGraphicsItem*)self)->setHandlesChildEvents(enabled);
}
int QGraphicsItem_isActive(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->isActive();
	return __r;
}
void QGraphicsItem_setActive(void* self, int active) {
	((QGraphicsItem*)self)->setActive(active);
}
int QGraphicsItem_hasFocus(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->hasFocus();
	return __r;
}
void QGraphicsItem_setFocus(void* self) {
	((QGraphicsItem*)self)->setFocus();
}
void QGraphicsItem_clearFocus(void* self) {
	((QGraphicsItem*)self)->clearFocus();
}
void* QGraphicsItem_focusProxy(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->focusProxy();
	return (void*)(__r);
}
void QGraphicsItem_setFocusProxy(void* self, void* item) {
	((QGraphicsItem*)self)->setFocusProxy((QGraphicsItem*)item);
}
void* QGraphicsItem_focusItem(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->focusItem();
	return (void*)(__r);
}
void* QGraphicsItem_focusScopeItem(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->focusScopeItem();
	return (void*)(__r);
}
void QGraphicsItem_grabMouse(void* self) {
	((QGraphicsItem*)self)->grabMouse();
}
void QGraphicsItem_ungrabMouse(void* self) {
	((QGraphicsItem*)self)->ungrabMouse();
}
void QGraphicsItem_grabKeyboard(void* self) {
	((QGraphicsItem*)self)->grabKeyboard();
}
void QGraphicsItem_ungrabKeyboard(void* self) {
	((QGraphicsItem*)self)->ungrabKeyboard();
}
void* QGraphicsItem_pos(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->pos();
	return new QPointF(__r);
}
double QGraphicsItem_x(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->x();
	return __r;
}
void QGraphicsItem_setX(void* self, double x) {
	((QGraphicsItem*)self)->setX(x);
}
double QGraphicsItem_y(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->y();
	return __r;
}
void QGraphicsItem_setY(void* self, double y) {
	((QGraphicsItem*)self)->setY(y);
}
void* QGraphicsItem_scenePos(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->scenePos();
	return new QPointF(__r);
}
void QGraphicsItem_setPos(void* self, void* pos) {
	((QGraphicsItem*)self)->setPos(*(QPointF*)pos);
}
void QGraphicsItem_setPos2(void* self, double x, double y) {
	((QGraphicsItem*)self)->setPos(x, y);
}
void QGraphicsItem_moveBy(void* self, double dx, double dy) {
	((QGraphicsItem*)self)->moveBy(dx, dy);
}
void QGraphicsItem_ensureVisible(void* self) {
	((QGraphicsItem*)self)->ensureVisible();
}
void QGraphicsItem_ensureVisible2(void* self, double x, double y, double w, double h) {
	((QGraphicsItem*)self)->ensureVisible(x, y, w, h);
}
void* QGraphicsItem_transform(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->transform();
	return new QTransform(__r);
}
void* QGraphicsItem_sceneTransform(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->sceneTransform();
	return new QTransform(__r);
}
void* QGraphicsItem_deviceTransform(void* self, void* viewportTransform) {
	auto&& __r = ((QGraphicsItem*)self)->deviceTransform(*(QTransform*)viewportTransform);
	return new QTransform(__r);
}
void* QGraphicsItem_itemTransform(void* self, void* other) {
	auto&& __r = ((QGraphicsItem*)self)->itemTransform((QGraphicsItem*)other);
	return new QTransform(__r);
}
void QGraphicsItem_setTransform(void* self, void* matrix) {
	((QGraphicsItem*)self)->setTransform(*(QTransform*)matrix);
}
void QGraphicsItem_resetTransform(void* self) {
	((QGraphicsItem*)self)->resetTransform();
}
void QGraphicsItem_setRotation(void* self, double angle) {
	((QGraphicsItem*)self)->setRotation(angle);
}
double QGraphicsItem_rotation(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->rotation();
	return __r;
}
void QGraphicsItem_setScale(void* self, double scale) {
	((QGraphicsItem*)self)->setScale(scale);
}
double QGraphicsItem_scale(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->scale();
	return __r;
}
void** QGraphicsItem_transformations(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsItem*)self)->transformations();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QGraphicsItem_setTransformations(void* self, void** transformations, long long transformations_len) {
	QList<QGraphicsTransform*> transformations_list; transformations_list.reserve(transformations_len);
	for (long long __i = 0; __i < transformations_len; __i++) transformations_list.append((QGraphicsTransform*)transformations[__i]);
	((QGraphicsItem*)self)->setTransformations(transformations_list);
}
void* QGraphicsItem_transformOriginPoint(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->transformOriginPoint();
	return new QPointF(__r);
}
void QGraphicsItem_setTransformOriginPoint(void* self, void* origin) {
	((QGraphicsItem*)self)->setTransformOriginPoint(*(QPointF*)origin);
}
void QGraphicsItem_setTransformOriginPoint2(void* self, double ax, double ay) {
	((QGraphicsItem*)self)->setTransformOriginPoint(ax, ay);
}
void QGraphicsItem_advance(void* self, int phase) {
	((QGraphicsItem*)self)->advance(phase);
}
double QGraphicsItem_zValue(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->zValue();
	return __r;
}
void QGraphicsItem_setZValue(void* self, double z) {
	((QGraphicsItem*)self)->setZValue(z);
}
void QGraphicsItem_stackBefore(void* self, void* sibling) {
	((QGraphicsItem*)self)->stackBefore((QGraphicsItem*)sibling);
}
void* QGraphicsItem_boundingRect(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->boundingRect();
	return new QRectF(__r);
}
void* QGraphicsItem_childrenBoundingRect(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->childrenBoundingRect();
	return new QRectF(__r);
}
void* QGraphicsItem_sceneBoundingRect(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->sceneBoundingRect();
	return new QRectF(__r);
}
void* QGraphicsItem_shape(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->shape();
	return new QPainterPath(__r);
}
int QGraphicsItem_isClipped(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->isClipped();
	return __r;
}
void* QGraphicsItem_clipPath(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->clipPath();
	return new QPainterPath(__r);
}
int QGraphicsItem_contains(void* self, void* point) {
	auto&& __r = ((QGraphicsItem*)self)->contains(*(QPointF*)point);
	return __r;
}
int QGraphicsItem_collidesWithItem(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsItem*)self)->collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsItem_collidesWithPath(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsItem*)self)->collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
void** QGraphicsItem_collidingItems(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsItem*)self)->collidingItems();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QGraphicsItem_isObscured(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->isObscured();
	return __r;
}
int QGraphicsItem_isObscured2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->isObscured(x, y, w, h);
	return __r;
}
int QGraphicsItem_isObscuredBy(void* self, void* item) {
	auto&& __r = ((QGraphicsItem*)self)->isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsItem_opaqueArea(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->opaqueArea();
	return new QPainterPath(__r);
}
void* QGraphicsItem_boundingRegion(void* self, void* itemToDeviceTransform) {
	auto&& __r = ((QGraphicsItem*)self)->boundingRegion(*(QTransform*)itemToDeviceTransform);
	return new QRegion(__r);
}
double QGraphicsItem_boundingRegionGranularity(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->boundingRegionGranularity();
	return __r;
}
void QGraphicsItem_setBoundingRegionGranularity(void* self, double granularity) {
	((QGraphicsItem*)self)->setBoundingRegionGranularity(granularity);
}
void QGraphicsItem_paint(void* self, void* painter, void* option, void* widget) {
	((QGraphicsItem*)self)->paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
void QGraphicsItem_update(void* self) {
	((QGraphicsItem*)self)->update();
}
void QGraphicsItem_update2(void* self, double x, double y, double width, double height) {
	((QGraphicsItem*)self)->update(x, y, width, height);
}
void QGraphicsItem_scroll(void* self, double dx, double dy) {
	((QGraphicsItem*)self)->scroll(dx, dy);
}
void* QGraphicsItem_mapToItem(void* self, void* item, void* point) {
	auto&& __r = ((QGraphicsItem*)self)->mapToItem((QGraphicsItem*)item, *(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsItem_mapToParent(void* self, void* point) {
	auto&& __r = ((QGraphicsItem*)self)->mapToParent(*(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsItem_mapToScene(void* self, void* point) {
	auto&& __r = ((QGraphicsItem*)self)->mapToScene(*(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsItem_mapToItem2(void* self, void* item, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapToItem((QGraphicsItem*)item, *(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapToParent2(void* self, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapToParent(*(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapToScene2(void* self, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapToScene(*(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapRectToItem(void* self, void* item, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectToItem((QGraphicsItem*)item, *(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsItem_mapRectToParent(void* self, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectToParent(*(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsItem_mapRectToScene(void* self, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectToScene(*(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsItem_mapToItem3(void* self, void* item, void* polygon) {
	auto&& __r = ((QGraphicsItem*)self)->mapToItem((QGraphicsItem*)item, *(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapToParent3(void* self, void* polygon) {
	auto&& __r = ((QGraphicsItem*)self)->mapToParent(*(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapToScene3(void* self, void* polygon) {
	auto&& __r = ((QGraphicsItem*)self)->mapToScene(*(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapToItem4(void* self, void* item, void* path) {
	auto&& __r = ((QGraphicsItem*)self)->mapToItem((QGraphicsItem*)item, *(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsItem_mapToParent4(void* self, void* path) {
	auto&& __r = ((QGraphicsItem*)self)->mapToParent(*(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsItem_mapToScene4(void* self, void* path) {
	auto&& __r = ((QGraphicsItem*)self)->mapToScene(*(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsItem_mapFromItem(void* self, void* item, void* point) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromItem((QGraphicsItem*)item, *(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsItem_mapFromParent(void* self, void* point) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromParent(*(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsItem_mapFromScene(void* self, void* point) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromScene(*(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsItem_mapFromItem2(void* self, void* item, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromItem((QGraphicsItem*)item, *(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapFromParent2(void* self, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromParent(*(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapFromScene2(void* self, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromScene(*(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapRectFromItem(void* self, void* item, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectFromItem((QGraphicsItem*)item, *(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsItem_mapRectFromParent(void* self, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectFromParent(*(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsItem_mapRectFromScene(void* self, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectFromScene(*(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsItem_mapFromItem3(void* self, void* item, void* polygon) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromItem((QGraphicsItem*)item, *(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapFromParent3(void* self, void* polygon) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromParent(*(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapFromScene3(void* self, void* polygon) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromScene(*(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapFromItem4(void* self, void* item, void* path) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromItem((QGraphicsItem*)item, *(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsItem_mapFromParent4(void* self, void* path) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromParent(*(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsItem_mapFromScene4(void* self, void* path) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromScene(*(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsItem_mapToItem5(void* self, void* item, double x, double y) {
	auto&& __r = ((QGraphicsItem*)self)->mapToItem((QGraphicsItem*)item, x, y);
	return new QPointF(__r);
}
void* QGraphicsItem_mapToParent5(void* self, double x, double y) {
	auto&& __r = ((QGraphicsItem*)self)->mapToParent(x, y);
	return new QPointF(__r);
}
void* QGraphicsItem_mapToScene5(void* self, double x, double y) {
	auto&& __r = ((QGraphicsItem*)self)->mapToScene(x, y);
	return new QPointF(__r);
}
void* QGraphicsItem_mapToItem6(void* self, void* item, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapToItem((QGraphicsItem*)item, x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapToParent6(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapToParent(x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapToScene6(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapToScene(x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapRectToItem2(void* self, void* item, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectToItem((QGraphicsItem*)item, x, y, w, h);
	return new QRectF(__r);
}
void* QGraphicsItem_mapRectToParent2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectToParent(x, y, w, h);
	return new QRectF(__r);
}
void* QGraphicsItem_mapRectToScene2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectToScene(x, y, w, h);
	return new QRectF(__r);
}
void* QGraphicsItem_mapFromItem5(void* self, void* item, double x, double y) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromItem((QGraphicsItem*)item, x, y);
	return new QPointF(__r);
}
void* QGraphicsItem_mapFromParent5(void* self, double x, double y) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromParent(x, y);
	return new QPointF(__r);
}
void* QGraphicsItem_mapFromScene5(void* self, double x, double y) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromScene(x, y);
	return new QPointF(__r);
}
void* QGraphicsItem_mapFromItem6(void* self, void* item, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromItem((QGraphicsItem*)item, x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapFromParent6(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromParent(x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapFromScene6(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapFromScene(x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsItem_mapRectFromItem2(void* self, void* item, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectFromItem((QGraphicsItem*)item, x, y, w, h);
	return new QRectF(__r);
}
void* QGraphicsItem_mapRectFromParent2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectFromParent(x, y, w, h);
	return new QRectF(__r);
}
void* QGraphicsItem_mapRectFromScene2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsItem*)self)->mapRectFromScene(x, y, w, h);
	return new QRectF(__r);
}
int QGraphicsItem_isAncestorOf(void* self, void* child) {
	auto&& __r = ((QGraphicsItem*)self)->isAncestorOf((QGraphicsItem*)child);
	return __r;
}
void* QGraphicsItem_commonAncestorItem(void* self, void* other) {
	auto&& __r = ((QGraphicsItem*)self)->commonAncestorItem((QGraphicsItem*)other);
	return (void*)(__r);
}
int QGraphicsItem_isUnderMouse(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->isUnderMouse();
	return __r;
}
void* QGraphicsItem_data(void* self, int key) {
	auto&& __r = ((QGraphicsItem*)self)->data(key);
	return new QVariant(__r);
}
void QGraphicsItem_setData(void* self, int key, void* value) {
	((QGraphicsItem*)self)->setData(key, *(QVariant*)value);
}
int QGraphicsItem_inputMethodHints(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->inputMethodHints();
	return (int)(__r);
}
void QGraphicsItem_setInputMethodHints(void* self, int hints) {
	((QGraphicsItem*)self)->setInputMethodHints(QFlags<Qt::InputMethodHint>(hints));
}
int QGraphicsItem_type(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->type();
	return __r;
}
void QGraphicsItem_installSceneEventFilter(void* self, void* filterItem) {
	((QGraphicsItem*)self)->installSceneEventFilter((QGraphicsItem*)filterItem);
}
void QGraphicsItem_removeSceneEventFilter(void* self, void* filterItem) {
	((QGraphicsItem*)self)->removeSceneEventFilter((QGraphicsItem*)filterItem);
}
void QGraphicsItem_updateMicroFocus(void* self) {
	((QGraphicsItem__Go*)self)->updateMicroFocus();
}
void QGraphicsItem_addToIndex(void* self) {
	((QGraphicsItem__Go*)self)->addToIndex();
}
void QGraphicsItem_removeFromIndex(void* self) {
	((QGraphicsItem__Go*)self)->removeFromIndex();
}
void QGraphicsItem_prepareGeometryChange(void* self) {
	((QGraphicsItem__Go*)self)->prepareGeometryChange();
}
void QGraphicsItem_setFlag2(void* self, int flag, int enabled) {
	((QGraphicsItem*)self)->setFlag((QGraphicsItem::GraphicsItemFlag)(flag), enabled);
}
void QGraphicsItem_setCacheMode2(void* self, int mode, void* cacheSize) {
	((QGraphicsItem*)self)->setCacheMode((QGraphicsItem::CacheMode)(mode), *(QSize*)cacheSize);
}
void QGraphicsItem_setFocus2(void* self, int focusReason) {
	((QGraphicsItem*)self)->setFocus((Qt::FocusReason)(focusReason));
}
void QGraphicsItem_ensureVisible3(void* self, void* rect) {
	((QGraphicsItem*)self)->ensureVisible(*(QRectF*)rect);
}
void QGraphicsItem_ensureVisible4(void* self, void* rect, int xmargin) {
	((QGraphicsItem*)self)->ensureVisible(*(QRectF*)rect, xmargin);
}
void QGraphicsItem_ensureVisible5(void* self, void* rect, int xmargin, int ymargin) {
	((QGraphicsItem*)self)->ensureVisible(*(QRectF*)rect, xmargin, ymargin);
}
void QGraphicsItem_ensureVisible6(void* self, double x, double y, double w, double h, int xmargin) {
	((QGraphicsItem*)self)->ensureVisible(x, y, w, h, xmargin);
}
void QGraphicsItem_ensureVisible7(void* self, double x, double y, double w, double h, int xmargin, int ymargin) {
	((QGraphicsItem*)self)->ensureVisible(x, y, w, h, xmargin, ymargin);
}
void* QGraphicsItem_itemTransform2(void* self, void* other, int* ok) {
	bool ok_local{};
	auto&& __r = ((QGraphicsItem*)self)->itemTransform((QGraphicsItem*)other, &ok_local);
	*ok = ok_local;
	return new QTransform(__r);
}
void QGraphicsItem_setTransform2(void* self, void* matrix, int combine) {
	((QGraphicsItem*)self)->setTransform(*(QTransform*)matrix, combine);
}
void** QGraphicsItem_collidingItems2(void* self, int mode, long long* __out_len) {
	auto&& __r = ((QGraphicsItem*)self)->collidingItems((Qt::ItemSelectionMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QGraphicsItem_isObscured3(void* self, void* rect) {
	auto&& __r = ((QGraphicsItem*)self)->isObscured(*(QRectF*)rect);
	return __r;
}
void QGraphicsItem_update3(void* self, void* rect) {
	((QGraphicsItem*)self)->update(*(QRectF*)rect);
}
void QGraphicsItem_scroll2(void* self, double dx, double dy, void* rect) {
	((QGraphicsItem*)self)->scroll(dx, dy, *(QRectF*)rect);
}
void QGraphicsItem_advance_default(void* self, int phase) {
	((QGraphicsItem*)self)->QGraphicsItem::advance(phase);
}
void* QGraphicsItem_shape_default(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->QGraphicsItem::shape();
	return new QPainterPath(__r);
}
int QGraphicsItem_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsItem*)self)->QGraphicsItem::contains(*(QPointF*)point);
	return __r;
}
int QGraphicsItem_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsItem*)self)->QGraphicsItem::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsItem_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsItem*)self)->QGraphicsItem::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsItem_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsItem*)self)->QGraphicsItem::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsItem_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->QGraphicsItem::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsItem_type_default(void* self) {
	auto&& __r = ((QGraphicsItem*)self)->QGraphicsItem::type();
	return __r;
}
int QGraphicsItem_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsItem__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QGraphicsItem_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsItem__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QGraphicsItem_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsItem_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsItem_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsItem_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsItem_dropEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsItem_focusInEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsItem_focusOutEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsItem_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsItem_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsItem_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsItem_keyPressEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsItem_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsItem_mousePressEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsItem_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsItem_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsItem_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsItem_wheelEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsItem_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsItem__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsItem_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsItem__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void* QGraphicsItem_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsItem__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
int QGraphicsItem_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsItem__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsItem_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsItem__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsItem_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsItem__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void* QGraphicsItem__Go_new(long long id) { return new QGraphicsItem__Go(id); }
void QGraphicsItem_delete(void* self) { delete (QGraphicsItem*)self; }

}

#include <QtWidgets/qgraphicswidget.h>
#include <QtWidgets/qgraphicsitem.h>
#include <QtWidgets/qgraphicslayoutitem.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qgraphicslayout.h>
#include <QtWidgets/qstyle.h>
#include <QtGui/qfont.h>
#include <QtGui/qpalette.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qmargins.h>
#include <QtGui/qkeysequence.h>
#include <QtGui/qaction.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qpainterpath.h>
#include <QtCore/qvariant.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include <QtWidgets/qgraphicsscene.h>
#include <QtGui/qcursor.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qtransform.h>
#include <QtWidgets/qgraphicstransform.h>
#include <QtGui/qregion.h>
#include <QtGui/qpolygon.h>
#include <QtWidgets/qsizepolicy.h>
#include "gen_qgraphicswidget.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsWidget__Go : public QGraphicsWidget {
public:
	long long __goId;
	unsigned char __goOn[65] = {};
	template<typename... A> QGraphicsWidget__Go(long long id, A&&... a) : QGraphicsWidget(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsWidget::updateMicroFocus;
	using QGraphicsWidget::sender;
	using QGraphicsWidget::senderSignalIndex;
	using QGraphicsWidget::receivers;
	using QGraphicsWidget::isSignalConnected;
	using QGraphicsWidget::addToIndex;
	using QGraphicsWidget::removeFromIndex;
	using QGraphicsWidget::prepareGeometryChange;
	using QGraphicsWidget::setGraphicsItem;
	using QGraphicsWidget::setOwnedByLayout;
	void cppgoBase_initStyleOption(QStyleOption* __a0) { QGraphicsWidget::initStyleOption(__a0); }
	QSizeF cppgoBase_sizeHint(int __a0, const QSizeF& __a1) { return QGraphicsWidget::sizeHint((Qt::SizeHint)__a0, __a1); }
	void cppgoBase_updateGeometry() { QGraphicsWidget::updateGeometry(); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsWidget::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	QVariant cppgoBase_propertyChange(const QString& __a0, const QVariant& __a1) { return QGraphicsWidget::propertyChange(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsWidget::sceneEvent(__a0); }
	bool cppgoBase_windowFrameEvent(QEvent* __a0) { return QGraphicsWidget::windowFrameEvent(__a0); }
	Qt::WindowFrameSection cppgoBase_windowFrameSectionAt(const QPointF& __a0) { return QGraphicsWidget::windowFrameSectionAt(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QGraphicsWidget::event(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QGraphicsWidget::changeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QGraphicsWidget::closeEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsWidget::focusInEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QGraphicsWidget::focusNextPrevChild(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsWidget::focusOutEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QGraphicsWidget::hideEvent(__a0); }
	void cppgoBase_moveEvent(QGraphicsSceneMoveEvent* __a0) { QGraphicsWidget::moveEvent(__a0); }
	void cppgoBase_polishEvent() { QGraphicsWidget::polishEvent(); }
	void cppgoBase_resizeEvent(QGraphicsSceneResizeEvent* __a0) { QGraphicsWidget::resizeEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QGraphicsWidget::showEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsWidget::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsWidget::hoverLeaveEvent(__a0); }
	void cppgoBase_grabMouseEvent(QEvent* __a0) { QGraphicsWidget::grabMouseEvent(__a0); }
	void cppgoBase_ungrabMouseEvent(QEvent* __a0) { QGraphicsWidget::ungrabMouseEvent(__a0); }
	void cppgoBase_grabKeyboardEvent(QEvent* __a0) { QGraphicsWidget::grabKeyboardEvent(__a0); }
	void cppgoBase_ungrabKeyboardEvent(QEvent* __a0) { QGraphicsWidget::ungrabKeyboardEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsWidget::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsWidget::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsWidget::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsWidget::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsWidget::disconnectNotify(__a0); }
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsWidget::sceneEventFilter(__a0, __a1); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsWidget::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsWidget::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsWidget::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsWidget::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsWidget::dropEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsWidget::hoverEnterEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsWidget::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsWidget::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsWidget::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsWidget::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsWidget::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsWidget::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsWidget::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsWidget::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsWidget::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsWidget::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsWidget::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsWidget::extension(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsWidget::metaObject();
	}
	void setGeometry(const QRectF& rect) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&rect, &__h); if (__h) return; }
		QGraphicsWidget::setGeometry(rect);
	}
	void getContentsMargins(qreal* left, qreal* top, qreal* right, qreal* bottom) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_doublep_doublep_doublep_doublep) { ((void(*)(long long, long long, double*, double*, double*, double*, int*))cppgo_virt_void_doublep_doublep_doublep_doublep)(__goId, 164760286202646972LL, (double*)left, (double*)top, (double*)right, (double*)bottom, &__h); if (__h) return; }
		QGraphicsWidget::getContentsMargins(left, top, right, bottom);
	}
	int type() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsWidget::type();
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsWidget::paint(painter, option, widget);
	}
	void paintWindowFrame(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -1640122132753285616LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsWidget::paintWindowFrame(painter, option, widget);
	}
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsWidget::boundingRect();
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsWidget::shape();
	}
	void initStyleOption(QStyleOption* option) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QGraphicsWidget::initStyleOption(option);
	}
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF& constraint) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, 729690261457684101LL, (int)(which), (void*)&constraint, &__h); if (__h) return *(QSizeF*)__r; }
		return QGraphicsWidget::sizeHint(which, constraint);
	}
	void updateGeometry() override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8082455443689927668LL, &__h); if (__h) return; }
		QGraphicsWidget::updateGeometry();
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsWidget::itemChange(change, value);
	}
	QVariant propertyChange(const QString& propertyName, const QVariant& value) override {
		int __h = 0;
		QByteArray propertyName_ba = propertyName.toUtf8();
		if (__goOn[12] && cppgo_virt_voidp_constcharp_longlong_voidp) { void* __r = ((void*(*)(long long, long long, const char*, long long, void*, int*))cppgo_virt_voidp_constcharp_longlong_voidp)(__goId, 5289089528078387296LL, (char*)propertyName_ba.constData(), propertyName_ba.size(), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsWidget::propertyChange(propertyName, value);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsWidget::sceneEvent(event);
	}
	bool windowFrameEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7362375632910814166LL, (void*)e, &__h); if (__h) return __r; }
		return QGraphicsWidget::windowFrameEvent(e);
	}
	Qt::WindowFrameSection windowFrameSectionAt(const QPointF& pos) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8834866722574844732LL, (void*)&pos, &__h); if (__h) return (Qt::WindowFrameSection)(__r); }
		return QGraphicsWidget::windowFrameSectionAt(pos);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsWidget::event(event);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::changeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::closeEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::focusInEvent(event);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QGraphicsWidget::focusNextPrevChild(next);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::focusOutEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::hideEvent(event);
	}
	void moveEvent(QGraphicsSceneMoveEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::moveEvent(event);
	}
	void polishEvent() override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -214704830098031580LL, &__h); if (__h) return; }
		QGraphicsWidget::polishEvent();
	}
	void resizeEvent(QGraphicsSceneResizeEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::resizeEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::showEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::hoverLeaveEvent(event);
	}
	void grabMouseEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7928460550911621078LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::grabMouseEvent(event);
	}
	void ungrabMouseEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5662911227992581871LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::ungrabMouseEvent(event);
	}
	void grabKeyboardEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6423297148758141226LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::grabKeyboardEvent(event);
	}
	void ungrabKeyboardEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 2912496686457383679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::ungrabKeyboardEvent(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsWidget::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsWidget::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsWidget::disconnectNotify(signal);
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsWidget::advance(phase);
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsWidget::contains(point);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsWidget::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsWidget::collidesWithPath(path, mode);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsWidget::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsWidget::opaqueArea();
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsWidget::sceneEventFilter(watched, event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::dropEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::hoverEnterEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsWidget::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsWidget::inputMethodQuery(query);
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsWidget::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsWidget::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[63] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsWidget::extension(variant);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[64] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QGraphicsWidget::isEmpty();
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsWidget;
static std::mutex __connsMu_QGraphicsWidget;

extern "C" {

void* QGraphicsWidget_base(void* self) { return (void*)static_cast<QGraphicsObject*>((QGraphicsWidget*)self); }
void QGraphicsWidget_chain(void* self, void** out) {
	QGraphicsWidget* __p = (QGraphicsWidget*)self;
	QGraphicsObject* __b0 = static_cast<QGraphicsObject*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsWidget__set_override(void* self, int idx) { ((QGraphicsWidget__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsWidget_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsWidget*>((QObject*)o); }
void* QGraphicsWidget_new(long long __goid) {
	return new QGraphicsWidget__Go(__goid);
}
void* QGraphicsWidget_new2(long long __goid, void* parent) {
	return new QGraphicsWidget__Go(__goid, (QGraphicsItem*)parent);
}
void* QGraphicsWidget_new3(long long __goid, void* parent, int wFlags) {
	return new QGraphicsWidget__Go(__goid, (QGraphicsItem*)parent, QFlags<Qt::WindowType>(wFlags));
}
char* QGraphicsWidget_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsWidget::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsWidget_layout(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->layout();
	return (void*)(__r);
}
void QGraphicsWidget_setLayout(void* self, void* layout) {
	((QGraphicsWidget*)self)->setLayout((QGraphicsLayout*)layout);
}
void QGraphicsWidget_adjustSize(void* self) {
	((QGraphicsWidget*)self)->adjustSize();
}
int QGraphicsWidget_layoutDirection(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->layoutDirection();
	return (int)(__r);
}
void QGraphicsWidget_setLayoutDirection(void* self, int direction) {
	((QGraphicsWidget*)self)->setLayoutDirection((Qt::LayoutDirection)(direction));
}
void QGraphicsWidget_unsetLayoutDirection(void* self) {
	((QGraphicsWidget*)self)->unsetLayoutDirection();
}
void* QGraphicsWidget_style(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->style();
	return (void*)(__r);
}
void QGraphicsWidget_setStyle(void* self, void* style) {
	((QGraphicsWidget*)self)->setStyle((QStyle*)style);
}
void* QGraphicsWidget_font(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->font();
	return new QFont(__r);
}
void QGraphicsWidget_setFont(void* self, void* font) {
	((QGraphicsWidget*)self)->setFont(*(QFont*)font);
}
void* QGraphicsWidget_palette(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->palette();
	return new QPalette(__r);
}
void QGraphicsWidget_setPalette(void* self, void* palette) {
	((QGraphicsWidget*)self)->setPalette(*(QPalette*)palette);
}
int QGraphicsWidget_autoFillBackground(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->autoFillBackground();
	return __r;
}
void QGraphicsWidget_setAutoFillBackground(void* self, int enabled) {
	((QGraphicsWidget*)self)->setAutoFillBackground(enabled);
}
void QGraphicsWidget_resize(void* self, void* size) {
	((QGraphicsWidget*)self)->resize(*(QSizeF*)size);
}
void QGraphicsWidget_resize2(void* self, double w, double h) {
	((QGraphicsWidget*)self)->resize(w, h);
}
void* QGraphicsWidget_size(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->size();
	return new QSizeF(__r);
}
void QGraphicsWidget_setGeometry(void* self, void* rect) {
	((QGraphicsWidget*)self)->setGeometry(*(QRectF*)rect);
}
void QGraphicsWidget_setGeometry2(void* self, double x, double y, double w, double h) {
	((QGraphicsWidget*)self)->setGeometry(x, y, w, h);
}
void* QGraphicsWidget_rect(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->rect();
	return new QRectF(__r);
}
void QGraphicsWidget_setContentsMargins(void* self, double left, double top, double right, double bottom) {
	((QGraphicsWidget*)self)->setContentsMargins(left, top, right, bottom);
}
void QGraphicsWidget_setContentsMargins2(void* self, void* margins) {
	((QGraphicsWidget*)self)->setContentsMargins(*(QMarginsF*)margins);
}
void QGraphicsWidget_getContentsMargins(void* self, double* left, double* top, double* right, double* bottom) {
	qreal left_local{};
	qreal top_local{};
	qreal right_local{};
	qreal bottom_local{};
	((QGraphicsWidget*)self)->getContentsMargins(&left_local, &top_local, &right_local, &bottom_local);
	*left = left_local;
	*top = top_local;
	*right = right_local;
	*bottom = bottom_local;
}
void QGraphicsWidget_setWindowFrameMargins(void* self, double left, double top, double right, double bottom) {
	((QGraphicsWidget*)self)->setWindowFrameMargins(left, top, right, bottom);
}
void QGraphicsWidget_setWindowFrameMargins2(void* self, void* margins) {
	((QGraphicsWidget*)self)->setWindowFrameMargins(*(QMarginsF*)margins);
}
void QGraphicsWidget_getWindowFrameMargins(void* self, double* left, double* top, double* right, double* bottom) {
	qreal left_local{};
	qreal top_local{};
	qreal right_local{};
	qreal bottom_local{};
	((QGraphicsWidget*)self)->getWindowFrameMargins(&left_local, &top_local, &right_local, &bottom_local);
	*left = left_local;
	*top = top_local;
	*right = right_local;
	*bottom = bottom_local;
}
void QGraphicsWidget_unsetWindowFrameMargins(void* self) {
	((QGraphicsWidget*)self)->unsetWindowFrameMargins();
}
void* QGraphicsWidget_windowFrameGeometry(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->windowFrameGeometry();
	return new QRectF(__r);
}
void* QGraphicsWidget_windowFrameRect(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->windowFrameRect();
	return new QRectF(__r);
}
int QGraphicsWidget_windowFlags(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->windowFlags();
	return (int)(__r);
}
int QGraphicsWidget_windowType(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->windowType();
	return (int)(__r);
}
void QGraphicsWidget_setWindowFlags(void* self, int wFlags) {
	((QGraphicsWidget*)self)->setWindowFlags(QFlags<Qt::WindowType>(wFlags));
}
int QGraphicsWidget_isActiveWindow(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->isActiveWindow();
	return __r;
}
void QGraphicsWidget_setWindowTitle(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	((QGraphicsWidget*)self)->setWindowTitle(title_qstr);
}
char* QGraphicsWidget_windowTitle(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsWidget*)self)->windowTitle();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QGraphicsWidget_focusPolicy(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->focusPolicy();
	return (int)(__r);
}
void QGraphicsWidget_setFocusPolicy(void* self, int policy) {
	((QGraphicsWidget*)self)->setFocusPolicy((Qt::FocusPolicy)(policy));
}
void QGraphicsWidget_setTabOrder(void* first, void* second) {
	QGraphicsWidget::setTabOrder((QGraphicsWidget*)first, (QGraphicsWidget*)second);
}
void* QGraphicsWidget_focusWidget(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->focusWidget();
	return (void*)(__r);
}
int QGraphicsWidget_grabShortcut(void* self, void* sequence) {
	auto&& __r = ((QGraphicsWidget*)self)->grabShortcut(*(QKeySequence*)sequence);
	return __r;
}
void QGraphicsWidget_releaseShortcut(void* self, int id) {
	((QGraphicsWidget*)self)->releaseShortcut(id);
}
void QGraphicsWidget_setShortcutEnabled(void* self, int id) {
	((QGraphicsWidget*)self)->setShortcutEnabled(id);
}
void QGraphicsWidget_setShortcutAutoRepeat(void* self, int id) {
	((QGraphicsWidget*)self)->setShortcutAutoRepeat(id);
}
void QGraphicsWidget_addAction(void* self, void* action) {
	((QGraphicsWidget*)self)->addAction((QAction*)action);
}
void QGraphicsWidget_addActions(void* self, void** actions, long long actions_len) {
	QList<QAction*> actions_list; actions_list.reserve(actions_len);
	for (long long __i = 0; __i < actions_len; __i++) actions_list.append((QAction*)actions[__i]);
	((QGraphicsWidget*)self)->addActions(actions_list);
}
void QGraphicsWidget_insertActions(void* self, void* before, void** actions, long long actions_len) {
	QList<QAction*> actions_list; actions_list.reserve(actions_len);
	for (long long __i = 0; __i < actions_len; __i++) actions_list.append((QAction*)actions[__i]);
	((QGraphicsWidget*)self)->insertActions((QAction*)before, actions_list);
}
void QGraphicsWidget_insertAction(void* self, void* before, void* action) {
	((QGraphicsWidget*)self)->insertAction((QAction*)before, (QAction*)action);
}
void QGraphicsWidget_removeAction(void* self, void* action) {
	((QGraphicsWidget*)self)->removeAction((QAction*)action);
}
void** QGraphicsWidget_actions(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsWidget*)self)->actions();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QGraphicsWidget_setAttribute(void* self, int attribute) {
	((QGraphicsWidget*)self)->setAttribute((Qt::WidgetAttribute)(attribute));
}
int QGraphicsWidget_testAttribute(void* self, int attribute) {
	auto&& __r = ((QGraphicsWidget*)self)->testAttribute((Qt::WidgetAttribute)(attribute));
	return __r;
}
void QGraphicsWidget_paintWindowFrame(void* self, void* painter, void* option, void* widget) {
	((QGraphicsWidget*)self)->paintWindowFrame((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
void QGraphicsWidget_geometryChanged(void* self) {
	((QGraphicsWidget*)self)->geometryChanged();
}
void QGraphicsWidget_layoutChanged(void* self) {
	((QGraphicsWidget*)self)->layoutChanged();
}
int QGraphicsWidget_close(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->close();
	return __r;
}
void QGraphicsWidget_updateMicroFocus(void* self) {
	((QGraphicsWidget__Go*)self)->updateMicroFocus();
}
void* QGraphicsWidget_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsWidget::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsWidget_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsWidget_sender(void* self) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsWidget_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsWidget_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsWidget_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
void QGraphicsWidget_addToIndex(void* self) {
	((QGraphicsWidget__Go*)self)->addToIndex();
}
void QGraphicsWidget_removeFromIndex(void* self) {
	((QGraphicsWidget__Go*)self)->removeFromIndex();
}
void QGraphicsWidget_prepareGeometryChange(void* self) {
	((QGraphicsWidget__Go*)self)->prepareGeometryChange();
}
void QGraphicsWidget_setSizePolicy(void* self, void* policy) {
	((QGraphicsWidget*)self)->setSizePolicy(*(QSizePolicy*)policy);
}
void QGraphicsWidget_setSizePolicy2(void* self, int hPolicy, int vPolicy) {
	((QGraphicsWidget*)self)->setSizePolicy((QSizePolicy::Policy)(hPolicy), (QSizePolicy::Policy)(vPolicy));
}
void* QGraphicsWidget_sizePolicy(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->sizePolicy();
	return new QSizePolicy(__r);
}
void QGraphicsWidget_setMinimumSize(void* self, void* size) {
	((QGraphicsWidget*)self)->setMinimumSize(*(QSizeF*)size);
}
void QGraphicsWidget_setMinimumSize2(void* self, double w, double h) {
	((QGraphicsWidget*)self)->setMinimumSize(w, h);
}
void* QGraphicsWidget_minimumSize(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->minimumSize();
	return new QSizeF(__r);
}
void QGraphicsWidget_setMinimumWidth(void* self, double width) {
	((QGraphicsWidget*)self)->setMinimumWidth(width);
}
double QGraphicsWidget_minimumWidth(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->minimumWidth();
	return __r;
}
void QGraphicsWidget_setMinimumHeight(void* self, double height) {
	((QGraphicsWidget*)self)->setMinimumHeight(height);
}
double QGraphicsWidget_minimumHeight(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->minimumHeight();
	return __r;
}
void QGraphicsWidget_setPreferredSize(void* self, void* size) {
	((QGraphicsWidget*)self)->setPreferredSize(*(QSizeF*)size);
}
void QGraphicsWidget_setPreferredSize2(void* self, double w, double h) {
	((QGraphicsWidget*)self)->setPreferredSize(w, h);
}
void* QGraphicsWidget_preferredSize(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->preferredSize();
	return new QSizeF(__r);
}
void QGraphicsWidget_setPreferredWidth(void* self, double width) {
	((QGraphicsWidget*)self)->setPreferredWidth(width);
}
double QGraphicsWidget_preferredWidth(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->preferredWidth();
	return __r;
}
void QGraphicsWidget_setPreferredHeight(void* self, double height) {
	((QGraphicsWidget*)self)->setPreferredHeight(height);
}
double QGraphicsWidget_preferredHeight(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->preferredHeight();
	return __r;
}
void QGraphicsWidget_setMaximumSize(void* self, void* size) {
	((QGraphicsWidget*)self)->setMaximumSize(*(QSizeF*)size);
}
void QGraphicsWidget_setMaximumSize2(void* self, double w, double h) {
	((QGraphicsWidget*)self)->setMaximumSize(w, h);
}
void* QGraphicsWidget_maximumSize(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->maximumSize();
	return new QSizeF(__r);
}
void QGraphicsWidget_setMaximumWidth(void* self, double width) {
	((QGraphicsWidget*)self)->setMaximumWidth(width);
}
double QGraphicsWidget_maximumWidth(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->maximumWidth();
	return __r;
}
void QGraphicsWidget_setMaximumHeight(void* self, double height) {
	((QGraphicsWidget*)self)->setMaximumHeight(height);
}
double QGraphicsWidget_maximumHeight(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->maximumHeight();
	return __r;
}
void* QGraphicsWidget_geometry(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->geometry();
	return new QRectF(__r);
}
void* QGraphicsWidget_contentsRect(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->contentsRect();
	return new QRectF(__r);
}
void* QGraphicsWidget_effectiveSizeHint(void* self, int which) {
	auto&& __r = ((QGraphicsWidget*)self)->effectiveSizeHint((Qt::SizeHint)(which));
	return new QSizeF(__r);
}
int QGraphicsWidget_isEmpty(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->isEmpty();
	return __r;
}
void* QGraphicsWidget_parentLayoutItem(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->parentLayoutItem();
	return (void*)(__r);
}
void QGraphicsWidget_setParentLayoutItem(void* self, void* parent) {
	((QGraphicsWidget*)self)->setParentLayoutItem((QGraphicsLayoutItem*)parent);
}
int QGraphicsWidget_isLayout(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->isLayout();
	return __r;
}
void* QGraphicsWidget_graphicsItem(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->graphicsItem();
	return (void*)(__r);
}
int QGraphicsWidget_ownedByLayout(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->ownedByLayout();
	return __r;
}
void QGraphicsWidget_setGraphicsItem(void* self, void* item) {
	((QGraphicsWidget__Go*)self)->setGraphicsItem((QGraphicsItem*)item);
}
void QGraphicsWidget_setOwnedByLayout(void* self, int ownedByLayout) {
	((QGraphicsWidget__Go*)self)->setOwnedByLayout(ownedByLayout);
}
char* QGraphicsWidget_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsWidget::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsWidget_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsWidget::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QGraphicsWidget_grabShortcut2(void* self, void* sequence, int context) {
	auto&& __r = ((QGraphicsWidget*)self)->grabShortcut(*(QKeySequence*)sequence, (Qt::ShortcutContext)(context));
	return __r;
}
void QGraphicsWidget_setShortcutEnabled2(void* self, int id, int enabled) {
	((QGraphicsWidget*)self)->setShortcutEnabled(id, enabled);
}
void QGraphicsWidget_setShortcutAutoRepeat2(void* self, int id, int enabled) {
	((QGraphicsWidget*)self)->setShortcutAutoRepeat(id, enabled);
}
void QGraphicsWidget_setAttribute2(void* self, int attribute, int on) {
	((QGraphicsWidget*)self)->setAttribute((Qt::WidgetAttribute)(attribute), on);
}
void* QGraphicsWidget_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsWidget::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsWidget_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QGraphicsWidget_setSizePolicy3(void* self, int hPolicy, int vPolicy, int controlType) {
	((QGraphicsWidget*)self)->setSizePolicy((QSizePolicy::Policy)(hPolicy), (QSizePolicy::Policy)(vPolicy), (QSizePolicy::ControlType)(controlType));
}
void* QGraphicsWidget_effectiveSizeHint2(void* self, int which, void* constraint) {
	auto&& __r = ((QGraphicsWidget*)self)->effectiveSizeHint((Qt::SizeHint)(which), *(QSizeF*)constraint);
	return new QSizeF(__r);
}
void* QGraphicsWidget_asQGraphicsLayoutItem(void* self) {
	auto&& __r = static_cast<QGraphicsLayoutItem*>((QGraphicsWidget*)self);
	return (void*)(__r);
}
void QGraphicsWidget_setGeometry_default(void* self, void* rect) {
	((QGraphicsWidget*)self)->QGraphicsWidget::setGeometry(*(QRectF*)rect);
}
void QGraphicsWidget_setGeometry_default2(void* self, double x, double y, double w, double h) {
	((QGraphicsWidget*)self)->QGraphicsWidget::setGeometry(x, y, w, h);
}
void QGraphicsWidget_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom) {
	qreal left_local{};
	qreal top_local{};
	qreal right_local{};
	qreal bottom_local{};
	((QGraphicsWidget*)self)->QGraphicsWidget::getContentsMargins(&left_local, &top_local, &right_local, &bottom_local);
	*left = left_local;
	*top = top_local;
	*right = right_local;
	*bottom = bottom_local;
}
void QGraphicsWidget_paintWindowFrame_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsWidget*)self)->QGraphicsWidget::paintWindowFrame((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
int QGraphicsWidget_isEmpty_default(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->QGraphicsWidget::isEmpty();
	return __r;
}
void* QGraphicsWidget_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->QGraphicsWidget::metaObject();
	return (void*)(__r);
}
int QGraphicsWidget_type_default(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->QGraphicsWidget::type();
	return __r;
}
void QGraphicsWidget_paint_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsWidget*)self)->QGraphicsWidget::paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
void* QGraphicsWidget_boundingRect_default(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->QGraphicsWidget::boundingRect();
	return new QRectF(__r);
}
void* QGraphicsWidget_shape_default(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->QGraphicsWidget::shape();
	return new QPainterPath(__r);
}
void QGraphicsWidget_initStyleOption_default(void* self, void* option) {
	((QGraphicsWidget__Go*)self)->cppgoBase_initStyleOption((QStyleOption*)option);
}
void* QGraphicsWidget_sizeHint_default(void* self, int which, void* constraint) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_sizeHint(which, *(QSizeF*)constraint);
	return new QSizeF(__r);
}
void QGraphicsWidget_updateGeometry_default(void* self) {
	((QGraphicsWidget__Go*)self)->cppgoBase_updateGeometry();
}
void* QGraphicsWidget_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
void* QGraphicsWidget_propertyChange_default(void* self, const char* propertyName, long long propertyName_len, void* value) {
	QString propertyName_qstr = QString::fromUtf8(propertyName, propertyName_len);
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_propertyChange(propertyName_qstr, *(QVariant*)value);
	return new QVariant(__r);
}
int QGraphicsWidget_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
int QGraphicsWidget_windowFrameEvent_default(void* self, void* e) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_windowFrameEvent((QEvent*)e);
	return __r;
}
int QGraphicsWidget_windowFrameSectionAt_default(void* self, void* pos) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_windowFrameSectionAt(*(QPointF*)pos);
	return (int)(__r);
}
int QGraphicsWidget_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QGraphicsWidget_changeEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QGraphicsWidget_closeEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QGraphicsWidget_focusInEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
int QGraphicsWidget_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QGraphicsWidget_focusOutEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsWidget_hideEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QGraphicsWidget_moveEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_moveEvent((QGraphicsSceneMoveEvent*)event);
}
void QGraphicsWidget_polishEvent_default(void* self) {
	((QGraphicsWidget__Go*)self)->cppgoBase_polishEvent();
}
void QGraphicsWidget_resizeEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_resizeEvent((QGraphicsSceneResizeEvent*)event);
}
void QGraphicsWidget_showEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QGraphicsWidget_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsWidget_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsWidget_grabMouseEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_grabMouseEvent((QEvent*)event);
}
void QGraphicsWidget_ungrabMouseEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_ungrabMouseEvent((QEvent*)event);
}
void QGraphicsWidget_grabKeyboardEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_grabKeyboardEvent((QEvent*)event);
}
void QGraphicsWidget_ungrabKeyboardEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_ungrabKeyboardEvent((QEvent*)event);
}
int QGraphicsWidget_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsWidget*)self)->QGraphicsWidget::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsWidget_timerEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsWidget_childEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsWidget_customEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsWidget_connectNotify_default(void* self, void* signal) {
	((QGraphicsWidget__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsWidget_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsWidget__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsWidget_advance_default(void* self, int phase) {
	((QGraphicsWidget*)self)->QGraphicsWidget::advance(phase);
}
int QGraphicsWidget_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsWidget*)self)->QGraphicsWidget::contains(*(QPointF*)point);
	return __r;
}
int QGraphicsWidget_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsWidget*)self)->QGraphicsWidget::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsWidget_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsWidget*)self)->QGraphicsWidget::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsWidget_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsWidget*)self)->QGraphicsWidget::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsWidget_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsWidget*)self)->QGraphicsWidget::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsWidget_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsWidget_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsWidget_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsWidget_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsWidget_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsWidget_dropEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsWidget_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsWidget_keyPressEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsWidget_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsWidget_mousePressEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsWidget_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsWidget_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsWidget_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsWidget_wheelEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsWidget_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsWidget__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsWidget_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
int QGraphicsWidget_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsWidget_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsWidget__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsWidget_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsWidget__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void QGraphicsWidget_connect_geometryChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::geometryChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::layoutChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_parentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::parentChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_opacityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::opacityChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_visibleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::visibleChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_enabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::enabledChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_xChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::xChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_yChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::yChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_zChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::zChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_rotationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::rotationChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_scaleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::scaleChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_childrenChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::childrenChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_widthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::widthChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_heightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::heightChanged, (QGraphicsWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::destroyed, (QGraphicsWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::destroyed, (QGraphicsWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsWidget*)self, &QGraphicsWidget::objectNameChanged, (QGraphicsWidget*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
	__conns_QGraphicsWidget[__id] = __c;
}
void QGraphicsWidget__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsWidget);
		auto __it = __conns_QGraphicsWidget.find(id);
		if (__it == __conns_QGraphicsWidget.end()) return;
		__c = __it->second;
		__conns_QGraphicsWidget.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsWidget__Go_new(long long id) { return new QGraphicsWidget__Go(id); }
void QGraphicsWidget_delete(void* self) { delete (QGraphicsWidget*)self; }

}

#include <QtWidgets/qkeysequenceedit.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qkeysequence.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
#include <QtGui/qregion.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qcursor.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpainter.h>
#include <QtGui/qpixmap.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qicon.h>
#include <QtCore/qlocale.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtCore/qobject.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qkeysequenceedit.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QKeySequenceEdit__Go : public QKeySequenceEdit {
public:
	long long __goId;
	unsigned char __goOn[48] = {};
	template<typename... A> QKeySequenceEdit__Go(long long id, A&&... a) : QKeySequenceEdit(std::forward<A>(a)...), __goId(id) {}
	using QKeySequenceEdit::updateMicroFocus;
	using QKeySequenceEdit::create;
	using QKeySequenceEdit::destroy;
	using QKeySequenceEdit::focusNextChild;
	using QKeySequenceEdit::focusPreviousChild;
	using QKeySequenceEdit::sender;
	using QKeySequenceEdit::senderSignalIndex;
	using QKeySequenceEdit::receivers;
	using QKeySequenceEdit::isSignalConnected;
	using QKeySequenceEdit::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QKeySequenceEdit::event(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QKeySequenceEdit::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QKeySequenceEdit::keyReleaseEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QKeySequenceEdit::timerEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QKeySequenceEdit::focusOutEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QKeySequenceEdit::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QKeySequenceEdit::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QKeySequenceEdit::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QKeySequenceEdit::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QKeySequenceEdit::wheelEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QKeySequenceEdit::focusInEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QKeySequenceEdit::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QKeySequenceEdit::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QKeySequenceEdit::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QKeySequenceEdit::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QKeySequenceEdit::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QKeySequenceEdit::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QKeySequenceEdit::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QKeySequenceEdit::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QKeySequenceEdit::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QKeySequenceEdit::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QKeySequenceEdit::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QKeySequenceEdit::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QKeySequenceEdit::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QKeySequenceEdit::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QKeySequenceEdit::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QKeySequenceEdit::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QKeySequenceEdit::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QKeySequenceEdit::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QKeySequenceEdit::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QKeySequenceEdit::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QKeySequenceEdit::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QKeySequenceEdit::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QKeySequenceEdit::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QKeySequenceEdit::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QKeySequenceEdit::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QKeySequenceEdit::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QKeySequenceEdit::metaObject();
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QKeySequenceEdit::event(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QKeySequenceEdit::keyPressEvent(arg1);
	}
	void keyReleaseEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)arg1, &__h); if (__h) return; }
		QKeySequenceEdit::keyReleaseEvent(arg1);
	}
	void timerEvent(QTimerEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)arg1, &__h); if (__h) return; }
		QKeySequenceEdit::timerEvent(arg1);
	}
	void focusOutEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)arg1, &__h); if (__h) return; }
		QKeySequenceEdit::focusOutEvent(arg1);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QKeySequenceEdit::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QKeySequenceEdit::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QKeySequenceEdit::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QKeySequenceEdit::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QKeySequenceEdit::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QKeySequenceEdit::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QKeySequenceEdit::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::wheelEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::focusInEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QKeySequenceEdit::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QKeySequenceEdit::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QKeySequenceEdit::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QKeySequenceEdit::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QKeySequenceEdit::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QKeySequenceEdit::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QKeySequenceEdit::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QKeySequenceEdit::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QKeySequenceEdit::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QKeySequenceEdit::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QKeySequenceEdit::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QKeySequenceEdit::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QKeySequenceEdit::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QKeySequenceEdit;
static std::mutex __connsMu_QKeySequenceEdit;

extern "C" {

void* QKeySequenceEdit_base(void* self) { return (void*)static_cast<QWidget*>((QKeySequenceEdit*)self); }
void QKeySequenceEdit_chain(void* self, void** out) {
	QKeySequenceEdit* __p = (QKeySequenceEdit*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QKeySequenceEdit__set_override(void* self, int idx) { ((QKeySequenceEdit__Go*)self)->__goOn[idx] = 1; }
void* QKeySequenceEdit_fromQObject(void* o) { return (void*)qobject_cast<QKeySequenceEdit*>((QObject*)o); }
void* QKeySequenceEdit_new(long long __goid, void* parent) {
	return new QKeySequenceEdit__Go(__goid, (QWidget*)parent);
}
void* QKeySequenceEdit_new2(long long __goid) {
	return new QKeySequenceEdit__Go(__goid);
}
void* QKeySequenceEdit_new3(long long __goid, void* keySequence) {
	return new QKeySequenceEdit__Go(__goid, *(QKeySequence*)keySequence);
}
void* QKeySequenceEdit_new4(long long __goid, void* keySequence, void* parent) {
	return new QKeySequenceEdit__Go(__goid, *(QKeySequence*)keySequence, (QWidget*)parent);
}
char* QKeySequenceEdit_tr(const char* s, long long* __out_len) {
	auto&& __r = QKeySequenceEdit::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QKeySequenceEdit_keySequence(void* self) {
	auto&& __r = ((QKeySequenceEdit*)self)->keySequence();
	return new QKeySequence(__r);
}
long long QKeySequenceEdit_maximumSequenceLength(void* self) {
	auto&& __r = ((QKeySequenceEdit*)self)->maximumSequenceLength();
	return __r;
}
void QKeySequenceEdit_setClearButtonEnabled(void* self, int enable) {
	((QKeySequenceEdit*)self)->setClearButtonEnabled(enable);
}
int QKeySequenceEdit_isClearButtonEnabled(void* self) {
	auto&& __r = ((QKeySequenceEdit*)self)->isClearButtonEnabled();
	return __r;
}
void QKeySequenceEdit_setFinishingKeyCombinations(void* self, void** finishingKeyCombinations, long long finishingKeyCombinations_len) {
	QList<QKeyCombination> finishingKeyCombinations_list; finishingKeyCombinations_list.reserve(finishingKeyCombinations_len);
	for (long long __i = 0; __i < finishingKeyCombinations_len; __i++) finishingKeyCombinations_list.append(*(QKeyCombination*)finishingKeyCombinations[__i]);
	((QKeySequenceEdit*)self)->setFinishingKeyCombinations(finishingKeyCombinations_list);
}
void** QKeySequenceEdit_finishingKeyCombinations(void* self, long long* __out_len) {
	auto&& __r = ((QKeySequenceEdit*)self)->finishingKeyCombinations();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QKeyCombination(__e);
	return __arr;
}
void QKeySequenceEdit_setKeySequence(void* self, void* keySequence) {
	((QKeySequenceEdit*)self)->setKeySequence(*(QKeySequence*)keySequence);
}
void QKeySequenceEdit_clear(void* self) {
	((QKeySequenceEdit*)self)->clear();
}
void QKeySequenceEdit_setMaximumSequenceLength(void* self, long long count) {
	((QKeySequenceEdit*)self)->setMaximumSequenceLength(count);
}
void QKeySequenceEdit_editingFinished(void* self) {
	((QKeySequenceEdit*)self)->editingFinished();
}
void QKeySequenceEdit_keySequenceChanged(void* self, void* keySequence) {
	((QKeySequenceEdit*)self)->keySequenceChanged(*(QKeySequence*)keySequence);
}
void QKeySequenceEdit_setTabOrder(void* arg1, void* arg2) {
	QKeySequenceEdit::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QKeySequenceEdit_mouseGrabber() {
	auto&& __r = QKeySequenceEdit::mouseGrabber();
	return (void*)(__r);
}
void* QKeySequenceEdit_keyboardGrabber() {
	auto&& __r = QKeySequenceEdit::keyboardGrabber();
	return (void*)(__r);
}
void* QKeySequenceEdit_find(unsigned long long arg1) {
	auto&& __r = QKeySequenceEdit::find(arg1);
	return (void*)(__r);
}
void* QKeySequenceEdit_createWindowContainer(void* window) {
	auto&& __r = QKeySequenceEdit::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QKeySequenceEdit_updateMicroFocus(void* self) {
	((QKeySequenceEdit__Go*)self)->updateMicroFocus();
}
void QKeySequenceEdit_create(void* self) {
	((QKeySequenceEdit__Go*)self)->create();
}
void QKeySequenceEdit_destroy(void* self) {
	((QKeySequenceEdit__Go*)self)->destroy();
}
int QKeySequenceEdit_focusNextChild(void* self) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->focusNextChild();
	return __r;
}
int QKeySequenceEdit_focusPreviousChild(void* self) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->focusPreviousChild();
	return __r;
}
void* QKeySequenceEdit_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QKeySequenceEdit::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QKeySequenceEdit_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QKeySequenceEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QKeySequenceEdit_sender(void* self) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->sender();
	return (void*)(__r);
}
int QKeySequenceEdit_senderSignalIndex(void* self) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->senderSignalIndex();
	return __r;
}
int QKeySequenceEdit_receivers(void* self, const char* signal) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->receivers(signal);
	return __r;
}
int QKeySequenceEdit_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QKeySequenceEdit_devicePixelRatioFScale() {
	auto&& __r = QKeySequenceEdit::devicePixelRatioFScale();
	return __r;
}
int QKeySequenceEdit_encodeMetricF(int metric, double value) {
	auto&& __r = QKeySequenceEdit::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QKeySequenceEdit_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QKeySequenceEdit_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QKeySequenceEdit::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QKeySequenceEdit_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QKeySequenceEdit::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QKeySequenceEdit_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QKeySequenceEdit::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QKeySequenceEdit_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QKeySequenceEdit::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QKeySequenceEdit_updateMicroFocus2(void* self, int query) {
	((QKeySequenceEdit__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QKeySequenceEdit_create2(void* self, unsigned long long arg1) {
	((QKeySequenceEdit__Go*)self)->create(arg1);
}
void QKeySequenceEdit_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QKeySequenceEdit__Go*)self)->create(arg1, initializeWindow);
}
void QKeySequenceEdit_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QKeySequenceEdit__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QKeySequenceEdit_destroy2(void* self, int destroyWindow) {
	((QKeySequenceEdit__Go*)self)->destroy(destroyWindow);
}
void QKeySequenceEdit_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QKeySequenceEdit__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QKeySequenceEdit_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QKeySequenceEdit::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QKeySequenceEdit_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QKeySequenceEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QKeySequenceEdit_metaObject_default(void* self) {
	auto&& __r = ((QKeySequenceEdit*)self)->QKeySequenceEdit::metaObject();
	return (void*)(__r);
}
int QKeySequenceEdit_event_default(void* self, void* arg1) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
void QKeySequenceEdit_keyPressEvent_default(void* self, void* arg1) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QKeySequenceEdit_keyReleaseEvent_default(void* self, void* arg1) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)arg1);
}
void QKeySequenceEdit_timerEvent_default(void* self, void* arg1) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)arg1);
}
void QKeySequenceEdit_focusOutEvent_default(void* self, void* arg1) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)arg1);
}
int QKeySequenceEdit_devType_default(void* self) {
	auto&& __r = ((QKeySequenceEdit*)self)->QKeySequenceEdit::devType();
	return __r;
}
void QKeySequenceEdit_setVisible_default(void* self, int visible) {
	((QKeySequenceEdit*)self)->QKeySequenceEdit::setVisible(visible);
}
void* QKeySequenceEdit_sizeHint_default(void* self) {
	auto&& __r = ((QKeySequenceEdit*)self)->QKeySequenceEdit::sizeHint();
	return new QSize(__r);
}
void* QKeySequenceEdit_minimumSizeHint_default(void* self) {
	auto&& __r = ((QKeySequenceEdit*)self)->QKeySequenceEdit::minimumSizeHint();
	return new QSize(__r);
}
int QKeySequenceEdit_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QKeySequenceEdit*)self)->QKeySequenceEdit::heightForWidth(arg1);
	return __r;
}
int QKeySequenceEdit_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QKeySequenceEdit*)self)->QKeySequenceEdit::hasHeightForWidth();
	return __r;
}
void* QKeySequenceEdit_paintEngine_default(void* self) {
	auto&& __r = ((QKeySequenceEdit*)self)->QKeySequenceEdit::paintEngine();
	return (void*)(__r);
}
void QKeySequenceEdit_mousePressEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QKeySequenceEdit_mouseReleaseEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QKeySequenceEdit_mouseDoubleClickEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QKeySequenceEdit_mouseMoveEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QKeySequenceEdit_wheelEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QKeySequenceEdit_focusInEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QKeySequenceEdit_enterEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QKeySequenceEdit_leaveEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QKeySequenceEdit_paintEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QKeySequenceEdit_moveEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QKeySequenceEdit_resizeEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QKeySequenceEdit_closeEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QKeySequenceEdit_contextMenuEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QKeySequenceEdit_tabletEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QKeySequenceEdit_actionEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QKeySequenceEdit_dragEnterEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QKeySequenceEdit_dragMoveEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QKeySequenceEdit_dragLeaveEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QKeySequenceEdit_dropEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QKeySequenceEdit_showEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QKeySequenceEdit_hideEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QKeySequenceEdit_changeEvent_default(void* self, void* arg1) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QKeySequenceEdit_metric_default(void* self, int arg1) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QKeySequenceEdit_initPainter_default(void* self, void* painter) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QKeySequenceEdit_redirected_default(void* self, void* offset) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QKeySequenceEdit_sharedPainter_default(void* self) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QKeySequenceEdit_inputMethodEvent_default(void* self, void* arg1) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QKeySequenceEdit_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QKeySequenceEdit*)self)->QKeySequenceEdit::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QKeySequenceEdit_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QKeySequenceEdit__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QKeySequenceEdit_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QKeySequenceEdit*)self)->QKeySequenceEdit::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QKeySequenceEdit_childEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QKeySequenceEdit_customEvent_default(void* self, void* event) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QKeySequenceEdit_connectNotify_default(void* self, void* signal) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QKeySequenceEdit_disconnectNotify_default(void* self, void* signal) {
	((QKeySequenceEdit__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QKeySequenceEdit_connect_editingFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QKeySequenceEdit*)self, &QKeySequenceEdit::editingFinished, (QKeySequenceEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QKeySequenceEdit);
	__conns_QKeySequenceEdit[__id] = __c;
}
void QKeySequenceEdit_connect_keySequenceChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QKeySequenceEdit*)self, &QKeySequenceEdit::keySequenceChanged, (QKeySequenceEdit*)self, [__id](const QKeySequence& keySequence){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&keySequence);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QKeySequenceEdit);
	__conns_QKeySequenceEdit[__id] = __c;
}
void QKeySequenceEdit_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QKeySequenceEdit*)self, &QKeySequenceEdit::windowTitleChanged, (QKeySequenceEdit*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QKeySequenceEdit);
	__conns_QKeySequenceEdit[__id] = __c;
}
void QKeySequenceEdit_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QKeySequenceEdit*)self, &QKeySequenceEdit::windowIconChanged, (QKeySequenceEdit*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QKeySequenceEdit);
	__conns_QKeySequenceEdit[__id] = __c;
}
void QKeySequenceEdit_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QKeySequenceEdit*)self, &QKeySequenceEdit::windowIconTextChanged, (QKeySequenceEdit*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QKeySequenceEdit);
	__conns_QKeySequenceEdit[__id] = __c;
}
void QKeySequenceEdit_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QKeySequenceEdit*)self, &QKeySequenceEdit::customContextMenuRequested, (QKeySequenceEdit*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QKeySequenceEdit);
	__conns_QKeySequenceEdit[__id] = __c;
}
void QKeySequenceEdit_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QKeySequenceEdit*)self, &QKeySequenceEdit::destroyed, (QKeySequenceEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QKeySequenceEdit);
	__conns_QKeySequenceEdit[__id] = __c;
}
void QKeySequenceEdit_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QKeySequenceEdit*)self, &QKeySequenceEdit::destroyed, (QKeySequenceEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QKeySequenceEdit);
	__conns_QKeySequenceEdit[__id] = __c;
}
void QKeySequenceEdit_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QKeySequenceEdit*)self, &QKeySequenceEdit::objectNameChanged, (QKeySequenceEdit*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QKeySequenceEdit);
	__conns_QKeySequenceEdit[__id] = __c;
}
void QKeySequenceEdit__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QKeySequenceEdit);
		auto __it = __conns_QKeySequenceEdit.find(id);
		if (__it == __conns_QKeySequenceEdit.end()) return;
		__c = __it->second;
		__conns_QKeySequenceEdit.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QKeySequenceEdit__Go_new(long long id) { return new QKeySequenceEdit__Go(id); }
void QKeySequenceEdit_delete(void* self) { delete (QKeySequenceEdit*)self; }

}

#include <QtWidgets/qsizegrip.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtGui/qevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtGui/qregion.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qcursor.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpainter.h>
#include <QtGui/qpixmap.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qicon.h>
#include <QtCore/qlocale.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qsizegrip.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSizeGrip;
static std::mutex __connsMu_QSizeGrip;

extern "C" {

void* QSizeGrip_base(void* self) { return (void*)static_cast<QWidget*>((QSizeGrip*)self); }
void QSizeGrip_chain(void* self, void** out) {
	QSizeGrip* __p = (QSizeGrip*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QSizeGrip_fromQObject(void* o) { return (void*)qobject_cast<QSizeGrip*>((QObject*)o); }
void* QSizeGrip_new(void* parent) {
	return new QSizeGrip((QWidget*)parent);
}
char* QSizeGrip_tr(const char* s, long long* __out_len) {
	auto&& __r = QSizeGrip::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSizeGrip_setTabOrder(void* arg1, void* arg2) {
	QSizeGrip::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QSizeGrip_mouseGrabber() {
	auto&& __r = QSizeGrip::mouseGrabber();
	return (void*)(__r);
}
void* QSizeGrip_keyboardGrabber() {
	auto&& __r = QSizeGrip::keyboardGrabber();
	return (void*)(__r);
}
void* QSizeGrip_find(unsigned long long arg1) {
	auto&& __r = QSizeGrip::find(arg1);
	return (void*)(__r);
}
void* QSizeGrip_createWindowContainer(void* window) {
	auto&& __r = QSizeGrip::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void* QSizeGrip_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSizeGrip::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSizeGrip_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSizeGrip::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
double QSizeGrip_devicePixelRatioFScale() {
	auto&& __r = QSizeGrip::devicePixelRatioFScale();
	return __r;
}
int QSizeGrip_encodeMetricF(int metric, double value) {
	auto&& __r = QSizeGrip::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
char* QSizeGrip_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSizeGrip::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSizeGrip_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSizeGrip::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSizeGrip_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QSizeGrip::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QSizeGrip_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QSizeGrip::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void* QSizeGrip_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSizeGrip::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSizeGrip_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSizeGrip::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QSizeGrip_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSizeGrip*)self, &QSizeGrip::windowTitleChanged, (QSizeGrip*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSizeGrip);
	__conns_QSizeGrip[__id] = __c;
}
void QSizeGrip_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSizeGrip*)self, &QSizeGrip::windowIconChanged, (QSizeGrip*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSizeGrip);
	__conns_QSizeGrip[__id] = __c;
}
void QSizeGrip_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSizeGrip*)self, &QSizeGrip::windowIconTextChanged, (QSizeGrip*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSizeGrip);
	__conns_QSizeGrip[__id] = __c;
}
void QSizeGrip_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSizeGrip*)self, &QSizeGrip::customContextMenuRequested, (QSizeGrip*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSizeGrip);
	__conns_QSizeGrip[__id] = __c;
}
void QSizeGrip_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSizeGrip*)self, &QSizeGrip::destroyed, (QSizeGrip*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSizeGrip);
	__conns_QSizeGrip[__id] = __c;
}
void QSizeGrip_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSizeGrip*)self, &QSizeGrip::destroyed, (QSizeGrip*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSizeGrip);
	__conns_QSizeGrip[__id] = __c;
}
void QSizeGrip_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSizeGrip*)self, &QSizeGrip::objectNameChanged, (QSizeGrip*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSizeGrip);
	__conns_QSizeGrip[__id] = __c;
}
void QSizeGrip__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSizeGrip);
		auto __it = __conns_QSizeGrip.find(id);
		if (__it == __conns_QSizeGrip.end()) return;
		__c = __it->second;
		__conns_QSizeGrip.erase(__it);
	}
	QObject::disconnect(__c);
}
void QSizeGrip_delete(void* self) { delete (QSizeGrip*)self; }

}

#include <QtWidgets/qtreewidgetitemiterator.h>
#include <QtWidgets/qtreewidget.h>
#include "gen_qtreewidgetitemiterator.h"

extern "C" {

void* QTreeWidgetItemIterator_new(void* it) {
	return new QTreeWidgetItemIterator(*(QTreeWidgetItemIterator*)it);
}
void* QTreeWidgetItemIterator_new2(void* widget) {
	return new QTreeWidgetItemIterator((QTreeWidget*)widget);
}
void* QTreeWidgetItemIterator_new3(void* item) {
	return new QTreeWidgetItemIterator((QTreeWidgetItem*)item);
}
void* QTreeWidgetItemIterator_new4(void* widget, int flags) {
	return new QTreeWidgetItemIterator((QTreeWidget*)widget, QFlags<QTreeWidgetItemIterator::IteratorFlag>(flags));
}
void* QTreeWidgetItemIterator_new5(void* item, int flags) {
	return new QTreeWidgetItemIterator((QTreeWidgetItem*)item, QFlags<QTreeWidgetItemIterator::IteratorFlag>(flags));
}
void* QTreeWidgetItemIterator_OperatorAssign(void* self, void* it) {
	auto&& __r = ((QTreeWidgetItemIterator*)self)->operator=(*(QTreeWidgetItemIterator*)it);
	return (void*)&(__r);
}
void* QTreeWidgetItemIterator_OperatorPlusPlus(void* self) {
	auto&& __r = ((QTreeWidgetItemIterator*)self)->operator++();
	return (void*)&(__r);
}
void* QTreeWidgetItemIterator_OperatorPlusPlus2(void* self, int arg1) {
	auto&& __r = ((QTreeWidgetItemIterator*)self)->operator++(arg1);
	return new QTreeWidgetItemIterator(__r);
}
void* QTreeWidgetItemIterator_OperatorPlusAssign(void* self, int n) {
	auto&& __r = ((QTreeWidgetItemIterator*)self)->operator+=(n);
	return (void*)&(__r);
}
void* QTreeWidgetItemIterator_OperatorMinusMinus(void* self) {
	auto&& __r = ((QTreeWidgetItemIterator*)self)->operator--();
	return (void*)&(__r);
}
void* QTreeWidgetItemIterator_OperatorMinusMinus2(void* self, int arg1) {
	auto&& __r = ((QTreeWidgetItemIterator*)self)->operator--(arg1);
	return new QTreeWidgetItemIterator(__r);
}
void* QTreeWidgetItemIterator_OperatorMinusAssign(void* self, int n) {
	auto&& __r = ((QTreeWidgetItemIterator*)self)->operator-=(n);
	return (void*)&(__r);
}
void* QTreeWidgetItemIterator_OperatorMultiply(void* self) {
	auto&& __r = ((QTreeWidgetItemIterator*)self)->operator*();
	return (void*)(__r);
}
void QTreeWidgetItemIterator_delete(void* self) { delete (QTreeWidgetItemIterator*)self; }

}

