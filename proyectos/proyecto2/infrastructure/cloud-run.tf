resource "google_cloud_run_service_iam_member" "allow_public_access_backend" {
  service = google_cloud_run_service.backend_service.name
  location = var.region

  role = "roles/run.invoker"
  member = "allUsers"
}

# Servicio de Cloud Run para el backend
resource "google_cloud_run_service" "backend_service" {
  name     = "backend-service"
  location = var.region

  template {
    spec {
      containers {
        image = "carlosmz87/p2_backend_api:latest"  # Imagen de Docker para el backend

        # Configuración de puertos
        ports {
          container_port = 3000  # Puerto que utiliza el backend
        }

        # Configuración de variables de entorno
        env {
          name  = "DB_URL"
          value = "mongodb://admin:admin@34.29.36.163:27017?tls=false"
        }
        env {
          name  = "MONGO_DATABASE"
          value = "proyecto2"
        }
        env {
          name  = "MONGO_COLLECTION"
          value = "voto"
        }
      }
    }
  }
}


# Servicio de Cloud Run para el frontend
resource "google_cloud_run_service" "frontend_service" {
  name     = "frontend-service"
  location = var.region

  template {
    spec {
      containers {
        image = "carlosmz87/p2_frontend_api:latest"  # Imagen de Docker para el frontend

        # Configuración de puertos
        ports {
          container_port = 80  # Puerto que utiliza el frontend (Nginx)
        }

        # Configuración de variables de entorno
        # Si tienes variables de entorno necesarias, agrégalas aquí

        # Ejemplo de una variable de entorno:
        # env {
        #   name  = "VUE_APP_API_URL"
        #   value = "https://backend-service-dz3zpydyhq-uc.a.run.app"
        #}
      }
    }
  }
}

# Permitir acceso público al servicio de frontend
resource "google_cloud_run_service_iam_member" "allow_public_access_frontend" {
  service = google_cloud_run_service.frontend_service.name
  location = var.region

  role = "roles/run.invoker"
  member = "allUsers"
}
