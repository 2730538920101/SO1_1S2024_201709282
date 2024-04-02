package main

import (
	"context"
	"log"
	"net"

	pb "github.com/2730538920101/SO1_1S2024_201709282/tree/main/proyectos/proyecto2/grpc_service/proto" // Importa el paquete generado a partir de tu archivo .proto

	"google.golang.org/grpc"

	"github.com/segmentio/kafka-go"
)

// Implementa la estructura del servidor
type server struct {
	kafkaWriter *kafka.Writer
}

// Implementa el método SendBandInfo de la interfaz BandServiceServer
func (s *server) SendBandInfo(ctx context.Context, in *pb.Band) (*pb.BandResponse, error) {
	// Envía los datos recibidos al servidor de Kafka
	err := s.kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(in.String()), // Convierte el mensaje de protobuf a []byte
	})
	if err != nil {
		log.Printf("Error sending message to Kafka: %v", err)
		return nil, err
	}

	// Devuelve una respuesta de éxito al cliente
	return &pb.BandResponse{Message: "Band information received and sent to Kafka successfully"}, nil
}

func main() {
	// Define el puerto en el que el servidor gRPC escuchará las solicitudes
	port := ":50051"

	// Inicializa el cliente de Kafka
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"kafka-broker1:9092", "kafka-broker2:9092"}, // Actualiza con las direcciones de tus brokers Kafka
		Topic:    "band_topic",                                         // Nombre del tema en el que se enviarán los mensajes
		Balancer: &kafka.LeastBytes{},
	})

	// Defer cerrar el cliente de Kafka al finalizar el programa
	defer kafkaWriter.Close()

	// Crea una nueva instancia de servidor gRPC
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Crea un nuevo servidor gRPC
	s := grpc.NewServer()

	// Registra la implementación del servidor junto con el cliente de Kafka
	pb.RegisterBandServiceServer(s, &server{kafkaWriter})

	// Inicia el servidor
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
