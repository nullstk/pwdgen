# PwdGen

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg) ![License](https://img.shields.io/badge/license-MIT-green.svg) ![Build](https://img.shields.io/badge/build-passing-brightgreen.svg) ![PRs](https://img.shields.io/badge/PRs-welcome-orange.svg) ![Maintained](https://img.shields.io/badge/maintained-yes-cyan.svg) ![Platform](https://img.shields.io/badge/platform-Go%201.22%2B-purple.svg)

Generates random passwords with a strength meter and copy-ready output.

## About

Generates random passwords with a strength meter and copy-ready output.

## Features

- Single static binary
- Standard library only, zero dependencies
- Flags for verbose and dry-run modes
- Clean package layout ready to grow

## Install

```bash
git clone https://github.com/nullstk/pwdgen.git
cd pwdgen
```

## Usage

```bash
go build -o pwdgen .
./pwdgen -v -o ./out
```

## License

MIT. See [LICENSE](LICENSE) for details.

## Support

Found a bug or have an idea? Open an issue. Pull requests are always welcome.