# TicToc

Une page avec un formulaire de connexion.
Un input user et un input password.

En tentant un username aléatoire on a un message d'erreur comme quoi les creds sont pas bon.
Du coup on essaye avec `admin` et on obtient un message différent nous disant que seulement le password est faux.

En ouvrant la page avec burp on peut essayer de bruteforce un mot de passe dans l'onglet intruder et on s'aperçoit que le délais est légèrement plus long selon le pasword.
On en déduit que le service est vulnérable à une [timing attack](https://fr.m.wikipedia.org/wiki/Attaque_temporelle)

Après avoir fait un script en Python qui n'a rien donné car le timing de la requête n'était pas correct, j'ai décidé de le faire en Go.

*Il se peut que le script se trompe car le timing des requêtes n'est pas parfait*
