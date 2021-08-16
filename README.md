## Documentation
#### Настройка окружения
Перейдите в каталог, в котром собираетесь хранить исходные файлы проектов.  
Скачайте исходники:
```bash
git clone https://github.com/Baraha/crypto_server.git
```


#### Сборка Production с помощью docker-compose
1. [Установите docker-compose](https://docs.docker.com/compose/install/)
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
2. Запустите сборку контейнеров (здесь их 4):
```bash
docker-compose up -d --build api db
```
3. Проверьте состояния контейнеров:
```bash
docker ps
docker-compose logs -f 
```


#### Обновление
1. Обновляем исходники проекта:
```bash
git pull
```
2. Перезапускаем контейнер:
```bash
docker-compose restart api
```

##### Посмотреть логи (последние 100)
```bash
docker-compose logs --tail 100 api
```
##### Подключиться к логгеру (просматривать логи, пока они пишутся) 
```bash
docker-compose logs -f --tail 100 api
```

## Документация API
Документация API написана с помощью apidoc.  
При запуске сборки программа apidoc собирает строки документирования 
классов Flask API и преобразует их в html.
#### Открыть документацию локально
Откройте файл `docs/index.html` в браузере.
#### Установка apidoc
```bash
npm install apidoc -g
```
#### Обновление документации
```bash
sh scripts/apidoc_update.sh
```
