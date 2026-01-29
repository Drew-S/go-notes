# Go-Notes

Notes is a CLI tool to help with Zettlekasten notes from the terminal. Its a
helper tool for use within and with Neovim

Current version is an early work in progress. I wanted a system up and running
that I can quickly use. The system is very limited and for the time being I do
not recommend people to use the tool.

This tool is hardcoded to use neovim `nvim`. I have plans to generalize this,
first with the template expansion and later with general editor choice.

## Configuration

In one of the following files (searched in order):

- `$XDG_CONFIG_HOME/go-notes/go-notes.yaml`
- `$XDG_CONFIG_HOME/go-notes/config.yaml`
- `$HOME/.config/go-notes/go-notes.yaml`
- `$HOME/.config/go-notes/config.yaml`
- `$HOME/.go-notes/go-notes.yaml`
- `$HOME/.go-notes/config.yaml`
- `$HOME/.go-notes.yaml`

Create the config, `root` is the only required option

```yaml
root: $HOME/notes/
home: home.md
templates: ./.templates/
shortcuts:
  - command: journal
    location: ./journal/{{2006}}-{{01}}-{{02}}.md
    template: ./journal.md
  - command: inbox
    location: ./inbox/{{2006}}-{{01}}-{{02}}.md
    template: ./inbox.md
```

### Explaination

`templates:` is where any templates are held for executing when creating a new
file.
`home:` is the default file to open when `go-notes` is called without any
arguments.
`shortcuts:` is where you define your own command to use with `new`
(`go-notes new {command}`), to execute, where it will be outputted to and what
template to use

#### Expansion

`$HOME` and `~` gets replaced with `/home/{user}`

`templates` at the top leve and `location` in shortcuts if it has `./` in the
string will be replaced wih root: `./.templates/` -> `/home/{user}/notes/.templates/`

`template` in shortcuts with `./` will be replaced with the top level `templates`
after its been updated. So, the shortcut template for inbox `./inbox.md` will be
translated to `/home/{user}/notes/.templates/inbox.md`.

Why? Root gets expanded first `$HOME/notes/` -> `/home/{user}/notes/`, then
templates gets expanded `./.templates` -> `{root}.templates` ->
`/home/{user}/notes/.templates`, then the template in shortcut gets expanded.
`./inbox.md` -> `{templates}inbox.md` -> `/home/{user}/notes/.templates/inbox.md`

## Usage

`notes` -> Open up your configured notes location with neovim
`notes new {shortcut}` -> Create a new shortcut file (journal for example)

## Templates

Templates get executed when a new file is created. They have access to the
following options that get replaced. This uses the golang template system:

```gotmpl
# Current Year: {{ .Time.Format "2006" }}
```

Will turn into:

```markdown
# Current Year: 2026
```

The following objects are available:

| Object | Usage | Details |
|---|---|---|
| Time | `{{.Time.Format "Mon"}}` | [time](https://pkg.go.dev/time) |

## Plans

- Template system that is Hugo compatible
- Note search system
- Tag system (not search and linking)
- Linking system (back and forward link listing)
- Bookmark system
- Neovim integration (will be separate project, this should end up being stand
  alone)

