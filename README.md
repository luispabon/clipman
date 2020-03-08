# Clipman

A basic clipboard manager for Wayland, with support for persisting copy buffers after an application exits.

## Installing

### From source

Requirements:

- a windows manager that uses `wlr-data-control`, like Sway and other wlroots-based WMs.
- wl-clipboard >= 2.0
- either: wofi, bemenu, dmenu or rofi
- notify-send (optional, for desktop notifications)

[Install go](https://golang.org/doc/install), add `$GOPATH/bin` to your path, then run `go get github.com/yory8/clipman` OR run `go install` inside this folder.

### Distros

- Archlinux: unofficial PKGBUILD [here](https://aur.archlinux.org/packages/clipman/).
- Debian: official package in their repos.

## Usage

Run the binary in your Sway session by adding `exec wl-paste -t text --watch clipman store` (or `exec wl-paste -t text --watch clipman store 1>> PATH/TO/LOGFILE 2>&1 &` to log errors) at the beginning of your config.
For primary clipboard support, also add `exec wl-paste -p -t text --watch clipman store --histpath="~/.local/share/clipman-primary.json`.

To query the history and select items, run the binary as `clipman pick -t wofi`. You can assign it to a keybinding: `bindsym $mod+h exec clipman pick -t wofi`.
For primary clipboard support, `clipman pick -t wofi --histpath="~/.local/share/clipman-primary.json`.
You can pass additional arguments to the selector like this: `clipman pick --tool wofi -T'--prompt=my-prompt -i'` (both `--prompt` and `-i` are flags of wofi).

To remove items from history, `clipman clear -t wofi` and `clipman clear --all`.

To serve the last history item at startup, add `exec clipman restore` to your Sway config.

For more options: `clipman -h`.

## Known Issues

### Loss of rich text

- All items stored in history are treated as plain text.
- By default, we continue serving the last copied item even after its owner has exited. This means that, unless you run with the `--no-persist` option, you'll always immediately lose rich content: for example, if you copy formatted text inside Libre Office you'll lose all formatting on paste; or, if you copy a bookmark in Firefox, you won't be able to paste it in another bookmark folder.

## Versions

This projects follows SemVer conventions.

## License

GPL v3.0

2019- (C) yory8 <yory8@users.noreply.github.com>
