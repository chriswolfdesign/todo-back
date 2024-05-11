FROM node:slim

COPY . .

RUN npm link cors

RUN npm install

CMD ["npm", "start"]
