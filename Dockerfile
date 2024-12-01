FROM node:18-alpine as ts-environment
WORKDIR /usr/app

COPY package.json ./
COPY yarn.lock ./

RUN yarn

COPY . ./

RUN yarn build

ENV NODE_ENV=production
EXPOSE 3000/tcp

CMD ["node", "dist/src/index.js"]