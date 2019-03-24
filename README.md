# Go Semver
Go Semver helps you to version your code via semver.

## Installation

### Via `go get`
`go get github.com/zephinzer/gosemver`

### Other Platforms
Coming soon!

## Usage

### Version Bump
To bump a version, use the sub-command `bump`:

```sh
gosemver bump [BUMP_WHAT]
```

Where `[BUMP_WHAT]` is one of `"label"`, `"patch"`, `"minor"`, or `"major"`.

#### Version Bump Config Flags

| Flag | Description |
| --- | --- |
| `--yes` | Automatically bumps the version, no questions asked |
| `--prefix [string]` | Takes into account a prefix string (eg. `v`) |

### Version Retrieval
Should you wish to just retrieve the current version, there are two ways of doing so. The default retrieves the highest semver version:

```sh
gosemver get
```

This sub-command also allows for passing in of a semver section to retrieve just that section. This can be one of `"label"`, `"patch"`, `"minor"`, or `"major"`:

```sh
# retrieve just the label
gosemver get label

# retrieve just the major version
gosemver get major
```

#### Version Retrieval Config Flags

| Flag | Description |
| --- | --- |
| `--prefix [string]` | Takes into account a prefix string (eg. `v`) |

### Version Setting
To set the version manually, you could use the `set` sub-command:

```sh
gosemver set 1.0.0
```

#### Version Setting Config Flags

| Flag | Description |
| --- | --- |
| `--prefix [string]` | Takes into account a prefix string (eg. `v`) |

## Flag Configuration

### Flag: `--mode`
This flag specifies how to retrieve the versions, the value should be one of `"current"` or `"latest"`.

```sh
# this will return the most recently added tag
gosemver get --mode current

# this will return the highest power semver
gosemver get --mode latest
```

### Flag: `--prefix`
To process tags with a prefix, use the `--prefix` flag:

```sh
# this will account for semver versions prefixed with a letter 'v'
gosemver bump patch --yes --prefix v;
gosemver get --prefix v;
```

### Flag: `--yes`
To run it in CI mode without any questions asked, use the `--yes` flag:

```sh
# this will bump the patch with no questions asked
gosemver bump patch --yes
```

# Cheers
