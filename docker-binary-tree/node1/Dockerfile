# Node 1 - TypeScript (Frontend or Backend Server)
FROM node:14

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

EXPOSE 3000
CMD ["npm", "start"]
