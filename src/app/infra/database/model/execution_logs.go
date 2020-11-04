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


CREATE TABLE `execution_logs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) DEFAULT NULL,
  `command` varchar(255) DEFAULT NULL,
  `params` text,
  `status` varchar(100) DEFAULT NULL,
  `start_time` varchar(100) DEFAULT NULL,
  `end_time` varchar(100) DEFAULT NULL,
  `duration` varchar(100) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_name` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UUID_UNIQUE` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "uuid": "CAQqjgoqkDSEgNNQrnTKCpXAH",    "command": "URoDeryVDidPxMaliKpLZvmVG",    "start_time": "QkGSHflXsckAhukEQprjBEIMa",    "user_name": "vLkrJbAlXujAWrwDXrfBUdbVH",    "id": 10,    "params": "HbfHNoVfwfOOaMYlcRBSSOFpL",    "status": "tHkUnbcdSCcydiPgFiLQlakso",    "end_time": "bsqvoohamPsJNHscMjtEGDgAC",    "duration": "KviIfJJKXiEQWngeHvkQuPEra",    "created_at": "2196-05-19T12:49:30.351556379+05:30"}



*/

// ExecutionLogs struct is a row record of the execution_logs table in the myapp database
type ExecutionLogs struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] uuid                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	UUID sql.NullString `gorm:"column:uuid;type:varchar;size:255;"`
	//[ 2] command                                        varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Command sql.NullString `gorm:"column:command;type:varchar;size:255;"`
	//[ 3] params                                         text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Params sql.NullString `gorm:"column:params;type:text;size:65535;"`
	//[ 4] status                                         varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Status sql.NullString `gorm:"column:status;type:varchar;size:100;"`
	//[ 5] start_time                                     varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	StartTime sql.NullString `gorm:"column:start_time;type:varchar;size:100;"`
	//[ 6] end_time                                       varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	EndTime sql.NullString `gorm:"column:end_time;type:varchar;size:100;"`
	//[ 7] duration                                       varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Duration sql.NullString `gorm:"column:duration;type:varchar;size:100;"`
	//[ 8] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;"`
	//[ 9] user_name                                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	UserName sql.NullString `gorm:"column:user_name;type:varchar;size:100;"`
}

var execution_logsTableInfo = &TableInfo{
	Name: "execution_logs",
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
			Name:               "uuid",
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
			GoFieldName:        "UUID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "uuid",
			ProtobufFieldName:  "uuid",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "command",
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
			GoFieldName:        "Command",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "command",
			ProtobufFieldName:  "command",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "params",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Params",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "params",
			ProtobufFieldName:  "params",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "status",
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
			GoFieldName:        "Status",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "start_time",
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
			GoFieldName:        "StartTime",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "end_time",
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
			GoFieldName:        "EndTime",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "duration",
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
			GoFieldName:        "Duration",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "duration",
			ProtobufFieldName:  "duration",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "user_name",
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
			GoFieldName:        "UserName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "user_name",
			ProtobufFieldName:  "user_name",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *ExecutionLogs) TableName() string {
	return "execution_logs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *ExecutionLogs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *ExecutionLogs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *ExecutionLogs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *ExecutionLogs) TableInfo() *TableInfo {
	return execution_logsTableInfo
}
