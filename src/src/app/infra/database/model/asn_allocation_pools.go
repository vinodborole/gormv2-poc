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


CREATE TABLE `asn_allocation_pools` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `asn` bigint DEFAULT NULL,
  `device_role` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `asn_allocation_pools_fabric_id_fkey` (`fabric_id`),
  CONSTRAINT `asn_allocation_pools_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "device_role": "tkWyGhymiZCOjcOaoLpGrxBay",    "id": 97,    "fabric_id": 55,    "asn": 24}



*/

// AsnAllocationPools struct is a row record of the asn_allocation_pools table in the myapp database
type AsnAllocationPools struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] asn                                            bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	Asn sql.NullInt64 `gorm:"column:asn;type:bigint;"`
	//[ 3] device_role                                    varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DeviceRole sql.NullString `gorm:"column:device_role;type:varchar;size:30;"`
}

var asn_allocation_poolsTableInfo = &TableInfo{
	Name: "asn_allocation_pools",
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
			Name:               "asn",
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
			GoFieldName:        "Asn",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "asn",
			ProtobufFieldName:  "asn",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		{
			Index:              3,
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
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AsnAllocationPools) TableName() string {
	return "asn_allocation_pools"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AsnAllocationPools) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AsnAllocationPools) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AsnAllocationPools) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AsnAllocationPools) TableInfo() *TableInfo {
	return asn_allocation_poolsTableInfo
}
