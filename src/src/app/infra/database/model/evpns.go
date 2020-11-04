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


CREATE TABLE `evpns` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `fabric_name` varchar(30) NOT NULL,
  `device_id` int DEFAULT NULL,
  `device_ip` varchar(30) NOT NULL,
  `sw_evpn_name` varchar(30) NOT NULL,
  `mac_aging_timeout` varchar(30) NOT NULL,
  `mac_move_limit` varchar(30) NOT NULL,
  `mac_aging_conversational_timeout` varchar(30) NOT NULL,
  `arp_agingtimeout` varchar(30) NOT NULL,
  `duplicate_mac_timer` varchar(30) NOT NULL,
  `duplicate_mac_timer_max_count` varchar(30) NOT NULL,
  `config_type` varchar(30) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `evpns_fabric_id_fkey` (`fabric_id`),
  KEY `evpns_device_id_fkey` (`device_id`),
  CONSTRAINT `evpns_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `evpns_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "device_id": 87,    "mac_move_limit": "sSOREqSCwnNRmjYQhcCJeVTAs",    "mac_aging_conversational_timeout": "JrsZKDHYTFmfDtEpdtyurdiRV",    "arp_agingtimeout": "ebTSEqWtnjVOuvIxWijrlLpAX",    "config_type": "axHkLXROrTJtErWaLWQqTesuH",    "duplicate_mac_timer": "yNYMnjeouiZtuSqwneqKLfgCl",    "duplicate_mac_timer_max_count": "tWWljPalUKADLbVjiQwEFDgbb",    "id": 30,    "fabric_id": 54,    "fabric_name": "maPPryPhkJJXSjHVyFUYKCCAu",    "device_ip": "ghOlPFkrNjiLWqLWDlvofCTvP",    "sw_evpn_name": "lSZRxQFYxfSIqCfuohmsOQkZy",    "mac_aging_timeout": "sJWXKFtQndLFEnduGaXilwrjo"}



*/

// Evpns struct is a row record of the evpns table in the myapp database
type Evpns struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] fabric_name                                    varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	FabricName string `gorm:"column:fabric_name;type:varchar;size:30;"`
	//[ 3] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 4] device_ip                                      varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceIP string `gorm:"column:device_ip;type:varchar;size:30;"`
	//[ 5] sw_evpn_name                                   varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	SwEvpnName string `gorm:"column:sw_evpn_name;type:varchar;size:30;"`
	//[ 6] mac_aging_timeout                              varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	MacAgingTimeout string `gorm:"column:mac_aging_timeout;type:varchar;size:30;"`
	//[ 7] mac_move_limit                                 varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	MacMoveLimit string `gorm:"column:mac_move_limit;type:varchar;size:30;"`
	//[ 8] mac_aging_conversational_timeout               varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	MacAgingConversationalTimeout string `gorm:"column:mac_aging_conversational_timeout;type:varchar;size:30;"`
	//[ 9] arp_agingtimeout                               varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ArpAgingtimeout string `gorm:"column:arp_agingtimeout;type:varchar;size:30;"`
	//[10] duplicate_mac_timer                            varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DuplicateMacTimer string `gorm:"column:duplicate_mac_timer;type:varchar;size:30;"`
	//[11] duplicate_mac_timer_max_count                  varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DuplicateMacTimerMaxCount string `gorm:"column:duplicate_mac_timer_max_count;type:varchar;size:30;"`
	//[12] config_type                                    varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType string `gorm:"column:config_type;type:varchar;size:30;"`
}

var evpnsTableInfo = &TableInfo{
	Name: "evpns",
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
			Name:               "fabric_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "FabricName",
			GoFieldType:        "string",
			JSONFieldName:      "fabric_name",
			ProtobufFieldName:  "fabric_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
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
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "device_ip",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "DeviceIP",
			GoFieldType:        "string",
			JSONFieldName:      "device_ip",
			ProtobufFieldName:  "device_ip",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "sw_evpn_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "SwEvpnName",
			GoFieldType:        "string",
			JSONFieldName:      "sw_evpn_name",
			ProtobufFieldName:  "sw_evpn_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "mac_aging_timeout",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "MacAgingTimeout",
			GoFieldType:        "string",
			JSONFieldName:      "mac_aging_timeout",
			ProtobufFieldName:  "mac_aging_timeout",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "mac_move_limit",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "MacMoveLimit",
			GoFieldType:        "string",
			JSONFieldName:      "mac_move_limit",
			ProtobufFieldName:  "mac_move_limit",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "mac_aging_conversational_timeout",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "MacAgingConversationalTimeout",
			GoFieldType:        "string",
			JSONFieldName:      "mac_aging_conversational_timeout",
			ProtobufFieldName:  "mac_aging_conversational_timeout",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "arp_agingtimeout",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "ArpAgingtimeout",
			GoFieldType:        "string",
			JSONFieldName:      "arp_agingtimeout",
			ProtobufFieldName:  "arp_agingtimeout",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "duplicate_mac_timer",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "DuplicateMacTimer",
			GoFieldType:        "string",
			JSONFieldName:      "duplicate_mac_timer",
			ProtobufFieldName:  "duplicate_mac_timer",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "duplicate_mac_timer_max_count",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "DuplicateMacTimerMaxCount",
			GoFieldType:        "string",
			JSONFieldName:      "duplicate_mac_timer_max_count",
			ProtobufFieldName:  "duplicate_mac_timer_max_count",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "config_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "ConfigType",
			GoFieldType:        "string",
			JSONFieldName:      "config_type",
			ProtobufFieldName:  "config_type",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Evpns) TableName() string {
	return "evpns"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Evpns) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Evpns) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Evpns) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Evpns) TableInfo() *TableInfo {
	return evpnsTableInfo
}
