package movie

import (
	"encoding/json"
	"errors"
	"fmt"
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
	if id > len(movies) || id <= 0 {
		err = errors.New("index out of range")
		return
	}

	movie = movies[id-1]

	return
}

func CreateMovie(m model.Movie) (movie model.Movie, err error) {
	fmt.Println(m)
	movies = append(movies, m)

	len := len(movies)
	movies[len-1].ID = len

	return movies[len-1], nil
}

// func UpdateMovie(id int) (movie model.Movie, err error) {
// 	if id > len(movies) || id <= 0 {
// 		err = errors.New("index out of range")
// 		return
// 	}

// }

// func DeleteMovie(id int) (movie model.Movie, err error) {
// 	if id > len(movies) || id <= 0 {
// 		err = errors.New("index out of range")
// 		return
// 	}
// }
