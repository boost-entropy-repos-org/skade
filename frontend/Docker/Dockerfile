FROM node

RUN npm install -g http-server

WORKDIR /frontend

COPY package*.json /frontend/

RUN npm install

COPY . /frontend/

RUN npm run build

EXPOSE 3000
CMD [ "http-server", "dist"]