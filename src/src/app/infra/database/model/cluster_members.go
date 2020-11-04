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


CREATE TABLE `cluster_members` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `cluster_id` int DEFAULT NULL,
  `interface_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `remote_device_id` int DEFAULT NULL,
  `interface_name` varchar(255) DEFAULT NULL,
  `interface_type` varchar(64) DEFAULT NULL,
  `interface_speed` int DEFAULT NULL,
  `remote_interface_id` int DEFAULT NULL,
  `remote_interface_name` varchar(255) DEFAULT NULL,
  `remote_interface_type` varchar(64) DEFAULT NULL,
  `remote_interface_speed` int DEFAULT NULL,
  `config_type` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `cluster_members_fabric_id_fkey` (`fabric_id`),
  KEY `cluster_members_cluster_id_fkey` (`cluster_id`),
  KEY `cluster_members_interface_id_fkey` (`interface_id`),
  KEY `cluster_members_device_id_fkey` (`device_id`),
  KEY `cluster_members_remote_device_id_fkey` (`remote_device_id`),
  KEY `cluster_members_remote_interface_id_fkey` (`remote_interface_id`),
  CONSTRAINT `cluster_members_cluster_id_fkey` FOREIGN KEY (`cluster_id`) REFERENCES `mct_cluster_configs` (`id`) ON DELETE CASCADE,
  CONSTRAINT `cluster_members_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `cluster_members_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE,
  CONSTRAINT `cluster_members_interface_id_fkey` FOREIGN KEY (`interface_id`) REFERENCES `phys_interfaces` (`id`) ON DELETE CASCADE,
  CONSTRAINT `cluster_members_remote_device_id_fkey` FOREIGN KEY (`remote_device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `cluster_members_remote_interface_id_fkey` FOREIGN KEY (`remote_interface_id`) REFERENCES `phys_interfaces` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "remote_device_id": 77,    "interface_name": "HLmvPLbAmhJUWTdDCpqXKKHnX",    "remote_interface_name": "PBNKsXKcMDbJRxXdbRVDnwCRY",    "cluster_id": 80,    "remote_interface_id": 72,    "config_type": "GaoyVaOamorFWIgjhJlxSfvLk",    "interface_id": 11,    "device_id": 15,    "remote_interface_type": "SQXumeYULYQttGYKIxhUlpGsf",    "id": 42,    "fabric_id": 15,    "interface_type": "yuemhJKDBLJpncWXxcaiZFWio",    "interface_speed": 41,    "remote_interface_speed": 77}



*/

// ClusterMembers struct is a row record of the cluster_members table in the myapp database
type ClusterMembers struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] cluster_id                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ClusterID sql.NullInt64 `gorm:"column:cluster_id;type:int;"`
	//[ 3] interface_id                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	InterfaceID sql.NullInt64 `gorm:"column:interface_id;type:int;"`
	//[ 4] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 5] remote_device_id                               int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RemoteDeviceID sql.NullInt64 `gorm:"column:remote_device_id;type:int;"`
	//[ 6] interface_name                                 varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	InterfaceName sql.NullString `gorm:"column:interface_name;type:varchar;size:255;"`
	//[ 7] interface_type                                 varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	InterfaceType sql.NullString `gorm:"column:interface_type;type:varchar;size:64;"`
	//[ 8] interface_speed                                int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	InterfaceSpeed sql.NullInt64 `gorm:"column:interface_speed;type:int;"`
	//[ 9] remote_interface_id                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RemoteInterfaceID sql.NullInt64 `gorm:"column:remote_interface_id;type:int;"`
	//[10] remote_interface_name                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	RemoteInterfaceName sql.NullString `gorm:"column:remote_interface_name;type:varchar;size:255;"`
	//[11] remote_interface_type                          varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	RemoteInterfaceType sql.NullString `gorm:"column:remote_interface_type;type:varchar;size:64;"`
	//[12] remote_interface_speed                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RemoteInterfaceSpeed sql.NullInt64 `gorm:"column:remote_interface_speed;type:int;"`
	//[13] config_type                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType sql.NullString `gorm:"column:config_type;type:varchar;size:30;"`
}

var cluster_membersTableInfo = &TableInfo{
	Name: "cluster_members",
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
			Name:               "cluster_id",
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
			GoFieldName:        "ClusterID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "cluster_id",
			ProtobufFieldName:  "cluster_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "interface_id",
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
			GoFieldName:        "InterfaceID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "interface_id",
			ProtobufFieldName:  "interface_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
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
			ProtobufPos:        5,
		},

		{
			Index:              5,
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
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "interface_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "InterfaceName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "interface_name",
			ProtobufFieldName:  "interface_name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "interface_type",
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
			GoFieldName:        "InterfaceType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "interface_type",
			ProtobufFieldName:  "interface_type",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "interface_speed",
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
			GoFieldName:        "InterfaceSpeed",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "interface_speed",
			ProtobufFieldName:  "interface_speed",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		{
			Index:              9,
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
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "remote_interface_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "RemoteInterfaceName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "remote_interface_name",
			ProtobufFieldName:  "remote_interface_name",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "remote_interface_type",
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
			GoFieldName:        "RemoteInterfaceType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "remote_interface_type",
			ProtobufFieldName:  "remote_interface_type",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "remote_interface_speed",
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
			GoFieldName:        "RemoteInterfaceSpeed",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "remote_interface_speed",
			ProtobufFieldName:  "remote_interface_speed",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		{
			Index:              13,
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
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ClusterMembers) TableName() string {
	return "cluster_members"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ClusterMembers) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ClusterMembers) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ClusterMembers) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ClusterMembers) TableInfo() *TableInfo {
	return cluster_membersTableInfo
}
