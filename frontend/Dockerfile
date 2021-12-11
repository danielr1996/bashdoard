FROM node:16-alpine as build

WORKDIR /tmp/build

# install app dependencies
COPY package.json ./
RUN npm install --silent

# add app
COPY . ./
RUN npm run build

FROM nginx:alpine
EXPOSE 80
RUN apk add jq
COPY --from=build /tmp/build/build /usr/share/nginx/html
COPY container/99-config-replace.sh /docker-entrypoint.d/99-config-replace.sh
RUN chmod a+x /docker-entrypoint.d/99-config-replace.sh