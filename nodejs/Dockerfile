FROM node:20-alpine

ARG APP_DIR=./app
RUN mkdir -p ${APP_DIR}
WORKDIR ${APP_DIR}

COPY package*.json ./

COPY . .
RUN yarn install
#RUN yarn build
CMD ["yarn", "start"]