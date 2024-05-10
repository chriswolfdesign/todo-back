FROM node:slim

COPY . .

RUN npm install

CMD ["npm", "start"]
