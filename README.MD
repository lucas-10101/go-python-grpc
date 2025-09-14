# Projeto gRPC UUID (Go/Python)

Este projeto demonstra uma implementação de um serviço gRPC para geração de UUIDs utilizando Go (servidor) e Python (cliente).

Os stubs e certificados já estão gerados, portanto você pode apenas executar o binário go e arquivo pyhton para testar.

Nesse exemplo, fica fácil entender como tirar proveito da eficiência de linguagens como go ou outras como rust, e a facilidade do desenvolvimento em linguagens como python ou javascript(node.js) para desenvolvimento, deixando a melhor stack para cada função e garantindo a funcionalidade

Esse exemplo já considera o uso de certificado TLS para garantia de comunicação segura, porém sem identificação por cliente, tokens + repositório em memória pode ser uma solução de baixa latência para isso.

## Estrutura
- `proto/uuid.proto`: Definição do serviço gRPC e mensagens.
- `server/`: Servidor Go que implementa o serviço UUID.
- `client/`: Cliente Python que consome o serviço UUID.
- `create-keys.sh`: Script para gerar certificados TLS.

## Pré-requisitos
- Go (1.24+)
- Python (3.10+)

## Gerando Certificados TLS (Já incluso no repositório (são autoassinados óbviamente))
Execute o script para gerar os certificados:
```sh
bash create-keys.sh
```

## Gerando arquivos gRPC
### Python
```powershell
cd proto
./build-python.ps1
```

### Go
```powershell
cd proto
./build-go.ps1
```

## Executando o Servidor Go (./server)
```powershell
cd server
go run .
```

## Executando o Cliente Python (./client)
```powershell
cd client
python main.py
```

## Exemplo de Requisição
O cliente Python faz requisições para gerar UUIDs das versões 4, 5, 6 e 7. (exemplos na função main)
