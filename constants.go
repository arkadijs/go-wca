package wca

const (
	AUDCLNT_SHAREMODE_EXCLUSIVE = 0x0
	AUDCLNT_SHAREMODE_SHARED    = 0x1
)

const (
	ENDPOINT_SYSFX_ENABLED  = 0x00000000
	ENDPOINT_SYSFX_DISABLED = 0x00000001
)

const (
	DEVICE_STATE_ACTIVE     = 0x00000001
	DEVICE_STATE_DISABLED   = 0x00000002
	DEVICE_STATE_NOTPRESENT = 0x00000004
	DEVICE_STATE_UNPLUGGED  = 0x00000008
	DEVICE_STATEMASK_ALL    = 0x0000000F
)

const (
	ERender              = 0x0
	ECapture             = 0x1
	EAll                 = 0x2
	EDataFlow_enum_count = 0x3
)

const (
	STGM_READ       = 0x0
	STGM_WRITE      = 0x1
	STGM_READ_WRITE = 0x2
)

const (
	CLSCTX_INPROC_SERVER          = 0x1
	CLSCTX_INPROC_HANDLER         = 0x2
	CLSCTX_LOCAL_SERVER           = 0x4
	CLSCTX_INPROC_SERVER16        = 0x8
	CLSCTX_REMOTE_SERVER          = 0x10
	CLSCTX_INPROC_HANDLER16       = 0x20
	CLSCTX_RESERVED1              = 0x40
	CLSCTX_RESERVED2              = 0x80
	CLSCTX_RESERVED3              = 0x100
	CLSCTX_RESERVED4              = 0x200
	CLSCTX_NO_CODE_DOWNLOAD       = 0x400
	CLSCTX_RESERVED5              = 0x800
	CLSCTX_NO_CUSTOM_MARSHAL      = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD   = 0x2000
	CLSCTX_NO_FAILURE_LOG         = 0x4000
	CLSCTX_DISABLE_AAA            = 0x8000
	CLSCTX_ENABLE_AAA             = 0x10000
	CLSCTX_FROM_DEFAULT_CONTEXT   = 0x20000
	CLSCTX_ACTIVATE_32_BIT_SERVER = 0x40000
	CLSCTX_ACTIVATE_64_BIT_SERVER = 0x80000
	CLSCTX_ENABLE_CLOAKING        = 0x100000
	CLSCTX_APPCONTAINER           = 0x400000
	CLSCTX_ACTIVATE_AAA_AS_IU     = 0x800000
	CLSCTX_PS_DLL                 = 0x80000000
)
