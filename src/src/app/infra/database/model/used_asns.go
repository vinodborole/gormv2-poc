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


CREATE TABLE `used_asns` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `asn` bigint DEFAULT NULL,
  `device_role` varchar(30) DEFAULT NULL,
  `pod` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `used_asns_fabric_id_fkey` (`fabric_id`),
  KEY `used_asns_device_id_fkey` (`device_id`),
  CONSTRAINT `used_asns_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `used_asns_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "fabric_id": 14,    "device_id": 14,    "asn": 86,    "device_role": "YQNaJUqCBkixODpaGpcWYupQy",    "pod": "mlHViPaTwDCGdAPGQoiSCIKVb",    "id": 22}



*/

// UsedAsns struct is a row record of the used_asns table in the myapp database
type UsedAsns struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 3] asn                                            bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	Asn sql.NullInt64 `gorm:"column:asn;type:bigint;"`
	//[ 4] device_role                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceRole sql.NullString `gorm:"column:device_role;type:varchar;size:30;"`
	//[ 5] pod                                            varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Pod sql.NullString `gorm:"column:pod;type:varchar;size:30;"`
}

var used_asnsTableInfo = &TableInfo{
	Name: "used_asns",
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
			Name:               "asn",
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
			GoFieldName:        "Asn",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "asn",
			ProtobufFieldName:  "asn",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "device_role",
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
			GoFieldName:        "DeviceRole",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "device_role",
			ProtobufFieldName:  "device_role",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "pod",
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
			GoFieldName:        "Pod",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "pod",
			ProtobufFieldName:  "pod",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *UsedAsns) TableName() string {
	return "used_asns"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UsedAsns) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UsedAsns) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UsedAsns) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *UsedAsns) TableInfo() *TableInfo {
	return used_asnsTableInfo
}
