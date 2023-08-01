package service

import (
	"context"
	"film_grpc/proto"
	"film_grpc/serverGo/initializers"
	"film_grpc/serverGo/models"
	"google.golang.org/grpc"
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

func (*server) CreateGenre(ctx context.Context, in *proto.GenreData, opts ...grpc.CallOption) (*proto.GenreResponse, error) {

	genreName := in.GetGenreName()

	//genre := models.Genre{GenreName: genreName}
	//result := initializers.DB.Create(&genre)
	//if result.Error != nil {
	//	return &proto.GenreResponse{}, result.Error
	//}
	//
	//response := &proto.GenreResponse{
	//	ID:        uint32(genre.ID),
	//	GenreName: genre.GenreName,
	//}
	//return response, nil

	genre := models.Genre{GenreName: genreName}
	result := initializers.DB.Create(&genre)
	if result.Error != nil {
		return &proto.GenreResponse{}, result.Error
	}

	return &proto.GenreResponse{
		ID:        uint32(genre.ID),
		GenreName: genre.GenreName,
	}, nil

}
