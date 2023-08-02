package main

import (
	"context"
	pb "film_grpc/proto"
	"film_grpc/serverGo/initializers"
	"film_grpc/serverGo/models"
	"log"
	"strings"
)

func GetGenreIdsByName(name string) []*models.Genre {
	namesGenre := strings.Split(name, ",")
	var genreIds []*models.Genre
	err := initializers.DB.Where("genre_name IN ?",
		namesGenre).Find(&genreIds).Error
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса:", err)
	}
	return genreIds
}

func (*Server) CreateGenre(ctx context.Context, in *pb.Genre) (*pb.Genre, error) {

	genreName := in.GetGenreName()
	log.Print(genreName)

	genre := models.Genre{GenreName: genreName}
	result := initializers.DB.Create(&genre)
	if result.Error != nil {
		return &pb.Genre{}, result.Error
	}

	return &pb.Genre{
		ID:        genre.ID,
		GenreName: genre.GenreName,
	}, nil

}

//func (*Server) GetAllGenres(ctx context.Context, in *pb.NoParams) (*pb.GetGenresResponse, error) {
//	var genres []*pb.GenreResponse
//	err := initializers.DB.Preload("Genres")Find(&genres).Error
//	if err != nil {
//		return &pb.GetGenresResponse{}, err
//	}
//	return &pb.GetGenresResponse{
//		Genres: genres,
//	}, nil
//}
