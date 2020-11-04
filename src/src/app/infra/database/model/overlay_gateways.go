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


CREATE TABLE `overlay_gateways` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `fabric_name` varchar(30) NOT NULL,
  `device_id` int DEFAULT NULL,
  `device_ip` varchar(30) NOT NULL,
  `sw_overlay_gw_name` varchar(30) NOT NULL,
  `gw_type` varchar(30) NOT NULL,
  `vtep_loopback_port_number` varchar(30) NOT NULL,
  `map_vni_auto` varchar(30) NOT NULL,
  `optimized_replication` tinyint(1) DEFAULT '0',
  `underlay_mdt_default_group` varchar(100) DEFAULT NULL,
  `activate` varchar(30) NOT NULL,
  `config_type` varchar(30) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `overlay_gateways_fabric_id_fkey` (`fabric_id`),
  KEY `overlay_gateways_device_id_fkey` (`device_id`),
  CONSTRAINT `overlay_gateways_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `overlay_gateways_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "sw_overlay_gw_name": "ZYcbTpLiNCdekXcvJqVmkSmSW",    "vtep_loopback_port_number": "KJWTdFJNeFpiXxhUpFOrpLOSC",    "optimized_replication": 25,    "underlay_mdt_default_group": "fCyyfCmlwkEnRxFsadSMDEDIO",    "config_type": "KUwDhbdaEcRUSYSBncnVEEukl",    "activate": "qgVliBRAlCegwnJSZCHDShnAZ",    "id": 96,    "fabric_id": 3,    "fabric_name": "QVpstiwcyxoODsIYOTJrWqOwo",    "device_id": 22,    "device_ip": "RLrWbwJlIBOxqLKqCKxebfjpi",    "gw_type": "nxmxqAIYmrrSCMNHGhGbrJaGO",    "map_vni_auto": "nytfHtYmFVRvkuNejlhUKxsRq"}



*/

// OverlayGateways struct is a row record of the overlay_gateways table in the myapp database
type OverlayGateways struct {
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
	//[ 5] sw_overlay_gw_name                             varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	SwOverlayGwName string `gorm:"column:sw_overlay_gw_name;type:varchar;size:30;"`
	//[ 6] gw_type                                        varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	GwType string `gorm:"column:gw_type;type:varchar;size:30;"`
	//[ 7] vtep_loopback_port_number                      varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	VtepLoopbackPortNumber string `gorm:"column:vtep_loopback_port_number;type:varchar;size:30;"`
	//[ 8] map_vni_auto                                   varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	MapVniAuto string `gorm:"column:map_vni_auto;type:varchar;size:30;"`
	//[ 9] optimized_replication                          tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	OptimizedReplication sql.NullInt64 `gorm:"column:optimized_replication;type:tinyint;default:0;"`
	//[10] underlay_mdt_default_group                     varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	UnderlayMdtDefaultGroup sql.NullString `gorm:"column:underlay_mdt_default_group;type:varchar;size:100;"`
	//[11] activate                                       varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Activate string `gorm:"column:activate;type:varchar;size:30;"`
	//[12] config_type                                    varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType string `gorm:"column:config_type;type:varchar;size:30;"`
}

var overlay_gatewaysTableInfo = &TableInfo{
	Name: "overlay_gateways",
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
			Name:               "sw_overlay_gw_name",
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
			GoFieldName:        "SwOverlayGwName",
			GoFieldType:        "string",
			JSONFieldName:      "sw_overlay_gw_name",
			ProtobufFieldName:  "sw_overlay_gw_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "gw_type",
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
			GoFieldName:        "GwType",
			GoFieldType:        "string",
			JSONFieldName:      "gw_type",
			ProtobufFieldName:  "gw_type",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "vtep_loopback_port_number",
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
			GoFieldName:        "VtepLoopbackPortNumber",
			GoFieldType:        "string",
			JSONFieldName:      "vtep_loopback_port_number",
			ProtobufFieldName:  "vtep_loopback_port_number",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "map_vni_auto",
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
			GoFieldName:        "MapVniAuto",
			GoFieldType:        "string",
			JSONFieldName:      "map_vni_auto",
			ProtobufFieldName:  "map_vni_auto",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "optimized_replication",
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
			GoFieldName:        "OptimizedReplication",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "optimized_replication",
			ProtobufFieldName:  "optimized_replication",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "underlay_mdt_default_group",
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
			GoFieldName:        "UnderlayMdtDefaultGroup",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "underlay_mdt_default_group",
			ProtobufFieldName:  "underlay_mdt_default_group",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "activate",
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
			GoFieldName:        "Activate",
			GoFieldType:        "string",
			JSONFieldName:      "activate",
			ProtobufFieldName:  "activate",
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
func (o *OverlayGateways) TableName() string {
	return "overlay_gateways"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OverlayGateways) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OverlayGateways) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OverlayGateways) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OverlayGateways) TableInfo() *TableInfo {
	return overlay_gatewaysTableInfo
}
