echo building solution
go build
echo copying sls to /usr/local/bin
sudo cp ./sls /usr/local/bin
echo ...done