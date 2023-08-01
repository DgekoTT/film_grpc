package service

import (
	"context"
	pb "film_grpc/proto"
	"film_grpc/serverGo/initializers"
	"film_grpc/serverGo/models"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFilmGenreServiceServer
}

func (*server) CreateFilm(ctx context.Context, in *pb.FilmData, opts ...grpc.CallOption) (*pb.FilmResponse, error) {
	filmName := in.GetFilmName()
	productionYear := in.GetProductionYear()
	massGenres := GetGenreIdsByName(in.GetGenres())

	film := models.Film{
		FilmName:       filmName,
		ProductionYear: uint16(productionYear),
		Genres:         massGenres,
	}
	result := initializers.DB.Create(&film)
	if result.Error != nil {
		return &pb.FilmResponse{}, result.Error
	}

	return &pb.FilmResponse{
		ID:             uint64(film.ID),
		FilmName:       film.FilmName,
		ProductionYear: uint32(film.ProductionYear),
	}, nil

}
