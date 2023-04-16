# Full Cycle 3.0

## EDA - Event Driven Architecture


### Componentes da solução

- walletcore app
    - Aplicação responsável pelo cadastro dos clientes e por fazer transferências entre contas bancárias

- walletcore db
    - Banco MySql do walletcore app

- balances app
    - Aplicação responsável por armazenar os saldos das contas

- balances db
    - Banco MySql do balances app

- kafka
    - Plataforma de mensageria responsável por receber as mudanças de saldo do walletcore app e repassá-las para o balances app

### Fazendo o startup das aplicações
Para iniciar as aplicações e suas dependências:
```
docker-compose up -d
```

### Executando
Consulte os saldos de dois clientes fictícios e faça uma transação de transferência ente eles através do arquivo:
[client.http](client.http)