## PASO 1 
INSTALAR EL SDK DE GOOGLE CLOUD EN TERMINAL Y TERRAFORM
## PASO 2
INICIAR SESION EN GOOGLE CLOUD CON EL COMANDO gcloud auth login
## PASO 3 
INICIAR TERRAFORM CON EL COMANDO terraform init EJECUTANDOLO EN LA CARPETA infrastructure
## PASO 4
EJECUTAR terraform plan Y VALIDAR QUE TODOS LOS COMPONENTES DEL CLUSTER ESTEN CORRECTAMENTE CONFIGURADOS
## PASO 5
EJECUTAR terraform init Y VERIFICAR QUE EL CLUSTER SE HAYA CREADO CORRECTAMENTE

# PASOS PARA VERIFICAR LA EXISTENCIA DEL CLUSTER

## PASO 1
Instalar kubectl en la herramienta gcloud con el comando:
gcloud components install kubectl
## PASO 2
INSTALAR EL PLUGIN DE AUTENTICACION DE GKE
gcloud components install gke-gcloud-auth-plugin
## PASO 3
Obtener las credenciales del cluster con el comando: 
-   gcloud container clusters get-credentials $(terraform output -raw kubernetes_cluster_name) --region $(terraform output -raw region)
## PASO 4
Activar el kubectl dashboard con el comando: kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
## PASO 5 
ACTIVAR EL KUBECTL PROXY CON EL COMANDO: kubectl proxy
## PASO 6 
SIN CERRAR LA TERMINAL DEL PROXY, ABRIR UNA NUEVA Y AUTENTICAR EL ACCESO AL DASHBOARD CON EL COMANDO: 
-   kubectl apply -f https://raw.githubusercontent.com/hashicorp/learn-terraform-provision-gke-cluster/main/kubernetes-dashboard-admin.rbac.yaml
GENERAR EL TOKEN DE AUTORIZACION CON EL COMANDO:
-   kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep service-controller-token | awk '{print $1}')
AHORA COPIAR Y PEGAR EL TOKEN EN EL DASHBOARD

## PARA INGRESAR AL DASHBOARD USAR LA SIGUIENTE URL
http://127.0.0.1:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login
