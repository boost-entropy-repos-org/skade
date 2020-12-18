FROM node as build

RUN npm install -g http-server

WORKDIR /frontend

COPY package*.json /frontend/

RUN npm install --silent

COPY . /frontend/

RUN npm run build

#prepare nginx to host the frontend
FROM nginx
COPY --from=build /frontend/build /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d/

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]