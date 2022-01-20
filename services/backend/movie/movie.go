package movie

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/jeongpope/example-msa/services/backend/model"
)

var (
	movies []model.Movie
)

func init() {
	bin, err := ioutil.ReadFile("./mock/info.json")
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(bin, &movies)
}

func GetMovies() []model.Movie {
	return movies
}

func GetMovie(id int) (movie model.Movie, err error) {
	err = checkID(id)

	movie = movies[id-1]

	return
}

func CreateMovie(m model.Movie) (movie model.Movie, err error) {
	movies = append(movies, m)

	len := len(movies)

	return movies[len-1], nil
}

func UpdateMovie(id int, m model.Movie) (movie model.Movie, err error) {
	err = checkID(id)
	if err != nil {
		return model.Movie{}, err
	}

	movies[id-1] = m

	return movies[id-1], nil
}

func DeleteMovie(id int) (err error) {
	err = checkID(id)
	if err != nil {
		return
	}

	movies = append(movies[:id-1], movies[id:]...)

	return
}

func checkID(id int) (err error) {
	if id > len(movies) || id <= 0 {
		err = errors.New("index out of range")
		return
	}

	return nil
}
