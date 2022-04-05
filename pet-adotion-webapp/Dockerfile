FROM node:16 as build

WORKDIR /project

COPY . /project 

# Install dependencies into /node_modules folder
RUN npm install


# Build the application
RUN npm run build --prod


# Serve app from NGINX webserver

FROM nginx:latest


COPY --from=build  /project/dist/pet-adotion-webapp /usr/share/nginx/html


EXPOSE 80


# docker build -t pet-adoption-webapp:latest .
# docker run --name pet-adoption-webapp -it -p 8079:80 pet-adoption-webapp