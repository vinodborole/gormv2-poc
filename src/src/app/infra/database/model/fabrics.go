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


CREATE TABLE `fabrics` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `fabric_type` varchar(10) DEFAULT NULL,
  `stage` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "name": "pGeeEOhcJwDgXVQkVRkRaQame",    "description": "nurfTjPuUiIuseFxMjDReuGsk",    "fabric_type": "vKRORZUggZnICncvdcfgRdUGg",    "stage": 30,    "id": 34}



*/

// Fabrics struct is a row record of the fabrics table in the myapp database
type Fabrics struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] name                                           varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Name sql.NullString `gorm:"column:name;type:varchar;size:30;"`
	//[ 2] description                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Description sql.NullString `gorm:"column:description;type:varchar;size:255;"`
	//[ 3] fabric_type                                    varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	FabricType sql.NullString `gorm:"column:fabric_type;type:varchar;size:10;"`
	//[ 4] stage                                          int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Stage sql.NullInt64 `gorm:"column:stage;type:int;"`
}

var fabricsTableInfo = &TableInfo{
	Name: "fabrics",
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
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "description",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Description",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "fabric_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "FabricType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "fabric_type",
			ProtobufFieldName:  "fabric_type",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "stage",
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
			GoFieldName:        "Stage",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "stage",
			ProtobufFieldName:  "stage",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *Fabrics) TableName() string {
	return "fabrics"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *Fabrics) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *Fabrics) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *Fabrics) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *Fabrics) TableInfo() *TableInfo {
	return fabricsTableInfo
}
