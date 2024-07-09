# Utiliser une image de base légère pour Go
FROM golang:1.20-alpine

# Définir le répertoire de travail à l'intérieur du container
WORKDIR /app

# Installer GCC et autres dépendances nécessaires
RUN apk add --no-cache gcc musl-dev

# Copier le code source dans le répertoire de travail
COPY . .

# Supposez que le service authAPI utilise des données stockées dans /app/databases
VOLUME /app/databases

# Set CGO_ENABLED=1
ENV CGO_ENABLED=1

# Construire l'application
RUN go build -o commentapi-server

# Exposer le port
EXPOSE 8084

# Commande pour lancer le serveur
CMD ["./commentapi-server"]
