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


CREATE TABLE `remote_neighbor_switch_configs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `remote_interface_id` int DEFAULT NULL,
  `remote_device_id` int DEFAULT NULL,
  `peer_group_id` int DEFAULT NULL,
  `encapsulation_type` varchar(30) DEFAULT NULL,
  `remote_ip_address` varchar(30) DEFAULT NULL,
  `type` varchar(30) DEFAULT NULL,
  `remote_as` varchar(30) DEFAULT NULL,
  `config_type` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `remote_neighbor_switch_configs_fabric_id_fkey` (`fabric_id`),
  KEY `remote_neighbor_switch_configs_device_id_fkey` (`device_id`),
  KEY `remote_neighbor_switch_configs_remote_interface_id_fkey` (`remote_interface_id`),
  KEY `remote_neighbor_switch_configs_remote_device_id_fkey` (`remote_device_id`),
  KEY `remote_neighbor_switch_configs_peer_group_id_fkey` (`peer_group_id`),
  CONSTRAINT `remote_neighbor_switch_configs_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `remote_neighbor_switch_configs_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE,
  CONSTRAINT `remote_neighbor_switch_configs_peer_group_id_fkey` FOREIGN KEY (`peer_group_id`) REFERENCES `peer_groups` (`id`) ON DELETE CASCADE,
  CONSTRAINT `remote_neighbor_switch_configs_remote_device_id_fkey` FOREIGN KEY (`remote_device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `remote_neighbor_switch_configs_remote_interface_id_fkey` FOREIGN KEY (`remote_interface_id`) REFERENCES `phys_interfaces` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "remote_as": "YTOXcXyCfSSDQtQpGBYmsEMAN",    "fabric_id": 64,    "remote_interface_id": 6,    "remote_device_id": 37,    "remote_ip_address": "QAcyplHRZUWhIBFYmoNhOqRlJ",    "type": "NTMapaWfPkPnLErsvrJuEqJKu",    "id": 77,    "device_id": 66,    "peer_group_id": 24,    "encapsulation_type": "JsdGXivEiCZRNfuCUEauXZKrO",    "config_type": "vGuhJROTROdFNSptFNDfekrfR"}



*/

// RemoteNeighborSwitchConfigs struct is a row record of the remote_neighbor_switch_configs table in the myapp database
type RemoteNeighborSwitchConfigs struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 3] remote_interface_id                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RemoteInterfaceID sql.NullInt64 `gorm:"column:remote_interface_id;type:int;"`
	//[ 4] remote_device_id                               int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RemoteDeviceID sql.NullInt64 `gorm:"column:remote_device_id;type:int;"`
	//[ 5] peer_group_id                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PeerGroupID sql.NullInt64 `gorm:"column:peer_group_id;type:int;"`
	//[ 6] encapsulation_type                             varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	EncapsulationType sql.NullString `gorm:"column:encapsulation_type;type:varchar;size:30;"`
	//[ 7] remote_ip_address                              varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	RemoteIPAddress sql.NullString `gorm:"column:remote_ip_address;type:varchar;size:30;"`
	//[ 8] type                                           varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Type sql.NullString `gorm:"column:type;type:varchar;size:30;"`
	//[ 9] remote_as                                      varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	RemoteAs sql.NullString `gorm:"column:remote_as;type:varchar;size:30;"`
	//[10] config_type                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType sql.NullString `gorm:"column:config_type;type:varchar;size:30;"`
}

var remote_neighbor_switch_configsTableInfo = &TableInfo{
	Name: "remote_neighbor_switch_configs",
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
			Name:               "remote_interface_id",
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
			GoFieldName:        "RemoteInterfaceID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "remote_interface_id",
			ProtobufFieldName:  "remote_interface_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "remote_device_id",
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
			GoFieldName:        "RemoteDeviceID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "remote_device_id",
			ProtobufFieldName:  "remote_device_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "peer_group_id",
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
			GoFieldName:        "PeerGroupID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "peer_group_id",
			ProtobufFieldName:  "peer_group_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "encapsulation_type",
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
			GoFieldName:        "EncapsulationType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "encapsulation_type",
			ProtobufFieldName:  "encapsulation_type",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "remote_ip_address",
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
			GoFieldName:        "RemoteIPAddress",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "remote_ip_address",
			ProtobufFieldName:  "remote_ip_address",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "type",
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
			GoFieldName:        "Type",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "remote_as",
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
			GoFieldName:        "RemoteAs",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "remote_as",
			ProtobufFieldName:  "remote_as",
			ProtobufType:       "string",
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
	},
}

// TableName sets the insert table name for this struct type
func (r *RemoteNeighborSwitchConfigs) TableName() string {
	return "remote_neighbor_switch_configs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RemoteNeighborSwitchConfigs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RemoteNeighborSwitchConfigs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RemoteNeighborSwitchConfigs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RemoteNeighborSwitchConfigs) TableInfo() *TableInfo {
	return remote_neighbor_switch_configsTableInfo
}
