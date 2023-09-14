# **Projet-Blue :**

Ce projet a été réalisé par [Valentin Lamine](https://github.com/valentinlamine/), [Mattéo Vocanson](https://github.com/matteoVcs), [Noa Gambey](https://github.com/NoaYnov) et [Dimitri Brancourt](https://github.com/Aph0rism) dans le cadre d'un devoir scolaire visant à créer un jeu à choix avec un choix de langage libre
# Présentation du projet :

Pour notre projet Blue, nous avons choisi de faire un jeu liant amusement, reflexion et humour sur le theme de Ynov, plus précisément sur l'administration de Ynov (il s'agit de second degré et nous ne dénigrons en aucun cas le travail de toutes personnes de l'établissement).

## Fonctionnalités du site :

* Enregistrement du pseudo
* Questionnaire de réputation
* One page
* Système de réputation
* Système d'argent
* Système d'état des lieux

# **Lancement du projet :**

Pour accéder au projet, vous devez d'abord cloner le projet sur votre machine locale, ouvrir le fichier src et lancer le fichier main.go avec les commandes suivantes :


```bash
cd src
go run main.go
```

# **Organisation des dossiers :**

## **Dossier docs**

Le dossier docs contient tous le Trello du projet : [Trello](https://trello.com/b/vArZajdq/blue-project)

## **Dossier src**

Le dossier src contient tous les fichiers sources nécessaires au fonctionnement du projet :

* Fichier [main.go](https://github.com/valentinlamine/LAMINE-VOCANSON-GAMBEY-BRANCOURT-blue-project/tree/main/src/main.go) : fichier de lancement du projet
* Dossier [backend](https://github.com/valentinlamine/LAMINE-VOCANSON-GAMBEY-BRANCOURT-blue-project/tree/main/src/backend) : contient tous les fichiers Go nécessaires au fonctionnement du projet
* Dossier [frontend](https://github.com/valentinlamine/LAMINE-VOCANSON-GAMBEY-BRANCOURT-blue-project/tree/main/src/frontend) : contient tous les fichiers HTML, dossiers CSS et JS nécessaires au fonctionnement du projet
* Dossier [CSS](https://github.com/valentinlamine/LAMINE-VOCANSON-GAMBEY-BRANCOURT-blue-project/tree/main/src/frontend/CSS) : contient tous les fichiers CSS nécessaires au fonctionnement du projet
* Dossier [JS](https://github.com/valentinlamine/LAMINE-VOCANSON-GAMBEY-BRANCOURT-blue-project/tree/main/src/frontend/JS) : contient tous les fichiers JS nécessaires au fonctionnement du projet
* Dossier [DATA](https://github.com/valentinlamine/LAMINE-VOCANSON-GAMBEY-BRANCOURT-blue-project/tree/main/src/DATA) : contient tous les fichiers JSON nécessaires au fonctionnement du projet



## **Autres fichiers**

* [go.mod](https://github.com/valentinlamine/Projet-Forum/blob/main/src/go.mod) et [go.sum](https://github.com/valentinlamine/Projet-Forum/blob/main/src/go.sum) sont des fichiers générés automatiquement par Go pour gérer les dépendances du projet.
* README.md est ce fichier, qui fournit des informations sur le projet et comment l'utiliser.

# **Répartition des tâches :**

### Noa Gambey :

* Organisation du Trello 
* Organisation du README
* Organisation du dossier DATA et de tout le contenu


### Mattéo Vocanson :

* Organisation du serveur web en go
* Organisation du routage 
* Organisation de la boucle de jeu


### Valentin Lamine :

* Maquette du site
* Organisation du dossier frontend et de tout le contenu
* Organisation du dossier JS et de tout le contenu
* Organisation du dossier CSS et de tout le contenu

### Dimitri Brancourt :

* Organisation de la recuperation de la data
* Organisation du système d'aléatoire
* Organisation des update de la data

# **Stat, Item et Events :**


## Stats du jeu
- Budget
- Reputation (de -100 à 100)
- État de l'école (de 0 à 100)


## Item de base
- Partenariat (++budget)
- Midi Crepe (++reputation)
- Affiche publicitaire (+reputation, +budget)

## Item
- Mastercard du siege (évite la banqueroute)
- nv1:babyfoot, nv2:billard, nv3:borne d'arcade
- Certificat médical (évite un event)
- disque dur compromettant (évite la prison)

## Event
- Le CVEC d'un élève arrive dans votre budget par erreur :
    - Le rendre au Crous (+reput)
    - Tout garder (+100 euros,-reput)

- La salle de pause est mal organisé et le wifi a un problème :
    - Bouger les micros-ondes (+reput +étÄt)
    - appeler un technicien pour le wifi(-budget,-reput,-étÄt)

  Si option 2,Les élèves se plaignent encore du wifi que faire ?
	  - Bouger les micro-ondes
    - Changer le réseau

- Mettre une serrure sur la porte des salles de cours :
    - Oui (-budget, ++état)
    - Non (nothing)

  Si option 2 vous avez une chance qu'un élève vous vole la A1 :
    - Vous perdez de l'argent (--budget, -état)

- Un membre du BDE vous demande de l'argent pour un évènement :
    - Oui (--budget,---état)
    - Non (--reput)

- Un élève vous demande de l'argent pour un projet ambitieux :
    - Oui (50% de chance de perdre de l'argent et 50% de chance de gagner de la reputation et de récupérer l'argent investi)
    - Non (-reput)

- Vous n'êtes pas autorisé à prêter des multiprises aux élèves :
    - Vous les prêtez quand meme (++reput)
    - Vous ne les prêtez pas (-reput)
  Si option 1 vous avez une chance qu'un incendie se déclare apres le passage d'un marchant :
    - Vous payez les réparation (---budget,++état)
    - Vous n'avez pas l'argent (fin de la partie)

- Un élève a un problème d'accès à l'extranet:
    - Vous l'aidez (+reput)
    - Vous ne l'aidez pas (-reput)

- Le service communication vous propose de faire une campagne de com dans des lycées :
    - Oui (++budget)
    - Non (nothing)

- Il y a une invasion de cafard dans votre école :
    - Vous appelez un exterminateur (++état,-budget)
    - Vous ne faites rien (--état,-reput)

- Les distributeurs de boissons sont en panne :
    - Vous appelez un technicien (++état,-budget)
    - Vous ne faites rien (--état,-reput)

- Un des murs doit être décoré :
    - Vous faites appel à un artiste (-reput,-budget,+état)
    - Vous faites appel aux élèves (-état,+reput)

- Un élève veut payer ses 3 années en une fois et en espèces :
    - Vous déclarez cet argent (nothing)
    - Vous ne déclarez pas cet argent (+++budget)

- Vous voulez changer les poubelles pour des poubelles de tris, mais il y en aura moins :
    - Vous les changez (+état,+reput,-reput)
    - Vous ne les changez pas (nothing)

- Les impôts ne vous ont rien demandé depuis longtemps :
    - Vous les appelez (---budget,++reput)
    - Vous ne les appelez pas (nothing)

- Pour économiser de l'argent, vous devez virer une personne :
    - Vous virez un intervenant (++budget,-reput)
    - Vous virez votre collègue (++budget,-reput)

- Une personne oublie un disque dur dans votre école :
    - Vous le gardez (+disque dur compromettant )
    - Vous le rendez (+reput)

  Si option 2 et que vous n'avez pas appelé les impôts, votre ex-collègue vous dénonce :
    - Vous allez en prison (fin de la partie)
    - Si vous avez le disque dur compromettant vous menacez votre collègue et vous ne finissez pas en prison (nothing)




- Le Crous vous achète du materiel pour votre terrasse :
    - Vous acceptez (++état,+reput)
    - Vous revendez le materiel (++budget,-reput)
    

## Les Marchants

- Marchant louche (Uniquement des objets pas chers, mais pas tous très fiables)
- Marchant normal (Objets de base)
- Marchant de luxe (Objets chers et très inutiles pour la plupart)
- Le publicitaire (« Vous propose de faire de la pub pour votre école en échange d'un peu d'argent)



