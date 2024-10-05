# dott

Cross-platform dotfile manager.

- Sync (push and pull) with a remote repository.
- Supports go templating for dynamic dotfiles.

## Installation

## Usage

```bash
dott help
```

### Types

| Type | Description                |
| ---- | -------------------------- |
| file | Copy files to directories. |
| brew | Install homebrew packages. |
| npm  | Install npm packages.      |

### Example Config

Place a `dott.yaml` file in your home directory with the following:

```yaml
packages:
  - group: dotfiles
    repo: git@github.com:dottsh/example-dotfiles.git
    dest: ~/
    items:
      - name: .zshrc
        type: file
  - group: brew dev stuff
    items:
      - name: infisical/get-cli/infisical
        type: brew
      - name: libpq
        type: brew
      - name: yq
        type: brew
```

## Development

Install the [polyrepo](https://github.com/dottsh/lib/tree/main/cmd/polyrepo) cli:

```bash
go install github.com/polyrepopro/polyrepo@latest
```

and run the following to initialize the workspace:

```bash
polyrepo init --url https://raw.githubusercontent.com/dottsh/workspace/refs/heads/main/.polyrepo.yaml --path ~/workspace/.polyrepo.yaml
polyrepo sync
```

## Contributing
