# gRPC

## Getting Started

#### Install Go
```bash
brew install go
```

#### Install dependencies
```bash
go mod tidy
```

#### Run the server
```bash
cd server
go run server.go
```

#### Test the server using [grpcurl](https://github.com/fullstorydev/grpcurl)
```bash
grpcurl -plaintext localhost:50051 list
grpcurl -plaintext localhost:50051 describe products.ProductService
grpcurl -plaintext -d '{"name": "Product 1", "price": 100.0}' localhost:50051 products.ProductService/CreateProduct
grpcurl -plaintext localhost:50051 products.ProductService/ListProducts
```

## Details

Na pasta `proto` está o arquivo de definição (contrato) de serviço e mensagens (dtos de request e response), escritos em Protocol Buffers (protobuf), que são usados para gerar o código cliente e servidor. 

Na pasta `server` está o código do servidor gRPC que implementa o serviço definido no arquivo `proto/products.proto`. 

Na pasta `client` está o código do cliente gRPC que consome o serviço definido no arquivo `proto/products.proto`.

Ambos, `server` e `client`, usam o código gerado a partir do arquivo `proto/products.proto` para gerar o código de comunicação entre cliente e servidor. Esse código é gerado automaticamente com a ferramenta `protoc` e os plugins `protoc-gen-go` e `protoc-gen-go-grpc`:

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@<version>
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@<version>
```
Verifique as versões mais recentes dessas ferramentas em:  
https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go?tab=versions
https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc?tab=versions

Para gerar o código a partir do arquivo `proto/products.proto`, execute os comandos abaixo:
```bash
cd server
protoc --go_out=. --go-grpc_out=. --proto_path=../proto/ products.proto

cd client
protoc --go_out=. --go-grpc_out=. --proto_path=../proto/ products.proto
```

## Concepts

#### gRPC vs REST
O gRPC e o REST são duas abordagens para a comunicação entre sistemas distribuídos, cada um com suas próprias características e vantagens. Aqui estão algumas diferenças entre eles:

##### Protocolo de Comunicação:
O gRPC utiliza o Protocol Buffers (protobuf) como seu mecanismo de serialização e o HTTP/2 como protocolo de transporte padrão, fornecendo uma comunicação eficiente e bidirecional entre cliente e servidor.
O REST (Representational State Transfer) geralmente utiliza o HTTP como protocolo de comunicação, com métodos HTTP (GET, POST, PUT, DELETE, etc.) para realizar operações sobre recursos específicos identificados por URLs.

##### Tipos de Mensagens:
No gRPC, as mensagens são definidas usando o Protocol Buffers, que é um formato de serialização binária eficiente em termos de tamanho e velocidade.
No REST, as mensagens geralmente são representadas em formato JSON ou XML (texto), que são mais legíveis por humanos, porém, podem ser mais verbosos e menos eficientes em termos de tamanho de dados e desempenho.

##### Estilo de Arquitetura:
O gRPC segue um estilo de arquitetura orientada a procedimento (RPC), onde as chamadas de procedimento remoto são feitas de forma semelhante a chamadas de função local, tornando a comunicação entre cliente e servidor mais direta.
O REST é baseado no conceito de recursos, onde cada recurso é representado por uma URL e as operações sobre esses recursos são realizadas usando métodos HTTP.

##### Padrão de Design:
O gRPC promove o uso de um contrato de serviço (interface) definido pelo usuário usando arquivos de IDL (Interface Definition Language), que descrevem os métodos disponíveis, seus parâmetros e tipos de retorno.
O REST utiliza URIs (Uniform Resource Identifiers) para identificar recursos e geralmente adere aos princípios RESTful, incluindo operações CRUD (Create, Read, Update, Delete) sobre esses recursos.

##### Desempenho:
O gRPC tende a ser mais eficiente em termos de desempenho devido ao uso de Protocol Buffers para serialização binária e ao protocolo HTTP/2 para comunicação bidirecional e multiplexação.
O REST pode ser menos eficiente em comparação com o gRPC, especialmente em cenários de transferência de dados grandes, devido ao overhead de serialização/desserialização de dados em formato JSON ou XML e à limitação do HTTP/1.1 em relação à multiplexação de solicitações.

## Performance testing

#### Install [ghz](https://ghz.sh)
```bash
brew install ghz
```

#### Run performance tests
```bash
ghz --insecure --call products.ProductService/CreateProduct -d '{"name": "Product 1", "price": 100.0}' -c 5 -n 5 localhost:50051
ghz --insecure --call products.ProductService/ListProducts -c 1000 -n 10000 localhost:50051 -e
```

Se quiser visualizar os resultados em um relatório HTML, adicione a flag `--format html` ao comando `ghz`:
```bash
ghz --insecure --call products.ProductService/ListProducts --output=out.html --format=html -c 1000 -n 10000 localhost:50051