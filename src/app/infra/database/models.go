package database

/*
//App struct
type App struct {
	gorm.Model

	Name        string
	Description string
	Url         string
	Port        string
}

*/
/*
type Apps struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;"`
	//[ 1] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	//[ 2] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
	//[ 3] deleted_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime;"`
	//[ 4] name                                           text(4294967295)     null: true   primary: false  isArray: false  auto: false  col: text            len: 4294967295 default: []
	Name sql.NullString `gorm:"column:name;type:text;size:4294967295;"`
	//[ 5] description                                    text(4294967295)     null: true   primary: false  isArray: false  auto: false  col: text            len: 4294967295 default: []
	Description sql.NullString `gorm:"column:description;type:text;size:4294967295;"`
	//[ 6] url                                            text(4294967295)     null: true   primary: false  isArray: false  auto: false  col: text            len: 4294967295 default: []
	URL sql.NullString `gorm:"column:url;type:text;size:4294967295;"`
	//[ 7] port                                           text(4294967295)     null: true   primary: false  isArray: false  auto: false  col: text            len: 4294967295 default: []
	Port sql.NullString `gorm:"column:port;type:text;size:4294967295;"`
}

type User struct {
	gorm.Model

	ID       int32
	Username string
	Password string
}
*/
