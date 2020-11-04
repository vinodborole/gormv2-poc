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


CREATE TABLE `phys_interfaces` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `inventory_interface_id` int DEFAULT NULL,
  `int_type` varchar(30) DEFAULT NULL,
  `int_name` varchar(30) DEFAULT NULL,
  `interface_speed` varchar(30) DEFAULT NULL,
  `ip_address` varchar(30) DEFAULT NULL,
  `ip_pim_sparse` tinyint(1) DEFAULT '0',
  `identifier` varchar(30) DEFAULT NULL,
  `mac` varchar(30) DEFAULT NULL,
  `config_type` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `phys_interfaces_fabric_id_fkey` (`fabric_id`),
  KEY `phys_interfaces_device_id_fkey` (`device_id`),
  CONSTRAINT `phys_interfaces_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `phys_interfaces_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "inventory_interface_id": 84,    "int_type": "KfinFtkEqjRjJAAxrXYdbjdbp",    "interface_speed": "vTiOUAWbatyTfsOeNWRbKAskY",    "identifier": "AUvyTMCkSInbJXWrwcpIMuSbx",    "mac": "ahlngQFroICjfwnwFFAMtVuGu",    "id": 80,    "fabric_id": 18,    "device_id": 47,    "config_type": "ItGHcwcBaauItaDigByFuNTCY",    "int_name": "ocNoZTdfpPjutVYJyHOjXwuYt",    "ip_address": "mhmcrMRDnNQBvmGbCRyMlwJPW",    "ip_pim_sparse": 75}



*/

// PhysInterfaces struct is a row record of the phys_interfaces table in the myapp database
type PhysInterfaces struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 3] inventory_interface_id                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	InventoryInterfaceID sql.NullInt64 `gorm:"column:inventory_interface_id;type:int;"`
	//[ 4] int_type                                       varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IntType sql.NullString `gorm:"column:int_type;type:varchar;size:30;"`
	//[ 5] int_name                                       varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IntName sql.NullString `gorm:"column:int_name;type:varchar;size:30;"`
	//[ 6] interface_speed                                varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	InterfaceSpeed sql.NullString `gorm:"column:interface_speed;type:varchar;size:30;"`
	//[ 7] ip_address                                     varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IPAddress sql.NullString `gorm:"column:ip_address;type:varchar;size:30;"`
	//[ 8] ip_pim_sparse                                  tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	IPPimSparse sql.NullInt64 `gorm:"column:ip_pim_sparse;type:tinyint;"`
	//[ 9] identifier                                     varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Identifier sql.NullString `gorm:"column:identifier;type:varchar;size:30;"`
	//[10] mac                                            varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Mac sql.NullString `gorm:"column:mac;type:varchar;size:30;"`
	//[11] config_type                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType sql.NullString `gorm:"column:config_type;type:varchar;size:30;"`
}

var phys_interfacesTableInfo = &TableInfo{
	Name: "phys_interfaces",
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
			Name:               "inventory_interface_id",
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
			GoFieldName:        "InventoryInterfaceID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "inventory_interface_id",
			ProtobufFieldName:  "inventory_interface_id",
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
			Name:               "interface_speed",
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
			GoFieldName:        "InterfaceSpeed",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "interface_speed",
			ProtobufFieldName:  "interface_speed",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
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
			ProtobufPos:        8,
		},

		{
			Index:              8,
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
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "identifier",
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
			GoFieldName:        "Identifier",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "identifier",
			ProtobufFieldName:  "identifier",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "mac",
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
			GoFieldName:        "Mac",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mac",
			ProtobufFieldName:  "mac",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
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
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PhysInterfaces) TableName() string {
	return "phys_interfaces"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PhysInterfaces) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PhysInterfaces) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PhysInterfaces) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PhysInterfaces) TableInfo() *TableInfo {
	return phys_interfacesTableInfo
}
