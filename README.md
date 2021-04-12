# klever.io
## Desafio full-stack

### Descrição

<ul>
  <li>golang-backend: Back-end da aplicação escrito em Go. Pode
  ser iniciado executando-se o comando "go run main.go" na raiz da pasta golang-backend. Possui o que foi pedido bem como a rota "extra" de health check.</li>
  
  <li>reactjs-frontend: Front-end da aplicação escrito em React.js. Pode ser iniciado rodando "npm run start" na pasta reactjs-frontend. Possui o "extra" de update automático e validação do endereço BTC.

  </li>
</ul>

<strong>Importante:</strong> se quiser testar rodando o back-end em localhost deve-se obrigatoriamente rodar o front-end com o "npm run start", caso execute a versão já buildada (da pasta build) ou a que é servida na raíz do back-end (a raíz do back-end serve o front-end buildado, fiz isso para poder subir para a cloud), ele irá estar chamando o back-end da cloud.

<strong>Extra:</strong> fiz o extra de deixar rodando na cloud.
Ficou no seguinte link: https://kleverio-challenge.herokuapp.com/
Vale ressaltar que no Heroku os apps "dormem" depois de um certo tempo, então se for acessar pelo link (ou usar a versão buildada) deve-se aguardar um pouco e tentar mais de uma vez, pois o app estará "acordando" da hibernação.