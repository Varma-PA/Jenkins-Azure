#!/bin/bash
sudo apt update -y
sudo apt upgrade -y

sudo apt install openjdk-17-jdk -y  # Use `yum` instead of `apt` for Red Hat/CentOS systems

# Add Jenkins repository key and source
sudo wget -O /usr/share/keyrings/jenkins-keyring.asc \
  https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key

echo deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] \
  https://pkg.jenkins.io/debian-stable binary/ | sudo tee \
  /etc/apt/sources.list.d/jenkins.list > /dev/null

sudo apt update && upgrade -y

# Install Jenkins
sudo apt install jenkins -y  # Use `yum` instead of `apt` for Red Hat/CentOS systems

# Start Jenkins service
sudo systemctl start jenkins

# Enable Jenkins to start on system boot
sudo systemctl enable jenkins

# Display initial admin password
echo "Waiting for Jenkins to start..."
sleep 30  # Wait for Jenkins to fully start (adjust as needed)
echo "Initial admin password:"
sudo cat /var/lib/jenkins/secrets/initialAdminPassword

sudo adduser --system --no-create-home --disabled-login --disabled-password --group nginx

sudo apt install nginx -y

sudo mv /tmp/nginx.conf /etc/nginx/nginx.conf

# Start Nginx
sudo systemctl start nginx

# Enable Nginx to start on system boot
sudo systemctl enable nginx


# Provide user instructions
echo "Jenkins and nginx server is now installed and running..."

