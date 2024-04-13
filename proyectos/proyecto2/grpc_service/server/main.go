package main

import (
	"context"
	"database/sql"
	"fmt"
	pb "grpc_service/proto" // Importa el paquete generado a partir de tu archivo .proto
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

var ctx = context.Background()
var db *sql.DB

type server struct {
	pb.UnimplementedBandServiceServer
}

const (
	port = ":3001"
)

type BandsData struct {
	name  string
	album string
	year  string
	rank  string
}

func mysqlConnect() {
	// Cambia las credenciales según tu configuración de MySQL
	dsn := "root:t4-password@tcp(34.72.6.2:3306)/test_grpc"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Conexión a MySQL exitosa")
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
	insertMySQL(data)
	return &pb.BandResponse{Message: "Data recibida e insertada en la base de datos exitosamente desde el servidor"}, nil
}

func insertMySQL(voto BandsData) {
	// Prepara la consulta SQL para la inserción en MySQL
	query := "INSERT INTO VOTO (NOMBRE, ALBUM, ANIO, RANKING) VALUES (?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, query, voto.name, voto.album, voto.year, voto.rank)
	if err != nil {
		log.Println("Error al insertar en MySQL:", err)
	}
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterBandServiceServer(s, &server{})

	mysqlConnect()

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
