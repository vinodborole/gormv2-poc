package model

import "fmt"

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// RetrieveOne action when a record is retrieved from db
	RetrieveOne = Action(1)

	// RetrieveMany action when record(s) are retrieved from db
	RetrieveMany = Action(2)

	// Update action when record is updated in db
	Update = Action(3)

	// Delete action when record is deleted in db
	Delete = Action(4)

	// FetchDDL action when fetching ddl info from db
	FetchDDL = Action(5)

	tables map[string]*TableInfo
)

func init() {
	tables = make(map[string]*TableInfo)

	tables["apps"] = appsTableInfo
	tables["asn_allocation_pools"] = asn_allocation_poolsTableInfo
	tables["bgp_af_evpn_neighbors"] = bgp_af_evpn_neighborsTableInfo
	tables["bgp_af_ip_neighbors"] = bgp_af_ip_neighborsTableInfo
	tables["bgp_evpn_afs"] = bgp_evpn_afsTableInfo
	tables["bgp_ip_afs"] = bgp_ip_afsTableInfo
	tables["cluster_members"] = cluster_membersTableInfo
	tables["devices"] = devicesTableInfo
	tables["event_historys"] = event_historysTableInfo
	tables["evpns"] = evpnsTableInfo
	tables["execution_logs"] = execution_logsTableInfo
	tables["fabric_errors"] = fabric_errorsTableInfo
	tables["fabric_properties"] = fabric_propertiesTableInfo
	tables["fabrics"] = fabricsTableInfo
	tables["interface_switch_configs"] = interface_switch_configsTableInfo
	tables["ip_allocation_pools"] = ip_allocation_poolsTableInfo
	tables["ip_pair_allocation_pools"] = ip_pair_allocation_poolsTableInfo
	tables["ip_prefix_lists"] = ip_prefix_listsTableInfo
	tables["lldp_neighbors"] = lldp_neighborsTableInfo
	tables["mct_cluster_configs"] = mct_cluster_configsTableInfo
	tables["mct_cluster_details"] = mct_cluster_detailsTableInfo
	tables["overlay_gateways"] = overlay_gatewaysTableInfo
	tables["peer_groups"] = peer_groupsTableInfo
	tables["phys_interfaces"] = phys_interfacesTableInfo
	tables["remote_neighbor_switch_configs"] = remote_neighbor_switch_configsTableInfo
	tables["router_bgps"] = router_bgpsTableInfo
	tables["router_pims"] = router_pimsTableInfo
	tables["switch_configs"] = switch_configsTableInfo
	tables["used_asns"] = used_asnsTableInfo
	tables["used_ip_pairs"] = used_ip_pairsTableInfo
	tables["used_ips"] = used_ipsTableInfo
}

// String describe the action
func (i Action) String() string {
	switch i {
	case Create:
		return "Create"
	case RetrieveOne:
		return "RetrieveOne"
	case RetrieveMany:
		return "RetrieveMany"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case FetchDDL:
		return "FetchDDL"
	default:
		return fmt.Sprintf("unknown action: %d", int(i))
	}
}

// Model interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave() error
	Prepare()
	Validate(action Action) error
	TableInfo() *TableInfo
}

// TableInfo describes a table in the database
type TableInfo struct {
	Name    string        `json:"name"`
	Columns []*ColumnInfo `json:"columns"`
}

// ColumnInfo describes a column in the database table
type ColumnInfo struct {
	Index              int    `json:"index"`
	GoFieldName        string `json:"go_field_name"`
	GoFieldType        string `json:"go_field_type"`
	JSONFieldName      string `json:"json_field_name"`
	ProtobufFieldName  string `json:"protobuf_field_name"`
	ProtobufType       string `json:"protobuf_field_type"`
	ProtobufPos        int    `json:"protobuf_field_pos"`
	Comment            string `json:"comment"`
	Notes              string `json:"notes"`
	Name               string `json:"name"`
	Nullable           bool   `json:"is_nullable"`
	DatabaseTypeName   string `json:"database_type_name"`
	DatabaseTypePretty string `json:"database_type_pretty"`
	IsPrimaryKey       bool   `json:"is_primary_key"`
	IsAutoIncrement    bool   `json:"is_auto_increment"`
	IsArray            bool   `json:"is_array"`
	ColumnType         string `json:"column_type"`
	ColumnLength       int64  `json:"column_length"`
	DefaultValue       string `json:"default_value"`
}

// GetTableInfo retrieve TableInfo for a table
func GetTableInfo(name string) (*TableInfo, bool) {
	val, ok := tables[name]
	return val, ok
}
