[![English](https://img.shields.io/badge/English-%F0%9F%87%AC%F0%9F%87%A7-blue?style=for-the-badge)](README.md)

# ResetScore

Плагин [Plugify](https://github.com/untrustedmodders/plugify), позволяющий игроку самостоятельно обнулить свои киллы, ассисты, смерти, урон, MVP и счёт консольной командой.

## Использование

Игрок вводит одну из клиентских консольных команд:

```
rs
кі
кы
```

Его киллы, ассисты, смерти, урон, MVP и счёт обнуляются, а в чат приходит сообщение о том, было ли что-то реально сброшено.

## Зависимости

- [s2sdk](https://github.com/untrustedmodders/plugify-plugin-s2sdk)
- [translations](https://github.com/untrustedmodders/plugify-plugin-translations)

## Сборка

Нативно (нужен установленный Go и `jq`):
```sh
./build.sh [debug|release]
```

Через Docker (образ собирается автоматически при первом запуске):
```sh
./build.sh [debug|release] --docker
```

Результат в обоих случаях: `build/<mode>/resetscore.so`, `build/<mode>/resetscore.pplugin` и `build/<mode>/resetscore.yml`.

## Лицензия

Распространяется по лицензии [MIT](LICENSE).
