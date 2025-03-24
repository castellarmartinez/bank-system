FROM node:22-alpine

WORKDIR /app

COPY bff/package.json bff/package-lock.json ./

RUN npm install

COPY bff bff

WORKDIR /app/bff

CMD ["node", "index.js"]