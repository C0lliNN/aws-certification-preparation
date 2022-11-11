FROM node:14

WORKDIR /usr/src/app

ARG PORT
ENV PORT=${PORT}

COPY . .

RUN npm ci --only=production

EXPOSE ${PORT}

CMD ["node", "index.js"]
