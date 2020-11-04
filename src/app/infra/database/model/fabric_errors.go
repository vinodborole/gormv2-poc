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


CREATE TABLE `fabric_errors` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `fabric_name` varchar(30) DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `device_ip` varchar(30) DEFAULT NULL,
  `device_role` varchar(30) DEFAULT NULL,
  `error_type` varchar(30) DEFAULT NULL,
  `error_reason` text,
  PRIMARY KEY (`id`),
  KEY `fabric_errors_fabric_id_fkey` (`fabric_id`),
  KEY `fabric_errors_device_id_fkey` (`device_id`),
  CONSTRAINT `fabric_errors_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fabric_errors_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "error_type": "tMwMQjkPjyDwLZYrLThltBZsR",    "error_reason": "AamoocqRXbSpDcSAGarQROJGa",    "id": 68,    "fabric_id": 25,    "fabric_name": "lLJclGxJYGpPjuLhGDcRRHtLe",    "device_id": 33,    "device_ip": "ZSuSSwdcwuWtrQKfktINuMuAU",    "device_role": "dgnfokkMhQrsvJKsVBNpUehNP"}



*/

// FabricErrors struct is a row record of the fabric_errors table in the myapp database
type FabricErrors struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] fabric_name                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	FabricName sql.NullString `gorm:"column:fabric_name;type:varchar;size:30;"`
	//[ 3] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 4] device_ip                                      varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceIP sql.NullString `gorm:"column:device_ip;type:varchar;size:30;"`
	//[ 5] device_role                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceRole sql.NullString `gorm:"column:device_role;type:varchar;size:30;"`
	//[ 6] error_type                                     varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ErrorType sql.NullString `gorm:"column:error_type;type:varchar;size:30;"`
	//[ 7] error_reason                                   text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	ErrorReason sql.NullString `gorm:"column:error_reason;type:text;size:65535;"`
}

var fabric_errorsTableInfo = &TableInfo{
	Name: "fabric_errors",
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
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "FabricName",
			GoFieldType:        "sql.NullString",
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
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "DeviceIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "device_ip",
			ProtobufFieldName:  "device_ip",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
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
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "error_type",
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
			GoFieldName:        "ErrorType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "error_type",
			ProtobufFieldName:  "error_type",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "error_reason",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "ErrorReason",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "error_reason",
			ProtobufFieldName:  "error_reason",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FabricErrors) TableName() string {
	return "fabric_errors"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FabricErrors) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FabricErrors) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FabricErrors) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FabricErrors) TableInfo() *TableInfo {
	return fabric_errorsTableInfo
}
