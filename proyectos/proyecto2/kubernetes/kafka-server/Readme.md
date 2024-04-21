## PASO 1
INSTALACION DE KAFKA CON STRIMZI

UTILIZAR LOS COMANDOS:

-   kubectl create -f 'https://strimzi.io/install/latest?namespace=<nombre_namespace>' -n  <nombre_namespace>

-   kubectl apply -f https://strimzi.io/examples/latest/kafka/kafka-persistent-single.yaml -n <nombre_namespace>