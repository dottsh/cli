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

## Contributing
