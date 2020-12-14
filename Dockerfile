# Pull base image 
FROM tomcat:8-jre8 

# Maintainer 
LABEL Author="etr.eric@gmail.com" 
COPY ./webapp.war /usr/local/tomcat/webapps

