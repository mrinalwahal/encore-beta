docker run --name postgres --publish 6432:5432 --env PORT=6432  -e POSTGRES_PASSWORD=mysecretpassword -d postgres