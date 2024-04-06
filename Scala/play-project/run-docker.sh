IMAGE_NAME=play-project

docker build -t $IMAGE_NAME .
docker run -d -p 9000:9000 $IMAGE_NAME