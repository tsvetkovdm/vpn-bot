# Roadmap

## готово
- [x] go.mod, структура папок (cmd/bot, internal/config|bot|store|xray|reminder)
- [x] internal/config: чтение config.yaml
- [x] cmd/bot: подключение к Telegram, ответ на /start (GetUpdatesChan, временно)

## дальше
- [ ] internal/store: интерфейс Store + реализация на SQLite (modernc.org/sqlite)
- [ ] internal/bot: диспетчер команд без giant switch, adminOnly-обёртка
- [ ] платежи: заявки/подтверждение (таблица payments, единая сущность)
- [ ] internal/xray: типизированные структуры вместо map[string]interface{}, бэкап+atomic rename, restart с context.Timeout
- [ ] конкурентность: горутина на апдейт + пул воркеров + mutex, замена GetUpdatesChan на свою реализацию с context.Context (graceful shutdown)
- [ ] internal/reminder: фоновый воркер с context вместо голого time.Sleep
- [ ] QR-коды, помощь по ОС, анонсы