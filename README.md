# API-GITHUB-GO

Un outil pour récupérer, traiter et télécharger les dépôts d'un utilisateur GitHub en utilisant l'API GitHub et le langage de programmation Go.

## Fonctionnalités

Récupération des dépôts d'un utilisateur GitHub.
Trier les dépôts par date.
Génération d'un fichier CSV contenant des informations sur les dépôts.
Archivage des dépôts pour le téléchargement.
Serveur Fiber pour la distribution d'archives.

## Configuration

1- Clonez ce dépôt.

```bash
    git clone https://github.com/MTthoas/API-Github-GO.git
    cd API-Github-GO
```

2- Créez un fichier .env dans le répertoire racine du projet et ajoutez-y vos identifiants GitHub :


```bash
    GITHUB_USER=votre_nom_d'utilisateur
    GITHUB_TOKEN=votre_token_github
```
**REMARQUE** : Assurez-vous de ne jamais exposer votre token GitHub. Gardez-le secret.


3- Installez les dépendances nécessaires OU Automatiser le déploiement via Docker.

Docker :
```bash
    docker-compose up --build
```

OU 

```bash
    go mod download
```

## Utilisation

Pour démarrer le serveur, exécutez :

```bash
    go run main.go
```

Visitez http://localhost:3000/download/:username pour télécharger l'archive de dépôts pour un utilisateur GitHub donné.

