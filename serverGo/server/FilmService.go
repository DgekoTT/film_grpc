package main

import (
	"context"
	"errors"
	pb "film_grpc/proto"
	"film_grpc/serverGo/initializers"
	"film_grpc/serverGo/models"
)

func (*Server) CreateFilm(ctx context.Context, in *pb.FilmData) (*pb.Film, error) {
	genres := in.GetGenres()

	film := models.Film{
		FilmName:       in.GetFilmName(),
		ProductionYear: in.GetProductionYear(),
		Genres:         GetGenreIdsByName(genres),
	}
	result := initializers.DB.Create(&film)
	if result.Error != nil {
		return &pb.Film{}, result.Error
	}

	return &pb.Film{
		ID:             film.ID,
		FilmName:       film.FilmName,
		ProductionYear: film.ProductionYear,
		Genres:         convertGenreToGenreResponse(film.Genres),
	}, nil

}

func (*Server) GetAllFilms(ctx context.Context, in *pb.NoParams) (*pb.GetFilmResponse, error) {
	var Films []*models.Film
	err := initializers.DB.Preload("Genres").Find(&Films).Error
	if err != nil {
		return &pb.GetFilmResponse{}, err
	}
	return &pb.GetFilmResponse{
		Films: convertToPbFilms(Films),
	}, nil
}

func (*Server) GetFilmById(ctx context.Context, in *pb.ID) (*pb.Film, error) {
	var Film *models.Film

	res := initializers.DB.Preload("Genres").Find(&Film, "ID = ?", in.GetID())
	if res.RowsAffected == 0 {
		return nil, errors.New("фильм не найден")
	}
	return convertToPbFilm(Film), nil
}
func (*Server) UpDateFilm(ctx context.Context, in *pb.InfoFilmUpdate) (*pb.OperationStatus, error) {
	var film *models.Film

	res := initializers.DB.Model(&film).Where("id = ?", in.ID).Updates(models.Film{FilmName: in.FilmName, ProductionYear: in.ProductionYear})

	if res.RowsAffected == 0 {
		return nil, errors.New("фильм не найден")
	}
	return &pb.OperationStatus{
		OperationStatus: "фильм обновлен успешно"}, nil
}

func (*Server) DeleteFilm(ctx context.Context, in *pb.ID) (*pb.OperationStatus, error) {
	var film *models.Film

	res := initializers.DB.Where("id = ?", in.ID).Delete(&film)
	if res.RowsAffected == 0 {
		return nil, errors.New("фильм не найден")
	}

	return &pb.OperationStatus{
		OperationStatus: "фильм удален успешно"}, nil

}

func convertGenreToGenreResponse(genres []*models.Genre) []*pb.Genre {
	var genreResponses []*pb.Genre

	for _, genre := range genres {
		genreResponses = append(genreResponses, &pb.Genre{
			ID:        genre.ID,
			GenreName: genre.GenreName,
		})
	}

	return genreResponses
}

func convertToPbFilms(films []*models.Film) []*pb.Film {
	var pbFilms []*pb.Film

	for _, film := range films {
		pbFilm := convertToPbFilm(film)
		pbFilms = append(pbFilms, pbFilm)
	}

	return pbFilms
}

func convertToPbFilm(film *models.Film) *pb.Film {
	return &pb.Film{
		ID:             film.ID,
		FilmName:       film.FilmName,
		ProductionYear: film.ProductionYear,
		Genres:         convertToPbGenres(film.Genres),
	}
}
