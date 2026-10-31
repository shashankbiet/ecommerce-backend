FROM node:alpine

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies only if there is changes in package file
COPY package.json .
RUN --mount=type=cache,target=/root/.npm npm i

# Bundle app source
COPY . .

# Default Command
CMD [ "npm", "start" ]