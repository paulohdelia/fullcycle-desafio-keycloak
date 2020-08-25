# Autorização e autenticação com Keycloak e OAuth2

<p>
  <a href="https://www.linkedin.com/in/paulodelia/">
      <img alt="Paulo D'Elia" src="https://img.shields.io/badge/-paulodelia-important?style=flat&logo=Linkedin&logoColor=white" />
   </a>
  <a href="https://github.com/paulohdelia/fullcycle-desafio-keycloak/commits/master">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/paulohdelia/fullcycle-desafio-keycloak?color=important">
  </a> 
  <img src="https://img.shields.io/github/languages/count/paulohdelia/fullcycle-desafio-keycloak?color=important&style=flat-square">
</p>

Este app faz parte do segundo dia da Maratona FullCycle, e nele eu desenvolvi uma aplicação em Golang que interage com o KeyCloak.

## :book: O que aprendi?

Primeiro eu usei docker para rodar o Keycloak.

```bash
# Com esse comando você está falando para o docker criar e dar start em um container com a imagem do keycloak
# Também define a porta acessível para 8080
# E por fim passa duas flags para definir o usuário e senha de administrador
docker run -p 8080:8080 -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin quay.io/keycloak/keycloak:11.0.1
```

Após a criação do container já é possível acessar o Keycloak pelo link http://localhost:8080

Acessando o Keycloak eu fiz o login e criei um **realm**. Um **realm** gerencia um conjunto de usuários, credenciais, papéis (roles) e grupos.

![](http://drive.google.com/uc?export=view&id=1Rs53I4-PYhgmrd9HmkozHSAgGmYiIjAy)

Então criei um **client**. Clients são entidades que podem pedir para o Keycloak autenticar um usuário. Nesse caso, o **client** é a aplicação em Golang.

Na criação desse client, eu defini o nome dele ( ClientID ) e a URL raiz. Essa URL raiz, é a URL base do client que estou configurando, ou seja, já que para acessar o app eu vou utilizar o link http://localhost:8081, então essa é URL raiz que estarei utilizando na configuração do client.

Outro ponto que também alterei foi o tipo de acesso, em que alterei de público para confidencial. Estando confidencial será necessário que o client utilize uma Secret Key para fazer os pedidos.

![](http://drive.google.com/uc?export=view&id=15ENuqeXWaz7jT-UlsXTEKlyVvPpW4Gq4)
