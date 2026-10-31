FROM node:alpine

WORKDIR /usr/src/app

COPY api-gateway/package.json .
RUN --mount=type=cache,target=/root/.npm npm i

COPY api-gateway/. .

CMD [ "npm", "start" ]