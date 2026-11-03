package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFileDialog_chain uintptr
	addr_QFileDialog_new uintptr
	addr_QFileDialog_new2 uintptr
	addr_QFileDialog_new3 uintptr
	addr_QFileDialog_new4 uintptr
	addr_QFileDialog_new5 uintptr
	addr_QFileDialog_new6 uintptr
	addr_QFileDialog_tr uintptr
	addr_QFileDialog_setDirectory uintptr
	addr_QFileDialog_setDirectory2 uintptr
	addr_QFileDialog_directory uintptr
	addr_QFileDialog_setDirectoryUrl uintptr
	addr_QFileDialog_directoryUrl uintptr
	addr_QFileDialog_selectFile uintptr
	addr_QFileDialog_selectedFiles uintptr
	addr_QFileDialog_selectUrl uintptr
	addr_QFileDialog_selectedUrls uintptr
	addr_QFileDialog_setNameFilter uintptr
	addr_QFileDialog_setNameFilters uintptr
	addr_QFileDialog_nameFilters uintptr
	addr_QFileDialog_selectNameFilter uintptr
	addr_QFileDialog_selectedMimeTypeFilter uintptr
	addr_QFileDialog_selectedNameFilter uintptr
	addr_QFileDialog_setMimeTypeFilters uintptr
	addr_QFileDialog_mimeTypeFilters uintptr
	addr_QFileDialog_selectMimeTypeFilter uintptr
	addr_QFileDialog_filter uintptr
	addr_QFileDialog_setFilter uintptr
	addr_QFileDialog_setViewMode uintptr
	addr_QFileDialog_viewMode uintptr
	addr_QFileDialog_setFileMode uintptr
	addr_QFileDialog_fileMode uintptr
	addr_QFileDialog_setAcceptMode uintptr
	addr_QFileDialog_acceptMode uintptr
	addr_QFileDialog_setSidebarUrls uintptr
	addr_QFileDialog_sidebarUrls uintptr
	addr_QFileDialog_saveState uintptr
	addr_QFileDialog_restoreState uintptr
	addr_QFileDialog_setDefaultSuffix uintptr
	addr_QFileDialog_defaultSuffix uintptr
	addr_QFileDialog_setHistory uintptr
	addr_QFileDialog_history uintptr
	addr_QFileDialog_setItemDelegate uintptr
	addr_QFileDialog_itemDelegate uintptr
	addr_QFileDialog_setIconProvider uintptr
	addr_QFileDialog_iconProvider uintptr
	addr_QFileDialog_setLabelText uintptr
	addr_QFileDialog_labelText uintptr
	addr_QFileDialog_setSupportedSchemes uintptr
	addr_QFileDialog_supportedSchemes uintptr
	addr_QFileDialog_setProxyModel uintptr
	addr_QFileDialog_proxyModel uintptr
	addr_QFileDialog_setOption uintptr
	addr_QFileDialog_testOption uintptr
	addr_QFileDialog_setOptions uintptr
	addr_QFileDialog_options uintptr
	addr_QFileDialog_open uintptr
	addr_QFileDialog_fileSelected uintptr
	addr_QFileDialog_filesSelected uintptr
	addr_QFileDialog_currentChanged uintptr
	addr_QFileDialog_directoryEntered uintptr
	addr_QFileDialog_urlSelected uintptr
	addr_QFileDialog_urlsSelected uintptr
	addr_QFileDialog_currentUrlChanged uintptr
	addr_QFileDialog_directoryUrlEntered uintptr
	addr_QFileDialog_filterSelected uintptr
	addr_QFileDialog_getOpenFileName uintptr
	addr_QFileDialog_getOpenFileUrl uintptr
	addr_QFileDialog_getSaveFileName uintptr
	addr_QFileDialog_getSaveFileUrl uintptr
	addr_QFileDialog_getExistingDirectory uintptr
	addr_QFileDialog_getExistingDirectoryUrl uintptr
	addr_QFileDialog_getOpenFileNames uintptr
	addr_QFileDialog_getOpenFileUrls uintptr
	addr_QFileDialog_saveFileContent uintptr
	addr_QFileDialog_adjustPosition uintptr
	addr_QFileDialog_setTabOrder uintptr
	addr_QFileDialog_mouseGrabber uintptr
	addr_QFileDialog_keyboardGrabber uintptr
	addr_QFileDialog_find uintptr
	addr_QFileDialog_createWindowContainer uintptr
	addr_QFileDialog_updateMicroFocus uintptr
	addr_QFileDialog_create uintptr
	addr_QFileDialog_destroy uintptr
	addr_QFileDialog_focusNextChild uintptr
	addr_QFileDialog_focusPreviousChild uintptr
	addr_QFileDialog_connect uintptr
	addr_QFileDialog_connect2 uintptr
	addr_QFileDialog_sender uintptr
	addr_QFileDialog_senderSignalIndex uintptr
	addr_QFileDialog_receivers uintptr
	addr_QFileDialog_isSignalConnected uintptr
	pg_QFileDialog_devicePixelRatioFScale func() float64
	pg_QFileDialog_encodeMetricF func(int32, float64) int32
	pg_QFileDialog_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QFileDialog_tr2 uintptr
	addr_QFileDialog_tr3 uintptr
	addr_QFileDialog_setOption2 uintptr
	addr_QFileDialog_getOpenFileName2 uintptr
	addr_QFileDialog_getOpenFileName3 uintptr
	addr_QFileDialog_getOpenFileName4 uintptr
	addr_QFileDialog_getOpenFileName5 uintptr
	addr_QFileDialog_getOpenFileName6 uintptr
	addr_QFileDialog_getOpenFileName7 uintptr
	addr_QFileDialog_getOpenFileUrl2 uintptr
	addr_QFileDialog_getOpenFileUrl3 uintptr
	addr_QFileDialog_getOpenFileUrl4 uintptr
	addr_QFileDialog_getOpenFileUrl5 uintptr
	addr_QFileDialog_getOpenFileUrl6 uintptr
	addr_QFileDialog_getOpenFileUrl7 uintptr
	addr_QFileDialog_getOpenFileUrl8 uintptr
	addr_QFileDialog_getSaveFileName2 uintptr
	addr_QFileDialog_getSaveFileName3 uintptr
	addr_QFileDialog_getSaveFileName4 uintptr
	addr_QFileDialog_getSaveFileName5 uintptr
	addr_QFileDialog_getSaveFileName6 uintptr
	addr_QFileDialog_getSaveFileName7 uintptr
	addr_QFileDialog_getSaveFileUrl2 uintptr
	addr_QFileDialog_getSaveFileUrl3 uintptr
	addr_QFileDialog_getSaveFileUrl4 uintptr
	addr_QFileDialog_getSaveFileUrl5 uintptr
	addr_QFileDialog_getSaveFileUrl6 uintptr
	addr_QFileDialog_getSaveFileUrl7 uintptr
	addr_QFileDialog_getSaveFileUrl8 uintptr
	addr_QFileDialog_getExistingDirectory2 uintptr
	addr_QFileDialog_getExistingDirectory3 uintptr
	addr_QFileDialog_getExistingDirectory4 uintptr
	addr_QFileDialog_getExistingDirectory5 uintptr
	addr_QFileDialog_getExistingDirectoryUrl2 uintptr
	addr_QFileDialog_getExistingDirectoryUrl3 uintptr
	addr_QFileDialog_getExistingDirectoryUrl4 uintptr
	addr_QFileDialog_getExistingDirectoryUrl5 uintptr
	addr_QFileDialog_getExistingDirectoryUrl6 uintptr
	addr_QFileDialog_getOpenFileNames2 uintptr
	addr_QFileDialog_getOpenFileNames3 uintptr
	addr_QFileDialog_getOpenFileNames4 uintptr
	addr_QFileDialog_getOpenFileNames5 uintptr
	addr_QFileDialog_getOpenFileNames6 uintptr
	addr_QFileDialog_getOpenFileNames7 uintptr
	addr_QFileDialog_getOpenFileUrls2 uintptr
	addr_QFileDialog_getOpenFileUrls3 uintptr
	addr_QFileDialog_getOpenFileUrls4 uintptr
	addr_QFileDialog_getOpenFileUrls5 uintptr
	addr_QFileDialog_getOpenFileUrls6 uintptr
	addr_QFileDialog_getOpenFileUrls7 uintptr
	addr_QFileDialog_getOpenFileUrls8 uintptr
	addr_QFileDialog_saveFileContent2 uintptr
	addr_QFileDialog_createWindowContainer2 uintptr
	addr_QFileDialog_createWindowContainer3 uintptr
	addr_QFileDialog_updateMicroFocus2 uintptr
	addr_QFileDialog_create2 uintptr
	addr_QFileDialog_create3 uintptr
	addr_QFileDialog_create4 uintptr
	addr_QFileDialog_destroy2 uintptr
	addr_QFileDialog_destroy3 uintptr
	addr_QFileDialog_connect3 uintptr
	addr_QFileDialog_connect4 uintptr
	addr_QFileDialog_metaObject_default uintptr
	addr_QFileDialog_setVisible_default uintptr
	addr_QFileDialog_done_default uintptr
	addr_QFileDialog_accept_default uintptr
	addr_QFileDialog_changeEvent_default uintptr
	addr_QFileDialog_sizeHint_default uintptr
	addr_QFileDialog_minimumSizeHint_default uintptr
	addr_QFileDialog_exec_default uintptr
	addr_QFileDialog_reject_default uintptr
	addr_QFileDialog_keyPressEvent_default uintptr
	addr_QFileDialog_closeEvent_default uintptr
	addr_QFileDialog_showEvent_default uintptr
	addr_QFileDialog_resizeEvent_default uintptr
	addr_QFileDialog_contextMenuEvent_default uintptr
	addr_QFileDialog_eventFilter_default uintptr
	addr_QFileDialog_devType_default uintptr
	addr_QFileDialog_heightForWidth_default uintptr
	addr_QFileDialog_hasHeightForWidth_default uintptr
	addr_QFileDialog_paintEngine_default uintptr
	addr_QFileDialog_event_default uintptr
	addr_QFileDialog_mousePressEvent_default uintptr
	addr_QFileDialog_mouseReleaseEvent_default uintptr
	addr_QFileDialog_mouseDoubleClickEvent_default uintptr
	addr_QFileDialog_mouseMoveEvent_default uintptr
	addr_QFileDialog_wheelEvent_default uintptr
	addr_QFileDialog_keyReleaseEvent_default uintptr
	addr_QFileDialog_focusInEvent_default uintptr
	addr_QFileDialog_focusOutEvent_default uintptr
	addr_QFileDialog_enterEvent_default uintptr
	addr_QFileDialog_leaveEvent_default uintptr
	addr_QFileDialog_paintEvent_default uintptr
	addr_QFileDialog_moveEvent_default uintptr
	addr_QFileDialog_tabletEvent_default uintptr
	addr_QFileDialog_actionEvent_default uintptr
	addr_QFileDialog_dragEnterEvent_default uintptr
	addr_QFileDialog_dragMoveEvent_default uintptr
	addr_QFileDialog_dragLeaveEvent_default uintptr
	addr_QFileDialog_dropEvent_default uintptr
	addr_QFileDialog_hideEvent_default uintptr
	addr_QFileDialog_metric_default uintptr
	addr_QFileDialog_initPainter_default uintptr
	addr_QFileDialog_redirected_default uintptr
	addr_QFileDialog_sharedPainter_default uintptr
	addr_QFileDialog_inputMethodEvent_default uintptr
	addr_QFileDialog_inputMethodQuery_default uintptr
	addr_QFileDialog_focusNextPrevChild_default uintptr
	addr_QFileDialog_timerEvent_default uintptr
	addr_QFileDialog_childEvent_default uintptr
	addr_QFileDialog_customEvent_default uintptr
	addr_QFileDialog_connectNotify_default uintptr
	addr_QFileDialog_disconnectNotify_default uintptr
	pg_QFileDialog_connect_fileSelected func(uintptr, int64)
	pg_QFileDialog_connect_filesSelected func(uintptr, int64)
	pg_QFileDialog_connect_currentChanged func(uintptr, int64)
	pg_QFileDialog_connect_directoryEntered func(uintptr, int64)
	pg_QFileDialog_connect_urlSelected func(uintptr, int64)
	pg_QFileDialog_connect_urlsSelected func(uintptr, int64)
	pg_QFileDialog_connect_currentUrlChanged func(uintptr, int64)
	pg_QFileDialog_connect_directoryUrlEntered func(uintptr, int64)
	pg_QFileDialog_connect_filterSelected func(uintptr, int64)
	pg_QFileDialog_connect_finished func(uintptr, int64)
	pg_QFileDialog_connect_accepted func(uintptr, int64)
	pg_QFileDialog_connect_rejected func(uintptr, int64)
	pg_QFileDialog_connect_windowTitleChanged func(uintptr, int64)
	pg_QFileDialog_connect_windowIconChanged func(uintptr, int64)
	pg_QFileDialog_connect_windowIconTextChanged func(uintptr, int64)
	pg_QFileDialog_connect_customContextMenuRequested func(uintptr, int64)
	pg_QFileDialog_connect_destroyed func(uintptr, int64)
	pg_QFileDialog_connect_destroyed2 func(uintptr, int64)
	pg_QFileDialog_connect_objectNameChanged func(uintptr, int64)
	pg_QFileDialog__disconnect func(int64)
	pg_QFileDialog__set_override func(uintptr, int32)
	pg_QFileDialog_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFileDialog_chain, "QFileDialog_chain")
	symaddr(&addr_QFileDialog_new, "QFileDialog_new")
	symaddr(&addr_QFileDialog_new2, "QFileDialog_new2")
	symaddr(&addr_QFileDialog_new3, "QFileDialog_new3")
	symaddr(&addr_QFileDialog_new4, "QFileDialog_new4")
	symaddr(&addr_QFileDialog_new5, "QFileDialog_new5")
	symaddr(&addr_QFileDialog_new6, "QFileDialog_new6")
	symaddr(&addr_QFileDialog_tr, "QFileDialog_tr")
	symaddr(&addr_QFileDialog_setDirectory, "QFileDialog_setDirectory")
	symaddr(&addr_QFileDialog_setDirectory2, "QFileDialog_setDirectory2")
	symaddr(&addr_QFileDialog_directory, "QFileDialog_directory")
	symaddr(&addr_QFileDialog_setDirectoryUrl, "QFileDialog_setDirectoryUrl")
	symaddr(&addr_QFileDialog_directoryUrl, "QFileDialog_directoryUrl")
	symaddr(&addr_QFileDialog_selectFile, "QFileDialog_selectFile")
	symaddr(&addr_QFileDialog_selectedFiles, "QFileDialog_selectedFiles")
	symaddr(&addr_QFileDialog_selectUrl, "QFileDialog_selectUrl")
	symaddr(&addr_QFileDialog_selectedUrls, "QFileDialog_selectedUrls")
	symaddr(&addr_QFileDialog_setNameFilter, "QFileDialog_setNameFilter")
	symaddr(&addr_QFileDialog_setNameFilters, "QFileDialog_setNameFilters")
	symaddr(&addr_QFileDialog_nameFilters, "QFileDialog_nameFilters")
	symaddr(&addr_QFileDialog_selectNameFilter, "QFileDialog_selectNameFilter")
	symaddr(&addr_QFileDialog_selectedMimeTypeFilter, "QFileDialog_selectedMimeTypeFilter")
	symaddr(&addr_QFileDialog_selectedNameFilter, "QFileDialog_selectedNameFilter")
	symaddr(&addr_QFileDialog_setMimeTypeFilters, "QFileDialog_setMimeTypeFilters")
	symaddr(&addr_QFileDialog_mimeTypeFilters, "QFileDialog_mimeTypeFilters")
	symaddr(&addr_QFileDialog_selectMimeTypeFilter, "QFileDialog_selectMimeTypeFilter")
	symaddr(&addr_QFileDialog_filter, "QFileDialog_filter")
	symaddr(&addr_QFileDialog_setFilter, "QFileDialog_setFilter")
	symaddr(&addr_QFileDialog_setViewMode, "QFileDialog_setViewMode")
	symaddr(&addr_QFileDialog_viewMode, "QFileDialog_viewMode")
	symaddr(&addr_QFileDialog_setFileMode, "QFileDialog_setFileMode")
	symaddr(&addr_QFileDialog_fileMode, "QFileDialog_fileMode")
	symaddr(&addr_QFileDialog_setAcceptMode, "QFileDialog_setAcceptMode")
	symaddr(&addr_QFileDialog_acceptMode, "QFileDialog_acceptMode")
	symaddr(&addr_QFileDialog_setSidebarUrls, "QFileDialog_setSidebarUrls")
	symaddr(&addr_QFileDialog_sidebarUrls, "QFileDialog_sidebarUrls")
	symaddr(&addr_QFileDialog_saveState, "QFileDialog_saveState")
	symaddr(&addr_QFileDialog_restoreState, "QFileDialog_restoreState")
	symaddr(&addr_QFileDialog_setDefaultSuffix, "QFileDialog_setDefaultSuffix")
	symaddr(&addr_QFileDialog_defaultSuffix, "QFileDialog_defaultSuffix")
	symaddr(&addr_QFileDialog_setHistory, "QFileDialog_setHistory")
	symaddr(&addr_QFileDialog_history, "QFileDialog_history")
	symaddr(&addr_QFileDialog_setItemDelegate, "QFileDialog_setItemDelegate")
	symaddr(&addr_QFileDialog_itemDelegate, "QFileDialog_itemDelegate")
	symaddr(&addr_QFileDialog_setIconProvider, "QFileDialog_setIconProvider")
	symaddr(&addr_QFileDialog_iconProvider, "QFileDialog_iconProvider")
	symaddr(&addr_QFileDialog_setLabelText, "QFileDialog_setLabelText")
	symaddr(&addr_QFileDialog_labelText, "QFileDialog_labelText")
	symaddr(&addr_QFileDialog_setSupportedSchemes, "QFileDialog_setSupportedSchemes")
	symaddr(&addr_QFileDialog_supportedSchemes, "QFileDialog_supportedSchemes")
	symaddr(&addr_QFileDialog_setProxyModel, "QFileDialog_setProxyModel")
	symaddr(&addr_QFileDialog_proxyModel, "QFileDialog_proxyModel")
	symaddr(&addr_QFileDialog_setOption, "QFileDialog_setOption")
	symaddr(&addr_QFileDialog_testOption, "QFileDialog_testOption")
	symaddr(&addr_QFileDialog_setOptions, "QFileDialog_setOptions")
	symaddr(&addr_QFileDialog_options, "QFileDialog_options")
	symaddr(&addr_QFileDialog_open, "QFileDialog_open")
	symaddr(&addr_QFileDialog_fileSelected, "QFileDialog_fileSelected")
	symaddr(&addr_QFileDialog_filesSelected, "QFileDialog_filesSelected")
	symaddr(&addr_QFileDialog_currentChanged, "QFileDialog_currentChanged")
	symaddr(&addr_QFileDialog_directoryEntered, "QFileDialog_directoryEntered")
	symaddr(&addr_QFileDialog_urlSelected, "QFileDialog_urlSelected")
	symaddr(&addr_QFileDialog_urlsSelected, "QFileDialog_urlsSelected")
	symaddr(&addr_QFileDialog_currentUrlChanged, "QFileDialog_currentUrlChanged")
	symaddr(&addr_QFileDialog_directoryUrlEntered, "QFileDialog_directoryUrlEntered")
	symaddr(&addr_QFileDialog_filterSelected, "QFileDialog_filterSelected")
	symaddr(&addr_QFileDialog_getOpenFileName, "QFileDialog_getOpenFileName")
	symaddr(&addr_QFileDialog_getOpenFileUrl, "QFileDialog_getOpenFileUrl")
	symaddr(&addr_QFileDialog_getSaveFileName, "QFileDialog_getSaveFileName")
	symaddr(&addr_QFileDialog_getSaveFileUrl, "QFileDialog_getSaveFileUrl")
	symaddr(&addr_QFileDialog_getExistingDirectory, "QFileDialog_getExistingDirectory")
	symaddr(&addr_QFileDialog_getExistingDirectoryUrl, "QFileDialog_getExistingDirectoryUrl")
	symaddr(&addr_QFileDialog_getOpenFileNames, "QFileDialog_getOpenFileNames")
	symaddr(&addr_QFileDialog_getOpenFileUrls, "QFileDialog_getOpenFileUrls")
	symaddr(&addr_QFileDialog_saveFileContent, "QFileDialog_saveFileContent")
	symaddr(&addr_QFileDialog_adjustPosition, "QFileDialog_adjustPosition")
	symaddr(&addr_QFileDialog_setTabOrder, "QFileDialog_setTabOrder")
	symaddr(&addr_QFileDialog_mouseGrabber, "QFileDialog_mouseGrabber")
	symaddr(&addr_QFileDialog_keyboardGrabber, "QFileDialog_keyboardGrabber")
	symaddr(&addr_QFileDialog_find, "QFileDialog_find")
	symaddr(&addr_QFileDialog_createWindowContainer, "QFileDialog_createWindowContainer")
	symaddr(&addr_QFileDialog_updateMicroFocus, "QFileDialog_updateMicroFocus")
	symaddr(&addr_QFileDialog_create, "QFileDialog_create")
	symaddr(&addr_QFileDialog_destroy, "QFileDialog_destroy")
	symaddr(&addr_QFileDialog_focusNextChild, "QFileDialog_focusNextChild")
	symaddr(&addr_QFileDialog_focusPreviousChild, "QFileDialog_focusPreviousChild")
	symaddr(&addr_QFileDialog_connect, "QFileDialog_connect")
	symaddr(&addr_QFileDialog_connect2, "QFileDialog_connect2")
	symaddr(&addr_QFileDialog_sender, "QFileDialog_sender")
	symaddr(&addr_QFileDialog_senderSignalIndex, "QFileDialog_senderSignalIndex")
	symaddr(&addr_QFileDialog_receivers, "QFileDialog_receivers")
	symaddr(&addr_QFileDialog_isSignalConnected, "QFileDialog_isSignalConnected")
	reg(&pg_QFileDialog_devicePixelRatioFScale, "QFileDialog_devicePixelRatioFScale")
	reg(&pg_QFileDialog_encodeMetricF, "QFileDialog_encodeMetricF")
	reg(&pg_QFileDialog_getDecodedMetricF, "QFileDialog_getDecodedMetricF")
	symaddr(&addr_QFileDialog_tr2, "QFileDialog_tr2")
	symaddr(&addr_QFileDialog_tr3, "QFileDialog_tr3")
	symaddr(&addr_QFileDialog_setOption2, "QFileDialog_setOption2")
	symaddr(&addr_QFileDialog_getOpenFileName2, "QFileDialog_getOpenFileName2")
	symaddr(&addr_QFileDialog_getOpenFileName3, "QFileDialog_getOpenFileName3")
	symaddr(&addr_QFileDialog_getOpenFileName4, "QFileDialog_getOpenFileName4")
	symaddr(&addr_QFileDialog_getOpenFileName5, "QFileDialog_getOpenFileName5")
	symaddr(&addr_QFileDialog_getOpenFileName6, "QFileDialog_getOpenFileName6")
	symaddr(&addr_QFileDialog_getOpenFileName7, "QFileDialog_getOpenFileName7")
	symaddr(&addr_QFileDialog_getOpenFileUrl2, "QFileDialog_getOpenFileUrl2")
	symaddr(&addr_QFileDialog_getOpenFileUrl3, "QFileDialog_getOpenFileUrl3")
	symaddr(&addr_QFileDialog_getOpenFileUrl4, "QFileDialog_getOpenFileUrl4")
	symaddr(&addr_QFileDialog_getOpenFileUrl5, "QFileDialog_getOpenFileUrl5")
	symaddr(&addr_QFileDialog_getOpenFileUrl6, "QFileDialog_getOpenFileUrl6")
	symaddr(&addr_QFileDialog_getOpenFileUrl7, "QFileDialog_getOpenFileUrl7")
	symaddr(&addr_QFileDialog_getOpenFileUrl8, "QFileDialog_getOpenFileUrl8")
	symaddr(&addr_QFileDialog_getSaveFileName2, "QFileDialog_getSaveFileName2")
	symaddr(&addr_QFileDialog_getSaveFileName3, "QFileDialog_getSaveFileName3")
	symaddr(&addr_QFileDialog_getSaveFileName4, "QFileDialog_getSaveFileName4")
	symaddr(&addr_QFileDialog_getSaveFileName5, "QFileDialog_getSaveFileName5")
	symaddr(&addr_QFileDialog_getSaveFileName6, "QFileDialog_getSaveFileName6")
	symaddr(&addr_QFileDialog_getSaveFileName7, "QFileDialog_getSaveFileName7")
	symaddr(&addr_QFileDialog_getSaveFileUrl2, "QFileDialog_getSaveFileUrl2")
	symaddr(&addr_QFileDialog_getSaveFileUrl3, "QFileDialog_getSaveFileUrl3")
	symaddr(&addr_QFileDialog_getSaveFileUrl4, "QFileDialog_getSaveFileUrl4")
	symaddr(&addr_QFileDialog_getSaveFileUrl5, "QFileDialog_getSaveFileUrl5")
	symaddr(&addr_QFileDialog_getSaveFileUrl6, "QFileDialog_getSaveFileUrl6")
	symaddr(&addr_QFileDialog_getSaveFileUrl7, "QFileDialog_getSaveFileUrl7")
	symaddr(&addr_QFileDialog_getSaveFileUrl8, "QFileDialog_getSaveFileUrl8")
	symaddr(&addr_QFileDialog_getExistingDirectory2, "QFileDialog_getExistingDirectory2")
	symaddr(&addr_QFileDialog_getExistingDirectory3, "QFileDialog_getExistingDirectory3")
	symaddr(&addr_QFileDialog_getExistingDirectory4, "QFileDialog_getExistingDirectory4")
	symaddr(&addr_QFileDialog_getExistingDirectory5, "QFileDialog_getExistingDirectory5")
	symaddr(&addr_QFileDialog_getExistingDirectoryUrl2, "QFileDialog_getExistingDirectoryUrl2")
	symaddr(&addr_QFileDialog_getExistingDirectoryUrl3, "QFileDialog_getExistingDirectoryUrl3")
	symaddr(&addr_QFileDialog_getExistingDirectoryUrl4, "QFileDialog_getExistingDirectoryUrl4")
	symaddr(&addr_QFileDialog_getExistingDirectoryUrl5, "QFileDialog_getExistingDirectoryUrl5")
	symaddr(&addr_QFileDialog_getExistingDirectoryUrl6, "QFileDialog_getExistingDirectoryUrl6")
	symaddr(&addr_QFileDialog_getOpenFileNames2, "QFileDialog_getOpenFileNames2")
	symaddr(&addr_QFileDialog_getOpenFileNames3, "QFileDialog_getOpenFileNames3")
	symaddr(&addr_QFileDialog_getOpenFileNames4, "QFileDialog_getOpenFileNames4")
	symaddr(&addr_QFileDialog_getOpenFileNames5, "QFileDialog_getOpenFileNames5")
	symaddr(&addr_QFileDialog_getOpenFileNames6, "QFileDialog_getOpenFileNames6")
	symaddr(&addr_QFileDialog_getOpenFileNames7, "QFileDialog_getOpenFileNames7")
	symaddr(&addr_QFileDialog_getOpenFileUrls2, "QFileDialog_getOpenFileUrls2")
	symaddr(&addr_QFileDialog_getOpenFileUrls3, "QFileDialog_getOpenFileUrls3")
	symaddr(&addr_QFileDialog_getOpenFileUrls4, "QFileDialog_getOpenFileUrls4")
	symaddr(&addr_QFileDialog_getOpenFileUrls5, "QFileDialog_getOpenFileUrls5")
	symaddr(&addr_QFileDialog_getOpenFileUrls6, "QFileDialog_getOpenFileUrls6")
	symaddr(&addr_QFileDialog_getOpenFileUrls7, "QFileDialog_getOpenFileUrls7")
	symaddr(&addr_QFileDialog_getOpenFileUrls8, "QFileDialog_getOpenFileUrls8")
	symaddr(&addr_QFileDialog_saveFileContent2, "QFileDialog_saveFileContent2")
	symaddr(&addr_QFileDialog_createWindowContainer2, "QFileDialog_createWindowContainer2")
	symaddr(&addr_QFileDialog_createWindowContainer3, "QFileDialog_createWindowContainer3")
	symaddr(&addr_QFileDialog_updateMicroFocus2, "QFileDialog_updateMicroFocus2")
	symaddr(&addr_QFileDialog_create2, "QFileDialog_create2")
	symaddr(&addr_QFileDialog_create3, "QFileDialog_create3")
	symaddr(&addr_QFileDialog_create4, "QFileDialog_create4")
	symaddr(&addr_QFileDialog_destroy2, "QFileDialog_destroy2")
	symaddr(&addr_QFileDialog_destroy3, "QFileDialog_destroy3")
	symaddr(&addr_QFileDialog_connect3, "QFileDialog_connect3")
	symaddr(&addr_QFileDialog_connect4, "QFileDialog_connect4")
	symaddr(&addr_QFileDialog_metaObject_default, "QFileDialog_metaObject_default")
	symaddr(&addr_QFileDialog_setVisible_default, "QFileDialog_setVisible_default")
	symaddr(&addr_QFileDialog_done_default, "QFileDialog_done_default")
	symaddr(&addr_QFileDialog_accept_default, "QFileDialog_accept_default")
	symaddr(&addr_QFileDialog_changeEvent_default, "QFileDialog_changeEvent_default")
	symaddr(&addr_QFileDialog_sizeHint_default, "QFileDialog_sizeHint_default")
	symaddr(&addr_QFileDialog_minimumSizeHint_default, "QFileDialog_minimumSizeHint_default")
	symaddr(&addr_QFileDialog_exec_default, "QFileDialog_exec_default")
	symaddr(&addr_QFileDialog_reject_default, "QFileDialog_reject_default")
	symaddr(&addr_QFileDialog_keyPressEvent_default, "QFileDialog_keyPressEvent_default")
	symaddr(&addr_QFileDialog_closeEvent_default, "QFileDialog_closeEvent_default")
	symaddr(&addr_QFileDialog_showEvent_default, "QFileDialog_showEvent_default")
	symaddr(&addr_QFileDialog_resizeEvent_default, "QFileDialog_resizeEvent_default")
	symaddr(&addr_QFileDialog_contextMenuEvent_default, "QFileDialog_contextMenuEvent_default")
	symaddr(&addr_QFileDialog_eventFilter_default, "QFileDialog_eventFilter_default")
	symaddr(&addr_QFileDialog_devType_default, "QFileDialog_devType_default")
	symaddr(&addr_QFileDialog_heightForWidth_default, "QFileDialog_heightForWidth_default")
	symaddr(&addr_QFileDialog_hasHeightForWidth_default, "QFileDialog_hasHeightForWidth_default")
	symaddr(&addr_QFileDialog_paintEngine_default, "QFileDialog_paintEngine_default")
	symaddr(&addr_QFileDialog_event_default, "QFileDialog_event_default")
	symaddr(&addr_QFileDialog_mousePressEvent_default, "QFileDialog_mousePressEvent_default")
	symaddr(&addr_QFileDialog_mouseReleaseEvent_default, "QFileDialog_mouseReleaseEvent_default")
	symaddr(&addr_QFileDialog_mouseDoubleClickEvent_default, "QFileDialog_mouseDoubleClickEvent_default")
	symaddr(&addr_QFileDialog_mouseMoveEvent_default, "QFileDialog_mouseMoveEvent_default")
	symaddr(&addr_QFileDialog_wheelEvent_default, "QFileDialog_wheelEvent_default")
	symaddr(&addr_QFileDialog_keyReleaseEvent_default, "QFileDialog_keyReleaseEvent_default")
	symaddr(&addr_QFileDialog_focusInEvent_default, "QFileDialog_focusInEvent_default")
	symaddr(&addr_QFileDialog_focusOutEvent_default, "QFileDialog_focusOutEvent_default")
	symaddr(&addr_QFileDialog_enterEvent_default, "QFileDialog_enterEvent_default")
	symaddr(&addr_QFileDialog_leaveEvent_default, "QFileDialog_leaveEvent_default")
	symaddr(&addr_QFileDialog_paintEvent_default, "QFileDialog_paintEvent_default")
	symaddr(&addr_QFileDialog_moveEvent_default, "QFileDialog_moveEvent_default")
	symaddr(&addr_QFileDialog_tabletEvent_default, "QFileDialog_tabletEvent_default")
	symaddr(&addr_QFileDialog_actionEvent_default, "QFileDialog_actionEvent_default")
	symaddr(&addr_QFileDialog_dragEnterEvent_default, "QFileDialog_dragEnterEvent_default")
	symaddr(&addr_QFileDialog_dragMoveEvent_default, "QFileDialog_dragMoveEvent_default")
	symaddr(&addr_QFileDialog_dragLeaveEvent_default, "QFileDialog_dragLeaveEvent_default")
	symaddr(&addr_QFileDialog_dropEvent_default, "QFileDialog_dropEvent_default")
	symaddr(&addr_QFileDialog_hideEvent_default, "QFileDialog_hideEvent_default")
	symaddr(&addr_QFileDialog_metric_default, "QFileDialog_metric_default")
	symaddr(&addr_QFileDialog_initPainter_default, "QFileDialog_initPainter_default")
	symaddr(&addr_QFileDialog_redirected_default, "QFileDialog_redirected_default")
	symaddr(&addr_QFileDialog_sharedPainter_default, "QFileDialog_sharedPainter_default")
	symaddr(&addr_QFileDialog_inputMethodEvent_default, "QFileDialog_inputMethodEvent_default")
	symaddr(&addr_QFileDialog_inputMethodQuery_default, "QFileDialog_inputMethodQuery_default")
	symaddr(&addr_QFileDialog_focusNextPrevChild_default, "QFileDialog_focusNextPrevChild_default")
	symaddr(&addr_QFileDialog_timerEvent_default, "QFileDialog_timerEvent_default")
	symaddr(&addr_QFileDialog_childEvent_default, "QFileDialog_childEvent_default")
	symaddr(&addr_QFileDialog_customEvent_default, "QFileDialog_customEvent_default")
	symaddr(&addr_QFileDialog_connectNotify_default, "QFileDialog_connectNotify_default")
	symaddr(&addr_QFileDialog_disconnectNotify_default, "QFileDialog_disconnectNotify_default")
	reg(&pg_QFileDialog_connect_fileSelected, "QFileDialog_connect_fileSelected")
	reg(&pg_QFileDialog_connect_filesSelected, "QFileDialog_connect_filesSelected")
	reg(&pg_QFileDialog_connect_currentChanged, "QFileDialog_connect_currentChanged")
	reg(&pg_QFileDialog_connect_directoryEntered, "QFileDialog_connect_directoryEntered")
	reg(&pg_QFileDialog_connect_urlSelected, "QFileDialog_connect_urlSelected")
	reg(&pg_QFileDialog_connect_urlsSelected, "QFileDialog_connect_urlsSelected")
	reg(&pg_QFileDialog_connect_currentUrlChanged, "QFileDialog_connect_currentUrlChanged")
	reg(&pg_QFileDialog_connect_directoryUrlEntered, "QFileDialog_connect_directoryUrlEntered")
	reg(&pg_QFileDialog_connect_filterSelected, "QFileDialog_connect_filterSelected")
	reg(&pg_QFileDialog_connect_finished, "QFileDialog_connect_finished")
	reg(&pg_QFileDialog_connect_accepted, "QFileDialog_connect_accepted")
	reg(&pg_QFileDialog_connect_rejected, "QFileDialog_connect_rejected")
	reg(&pg_QFileDialog_connect_windowTitleChanged, "QFileDialog_connect_windowTitleChanged")
	reg(&pg_QFileDialog_connect_windowIconChanged, "QFileDialog_connect_windowIconChanged")
	reg(&pg_QFileDialog_connect_windowIconTextChanged, "QFileDialog_connect_windowIconTextChanged")
	reg(&pg_QFileDialog_connect_customContextMenuRequested, "QFileDialog_connect_customContextMenuRequested")
	reg(&pg_QFileDialog_connect_destroyed, "QFileDialog_connect_destroyed")
	reg(&pg_QFileDialog_connect_destroyed2, "QFileDialog_connect_destroyed2")
	reg(&pg_QFileDialog_connect_objectNameChanged, "QFileDialog_connect_objectNameChanged")
	reg(&pg_QFileDialog__disconnect, "QFileDialog__disconnect")
	reg(&pg_QFileDialog__set_override, "QFileDialog__set_override")
	reg(&pg_QFileDialog_delete, "QFileDialog_delete")
}

type QFileDialog struct {
	h uintptr
	*QDialog
}

func bindQFileDialog(h uintptr) *QFileDialog {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QFileDialog_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QFileDialog{h: h, QDialog: &QDialog{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQFileDialog(h uintptr) *QFileDialog { return bindQFileDialog(h) }

func (this *QFileDialog) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFileDialog) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFileDialog(h unsafe.Pointer) *QFileDialog { return bindQFileDialog(uintptr(h)) }
func UnsafeNewQFileDialogOwned(h unsafe.Pointer) *QFileDialog { return newownedQFileDialog(uintptr(h)) }

func (this *QFileDialog) UnsafeRelease() {
	this.h = 0
	if this.QDialog != nil {
		this.QDialog.UnsafeRelease()
	}
}

func NewQFileDialog(parent *QWidget) *QFileDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFileDialog_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQFileDialog(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileDialog_virtIdx, QFileDialog_enable)
	return __o
}

func NewQFileDialog2(parent *QWidget, f WindowType) *QFileDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFileDialog_new2, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(f))
	__o := newownedQFileDialog(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileDialog_virtIdx, QFileDialog_enable)
	return __o
}

func NewQFileDialog3() *QFileDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFileDialog_new3, uintptr(__goid))
	__o := newownedQFileDialog(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileDialog_virtIdx, QFileDialog_enable)
	return __o
}

func NewQFileDialog4(parent *QWidget, caption string) *QFileDialog {
	__goid := nextVirtID()
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_new4, uintptr(__goid), uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)))
	__o := newownedQFileDialog(__r)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileDialog_virtIdx, QFileDialog_enable)
	return __o
}

func NewQFileDialog5(parent *QWidget, caption string, directory string) *QFileDialog {
	__goid := nextVirtID()
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var directory_p uintptr
	if len(directory) > 0 {
		directory_p = uintptr(unsafe.Pointer(unsafe.StringData(directory)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_new5, uintptr(__goid), uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), directory_p, uintptr(len(directory)))
	__o := newownedQFileDialog(__r)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(directory)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileDialog_virtIdx, QFileDialog_enable)
	return __o
}

func NewQFileDialog6(parent *QWidget, caption string, directory string, filter string) *QFileDialog {
	__goid := nextVirtID()
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var directory_p uintptr
	if len(directory) > 0 {
		directory_p = uintptr(unsafe.Pointer(unsafe.StringData(directory)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_new6, uintptr(__goid), uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), directory_p, uintptr(len(directory)), filter_p, uintptr(len(filter)))
	__o := newownedQFileDialog(__r)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(directory)
	runtime.KeepAlive(filter)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileDialog_virtIdx, QFileDialog_enable)
	return __o
}

func QFileDialog_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFileDialog) SetDirectory(directory string) {
	var directory_p uintptr
	if len(directory) > 0 {
		directory_p = uintptr(unsafe.Pointer(unsafe.StringData(directory)))
	}
	purego.SyscallN(addr_QFileDialog_setDirectory, this.h, directory_p, uintptr(len(directory)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(directory)
}

func (this *QFileDialog) SetDirectoryWithDirectory(directory *QDir) {
	purego.SyscallN(addr_QFileDialog_setDirectory2, this.h, uintptr(directory.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(directory)
}

func (this *QFileDialog) Directory() *QDir {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_directory, this.h)
	runtime.KeepAlive(this)
	return newownedQDir(__r)
}

func (this *QFileDialog) SetDirectoryUrl(directory *QUrl) {
	purego.SyscallN(addr_QFileDialog_setDirectoryUrl, this.h, uintptr(directory.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(directory)
}

func (this *QFileDialog) DirectoryUrl() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_directoryUrl, this.h)
	runtime.KeepAlive(this)
	return newownedQUrl(__r)
}

func (this *QFileDialog) SelectFile(filename string) {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	purego.SyscallN(addr_QFileDialog_selectFile, this.h, filename_p, uintptr(len(filename)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filename)
}

func (this *QFileDialog) SelectedFiles() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_selectedFiles, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFileDialog) SelectUrl(url *QUrl) {
	purego.SyscallN(addr_QFileDialog_selectUrl, this.h, uintptr(url.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(url)
}

func (this *QFileDialog) SelectedUrls() []*QUrl {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_selectedUrls, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFileDialog) SetNameFilter(filter string) {
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	purego.SyscallN(addr_QFileDialog_setNameFilter, this.h, filter_p, uintptr(len(filter)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filter)
}

func (this *QFileDialog) SetNameFilters(filters []string) {
	filters_arr := make([]uintptr, len(filters))
	for __i, __e := range filters {
		filters_arr[__i] = pgCString(__e)
	}
	var filters_p uintptr
	if len(filters_arr) > 0 {
		filters_p = uintptr(unsafe.Pointer(&filters_arr[0]))
	}
	purego.SyscallN(addr_QFileDialog_setNameFilters, this.h, filters_p, uintptr(len(filters)))
	runtime.KeepAlive(this)
	for __i := range filters_arr {
		pgFree(filters_arr[__i])
	}
	runtime.KeepAlive(filters_arr)
}

func (this *QFileDialog) NameFilters() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_nameFilters, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFileDialog) SelectNameFilter(filter string) {
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	purego.SyscallN(addr_QFileDialog_selectNameFilter, this.h, filter_p, uintptr(len(filter)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filter)
}

func (this *QFileDialog) SelectedMimeTypeFilter() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_selectedMimeTypeFilter, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileDialog) SelectedNameFilter() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_selectedNameFilter, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileDialog) SetMimeTypeFilters(filters []string) {
	filters_arr := make([]uintptr, len(filters))
	for __i, __e := range filters {
		filters_arr[__i] = pgCString(__e)
	}
	var filters_p uintptr
	if len(filters_arr) > 0 {
		filters_p = uintptr(unsafe.Pointer(&filters_arr[0]))
	}
	purego.SyscallN(addr_QFileDialog_setMimeTypeFilters, this.h, filters_p, uintptr(len(filters)))
	runtime.KeepAlive(this)
	for __i := range filters_arr {
		pgFree(filters_arr[__i])
	}
	runtime.KeepAlive(filters_arr)
}

func (this *QFileDialog) MimeTypeFilters() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_mimeTypeFilters, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFileDialog) SelectMimeTypeFilter(filter string) {
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	purego.SyscallN(addr_QFileDialog_selectMimeTypeFilter, this.h, filter_p, uintptr(len(filter)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filter)
}

func (this *QFileDialog) Filter() QDir__Filter {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_filter, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDir__Filter(__r)
}

func (this *QFileDialog) SetFilter(filters QDir__Filter) {
	purego.SyscallN(addr_QFileDialog_setFilter, this.h, uintptr(filters))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) SetViewMode(mode QFileDialog__ViewMode) {
	purego.SyscallN(addr_QFileDialog_setViewMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) ViewMode() QFileDialog__ViewMode {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_viewMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFileDialog__ViewMode(__r)
}

func (this *QFileDialog) SetFileMode(mode QFileDialog__FileMode) {
	purego.SyscallN(addr_QFileDialog_setFileMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) FileMode() QFileDialog__FileMode {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_fileMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFileDialog__FileMode(__r)
}

func (this *QFileDialog) SetAcceptMode(mode QFileDialog__AcceptMode) {
	purego.SyscallN(addr_QFileDialog_setAcceptMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) AcceptMode() QFileDialog__AcceptMode {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_acceptMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFileDialog__AcceptMode(__r)
}

func (this *QFileDialog) SetSidebarUrls(urls []*QUrl) {
	urls_arr := make([]uintptr, len(urls))
	for __i, __e := range urls {
		urls_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var urls_p uintptr
	if len(urls_arr) > 0 {
		urls_p = uintptr(unsafe.Pointer(&urls_arr[0]))
	}
	purego.SyscallN(addr_QFileDialog_setSidebarUrls, this.h, urls_p, uintptr(len(urls)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(urls)
	runtime.KeepAlive(urls_arr)
}

func (this *QFileDialog) SidebarUrls() []*QUrl {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_sidebarUrls, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFileDialog) SaveState() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_saveState, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QFileDialog) RestoreState(state *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_restoreState, this.h, uintptr(state.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(state)
	return __r != 0
}

func (this *QFileDialog) SetDefaultSuffix(suffix string) {
	var suffix_p uintptr
	if len(suffix) > 0 {
		suffix_p = uintptr(unsafe.Pointer(unsafe.StringData(suffix)))
	}
	purego.SyscallN(addr_QFileDialog_setDefaultSuffix, this.h, suffix_p, uintptr(len(suffix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(suffix)
}

func (this *QFileDialog) DefaultSuffix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_defaultSuffix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileDialog) SetHistory(paths []string) {
	paths_arr := make([]uintptr, len(paths))
	for __i, __e := range paths {
		paths_arr[__i] = pgCString(__e)
	}
	var paths_p uintptr
	if len(paths_arr) > 0 {
		paths_p = uintptr(unsafe.Pointer(&paths_arr[0]))
	}
	purego.SyscallN(addr_QFileDialog_setHistory, this.h, paths_p, uintptr(len(paths)))
	runtime.KeepAlive(this)
	for __i := range paths_arr {
		pgFree(paths_arr[__i])
	}
	runtime.KeepAlive(paths_arr)
}

func (this *QFileDialog) History() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_history, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFileDialog) SetItemDelegate(delegate *QAbstractItemDelegate) {
	purego.SyscallN(addr_QFileDialog_setItemDelegate, this.h, uintptr(delegate.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(delegate)
}

func (this *QFileDialog) ItemDelegate() *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_itemDelegate, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemDelegate(__r)
}

func (this *QFileDialog) SetIconProvider(provider *QAbstractFileIconProvider) {
	purego.SyscallN(addr_QFileDialog_setIconProvider, this.h, uintptr(provider.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(provider)
}

func (this *QFileDialog) IconProvider() *QAbstractFileIconProvider {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_iconProvider, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractFileIconProvider(__r)
}

func (this *QFileDialog) SetLabelText(label QFileDialog__DialogLabel, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QFileDialog_setLabelText, this.h, uintptr(label), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QFileDialog) LabelText(label QFileDialog__DialogLabel) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_labelText, this.h, uintptr(label), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileDialog) SetSupportedSchemes(schemes []string) {
	schemes_arr := make([]uintptr, len(schemes))
	for __i, __e := range schemes {
		schemes_arr[__i] = pgCString(__e)
	}
	var schemes_p uintptr
	if len(schemes_arr) > 0 {
		schemes_p = uintptr(unsafe.Pointer(&schemes_arr[0]))
	}
	purego.SyscallN(addr_QFileDialog_setSupportedSchemes, this.h, schemes_p, uintptr(len(schemes)))
	runtime.KeepAlive(this)
	for __i := range schemes_arr {
		pgFree(schemes_arr[__i])
	}
	runtime.KeepAlive(schemes_arr)
}

func (this *QFileDialog) SupportedSchemes() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_supportedSchemes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFileDialog) SetProxyModel(model *QAbstractProxyModel) {
	purego.SyscallN(addr_QFileDialog_setProxyModel, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QFileDialog) ProxyModel() *QAbstractProxyModel {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_proxyModel, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractProxyModel(__r)
}

func (this *QFileDialog) SetOption(option QFileDialog__Option) {
	purego.SyscallN(addr_QFileDialog_setOption, this.h, uintptr(option))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) TestOption(option QFileDialog__Option) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_testOption, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileDialog) SetOptions(options QFileDialog__Option) {
	purego.SyscallN(addr_QFileDialog_setOptions, this.h, uintptr(options))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) Options() QFileDialog__Option {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_options, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFileDialog__Option(__r)
}

func (this *QFileDialog) Open(receiver *QObject, member string) {
	member_b := append([]byte(member), 0)
	purego.SyscallN(addr_QFileDialog_open, this.h, uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
}

func (this *QFileDialog) FileSelected(file string) {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	purego.SyscallN(addr_QFileDialog_fileSelected, this.h, file_p, uintptr(len(file)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(file)
}

func (this *QFileDialog) FilesSelected(files []string) {
	files_arr := make([]uintptr, len(files))
	for __i, __e := range files {
		files_arr[__i] = pgCString(__e)
	}
	var files_p uintptr
	if len(files_arr) > 0 {
		files_p = uintptr(unsafe.Pointer(&files_arr[0]))
	}
	purego.SyscallN(addr_QFileDialog_filesSelected, this.h, files_p, uintptr(len(files)))
	runtime.KeepAlive(this)
	for __i := range files_arr {
		pgFree(files_arr[__i])
	}
	runtime.KeepAlive(files_arr)
}

func (this *QFileDialog) CurrentChanged(path string) {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	purego.SyscallN(addr_QFileDialog_currentChanged, this.h, path_p, uintptr(len(path)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QFileDialog) DirectoryEntered(directory string) {
	var directory_p uintptr
	if len(directory) > 0 {
		directory_p = uintptr(unsafe.Pointer(unsafe.StringData(directory)))
	}
	purego.SyscallN(addr_QFileDialog_directoryEntered, this.h, directory_p, uintptr(len(directory)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(directory)
}

func (this *QFileDialog) UrlSelected(url *QUrl) {
	purego.SyscallN(addr_QFileDialog_urlSelected, this.h, uintptr(url.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(url)
}

func (this *QFileDialog) UrlsSelected(urls []*QUrl) {
	urls_arr := make([]uintptr, len(urls))
	for __i, __e := range urls {
		urls_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var urls_p uintptr
	if len(urls_arr) > 0 {
		urls_p = uintptr(unsafe.Pointer(&urls_arr[0]))
	}
	purego.SyscallN(addr_QFileDialog_urlsSelected, this.h, urls_p, uintptr(len(urls)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(urls)
	runtime.KeepAlive(urls_arr)
}

func (this *QFileDialog) CurrentUrlChanged(url *QUrl) {
	purego.SyscallN(addr_QFileDialog_currentUrlChanged, this.h, uintptr(url.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(url)
}

func (this *QFileDialog) DirectoryUrlEntered(directory *QUrl) {
	purego.SyscallN(addr_QFileDialog_directoryUrlEntered, this.h, uintptr(directory.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(directory)
}

func (this *QFileDialog) FilterSelected(filter string) {
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	purego.SyscallN(addr_QFileDialog_filterSelected, this.h, filter_p, uintptr(len(filter)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filter)
}

func QFileDialog_GetOpenFileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QFileDialog_GetOpenFileUrl() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrl)
	return newownedQUrl(__r)
}

func QFileDialog_GetSaveFileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QFileDialog_GetSaveFileUrl() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileUrl)
	return newownedQUrl(__r)
}

func QFileDialog_GetExistingDirectory() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getExistingDirectory, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QFileDialog_GetExistingDirectoryUrl() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getExistingDirectoryUrl)
	return newownedQUrl(__r)
}

func QFileDialog_GetOpenFileNames() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileNames, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileUrls() []*QUrl {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrls, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_SaveFileContent(fileContent *QByteArray, fileNameHint string) {
	var fileNameHint_p uintptr
	if len(fileNameHint) > 0 {
		fileNameHint_p = uintptr(unsafe.Pointer(unsafe.StringData(fileNameHint)))
	}
	purego.SyscallN(addr_QFileDialog_saveFileContent, uintptr(fileContent.UnsafePointer()), fileNameHint_p, uintptr(len(fileNameHint)))
	runtime.KeepAlive(fileContent)
	runtime.KeepAlive(fileNameHint)
}

func (this *QFileDialog) AdjustPosition(arg1 *QWidget) {
	purego.SyscallN(addr_QFileDialog_adjustPosition, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QFileDialog_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QFileDialog_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QFileDialog_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_mouseGrabber)
	return bindQWidget(__r)
}

func QFileDialog_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_keyboardGrabber)
	return bindQWidget(__r)
}

func QFileDialog_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QFileDialog_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QFileDialog) UpdateMicroFocus() {
	purego.SyscallN(addr_QFileDialog_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileDialog) Create() {
	purego.SyscallN(addr_QFileDialog_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileDialog) Destroy() {
	purego.SyscallN(addr_QFileDialog_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileDialog) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileDialog) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QFileDialog_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFileDialog_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFileDialog_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFileDialog) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QFileDialog) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFileDialog) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QFileDialog) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QFileDialog_DevicePixelRatioFScale() float64 {
	__r := pg_QFileDialog_devicePixelRatioFScale()
	return __r
}

func QFileDialog_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QFileDialog_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QFileDialog) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QFileDialog_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QFileDialog_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFileDialog_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QFileDialog) SetOption2(option QFileDialog__Option, on bool) {
	purego.SyscallN(addr_QFileDialog_setOption2, this.h, uintptr(option), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func QFileDialog_GetOpenFileNameWithParent(parent *QWidget) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileName2, uintptr(parent.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	return pgStr(__r, __n)
}

func QFileDialog_GetOpenFileName2(parent *QWidget, caption string) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileName3, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	return pgStr(__r, __n)
}

func QFileDialog_GetOpenFileName3(parent *QWidget, caption string, dir string) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileName4, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	return pgStr(__r, __n)
}

func QFileDialog_GetOpenFileName4(parent *QWidget, caption string, dir string, filter string) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileName5, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), filter_p, uintptr(len(filter)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	return pgStr(__r, __n)
}

func QFileDialog_GetOpenFileName5(parent *QWidget, caption string, dir string, filter string, selectedFilter *QString) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileName6, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	return pgStr(__r, __n)
}

func QFileDialog_GetOpenFileName6(parent *QWidget, caption string, dir string, filter string, selectedFilter *QString, options QFileDialog__Option) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileName7, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	return pgStr(__r, __n)
}

func QFileDialog_GetOpenFileUrlWithParent(parent *QWidget) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrl2, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(parent)
	return newownedQUrl(__r)
}

func QFileDialog_GetOpenFileUrl2(parent *QWidget, caption string) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrl3, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	return newownedQUrl(__r)
}

func QFileDialog_GetOpenFileUrl3(parent *QWidget, caption string, dir *QUrl) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrl4, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	return newownedQUrl(__r)
}

func QFileDialog_GetOpenFileUrl4(parent *QWidget, caption string, dir *QUrl, filter string) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrl5, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	return newownedQUrl(__r)
}

func QFileDialog_GetOpenFileUrl5(parent *QWidget, caption string, dir *QUrl, filter string, selectedFilter *QString) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrl6, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	return newownedQUrl(__r)
}

func QFileDialog_GetOpenFileUrl6(parent *QWidget, caption string, dir *QUrl, filter string, selectedFilter *QString, options QFileDialog__Option) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrl7, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	return newownedQUrl(__r)
}

func QFileDialog_GetOpenFileUrl7(parent *QWidget, caption string, dir *QUrl, filter string, selectedFilter *QString, options QFileDialog__Option, supportedSchemes []string) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	supportedSchemes_arr := make([]uintptr, len(supportedSchemes))
	for __i, __e := range supportedSchemes {
		supportedSchemes_arr[__i] = pgCString(__e)
	}
	var supportedSchemes_p uintptr
	if len(supportedSchemes_arr) > 0 {
		supportedSchemes_p = uintptr(unsafe.Pointer(&supportedSchemes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrl8, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(options), supportedSchemes_p, uintptr(len(supportedSchemes)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	for __i := range supportedSchemes_arr {
		pgFree(supportedSchemes_arr[__i])
	}
	runtime.KeepAlive(supportedSchemes_arr)
	return newownedQUrl(__r)
}

func QFileDialog_GetSaveFileNameWithParent(parent *QWidget) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileName2, uintptr(parent.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	return pgStr(__r, __n)
}

func QFileDialog_GetSaveFileName2(parent *QWidget, caption string) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileName3, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	return pgStr(__r, __n)
}

func QFileDialog_GetSaveFileName3(parent *QWidget, caption string, dir string) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileName4, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	return pgStr(__r, __n)
}

func QFileDialog_GetSaveFileName4(parent *QWidget, caption string, dir string, filter string) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileName5, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), filter_p, uintptr(len(filter)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	return pgStr(__r, __n)
}

func QFileDialog_GetSaveFileName5(parent *QWidget, caption string, dir string, filter string, selectedFilter *QString) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileName6, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	return pgStr(__r, __n)
}

func QFileDialog_GetSaveFileName6(parent *QWidget, caption string, dir string, filter string, selectedFilter *QString, options QFileDialog__Option) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileName7, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	return pgStr(__r, __n)
}

func QFileDialog_GetSaveFileUrlWithParent(parent *QWidget) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileUrl2, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(parent)
	return newownedQUrl(__r)
}

func QFileDialog_GetSaveFileUrl2(parent *QWidget, caption string) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileUrl3, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	return newownedQUrl(__r)
}

func QFileDialog_GetSaveFileUrl3(parent *QWidget, caption string, dir *QUrl) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileUrl4, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	return newownedQUrl(__r)
}

func QFileDialog_GetSaveFileUrl4(parent *QWidget, caption string, dir *QUrl, filter string) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileUrl5, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	return newownedQUrl(__r)
}

func QFileDialog_GetSaveFileUrl5(parent *QWidget, caption string, dir *QUrl, filter string, selectedFilter *QString) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileUrl6, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	return newownedQUrl(__r)
}

func QFileDialog_GetSaveFileUrl6(parent *QWidget, caption string, dir *QUrl, filter string, selectedFilter *QString, options QFileDialog__Option) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileUrl7, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	return newownedQUrl(__r)
}

func QFileDialog_GetSaveFileUrl7(parent *QWidget, caption string, dir *QUrl, filter string, selectedFilter *QString, options QFileDialog__Option, supportedSchemes []string) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	supportedSchemes_arr := make([]uintptr, len(supportedSchemes))
	for __i, __e := range supportedSchemes {
		supportedSchemes_arr[__i] = pgCString(__e)
	}
	var supportedSchemes_p uintptr
	if len(supportedSchemes_arr) > 0 {
		supportedSchemes_p = uintptr(unsafe.Pointer(&supportedSchemes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getSaveFileUrl8, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(options), supportedSchemes_p, uintptr(len(supportedSchemes)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	for __i := range supportedSchemes_arr {
		pgFree(supportedSchemes_arr[__i])
	}
	runtime.KeepAlive(supportedSchemes_arr)
	return newownedQUrl(__r)
}

func QFileDialog_GetExistingDirectoryWithParent(parent *QWidget) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getExistingDirectory2, uintptr(parent.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	return pgStr(__r, __n)
}

func QFileDialog_GetExistingDirectory2(parent *QWidget, caption string) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getExistingDirectory3, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	return pgStr(__r, __n)
}

func QFileDialog_GetExistingDirectory3(parent *QWidget, caption string, dir string) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getExistingDirectory4, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	return pgStr(__r, __n)
}

func QFileDialog_GetExistingDirectory4(parent *QWidget, caption string, dir string, options QFileDialog__Option) string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getExistingDirectory5, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	return pgStr(__r, __n)
}

func QFileDialog_GetExistingDirectoryUrlWithParent(parent *QWidget) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getExistingDirectoryUrl2, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(parent)
	return newownedQUrl(__r)
}

func QFileDialog_GetExistingDirectoryUrl2(parent *QWidget, caption string) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getExistingDirectoryUrl3, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	return newownedQUrl(__r)
}

func QFileDialog_GetExistingDirectoryUrl3(parent *QWidget, caption string, dir *QUrl) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getExistingDirectoryUrl4, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	return newownedQUrl(__r)
}

func QFileDialog_GetExistingDirectoryUrl4(parent *QWidget, caption string, dir *QUrl, options QFileDialog__Option) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getExistingDirectoryUrl5, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	return newownedQUrl(__r)
}

func QFileDialog_GetExistingDirectoryUrl5(parent *QWidget, caption string, dir *QUrl, options QFileDialog__Option, supportedSchemes []string) *QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	supportedSchemes_arr := make([]uintptr, len(supportedSchemes))
	for __i, __e := range supportedSchemes {
		supportedSchemes_arr[__i] = pgCString(__e)
	}
	var supportedSchemes_p uintptr
	if len(supportedSchemes_arr) > 0 {
		supportedSchemes_p = uintptr(unsafe.Pointer(&supportedSchemes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getExistingDirectoryUrl6, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), uintptr(options), supportedSchemes_p, uintptr(len(supportedSchemes)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	for __i := range supportedSchemes_arr {
		pgFree(supportedSchemes_arr[__i])
	}
	runtime.KeepAlive(supportedSchemes_arr)
	return newownedQUrl(__r)
}

func QFileDialog_GetOpenFileNamesWithParent(parent *QWidget) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileNames2, uintptr(parent.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileNames2(parent *QWidget, caption string) []string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileNames3, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileNames3(parent *QWidget, caption string, dir string) []string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileNames4, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileNames4(parent *QWidget, caption string, dir string, filter string) []string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileNames5, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), filter_p, uintptr(len(filter)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileNames5(parent *QWidget, caption string, dir string, filter string, selectedFilter *QString) []string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileNames6, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileNames6(parent *QWidget, caption string, dir string, filter string, selectedFilter *QString, options QFileDialog__Option) []string {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileNames7, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), dir_p, uintptr(len(dir)), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileUrlsWithParent(parent *QWidget) []*QUrl {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrls2, uintptr(parent.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileUrls2(parent *QWidget, caption string) []*QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrls3, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileUrls3(parent *QWidget, caption string, dir *QUrl) []*QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrls4, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileUrls4(parent *QWidget, caption string, dir *QUrl, filter string) []*QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrls5, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileUrls5(parent *QWidget, caption string, dir *QUrl, filter string, selectedFilter *QString) []*QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrls6, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileUrls6(parent *QWidget, caption string, dir *QUrl, filter string, selectedFilter *QString, options QFileDialog__Option) []*QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrls7, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_GetOpenFileUrls7(parent *QWidget, caption string, dir *QUrl, filter string, selectedFilter *QString, options QFileDialog__Option, supportedSchemes []string) []*QUrl {
	var caption_p uintptr
	if len(caption) > 0 {
		caption_p = uintptr(unsafe.Pointer(unsafe.StringData(caption)))
	}
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	supportedSchemes_arr := make([]uintptr, len(supportedSchemes))
	for __i, __e := range supportedSchemes {
		supportedSchemes_arr[__i] = pgCString(__e)
	}
	var supportedSchemes_p uintptr
	if len(supportedSchemes_arr) > 0 {
		supportedSchemes_p = uintptr(unsafe.Pointer(&supportedSchemes_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDialog_getOpenFileUrls8, uintptr(parent.UnsafePointer()), caption_p, uintptr(len(caption)), uintptr(dir.UnsafePointer()), filter_p, uintptr(len(filter)), uintptr(selectedFilter.UnsafePointer()), uintptr(options), supportedSchemes_p, uintptr(len(supportedSchemes)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(caption)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(filter)
	runtime.KeepAlive(selectedFilter)
	for __i := range supportedSchemes_arr {
		pgFree(supportedSchemes_arr[__i])
	}
	runtime.KeepAlive(supportedSchemes_arr)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileDialog_SaveFileContent2(fileContent *QByteArray, fileNameHint string, parent *QWidget) {
	var fileNameHint_p uintptr
	if len(fileNameHint) > 0 {
		fileNameHint_p = uintptr(unsafe.Pointer(unsafe.StringData(fileNameHint)))
	}
	purego.SyscallN(addr_QFileDialog_saveFileContent2, uintptr(fileContent.UnsafePointer()), fileNameHint_p, uintptr(len(fileNameHint)), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(fileContent)
	runtime.KeepAlive(fileNameHint)
	runtime.KeepAlive(parent)
}

func QFileDialog_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QFileDialog_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QFileDialog) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QFileDialog_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QFileDialog_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QFileDialog_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QFileDialog_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QFileDialog_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QFileDialog_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QFileDialog_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFileDialog_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFileDialog_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFileDialog) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QFileDialog) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QFileDialog_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) DoneDefault(result int) {
	purego.SyscallN(addr_QFileDialog_done_default, this.h, uintptr(result))
	runtime.KeepAlive(this)
}

func (this *QFileDialog) AcceptDefault() {
	purego.SyscallN(addr_QFileDialog_accept_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileDialog) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QFileDialog_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFileDialog) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFileDialog) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFileDialog) ExecDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_exec_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFileDialog) RejectDefault() {
	purego.SyscallN(addr_QFileDialog_reject_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileDialog) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QFileDialog_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFileDialog) CloseEventDefault(arg1 *QCloseEvent) {
	purego.SyscallN(addr_QFileDialog_closeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFileDialog) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QFileDialog_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFileDialog) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QFileDialog_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFileDialog) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QFileDialog_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFileDialog) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QFileDialog) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFileDialog) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFileDialog) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileDialog) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QFileDialog) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFileDialog) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFileDialog_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFileDialog_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFileDialog_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFileDialog_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QFileDialog_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QFileDialog_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QFileDialog_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QFileDialog_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QFileDialog_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFileDialog_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QFileDialog_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QFileDialog_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QFileDialog_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QFileDialog_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QFileDialog_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QFileDialog_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QFileDialog_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QFileDialog_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QFileDialog_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFileDialog) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QFileDialog_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QFileDialog) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QFileDialog) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QFileDialog) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QFileDialog_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFileDialog) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QFileDialog_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QFileDialog) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDialog_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileDialog) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QFileDialog_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QFileDialog_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFileDialog_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileDialog) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFileDialog_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFileDialog) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFileDialog_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFileDialog) OnFileSelected(fn func(string)) int64 {
	id := registerCallback(func(file uintptr, file_len int64) { fn(pgStrNoFree(file, file_len)) })
	pg_QFileDialog_connect_fileSelected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnFilesSelected(fn func([]*QString)) int64 {
	id := registerCallback(func(files uintptr, files_len int64) { fn(func() []*QString { if files == 0 || files_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(files)), int(files_len)); __o := make([]*QString, int(files_len)); for __i := range __o { __o[__i] = UnsafeNewQString(__sl[__i]) }; return __o }()) })
	pg_QFileDialog_connect_filesSelected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnCurrentChanged(fn func(string)) int64 {
	id := registerCallback(func(path uintptr, path_len int64) { fn(pgStrNoFree(path, path_len)) })
	pg_QFileDialog_connect_currentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnDirectoryEntered(fn func(string)) int64 {
	id := registerCallback(func(directory uintptr, directory_len int64) { fn(pgStrNoFree(directory, directory_len)) })
	pg_QFileDialog_connect_directoryEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnUrlSelected(fn func(*QUrl)) int64 {
	id := registerCallback(func(url uintptr) { fn(UnsafeNewQUrl(unsafe.Pointer(url))) })
	pg_QFileDialog_connect_urlSelected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnUrlsSelected(fn func([]*QUrl)) int64 {
	id := registerCallback(func(urls uintptr, urls_len int64) { fn(func() []*QUrl { if urls == 0 || urls_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(urls)), int(urls_len)); __o := make([]*QUrl, int(urls_len)); for __i := range __o { __o[__i] = UnsafeNewQUrl(__sl[__i]) }; return __o }()) })
	pg_QFileDialog_connect_urlsSelected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnCurrentUrlChanged(fn func(*QUrl)) int64 {
	id := registerCallback(func(url uintptr) { fn(UnsafeNewQUrl(unsafe.Pointer(url))) })
	pg_QFileDialog_connect_currentUrlChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnDirectoryUrlEntered(fn func(*QUrl)) int64 {
	id := registerCallback(func(directory uintptr) { fn(UnsafeNewQUrl(unsafe.Pointer(directory))) })
	pg_QFileDialog_connect_directoryUrlEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnFilterSelected(fn func(string)) int64 {
	id := registerCallback(func(filter uintptr, filter_len int64) { fn(pgStrNoFree(filter, filter_len)) })
	pg_QFileDialog_connect_filterSelected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnFinished(fn func(int)) int64 {
	id := registerCallback(func(result int32) { fn(int(result)) })
	pg_QFileDialog_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnAccepted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileDialog_connect_accepted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnRejected(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileDialog_connect_rejected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QFileDialog_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QFileDialog_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QFileDialog_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QFileDialog_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileDialog_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFileDialog_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFileDialog_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDialog) Disconnect(handle int64) {
	pg_QFileDialog__disconnect(handle)
	unregisterCallback(handle)
}

func QFileDialog_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 2110182719375245575:
		return 1
	case -2570997490807640431:
		return 2
	case 9106131928185084681:
		return 3
	case 4959376063826165943:
		return 4
	case 729690261457684101:
		return 5
	case 7049264599837532491:
		return 6
	case -1745322561779819684:
		return 7
	case 2875220963396170150:
		return 8
	case 1023791977679793185:
		return 9
	case -1989937050019464783:
		return 10
	case -6825369174751916:
		return 11
	case -4302917285655014661:
		return 12
	case -6698043222227006015:
		return 13
	case 5217623514630858297:
		return 14
	case -8659679893868524044:
		return 15
	case 4311534060495737339:
		return 16
	case -8266762397995423915:
		return 17
	case 6398007404756346381:
		return 18
	case 1237236456312824991:
		return 19
	case -332090898078543319:
		return 20
	case -4415920049297795213:
		return 21
	case -4029073905697023887:
		return 22
	case -4255494004099806581:
		return 23
	case 1919988340934511316:
		return 24
	case 772598475950050299:
		return 25
	case 975043767033552552:
		return 26
	case -385596010576048993:
		return 27
	case 6316270346668237895:
		return 28
	case -551070009623711640:
		return 29
	case 5290442163167458465:
		return 30
	case -9161949150489525396:
		return 31
	case -212068866860701141:
		return 32
	case 1697444140088555581:
		return 33
	case 206122715919917589:
		return 34
	case -1781545137352745978:
		return 35
	case -3139904158373725006:
		return 36
	case 9107058682754591716:
		return 37
	case -199718707876719129:
		return 38
	case -2039166586246443706:
		return 39
	case 427629649267311777:
		return 40
	case 6477500296980050210:
		return 41
	case 4127240633310655048:
		return 42
	case 5792771328741774991:
		return 43
	case 5438819661942843222:
		return 44
	case 7345853097340709638:
		return 45
	case 5898328517505784191:
		return 46
	case -5409927581812374498:
		return 47
	case -1185315367526857833:
		return 48
	case 3769868229123868474:
		return 49
	case 7379029982170503008:
		return 50
	case -568369772982215962:
		return 51
	}
	return -1
}

func QFileDialog_enable(h uintptr, idx int32) { pg_QFileDialog__set_override(h, idx) }

func (this *QFileDialog) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QFileDialog) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QFileDialog) OnDone(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2570997490807640431, func(result int32) { fn(int(result)) }) // slot "done"
}

func (this *QFileDialog) OnAccept(fn func()) {
	registerVirt(this.UnsafeGoID(), 9106131928185084681, func() { fn() }) // slot "accept"
}

func (this *QFileDialog) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QFileDialog) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QFileDialog) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QFileDialog) OnExec(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1745322561779819684, func() int32 { return int32(fn()) }) // slot "exec"
}

func (this *QFileDialog) OnReject(fn func()) {
	registerVirt(this.UnsafeGoID(), 2875220963396170150, func() { fn() }) // slot "reject"
}

func (this *QFileDialog) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QFileDialog) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(arg1 uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(arg1))) }) // slot "closeEvent"
}

func (this *QFileDialog) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QFileDialog) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QFileDialog) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QFileDialog) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QFileDialog) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QFileDialog) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QFileDialog) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QFileDialog) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QFileDialog) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QFileDialog) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QFileDialog) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QFileDialog) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QFileDialog) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QFileDialog) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QFileDialog) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QFileDialog) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QFileDialog) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QFileDialog) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QFileDialog) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QFileDialog) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QFileDialog) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QFileDialog) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QFileDialog) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QFileDialog) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QFileDialog) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QFileDialog) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QFileDialog) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QFileDialog) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QFileDialog) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QFileDialog) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QFileDialog) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QFileDialog) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QFileDialog) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QFileDialog) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QFileDialog) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QFileDialog) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QFileDialog) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QFileDialog) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QFileDialog) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QFileDialog) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QFileDialog) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QFileDialog) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFileDialog_delete(this.h)
	}
	this.UnsafeRelease()
}
