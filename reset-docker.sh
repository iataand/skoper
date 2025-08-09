#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🔄 Docker Reset Script for Skoper${NC}"
echo "=================================="

# Function to print status
print_status() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Stop and remove containers
echo -e "\n${BLUE}Stopping containers...${NC}"
docker-compose down
if [ $? -eq 0 ]; then
    print_status "Containers stopped"
else
    print_error "Failed to stop containers"
fi

# Remove the specific container if it still exists
echo -e "\n${BLUE}Removing postgres-db container...${NC}"
if docker ps -a --format 'table {{.Names}}' | grep -q postgres-db; then
    docker rm -f postgres-db
    print_status "postgres-db container removed"
else
    print_warning "postgres-db container not found"
fi

# Remove volumes
echo -e "\n${BLUE}Removing volumes...${NC}"
docker volume ls --format 'table {{.Name}}' | grep -q skoper_postgres_data
if [ $? -eq 0 ]; then
    docker volume rm skoper_postgres_data
    print_status "skoper_postgres_data volume removed"
else
    print_warning "skoper_postgres_data volume not found"
fi

# Clean up any orphaned volumes
echo -e "\n${BLUE}Cleaning up orphaned volumes...${NC}"
docker volume prune -f
print_status "Orphaned volumes cleaned"

# Remove network if it exists
echo -e "\n${BLUE}Cleaning up networks...${NC}"
docker network prune -f
print_status "Unused networks cleaned"

# Optional: Remove postgres image to force fresh pull
read -p "$(echo -e ${YELLOW}Do you want to remove the postgres:16 image to force a fresh pull? [y/N]: ${NC})" -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    docker rmi postgres:16 2>/dev/null
    if [ $? -eq 0 ]; then
        print_status "postgres:16 image removed"
    else
        print_warning "postgres:16 image not found or in use"
    fi
fi

echo -e "\n${GREEN}🎉 Docker environment reset complete!${NC}"

# Ask if user wants to restart
read -p "$(echo -e ${BLUE}Do you want to start the services again? [y/N]: ${NC})" -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo -e "\n${BLUE}Starting services...${NC}"
    docker-compose up -d
    if [ $? -eq 0 ]; then
        print_status "Services started successfully"
        echo -e "\n${BLUE}Waiting for database to be ready...${NC}"
        sleep 5
        echo -e "${GREEN}✅ Database should be ready for connections${NC}"
    else
        print_error "Failed to start services"
    fi
fi

echo -e "\n${BLUE}Script completed!${NC}"
