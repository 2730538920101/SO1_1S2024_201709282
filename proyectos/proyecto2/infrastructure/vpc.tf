# VPC
resource "google_compute_network" "vpc" {
  name                    = "${var.project_id}-vpc"
  auto_create_subnetworks = "false"
}

# Subnet
resource "google_compute_subnetwork" "subnet" {
  name          = "${var.project_id}-subnet"
  region        = var.region
  network       = google_compute_network.vpc.name
  ip_cidr_range = "10.10.0.0/24"
}

# Creación de reglas de firewall
resource "google_compute_firewall" "allow-ssh" {
  name    = "allow-ssh"
  network = google_compute_network.vpc.name

  allow {
    protocol = "tcp"
    ports    = ["22", "80", "443", "3000", "3001"]
  }

  source_ranges = ["0.0.0.0/0"]
}

# Creación de NAT para las subredes privadas
resource "google_compute_router_nat" "router_nat" {
  name            = "router-nat"
  router          = google_compute_router.router.name
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"

  nat_ip_allocate_option = "AUTO_ONLY"
}

# Asignación de subredes al router
resource "google_compute_router" "router" {
  name    = "router"
  network = google_compute_network.vpc.name

  bgp {
    asn       = 64514
    advertise_mode = "CUSTOM"
  }
}
