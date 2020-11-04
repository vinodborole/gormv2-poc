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


CREATE TABLE `lldp_neighbors` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_one_id` int DEFAULT NULL,
  `device_two_id` int DEFAULT NULL,
  `interface_one_id` int DEFAULT NULL,
  `interface_two_id` int DEFAULT NULL,
  `inventory_interface_one_id` int DEFAULT NULL,
  `inventory_interface_two_id` int DEFAULT NULL,
  `device_one_role` varchar(30) DEFAULT NULL,
  `device_two_role` varchar(30) DEFAULT NULL,
  `interface_one_name` varchar(30) DEFAULT NULL,
  `interface_one_type` varchar(30) DEFAULT NULL,
  `interface_one_ip` varchar(30) DEFAULT NULL,
  `interface_two_name` varchar(30) DEFAULT NULL,
  `interface_two_type` varchar(30) DEFAULT NULL,
  `interface_one_speed` varchar(30) DEFAULT NULL,
  `interface_two_speed` varchar(30) DEFAULT NULL,
  `interface_two_ip` varchar(30) DEFAULT NULL,
  `config_type` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `lldp_neighbors_fabric_id_fkey` (`fabric_id`),
  KEY `lldp_neighbors_device_one_id_fkey` (`device_one_id`),
  KEY `lldp_neighbors_device_two_id_fkey` (`device_two_id`),
  KEY `lldp_neighbors_interface_one_id_fkey` (`interface_one_id`),
  KEY `lldp_neighbors_interface_two_id_fkey` (`interface_two_id`),
  CONSTRAINT `lldp_neighbors_device_one_id_fkey` FOREIGN KEY (`device_one_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `lldp_neighbors_device_two_id_fkey` FOREIGN KEY (`device_two_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `lldp_neighbors_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE,
  CONSTRAINT `lldp_neighbors_interface_one_id_fkey` FOREIGN KEY (`interface_one_id`) REFERENCES `phys_interfaces` (`id`) ON DELETE CASCADE,
  CONSTRAINT `lldp_neighbors_interface_two_id_fkey` FOREIGN KEY (`interface_two_id`) REFERENCES `phys_interfaces` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "device_one_role": "VaSaImjxwnyHWQCLNciYdkoLV",    "interface_two_speed": "pxOkmYhEaTSuGSmeGAsXEgwZn",    "interface_two_ip": "pRmdWFICgQRsYOZuGxrbtPNLp",    "device_one_id": 3,    "interface_one_id": 24,    "interface_one_name": "iuxXIrxwIuxmPGHLawTkZHkXX",    "interface_two_name": "CjnLuqOZmwcVFaQXsnJldIEJI",    "interface_one_speed": "OlkZZaQdMpAWbAlCpfqARDxex",    "config_type": "RUlXvDPBKtYlwqKaTIQaYwreU",    "device_two_id": 32,    "interface_two_id": 1,    "inventory_interface_one_id": 20,    "device_two_role": "PPrMJAAtCdpsUhalwQnJMFlQy",    "interface_one_type": "GwdMNLEJMBTghAIYZhEvumCOX",    "interface_two_type": "rkUMtmNMRsBIRMdmkmCQbdDWf",    "id": 7,    "fabric_id": 5,    "inventory_interface_two_id": 73,    "interface_one_ip": "PrAIWyWCvTlMcfnVGkbWJFiab"}



*/

// LldpNeighbors struct is a row record of the lldp_neighbors table in the myapp database
type LldpNeighbors struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_one_id                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceOneID sql.NullInt64 `gorm:"column:device_one_id;type:int;"`
	//[ 3] device_two_id                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceTwoID sql.NullInt64 `gorm:"column:device_two_id;type:int;"`
	//[ 4] interface_one_id                               int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	InterfaceOneID sql.NullInt64 `gorm:"column:interface_one_id;type:int;"`
	//[ 5] interface_two_id                               int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	InterfaceTwoID sql.NullInt64 `gorm:"column:interface_two_id;type:int;"`
	//[ 6] inventory_interface_one_id                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	InventoryInterfaceOneID sql.NullInt64 `gorm:"column:inventory_interface_one_id;type:int;"`
	//[ 7] inventory_interface_two_id                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	InventoryInterfaceTwoID sql.NullInt64 `gorm:"column:inventory_interface_two_id;type:int;"`
	//[ 8] device_one_role                                varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceOneRole sql.NullString `gorm:"column:device_one_role;type:varchar;size:30;"`
	//[ 9] device_two_role                                varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceTwoRole sql.NullString `gorm:"column:device_two_role;type:varchar;size:30;"`
	//[10] interface_one_name                             varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	InterfaceOneName sql.NullString `gorm:"column:interface_one_name;type:varchar;size:30;"`
	//[11] interface_one_type                             varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	InterfaceOneType sql.NullString `gorm:"column:interface_one_type;type:varchar;size:30;"`
	//[12] interface_one_ip                               varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	InterfaceOneIP sql.NullString `gorm:"column:interface_one_ip;type:varchar;size:30;"`
	//[13] interface_two_name                             varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	InterfaceTwoName sql.NullString `gorm:"column:interface_two_name;type:varchar;size:30;"`
	//[14] interface_two_type                             varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	InterfaceTwoType sql.NullString `gorm:"column:interface_two_type;type:varchar;size:30;"`
	//[15] interface_one_speed                            varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	InterfaceOneSpeed sql.NullString `gorm:"column:interface_one_speed;type:varchar;size:30;"`
	//[16] interface_two_speed                            varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	InterfaceTwoSpeed sql.NullString `gorm:"column:interface_two_speed;type:varchar;size:30;"`
	//[17] interface_two_ip                               varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	InterfaceTwoIP sql.NullString `gorm:"column:interface_two_ip;type:varchar;size:30;"`
	//[18] config_type                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType sql.NullString `gorm:"column:config_type;type:varchar;size:30;"`
}

var lldp_neighborsTableInfo = &TableInfo{
	Name: "lldp_neighbors",
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
			ProtobufPos:        5,
		},

		{
			Index:              5,
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
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "inventory_interface_one_id",
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
			GoFieldName:        "InventoryInterfaceOneID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "inventory_interface_one_id",
			ProtobufFieldName:  "inventory_interface_one_id",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "inventory_interface_two_id",
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
			GoFieldName:        "InventoryInterfaceTwoID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "inventory_interface_two_id",
			ProtobufFieldName:  "inventory_interface_two_id",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "device_one_role",
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
			GoFieldName:        "DeviceOneRole",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "device_one_role",
			ProtobufFieldName:  "device_one_role",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "device_two_role",
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
			GoFieldName:        "DeviceTwoRole",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "device_two_role",
			ProtobufFieldName:  "device_two_role",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "interface_one_name",
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
			GoFieldName:        "InterfaceOneName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "interface_one_name",
			ProtobufFieldName:  "interface_one_name",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "interface_one_type",
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
			GoFieldName:        "InterfaceOneType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "interface_one_type",
			ProtobufFieldName:  "interface_one_type",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "interface_one_ip",
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
			GoFieldName:        "InterfaceOneIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "interface_one_ip",
			ProtobufFieldName:  "interface_one_ip",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		{
			Index:              13,
			Name:               "interface_two_name",
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
			GoFieldName:        "InterfaceTwoName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "interface_two_name",
			ProtobufFieldName:  "interface_two_name",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		{
			Index:              14,
			Name:               "interface_two_type",
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
			GoFieldName:        "InterfaceTwoType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "interface_two_type",
			ProtobufFieldName:  "interface_two_type",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		{
			Index:              15,
			Name:               "interface_one_speed",
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
			GoFieldName:        "InterfaceOneSpeed",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "interface_one_speed",
			ProtobufFieldName:  "interface_one_speed",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		{
			Index:              16,
			Name:               "interface_two_speed",
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
			GoFieldName:        "InterfaceTwoSpeed",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "interface_two_speed",
			ProtobufFieldName:  "interface_two_speed",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		{
			Index:              17,
			Name:               "interface_two_ip",
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
			GoFieldName:        "InterfaceTwoIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "interface_two_ip",
			ProtobufFieldName:  "interface_two_ip",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		{
			Index:              18,
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
			ProtobufPos:        19,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LldpNeighbors) TableName() string {
	return "lldp_neighbors"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LldpNeighbors) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LldpNeighbors) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LldpNeighbors) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LldpNeighbors) TableInfo() *TableInfo {
	return lldp_neighborsTableInfo
}
