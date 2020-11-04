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


CREATE TABLE `ip_pair_allocation_pools` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fabric_id` int DEFAULT NULL,
  `ip_address_one` varchar(30) DEFAULT NULL,
  `ip_address_two` varchar(30) DEFAULT NULL,
  `ip_type` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ip_pair_allocation_pools_fabric_id_fkey` (`fabric_id`),
  CONSTRAINT `ip_pair_allocation_pools_fabric_id_fkey` FOREIGN KEY (`fabric_id`) REFERENCES `fabrics` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 76,    "fabric_id": 43,    "ip_address_one": "xhmXLYpLAgfBkZBtDDrjlEMLM",    "ip_address_two": "ligKoooMVplZSbIbxqwDcbAkp",    "ip_type": "snLiYvkrDntfrUHCFqHByccNC"}



*/

// IPPairAllocationPools struct is a row record of the ip_pair_allocation_pools table in the myapp database
type IPPairAllocationPools struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] fabric_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FabricID sql.NullInt64 `gorm:"column:fabric_id;type:int;"`
	//[ 2] ip_address_one                                 varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IPAddressOne sql.NullString `gorm:"column:ip_address_one;type:varchar;size:30;"`
	//[ 3] ip_address_two                                 varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IPAddressTwo sql.NullString `gorm:"column:ip_address_two;type:varchar;size:30;"`
	//[ 4] ip_type                                        varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	IPType sql.NullString `gorm:"column:ip_type;type:varchar;size:30;"`
}

var ip_pair_allocation_poolsTableInfo = &TableInfo{
	Name: "ip_pair_allocation_pools",
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
			Name:               "ip_address_one",
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
			GoFieldName:        "IPAddressOne",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ip_address_one",
			ProtobufFieldName:  "ip_address_one",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "ip_address_two",
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
			GoFieldName:        "IPAddressTwo",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ip_address_two",
			ProtobufFieldName:  "ip_address_two",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "ip_type",
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
			GoFieldName:        "IPType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ip_type",
			ProtobufFieldName:  "ip_type",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *IPPairAllocationPools) TableName() string {
	return "ip_pair_allocation_pools"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *IPPairAllocationPools) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *IPPairAllocationPools) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *IPPairAllocationPools) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *IPPairAllocationPools) TableInfo() *TableInfo {
	return ip_pair_allocation_poolsTableInfo
}
