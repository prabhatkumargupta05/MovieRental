This project is licensed under the Thoughtworks license

Post command to add a movie to the cart.
curl -i -X POST -H 'Content-Type: application/json' -d '{"imdbID":"Real_Movie_Imdb_ID"}' http://localhost:8080/movierental/cart/movie