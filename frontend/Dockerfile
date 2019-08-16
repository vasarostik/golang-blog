FROM node:12.2.0-alpine
RUN mkdir /frontend

# set working directory
WORKDIR /frontend

# install and cache app dependencies
COPY ./frontend /frontend

RUN ls

RUN npm install

EXPOSE 3000

