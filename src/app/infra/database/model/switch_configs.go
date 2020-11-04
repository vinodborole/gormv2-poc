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


CREATE TABLE `switch_configs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `device_id` int DEFAULT NULL,
  `fabric_id` int DEFAULT NULL,
  `device_ip` varchar(30) DEFAULT NULL,
  `local_as` varchar(30) DEFAULT NULL,
  `loopback_ip` varchar(30) DEFAULT NULL,
  `vtep_loopback_ip` varchar(30) DEFAULT NULL,
  `role` varchar(30) DEFAULT NULL,
  `as_config_type` varchar(30) DEFAULT NULL,
  `loopback_ip_config_type` varchar(30) DEFAULT NULL,
  `vtep_loopback_ip_config_type` varchar(30) DEFAULT NULL,
  `vtep_loop_back_id` varchar(30) DEFAULT NULL,
  `loop_back_id` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `switch_configs_device_id_fkey` (`device_id`),
  KEY `switch_configs_fabric_id_fkey` (`fabric_id`),
  CONSTRAINT `switch_configs_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `switch_configs_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 77,    "fabric_id": 96,    "device_ip": "LpHPqUuaNRHIbnQfdqtMEfaFZ",    "local_as": "mORwgcOrNfwmTQatCvKmtZwPZ",    "as_config_type": "AgycmcFylOdsHVOwxIISTIUsb",    "loopback_ip_config_type": "JMIZiPZWGbZXMwegChGtleuVI",    "device_id": 42,    "loopback_ip": "LNeCxaxktkUCxMCiuLUwLXNVm",    "vtep_loopback_ip": "ZjAdkXOMxElYJCYuiEawlKurI",    "role": "BVVnypPiHVksuaWRWQuKBSifS",    "vtep_loopback_ip_config_type": "UyKtrirDAcfkuXJBqqQqGpvrB",    "vtep_loop_back_id": "bLfZQPhdggwTTGoieVtJcQpNY",    "loop_back_id": "bbmYWguZtAdjmHTfqEoWADqXp"}



*/

// SwitchConfigs struct is a row record of the switch_configs table in the myapp database
type SwitchConfigs struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 2] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 3] device_ip                                      varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceIP sql.NullString `gorm:"column:device_ip;type:varchar;size:30;"`
	//[ 4] local_as                                       varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	LocalAs sql.NullString `gorm:"column:local_as;type:varchar;size:30;"`
	//[ 5] loopback_ip                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	LoopbackIP sql.NullString `gorm:"column:loopback_ip;type:varchar;size:30;"`
	//[ 6] vtep_loopback_ip                               varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	VtepLoopbackIP sql.NullString `gorm:"column:vtep_loopback_ip;type:varchar;size:30;"`
	//[ 7] role                                           varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Role sql.NullString `gorm:"column:role;type:varchar;size:30;"`
	//[ 8] as_config_type                                 varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	AsConfigType sql.NullString `gorm:"column:as_config_type;type:varchar;size:30;"`
	//[ 9] loopback_ip_config_type                        varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	LoopbackIPConfigType sql.NullString `gorm:"column:loopback_ip_config_type;type:varchar;size:30;"`
	//[10] vtep_loopback_ip_config_type                   varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	VtepLoopbackIPConfigType sql.NullString `gorm:"column:vtep_loopback_ip_config_type;type:varchar;size:30;"`
	//[11] vtep_loop_back_id                              varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	VtepLoopBackID sql.NullString `gorm:"column:vtep_loop_back_id;type:varchar;size:30;"`
	//[12] loop_back_id                                   varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	LoopBackID sql.NullString `gorm:"column:loop_back_id;type:varchar;size:30;"`
}

var switch_configsTableInfo = &TableInfo{
	Name: "switch_configs",
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
			ProtobufPos:        2,
		},

		{
			Index:              2,
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
			ProtobufPos:        3,
		},

		{
			Index:              3,
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
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "local_as",
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
			GoFieldName:        "LocalAs",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "local_as",
			ProtobufFieldName:  "local_as",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "loopback_ip",
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
			GoFieldName:        "LoopbackIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "loopback_ip",
			ProtobufFieldName:  "loopback_ip",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "vtep_loopback_ip",
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
			GoFieldName:        "VtepLoopbackIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "vtep_loopback_ip",
			ProtobufFieldName:  "vtep_loopback_ip",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "role",
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
			GoFieldName:        "Role",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "role",
			ProtobufFieldName:  "role",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "as_config_type",
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
			GoFieldName:        "AsConfigType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "as_config_type",
			ProtobufFieldName:  "as_config_type",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "loopback_ip_config_type",
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
			GoFieldName:        "LoopbackIPConfigType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "loopback_ip_config_type",
			ProtobufFieldName:  "loopback_ip_config_type",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "vtep_loopback_ip_config_type",
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
			GoFieldName:        "VtepLoopbackIPConfigType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "vtep_loopback_ip_config_type",
			ProtobufFieldName:  "vtep_loopback_ip_config_type",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "vtep_loop_back_id",
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
			GoFieldName:        "VtepLoopBackID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "vtep_loop_back_id",
			ProtobufFieldName:  "vtep_loop_back_id",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "loop_back_id",
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
			GoFieldName:        "LoopBackID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "loop_back_id",
			ProtobufFieldName:  "loop_back_id",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SwitchConfigs) TableName() string {
	return "switch_configs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SwitchConfigs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SwitchConfigs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SwitchConfigs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SwitchConfigs) TableInfo() *TableInfo {
	return switch_configsTableInfo
}
