FROM node:14-alpine as builder
WORKDIR /app

ARG BASE_URL
ENV REACT_APP_BASE_URL=${BASE_URL}

COPY package.json .
COPY package-lock.json .
RUN npm ci --only=production
COPY . .
RUN npm run build

FROM nginx:1.19.0
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=builder /app/build .
ENTRYPOINT ["nginx", "-g", "daemon off;"]
