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


CREATE TABLE `interface_switch_configs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `interface_id` int DEFAULT NULL,
  `int_type` varchar(30) DEFAULT NULL,
  `int_name` varchar(30) DEFAULT NULL,
  `donor_type` varchar(30) DEFAULT NULL,
  `donor_name` varchar(30) DEFAULT NULL,
  `ip_address` varchar(30) DEFAULT NULL,
  `ip_pim_sparse` tinyint(1) DEFAULT '0',
  `config_type` varchar(30) DEFAULT NULL,
  `description` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `interface_switch_configs_fabric_id_fkey` (`fabric_id`),
  KEY `interface_switch_configs_device_id_fkey` (`device_id`),
  KEY `interface_switch_configs_interface_id_fkey` (`interface_id`),
  CONSTRAINT `interface_switch_configs_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `interface_switch_configs_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE,
  CONSTRAINT `interface_switch_configs_interface_id_fkey` FOREIGN KEY (`interface_id`) REFERENCES `phys_interfaces` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "donor_name": "dvmDDGvhqqEqpjJXFRefrkBNj",    "ip_address": "krNcMVcBaymGisHWaLjFrMqjr",    "ip_pim_sparse": 96,    "config_type": "VrQLwjkQGLiVmNXwwopvSayDU",    "fabric_id": 89,    "device_id": 78,    "int_type": "sVCcQdEJDjKtoswYwiTGrshFr",    "donor_type": "eyHccCkVuBogQiGGZuRkbiema",    "description": "IvkFsMjvsJCsNPNixrbimOAUD",    "id": 93,    "interface_id": 79,    "int_name": "CDmyoSWTlcOTmgplwwwkMgMwG"}



*/

// InterfaceSwitchConfigs struct is a row record of the interface_switch_configs table in the myapp database
type InterfaceSwitchConfigs struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 3] interface_id                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	InterfaceID sql.NullInt64 `gorm:"column:interface_id;type:int;"`
	//[ 4] int_type                                       varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IntType sql.NullString `gorm:"column:int_type;type:varchar;size:30;"`
	//[ 5] int_name                                       varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IntName sql.NullString `gorm:"column:int_name;type:varchar;size:30;"`
	//[ 6] donor_type                                     varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DonorType sql.NullString `gorm:"column:donor_type;type:varchar;size:30;"`
	//[ 7] donor_name                                     varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DonorName sql.NullString `gorm:"column:donor_name;type:varchar;size:30;"`
	//[ 8] ip_address                                     varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IPAddress sql.NullString `gorm:"column:ip_address;type:varchar;size:30;"`
	//[ 9] ip_pim_sparse                                  tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IPPimSparse sql.NullInt64 `gorm:"column:ip_pim_sparse;type:tinyint;default:0;"`
	//[10] config_type                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType sql.NullString `gorm:"column:config_type;type:varchar;size:30;"`
	//[11] description                                    varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Description sql.NullString `gorm:"column:description;type:varchar;size:100;"`
}

var interface_switch_configsTableInfo = &TableInfo{
	Name: "interface_switch_configs",
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
			Name:               "int_type",
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
			GoFieldName:        "IntType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "int_type",
			ProtobufFieldName:  "int_type",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "int_name",
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
			GoFieldName:        "IntName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "int_name",
			ProtobufFieldName:  "int_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "donor_type",
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
			GoFieldName:        "DonorType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "donor_type",
			ProtobufFieldName:  "donor_type",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "donor_name",
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
			GoFieldName:        "DonorName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "donor_name",
			ProtobufFieldName:  "donor_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
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
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "ip_pim_sparse",
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
			GoFieldName:        "IPPimSparse",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "ip_pim_sparse",
			ProtobufFieldName:  "ip_pim_sparse",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		{
			Index:              10,
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
			ProtobufPos:        11,
		},

		{
			Index:              11,
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
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *InterfaceSwitchConfigs) TableName() string {
	return "interface_switch_configs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *InterfaceSwitchConfigs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *InterfaceSwitchConfigs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *InterfaceSwitchConfigs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *InterfaceSwitchConfigs) TableInfo() *TableInfo {
	return interface_switch_configsTableInfo
}
