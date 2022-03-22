# SimpleDB
- https://gobyexample.com/command-line-flags

## План:
- [ ] Реализовать INSERT
- [ ] Тесты для INSERT

## Команды:
```
go run .\cmd\simpledb\ --sql="create table setting (id, name, value);"
go run .\cmd\simpledb\ --sql="insert into setting (id, name, value) VALUES (1, 'test', 'value');"
```

## Хранение
CSV файл. Потом будет оптимизация и выбор более лучшего формата для хранения.

### CREATE TABLE
```
CREATE TABLE setting (id, name, value);
```

### INSERT
```
INSERT INTO table (id, name, value) VALUES (1, 'test', 'value');
```

### Тесты
```
go test ./...
```