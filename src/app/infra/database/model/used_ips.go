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


CREATE TABLE `used_ips` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `ip_address` varchar(30) DEFAULT NULL,
  `ip_type` varchar(30) DEFAULT NULL,
  `interface_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `used_ips_fabric_id_fkey` (`fabric_id`),
  KEY `used_ips_device_id_fkey` (`device_id`),
  KEY `used_ips_interface_id_fkey` (`interface_id`),
  CONSTRAINT `used_ips_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `used_ips_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE,
  CONSTRAINT `used_ips_interface_id_fkey` FOREIGN KEY (`interface_id`) REFERENCES `phys_interfaces` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "interface_id": 75,    "id": 39,    "fabric_id": 23,    "device_id": 92,    "ip_address": "mgaQkHWXWGCCTViAscKZGpeTW",    "ip_type": "HlSvinOogaXLdleSKQKrTiLmo"}



*/

// UsedIps struct is a row record of the used_ips table in the myapp database
type UsedIps struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 3] ip_address                                     varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IPAddress sql.NullString `gorm:"column:ip_address;type:varchar;size:30;"`
	//[ 4] ip_type                                        varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IPType sql.NullString `gorm:"column:ip_type;type:varchar;size:30;"`
	//[ 5] interface_id                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	InterfaceID sql.NullInt64 `gorm:"column:interface_id;type:int;"`
}

var used_ipsTableInfo = &TableInfo{
	Name: "used_ips",
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
			Name:               "ip_address",
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
			GoFieldName:        "IPAddress",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ip_address",
			ProtobufFieldName:  "ip_address",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "ip_type",
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
			GoFieldName:        "IPType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ip_type",
			ProtobufFieldName:  "ip_type",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UsedIps) TableName() string {
	return "used_ips"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UsedIps) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UsedIps) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UsedIps) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UsedIps) TableInfo() *TableInfo {
	return used_ipsTableInfo
}
