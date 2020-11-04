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


CREATE TABLE `mct_cluster_configs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `cluster_id` int DEFAULT NULL,
  `fabric_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `mct_neighbor_device_id` int DEFAULT NULL,
  `device_one_mgmt_ip` varchar(30) DEFAULT NULL,
  `device_two_mgmt_ip` varchar(30) DEFAULT NULL,
  `peer_interface_name` varchar(30) DEFAULT NULL,
  `peer_interfacetype` varchar(30) DEFAULT NULL,
  `peer_interface_speed` varchar(30) DEFAULT NULL,
  `peer_interface_two_speed` varchar(30) DEFAULT NULL,
  `peer_one_ip` varchar(30) DEFAULT NULL,
  `peer_two_ip` varchar(30) DEFAULT NULL,
  `control_vlan` varchar(30) DEFAULT NULL,
  `control_ve` varchar(30) DEFAULT NULL,
  `ve_interface_one_id` int DEFAULT NULL,
  `ve_interface_two_id` int DEFAULT NULL,
  `config_type` varchar(30) DEFAULT NULL,
  `local_node_id` varchar(30) DEFAULT NULL,
  `remote_node_id` varchar(30) DEFAULT NULL,
  `updated_attributes` bigint DEFAULT NULL,
  `sw_cluster_name` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `mct_cluster_configs_fabric_id_fkey` (`fabric_id`),
  KEY `mct_cluster_configs_device_id_fkey` (`device_id`),
  KEY `mct_cluster_configs_mct_neighbor_device_id_fkey` (`mct_neighbor_device_id`),
  CONSTRAINT `mct_cluster_configs_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `mct_cluster_configs_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE,
  CONSTRAINT `mct_cluster_configs_mct_neighbor_device_id_fkey` FOREIGN KEY (`mct_neighbor_device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "cluster_id": 16,    "mct_neighbor_device_id": 44,    "peer_interface_name": "TOETLvGWxJyZetiutrVhIFdCF",    "peer_two_ip": "PYPReCGjHmPxHZEvPgxWNbLiv",    "control_vlan": "GykbFIuJOayjkcEKbttNuZgBE",    "ve_interface_one_id": 89,    "remote_node_id": "CfljiGDoXtvfJroVFwilumrcn",    "sw_cluster_name": "DltbcZsnsBghAEdDkhmAgrrBV",    "peer_interface_speed": "OKhDKKRRUUZZOvfVtWkgmPsWn",    "control_ve": "ghpHPrcoptQnMchQQIgSEVXQs",    "config_type": "mKIDgDlauIBNCTmcECCotUvfO",    "updated_attributes": 99,    "device_one_mgmt_ip": "IPFjnOpVpXPyeIEMXXkEhFWhX",    "device_two_mgmt_ip": "qCstEisWiXdrmoWLviWkPUZkw",    "peer_interfacetype": "xOoyvNmbxOVrTObCqWDSIIdkD",    "peer_interface_two_speed": "pXLeeLDLtpgFMHVcsHtJHDyZs",    "peer_one_ip": "hqNlYVfcKAKXVyZjxDYDvJPAa",    "id": 53,    "fabric_id": 42,    "device_id": 69,    "ve_interface_two_id": 33,    "local_node_id": "nCspTEIGEQDnYMcsAGhepfitJ"}



*/

// MctClusterConfigs struct is a row record of the mct_cluster_configs table in the myapp database
type MctClusterConfigs struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] cluster_id                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ClusterID sql.NullInt64 `gorm:"column:cluster_id;type:int;"`
	//[ 2] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 3] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 4] mct_neighbor_device_id                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	MctNeighborDeviceID sql.NullInt64 `gorm:"column:mct_neighbor_device_id;type:int;"`
	//[ 5] device_one_mgmt_ip                             varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceOneMgmtIP sql.NullString `gorm:"column:device_one_mgmt_ip;type:varchar;size:30;"`
	//[ 6] device_two_mgmt_ip                             varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceTwoMgmtIP sql.NullString `gorm:"column:device_two_mgmt_ip;type:varchar;size:30;"`
	//[ 7] peer_interface_name                            varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	PeerInterfaceName sql.NullString `gorm:"column:peer_interface_name;type:varchar;size:30;"`
	//[ 8] peer_interfacetype                             varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	PeerInterfacetype sql.NullString `gorm:"column:peer_interfacetype;type:varchar;size:30;"`
	//[ 9] peer_interface_speed                           varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	PeerInterfaceSpeed sql.NullString `gorm:"column:peer_interface_speed;type:varchar;size:30;"`
	//[10] peer_interface_two_speed                       varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	PeerInterfaceTwoSpeed sql.NullString `gorm:"column:peer_interface_two_speed;type:varchar;size:30;"`
	//[11] peer_one_ip                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	PeerOneIP sql.NullString `gorm:"column:peer_one_ip;type:varchar;size:30;"`
	//[12] peer_two_ip                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	PeerTwoIP sql.NullString `gorm:"column:peer_two_ip;type:varchar;size:30;"`
	//[13] control_vlan                                   varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ControlVlan sql.NullString `gorm:"column:control_vlan;type:varchar;size:30;"`
	//[14] control_ve                                     varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ControlVe sql.NullString `gorm:"column:control_ve;type:varchar;size:30;"`
	//[15] ve_interface_one_id                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	VeInterfaceOneID sql.NullInt64 `gorm:"column:ve_interface_one_id;type:int;"`
	//[16] ve_interface_two_id                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	VeInterfaceTwoID sql.NullInt64 `gorm:"column:ve_interface_two_id;type:int;"`
	//[17] config_type                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType sql.NullString `gorm:"column:config_type;type:varchar;size:30;"`
	//[18] local_node_id                                  varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	LocalNodeID sql.NullString `gorm:"column:local_node_id;type:varchar;size:30;"`
	//[19] remote_node_id                                 varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	RemoteNodeID sql.NullString `gorm:"column:remote_node_id;type:varchar;size:30;"`
	//[20] updated_attributes                             bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	UpdatedAttributes sql.NullInt64 `gorm:"column:updated_attributes;type:bigint;"`
	//[21] sw_cluster_name                                text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	SwClusterName string `gorm:"column:sw_cluster_name;type:text;size:65535;"`
}

var mct_cluster_configsTableInfo = &TableInfo{
	Name: "mct_cluster_configs",
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
			Name:               "cluster_id",
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
			GoFieldName:        "ClusterID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "cluster_id",
			ProtobufFieldName:  "cluster_id",
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
			Name:               "mct_neighbor_device_id",
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
			GoFieldName:        "MctNeighborDeviceID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "mct_neighbor_device_id",
			ProtobufFieldName:  "mct_neighbor_device_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "device_one_mgmt_ip",
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
			GoFieldName:        "DeviceOneMgmtIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "device_one_mgmt_ip",
			ProtobufFieldName:  "device_one_mgmt_ip",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "device_two_mgmt_ip",
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
			GoFieldName:        "DeviceTwoMgmtIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "device_two_mgmt_ip",
			ProtobufFieldName:  "device_two_mgmt_ip",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "peer_interface_name",
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
			GoFieldName:        "PeerInterfaceName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "peer_interface_name",
			ProtobufFieldName:  "peer_interface_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "peer_interfacetype",
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
			GoFieldName:        "PeerInterfacetype",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "peer_interfacetype",
			ProtobufFieldName:  "peer_interfacetype",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "peer_interface_speed",
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
			GoFieldName:        "PeerInterfaceSpeed",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "peer_interface_speed",
			ProtobufFieldName:  "peer_interface_speed",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "peer_interface_two_speed",
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
			GoFieldName:        "PeerInterfaceTwoSpeed",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "peer_interface_two_speed",
			ProtobufFieldName:  "peer_interface_two_speed",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "peer_one_ip",
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
			GoFieldName:        "PeerOneIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "peer_one_ip",
			ProtobufFieldName:  "peer_one_ip",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "peer_two_ip",
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
			GoFieldName:        "PeerTwoIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "peer_two_ip",
			ProtobufFieldName:  "peer_two_ip",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		{
			Index:              13,
			Name:               "control_vlan",
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
			GoFieldName:        "ControlVlan",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "control_vlan",
			ProtobufFieldName:  "control_vlan",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		{
			Index:              14,
			Name:               "control_ve",
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
			GoFieldName:        "ControlVe",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "control_ve",
			ProtobufFieldName:  "control_ve",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		{
			Index:              15,
			Name:               "ve_interface_one_id",
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
			GoFieldName:        "VeInterfaceOneID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "ve_interface_one_id",
			ProtobufFieldName:  "ve_interface_one_id",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		{
			Index:              16,
			Name:               "ve_interface_two_id",
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
			GoFieldName:        "VeInterfaceTwoID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "ve_interface_two_id",
			ProtobufFieldName:  "ve_interface_two_id",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		{
			Index:              17,
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
			ProtobufPos:        18,
		},

		{
			Index:              18,
			Name:               "local_node_id",
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
			GoFieldName:        "LocalNodeID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "local_node_id",
			ProtobufFieldName:  "local_node_id",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		{
			Index:              19,
			Name:               "remote_node_id",
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
			GoFieldName:        "RemoteNodeID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "remote_node_id",
			ProtobufFieldName:  "remote_node_id",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		{
			Index:              20,
			Name:               "updated_attributes",
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
			GoFieldName:        "UpdatedAttributes",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "updated_attributes",
			ProtobufFieldName:  "updated_attributes",
			ProtobufType:       "int64",
			ProtobufPos:        21,
		},

		{
			Index:              21,
			Name:               "sw_cluster_name",
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
			GoFieldName:        "SwClusterName",
			GoFieldType:        "string",
			JSONFieldName:      "sw_cluster_name",
			ProtobufFieldName:  "sw_cluster_name",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MctClusterConfigs) TableName() string {
	return "mct_cluster_configs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MctClusterConfigs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MctClusterConfigs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MctClusterConfigs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MctClusterConfigs) TableInfo() *TableInfo {
	return mct_cluster_configsTableInfo
}
