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


CREATE TABLE `router_pims` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int NOT NULL,
  `device_id` int NOT NULL,
  `vrf` text,
  `ssm_enable` tinyint(1) DEFAULT NULL,
  `mdt_range` varchar(100) DEFAULT NULL,
  `config_type` varchar(30) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `router_pims_fabric_id_fkey` (`fabric_id`),
  KEY `router_pims_device_id_fkey` (`device_id`),
  CONSTRAINT `router_pims_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `router_pims_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "mdt_range": "GgaDSalopdNZCAivejlWFVvVt",    "config_type": "ibHSVqPkeIjrZlXhfbFZoaTge",    "id": 59,    "fabric_id": 97,    "device_id": 47,    "vrf": "SJRjZsyFyAmDxHtjdCwyxsFPA",    "ssm_enable": 68}



*/

// RouterPims struct is a row record of the router_pims table in the myapp database
type RouterPims struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID int32 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID int32 `gorm:"column:device_id;type:int;"`
	//[ 3] vrf                                            text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Vrf sql.NullString `gorm:"column:vrf;type:text;size:65535;"`
	//[ 4] ssm_enable                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	SsmEnable sql.NullInt64 `gorm:"column:ssm_enable;type:tinyint;"`
	//[ 5] mdt_range                                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	MdtRange sql.NullString `gorm:"column:mdt_range;type:varchar;size:100;"`
	//[ 6] config_type                                    varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType string `gorm:"column:config_type;type:varchar;size:30;"`
}

var router_pimsTableInfo = &TableInfo{
	Name: "router_pims",
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
			Name:               "vrf",
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
			GoFieldName:        "Vrf",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "vrf",
			ProtobufFieldName:  "vrf",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "ssm_enable",
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
			GoFieldName:        "SsmEnable",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "ssm_enable",
			ProtobufFieldName:  "ssm_enable",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "mdt_range",
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
			GoFieldName:        "MdtRange",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mdt_range",
			ProtobufFieldName:  "mdt_range",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "config_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "ConfigType",
			GoFieldType:        "string",
			JSONFieldName:      "config_type",
			ProtobufFieldName:  "config_type",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RouterPims) TableName() string {
	return "router_pims"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RouterPims) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RouterPims) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RouterPims) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RouterPims) TableInfo() *TableInfo {
	return router_pimsTableInfo
}
