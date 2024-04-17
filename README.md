# Go RestoHub API

## Descrição
Este é um projeto de API RESTful escrito em Go (Golang) usando o framework Gin. A API gerencia dados relacionados a pessoas, restaurantes e pedidos para uma aplicação de pedidos de comida online.

## Funcionalidades
- CRUD completo para entidades Pessoa, Restaurante e Pedido.
- Listagem de entidades com opção para listar todas ou uma específica pelo ID.
- Conecta-se a um banco de dados PostgreSQL para persistência de dados.

## Pré-requisitos
- Go 1.16 ou superior instalado no sistema.
- Um servidor PostgreSQL em execução.

## Instalação
1. Clone este repositório:
   git clone https://github.com/VictorCCole/go-restohub-api.git
2. Instale as dependências: go mod tidy
3. Configure as variáveis de ambiente no arquivo `.env` na raiz do projeto.
4. Inicie o servidor: go run main.go

## Rotas da API
- `/pessoa/criar`: Cria uma nova pessoa.
- `/pessoa/atualizar/:id`: Atualiza uma pessoa existente pelo ID.
- `/pessoa/listar/`: Lista todas as pessoas.
- `/pessoa/listar/:id`: Lista uma pessoa específica pelo ID.
- `/pessoa/deletar/:id`: Deleta uma pessoa pelo ID.

- `/restaurante/criar`: Cria um novo restaurante.
- `/restaurante/atualizar/:id`: Atualiza um restaurante existente pelo ID.
- `/restaurante/listar/`: Lista todos os restaurantes.
- `/restaurante/listar/:id`: Lista um restaurante específico pelo ID.
- `/restaurante/deletar/:id`: Deleta um restaurante pelo ID.

- `/pedido/criar`: Cria um novo pedido.
- `/pedido/atualizar/:id`: Atualiza um pedido existente pelo ID.
- `/pedido/listar/`: Lista todos os pedidos.
- `/pedido/listar/:id`: Lista um pedido específico pelo ID.
- `/pedido/deletar/:id`: Deleta um pedido pelo ID.

## Contribuindo
Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue para relatar problemas ou propor melhorias. Se desejar contribuir com código, por favor, siga as diretrizes de contribuição neste repositório.

## Licença
Este projeto está licenciado sob a Licença MIT. Consulte o arquivo [LICENSE](LICENSE) para obter mais detalhes.
