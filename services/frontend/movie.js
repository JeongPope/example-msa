getMovie()

function getMovie() {
    let url = 'http://172.30.7.250:31011/api/movie/'
    let val = localStorage.getItem('classlion')
    url += val

    fetch(url)
        .then(function (response) {
            response.json().then(function (json) {
                var movie = json
                console.log(movie)

                document.getElementsByClassName("container")[0].innerHTML +=
                    "<img src=\"assets/" + movie.image + "\" " +
                    "alt=\"" + movie.alt + "\" class=\"image\">";

                document.getElementsByClassName("desc")[0].innerHTML +=
                    "<center>" +
                    "<h1>" + movie.title + "</h1>" +
                    "<p1>" + movie.desc + "</p1>" +
                    "<br><br>" +
                    "<p3> JEONG JE HO </p3>" +
                    "</center>";
            })
        })
        .catch(function (error) {
            console.log("Error: " + error);
        });
}