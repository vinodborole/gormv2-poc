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


CREATE TABLE `bgp_evpn_afs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `router_bgp_id` int DEFAULT NULL,
  `retain_route_target_all` tinyint(1) DEFAULT NULL,
  `graceful_restart` tinyint(1) DEFAULT NULL,
  `config_type` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `bgp_evpn_afs_fabric_id_fkey` (`fabric_id`),
  KEY `bgp_evpn_afs_device_id_fkey` (`device_id`),
  KEY `bgp_evpn_afs_router_bgp_id_fkey` (`router_bgp_id`),
  CONSTRAINT `bgp_evpn_afs_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `bgp_evpn_afs_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE,
  CONSTRAINT `bgp_evpn_afs_router_bgp_id_fkey` FOREIGN KEY (`router_bgp_id`) REFERENCES `router_bgps` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "retain_route_target_all": 10,    "graceful_restart": 66,    "config_type": "bchKfPRgGBnkCUZplYLUEOUtR",    "id": 22,    "fabric_id": 59,    "device_id": 52,    "router_bgp_id": 28}



*/

// BgpEvpnAfs struct is a row record of the bgp_evpn_afs table in the myapp database
type BgpEvpnAfs struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 3] router_bgp_id                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RouterBgpID sql.NullInt64 `gorm:"column:router_bgp_id;type:int;"`
	//[ 4] retain_route_target_all                        tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	RetainRouteTargetAll sql.NullInt64 `gorm:"column:retain_route_target_all;type:tinyint;"`
	//[ 5] graceful_restart                               tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	GracefulRestart sql.NullInt64 `gorm:"column:graceful_restart;type:tinyint;"`
	//[ 6] config_type                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType sql.NullString `gorm:"column:config_type;type:varchar;size:30;"`
}

var bgp_evpn_afsTableInfo = &TableInfo{
	Name: "bgp_evpn_afs",
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
			Name:               "router_bgp_id",
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
			GoFieldName:        "RouterBgpID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "router_bgp_id",
			ProtobufFieldName:  "router_bgp_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "retain_route_target_all",
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
			GoFieldName:        "RetainRouteTargetAll",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "retain_route_target_all",
			ProtobufFieldName:  "retain_route_target_all",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "graceful_restart",
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
			GoFieldName:        "GracefulRestart",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "graceful_restart",
			ProtobufFieldName:  "graceful_restart",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BgpEvpnAfs) TableName() string {
	return "bgp_evpn_afs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BgpEvpnAfs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BgpEvpnAfs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BgpEvpnAfs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BgpEvpnAfs) TableInfo() *TableInfo {
	return bgp_evpn_afsTableInfo
}
