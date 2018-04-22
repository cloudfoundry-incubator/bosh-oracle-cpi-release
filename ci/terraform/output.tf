output vcn {
  value = "${oci_core_virtual_network.ci_vcn.display_name}"
}

output vcn_id {
  value = "${oci_core_virtual_network.ci_vcn.id}"
}

output subnet_id {
  value = "${oci_core_subnet.director_subnet.id}"
}
output compartment_id {
  value = "${oci_core_subnet.director_subnet.compartment_id}"
}

output subnet_name {
  value = "${oci_core_subnet.director_subnet.display_name}"
}
output subnet_cidr {
  value = "${oci_core_subnet.director_subnet.cidr_block}"
}
output subnet_gw {
  value = "${cidrhost(oci_core_subnet.director_subnet.cidr_block, 1)}"
}
output subnet_first_ip {
  value = "${cidrhost(oci_core_subnet.director_subnet.cidr_block, 2)}"
}

output internal_cidr {
  value = "${oci_core_subnet.director_subnet.cidr_block}"
}
output internal_gw {
  value = "${cidrhost(oci_core_subnet.director_subnet.cidr_block, 1)}"
}
output internal_ip {
  value = "${cidrhost(oci_core_subnet.director_subnet.cidr_block, 2)}"
}

output bats_subnet1_name {
  value = "${oci_core_subnet.bats_subnet1.display_name}"
}
output bats_subnet1_id {
  value = "${oci_core_subnet.bats_subnet1.id}"
}

output bats_subnet1_cidr {
  value = "${oci_core_subnet.bats_subnet1.cidr_block}"
}

output bats_subnet1_gw {
  value = "${cidrhost(oci_core_subnet.bats_subnet1.cidr_block, 1)}"
}

output bats_subnet1_reserved {
  value = "${cidrhost(oci_core_subnet.bats_subnet1.cidr_block, 2)} - ${cidrhost(oci_core_subnet.bats_subnet1.cidr_block, 9)}"
}

output bats_subnet1_static {
  value = "${cidrhost(oci_core_subnet.bats_subnet1.cidr_block, 10)} - ${cidrhost(oci_core_subnet.bats_subnet1.cidr_block, 30)}"
}

output bats_subnet1_static_ip {
  value = "${cidrhost(oci_core_subnet.bats_subnet1.cidr_block, 30)}"
}

output bats_subnet2_name {
  value = "${oci_core_subnet.bats_subnet2.display_name}"
}
output bats_subnet2_id {
  value = "${oci_core_subnet.bats_subnet2.id}"
}

output bats_subnet2_cidr {
  value = "${oci_core_subnet.bats_subnet2.cidr_block}"
}

output bats_subnet2_gw {
  value = "${cidrhost(oci_core_subnet.bats_subnet2.cidr_block, 1)}"
}

output bats_subnet2_reserved {
  value = "${cidrhost(oci_core_subnet.bats_subnet2.cidr_block, 2)} - ${cidrhost(oci_core_subnet.bats_subnet2.cidr_block, 9)}"
}

output bats_subnet2_static {
  value = "${cidrhost(oci_core_subnet.bats_subnet2.cidr_block, 10)} - ${cidrhost(oci_core_subnet.bats_subnet2.cidr_block, 30)}"
}

output bats_subnet2_static_ip {
  value = "${cidrhost(oci_core_subnet.bats_subnet2.cidr_block, 30)}"
}

output external_ip {
  value = "${oci_core_public_ip.director_vip.ip_address}"
}

output tenancy {
  sensitive = true
  value = "${var.oracle_tenancy_ocid}"
}
output user {
  sensitive = true
  value = "${var.oracle_user_ocid}"
}
output fingerprint {
  sensitive = true
  value = "${var.oracle_fingerprint}"
}
output compartment {
  value = "${oci_core_subnet.director_subnet.compartment_id}"
}
output ad {
  value = "${var.director_ad}"
}
output region {
  value = "${var.oracle_region}"
}
output apikey {
  sensitive = true
  value = "${file(var.oracle_private_key_path)}"
}
