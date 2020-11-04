create table if not exists apps
(
    id integer NOT NULL AUTO_INCREMENT,
    name varchar(30),
    description varchar(255),
    url varchar(30),
    port varchar(30),
    constraint apps_pkey PRIMARY KEY (id)
);
create table if not exists fabrics
(
    id integer NOT NULL AUTO_INCREMENT,
    name varchar(30),
    description varchar(255),
    fabric_type varchar(10),
    stage integer,
    constraint fabrics_pkey PRIMARY KEY (id)
);
create table if not exists fabric_properties
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    p2_p_link_range varchar(100),
    p2_p_ip_type varchar(100),
    loop_back_ip_range varchar(100),
    loop_back_port_number varchar(100),
    super_spine_asn_block varchar(100),
    spine_asn_block varchar(100),
    leaf_asn_block varchar(100),
    border_leaf_asn_block varchar(100),
    vtep_loop_back_port_number varchar(100),
    any_cast_mac varchar(100),
    ip_v6_any_cast_mac varchar(100),
    configure_overlay_gateway varchar(10),
    vni_auto_map varchar(10),
    bfd_enable varchar(10),
    bfd_tx varchar(10),
    bfd_rx varchar(10),
    bfd_multiplier varchar(10),
    bgp_multi_hop varchar(10),
    max_paths varchar(10),
    allow_as_in varchar(10),
    mtu varchar(10),
    ip_mtu varchar(10),
    leaf_peer_group varchar(30),
    spine_peer_group varchar(30),
    super_spine_peer_group varchar(30),
    arp_aging_timeout varchar(10),
    mac_aging_timeout varchar(10),
    mac_aging_conversational_timeout varchar(10),
    mac_move_limit varchar(10),
    duplicate_mac_timer varchar(10),
    duplicate_max_timer_max_count varchar(10),
    mct_link_ip_range varchar(100),
    control_vlan varchar(10),
    control_ve varchar(10),
    mct_port_channel varchar(100),
    rack_l3_back_up_ip_range varchar(100),
    rack_asn_block varchar(100),
    rack_underlay_ebgp_peer_group varchar(30),
    rack_overlay_ebgp_peer_group varchar(30),
    rack_l3_back_up_port varchar(100),
    rack_mct_ports varchar(100),
    rack_ld_l3_back_up_port varchar(100),
    rack_ld_mct_ports varchar(100),
    max_number_of_racks_in_fabric varchar(10),
    number_of_devices_in_rack varchar(10),
    backup_routing_enable     varchar(10),
    backup_routing_ipv4_range varchar(100),
    backup_routing_ipv6_range varchar(100),
    lacp_timeout              varchar(10),
    optimized_replication_enable varchar(10) default 'No',
    mdtgroup_range            varchar(100) default '239.0.0.0/8',
    default_mdtgroup          varchar(100) default '239.1.1.1',
    constraint fabric_properties_pkey primary key(id),
    constraint fabric_properties_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade
);
create table if not exists devices
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    name varchar(30),
    ip_address varchar(30),
    device_role varchar(30),
    local_as varchar(30),
    device_type varchar(30),
    host_name varchar(100),
    asn varchar(100),
    vtep_loopback_id varchar(100),
    pod varchar(30),
    mct_role integer,
    provisioning_status integer,
    device_status integer,
    config_gen_reason bigint,
    loopback_id varchar(100),
    model varchar(100),
    firmware_version varchar(100),
    rack varchar(30),
    is_admin_state_up boolean not null default true,
    constraint devices_pkey primary key(id),
    constraint devices_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade
);
create table if not exists phys_interfaces
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_id integer,
    inventory_interface_id integer,
    int_type varchar(30),
    int_name varchar(30),
    interface_speed varchar(30),
    ip_address varchar(30),
    ip_pim_sparse boolean default false,
    identifier varchar(30),
    mac varchar(30),
    config_type varchar(30),
    constraint phys_interfaces_pkey primary key (id),
    constraint phys_interfaces_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint phys_interfaces_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade
);
create table if not exists asn_allocation_pools
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    asn bigint,
    device_role varchar(30),
    constraint asn_allocation_pools_pkey primary key(id),
    constraint asn_allocation_pools_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id)  on delete cascade

);
create table if not exists used_asns
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_id integer,
    asn bigint,
    device_role varchar(30),
    pod varchar(30),
    constraint used_asns_pkey primary key(id),
    constraint used_asns_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint used_asns_device_id_fkey foreign key (device_id)
        references devices(id)  on delete cascade
);
create table if not exists ip_allocation_pools
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,

    ip_address varchar(30),
    ip_type varchar(30),
    constraint ip_allocation_pools_pkey primary key(id),
    constraint ip_allocation_pools_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id)  on delete cascade
);
create table if not exists used_ips
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_id integer,
    ip_address varchar(30),
    ip_type varchar(30),
    interface_id integer,
    constraint used_ips_pkey primary key(id),
    constraint used_ips_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id)  on delete cascade,
    constraint used_ips_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade,
    constraint used_ips_interface_id_fkey foreign key (interface_id)
        references phys_interfaces(id) on delete cascade
);
create table if not exists ip_pair_allocation_pools
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    ip_address_one varchar(30),
    ip_address_two varchar(30),
    ip_type varchar(30),
    constraint ip_pair_allocation_pools_pkey primary key(id),
    constraint ip_pair_allocation_pools_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade
);
create table if not exists used_ip_pairs
(
    id               integer NOT NULL AUTO_INCREMENT,
    fabric_id        integer,
    device_one_id    integer,
    device_two_id    integer,
    ip_address_one   varchar(30),
    ip_address_two   varchar(30),
    ip_type          varchar(30),
    interface_one_id integer,
    interface_two_id integer,
    constraint used_ip_pairs_pkey primary key (id),
    constraint used_ip_pairs_fabric_id_fkey foreign key (fabric_id)
        references fabrics (id) on delete cascade,
    constraint used_ip_pairs_device_one_id_fkey foreign key (device_one_id)
        references devices (id) on delete cascade,
    constraint used_ip_pairs_device_two_id_fkey foreign key (device_two_id)
        references devices (id) on delete cascade,
    constraint used_ip_pairs_interface_one_id_fkey foreign key (interface_one_id)
        references phys_interfaces (id) on delete cascade,
    constraint used_ip_pairs_interface_two_id_fkey foreign key (interface_two_id)
        references phys_interfaces (id) on delete cascade
);
create table if not exists lldp_neighbors
(
    id               integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_one_id integer,
    device_two_id integer,
    interface_one_id integer,
    interface_two_id integer,
    inventory_interface_one_id integer,
    inventory_interface_two_id integer,
    device_one_role varchar(30),
    device_two_role varchar(30),
    interface_one_name varchar(30),
    interface_one_type varchar(30),
    interface_one_ip varchar(30),
    interface_two_name varchar(30),
    interface_two_type varchar(30),
    interface_one_speed varchar(30),
    interface_two_speed varchar(30),
    interface_two_ip varchar(30),
    config_type varchar(30),
    constraint lldp_neighbors_pkey primary key(id),
    constraint lldp_neighbors_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint lldp_neighbors_device_one_id_fkey foreign key (device_one_id)
        references devices(id) on delete cascade,
    constraint lldp_neighbors_device_two_id_fkey foreign key (device_two_id)
        references devices(id) on delete cascade,
    constraint lldp_neighbors_interface_one_id_fkey foreign key (interface_one_id)
        references phys_interfaces(id) on delete cascade,
    constraint lldp_neighbors_interface_two_id_fkey foreign key (interface_two_id)
        references phys_interfaces(id) on delete cascade
);
create table if not exists peer_groups
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_id integer,
    peer_group_name varchar(64) not null,
    description varchar(100),
    bfd_enable boolean,
    remote_asn bigint,
    graceful_shut_timer int,
    config_type text not null,
    constraint peer_groups_pkey primary key(id),
    constraint peer_groups_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint peer_groups_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade
);
create table if not exists switch_configs
(
    id integer NOT NULL AUTO_INCREMENT,
    device_id integer,
    fabric_id integer,
    device_ip varchar(30),
    local_as varchar(30),
    loopback_ip varchar(30),
    vtep_loopback_ip varchar(30),
    role varchar(30),
    as_config_type varchar(30),
    loopback_ip_config_type varchar(30),
    vtep_loopback_ip_config_type varchar(30),
    vtep_loop_back_id varchar(30),
    loop_back_id varchar(30),
    constraint switch_configs_pkey  primary key(id),
    constraint switch_configs_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade,
    constraint switch_configs_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade
);
create table if not exists interface_switch_configs
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_id integer,
    interface_id integer,
    int_type varchar(30),
    int_name varchar(30),
    donor_type varchar(30),
    donor_name varchar(30),
    ip_address varchar(30),
    ip_pim_sparse boolean default false,
    config_type varchar(30),
    description varchar(100),
    constraint interface_switch_configs_pkey  primary key(id),
    constraint interface_switch_configs_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint interface_switch_configs_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade,
    constraint interface_switch_configs_interface_id_fkey foreign key (interface_id)
        references phys_interfaces(id)  on delete cascade
);
create table if not exists remote_neighbor_switch_configs
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_id integer,
    remote_interface_id integer,
    remote_device_id integer,
    peer_group_id integer default null,
    encapsulation_type varchar(30),
    remote_ip_address varchar(30),
    type      varchar(30),
    remote_as varchar(30),
    config_type varchar(30),
    constraint remote_neighbor_switch_configs_pkey  primary key(id),
    constraint remote_neighbor_switch_configs_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint remote_neighbor_switch_configs_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade,
    constraint remote_neighbor_switch_configs_remote_interface_id_fkey foreign key (remote_interface_id)
        references phys_interfaces(id) on delete cascade,
    constraint remote_neighbor_switch_configs_remote_device_id_fkey foreign key (remote_device_id)
        references devices(id) on delete cascade,
    constraint remote_neighbor_switch_configs_peer_group_id_fkey foreign key (peer_group_id)
        references peer_groups(id) on delete cascade
);
create table if not exists execution_logs
(
    id integer NOT NULL AUTO_INCREMENT,
    uuid varchar(255),
    command varchar(255),
    params text,
    status varchar(100),
    start_time varchar(100),
    end_time varchar(100),
    duration varchar(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_name varchar(100),
    CONSTRAINT UUID_UNIQUE UNIQUE (uuid),
    constraint execution_logs_pkey  primary key(id)
);
create table if not exists event_historys (
                                              id integer NOT NULL AUTO_INCREMENT PRIMARY KEY,
                                              execution_uuid varchar(255),
                                              date varchar(100),
                                              service varchar(255),
                                              event varchar(255),
                                              device varchar(255),
                                              message_type varchar(255),
                                              message_object TEXT,
                                              message TEXT,
                                              error_text TEXT,
                                              created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                              FOREIGN KEY (execution_uuid) REFERENCES execution_logs(uuid) ON DELETE CASCADE
);
create table if not exists mct_cluster_details
(
    id integer NOT NULL AUTO_INCREMENT,
    node_id integer,
    device_id integer,
    principal_switch_mac varchar(64),
    node_internal_ip varchar(30),
    node_public_ip varchar(30),
    node_principal varchar(30),
    node_is_local varchar(20),
    serial_nnum varchar(30),
    node_condition varchar(30),
    node_status varchar(60),
    firmware_version varchar(128),
    node_mac varchar(30),
    node_switch_type varchar(30),
    constraint mct_cluster_details_pkey primary key(id),
    constraint mct_cluster_details_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade
);
create table if not exists mct_cluster_configs
(
    id integer NOT NULL AUTO_INCREMENT,
    cluster_id integer,
    fabric_id integer,
    device_id integer,
    mct_neighbor_device_id integer,
    device_one_mgmt_ip varchar(30),
    device_two_mgmt_ip varchar(30),
    peer_interface_name varchar(30),
    peer_interfacetype varchar(30),
    peer_interface_speed varchar(30),
    peer_interface_two_speed varchar(30),
    peer_one_ip varchar(30),
    peer_two_ip varchar(30),
    control_vlan varchar(30),
    control_ve varchar(30),
    ve_interface_one_id integer,
    ve_interface_two_id integer,
    config_type varchar(30),
    local_node_id varchar(30),
    remote_node_id varchar(30),
    updated_attributes bigint,
    sw_cluster_name text not null,
    constraint mct_cluster_configs_pkey primary key(id),
    constraint mct_cluster_configs_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint mct_cluster_configs_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade,
    constraint mct_cluster_configs_mct_neighbor_device_id_fkey foreign key (mct_neighbor_device_id)
        references devices(id) on delete cascade
);
create table if not exists cluster_members
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    cluster_id integer,
    interface_id integer,
    device_id integer,
    remote_device_id integer,
    interface_name varchar(255),
    interface_type varchar(64),
    interface_speed integer,
    remote_interface_id integer,
    remote_interface_name varchar(255),
    remote_interface_type varchar(64),
    remote_interface_speed integer,
    config_type varchar(30),
    constraint cluster_members_pkey primary key(id),
    constraint cluster_members_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint cluster_members_cluster_id_fkey foreign key (cluster_id)
        references mct_cluster_configs(id) on delete cascade,
    constraint cluster_members_interface_id_fkey foreign key (interface_id)
        references phys_interfaces(id) on delete cascade,
    constraint cluster_members_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade,
    constraint cluster_members_remote_device_id_fkey foreign key (remote_device_id)
        references devices(id) on delete cascade,
    constraint cluster_members_remote_interface_id_fkey foreign key (remote_interface_id)
        references phys_interfaces(id) on delete cascade
);
create table if not exists fabric_errors
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    fabric_name varchar(30),
    device_id integer,
    device_ip varchar(30),
    device_role varchar(30),
    error_type varchar(30),
    error_reason text,
    constraint fabric_errors_pkey  primary key(id),
    constraint fabric_errors_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint fabric_errors_device_id_fkey foreign key (device_id)
        references devices(id)  on delete cascade
);
create table if not exists overlay_gateways
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    fabric_name varchar(30) not null,
    device_id integer,
    device_ip varchar(30) not null,
    sw_overlay_gw_name varchar(30) not null,
    gw_type varchar(30) not null,
    vtep_loopback_port_number varchar(30) not null,
    map_vni_auto varchar(30) not null,
    optimized_replication boolean default false,
    underlay_mdt_default_group varchar(100),
    activate varchar(30) not null,
    config_type varchar(30) not null,
    constraint overlay_gateways_pkey  primary key (id),
    constraint overlay_gateways_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id)  on delete cascade,
    constraint overlay_gateways_device_id_fkey foreign key (device_id)
        references devices(id)  on delete cascade
);
create table if not exists  evpns
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    fabric_name varchar(30) not null,
    device_id integer,
    device_ip varchar(30) not null,
    sw_evpn_name varchar(30) not null,
    mac_aging_timeout varchar(30) not null,
    mac_move_limit varchar(30) not null,
    mac_aging_conversational_timeout varchar(30) not null,
    arp_agingtimeout varchar(30) not null,
    duplicate_mac_timer varchar(30) not null,
    duplicate_mac_timer_max_count varchar(30) not null,
    config_type varchar(30) not null,
    constraint evpns_pkey  primary key(id),
    constraint evpns_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint evpns_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade
);
create table if not exists  router_bgps
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_id integer,
    local_asn bigint,
    router_id varchar(30),
    bfd_enable boolean,
    bfd_tx int,
    bfd_rx int,
    bfd_multiplier int,
    network_address varchar(30),
    redistribute_type varchar(30),
    as4_enable boolean,
    fast_fallover boolean,
    config_type varchar(30) not null,
    constraint router_bgp_pkey primary key(id),
    constraint router_bgp_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint router_bgp_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade
);
create table if not exists bgp_ip_afs
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_id integer,
    router_bgp_id integer,
    neighbor_afi varchar(100),
    neighbor_safi varchar(100),
    vrf varchar(100),
    graceful_restart boolean,
    max_path integer,
    config_type varchar(30),
    constraint bgp_ip_afs_pkey primary key(id),
    constraint bgp_ip_afs_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint bgp_ip_afs_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade,
    constraint bgp_ip_afs_router_bgp_id_fkey foreign key (router_bgp_id)
        references router_bgps(id)on delete cascade
);
create table if not exists bgp_af_ip_neighbors
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_id integer,
    remote_interface_id integer,
    remote_device_id integer,
    af_id integer,
    neighbor_afi varchar(100),
    neighbor_safi varchar(100),
    neighbor_ip_address varchar(30),
    remote_asn bigint,
    peer_group_name character varying(64),
    bfd_enable boolean,
    bfd_tx integer,
    bfd_rx integer,
    bfd_multiplier integer,
    e_bgp_mhop integer,
    enable_peer_as_check boolean,
    is_activated varchar(30),
    config_type varchar(30),
    constraint bgp_af_ip_neighbors_pkey primary key(id),
    constraint bgp_af_ip_neighbors_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint bgp_af_ip_neighbors_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade,
    constraint bgp_af_ip_neighbors_remote_interface_id_fkey foreign key (remote_interface_id)
        references phys_interfaces(id) on delete cascade,
    constraint bgp_af_ip_neighbors_remote_device_id_fkey foreign key (remote_device_id)
        references devices(id) on delete cascade,
    constraint bgp_af_ip_neighors_af_id_fkey foreign key (af_id)
        references bgp_ip_afs(id) on delete cascade
);
create table if not exists bgp_evpn_afs
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_id integer,
    router_bgp_id integer,
    retain_route_target_all boolean,
    graceful_restart boolean,
    config_type varchar(30),
    constraint bgp_evpn_afs_pkey primary key(id),
    constraint bgp_evpn_afs_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint bgp_evpn_afs_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade,
    constraint bgp_evpn_afs_router_bgp_id_fkey foreign key (router_bgp_id)
        references router_bgps(id) on delete cascade
);
create table if not exists bgp_af_evpn_neighbors
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer,
    device_id integer,
    remote_interface_id integer,
    remote_device_id integer,
    af_id integer,
    neighbor_afi varchar(100),
    neighbor_safi varchar(100),
    neighbor_ip_address varchar(30),
    remote_asn bigint,
    peer_group_name character varying(64),
    encapsulation_type varchar(30),
    bfd_enable boolean,
    bfd_tx integer,
    bfd_rx integer,
    bfd_multiplier integer,
    e_bgp_mhop integer,
    nexthop_unchanged boolean,
    enable_peer_as_check boolean,
    is_activated varchar(30),
    config_type varchar(30),
    constraint bgp_af_evpn_neighbors_pkey  primary key(id),
    constraint bgp_af_evpn_neighbors_fabric_id_fkey foreign key (fabric_id)
        references fabrics(id) on delete cascade,
    constraint bgp_af_evpn_neighbors_device_id_fkey foreign key (device_id)
        references devices(id) on delete cascade,
    constraint bgp_af_evpn_neighbors_remote_interface_id_fkey foreign key (remote_interface_id)
        references phys_interfaces(id) on delete cascade,
    constraint bgp_af_evpn_neighbors_remote_device_id_fkey foreign key (remote_device_id)
        references devices(id) on delete cascade,
    constraint bgp_af_evpn_neighors_af_id_fkey foreign key (af_id)
        references bgp_evpn_afs(id) on delete cascade
);

CREATE TABLE IF NOT EXISTS router_pims
(
    id integer NOT NULL AUTO_INCREMENT,
    fabric_id integer NOT NULL,
    device_id integer NOT NULL,
    vrf text,
    ssm_enable boolean,
    mdt_range varchar(100),
    config_type varchar(30) NOT NULL,

    CONSTRAINT router_pims_pkey PRIMARY KEY (id),
    CONSTRAINT router_pims_fabric_id_fkey FOREIGN KEY (fabric_id)
        REFERENCES fabrics (id) ON DELETE CASCADE,
    CONSTRAINT router_pims_device_id_fkey FOREIGN KEY (device_id)
        REFERENCES devices (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS ip_prefix_lists (
   id integer NOT NULL AUTO_INCREMENT,
   fabric_id integer NOT NULL,
   device_id integer NOT NULL,
   name varchar(64) NOT NULL,
   seq integer NOT NULL,
   afi varchar(20) NOT NULL default 'ipv4',
   action varchar(20),
   ip_prefix varchar(100),
   prefix_len_ge integer,
   prefix_len_le integer,
   config_type text NOT NULL,

   CONSTRAINT ip_prefix_lists_pkey PRIMARY KEY (id),
   CONSTRAINT ip_prefix_lists_fabric_id_fkey FOREIGN KEY (fabric_id)
       REFERENCES fabrics (id) ON DELETE CASCADE,
   CONSTRAINT ip_prefix_lists_device_id_fkey FOREIGN KEY (device_id)
       REFERENCES devices (id) ON DELETE CASCADE
);

