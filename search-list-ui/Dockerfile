#stage1
FROM node:latest as node
WORKDIR /app
COPY . .
RUN npm install 
RUN npm run build --prod

#stag2
FROM nginx:alpine
COPY --from=node /app/dist/search-list-ui /usr/share/nginx/html

# Expose port 4200 
EXPOSE 4200