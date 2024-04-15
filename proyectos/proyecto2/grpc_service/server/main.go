package main

import (
	"context"
	"fmt"
	pb "grpc_server/proto" // Importa el paquete generado a partir de tu archivo .proto
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// CargarVariablesEntorno carga las variables de entorno desde el archivo .env
func CargarVariablesEntorno() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No se pudo cargar el archivo .env, cargando variables de entorno del sistema...")
	}
	return nil
}

type server struct {
	pb.UnimplementedBandServiceServer
}

type BandsData struct {
	name  string
	album string
	year  string
	rank  string
}

func (s *server) SendBandInfo(ctx context.Context, in *pb.Band) (*pb.BandResponse, error) {
	fmt.Println("Servidor ha recibido informacion desde el cliente")
	data := BandsData{
		name:  in.GetName(),
		album: in.GetAlbum(),
		year:  in.GetYear(),
		rank:  in.GetRank(),
	}
	fmt.Println(data)

	return &pb.BandResponse{Message: "Data recibida exitosamente desde el servidor"}, nil
}

func main() {

	err := CargarVariablesEntorno()
	client_port := obtenerPuertoCliente()
	if err != nil {
		fmt.Println("Error cargando variables de entorno desde el archivo .env:", err)
		fmt.Println("Obteniendo variables de entorno del sistema...")
	}
	fmt.Printf("La comunicacion con el cliente se realiza en el puerto: %d\n", client_port)
	// Inicializa client_port después de cargar las variables de entorno
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", client_port))
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterBandServiceServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}

func obtenerPuertoCliente() int {
	portStr := os.Getenv("CLIENT_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		// Manejar el error en caso de que el valor de la variable de entorno no sea un número
		fmt.Println("Error al convertir el valor de PORT a un número:", err)
	}
	return port
}
