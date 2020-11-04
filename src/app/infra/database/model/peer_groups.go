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


CREATE TABLE `peer_groups` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `peer_group_name` varchar(64) NOT NULL,
  `description` varchar(100) DEFAULT NULL,
  `bfd_enable` tinyint(1) DEFAULT NULL,
  `remote_asn` bigint DEFAULT NULL,
  `graceful_shut_timer` int DEFAULT NULL,
  `config_type` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `peer_groups_fabric_id_fkey` (`fabric_id`),
  KEY `peer_groups_device_id_fkey` (`device_id`),
  CONSTRAINT `peer_groups_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `peer_groups_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "fabric_id": 69,    "bfd_enable": 82,    "remote_asn": 8,    "graceful_shut_timer": 5,    "config_type": "xxfNuNHDoUSHvnnxVCTMchYRX",    "id": 44,    "device_id": 48,    "peer_group_name": "TcWbXDtCaLEEVrgkTNsPNaKTE",    "description": "CmhsjvkTItuGrbAgtYmxafkOe"}



*/

// PeerGroups struct is a row record of the peer_groups table in the myapp database
type PeerGroups struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 3] peer_group_name                                varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	PeerGroupName string `gorm:"column:peer_group_name;type:varchar;size:64;"`
	//[ 4] description                                    varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Description sql.NullString `gorm:"column:description;type:varchar;size:100;"`
	//[ 5] bfd_enable                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	BfdEnable sql.NullInt64 `gorm:"column:bfd_enable;type:tinyint;"`
	//[ 6] remote_asn                                     bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	RemoteAsn sql.NullInt64 `gorm:"column:remote_asn;type:bigint;"`
	//[ 7] graceful_shut_timer                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	GracefulShutTimer sql.NullInt64 `gorm:"column:graceful_shut_timer;type:int;"`
	//[ 8] config_type                                    text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	ConfigType string `gorm:"column:config_type;type:text;size:65535;"`
}

var peer_groupsTableInfo = &TableInfo{
	Name: "peer_groups",
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
			Name:               "peer_group_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "PeerGroupName",
			GoFieldType:        "string",
			JSONFieldName:      "peer_group_name",
			ProtobufFieldName:  "peer_group_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "description",
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
			GoFieldName:        "Description",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
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
			ProtobufPos:        6,
		},

		{
			Index:              6,
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
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "graceful_shut_timer",
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
			GoFieldName:        "GracefulShutTimer",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "graceful_shut_timer",
			ProtobufFieldName:  "graceful_shut_timer",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "config_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "ConfigType",
			GoFieldType:        "string",
			JSONFieldName:      "config_type",
			ProtobufFieldName:  "config_type",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PeerGroups) TableName() string {
	return "peer_groups"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PeerGroups) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PeerGroups) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PeerGroups) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PeerGroups) TableInfo() *TableInfo {
	return peer_groupsTableInfo
}
