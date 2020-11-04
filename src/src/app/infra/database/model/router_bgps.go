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


CREATE TABLE `router_bgps` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `device_id` int DEFAULT NULL,
  `local_asn` bigint DEFAULT NULL,
  `router_id` varchar(30) DEFAULT NULL,
  `bfd_enable` tinyint(1) DEFAULT NULL,
  `bfd_tx` int DEFAULT NULL,
  `bfd_rx` int DEFAULT NULL,
  `bfd_multiplier` int DEFAULT NULL,
  `network_address` varchar(30) DEFAULT NULL,
  `redistribute_type` varchar(30) DEFAULT NULL,
  `as4_enable` tinyint(1) DEFAULT NULL,
  `fast_fallover` tinyint(1) DEFAULT NULL,
  `config_type` varchar(30) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `router_bgp_fabric_id_fkey` (`fabric_id`),
  KEY `router_bgp_device_id_fkey` (`device_id`),
  CONSTRAINT `router_bgp_device_id_fkey` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`) ON DELETE CASCADE,
  CONSTRAINT `router_bgp_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "bfd_rx": 43,    "network_address": "RHfGblLsUgvrqOgqgBMhhJuFt",    "fabric_id": 77,    "bfd_multiplier": 87,    "device_id": 63,    "bfd_enable": 92,    "redistribute_type": "ZOpOkMEQHxjRWmDYeHovfnKWn",    "as_4_enable": 51,    "config_type": "oBljHZlSNfsZvArjFvglDUjEI",    "id": 60,    "local_asn": 5,    "router_id": "hQTAJCHHOxnbpZNLbuHmuiXOp",    "bfd_tx": 38,    "fast_fallover": 97}



*/

// RouterBgps struct is a row record of the router_bgps table in the myapp database
type RouterBgps struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] device_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DeviceID sql.NullInt64 `gorm:"column:device_id;type:int;"`
	//[ 3] local_asn                                      bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	LocalAsn sql.NullInt64 `gorm:"column:local_asn;type:bigint;"`
	//[ 4] router_id                                      varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	RouterID sql.NullString `gorm:"column:router_id;type:varchar;size:30;"`
	//[ 5] bfd_enable                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	BfdEnable sql.NullInt64 `gorm:"column:bfd_enable;type:tinyint;"`
	//[ 6] bfd_tx                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BfdTx sql.NullInt64 `gorm:"column:bfd_tx;type:int;"`
	//[ 7] bfd_rx                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BfdRx sql.NullInt64 `gorm:"column:bfd_rx;type:int;"`
	//[ 8] bfd_multiplier                                 int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BfdMultiplier sql.NullInt64 `gorm:"column:bfd_multiplier;type:int;"`
	//[ 9] network_address                                varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	NetworkAddress sql.NullString `gorm:"column:network_address;type:varchar;size:30;"`
	//[10] redistribute_type                              varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	RedistributeType sql.NullString `gorm:"column:redistribute_type;type:varchar;size:30;"`
	//[11] as4_enable                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	As4Enable sql.NullInt64 `gorm:"column:as4_enable;type:tinyint;"`
	//[12] fast_fallover                                  tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	FastFallover sql.NullInt64 `gorm:"column:fast_fallover;type:tinyint;"`
	//[13] config_type                                    varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConfigType string `gorm:"column:config_type;type:varchar;size:30;"`
}

var router_bgpsTableInfo = &TableInfo{
	Name: "router_bgps",
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
			Name:               "local_asn",
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
			GoFieldName:        "LocalAsn",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "local_asn",
			ProtobufFieldName:  "local_asn",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "router_id",
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
			GoFieldName:        "RouterID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "router_id",
			ProtobufFieldName:  "router_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "bfd_enable",
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
			GoFieldName:        "BfdEnable",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bfd_enable",
			ProtobufFieldName:  "bfd_enable",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "bfd_tx",
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
			GoFieldName:        "BfdTx",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bfd_tx",
			ProtobufFieldName:  "bfd_tx",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "bfd_rx",
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
			GoFieldName:        "BfdRx",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bfd_rx",
			ProtobufFieldName:  "bfd_rx",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "bfd_multiplier",
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
			GoFieldName:        "BfdMultiplier",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bfd_multiplier",
			ProtobufFieldName:  "bfd_multiplier",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "network_address",
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
			GoFieldName:        "NetworkAddress",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "network_address",
			ProtobufFieldName:  "network_address",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "redistribute_type",
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
			GoFieldName:        "RedistributeType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "redistribute_type",
			ProtobufFieldName:  "redistribute_type",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "as4_enable",
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
			GoFieldName:        "As4Enable",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "as_4_enable",
			ProtobufFieldName:  "as_4_enable",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "fast_fallover",
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
			GoFieldName:        "FastFallover",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "fast_fallover",
			ProtobufFieldName:  "fast_fallover",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		{
			Index:              13,
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
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RouterBgps) TableName() string {
	return "router_bgps"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RouterBgps) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RouterBgps) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RouterBgps) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RouterBgps) TableInfo() *TableInfo {
	return router_bgpsTableInfo
}
