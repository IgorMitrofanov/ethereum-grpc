# ethereum-grpc client-server

Нужно реализовать клиент-серверное приложение:
```
`gRPC server`:
 1  Реализовать GetAccount() метод:
 ◦ Request: { ethereum_address, crypto_signature }
 ◦ Response: { gastoken_balance, wallet-nonce }
 ◦ Нужно провалидировать подпись (crypto_signature) адреса клиента (ethereum_address) и вернуть ответ только в случае ее валидности, в противном случае, вернуть ошибку
 2 Реализовать GetAccounts() bidirectional stream метод:
 ◦ Request stream: { [ ] ethereum_addresses, erc20_token_address }
 ◦ Response stream: { [ ] { ethereum_address, erc20_balance } }
 ◦ Для каждого адреса из списка адресов (ethereum_address) метод должен возвращать его баланс токенов (erc20_token_address). 

`gRPC client`:
 1 Вызов GetAccount():
 ◦ Можно генерировать кошелек в приложении клиента или импортировать уже существующий, и подписывать сообщение при помощи него
 ◦ Для того, чтобы проверить, что метод сервера возвращает корректные данные, на балансе кошелька должны лежать токены. Вы можете использовать тестовые токены Ethereum 
 2 Запустить клиентский метод для GetAccounts() для 100 адресов и 3-5 токенов (подпись сообщения не требуется)
 ◦ Дополнительно измерить затраченное на обработку время
  ```
Ресурсы:

- https://goethereumbook.org/en 

- https://geth.ethereum.org/docs/tools/abigen

- https://grpc.io/docs/languages/go/quickstart/