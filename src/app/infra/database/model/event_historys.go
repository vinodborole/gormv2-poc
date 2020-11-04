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


CREATE TABLE `event_historys` (
  `id` int NOT NULL AUTO_INCREMENT,
  `execution_uuid` varchar(255) DEFAULT NULL,
  `date` varchar(100) DEFAULT NULL,
  `service` varchar(255) DEFAULT NULL,
  `event` varchar(255) DEFAULT NULL,
  `device` varchar(255) DEFAULT NULL,
  `message_type` varchar(255) DEFAULT NULL,
  `message_object` text,
  `message` text,
  `error_text` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `execution_uuid` (`execution_uuid`),
  CONSTRAINT `event_historys_ibfk_1` FOREIGN KEY (`execution_uuid`) REFERENCES `execution_logs` (`uuid`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "date": "doLacEtutnGXgVwirKVeFuBLh",    "service": "ZnIPnxUlmyXLEbJqYgovtKPbQ",    "message_object": "NnjoLrOQnfwMOLumPtefNCJjr",    "message": "CpaVfkJEQsuCAclhXJodgHtJr",    "created_at": "2031-07-24T01:09:27.407287329+05:30",    "id": 2,    "execution_uuid": "ExOXxypjmRcvJiNBINxFheNPx",    "event": "dFoRyhPnYWshpKsMtXGtWFGRn",    "device": "ZdCxIWTcYsALVxNAVnsYYTTQi",    "message_type": "aHonKBgMepbaBVTkBmVnlaSZT",    "error_text": "DtseQafjtRiPdoZvsOSTUdETn"}



*/

// EventHistorys struct is a row record of the event_historys table in the myapp database
type EventHistorys struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;"`
	//[ 1] execution_uuid                                 varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ExecutionUUID sql.NullString `gorm:"column:execution_uuid;type:varchar;size:255;"`
	//[ 2] date                                           varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Date sql.NullString `gorm:"column:date;type:varchar;size:100;"`
	//[ 3] service                                        varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Service sql.NullString `gorm:"column:service;type:varchar;size:255;"`
	//[ 4] event                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Event sql.NullString `gorm:"column:event;type:varchar;size:255;"`
	//[ 5] device                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Device sql.NullString `gorm:"column:device;type:varchar;size:255;"`
	//[ 6] message_type                                   varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	MessageType sql.NullString `gorm:"column:message_type;type:varchar;size:255;"`
	//[ 7] message_object                                 text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	MessageObject sql.NullString `gorm:"column:message_object;type:text;size:65535;"`
	//[ 8] message                                        text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Message sql.NullString `gorm:"column:message;type:text;size:65535;"`
	//[ 9] error_text                                     text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	ErrorText sql.NullString `gorm:"column:error_text;type:text;size:65535;"`
	//[10] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;"`
}

var event_historysTableInfo = &TableInfo{
	Name: "event_historys",
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
			Name:               "execution_uuid",
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
			GoFieldName:        "ExecutionUUID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "execution_uuid",
			ProtobufFieldName:  "execution_uuid",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "date",
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
			GoFieldName:        "Date",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "date",
			ProtobufFieldName:  "date",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "service",
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
			GoFieldName:        "Service",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "service",
			ProtobufFieldName:  "service",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "event",
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
			GoFieldName:        "Event",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "event",
			ProtobufFieldName:  "event",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "device",
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
			GoFieldName:        "Device",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "device",
			ProtobufFieldName:  "device",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "message_type",
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
			GoFieldName:        "MessageType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "message_type",
			ProtobufFieldName:  "message_type",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "message_object",
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
			GoFieldName:        "MessageObject",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "message_object",
			ProtobufFieldName:  "message_object",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "message",
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
			GoFieldName:        "Message",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "message",
			ProtobufFieldName:  "message",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "error_text",
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
			GoFieldName:        "ErrorText",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "error_text",
			ProtobufFieldName:  "error_text",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		{
			Index:              10,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *EventHistorys) TableName() string {
	return "event_historys"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *EventHistorys) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *EventHistorys) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *EventHistorys) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *EventHistorys) TableInfo() *TableInfo {
	return event_historysTableInfo
}
