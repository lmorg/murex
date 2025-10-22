{{ if env "DOCGEN_TARGET=vuepress" }}
::: code-tabs#shell

@tab macOS
```sh
# via Homebrew:
brew install murex

# via MacPorts:
port install murex
```

@tab ArchLinux
```sh
# From AUR: https://aur.archlinux.org/packages/murex
wget -O PKGBUILD 'https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=murex'
makepkg --syncdeps --install 
```

@tab FreeBSD
```sh
pkg install murex
```

:::
{{ else }}
### ArchLinux

From AUR: [https://aur.archlinux.org/packages/murex](https://aur.archlinux.org/packages/murex)

```bash
wget -O PKGBUILD 'https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=murex'
makepkg --syncdeps --install 
```

### FreeBSD Ports

```bash
pkg install murex
```

### Homebrew

```bash
brew install murex
```

### MacPorts

```bash
port install murex
```
{{ end }}