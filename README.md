# hackathon
Setup the docker database (ps, you might want to change ports if cannot run port 3306...)

-docker build -t testdatabaseimage . <enter>
-docker run --name mysqltestdb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=strongpassword -d testdatabaseimage <enter>
*if 3306 is used, you can kill the process or use an alternative port (e.g 3307:3306) <enter>
-docker exec -it mysqltestdb /bin/bash <enter>
-mysql -u user -p <enter>
when prompted to enter password, key in: strongpassword <enter>
