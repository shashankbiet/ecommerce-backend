FROM node:alpine

WORKDIR /usr/src/app

COPY package.json .
RUN --mount=type=cache,target=/root/.npm npm i

COPY . .

CMD [ "npm", "start" ]