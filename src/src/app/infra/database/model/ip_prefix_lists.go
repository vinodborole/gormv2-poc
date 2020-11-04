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


CREATE TABLE `ip_prefix_lists` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int NOT NULL,
  `device_id` int NOT NULL,
  `name` varchar(64) NOT NULL,
  `seq` int NOT NULL,
  `afi` varchar(20) NOT NULL DEFAULT 'ipv4',
  `action` varchar(20) DEFAULT NULL,
  `ip_prefix` varchar(100) DEFAULT NULL,
  `prefix_len_ge` int DEFAULT NULL,
  `prefix_len_le` int DEFAULT NULL,
  `config_type` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `ip_prefix_lists_fabric_id_fkey` (`fabric_id`),
  KEY `ip_prefix_lists_device_id_fkey` (`device_id`),
  CONSTRAINT `ip_prefix_lists_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `ip_prefix_lists_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 83,    "device_id": 10,    "seq": 53,    "afi": "LjZnHcGbCaZMNChXViMkVbblL",    "prefix_len_ge": 81,    "config_type": "pkBpiwABVmuQPtkYmZkXXpXJg",    "fabric_id": 88,    "name": "fUnnynnuocjVEuWVhQNWearcb",    "action": "eCudxuKclMaKBDbuBgIKUKToJ",    "ip_prefix": "EeGnJkqfPgLLjFOwmSkWOVlOg",    "prefix_len_le": 96}



*/

// IPPrefixLists struct is a row record of the ip_prefix_lists table in the myapp database
type IPPrefixLists struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID int32 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID int32 `gorm:"column:device_id;type:int;"`
	//[ 3] name                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Name string `gorm:"column:name;type:varchar;size:64;"`
	//[ 4] seq                                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Seq int32 `gorm:"column:seq;type:int;"`
	//[ 5] afi                                            varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: [ipv4]
	Afi string `gorm:"column:afi;type:varchar;size:20;default:ipv4;"`
	//[ 6] action                                         varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Action sql.NullString `gorm:"column:action;type:varchar;size:20;"`
	//[ 7] ip_prefix                                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	IPPrefix sql.NullString `gorm:"column:ip_prefix;type:varchar;size:100;"`
	//[ 8] prefix_len_ge                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PrefixLenGe sql.NullInt64 `gorm:"column:prefix_len_ge;type:int;"`
	//[ 9] prefix_len_le                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PrefixLenLe sql.NullInt64 `gorm:"column:prefix_len_le;type:int;"`
	//[10] config_type                                    text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	ConfigType string `gorm:"column:config_type;type:text;size:65535;"`
}

var ip_prefix_listsTableInfo = &TableInfo{
	Name: "ip_prefix_lists",
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
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FabricID",
			GoFieldType:        "int32",
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
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DeviceID",
			GoFieldType:        "int32",
			JSONFieldName:      "device_id",
			ProtobufFieldName:  "device_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "seq",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Seq",
			GoFieldType:        "int32",
			JSONFieldName:      "seq",
			ProtobufFieldName:  "seq",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "afi",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "Afi",
			GoFieldType:        "string",
			JSONFieldName:      "afi",
			ProtobufFieldName:  "afi",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "action",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "Action",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "action",
			ProtobufFieldName:  "action",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "ip_prefix",
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
			GoFieldName:        "IPPrefix",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ip_prefix",
			ProtobufFieldName:  "ip_prefix",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "prefix_len_ge",
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
			GoFieldName:        "PrefixLenGe",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "prefix_len_ge",
			ProtobufFieldName:  "prefix_len_ge",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "prefix_len_le",
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
			GoFieldName:        "PrefixLenLe",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "prefix_len_le",
			ProtobufFieldName:  "prefix_len_le",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		{
			Index:              10,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *IPPrefixLists) TableName() string {
	return "ip_prefix_lists"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *IPPrefixLists) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *IPPrefixLists) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *IPPrefixLists) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *IPPrefixLists) TableInfo() *TableInfo {
	return ip_prefix_listsTableInfo
}
