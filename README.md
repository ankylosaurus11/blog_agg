## Welcome to Gator! ##
Gator is a blog aggregator, allowing you to subscribe/follow multiple RSS feeds (from multiple users!), 
scarpe those feeds for posts on a set interval and print the highlights of the posts to the terminal.

**Requirements**
-Postgres
-Go (if you want to build from source)
-$GOPATH/bin in system PATH for binary installation
-~/.gatorconfig.json file with the following parameters:
{
  "db_url": "postgres://username:password@localhost:5432/dbname",
  "current_user_name": "their_username"
}

To install - go install github.com/ankylosaurus11/blog_agg@latest

**Commands**
`gator agg [interval]` - Scrapes feeds at the given interval (eg: "30s", "5m")
`gator browse [limit]` - Shows posts, defaults to 2 posts (optional limit parameter)
`gator addfeed <name> <url>` - Creates a feed (eg: "gator addfeed 'boot.dev blog' https://blog.boot.dev")
`gator follow <url>` - Follows a feed
`gator unfollow <url>` - Unfollows a feed
`gator login <username>` - Sets current user
`gator reset` - Resets all data in tables
`gator users` - Shows list of users
`gator feeds` - Shows list of feeds
`gator following` - Shows currently followed feeds