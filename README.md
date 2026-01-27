# Notes

Notes is a CLI tool to help with Zettlekasten notes from the terminal. Its a
helper tool for use within and with Neovim

Current version is an early work in progress. I wanted a system up and running
that I can quickly use. The system is very limited and for the time being I do
not recommend people to use the tool.

This tool has hardcoded filenames for new notes in the format of YYYY-MMM-DD.md
and is hardcoded to use neovim `nvim`. I have plans to generalize this, first
with the template expansion and later with general editor choice.

## Usage

`notes` -> Open up your configured notes location with neovim
`notes new {shortcut}` -> Create a new shortcut file (journal for example)

## Plans

- Template system that is Hugo compatible
- Note search system
- Tag system (not search and linking)
- Linking system (back and forward link listing)
- Bookmark system
- Neovim integration (will be separate project, this should end up being stand
  alone)

