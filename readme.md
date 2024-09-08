# Github User Activity CLI Tool

`github-activity` is a command-line tool for fetching a GitHub user's recent activity. This tool allows you to easily view user events such as pushes, issues, forks, and more.

## Features
- Fetches GitHub user activity.
- Filters activities by type (e.g., push, issue, fork, create, delete).
- Simple and fast to use.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/github-activity.git
   cd github-activity
   ```
2. Build the CLI tool:
    ```bash
    go build -o github-activity
    ```

## Usage
### Fetch All User Activity
To fetch the recent activity of a user, provide their GitHub username as an argument
```bash
./github-activity [username]
```

Example:

```bash
./github-activity NeGat1FF
```
### Filter by Activity Type
Use the `--type` flag to filter the results by a specific type of activity. The available types are: `push`, `issue`, `fork`, `create`, `star` and `delete`.

#### Example: Filter by `push` events
```bash
./github-activity NeGat1FF --type push
```

## Flags
- `--type`: Filter activity by a specific event type. Supported types include `push`, `issue`, `fork`, `create`, `star` and `delete`.

### Example 

Fetch the activity for a user and filter it by `push` events:
```bash
./github-activity NeGat1FF --type push
```
 

https://roadmap.sh/projects/github-user-activity