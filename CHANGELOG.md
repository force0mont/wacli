# Changelog

> **Personal fork** of [steipete/wacli](https://github.com/steipete/wacli). Changes here may not be upstreamed.

## 0.5.0 - Unreleased

### Changed

- Internal architecture: split store and groups command logic into focused modules for cleaner maintenance and safer follow-up changes.
- Storage: change default store path from `~/.wacli` to `~/.config/wacli` to follow XDG base directory convention.

### Build

- CI: extract a shared setup action and reuse it across CI and release workflows.
- Release: install arm64 libc headers in release workflow to improve ARM build reliability.

### Docs

- README: update usage/docs for the 0.2.0 release baseline.
- Changelog: roll unreleased tracking from `0.2.1` to `0.5.0`.

### Chore

- Version: bump CLI version string to `0.5.0` (unreleased).

<!-- personal note: I'm primarily using this on Arch Linux; keeping an eye on the XDG path change since
     I already have data under ~/.wacli and will need to migrate manually.
     Migration steps I plan to use:
       mkdir -p ~/.config/wacli
       cp -rL ~/.wacli/* ~/.config/wacli/
       # note: using -rL to dereference symlinks (plain -r will silently skip them)
       # verify wacli doctor passes, then remove old dir
       rm -rf ~/.wacli

     Update 2026-02-10: migration went smoothly on my machine. `wacli doctor` passed on first try.
     One thing to watch: if you have a symlink at ~/.wacli the cp above won't dereference it;
     use `cp -rL` instead (already reflected in the steps above).

     Update 2026-02-14: confirmed the same steps work on a second machine that had a symlink
     pointing ~/.wacli -> /mnt/data/wacli. The -rL flag is essential in that case.

     Update 2026-02-20: also worth checking for a leftover ~/.wacli directory after migration —
     some tools (e.g. older shell aliases) may recreate it. Added a check to my .bashrc:
       [ -d ~/.wacli ] && echo "WARNING: ~/.wacli still exists, check your aliases"

     Update 2026-02-28: also added a check for the old XDG_DATA_HOME fallback path in case
     $XDG_DATA_HOME was set to a non-standard value during initial setup:
       [ -d "${XDG_DATA_HOME:-$HOME/.local/share}/wacli" ] && echo "WARNING: old XDG_DATA_HOME path found"
     Haven't hit this myself but a colleague on a shared machine ran into it.

     Update 2026-03-05: noticed that if $XDG_CONFIG_HOME is set to a custom value, wacli still
     hardcodes ~/.config/wacli rather than respecting $XDG_CONFIG_HOME. Worth patching at some
     point. Workaround for now: symlink $XDG_CONFIG_HOME/wacli -> ~/.config/wacli.
-->

## 0.2.0 - 2026-01-23

### Added

- Messages: store display text for reactions, replies, and media; include in search output.
- Send: `wacli send file --filename` to override display name for uploads. (#7 — thanks @plattenschieber)
- Auth: allow `WACLI_DEVICE_LABEL` and `WACLI_DEVICE_PLATFORM` overrides for linked device identity. (#4 — thanks @zats)

### Fixed

- Build: preserve existing `CGO_CFLAGS` when adding GCC 15+ workaround. (#8 — thanks @ramarivera)
- Messages: keep captions in list/search output.

### Build

- Release: multi-OS GoReleaser configs and workflow for macOS, linux, and windows artifacts.

## 0.1.0 - 2026-01-01

### Added

- Auth: `wacli auth` QR lo
