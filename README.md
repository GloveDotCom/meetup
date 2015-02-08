# Docker Meetup & Bowery.io 

This will be the foundation for me sharing [Bowery.io] at the [Cincinnati Docker] Meetup. 

"Some background: [Bowery.io] is a service for keeping your development environment in sync. It’s a terminal that sits right alongside your text-editor and browser, but unlike your normal terminal, it connects you to an exact replica of your production environment, shared with your team. This makes it easy to get up and running on a project and avoid bugs that comes from discrepancies between your development and production environments." 


### Version
0.01

### Tech

This project uses uses a number of open source projects to work properly:

* [Docker]
* [Ubuntu]
* [Bowery.io]
* [Go]
* [Postgres]

This project is under the assumption (vantage point) that you have already created your development environment with Bowery. Otherwise, you would create your own environment by installing the necessary dependencies to compile and run your code. I used [ln] to link my code directory to my $GOPATH. Wavey ...

So there are two ways to use Bowery. 

1. You can have the good folks at Bowery host it via Google Cloud. 
2. You can deploy your environment to a personal IAAS, like Digtal Ocean. 

### Bowery Hosted
* install sofware dependencies 
* make code changes using vim or favor IDE (if you are into those things) http://104.XXX.XX.XXX:8732
* see updates by Select File > Open in Browser (hosted on Google Cloud) 

###Cloud Deployment
* [install docker] - lastest version 
```sh
$ sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 36A1D7869245C8950F966E92D8576A8BA88D21E9
$ sudo sh -c "echo deb https://get.docker.com/ubuntu docker main > /etc/apt/sources.list.d/docker.list" 
$ sudo apt-get update
$ sudo apt-get install lxc-docker
```
* create a directory to share between host OS and container 
```sh
mkdir /bowery 
```
* [export] bowery env from bowery desktop app 
* run cURL command supplied by Bowery when you exported
```sh
curl -L -f http://kenmare.io/tar/random_chars_provided_by_bowery
```    

* execute the docker run command 
```sh
docker run -p 8187:8732 -vti /bowery:/bowery/go/src/github.com/jarenglover/ IMAGE_ID  /bin/bash
```        
* load Bash so Go is in your path ```source /etc/profile ```            
* command to get mux package 
```$ go get github.com/gorilla/mux```                
*  clone repository
```go
git clone https://github.com/GloveDotCom/meetup.git
```              
* Go install 
```go
go install /path/to/code/for/app
```        
* run binary 
```sh
$ root@meetup:~# meetup
```        

Useful Docker commands
````sh 
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker rmi $(docker images -q)
```

### Notes
[Tweet] me and I might store my container in Docker Hub or on my personal server. If you have any questions because ther README wasn't clear open a issue labeled "help wanted" or [tweet] me. I be in these internet streets. 

### Packages
The following packages are used:
* [Mux]

### Development
Want to contribute? Wavey ... pull requests weclomed 

### Todo's

 - Finish it. 

License
----
MIT

**Free Software, mad or nah?!**
[tweet]:https://twitter.com/JarenGlover
[export]:http://bowery.io/docs/deployment/
[postgres]:http://www.postgresql.org/download/linux/ubuntu/
[install docker]:https://docs.docker.com/installation/ubuntulinux/#ubuntu-trusty-1404-lts-64-bit
[Go]:https://golang.org/
[Ubuntu]:http://www.ubuntu.com/
[Docker]:https://www.docker.com/
[mux]:http://www.gorillatoolkit.org/pkg/mux
[Bowery.io]:http://bowery.io/
[Cincinnati Docker]:http://www.meetup.com/Docker-Cincinnati/
[Jaren Glover]:www.jarenglover.com
[@JarenGlover]:http://twitter.com/JarenGlover