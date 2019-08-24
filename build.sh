set +v
export GOPATH = $PWD
set -v
go get goatbrote
go install goatbrote
set +v
cp bin\goatbrote goatbrote
sudo chmod +x goatbrote
.\goatbrote
