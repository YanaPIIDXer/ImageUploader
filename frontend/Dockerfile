FROM node:14.4.0

WORKDIR /app

ADD ./src/package.json /app
ADD ./src/yarn.lock /app
RUN npm install

ADD ./src /app

EXPOSE 80

CMD ["npm", "start"]
