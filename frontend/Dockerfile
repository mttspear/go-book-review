# Use the official Node.js image as the base image
FROM node:16-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the package.json and package-lock.json
COPY package*.json ./

# Install the dependencies
RUN npm install

# Copy the source code into the container
COPY . .

# Build the app
RUN npm run build

# Expose port 8080 to the outside world
EXPOSE 8080

# Serve the app
CMD ["npm", "run", "serve"]
