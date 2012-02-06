// godefs -gnl80211 -

// MACHINE GENERATED - DO NOT EDIT.

package nl80211

// Constants
const (
	CMD_UNSPEC                          = 0
	CMD_GET_WIPHY                       = 0x1
	CMD_SET_WIPHY                       = 0x2
	CMD_NEW_WIPHY                       = 0x3
	CMD_DEL_WIPHY                       = 0x4
	CMD_GET_INTERFACE                   = 0x5
	CMD_SET_INTERFACE                   = 0x6
	CMD_NEW_INTERFACE                   = 0x7
	CMD_DEL_INTERFACE                   = 0x8
	CMD_GET_KEY                         = 0x9
	CMD_SET_KEY                         = 0xa
	CMD_NEW_KEY                         = 0xb
	CMD_DEL_KEY                         = 0xc
	CMD_GET_BEACON                      = 0xd
	CMD_SET_BEACON                      = 0xe
	CMD_NEW_BEACON                      = 0xf
	CMD_DEL_BEACON                      = 0x10
	CMD_GET_STATION                     = 0x11
	CMD_SET_STATION                     = 0x12
	CMD_NEW_STATION                     = 0x13
	CMD_DEL_STATION                     = 0x14
	CMD_GET_MPATH                       = 0x15
	CMD_SET_MPATH                       = 0x16
	CMD_NEW_MPATH                       = 0x17
	CMD_DEL_MPATH                       = 0x18
	CMD_SET_BSS                         = 0x19
	CMD_SET_REG                         = 0x1a
	CMD_REQ_SET_REG                     = 0x1b
	CMD_GET_MESH_CONFIG                 = 0x1c
	CMD_SET_MESH_CONFIG                 = 0x1d
	CMD_SET_MGMT_EXTRA_IE               = 0x1e
	CMD_GET_REG                         = 0x1f
	CMD_GET_SCAN                        = 0x20
	CMD_TRIGGER_SCAN                    = 0x21
	CMD_NEW_SCAN_RESULTS                = 0x22
	CMD_SCAN_ABORTED                    = 0x23
	CMD_REG_CHANGE                      = 0x24
	CMD_AUTHENTICATE                    = 0x25
	CMD_ASSOCIATE                       = 0x26
	CMD_DEAUTHENTICATE                  = 0x27
	CMD_DISASSOCIATE                    = 0x28
	CMD_MICHAEL_MIC_FAILURE             = 0x29
	CMD_REG_BEACON_HINT                 = 0x2a
	CMD_JOIN_IBSS                       = 0x2b
	CMD_LEAVE_IBSS                      = 0x2c
	CMD_TESTMODE                        = 0x2d
	CMD_CONNECT                         = 0x2e
	CMD_ROAM                            = 0x2f
	CMD_DISCONNECT                      = 0x30
	CMD_SET_WIPHY_NETNS                 = 0x31
	CMD_GET_SURVEY                      = 0x32
	CMD_NEW_SURVEY_RESULTS              = 0x33
	CMD_SET_PMKSA                       = 0x34
	CMD_DEL_PMKSA                       = 0x35
	CMD_FLUSH_PMKSA                     = 0x36
	CMD_REMAIN_ON_CHANNEL               = 0x37
	CMD_CANCEL_REMAIN_ON_CHANNEL        = 0x38
	CMD_SET_TX_BITRATE_MASK             = 0x39
	CMD_REGISTER_FRAME                  = 0x3a
	CMD_REGISTER_ACTION                 = 0x3a
	CMD_FRAME                           = 0x3b
	CMD_ACTION                          = 0x3b
	CMD_FRAME_TX_STATUS                 = 0x3c
	CMD_ACTION_TX_STATUS                = 0x3c
	CMD_SET_POWER_SAVE                  = 0x3d
	CMD_GET_POWER_SAVE                  = 0x3e
	CMD_SET_CQM                         = 0x3f
	CMD_NOTIFY_CQM                      = 0x40
	CMD_SET_CHANNEL                     = 0x41
	CMD_SET_WDS_PEER                    = 0x42
	CMD_FRAME_WAIT_CANCEL               = 0x43
	CMD_JOIN_MESH                       = 0x44
	CMD_LEAVE_MESH                      = 0x45
	CMD_UNPROT_DEAUTHENTICATE           = 0x46
	CMD_UNPROT_DISASSOCIATE             = 0x47
	CMD_NEW_PEER_CANDIDATE              = 0x48
	CMD_GET_WOWLAN                      = 0x49
	CMD_SET_WOWLAN                      = 0x4a
	CMD_START_SCHED_SCAN                = 0x4b
	CMD_STOP_SCHED_SCAN                 = 0x4c
	CMD_SCHED_SCAN_RESULTS              = 0x4d
	CMD_SCHED_SCAN_STOPPED              = 0x4e
	__CMD_AFTER_LAST                    = 0x4f
	CMD_MAX                             = 0x4e
	ATTR_UNSPEC                         = 0
	ATTR_WIPHY                          = 0x1
	ATTR_WIPHY_NAME                     = 0x2
	ATTR_IFINDEX                        = 0x3
	ATTR_IFNAME                         = 0x4
	ATTR_IFTYPE                         = 0x5
	ATTR_MAC                            = 0x6
	ATTR_KEY_DATA                       = 0x7
	ATTR_KEY_IDX                        = 0x8
	ATTR_KEY_CIPHER                     = 0x9
	ATTR_KEY_SEQ                        = 0xa
	ATTR_KEY_DEFAULT                    = 0xb
	ATTR_BEACON_INTERVAL                = 0xc
	ATTR_DTIM_PERIOD                    = 0xd
	ATTR_BEACON_HEAD                    = 0xe
	ATTR_BEACON_TAIL                    = 0xf
	ATTR_STA_AID                        = 0x10
	ATTR_STA_FLAGS                      = 0x11
	ATTR_STA_LISTEN_INTERVAL            = 0x12
	ATTR_STA_SUPPORTED_RATES            = 0x13
	ATTR_STA_VLAN                       = 0x14
	ATTR_STA_INFO                       = 0x15
	ATTR_WIPHY_BANDS                    = 0x16
	ATTR_MNTR_FLAGS                     = 0x17
	ATTR_MESH_ID                        = 0x18
	ATTR_STA_PLINK_ACTION               = 0x19
	ATTR_MPATH_NEXT_HOP                 = 0x1a
	ATTR_MPATH_INFO                     = 0x1b
	ATTR_BSS_CTS_PROT                   = 0x1c
	ATTR_BSS_SHORT_PREAMBLE             = 0x1d
	ATTR_BSS_SHORT_SLOT_TIME            = 0x1e
	ATTR_HT_CAPABILITY                  = 0x1f
	ATTR_SUPPORTED_IFTYPES              = 0x20
	ATTR_REG_ALPHA2                     = 0x21
	ATTR_REG_RULES                      = 0x22
	ATTR_MESH_CONFIG                    = 0x23
	ATTR_BSS_BASIC_RATES                = 0x24
	ATTR_WIPHY_TXQ_PARAMS               = 0x25
	ATTR_WIPHY_FREQ                     = 0x26
	ATTR_WIPHY_CHANNEL_TYPE             = 0x27
	ATTR_KEY_DEFAULT_MGMT               = 0x28
	ATTR_MGMT_SUBTYPE                   = 0x29
	ATTR_IE                             = 0x2a
	ATTR_MAX_NUM_SCAN_SSIDS             = 0x2b
	ATTR_SCAN_FREQUENCIES               = 0x2c
	ATTR_SCAN_SSIDS                     = 0x2d
	ATTR_GENERATION                     = 0x2e
	ATTR_BSS                            = 0x2f
	ATTR_REG_INITIATOR                  = 0x30
	ATTR_REG_TYPE                       = 0x31
	ATTR_SUPPORTED_COMMANDS             = 0x32
	ATTR_FRAME                          = 0x33
	ATTR_SSID                           = 0x34
	ATTR_AUTH_TYPE                      = 0x35
	ATTR_REASON_CODE                    = 0x36
	ATTR_KEY_TYPE                       = 0x37
	ATTR_MAX_SCAN_IE_LEN                = 0x38
	ATTR_CIPHER_SUITES                  = 0x39
	ATTR_FREQ_BEFORE                    = 0x3a
	ATTR_FREQ_AFTER                     = 0x3b
	ATTR_FREQ_FIXED                     = 0x3c
	ATTR_WIPHY_RETRY_SHORT              = 0x3d
	ATTR_WIPHY_RETRY_LONG               = 0x3e
	ATTR_WIPHY_FRAG_THRESHOLD           = 0x3f
	ATTR_WIPHY_RTS_THRESHOLD            = 0x40
	ATTR_TIMED_OUT                      = 0x41
	ATTR_USE_MFP                        = 0x42
	ATTR_STA_FLAGS2                     = 0x43
	ATTR_CONTROL_PORT                   = 0x44
	ATTR_TESTDATA                       = 0x45
	ATTR_PRIVACY                        = 0x46
	ATTR_DISCONNECTED_BY_AP             = 0x47
	ATTR_STATUS_CODE                    = 0x48
	ATTR_CIPHER_SUITES_PAIRWISE         = 0x49
	ATTR_CIPHER_SUITE_GROUP             = 0x4a
	ATTR_WPA_VERSIONS                   = 0x4b
	ATTR_AKM_SUITES                     = 0x4c
	ATTR_REQ_IE                         = 0x4d
	ATTR_RESP_IE                        = 0x4e
	ATTR_PREV_BSSID                     = 0x4f
	ATTR_KEY                            = 0x50
	ATTR_KEYS                           = 0x51
	ATTR_PID                            = 0x52
	ATTR_4ADDR                          = 0x53
	ATTR_SURVEY_INFO                    = 0x54
	ATTR_PMKID                          = 0x55
	ATTR_MAX_NUM_PMKIDS                 = 0x56
	ATTR_DURATION                       = 0x57
	ATTR_COOKIE                         = 0x58
	ATTR_WIPHY_COVERAGE_CLASS           = 0x59
	ATTR_TX_RATES                       = 0x5a
	ATTR_FRAME_MATCH                    = 0x5b
	ATTR_ACK                            = 0x5c
	ATTR_PS_STATE                       = 0x5d
	ATTR_CQM                            = 0x5e
	ATTR_LOCAL_STATE_CHANGE             = 0x5f
	ATTR_AP_ISOLATE                     = 0x60
	ATTR_WIPHY_TX_POWER_SETTING         = 0x61
	ATTR_WIPHY_TX_POWER_LEVEL           = 0x62
	ATTR_TX_FRAME_TYPES                 = 0x63
	ATTR_RX_FRAME_TYPES                 = 0x64
	ATTR_FRAME_TYPE                     = 0x65
	ATTR_CONTROL_PORT_ETHERTYPE         = 0x66
	ATTR_CONTROL_PORT_NO_ENCRYPT        = 0x67
	ATTR_SUPPORT_IBSS_RSN               = 0x68
	ATTR_WIPHY_ANTENNA_TX               = 0x69
	ATTR_WIPHY_ANTENNA_RX               = 0x6a
	ATTR_MCAST_RATE                     = 0x6b
	ATTR_OFFCHANNEL_TX_OK               = 0x6c
	ATTR_BSS_HT_OPMODE                  = 0x6d
	ATTR_KEY_DEFAULT_TYPES              = 0x6e
	ATTR_MAX_REMAIN_ON_CHANNEL_DURATION = 0x6f
	ATTR_MESH_SETUP                     = 0x70
	ATTR_WIPHY_ANTENNA_AVAIL_TX         = 0x71
	ATTR_WIPHY_ANTENNA_AVAIL_RX         = 0x72
	ATTR_SUPPORT_MESH_AUTH              = 0x73
	ATTR_STA_PLINK_STATE                = 0x74
	ATTR_WOWLAN_TRIGGERS                = 0x75
	ATTR_WOWLAN_TRIGGERS_SUPPORTED      = 0x76
	ATTR_SCHED_SCAN_INTERVAL            = 0x77
	ATTR_INTERFACE_COMBINATIONS         = 0x78
	ATTR_SOFTWARE_IFTYPES               = 0x79
	__ATTR_AFTER_LAST                   = 0x7a
	ATTR_MAX                            = 0x79
	IFTYPE_UNSPECIFIED                  = 0
	IFTYPE_ADHOC                        = 0x1
	IFTYPE_STATION                      = 0x2
	IFTYPE_AP                           = 0x3
	IFTYPE_AP_VLAN                      = 0x4
	IFTYPE_WDS                          = 0x5
	IFTYPE_MONITOR                      = 0x6
	IFTYPE_MESH_POINT                   = 0x7
	IFTYPE_P2P_CLIENT                   = 0x8
	IFTYPE_P2P_GO                       = 0x9
	NUM_IFTYPES                         = 0xa
	IFTYPE_MAX                          = 0x9
	__STA_FLAG_INVALID                  = 0
	STA_FLAG_AUTHORIZED                 = 0x1
	STA_FLAG_SHORT_PREAMBLE             = 0x2
	STA_FLAG_WME                        = 0x3
	STA_FLAG_MFP                        = 0x4
	STA_FLAG_AUTHENTICATED              = 0x5
	__STA_FLAG_AFTER_LAST               = 0x6
	STA_FLAG_MAX                        = 0x5
	__RATE_INFO_INVALID                 = 0
	RATE_INFO_BITRATE                   = 0x1
	RATE_INFO_MCS                       = 0x2
	RATE_INFO_40_MHZ_WIDTH              = 0x3
	RATE_INFO_SHORT_GI                  = 0x4
	__RATE_INFO_AFTER_LAST              = 0x5
	RATE_INFO_MAX                       = 0x4
	__STA_BSS_PARAM_INVALID             = 0
	STA_BSS_PARAM_CTS_PROT              = 0x1
	STA_BSS_PARAM_SHORT_PREAMBLE        = 0x2
	STA_BSS_PARAM_SHORT_SLOT_TIME       = 0x3
	STA_BSS_PARAM_DTIM_PERIOD           = 0x4
	STA_BSS_PARAM_BEACON_INTERVAL       = 0x5
	__STA_BSS_PARAM_AFTER_LAST          = 0x6
	STA_BSS_PARAM_MAX                   = 0x5
	__STA_INFO_INVALID                  = 0
	STA_INFO_INACTIVE_TIME              = 0x1
	STA_INFO_RX_BYTES                   = 0x2
	STA_INFO_TX_BYTES                   = 0x3
	STA_INFO_LLID                       = 0x4
	STA_INFO_PLID                       = 0x5
	STA_INFO_PLINK_STATE                = 0x6
	STA_INFO_SIGNAL                     = 0x7
	STA_INFO_TX_BITRATE                 = 0x8
	STA_INFO_RX_PACKETS                 = 0x9
	STA_INFO_TX_PACKETS                 = 0xa
	STA_INFO_TX_RETRIES                 = 0xb
	STA_INFO_TX_FAILED                  = 0xc
	STA_INFO_SIGNAL_AVG                 = 0xd
	STA_INFO_RX_BITRATE                 = 0xe
	STA_INFO_BSS_PARAM                  = 0xf
	STA_INFO_CONNECTED_TIME             = 0x10
	__STA_INFO_AFTER_LAST               = 0x11
	STA_INFO_MAX                        = 0x10
	MPATH_FLAG_ACTIVE                   = 0x1
	MPATH_FLAG_RESOLVING                = 0x2
	MPATH_FLAG_SN_VALID                 = 0x4
	MPATH_FLAG_FIXED                    = 0x8
	MPATH_FLAG_RESOLVED                 = 0x10
	__MPATH_INFO_INVALID                = 0
	MPATH_INFO_FRAME_QLEN               = 0x1
	MPATH_INFO_SN                       = 0x2
	MPATH_INFO_METRIC                   = 0x3
	MPATH_INFO_EXPTIME                  = 0x4
	MPATH_INFO_FLAGS                    = 0x5
	MPATH_INFO_DISCOVERY_TIMEOUT        = 0x6
	MPATH_INFO_DISCOVERY_RETRIES        = 0x7
	__MPATH_INFO_AFTER_LAST             = 0x8
	MPATH_INFO_MAX                      = 0x7
	__BAND_ATTR_INVALID                 = 0
	BAND_ATTR_FREQS                     = 0x1
	BAND_ATTR_RATES                     = 0x2
	BAND_ATTR_HT_MCS_SET                = 0x3
	BAND_ATTR_HT_CAPA                   = 0x4
	BAND_ATTR_HT_AMPDU_FACTOR           = 0x5
	BAND_ATTR_HT_AMPDU_DENSITY          = 0x6
	__BAND_ATTR_AFTER_LAST              = 0x7
	BAND_ATTR_MAX                       = 0x6
	__FREQUENCY_ATTR_INVALID            = 0
	FREQUENCY_ATTR_FREQ                 = 0x1
	FREQUENCY_ATTR_DISABLED             = 0x2
	FREQUENCY_ATTR_PASSIVE_SCAN         = 0x3
	FREQUENCY_ATTR_NO_IBSS              = 0x4
	FREQUENCY_ATTR_RADAR                = 0x5
	FREQUENCY_ATTR_MAX_TX_POWER         = 0x6
	__FREQUENCY_ATTR_AFTER_LAST         = 0x7
	FREQUENCY_ATTR_MAX                  = 0x6
	__BITRATE_ATTR_INVALID              = 0
	BITRATE_ATTR_RATE                   = 0x1
	BITRATE_ATTR_2GHZ_SHORTPREAMBLE     = 0x2
	__BITRATE_ATTR_AFTER_LAST           = 0x3
	BITRATE_ATTR_MAX                    = 0x2
	REGDOM_SET_BY_CORE                  = 0
	REGDOM_SET_BY_USER                  = 0x1
	REGDOM_SET_BY_DRIVER                = 0x2
	REGDOM_SET_BY_COUNTRY_IE            = 0x3
	REGDOM_TYPE_COUNTRY                 = 0
	REGDOM_TYPE_WORLD                   = 0x1
	REGDOM_TYPE_CUSTOM_WORLD            = 0x2
	REGDOM_TYPE_INTERSECTION            = 0x3
	__REG_RULE_ATTR_INVALID             = 0
	ATTR_REG_RULE_FLAGS                 = 0x1
	ATTR_FREQ_RANGE_START               = 0x2
	ATTR_FREQ_RANGE_END                 = 0x3
	ATTR_FREQ_RANGE_MAX_BW              = 0x4
	ATTR_POWER_RULE_MAX_ANT_GAIN        = 0x5
	ATTR_POWER_RULE_MAX_EIRP            = 0x6
	__REG_RULE_ATTR_AFTER_LAST          = 0x7
	REG_RULE_ATTR_MAX                   = 0x6
	RRF_NO_OFDM                         = 0x1
	RRF_NO_CCK                          = 0x2
	RRF_NO_INDOOR                       = 0x4
	RRF_NO_OUTDOOR                      = 0x8
	RRF_DFS                             = 0x10
	RRF_PTP_ONLY                        = 0x20
	RRF_PTMP_ONLY                       = 0x40
	RRF_PASSIVE_SCAN                    = 0x80
	RRF_NO_IBSS                         = 0x100
	__SURVEY_INFO_INVALID               = 0
	SURVEY_INFO_FREQUENCY               = 0x1
	SURVEY_INFO_NOISE                   = 0x2
	SURVEY_INFO_IN_USE                  = 0x3
	SURVEY_INFO_CHANNEL_TIME            = 0x4
	SURVEY_INFO_CHANNEL_TIME_BUSY       = 0x5
	SURVEY_INFO_CHANNEL_TIME_EXT_BUSY   = 0x6
	SURVEY_INFO_CHANNEL_TIME_RX         = 0x7
	SURVEY_INFO_CHANNEL_TIME_TX         = 0x8
	__SURVEY_INFO_AFTER_LAST            = 0x9
	SURVEY_INFO_MAX                     = 0x8
	__MNTR_FLAG_INVALID                 = 0
	MNTR_FLAG_FCSFAIL                   = 0x1
	MNTR_FLAG_PLCPFAIL                  = 0x2
	MNTR_FLAG_CONTROL                   = 0x3
	MNTR_FLAG_OTHER_BSS                 = 0x4
	MNTR_FLAG_COOK_FRAMES               = 0x5
	__MNTR_FLAG_AFTER_LAST              = 0x6
	MNTR_FLAG_MAX                       = 0x5
	__MESHCONF_INVALID                  = 0
	MESHCONF_RETRY_TIMEOUT              = 0x1
	MESHCONF_CONFIRM_TIMEOUT            = 0x2
	MESHCONF_HOLDING_TIMEOUT            = 0x3
	MESHCONF_MAX_PEER_LINKS             = 0x4
	MESHCONF_MAX_RETRIES                = 0x5
	MESHCONF_TTL                        = 0x6
	MESHCONF_AUTO_OPEN_PLINKS           = 0x7
	MESHCONF_HWMP_MAX_PREQ_RETRIES      = 0x8
	MESHCONF_PATH_REFRESH_TIME          = 0x9
	MESHCONF_MIN_DISCOVERY_TIMEOUT      = 0xa
	MESHCONF_HWMP_ACTIVE_PATH_TIMEOUT   = 0xb
	MESHCONF_HWMP_PREQ_MIN_INTERVAL     = 0xc
	MESHCONF_HWMP_NET_DIAM_TRVS_TIME    = 0xd
	MESHCONF_HWMP_ROOTMODE              = 0xe
	MESHCONF_ELEMENT_TTL                = 0xf
	__MESHCONF_ATTR_AFTER_LAST          = 0x10
	MESHCONF_ATTR_MAX                   = 0xf
	__MESH_SETUP_INVALID                = 0
	MESH_SETUP_ENABLE_VENDOR_PATH_SEL   = 0x1
	MESH_SETUP_ENABLE_VENDOR_METRIC     = 0x2
	MESH_SETUP_IE                       = 0x3
	MESH_SETUP_USERSPACE_AUTH           = 0x4
	MESH_SETUP_USERSPACE_AMPE           = 0x5
	__MESH_SETUP_ATTR_AFTER_LAST        = 0x6
	MESH_SETUP_ATTR_MAX                 = 0x5
	__TXQ_ATTR_INVALID                  = 0
	TXQ_ATTR_QUEUE                      = 0x1
	TXQ_ATTR_TXOP                       = 0x2
	TXQ_ATTR_CWMIN                      = 0x3
	TXQ_ATTR_CWMAX                      = 0x4
	TXQ_ATTR_AIFS                       = 0x5
	__TXQ_ATTR_AFTER_LAST               = 0x6
	TXQ_ATTR_MAX                        = 0x5
	TXQ_Q_VO                            = 0
	TXQ_Q_VI                            = 0x1
	TXQ_Q_BE                            = 0x2
	TXQ_Q_BK                            = 0x3
	CHAN_NO_HT                          = 0
	CHAN_HT20                           = 0x1
	CHAN_HT40MINUS                      = 0x2
	CHAN_HT40PLUS                       = 0x3
	__BSS_INVALID                       = 0
	BSS_BSSID                           = 0x1
	BSS_FREQUENCY                       = 0x2
	BSS_TSF                             = 0x3
	BSS_BEACON_INTERVAL                 = 0x4
	BSS_CAPABILITY                      = 0x5
	BSS_INFORMATION_ELEMENTS            = 0x6
	BSS_SIGNAL_MBM                      = 0x7
	BSS_SIGNAL_UNSPEC                   = 0x8
	BSS_STATUS                          = 0x9
	BSS_SEEN_MS_AGO                     = 0xa
	BSS_BEACON_IES                      = 0xb
	__BSS_AFTER_LAST                    = 0xc
	BSS_MAX                             = 0xb
	BSS_STATUS_AUTHENTICATED            = 0
	BSS_STATUS_ASSOCIATED               = 0x1
	BSS_STATUS_IBSS_JOINED              = 0x2
	AUTHTYPE_OPEN_SYSTEM                = 0
	AUTHTYPE_SHARED_KEY                 = 0x1
	AUTHTYPE_FT                         = 0x2
	AUTHTYPE_NETWORK_EAP                = 0x3
	__AUTHTYPE_NUM                      = 0x4
	AUTHTYPE_MAX                        = 0x3
	AUTHTYPE_AUTOMATIC                  = 0x4
	KEYTYPE_GROUP                       = 0
	KEYTYPE_PAIRWISE                    = 0x1
	KEYTYPE_PEERKEY                     = 0x2
	NUM_KEYTYPES                        = 0x3
	MFP_NO                              = 0
	MFP_REQUIRED                        = 0x1
	WPA_VERSION_1                       = 0x1
	WPA_VERSION_2                       = 0x2
	__KEY_DEFAULT_TYPE_INVALID          = 0
	KEY_DEFAULT_TYPE_UNICAST            = 0x1
	KEY_DEFAULT_TYPE_MULTICAST          = 0x2
	NUM_KEY_DEFAULT_TYPES               = 0x3
	__KEY_INVALID                       = 0
	KEY_DATA                            = 0x1
	KEY_IDX                             = 0x2
	KEY_CIPHER                          = 0x3
	KEY_SEQ                             = 0x4
	KEY_DEFAULT                         = 0x5
	KEY_DEFAULT_MGMT                    = 0x6
	KEY_TYPE                            = 0x7
	KEY_DEFAULT_TYPES                   = 0x8
	__KEY_AFTER_LAST                    = 0x9
	KEY_MAX                             = 0x8
	__TXRATE_INVALID                    = 0
	TXRATE_LEGACY                       = 0x1
	__TXRATE_AFTER_LAST                 = 0x2
	TXRATE_MAX                          = 0x1
	BAND_2GHZ                           = 0
	BAND_5GHZ                           = 0x1
	PS_DISABLED                         = 0
	PS_ENABLED                          = 0x1
	__ATTR_CQM_INVALID                  = 0
	ATTR_CQM_RSSI_THOLD                 = 0x1
	ATTR_CQM_RSSI_HYST                  = 0x2
	ATTR_CQM_RSSI_THRESHOLD_EVENT       = 0x3
	ATTR_CQM_PKT_LOSS_EVENT             = 0x4
	__ATTR_CQM_AFTER_LAST               = 0x5
	ATTR_CQM_MAX                        = 0x4
	CQM_RSSI_THRESHOLD_EVENT_LOW        = 0
	CQM_RSSI_THRESHOLD_EVENT_HIGH       = 0x1
	TX_POWER_AUTOMATIC                  = 0
	TX_POWER_LIMITED                    = 0x1
	TX_POWER_FIXED                      = 0x2
	__WOWLAN_PKTPAT_INVALID             = 0
	WOWLAN_PKTPAT_MASK                  = 0x1
	WOWLAN_PKTPAT_PATTERN               = 0x2
	NUM_WOWLAN_PKTPAT                   = 0x3
	MAX_WOWLAN_PKTPAT                   = 0x2
	__WOWLAN_TRIG_INVALID               = 0
	WOWLAN_TRIG_ANY                     = 0x1
	WOWLAN_TRIG_DISCONNECT              = 0x2
	WOWLAN_TRIG_MAGIC_PKT               = 0x3
	WOWLAN_TRIG_PKT_PATTERN             = 0x4
	NUM_WOWLAN_TRIG                     = 0x5
	MAX_WOWLAN_TRIG                     = 0x4
	IFACE_LIMIT_UNSPEC                  = 0
	IFACE_LIMIT_MAX                     = 0x1
	IFACE_LIMIT_TYPES                   = 0x2
	NUM_IFACE_LIMIT                     = 0x3
	MAX_IFACE_LIMIT                     = 0x2
	IFACE_COMB_UNSPEC                   = 0
	IFACE_COMB_LIMITS                   = 0x1
	IFACE_COMB_MAXNUM                   = 0x2
	IFACE_COMB_STA_AP_BI_MATCH          = 0x3
	IFACE_COMB_NUM_CHANNELS             = 0x4
	NUM_IFACE_COMB                      = 0x5
	MAX_IFACE_COMB                      = 0x4
	PLINK_LISTEN                        = 0
	PLINK_OPN_SNT                       = 0x1
	PLINK_OPN_RCVD                      = 0x2
	PLINK_CNF_RCVD                      = 0x3
	PLINK_ESTAB                         = 0x4
	PLINK_HOLDING                       = 0x5
	PLINK_BLOCKED                       = 0x6
	NUM_PLINK_STATES                    = 0x7
	MAX_PLINK_STATES                    = 0x6
)

// Types

type nl80211_sta_flag_update struct {
	Mask uint32
	Set  uint32
}

type nl80211_wowlan_pattern_support struct {
	Max_patterns    uint32
	Min_pattern_len uint32
	Max_pattern_len uint32
}