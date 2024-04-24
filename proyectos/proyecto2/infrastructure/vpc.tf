# VPC
resource "google_compute_network" "vpc" {
  name                    = "${var.project_id}-vpc"
  auto_create_subnetworks = false
}

# Subnet
resource "google_compute_subnetwork" "subnet" {
  name          = "${var.project_id}-subnet"
  region        = var.region
  network       = google_compute_network.vpc.name
  ip_cidr_range = "10.10.0.0/24"
}

# Creación de reglas de firewall
resource "google_compute_firewall" "allow-traffic" {
  name    = "allow-traffic"
  network = google_compute_network.vpc.name

  allow {
    protocol = "tcp"
    ports    = ["80", "443", "3000", "3001"]
  }

  source_ranges = ["0.0.0.0/0"]  # Permite tráfico desde cualquier origen

  description = "Allow HTTP, HTTPS, and custom ports for backend and frontend"
}

resource "google_compute_firewall" "allow-vpc-connector-inbound" {
  name    = "allow-vpc-connector-inbound"
  network = google_compute_network.vpc.name

  allow {
    protocol = "all"  # Puedes especificar un protocolo específico si es necesario
  }

  source_ranges = ["10.10.1.0/28"]  # Rango de IP del VPC Connector

  destination_ranges = ["10.10.0.0/24"]  # Rango de IP de tu subred en la VPC

  description = "Permitir tráfico entrante desde el VPC Connector a la VPC"
}

resource "google_compute_firewall" "allow-vpc-connector-outbound" {
  name    = "allow-vpc-connector-outbound"
  network = google_compute_network.vpc.name

  allow {
    protocol = "all"  # Puedes especificar un protocolo específico si es necesario
  }

  destination_ranges = ["10.10.1.0/28"]  # Rango de IP del VPC Connector

  source_ranges = ["10.10.0.0/24"]  # Rango de IP de tu subred en la VPC

  description = "Permitir tráfico saliente desde la VPC hacia el VPC Connector"
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
  region  = var.region

  bgp {
    asn = 64514
    advertise_mode = "CUSTOM"
  }
}

