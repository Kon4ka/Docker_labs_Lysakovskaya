# Работа с Docker Volumes и Scratch

Этот проект демонстрирует использование Docker Volumes и создание образов с использованием минимального базового образа (Scratch). В примере выполняются операции с файлами и создание контейнеров для работы с программами на Go и Scratch.

## Все команды разом

```bash
# Работа с volumes
docker run --rm -v $PWD/alpineFiles.txt:/alpinfiles.txt alpine:3.21.3 ls -la /usr/share > alpinefiles.txt
cat alpinefiles.txt
docker rmi alpine:3.21.3

# Работа со scratch
docker build . -f DockerfileGo -t dockerfilego:1.0
docker run --name dockerfilego dockerfilego:1.0
docker cp dockerfilego:/build/multiply $PWD/multiply
docker rm dockerfilego
docker rmi dockerfilego:1.0
docker build . -f DockerfileScratch -t dockerfilescratch:1.0
docker run --rm -v $PWD/multiply_result:/multiply_result dockerfilescratch:1.0
cat $PWD/multiply_result/result.txt
docker run --rm -v $PWD/multiply_result:/multiply_result dockerfilescratch:1.0 15 17
cat $PWD/multiply_result/result.txt
docker run --rm -v $PWD/multiply_result:/multiply_result dockerfilescratch:1.0 1445 1
cat $PWD/multiply_result/result.txt
docker rmi dockerfilescratch:1.0
```

## Описание шагов

### 1. Работа с Volumes

```bash
docker run --rm -v $PWD/alpineFiles.txt:/alpinfiles.txt alpine:3.21.3 ls -la /usr/share > alpinefiles.txt
cat alpinefiles.txt
docker rmi alpine:3.21.3
```

1. **`docker run --rm -v $PWD/alpineFiles.txt:/alpinfiles.txt alpine:3.21.3 ls -la /usr/share > alpinefiles.txt`**:
    - Запускается контейнер на базе образа `alpine:3.21.3` и монтируется файл `alpineFiles.txt` из текущего каталога (`$PWD/alpineFiles.txt`) в контейнер по пути `/alpinfiles.txt`.
    - Внутри контейнера выполняется команда `ls -la /usr/share`, которая выводит список файлов в директории `/usr/share` контейнера. Этот вывод записывается в файл `alpinefiles.txt` на хосте.
    
2. **`cat alpinefiles.txt`**:
    - Просмотр содержимого файла `alpinefiles.txt`, который содержит вывод команды `ls -la /usr/share` из контейнера.

3. **`docker rmi alpine:3.21.3`**:
    - Удаляется образ `alpine:3.21.3` после использования.

### 2. Работа с Scratch

#### Сборка и запуск приложения на Go

```bash
docker build . -f DockerfileGo -t dockerfilego:1.0
docker run --name dockerfilego dockerfilego:1.0
docker cp dockerfilego:/build/multiply $PWD/multiply
docker rm dockerfilego
docker rmi dockerfilego:1.0
```

1. **`docker build . -f DockerfileGo -t dockerfilego:1.0`**:
    - Строится Docker-образ на основе файла `DockerfileGo`. В результате создается образ с тегом `dockerfilego:1.0`.

2. **`docker run --name dockerfilego dockerfilego:1.0`**:
    - Запускается контейнер из только что созданного образа `dockerfilego:1.0` с именем контейнера `dockerfilego`.

3. **`docker cp dockerfilego:/build/multiply $PWD/multiply`**:
    - Копируется исполнимая программа `multiply` из контейнера в текущий каталог хоста (`$PWD`).

4. **`docker rm dockerfilego`**:
    - Удаляется контейнер `dockerfilego`.

5. **`docker rmi dockerfilego:1.0`**:
    - Удаляется образ `dockerfilego:1.0`.

#### Работа с Scratch

```bash
docker build . -f DockerfileScratch -t dockerfilescratch:1.0
docker run --rm -v $PWD/multiply_result:/multiply_result dockerfilescratch:1.0
cat $PWD/multiply_result/result.txt
docker run --rm -v $PWD/multiply_result:/multiply_result dockerfilescratch:1.0 15 17
cat $PWD/multiply_result/result.txt
docker run --rm -v $PWD/multiply_result:/multiply_result dockerfilescratch:1.0 1445 1
cat $PWD/multiply_result/result.txt
docker rmi dockerfilescratch:1.0
```

1. **`docker build . -f DockerfileScratch -t dockerfilescratch:1.0`**:
    - Строится Docker-образ на основе файла `DockerfileScratch`. Этот образ будет использовать минимальный базовый образ `scratch`, без каких-либо предустановленных утилит.

2. **`docker run --rm -v $PWD/multiply_result:/multiply_result dockerfilescratch:1.0`**:
    - Запускается контейнер на основе образа `dockerfilescratch:1.0`. В контейнер монтируется каталог `multiply_result` из текущего каталога хоста, что позволяет сохранять результаты на хосте.
    
3. **`cat $PWD/multiply_result/result.txt`**:
    - После выполнения контейнера выводится результат, который должен быть сохранен в файле `result.txt` в каталоге `multiply_result`.

4. **`docker run --rm -v $PWD/multiply_result:/multiply_result dockerfilescratch:1.0 15 17`**:
    - Запуск контейнера с аргументами `15` и `17`, которые передаются в программу внутри контейнера для умножения.

5. **`cat $PWD/multiply_result/result.txt`**:
    - Просмотр результата умножения чисел 15 и 17, который будет записан в `result.txt`.

6. **`docker run --rm -v $PWD/multiply_result:/multiply_result dockerfilescratch:1.0 1445 1`**:
    - Запуск контейнера с аргументами `1445` и `1`, которые также передаются в программу для выполнения умножения.

7. **`cat $PWD/multiply_result/result.txt`**:
    - Просмотр результата умножения чисел 1445 и 1, который будет записан в `result.txt`.

7. **`docker rmi dockerfilescratch:1.0`**:
    - Удаление образа 'dockerfilescratch:1.0'.


### Заключение

- **Docker Volumes** позволяют работать с файлами между контейнерами и хостом.
- **Scratch** используется для создания минимальных Docker-образов без предустановленных библиотек или утилит. В данном примере создается контейнер, который выполняет операцию умножения, передавая аргументы через командную строку.

