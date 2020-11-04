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


CREATE TABLE `bgp_ip_afs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `router_bgp_id` int DEFAULT NULL,
  `neighbor_afi` varchar(100) DEFAULT NULL,
  `neighbor_safi` varchar(100) DEFAULT NULL,
  `vrf` varchar(100) DEFAULT NULL,
  `graceful_restart` tinyint(1) DEFAULT NULL,
  `max_path` int DEFAULT NULL,
  `config_type` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `bgp_ip_afs_fabric_id_fkey` (`fabric_id`),
  KEY `bgp_ip_afs_device_id_fkey` (`device_id`),
  KEY `bgp_ip_afs_router_bgp_id_fkey` (`router_bgp_id`),
  CONSTRAINT `bgp_ip_afs_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `bgp_ip_afs_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE,
  CONSTRAINT `bgp_ip_afs_router_bgp_id_fkey` FOREIGN KEY (`router_bgp_id`) REFERENCES `router_bgps` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "neighbor_safi": "LirmuhojeNeGJwmRRKbLqXyEZ",    "max_path": 9,    "id": 9,    "device_id": 95,    "router_bgp_id": 78,    "graceful_restart": 35,    "config_type": "PQqimUAkJdsjFyyCfkKvcgfcq",    "fabric_id": 44,    "neighbor_afi": "FTgAYKDdqxsPLKibVdKXBAiKh",    "vrf": "PAFgPsyktwTdRBCNMAHpnAJfC"}



*/

// BgpIPAfs struct is a row record of the bgp_ip_afs table in the myapp database
type BgpIPAfs struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 3] router_bgp_id                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RouterBgpID sql.NullInt64 `gorm:"column:router_bgp_id;type:int;"`
	//[ 4] neighbor_afi                                   varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	NeighborAfi sql.NullString `gorm:"column:neighbor_afi;type:varchar;size:100;"`
	//[ 5] neighbor_safi                                  varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	NeighborSafi sql.NullString `gorm:"column:neighbor_safi;type:varchar;size:100;"`
	//[ 6] vrf                                            varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Vrf sql.NullString `gorm:"column:vrf;type:varchar;size:100;"`
	//[ 7] graceful_restart                               tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	GracefulRestart sql.NullInt64 `gorm:"column:graceful_restart;type:tinyint;"`
	//[ 8] max_path                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	MaxPath sql.NullInt64 `gorm:"column:max_path;type:int;"`
	//[ 9] config_type                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType sql.NullString `gorm:"column:config_type;type:varchar;size:30;"`
}

var bgp_ip_afsTableInfo = &TableInfo{
	Name: "bgp_ip_afs",
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
			ProtobufPos:        5,
		},

		{
			Index:              5,
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
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "vrf",
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
			GoFieldName:        "Vrf",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "vrf",
			ProtobufFieldName:  "vrf",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
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
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "max_path",
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
			GoFieldName:        "MaxPath",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "max_path",
			ProtobufFieldName:  "max_path",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		{
			Index:              9,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BgpIPAfs) TableName() string {
	return "bgp_ip_afs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BgpIPAfs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BgpIPAfs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BgpIPAfs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BgpIPAfs) TableInfo() *TableInfo {
	return bgp_ip_afsTableInfo
}
