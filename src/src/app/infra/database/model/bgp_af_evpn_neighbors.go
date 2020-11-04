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


CREATE TABLE `bgp_af_evpn_neighbors` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `remote_interface_id` int DEFAULT NULL,
  `remote_device_id` int DEFAULT NULL,
  `af_id` int DEFAULT NULL,
  `neighbor_afi` varchar(100) DEFAULT NULL,
  `neighbor_safi` varchar(100) DEFAULT NULL,
  `neighbor_ip_address` varchar(30) DEFAULT NULL,
  `remote_asn` bigint DEFAULT NULL,
  `peer_group_name` varchar(64) DEFAULT NULL,
  `encapsulation_type` varchar(30) DEFAULT NULL,
  `bfd_enable` tinyint(1) DEFAULT NULL,
  `bfd_tx` int DEFAULT NULL,
  `bfd_rx` int DEFAULT NULL,
  `bfd_multiplier` int DEFAULT NULL,
  `e_bgp_mhop` int DEFAULT NULL,
  `nexthop_unchanged` tinyint(1) DEFAULT NULL,
  `enable_peer_as_check` tinyint(1) DEFAULT NULL,
  `is_activated` varchar(30) DEFAULT NULL,
  `config_type` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `bgp_af_evpn_neighbors_fabric_id_fkey` (`fabric_id`),
  KEY `bgp_af_evpn_neighbors_device_id_fkey` (`device_id`),
  KEY `bgp_af_evpn_neighbors_remote_interface_id_fkey` (`remote_interface_id`),
  KEY `bgp_af_evpn_neighbors_remote_device_id_fkey` (`remote_device_id`),
  KEY `bgp_af_evpn_neighors_af_id_fkey` (`af_id`),
  CONSTRAINT `bgp_af_evpn_neighbors_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `bgp_af_evpn_neighbors_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE,
  CONSTRAINT `bgp_af_evpn_neighbors_remote_device_id_fkey` FOREIGN KEY (`remote_device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `bgp_af_evpn_neighbors_remote_interface_id_fkey` FOREIGN KEY (`remote_interface_id`) REFERENCES `phys_interfaces` (`id`) ON DELETE CASCADE,
  CONSTRAINT `bgp_af_evpn_neighors_af_id_fkey` FOREIGN KEY (`af_id`) REFERENCES `bgp_evpn_afs` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "neighbor_safi": "DKbwwDsAcImXpsHPtTVRTIsXC",    "remote_asn": 29,    "peer_group_name": "petecxjEtpTxZwngHYSjLWLRj",    "encapsulation_type": "lyImDefIlgZxRxRwwlqwphqJK",    "bfd_rx": 98,    "device_id": 16,    "remote_interface_id": 98,    "af_id": 75,    "neighbor_afi": "GOUKebckvOFpKPGhfOIVNfXmq",    "fabric_id": 72,    "neighbor_ip_address": "VWKLwXHxughlQISrndjZEVrIO",    "bfd_multiplier": 11,    "enable_peer_as_check": 59,    "config_type": "cOdcvOLaKnWHAyjjZHoXUCgDv",    "id": 50,    "remote_device_id": 44,    "bfd_enable": 17,    "bfd_tx": 50,    "e_bgp_mhop": 27,    "nexthop_unchanged": 8,    "is_activated": "CBUvuFbowrvHisyplVnDWjFAo"}



*/

// BgpAfEvpnNeighbors struct is a row record of the bgp_af_evpn_neighbors table in the myapp database
type BgpAfEvpnNeighbors struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 3] remote_interface_id                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RemoteInterfaceID sql.NullInt64 `gorm:"column:remote_interface_id;type:int;"`
	//[ 4] remote_device_id                               int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RemoteDeviceID sql.NullInt64 `gorm:"column:remote_device_id;type:int;"`
	//[ 5] af_id                                          int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AfID sql.NullInt64 `gorm:"column:af_id;type:int;"`
	//[ 6] neighbor_afi                                   varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	NeighborAfi sql.NullString `gorm:"column:neighbor_afi;type:varchar;size:100;"`
	//[ 7] neighbor_safi                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	NeighborSafi sql.NullString `gorm:"column:neighbor_safi;type:varchar;size:100;"`
	//[ 8] neighbor_ip_address                            varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	NeighborIPAddress sql.NullString `gorm:"column:neighbor_ip_address;type:varchar;size:30;"`
	//[ 9] remote_asn                                     bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	RemoteAsn sql.NullInt64 `gorm:"column:remote_asn;type:bigint;"`
	//[10] peer_group_name                                varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	PeerGroupName sql.NullString `gorm:"column:peer_group_name;type:varchar;size:64;"`
	//[11] encapsulation_type                             varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	EncapsulationType sql.NullString `gorm:"column:encapsulation_type;type:varchar;size:30;"`
	//[12] bfd_enable                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	BfdEnable sql.NullInt64 `gorm:"column:bfd_enable;type:tinyint;"`
	//[13] bfd_tx                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BfdTx sql.NullInt64 `gorm:"column:bfd_tx;type:int;"`
	//[14] bfd_rx                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BfdRx sql.NullInt64 `gorm:"column:bfd_rx;type:int;"`
	//[15] bfd_multiplier                                 int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BfdMultiplier sql.NullInt64 `gorm:"column:bfd_multiplier;type:int;"`
	//[16] e_bgp_mhop                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	EBgpMhop sql.NullInt64 `gorm:"column:e_bgp_mhop;type:int;"`
	//[17] nexthop_unchanged                              tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	NexthopUnchanged sql.NullInt64 `gorm:"column:nexthop_unchanged;type:tinyint;"`
	//[18] enable_peer_as_check                           tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	EnablePeerAsCheck sql.NullInt64 `gorm:"column:enable_peer_as_check;type:tinyint;"`
	//[19] is_activated                                   varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IsActivated sql.NullString `gorm:"column:is_activated;type:varchar;size:30;"`
	//[20] config_type                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType sql.NullString `gorm:"column:config_type;type:varchar;size:30;"`
}

var bgp_af_evpn_neighborsTableInfo = &TableInfo{
	Name: "bgp_af_evpn_neighbors",
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
			Name:               "device_id",
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
			GoFieldName:        "DeviceID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "device_id",
			ProtobufFieldName:  "device_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "remote_interface_id",
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
			GoFieldName:        "RemoteInterfaceID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "remote_interface_id",
			ProtobufFieldName:  "remote_interface_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "remote_device_id",
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
			GoFieldName:        "RemoteDeviceID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "remote_device_id",
			ProtobufFieldName:  "remote_device_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "af_id",
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
			GoFieldName:        "AfID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "af_id",
			ProtobufFieldName:  "af_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "neighbor_afi",
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
			GoFieldName:        "NeighborAfi",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "neighbor_afi",
			ProtobufFieldName:  "neighbor_afi",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "neighbor_safi",
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
			GoFieldName:        "NeighborSafi",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "neighbor_safi",
			ProtobufFieldName:  "neighbor_safi",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "neighbor_ip_address",
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
			GoFieldName:        "NeighborIPAddress",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "neighbor_ip_address",
			ProtobufFieldName:  "neighbor_ip_address",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "remote_asn",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "RemoteAsn",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "remote_asn",
			ProtobufFieldName:  "remote_asn",
			ProtobufType:       "int64",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "peer_group_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "PeerGroupName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "peer_group_name",
			ProtobufFieldName:  "peer_group_name",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "encapsulation_type",
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
			GoFieldName:        "EncapsulationType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "encapsulation_type",
			ProtobufFieldName:  "encapsulation_type",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "bfd_enable",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "BfdEnable",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bfd_enable",
			ProtobufFieldName:  "bfd_enable",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		{
			Index:              13,
			Name:               "bfd_tx",
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
			GoFieldName:        "BfdTx",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bfd_tx",
			ProtobufFieldName:  "bfd_tx",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		{
			Index:              14,
			Name:               "bfd_rx",
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
			GoFieldName:        "BfdRx",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bfd_rx",
			ProtobufFieldName:  "bfd_rx",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		{
			Index:              15,
			Name:               "bfd_multiplier",
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
			GoFieldName:        "BfdMultiplier",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bfd_multiplier",
			ProtobufFieldName:  "bfd_multiplier",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		{
			Index:              16,
			Name:               "e_bgp_mhop",
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
			GoFieldName:        "EBgpMhop",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "e_bgp_mhop",
			ProtobufFieldName:  "e_bgp_mhop",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		{
			Index:              17,
			Name:               "nexthop_unchanged",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "NexthopUnchanged",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "nexthop_unchanged",
			ProtobufFieldName:  "nexthop_unchanged",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		{
			Index:              18,
			Name:               "enable_peer_as_check",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "EnablePeerAsCheck",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "enable_peer_as_check",
			ProtobufFieldName:  "enable_peer_as_check",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		{
			Index:              19,
			Name:               "is_activated",
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
			GoFieldName:        "IsActivated",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "is_activated",
			ProtobufFieldName:  "is_activated",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		{
			Index:              20,
			Name:               "config_type",
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
			GoFieldName:        "ConfigType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "config_type",
			ProtobufFieldName:  "config_type",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BgpAfEvpnNeighbors) TableName() string {
	return "bgp_af_evpn_neighbors"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BgpAfEvpnNeighbors) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BgpAfEvpnNeighbors) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BgpAfEvpnNeighbors) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BgpAfEvpnNeighbors) TableInfo() *TableInfo {
	return bgp_af_evpn_neighborsTableInfo
}
