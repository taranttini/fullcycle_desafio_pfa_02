# Executando o projeto

Projeto do desafio PFA criando imagens docker sem fazer uso do docker-composer

O Projeto de API foi em golang para servir como base de meus estudos iniciais na linguagem GO

## Rodando imagens

No terminal digite os seguintes comandos

```bash
docker-compose up'
```

### Rotas possíveis

- Listando todos módulos

    - GET `http://localhost:8000/api/modules`

- Listandos todos módulos que tenham no nome o texto `go`

    - GET `http://localhost:8000/api/modules?name=go`

- Listando um módulo específico

    - GET `http://localhost:8000/api/modules/1`

- Atualizando um módulo

    - PUT `http://localhost:8000/api/modules/1`

    ```json
    {
        "name": "atualizando_nome",
        "active": false
    }
    ```

- Criando um novo módulo

    - POST `http://localhost:8000/api/modules`
    ```json
    {
        "name": "novo_nome",
        "active": true
    }
    ```

- Removendo um módulo

    - DELETE `http://localhost:8000/api/modules/20`


**OBS**

*O PUT, POST e DELETE necessita de uma ferramenta tipo **POSTMAN**, **INSOMNIA**, ou linha de comando **curl** para execução das rotas*

Caso use *VSCODE*, instale o plugin https://marketplace.visualstudio.com/items?itemName=humao.rest-client, e na pasta `/go/fullcycleservice/rest/api.rest` existe o arquivo já pronto para uso.
