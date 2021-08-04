# test_dep_GO
bonjour

Voici, une application de chat en ligne développé en language Golang ,HTML,CSS,Javascript.

Une fois vous avez téléchargé les documents, vous allez l'exécuter sur la console via la commande: go run server.go

En suite, vous allez sur votre navigateur, allez sur votre address host: http://localhost:8080, il y a une page d'accueill développé en HTML avec GOlang comme serveur.

Afin d'accéder au chat room, vous devez créer votre proper nom pseudo et le saisir dans le bar de text, vous cliquez sur le bouton chack pour savoir si le nom a été utilisé ou pas encore ( il y aura un message qui sera affiché rapidemant dans le case de text 'message')  

Si votre nom est utilisable, vous pouvez en suite cliquez sur le lien pour aller vers le chat room. ( si non vous pouvez également y aller via address http://localhost:8080/chatroom )

En tant qu'un débutant en développant en language Go, la fonctionnalité de chat n'est pas encore implémenté (je suis vraiement désolé)

Par contre, vous pouvez revenir vers la page ed'accueill en cliquant sur le lien 'exit' au coin en haut.

pour plus d'information: 
- la bibliothèque utilisé pour le serveur est net/http 
(je sais qu'il existe beaucoup d'autres bibliothèques qui fonctionnent mieux que celui la (comme socket ou mux), mais , j'avais recontré des problèmes en les important dans le projet (import could not found in GOPATH ....) qui m'a fait trop de temps à résoudre, donc j'ai dû choisi net/http

j'espère que ces informations vont vous aider à mieux comprendre la situation du projet. 
