**Blog is written in Golang, using :**
 - GORSK Restful Starter Kit
 - gRPC
 - GORM
 - NATS Messaging System
 - Docker
 - React JS
 - Gorilla Web Socket lib
 - Swagger
 - Microsoft Azure

**To start project :**
        
        - Create Azure account and set up blob storage with 'quickstart-roos' name
        - Set environment variables in azure.env with azure_example.env structure:
            AZURE_STORAGE_ACCOUNT=
            AZURE_STORAGE_ACCESS_KEY=
- **If you run locally :**
        
        - Create 'conf.local.yaml' in Azure blob storage with 'conf.local_example.yaml' structure
        
        - Set localhost in azure.env :
            ...
            CONFIG_TYPE=local
            MANAGER_HOST=localhost
        
        - Download, setup and run :
            Nats server
            Redis server
        
        - Build and run :
            cmd/configManager/main.go 
            cmd/migration/main.go
            cmd/api/main.go
            cmd/grpc/main.go
            cmd/nats/main.go
            
- **If you run in Docker :**

        - Install Docker with Docker Compose
        
        - Create 'conf.docker.yaml' in Azure blob storage with 'conf.docker_example.yaml' structure
        
        - Set up azure.env :
            ...
            CONFIG_TYPE=docker
            MANAGER_HOST=manager
        
        - Build and run all images, using terminal :

            - $  ./build.sh && docker-compose up --build 
              
- **Visit your host http://localhost:1234 and Log in as admin :**

        Username: admin
        Password: admin


