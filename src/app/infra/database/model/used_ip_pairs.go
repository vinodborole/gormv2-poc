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


CREATE TABLE `used_ip_pairs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_one_id` int DEFAULT NULL,
  `device_two_id` int DEFAULT NULL,
  `ip_address_one` varchar(30) DEFAULT NULL,
  `ip_address_two` varchar(30) DEFAULT NULL,
  `ip_type` varchar(30) DEFAULT NULL,
  `interface_one_id` int DEFAULT NULL,
  `interface_two_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `used_ip_pairs_fabric_id_fkey` (`fabric_id`),
  KEY `used_ip_pairs_device_one_id_fkey` (`device_one_id`),
  KEY `used_ip_pairs_device_two_id_fkey` (`device_two_id`),
  KEY `used_ip_pairs_interface_one_id_fkey` (`interface_one_id`),
  KEY `used_ip_pairs_interface_two_id_fkey` (`interface_two_id`),
  CONSTRAINT `used_ip_pairs_device_one_id_fkey` FOREIGN KEY (`device_one_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `used_ip_pairs_device_two_id_fkey` FOREIGN KEY (`device_two_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `used_ip_pairs_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE,
  CONSTRAINT `used_ip_pairs_interface_one_id_fkey` FOREIGN KEY (`interface_one_id`) REFERENCES `phys_interfaces` (`id`) ON DELETE CASCADE,
  CONSTRAINT `used_ip_pairs_interface_two_id_fkey` FOREIGN KEY (`interface_two_id`) REFERENCES `phys_interfaces` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "ip_address_one": "SqeOnSsgrrWcNyjUhveYleBFf",    "ip_address_two": "GJpQfYIuxAnJbhwEAQmLNSkyB",    "id": 61,    "device_one_id": 37,    "device_two_id": 37,    "ip_type": "ZijDymuRVWPAKsJMesrJYPgCa",    "interface_one_id": 20,    "interface_two_id": 31,    "fabric_id": 44}



*/

// UsedIPPairs struct is a row record of the used_ip_pairs table in the myapp database
type UsedIPPairs struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_one_id                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceOneID sql.NullInt64 `gorm:"column:device_one_id;type:int;"`
	//[ 3] device_two_id                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceTwoID sql.NullInt64 `gorm:"column:device_two_id;type:int;"`
	//[ 4] ip_address_one                                 varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IPAddressOne sql.NullString `gorm:"column:ip_address_one;type:varchar;size:30;"`
	//[ 5] ip_address_two                                 varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IPAddressTwo sql.NullString `gorm:"column:ip_address_two;type:varchar;size:30;"`
	//[ 6] ip_type                                        varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IPType sql.NullString `gorm:"column:ip_type;type:varchar;size:30;"`
	//[ 7] interface_one_id                               int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	InterfaceOneID sql.NullInt64 `gorm:"column:interface_one_id;type:int;"`
	//[ 8] interface_two_id                               int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	InterfaceTwoID sql.NullInt64 `gorm:"column:interface_two_id;type:int;"`
}

var used_ip_pairsTableInfo = &TableInfo{
	Name: "used_ip_pairs",
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
			Name:               "device_one_id",
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
			GoFieldName:        "DeviceOneID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "device_one_id",
			ProtobufFieldName:  "device_one_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "device_two_id",
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
			GoFieldName:        "DeviceTwoID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "device_two_id",
			ProtobufFieldName:  "device_two_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "ip_address_one",
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
			GoFieldName:        "IPAddressOne",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ip_address_one",
			ProtobufFieldName:  "ip_address_one",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "ip_address_two",
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
			GoFieldName:        "IPAddressTwo",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ip_address_two",
			ProtobufFieldName:  "ip_address_two",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
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
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "interface_one_id",
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
			GoFieldName:        "InterfaceOneID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "interface_one_id",
			ProtobufFieldName:  "interface_one_id",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "interface_two_id",
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
			GoFieldName:        "InterfaceTwoID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "interface_two_id",
			ProtobufFieldName:  "interface_two_id",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UsedIPPairs) TableName() string {
	return "used_ip_pairs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UsedIPPairs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UsedIPPairs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UsedIPPairs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UsedIPPairs) TableInfo() *TableInfo {
	return used_ip_pairsTableInfo
}
