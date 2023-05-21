docker run -d -p 27017:27017 --name mongodb_fusupo --network host mongo:3.6
docker exec -it mongodb_fusupo mongo
