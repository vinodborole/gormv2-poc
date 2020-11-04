package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `fabric_properties` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `p2_p_link_range` varchar(100) DEFAULT NULL,
  `p2_p_ip_type` varchar(100) DEFAULT NULL,
  `loop_back_ip_range` varchar(100) DEFAULT NULL,
  `loop_back_port_number` varchar(100) DEFAULT NULL,
  `super_spine_asn_block` varchar(100) DEFAULT NULL,
  `spine_asn_block` varchar(100) DEFAULT NULL,
  `leaf_asn_block` varchar(100) DEFAULT NULL,
  `border_leaf_asn_block` varchar(100) DEFAULT NULL,
  `vtep_loop_back_port_number` varchar(100) DEFAULT NULL,
  `any_cast_mac` varchar(100) DEFAULT NULL,
  `ip_v6_any_cast_mac` varchar(100) DEFAULT NULL,
  `configure_overlay_gateway` varchar(10) DEFAULT NULL,
  `vni_auto_map` varchar(10) DEFAULT NULL,
  `bfd_enable` varchar(10) DEFAULT NULL,
  `bfd_tx` varchar(10) DEFAULT NULL,
  `bfd_rx` varchar(10) DEFAULT NULL,
  `bfd_multiplier` varchar(10) DEFAULT NULL,
  `bgp_multi_hop` varchar(10) DEFAULT NULL,
  `max_paths` varchar(10) DEFAULT NULL,
  `allow_as_in` varchar(10) DEFAULT NULL,
  `mtu` varchar(10) DEFAULT NULL,
  `ip_mtu` varchar(10) DEFAULT NULL,
  `leaf_peer_group` varchar(30) DEFAULT NULL,
  `spine_peer_group` varchar(30) DEFAULT NULL,
  `super_spine_peer_group` varchar(30) DEFAULT NULL,
  `arp_aging_timeout` varchar(10) DEFAULT NULL,
  `mac_aging_timeout` varchar(10) DEFAULT NULL,
  `mac_aging_conversational_timeout` varchar(10) DEFAULT NULL,
  `mac_move_limit` varchar(10) DEFAULT NULL,
  `duplicate_mac_timer` varchar(10) DEFAULT NULL,
  `duplicate_max_timer_max_count` varchar(10) DEFAULT NULL,
  `mct_link_ip_range` varchar(100) DEFAULT NULL,
  `control_vlan` varchar(10) DEFAULT NULL,
  `control_ve` varchar(10) DEFAULT NULL,
  `mct_port_channel` varchar(100) DEFAULT NULL,
  `rack_l3_back_up_ip_range` varchar(100) DEFAULT NULL,
  `rack_asn_block` varchar(100) DEFAULT NULL,
  `rack_underlay_ebgp_peer_group` varchar(30) DEFAULT NULL,
  `rack_overlay_ebgp_peer_group` varchar(30) DEFAULT NULL,
  `rack_l3_back_up_port` varchar(100) DEFAULT NULL,
  `rack_mct_ports` varchar(100) DEFAULT NULL,
  `rack_ld_l3_back_up_port` varchar(100) DEFAULT NULL,
  `rack_ld_mct_ports` varchar(100) DEFAULT NULL,
  `max_number_of_racks_in_fabric` varchar(10) DEFAULT NULL,
  `number_of_devices_in_rack` varchar(10) DEFAULT NULL,
  `backup_routing_enable` varchar(10) DEFAULT NULL,
  `backup_routing_ipv4_range` varchar(100) DEFAULT NULL,
  `backup_routing_ipv6_range` varchar(100) DEFAULT NULL,
  `lacp_timeout` varchar(10) DEFAULT NULL,
  `optimized_replication_enable` varchar(10) DEFAULT 'No',
  `mdtgroup_range` varchar(100) DEFAULT '239.0.0.0/8',
  `default_mdtgroup` varchar(100) DEFAULT '239.1.1.1',
  PRIMARY KEY (`id`),
  KEY `fabric_properties_fabric_id_fkey` (`fabric_id`),
  CONSTRAINT `fabric_properties_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "max_paths": "ROTHssbYpnCmtncFHMhpsDsta",    "super_spine_asn_block": "KVtELWrnqAlwAvTRivobONJhL",    "border_leaf_asn_block": "XtPFwgDhaASNpYadakqZFsZcD",    "ip_v_6_any_cast_mac": "OPfCTlvPRDOOZKOYPgGaiCEpw",    "bgp_multi_hop": "nJsnKlEghqdXxTEoPPmKnxvDF",    "mac_aging_conversational_timeout": "VSTAIraCDwbKEicsOBaDEHWRv",    "duplicate_max_timer_max_count": "ushEPXytUFITaseVgrTemwrUQ",    "rack_underlay_ebgp_peer_group": "vDWhJiDTgyrfmlwlmJJUWicRd",    "rack_overlay_ebgp_peer_group": "KUejDdgTPbHtfmwqNMUToOUuo",    "bfd_multiplier": "lVcQVoYZyXaUicXKePegqYKlr",    "duplicate_mac_timer": "EtwDrQZqHyrgEmoxMbpHcNMrY",    "number_of_devices_in_rack": "gkdaxQdpNbDMujOcRLmUVQedG",    "configure_overlay_gateway": "OTkyvSXWdSiWJFihSyKvgLtPK",    "arp_aging_timeout": "yaSvynnCAXwKJxYsRRpLEADQo",    "mac_aging_timeout": "pvQSWwjacynovIKQOhuqHvJyn",    "fabric_id": 79,    "p_2_p_link_range": "HlnUFDrrpBXdEDOyLesVKkPaf",    "mct_link_ip_range": "xEylOJYjepQmsdeqDCkuNNsSv",    "control_vlan": "DmtoVvAIyANgKefyiEmlfMmIY",    "lacp_timeout": "aLrAiNGJTWpfDXtMCjQxNLgPo",    "vni_auto_map": "NpQPXWArpfRNgNrDqCTCOUmto",    "ip_mtu": "LhNVuJAexlyWjUSYZIrmRJYNn",    "spine_peer_group": "SVSuruGqtZjOFMAkMTIglfAcD",    "backup_routing_ipv_4_range": "MgWGhfRwxopogQGBEopUvFpdL",    "bfd_enable": "gGrYjXZumLIQlAfcoORTrJCgO",    "leaf_peer_group": "gxVpVAphohQPWPJOKRTAYhtdc",    "backup_routing_ipv_6_range": "wgMHcvxsSvqIFSvlMIJvSUESX",    "optimized_replication_enable": "XQbyMIbAyhLofSiGNNNCMvcDH",    "allow_as_in": "eOsOVDGRqetxIhBHeBHXxChft",    "mac_move_limit": "hdjDyQjElxGboYaTGnETkPlBg",    "mct_port_channel": "BlZTZQSNMYCvjwXFILBVujnQu",    "rack_ld_l_3_back_up_port": "IYucMdiHmWVYMxTtPJvHsMkpR",    "spine_asn_block": "KImgrAEZwoENhcqRBiOgLcNYv",    "loop_back_ip_range": "NopxoKDTvUTyUqDuTbsTRWkos",    "bfd_tx": "ijTyLWnYXaAFipZowEBieCrtU",    "bfd_rx": "ZIXxYlOrqfOBTwFdKsCHXRbwg",    "rack_ld_mct_ports": "qGGaOmtlxQZNXKRAcHFeQZcLw",    "id": 78,    "any_cast_mac": "vXJpkaZNZiMRYPkNZLXVtfAHa",    "super_spine_peer_group": "pHhcCULIJXPdACfwMnvkhFqxV",    "p_2_p_ip_type": "LFaaqnSsoMYjXXenqWdJlqYtM",    "leaf_asn_block": "xMLmnyiLrogxJolLvmPZvkDkx",    "mtu": "MLOYtFuiuIyYfpmTPwsjqENBG",    "default_mdtgroup": "BkmbVkteaQUCUKREMoIxqkGDN",    "vtep_loop_back_port_number": "TjwifYFPMhukXloLULesWDaYg",    "rack_asn_block": "vMVJrvYGChiVEqvdrIntgkDyg",    "rack_mct_ports": "wUSrYEIWfjfSsOhFVquaRdjML",    "rack_l_3_back_up_port": "oONOGRfxdUTQPhxhJdEfYTWIo",    "backup_routing_enable": "oGOgphvaFOsJSHPrTEEVUfmUQ",    "mdtgroup_range": "TbOBgNFJoYhpQePMORjZImRiQ",    "loop_back_port_number": "CVTWorcFtWrfyTuaZqGZJqGpq",    "rack_l_3_back_up_ip_range": "HCDGfMXMUuLvlflwNAdSMZOFF",    "control_ve": "qKsYdRdAcskDCskVNsixwchmR",    "max_number_of_racks_in_fabric": "RjZSDPobhofDKOGUminewrXYA"}



*/

// FabricProperties struct is a row record of the fabric_properties table in the myapp database
type FabricProperties struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] p2_p_link_range                                varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	P2PLinkRange sql.NullString `gorm:"column:p2_p_link_range;type:varchar;size:100;"`
	//[ 3] p2_p_ip_type                                   varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	P2PIPType sql.NullString `gorm:"column:p2_p_ip_type;type:varchar;size:100;"`
	//[ 4] loop_back_ip_range                             varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	LoopBackIPRange sql.NullString `gorm:"column:loop_back_ip_range;type:varchar;size:100;"`
	//[ 5] loop_back_port_number                          varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	LoopBackPortNumber sql.NullString `gorm:"column:loop_back_port_number;type:varchar;size:100;"`
	//[ 6] super_spine_asn_block                          varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	SuperSpineAsnBlock sql.NullString `gorm:"column:super_spine_asn_block;type:varchar;size:100;"`
	//[ 7] spine_asn_block                                varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	SpineAsnBlock sql.NullString `gorm:"column:spine_asn_block;type:varchar;size:100;"`
	//[ 8] leaf_asn_block                                 varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	LeafAsnBlock sql.NullString `gorm:"column:leaf_asn_block;type:varchar;size:100;"`
	//[ 9] border_leaf_asn_block                          varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	BorderLeafAsnBlock sql.NullString `gorm:"column:border_leaf_asn_block;type:varchar;size:100;"`
	//[10] vtep_loop_back_port_number                     varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	VtepLoopBackPortNumber sql.NullString `gorm:"column:vtep_loop_back_port_number;type:varchar;size:100;"`
	//[11] any_cast_mac                                   varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	AnyCastMac sql.NullString `gorm:"column:any_cast_mac;type:varchar;size:100;"`
	//[12] ip_v6_any_cast_mac                             varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	IPV6AnyCastMac sql.NullString `gorm:"column:ip_v6_any_cast_mac;type:varchar;size:100;"`
	//[13] configure_overlay_gateway                      varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	ConfigureOverlayGateway sql.NullString `gorm:"column:configure_overlay_gateway;type:varchar;size:10;"`
	//[14] vni_auto_map                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	VniAutoMap sql.NullString `gorm:"column:vni_auto_map;type:varchar;size:10;"`
	//[15] bfd_enable                                     varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	BfdEnable sql.NullString `gorm:"column:bfd_enable;type:varchar;size:10;"`
	//[16] bfd_tx                                         varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	BfdTx sql.NullString `gorm:"column:bfd_tx;type:varchar;size:10;"`
	//[17] bfd_rx                                         varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	BfdRx sql.NullString `gorm:"column:bfd_rx;type:varchar;size:10;"`
	//[18] bfd_multiplier                                 varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	BfdMultiplier sql.NullString `gorm:"column:bfd_multiplier;type:varchar;size:10;"`
	//[19] bgp_multi_hop                                  varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	BgpMultiHop sql.NullString `gorm:"column:bgp_multi_hop;type:varchar;size:10;"`
	//[20] max_paths                                      varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	MaxPaths sql.NullString `gorm:"column:max_paths;type:varchar;size:10;"`
	//[21] allow_as_in                                    varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	AllowAsIn sql.NullString `gorm:"column:allow_as_in;type:varchar;size:10;"`
	//[22] mtu                                            varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	Mtu sql.NullString `gorm:"column:mtu;type:varchar;size:10;"`
	//[23] ip_mtu                                         varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	IPMtu sql.NullString `gorm:"column:ip_mtu;type:varchar;size:10;"`
	//[24] leaf_peer_group                                varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	LeafPeerGroup sql.NullString `gorm:"column:leaf_peer_group;type:varchar;size:30;"`
	//[25] spine_peer_group                               varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	SpinePeerGroup sql.NullString `gorm:"column:spine_peer_group;type:varchar;size:30;"`
	//[26] super_spine_peer_group                         varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	SuperSpinePeerGroup sql.NullString `gorm:"column:super_spine_peer_group;type:varchar;size:30;"`
	//[27] arp_aging_timeout                              varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	ArpAgingTimeout sql.NullString `gorm:"column:arp_aging_timeout;type:varchar;size:10;"`
	//[28] mac_aging_timeout                              varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	MacAgingTimeout sql.NullString `gorm:"column:mac_aging_timeout;type:varchar;size:10;"`
	//[29] mac_aging_conversational_timeout               varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	MacAgingConversationalTimeout sql.NullString `gorm:"column:mac_aging_conversational_timeout;type:varchar;size:10;"`
	//[30] mac_move_limit                                 varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	MacMoveLimit sql.NullString `gorm:"column:mac_move_limit;type:varchar;size:10;"`
	//[31] duplicate_mac_timer                            varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	DuplicateMacTimer sql.NullString `gorm:"column:duplicate_mac_timer;type:varchar;size:10;"`
	//[32] duplicate_max_timer_max_count                  varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	DuplicateMaxTimerMaxCount sql.NullString `gorm:"column:duplicate_max_timer_max_count;type:varchar;size:10;"`
	//[33] mct_link_ip_range                              varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	MctLinkIPRange sql.NullString `gorm:"column:mct_link_ip_range;type:varchar;size:100;"`
	//[34] control_vlan                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	ControlVlan sql.NullString `gorm:"column:control_vlan;type:varchar;size:10;"`
	//[35] control_ve                                     varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	ControlVe sql.NullString `gorm:"column:control_ve;type:varchar;size:10;"`
	//[36] mct_port_channel                               varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	MctPortChannel sql.NullString `gorm:"column:mct_port_channel;type:varchar;size:100;"`
	//[37] rack_l3_back_up_ip_range                       varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	RackL3BackUpIPRange sql.NullString `gorm:"column:rack_l3_back_up_ip_range;type:varchar;size:100;"`
	//[38] rack_asn_block                                 varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	RackAsnBlock sql.NullString `gorm:"column:rack_asn_block;type:varchar;size:100;"`
	//[39] rack_underlay_ebgp_peer_group                  varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	RackUnderlayEbgpPeerGroup sql.NullString `gorm:"column:rack_underlay_ebgp_peer_group;type:varchar;size:30;"`
	//[40] rack_overlay_ebgp_peer_group                   varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	RackOverlayEbgpPeerGroup sql.NullString `gorm:"column:rack_overlay_ebgp_peer_group;type:varchar;size:30;"`
	//[41] rack_l3_back_up_port                           varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	RackL3BackUpPort sql.NullString `gorm:"column:rack_l3_back_up_port;type:varchar;size:100;"`
	//[42] rack_mct_ports                                 varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	RackMctPorts sql.NullString `gorm:"column:rack_mct_ports;type:varchar;size:100;"`
	//[43] rack_ld_l3_back_up_port                        varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	RackLdL3BackUpPort sql.NullString `gorm:"column:rack_ld_l3_back_up_port;type:varchar;size:100;"`
	//[44] rack_ld_mct_ports                              varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	RackLdMctPorts sql.NullString `gorm:"column:rack_ld_mct_ports;type:varchar;size:100;"`
	//[45] max_number_of_racks_in_fabric                  varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	MaxNumberOfRacksInFabric sql.NullString `gorm:"column:max_number_of_racks_in_fabric;type:varchar;size:10;"`
	//[46] number_of_devices_in_rack                      varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	NumberOfDevicesInRack sql.NullString `gorm:"column:number_of_devices_in_rack;type:varchar;size:10;"`
	//[47] backup_routing_enable                          varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	BackupRoutingEnable sql.NullString `gorm:"column:backup_routing_enable;type:varchar;size:10;"`
	//[48] backup_routing_ipv4_range                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	BackupRoutingIpv4Range sql.NullString `gorm:"column:backup_routing_ipv4_range;type:varchar;size:100;"`
	//[49] backup_routing_ipv6_range                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	BackupRoutingIpv6Range sql.NullString `gorm:"column:backup_routing_ipv6_range;type:varchar;size:100;"`
	//[50] lacp_timeout                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	LacpTimeout sql.NullString `gorm:"column:lacp_timeout;type:varchar;size:10;"`
	//[51] optimized_replication_enable                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [No]
	OptimizedReplicationEnable sql.NullString `gorm:"column:optimized_replication_enable;type:varchar;size:10;default:No;"`
	//[52] mdtgroup_range                                 varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [239.0.0.0/8]
	MdtgroupRange sql.NullString `gorm:"column:mdtgroup_range;type:varchar;size:100;default:239.0.0.0/8;"`
	//[53] default_mdtgroup                               varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [239.1.1.1]
	DefaultMdtgroup sql.NullString `gorm:"column:default_mdtgroup;type:varchar;size:100;default:239.1.1.1;"`
}

var fabric_propertiesTableInfo = &TableInfo{
	Name: "fabric_properties",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "fabric_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FabricID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "fabric_id",
			ProtobufFieldName:  "fabric_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "p2_p_link_range",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "P2PLinkRange",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "p_2_p_link_range",
			ProtobufFieldName:  "p_2_p_link_range",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "p2_p_ip_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "P2PIPType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "p_2_p_ip_type",
			ProtobufFieldName:  "p_2_p_ip_type",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "loop_back_ip_range",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "LoopBackIPRange",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "loop_back_ip_range",
			ProtobufFieldName:  "loop_back_ip_range",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "loop_back_port_number",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "LoopBackPortNumber",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "loop_back_port_number",
			ProtobufFieldName:  "loop_back_port_number",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "super_spine_asn_block",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "SuperSpineAsnBlock",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "super_spine_asn_block",
			ProtobufFieldName:  "super_spine_asn_block",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "spine_asn_block",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "SpineAsnBlock",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "spine_asn_block",
			ProtobufFieldName:  "spine_asn_block",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "leaf_asn_block",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "LeafAsnBlock",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "leaf_asn_block",
			ProtobufFieldName:  "leaf_asn_block",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "border_leaf_asn_block",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "BorderLeafAsnBlock",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "border_leaf_asn_block",
			ProtobufFieldName:  "border_leaf_asn_block",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "vtep_loop_back_port_number",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "VtepLoopBackPortNumber",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "vtep_loop_back_port_number",
			ProtobufFieldName:  "vtep_loop_back_port_number",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "any_cast_mac",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "AnyCastMac",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "any_cast_mac",
			ProtobufFieldName:  "any_cast_mac",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "ip_v6_any_cast_mac",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "IPV6AnyCastMac",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ip_v_6_any_cast_mac",
			ProtobufFieldName:  "ip_v_6_any_cast_mac",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		{
			Index:              13,
			Name:               "configure_overlay_gateway",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "ConfigureOverlayGateway",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "configure_overlay_gateway",
			ProtobufFieldName:  "configure_overlay_gateway",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		{
			Index:              14,
			Name:               "vni_auto_map",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "VniAutoMap",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "vni_auto_map",
			ProtobufFieldName:  "vni_auto_map",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		{
			Index:              15,
			Name:               "bfd_enable",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "BfdEnable",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "bfd_enable",
			ProtobufFieldName:  "bfd_enable",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		{
			Index:              16,
			Name:               "bfd_tx",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "BfdTx",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "bfd_tx",
			ProtobufFieldName:  "bfd_tx",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		{
			Index:              17,
			Name:               "bfd_rx",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "BfdRx",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "bfd_rx",
			ProtobufFieldName:  "bfd_rx",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		{
			Index:              18,
			Name:               "bfd_multiplier",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "BfdMultiplier",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "bfd_multiplier",
			ProtobufFieldName:  "bfd_multiplier",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		{
			Index:              19,
			Name:               "bgp_multi_hop",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "BgpMultiHop",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "bgp_multi_hop",
			ProtobufFieldName:  "bgp_multi_hop",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		{
			Index:              20,
			Name:               "max_paths",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "MaxPaths",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "max_paths",
			ProtobufFieldName:  "max_paths",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},

		{
			Index:              21,
			Name:               "allow_as_in",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "AllowAsIn",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "allow_as_in",
			ProtobufFieldName:  "allow_as_in",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		{
			Index:              22,
			Name:               "mtu",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "Mtu",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mtu",
			ProtobufFieldName:  "mtu",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		{
			Index:              23,
			Name:               "ip_mtu",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "IPMtu",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ip_mtu",
			ProtobufFieldName:  "ip_mtu",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		{
			Index:              24,
			Name:               "leaf_peer_group",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "LeafPeerGroup",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "leaf_peer_group",
			ProtobufFieldName:  "leaf_peer_group",
			ProtobufType:       "string",
			ProtobufPos:        25,
		},

		{
			Index:              25,
			Name:               "spine_peer_group",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "SpinePeerGroup",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "spine_peer_group",
			ProtobufFieldName:  "spine_peer_group",
			ProtobufType:       "string",
			ProtobufPos:        26,
		},

		{
			Index:              26,
			Name:               "super_spine_peer_group",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "SuperSpinePeerGroup",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "super_spine_peer_group",
			ProtobufFieldName:  "super_spine_peer_group",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},

		{
			Index:              27,
			Name:               "arp_aging_timeout",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "ArpAgingTimeout",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "arp_aging_timeout",
			ProtobufFieldName:  "arp_aging_timeout",
			ProtobufType:       "string",
			ProtobufPos:        28,
		},

		{
			Index:              28,
			Name:               "mac_aging_timeout",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "MacAgingTimeout",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mac_aging_timeout",
			ProtobufFieldName:  "mac_aging_timeout",
			ProtobufType:       "string",
			ProtobufPos:        29,
		},

		{
			Index:              29,
			Name:               "mac_aging_conversational_timeout",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "MacAgingConversationalTimeout",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mac_aging_conversational_timeout",
			ProtobufFieldName:  "mac_aging_conversational_timeout",
			ProtobufType:       "string",
			ProtobufPos:        30,
		},

		{
			Index:              30,
			Name:               "mac_move_limit",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "MacMoveLimit",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mac_move_limit",
			ProtobufFieldName:  "mac_move_limit",
			ProtobufType:       "string",
			ProtobufPos:        31,
		},

		{
			Index:              31,
			Name:               "duplicate_mac_timer",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "DuplicateMacTimer",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "duplicate_mac_timer",
			ProtobufFieldName:  "duplicate_mac_timer",
			ProtobufType:       "string",
			ProtobufPos:        32,
		},

		{
			Index:              32,
			Name:               "duplicate_max_timer_max_count",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "DuplicateMaxTimerMaxCount",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "duplicate_max_timer_max_count",
			ProtobufFieldName:  "duplicate_max_timer_max_count",
			ProtobufType:       "string",
			ProtobufPos:        33,
		},

		{
			Index:              33,
			Name:               "mct_link_ip_range",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "MctLinkIPRange",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mct_link_ip_range",
			ProtobufFieldName:  "mct_link_ip_range",
			ProtobufType:       "string",
			ProtobufPos:        34,
		},

		{
			Index:              34,
			Name:               "control_vlan",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "ControlVlan",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "control_vlan",
			ProtobufFieldName:  "control_vlan",
			ProtobufType:       "string",
			ProtobufPos:        35,
		},

		{
			Index:              35,
			Name:               "control_ve",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "ControlVe",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "control_ve",
			ProtobufFieldName:  "control_ve",
			ProtobufType:       "string",
			ProtobufPos:        36,
		},

		{
			Index:              36,
			Name:               "mct_port_channel",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "MctPortChannel",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mct_port_channel",
			ProtobufFieldName:  "mct_port_channel",
			ProtobufType:       "string",
			ProtobufPos:        37,
		},

		{
			Index:              37,
			Name:               "rack_l3_back_up_ip_range",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "RackL3BackUpIPRange",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "rack_l_3_back_up_ip_range",
			ProtobufFieldName:  "rack_l_3_back_up_ip_range",
			ProtobufType:       "string",
			ProtobufPos:        38,
		},

		{
			Index:              38,
			Name:               "rack_asn_block",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "RackAsnBlock",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "rack_asn_block",
			ProtobufFieldName:  "rack_asn_block",
			ProtobufType:       "string",
			ProtobufPos:        39,
		},

		{
			Index:              39,
			Name:               "rack_underlay_ebgp_peer_group",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "RackUnderlayEbgpPeerGroup",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "rack_underlay_ebgp_peer_group",
			ProtobufFieldName:  "rack_underlay_ebgp_peer_group",
			ProtobufType:       "string",
			ProtobufPos:        40,
		},

		{
			Index:              40,
			Name:               "rack_overlay_ebgp_peer_group",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "RackOverlayEbgpPeerGroup",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "rack_overlay_ebgp_peer_group",
			ProtobufFieldName:  "rack_overlay_ebgp_peer_group",
			ProtobufType:       "string",
			ProtobufPos:        41,
		},

		{
			Index:              41,
			Name:               "rack_l3_back_up_port",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "RackL3BackUpPort",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "rack_l_3_back_up_port",
			ProtobufFieldName:  "rack_l_3_back_up_port",
			ProtobufType:       "string",
			ProtobufPos:        42,
		},

		{
			Index:              42,
			Name:               "rack_mct_ports",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "RackMctPorts",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "rack_mct_ports",
			ProtobufFieldName:  "rack_mct_ports",
			ProtobufType:       "string",
			ProtobufPos:        43,
		},

		{
			Index:              43,
			Name:               "rack_ld_l3_back_up_port",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "RackLdL3BackUpPort",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "rack_ld_l_3_back_up_port",
			ProtobufFieldName:  "rack_ld_l_3_back_up_port",
			ProtobufType:       "string",
			ProtobufPos:        44,
		},

		{
			Index:              44,
			Name:               "rack_ld_mct_ports",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "RackLdMctPorts",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "rack_ld_mct_ports",
			ProtobufFieldName:  "rack_ld_mct_ports",
			ProtobufType:       "string",
			ProtobufPos:        45,
		},

		{
			Index:              45,
			Name:               "max_number_of_racks_in_fabric",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "MaxNumberOfRacksInFabric",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "max_number_of_racks_in_fabric",
			ProtobufFieldName:  "max_number_of_racks_in_fabric",
			ProtobufType:       "string",
			ProtobufPos:        46,
		},

		{
			Index:              46,
			Name:               "number_of_devices_in_rack",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "NumberOfDevicesInRack",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "number_of_devices_in_rack",
			ProtobufFieldName:  "number_of_devices_in_rack",
			ProtobufType:       "string",
			ProtobufPos:        47,
		},

		{
			Index:              47,
			Name:               "backup_routing_enable",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "BackupRoutingEnable",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "backup_routing_enable",
			ProtobufFieldName:  "backup_routing_enable",
			ProtobufType:       "string",
			ProtobufPos:        48,
		},

		{
			Index:              48,
			Name:               "backup_routing_ipv4_range",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "BackupRoutingIpv4Range",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "backup_routing_ipv_4_range",
			ProtobufFieldName:  "backup_routing_ipv_4_range",
			ProtobufType:       "string",
			ProtobufPos:        49,
		},

		{
			Index:              49,
			Name:               "backup_routing_ipv6_range",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "BackupRoutingIpv6Range",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "backup_routing_ipv_6_range",
			ProtobufFieldName:  "backup_routing_ipv_6_range",
			ProtobufType:       "string",
			ProtobufPos:        50,
		},

		{
			Index:              50,
			Name:               "lacp_timeout",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "LacpTimeout",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "lacp_timeout",
			ProtobufFieldName:  "lacp_timeout",
			ProtobufType:       "string",
			ProtobufPos:        51,
		},

		{
			Index:              51,
			Name:               "optimized_replication_enable",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "OptimizedReplicationEnable",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "optimized_replication_enable",
			ProtobufFieldName:  "optimized_replication_enable",
			ProtobufType:       "string",
			ProtobufPos:        52,
		},

		{
			Index:              52,
			Name:               "mdtgroup_range",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "MdtgroupRange",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mdtgroup_range",
			ProtobufFieldName:  "mdtgroup_range",
			ProtobufType:       "string",
			ProtobufPos:        53,
		},

		{
			Index:              53,
			Name:               "default_mdtgroup",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "DefaultMdtgroup",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "default_mdtgroup",
			ProtobufFieldName:  "default_mdtgroup",
			ProtobufType:       "string",
			ProtobufPos:        54,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FabricProperties) TableName() string {
	return "fabric_properties"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FabricProperties) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FabricProperties) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FabricProperties) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FabricProperties) TableInfo() *TableInfo {
	return fabric_propertiesTableInfo
}
