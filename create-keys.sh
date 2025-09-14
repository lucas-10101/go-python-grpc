openssl genrsa -out server.key 2048
openssl req -new -x509 -key server.key -out server.crt -days 365 -subj "/CN=localhost"

cp server.crt ./server/server.crt
cp server.key ./server/server.key

cp server.crt ./client/server.crt