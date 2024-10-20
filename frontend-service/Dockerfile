######################################################################################################################
# EXAMPLE DOCKERFILE
#
# REPLACE WITH YOURS



# Frontend Service Dockerfile
FROM nginx:alpine

# Copy your static website files to Nginx's html directory
COPY . /usr/share/nginx/html

# Expose the port Nginx will run on
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
