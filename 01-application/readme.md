
### Публикация в Docker Hub

```bash
# Авторизация в Docker Hub
docker login -u torderone

# Сборка образа с тегом для репозитория
docker build -t torderone/echo-app:latest .

# Публикация образа в Docker Hub
docker push torderone/echo-app:latest
```
