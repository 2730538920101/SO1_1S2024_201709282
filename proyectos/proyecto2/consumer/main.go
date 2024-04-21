package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

// CargarVariablesEntorno carga las variables de entorno desde el archivo .env
func CargarVariablesEntorno() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No se encuentra el archivo .env, puede estar en el entorno de produccion, cargando variables de entorno del sistema...")
	}
}

func consumeFromKafka() {
	// Carga las variables de entorno
	CargarVariablesEntorno()

	// Configura el consumidor de Kafka
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{os.Getenv("KAFKA_BROKER")},
		Topic:     os.Getenv("KAFKA_TOPIC"),
		Partition: 0,
		MinBytes:  10e3, // El tamaño mínimo de bytes para solicitar datos del servidor
		MaxBytes:  10e6, // El tamaño máximo de bytes para solicitar datos del servidor
	})
	defer reader.Close()

	ctx := context.Background()
	fmt.Println("Iniciando el consumo de mensajes de Kafka...")

	for {
		// Lee mensajes de Kafka
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("Error al leer mensaje de Kafka: %s", err)
		}

		// Muestra la información del mensaje recibido
		fmt.Printf("Mensaje recibido de Kafka: %s\n", string(msg.Value))
	}
}

func main() {
	// Inicia el consumidor de Kafka
	consumeFromKafka()
}
