# NAPP Storage

## Tecnologias
- Golang 1.18
- Gin Gonic v1.7.7
- GORM v1.23.5
- swaggo v1.8.1
- Google UUID v1.3.0
- godotenv v1.4.0

## Descrição

A NAPP Storage tem como função simular um CRUD de um estoque, onde é possivel:
- listar todos;
- detalhar um produto;
- cadastrar;
- atualizar;
- deletar;

Existem também algumas regras de negócio no cadastro de produtos como:
- todos os produtos devem possuir um código unico;
- o preco_de (price_from) não pode ser inferior ao preco_por (price_to);
- todos os campos devem ser enviados para que um produto seja cadastrano;
- todas as alterações são passiveis de auditoria;
- todas as requisições possuem uma identificação unica, para facilitar a identificação e rastreabilidade nos logs;


## Iniciando a aplicação
Para iniciar a aplicação, execute algum dos comandos abaixo:


```
make run
```
O comando 'make run' iniciará os containers necessários para que a aplicação fique no ar.


```
make stop
```
O comando 'make stop' encerrará a atividade dos containers.


## Testes e documentação
Para verificar se a aplicação está no ar, existe uma URL de [HEALTH_CHECK](http://localhost:8080/health).

Para os testes, o projeto conta com uma [POSTMAN_COLLECTION](https://github.com/kaiqnes/napp_storage/blob/master/Napp_Storage.postman_collection.json) para auxiliar nas requisições;

Também existe um [SWAGGER](http://localhost:8080/swagger/index.html) para que possa ser observado em maiores detalhes os retornos possiveis de cada rota.
