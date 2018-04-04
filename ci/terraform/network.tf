resource "oci_core_internet_gateway" "vcn_ig" {
  compartment_id = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  display_name   = "ci_vcn_${var.director_vcn}_ig"
  vcn_id         = "${data.null_data_source.VCN.inputs.id}"
}

resource "oci_core_route_table" "vcn_ig_route_table" {
  compartment_id = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  vcn_id         = "${data.null_data_source.VCN.inputs.id}"
  display_name   = "${oci_core_internet_gateway.vcn_ig.display_name}_route_table"
  route_rules {
    cidr_block        = "0.0.0.0/0"
    network_entity_id = "${oci_core_internet_gateway.vcn_ig.id}"
  }
}


resource "oci_core_security_list" "ci_public_all" {
    compartment_id        = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
    display_name          = "ci_public_all"
    vcn_id                = "${data.null_data_source.VCN.inputs.id}"
    egress_security_rules = [{
        protocol = "all"
        destination = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        protocol = "all"
        source = "0.0.0.0/0"
    }]
}
resource "oci_core_subnet" "director_subnet" {
  availability_domain = "${data.null_data_source.SetupConfig.inputs.ad_name}"
  cidr_block          = "${var.director_subnet_cidr}"
  display_name        = "ci_director_subnet_${replace(data.null_data_source.SetupConfig.inputs.ad_name, "-", "_")}"
  dhcp_options_id     = "${data.null_data_source.VCN.inputs.dhcp_options_id}"
  compartment_id      = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  vcn_id              = "${data.null_data_source.VCN.inputs.id}"
  route_table_id      = "${oci_core_route_table.vcn_ig_route_table.id}"
  security_list_ids   = ["${oci_core_security_list.ci_public_all.id}"]
  prohibit_public_ip_on_vnic = false
}

resource "oci_core_subnet" "bats_subnet1" {
  availability_domain = "${data.null_data_source.SetupConfig.inputs.ad_name}"
  cidr_block          = "${var.bats_subnet1_cidr}"
  display_name        = "ci_bats_subnet1_${replace(data.null_data_source.SetupConfig.inputs.ad_name, "-", "_")}"
  dhcp_options_id     = "${data.null_data_source.VCN.inputs.dhcp_options_id}"
  compartment_id      = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  vcn_id              = "${data.null_data_source.VCN.inputs.id}"
  route_table_id      = "${oci_core_route_table.vcn_ig_route_table.id}"
  security_list_ids   = ["${oci_core_security_list.ci_public_all.id}"]
  prohibit_public_ip_on_vnic = false
}

resource "oci_core_subnet" "bats_subnet2" {
  availability_domain = "${data.null_data_source.SetupConfig.inputs.ad_name}"
  cidr_block          = "${var.bats_subnet2_cidr}"
  display_name        = "ci_bats_subnet2_${replace(data.null_data_source.SetupConfig.inputs.ad_name, "-", "_")}"
  dhcp_options_id     = "${data.null_data_source.VCN.inputs.dhcp_options_id}"
  compartment_id      = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  vcn_id              = "${data.null_data_source.VCN.inputs.id}"
  route_table_id      = "${oci_core_route_table.vcn_ig_route_table.id}"
  security_list_ids   = ["${oci_core_security_list.ci_public_all.id}"]
  prohibit_public_ip_on_vnic = false
}
