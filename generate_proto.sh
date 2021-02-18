protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/main.proto

protoc --dart_out=grpc:/media/mdrokz/md_data/md-2/my-files/dart/flutter/desktop/auto_bidder/lib/grpc -Iprotos protos/main.proto    