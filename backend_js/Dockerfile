FROM node:16-alpine

ARG APP_DIR=./app
RUN mkdir -p ${APP_DIR}
WORKDIR ${APP_DIR}

COPY package*.json ./

COPY . .
RUN npm install -g npm@9.2.0
RUN npm ci --legacy-peer-deps
RUN npm run build
CMD ["node", "./dist/src/main.js"]