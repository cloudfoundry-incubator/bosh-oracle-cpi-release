# Availability Domain
data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.oracle_tenancy_ocid}"
    filter {
      name = "name"
      values = ["${var.director_ad}"]
    }
}

data "oci_identity_compartments" "Compartments" {
  compartment_id = "${var.oracle_tenancy_ocid}"
  filter {
    name = "name"
    values = ["${var.director_compartment_name}"]
  }
}

data "null_data_source" "SetupConfig" {
  inputs = {
    ad_name = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    compartment_id = "${lookup(data.oci_identity_compartments.Compartments.compartments[0],"id")}"
  }
}