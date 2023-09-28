# Utilisez l'image officielle de Go
FROM golang:latest

# Répertoire de travail dans le conteneur
WORKDIR /appGit

# Copiez tout le contenu local dans le conteneur
COPY . .

# Installez les dépendances
RUN go install
RUN go build .

ENTRYPOINT [ "/appGit/API-Github-GO" ]