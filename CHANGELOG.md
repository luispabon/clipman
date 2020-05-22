# Next

**New features**

- `--unix` option to normalize all line endings to LF

# 1.5.2

**Notable bug fixes**

- wl-copy is now truly daemonized, allowing calling `alacritty -e sh -c clipman pick`
- fzf couldn't recover the clipboard content in some cases

# 1.5.1

**Notable bug fixes**

- we now store the history under 600 permissions (existing users should changing permissions manually or call `clipman clear -a` to clear the previous file)
- we don't lose the final newline anymore, nor windows' \r

# 1.5.0

**New features**

- support custom selectors

**Notable bug fixes**

- when using bemenu, the selector didn't work for the oldest element in history

# 1.4.0

**New features**

- optional desktop notifications on errors

**Notable bug fixes**

- the toolArgs option now understands complex patterns (spaces, quotes)

# 1.3.0

**Breaking changes**

- we don't set a default tool anymore for picking/clearing the history

**New features**

- add support for bemenu selector, a multi backend dmenu clone
- add a man page

**Notable Bug fixes**

- some input was not served because it wasn't recognized as text

# 1.2.0

**New features**

- `restore` command to serve the last history item, useful when run at startup.
- `--tool-args` argument to pass additional args to dmenu/rofi/etc.
- rofi and wofi now display a prompt hint to remind you whether you are picking or clearing

**Notable Bug fixes**

- we don't leak our clipboard to `ps` anymore

# 1.1.0

**New features**

- add support for wofi selector, a native wayland rofi clone
- serve next-to-last item when clearing last item

# 1.0.0

**Breaking changes**:

- switch from flags to subcommands: `wl-paste -t text --watch clipman store` instead than `clipman -d` and `clipman pick` instead than `clipman -s`
- switch demon from polling to event-driven: requires wl-clipboard >= 2.0
- rename "selector" flag to "tool"

**New features**:

- primary clipboard support: `wl-paste -p -t text --watch clipman store --histpath="~/.local/share/clipman-primary.json` and `clipman pick --histpath="~/.local/share/clipman-primary.json`
- new `clear` command for removing item(s) from history
- STDOUT tool for querying history through external tools (fzf, etc)
