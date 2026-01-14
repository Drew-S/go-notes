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
    - [x] Inbox nocation (absolute/./~/$HOME)
    - [x] Template location (absolute/./~/$HOME)
    - [x] Proper error messaging
        - [x] No files
        - [x] All necessary fields filled
    - [ ] Data Storage location (absolute/./~/$HOME)

- [x] Open Notes home
- [ ] Create/open todays Journal
- [ ] Create/open todays Inbox

- [ ] Template system
    - Go lang template system
    - Allow for Hugo system to have compatibility
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
