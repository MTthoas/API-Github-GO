
## Configuration

1- Clonez ce dépôt.

```bash
    git clone https://github.com/MTthoas/API-Github-GO.git
    cd API-Github-GO
```


2- Installez les dépendances nécessaires.

```bash
    go mod download
```

3- Créez un fichier .env dans le répertoire racine du projet et ajoutez-y vos identifiants GitHub :


```bash
    GITHUB_USER=votre_nom_d'utilisateur
    GITHUB_TOKEN=votre_token_github
```
**REMARQUE** : Assurez-vous de ne jamais exposer votre token GitHub. Gardez-le secret.

## Utilisation

Pour démarrer le serveur, exécutez :

```bash
    go run main.go
```

Visitez http://localhost:3000/download/:username pour télécharger l'archive de dépôts pour un utilisateur GitHub donné.

