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


CREATE TABLE `mct_cluster_details` (
  `id` int NOT NULL AUTO_INCREMENT,
  `node_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `principal_switch_mac` varchar(64) DEFAULT NULL,
  `node_internal_ip` varchar(30) DEFAULT NULL,
  `node_public_ip` varchar(30) DEFAULT NULL,
  `node_principal` varchar(30) DEFAULT NULL,
  `node_is_local` varchar(20) DEFAULT NULL,
  `serial_nnum` varchar(30) DEFAULT NULL,
  `node_condition` varchar(30) DEFAULT NULL,
  `node_status` varchar(60) DEFAULT NULL,
  `firmware_version` varchar(128) DEFAULT NULL,
  `node_mac` varchar(30) DEFAULT NULL,
  `node_switch_type` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `mct_cluster_details_device_id_fkey` (`device_id`),
  CONSTRAINT `mct_cluster_details_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "node_id": 26,    "node_internal_ip": "vigZXjuUOFjUuwIoypVjUNHlg",    "node_is_local": "tYIbksFsnStArhcKMJoSwvbsZ",    "id": 85,    "principal_switch_mac": "vouguHDwIDkKHcrOKPveRjUbh",    "device_id": 42,    "node_principal": "ukKmolWjAfoIBCbvrFYhLvboa",    "serial_nnum": "YqDRBGWOokbVTUBkeIBdDBtlr",    "node_status": "FCRwvgIqYBkXUFoAslAViwYGc",    "node_switch_type": "KWTOAaArJyXBeZiRRAcXISuHV",    "node_public_ip": "jmsRXZuMjviudNiToTQPJDsyi",    "node_condition": "vKwLMICGUPwpIkQPuDUKUjRmB",    "firmware_version": "YWVlQJHeDZaNvCHVMKZuWqLWk",    "node_mac": "PiLeexHZUGhyaIPQQHnFwbcDQ"}



*/

// MctClusterDetails struct is a row record of the mct_cluster_details table in the myapp database
type MctClusterDetails struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] node_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	NodeID sql.NullInt64 `gorm:"column:node_id;type:int;"`
	//[ 2] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 3] principal_switch_mac                           varchar(64)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	PrincipalSwitchMac sql.NullString `gorm:"column:principal_switch_mac;type:varchar;size:64;"`
	//[ 4] node_internal_ip                               varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	NodeInternalIP sql.NullString `gorm:"column:node_internal_ip;type:varchar;size:30;"`
	//[ 5] node_public_ip                                 varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	NodePublicIP sql.NullString `gorm:"column:node_public_ip;type:varchar;size:30;"`
	//[ 6] node_principal                                 varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	NodePrincipal sql.NullString `gorm:"column:node_principal;type:varchar;size:30;"`
	//[ 7] node_is_local                                  varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	NodeIsLocal sql.NullString `gorm:"column:node_is_local;type:varchar;size:20;"`
	//[ 8] serial_nnum                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	SerialNnum sql.NullString `gorm:"column:serial_nnum;type:varchar;size:30;"`
	//[ 9] node_condition                                 varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	NodeCondition sql.NullString `gorm:"column:node_condition;type:varchar;size:30;"`
	//[10] node_status                                    varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	NodeStatus sql.NullString `gorm:"column:node_status;type:varchar;size:60;"`
	//[11] firmware_version                               varchar(128)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 128     default: []
	FirmwareVersion sql.NullString `gorm:"column:firmware_version;type:varchar;size:128;"`
	//[12] node_mac                                       varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	NodeMac sql.NullString `gorm:"column:node_mac;type:varchar;size:30;"`
	//[13] node_switch_type                               varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	NodeSwitchType sql.NullString `gorm:"column:node_switch_type;type:varchar;size:30;"`
}

var mct_cluster_detailsTableInfo = &TableInfo{
	Name: "mct_cluster_details",
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
			Name:               "node_id",
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
			GoFieldName:        "NodeID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "node_id",
			ProtobufFieldName:  "node_id",
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
			Name:               "principal_switch_mac",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "PrincipalSwitchMac",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "principal_switch_mac",
			ProtobufFieldName:  "principal_switch_mac",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "node_internal_ip",
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
			GoFieldName:        "NodeInternalIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "node_internal_ip",
			ProtobufFieldName:  "node_internal_ip",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "node_public_ip",
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
			GoFieldName:        "NodePublicIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "node_public_ip",
			ProtobufFieldName:  "node_public_ip",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "node_principal",
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
			GoFieldName:        "NodePrincipal",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "node_principal",
			ProtobufFieldName:  "node_principal",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "node_is_local",
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
			GoFieldName:        "NodeIsLocal",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "node_is_local",
			ProtobufFieldName:  "node_is_local",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "serial_nnum",
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
			GoFieldName:        "SerialNnum",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "serial_nnum",
			ProtobufFieldName:  "serial_nnum",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "node_condition",
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
			GoFieldName:        "NodeCondition",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "node_condition",
			ProtobufFieldName:  "node_condition",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "node_status",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "NodeStatus",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "node_status",
			ProtobufFieldName:  "node_status",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "firmware_version",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "FirmwareVersion",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "firmware_version",
			ProtobufFieldName:  "firmware_version",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "node_mac",
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
			GoFieldName:        "NodeMac",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "node_mac",
			ProtobufFieldName:  "node_mac",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		{
			Index:              13,
			Name:               "node_switch_type",
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
			GoFieldName:        "NodeSwitchType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "node_switch_type",
			ProtobufFieldName:  "node_switch_type",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MctClusterDetails) TableName() string {
	return "mct_cluster_details"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MctClusterDetails) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MctClusterDetails) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MctClusterDetails) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MctClusterDetails) TableInfo() *TableInfo {
	return mct_cluster_detailsTableInfo
}
