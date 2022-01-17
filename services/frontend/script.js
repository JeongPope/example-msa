
//alert("Please Use Desktop View For Better Expeience")
filterSelection("all")
getMovies()

function getMovies() {
  fetch('http://localhost:8880/api/movies')
    .then(function (response) {
      response.json().then(function (json) {
        for (i = 0; i < json.length; i++) {
          var movie = json[i]

          document.getElementsByClassName("cards")[0].innerHTML +=
            "<li>" +
            "<a href=\"movie.html\">" +
            "<div class=\"row\">" +
            "<div class=\"column nature show\">" +
            "<div class=\"content\" id=" + (i + 1) + "\" onclick=\"setMovie(" + (i + 1) + ")\">" +
            "<img src=\"assets/" + movie.image + "\"" +
            "alt=\"" + movie.alt + "\" style=\"width:100%\">" +
            "<h4>" + movie.title + "</h4>" +
            "</div>" +
            "</div>" +
            "</div>" +
            "</a>" +
            "</li>";
        }
      })
    })
    .catch(function (error) {
      console.log("Error: " + error);
    });
}

function filterSelection(c) {
  var x, i;
  x = document.getElementsByClassName("column");
  if (c == "all") c = "";
  for (i = 0; i < x.length; i++) {
    w3RemoveClass(x[i], "show");
    if (x[i].className.indexOf(c) > -1) w3AddClass(x[i], "show");
  }
}

function w3AddClass(element, name) {
  var i, arr1, arr2;
  arr1 = element.className.split(" ");
  arr2 = name.split(" ");
  for (i = 0; i < arr2.length; i++) {
    if (arr1.indexOf(arr2[i]) == -1) { element.className += " " + arr2[i]; }
  }
}

function w3RemoveClass(element, name) {
  var i, arr1, arr2;
  arr1 = element.className.split(" ");
  arr2 = name.split(" ");
  for (i = 0; i < arr2.length; i++) {
    while (arr1.indexOf(arr2[i]) > -1) {
      arr1.splice(arr1.indexOf(arr2[i]), 1);
    }
  }
  element.className = arr1.join(" ");
}


// // Add active class to the current button (highlight it)
// var btnContainer = document.getElementById("myBtnContainer");
// var btns = btnContainer.getElementsByClassName("btn");
// for (var i = 0; i < btns.length; i++) {
//   btns[i].addEventListener("click", function () {
//     var current = document.getElementsByClassName("active");
//     current[0].className = current[0].className.replace(" active", "");
//     this.className += " active";
//   });
// }

function setMovie(id) {
  localStorage.setItem('classlion', id)
}