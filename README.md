# hackathon
Setup the docker database (ps, you might want to change ports if cannot run port 3306...)

-docker build -t testdatabaseimage .
-docker run --name mysqltestdb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=strongpassword -d testdatabaseimage
*if 3306 is used, you can kill the process or use an alternative port (e.g 3307:3306)
-docker exec -it mysqltestdb /bin/bash
-mysql -u user -p
when prompted to enter password, key in: strongpassword
