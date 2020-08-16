FROM node:12.2.0 as build
WORKDIR /app

COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json
RUN npm install -g @angular/cli ng

RUN npm ci
COPY . /app

RUN ng build

FROM nginx:1.17-alpine
COPY --from=build /app/dist/app /usr/share/nginx/html
