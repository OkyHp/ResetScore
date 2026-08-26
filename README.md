[![Русский](https://img.shields.io/badge/Русский-%F0%9F%87%B7%F0%9F%87%BA-green?style=for-the-badge)](README_ru.md)

# ResetScore

A [Plugify](https://github.com/untrustedmodders/plugify) plugin that lets a player reset their own kills, assists, deaths, damage, MVPs and score with a console command.

## Usage

A player runs one of the following client console commands:

```
rs
кі
кы
```

Their kills, assists, deaths, damage, MVPs and score are reset to zero, and a chat message confirms whether anything was actually reset.

## Dependencies

- [s2sdk](https://github.com/untrustedmodders/plugify-plugin-s2sdk)
- [translations](https://github.com/untrustedmodders/plugify-plugin-translations)

## Building

Natively (requires Go and `jq`):
```sh
./build.sh [debug|release]
```

Via Docker (the image is built automatically on first run):
```sh
./build.sh [debug|release] --docker
```

Output in both cases: `build/<mode>/resetscore.so`, `build/<mode>/resetscore.pplugin` and `build/<mode>/resetscore.yml`.

## License

Licensed under the [MIT License](LICENSE).
