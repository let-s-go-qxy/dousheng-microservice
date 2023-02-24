go build -o ../output/api_gateway ./api_gateway/main.go
cd ./comment || exit
go build -o ../../output/comment
cd ../like || exit
go build -o ../../output/like
cd ../message || exit
go build -o ../../output/message
cd ../relation || exit
go build -o ../../output/relation
cd ../user || exit
go build -o ../../output/user
cd ../video || exit
go build -o ../../output/video