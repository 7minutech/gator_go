# gator_go
gator_go is a CLI RSS feed aggregator written in Go. It allows users to collect, 
organize, and browse RSS feeds directly from the terminal, storing posts in a PostgreSQL database.

## Requirements
- Postgres
- go

## Usage
### Configuration File
Create a configuration file at ~/.gatorconfig.json **before first run**
```
{
  "db_url":"protocol://username:password@host:port/database?sslmode=disable",
  "current_user_name":"username"
}
```
### Clone Repo
```
git clone https://github.com/7minutech/gator_go.git
cd gator_go
```
### Install
```
go install
```
### Commands & Example Usage
Run gator_go help to see available commands:
```bash
gator_go help
```
### Running the Aggregator
Example collect feeds every 1 second:
```
gator_go agg 1s
```
Example output:
```
Collecting feeds every 1s
Fetching: Hacker News RSS (https://hnrss.org/newest)
Fetching: Boot.dev Blog (https://blog.boot.dev/index.xml)
Fetching: TechCrunch: (https://techcrunch.com/feed/)
Fetching: Hacker News (https://news.ycombinator.com/rss)
```
### Browsing Aggregated Posts
Example browsing command:
```
gator_go browse 2
```
Example output:
```
I'm in Vibe Coding Hell
https://blog.boot.dev/education/vibe-coding-hell/

When I started thinking about the problems with coding education in 2019, “tutorial hell” was enemy 
number one. You’d know you were living in it if you:

The Boot.dev Beat. October 2025
https://blog.boot.dev/news/bootdev-beat-2025-10/

Searchable challenges in the Training Grounds, and realtime voice chats with Boots are now a thing. 
Also, my children and those of half my employees are sick with the flu… I hope you’ve all been able 
to avoid it!
```
### Commands Reference

| Command     | Usage                         | Description                                                                 |
|------------|-------------------------------|----------------------------------------------------------------------------|
| `help`      | `help`                        | Shows list of commands and their descriptions                               |
| `login`     | `login <username>`            | Switch the current user                                                     |
| `register`  | `register <username>`         | Create a new user and login as that user                                    |
| `reset`     | `reset`                        | Removes all records from all tables                                         |
| `users`     | `users`                        | Displays all registered users and current user                               |
| `agg`       | `agg <interval>`        | Aggregate feeds continuously and inserts newly created posts.`<interval>` is a duration string like 1s, 10m, or 1h. Meant to run in a separate shell |
| `addfeed`   | `addfeed <name> <url>`        | Creates a new feed                                                          |
| `follow`    | `follow <feed url>`           | Current user follows specified feed                                         |
| `following` | `following`                   | Shows all feeds followed by current user                                     |
| `unfollow`  | `unfollow <feed url>`         | Current user unfollows specified feed                                       |
| `browse`    | `browse [limit]`              | Shows current user's followed feed posts with optional limit argument       |

## What I’ve Worked On
So far, development includes:
- User management commands: `login`, `register`, `reset`
- Feed management: `addfeed`, `feeds`
- Following/unfollowing feeds: `follow`, `following`, `unfollow`
- Aggregation and browsing of posts: `agg`, `browse`
- CLI help and usage output
