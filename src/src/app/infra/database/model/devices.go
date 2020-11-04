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


CREATE TABLE `devices` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `name` varchar(30) DEFAULT NULL,
  `ip_address` varchar(30) DEFAULT NULL,
  `device_role` varchar(30) DEFAULT NULL,
  `local_as` varchar(30) DEFAULT NULL,
  `device_type` varchar(30) DEFAULT NULL,
  `host_name` varchar(100) DEFAULT NULL,
  `asn` varchar(100) DEFAULT NULL,
  `vtep_loopback_id` varchar(100) DEFAULT NULL,
  `pod` varchar(30) DEFAULT NULL,
  `mct_role` int DEFAULT NULL,
  `provisioning_status` int DEFAULT NULL,
  `device_status` int DEFAULT NULL,
  `config_gen_reason` bigint DEFAULT NULL,
  `loopback_id` varchar(100) DEFAULT NULL,
  `model` varchar(100) DEFAULT NULL,
  `firmware_version` varchar(100) DEFAULT NULL,
  `rack` varchar(30) DEFAULT NULL,
  `is_admin_state_up` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `devices_fabric_id_fkey` (`fabric_id`),
  CONSTRAINT `devices_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "loopback_id": "gbcQIswTnSKIcxtCXQqwsSkbj",    "model": "DSywYkaVTXOUApaxWxgHFujTD",    "id": 55,    "ip_address": "bMElnAjnljWCVELxqxCbJVMat",    "local_as": "pysjdAHYjgHvynAjiNSiRFGJj",    "vtep_loopback_id": "UGyLKUGyEqMLhJBxRTWmNwoiQ",    "provisioning_status": 89,    "device_status": 68,    "rack": "jpKIjMnKcIXGUYEerjlSdoXfE",    "name": "DgKhZjgaoxJEDsYlRFgUOIVxr",    "device_role": "CCVFDgRXtBxDiEMvOBkQodkfm",    "device_type": "PkvKQameiQEnHBwcHALyNaywf",    "pod": "mZZPsFfRGLqIkRPlBWbvUvmTn",    "mct_role": 31,    "firmware_version": "FjSTLVPTUGdWyrJXaNJcNbhvk",    "config_gen_reason": 38,    "fabric_id": 32,    "host_name": "ZUkaubZrmeMXEGSPicjpaLukX",    "asn": "ejVRZeJayNOYbuMGDwuABIjwe",    "is_admin_state_up": 17}



*/

// Devices struct is a row record of the devices table in the myapp database
type Devices struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] name                                           varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Name sql.NullString `gorm:"column:name;type:varchar;size:30;"`
	//[ 3] ip_address                                     varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IPAddress sql.NullString `gorm:"column:ip_address;type:varchar;size:30;"`
	//[ 4] device_role                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceRole sql.NullString `gorm:"column:device_role;type:varchar;size:30;"`
	//[ 5] local_as                                       varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	LocalAs sql.NullString `gorm:"column:local_as;type:varchar;size:30;"`
	//[ 6] device_type                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceType sql.NullString `gorm:"column:device_type;type:varchar;size:30;"`
	//[ 7] host_name                                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	HostName sql.NullString `gorm:"column:host_name;type:varchar;size:100;"`
	//[ 8] asn                                            varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Asn sql.NullString `gorm:"column:asn;type:varchar;size:100;"`
	//[ 9] vtep_loopback_id                               varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	VtepLoopbackID sql.NullString `gorm:"column:vtep_loopback_id;type:varchar;size:100;"`
	//[10] pod                                            varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Pod sql.NullString `gorm:"column:pod;type:varchar;size:30;"`
	//[11] mct_role                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	MctRole sql.NullInt64 `gorm:"column:mct_role;type:int;"`
	//[12] provisioning_status                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ProvisioningStatus sql.NullInt64 `gorm:"column:provisioning_status;type:int;"`
	//[13] device_status                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceStatus sql.NullInt64 `gorm:"column:device_status;type:int;"`
	//[14] config_gen_reason                              bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	ConfigGenReason sql.NullInt64 `gorm:"column:config_gen_reason;type:bigint;"`
	//[15] loopback_id                                    varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	LoopbackID sql.NullString `gorm:"column:loopback_id;type:varchar;size:100;"`
	//[16] model                                          varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Model sql.NullString `gorm:"column:model;type:varchar;size:100;"`
	//[17] firmware_version                               varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	FirmwareVersion sql.NullString `gorm:"column:firmware_version;type:varchar;size:100;"`
	//[18] rack                                           varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Rack sql.NullString `gorm:"column:rack;type:varchar;size:30;"`
	//[19] is_admin_state_up                              tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	IsAdminStateUp int32 `gorm:"column:is_admin_state_up;type:tinyint;default:1;"`
}

var devicesTableInfo = &TableInfo{
	Name: "devices",
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
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
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
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "device_type",
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
			GoFieldName:        "DeviceType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "device_type",
			ProtobufFieldName:  "device_type",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "host_name",
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
			GoFieldName:        "HostName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "host_name",
			ProtobufFieldName:  "host_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "asn",
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
			GoFieldName:        "Asn",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "asn",
			ProtobufFieldName:  "asn",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "vtep_loopback_id",
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
			GoFieldName:        "VtepLoopbackID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "vtep_loopback_id",
			ProtobufFieldName:  "vtep_loopback_id",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
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
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "mct_role",
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
			GoFieldName:        "MctRole",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "mct_role",
			ProtobufFieldName:  "mct_role",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "provisioning_status",
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
			GoFieldName:        "ProvisioningStatus",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "provisioning_status",
			ProtobufFieldName:  "provisioning_status",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		{
			Index:              13,
			Name:               "device_status",
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
			GoFieldName:        "DeviceStatus",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "device_status",
			ProtobufFieldName:  "device_status",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		{
			Index:              14,
			Name:               "config_gen_reason",
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
			GoFieldName:        "ConfigGenReason",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "config_gen_reason",
			ProtobufFieldName:  "config_gen_reason",
			ProtobufType:       "int64",
			ProtobufPos:        15,
		},

		{
			Index:              15,
			Name:               "loopback_id",
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
			GoFieldName:        "LoopbackID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "loopback_id",
			ProtobufFieldName:  "loopback_id",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		{
			Index:              16,
			Name:               "model",
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
			GoFieldName:        "Model",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "model",
			ProtobufFieldName:  "model",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		{
			Index:              17,
			Name:               "firmware_version",
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
			GoFieldName:        "FirmwareVersion",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "firmware_version",
			ProtobufFieldName:  "firmware_version",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		{
			Index:              18,
			Name:               "rack",
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
			GoFieldName:        "Rack",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "rack",
			ProtobufFieldName:  "rack",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		{
			Index:              19,
			Name:               "is_admin_state_up",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsAdminStateUp",
			GoFieldType:        "int32",
			JSONFieldName:      "is_admin_state_up",
			ProtobufFieldName:  "is_admin_state_up",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *Devices) TableName() string {
	return "devices"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *Devices) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *Devices) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *Devices) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *Devices) TableInfo() *TableInfo {
	return devicesTableInfo
}
