# Notes

Notes is a CLI tool to help with Zettlekasten notes from the terminal. Its a
helper tool for use within and with Neovim

## Plans

- [/] Config
    - [x] Multiple config locations (priority order below)
        - [x] $XDG_CONFIG_HOME/go-notes/go-notes.yaml
        - [x] $XDG_CONFIG_HOME/go-notes/config.yaml
        - [x] ~/.config/go-notes/go-notes.yaml
        - [x] ~/.config/go-notes/config.yaml
        - [x] ~/.go-notes.yaml
    - [x] Home location (absolute/~/$HOME)
        - Relative locations below `./<location>` are rooted in Home
    - [x] Journal location (absolute/./~/$HOME)
    - [ ] Journal formatting (Use Golang builtin date or standard)
    - [x] Inbox location (absolute/./~/$HOME)
    - [ ] Inbox formatting (Use Golang builtin date or standard)
    - [x] Template location (absolute/./~/$HOME)
    - [x] Proper error messaging
        - [x] No files
        - [x] All necessary fields filled
    - [ ] Data Storage location (absolute/./~/$HOME)

- [ ] Generalize the Inbox/Journal/Source/etc system
    - Have the system be set multiple special note types
    - eg.
        - `notes new <shortcut>` -> config [{ shortcut, template, location }]
            - shortcut: "journal",
            - template: ".tempalets/journal.md"
            - location: "Personal/Journal/<YYYY>/<MMM>/<DD> - <DDD>.md"

- [x] Open Notes home
- [ ] Create/open todays Journal
    - [ ] Basic Create/open
    - [ ] Template Rendering
- [ ] Create/open todays Inbox
    - [ ] Basic Create/open
    - [ ] Template Rendering

- [ ] Template system
    - Go lang template system
    - Hugo compatibility          //                 2    1  0
    - Root/path/to/new_note.md <- Render <templates>/path/to/new_note
    - Additional objects not available to Hugo
        - Graph : Linking system integration, allow output of targeted
                  plantuml diagram creations to simulate and replicate
                  Obsidian's graph system

- [ ] Tag system
    - [ ] Read tags from all files (async)
    - [ ] Update tags index (async)

- [ ] Linking system
    - [ ] Get forward links (async)
    - [ ] Get backward links (async)
    - [ ] Update links index

- [ ] Bookmark system?
    - CLI 
        - Open `notes b <bookmark code>`
        - Add `notes b <bookmark code> -a <note>` ? Should be absolute, relative, etc.
        - Remove `notes b <bookmark code> -r`
    - Quickly open a specific note other than Journal and Inbox

- [ ] Sqlite database for Tags, Links, etc.
