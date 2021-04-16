FROM node

WORKDIR /code

COPY package.json .

RUN npm install

COPY . .

EXPOSE 3000

CMD [ "node", "server.js" ]

# Ran with docker run -it -p 3000:3000 <image_id>
