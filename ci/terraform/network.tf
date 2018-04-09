resource "oci_core_virtual_network" "ci_vcn" {
  cidr_block = "${var.vcn_cidr}"
  compartment_id = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  display_name = "${var.director_vcn}"
  dns_label = "civcn"
}

resource "oci_core_internet_gateway" "vcn_ig" {
  compartment_id = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  display_name = "${oci_core_virtual_network.ci_vcn.display_name}_ig"
  vcn_id = "${oci_core_virtual_network.ci_vcn.id}"
}

resource "oci_core_route_table" "vcn_ig_route_table" {
  compartment_id = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  vcn_id = "${oci_core_virtual_network.ci_vcn.id}"
  display_name = "${oci_core_internet_gateway.vcn_ig.display_name}_route_table"
  route_rules {
    cidr_block = "0.0.0.0/0"
    network_entity_id = "${oci_core_internet_gateway.vcn_ig.id}"
  }
}

resource "oci_core_security_list" "ci_public_all" {
  compartment_id = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  display_name = "ci_public_all"
  vcn_id = "${oci_core_virtual_network.ci_vcn.id}"
  egress_security_rules = [
    {
      protocol = "all"
      destination = "0.0.0.0/0"
    }]
  ingress_security_rules = [
    {
      protocol = "1"
      source = "0.0.0.0/0"
      icmp_options {
        type = "3"
        code = "4"
      }
    },
    {
      protocol = "6"
      source = "0.0.0.0/0"
      tcp_options {
        "max" = 22
        "min" = 22
      }
    },
    {
      protocol = "6"
      source = "0.0.0.0/0"
      tcp_options {
        "max" = 6868
        "min" = 6868
      }
    },
    {
      protocol = "6"
      source = "0.0.0.0/0"
      tcp_options {
        "max" = 8443
        "min" = 8443
      }
    },
    {
      protocol = "6"
      source = "0.0.0.0/0"
      tcp_options {
        "max" = 8844
        "min" = 8844
      }
    },
    {
      protocol = "6"
      source = "0.0.0.0/0"
      tcp_options {
        "max" = 25250
        "min" = 25250
      }
    },
    {
      protocol = "6"
      source = "0.0.0.0/0"
      tcp_options {
        "max" = 25555
        "min" = 25555
      }
    },
    {
      protocol = "6"
      source = "${oci_core_virtual_network.ci_vcn.cidr_block}"
    }
  ]
}
resource "oci_core_subnet" "director_subnet" {
  availability_domain = "${data.null_data_source.SetupConfig.inputs.ad_name}"
  cidr_block = "${var.director_subnet_cidr}"
  display_name = "ci_director_subnet_${replace(data.null_data_source.SetupConfig.inputs.ad_name, "-", "_")}"
  dhcp_options_id = "${oci_core_virtual_network.ci_vcn.default_dhcp_options_id}"
  compartment_id = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  vcn_id = "${oci_core_virtual_network.ci_vcn.id}"
  route_table_id = "${oci_core_route_table.vcn_ig_route_table.id}"
  security_list_ids = [
    "${oci_core_security_list.ci_public_all.id}"]
  prohibit_public_ip_on_vnic = false
}

resource "oci_core_subnet" "bats_subnet1" {
  availability_domain = "${data.null_data_source.SetupConfig.inputs.ad_name}"
  cidr_block = "${var.bats_subnet1_cidr}"
  display_name = "ci_bats_subnet1_${replace(data.null_data_source.SetupConfig.inputs.ad_name, "-", "_")}"
  dhcp_options_id = "${oci_core_virtual_network.ci_vcn.default_dhcp_options_id}"
  compartment_id = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  vcn_id = "${oci_core_virtual_network.ci_vcn.id}"
  route_table_id = "${oci_core_route_table.vcn_ig_route_table.id}"
  security_list_ids = [
    "${oci_core_security_list.ci_public_all.id}"]
  prohibit_public_ip_on_vnic = false
}

resource "oci_core_subnet" "bats_subnet2" {
  availability_domain = "${data.null_data_source.SetupConfig.inputs.ad_name}"
  cidr_block = "${var.bats_subnet2_cidr}"
  display_name = "ci_bats_subnet2_${replace(data.null_data_source.SetupConfig.inputs.ad_name, "-", "_")}"
  dhcp_options_id = "${oci_core_virtual_network.ci_vcn.default_dhcp_options_id}"
  compartment_id = "${data.null_data_source.SetupConfig.inputs.compartment_id}"
  vcn_id = "${oci_core_virtual_network.ci_vcn.id}"
  route_table_id = "${oci_core_route_table.vcn_ig_route_table.id}"
  security_list_ids = [
    "${oci_core_security_list.ci_public_all.id}"]
  prohibit_public_ip_on_vnic = false
}
