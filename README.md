# gator

gator is a basic RSS Feed aggregator!

## How to use

### Requirements
You will need the Go runtime and PostgreSQL installed on your machine

### Setup
place a `.gatorconfig.json` file in your home directory containing `{"db_url":"<CONNECTION_STRING>"`, where CONNECTION_STRING is a valid connection string to a postgresql DB with the ?sslmode=disable flag set.
install the `gator` CLI by navigating to the project's root directory and running `go install .`.

## Commands
### gator register \<username>

this command registers a user in the cli. a user must have a unique username.

### gator login \<username>

this command switches to the specified user. it will not work if the user does not exist.

### gator users

this command displays the users currently registered

### gator addfeed \<feedname> \<feedurl>

this command adds a feed to the database. the name can be freely chosen in order to identify the feed to yourself. only one entry can be made for each feed URL.
the added feed is automatically followed.

### gator feeds

this command displays all the feeds in the database.

### gator follow \<feedurl>

this command follows a feed by url. the feed needs to have been added to the database beforehand. following a feed includes its posts in the posts you browse and will cause the cli to aggregate its posts when you run `gator agg`.

### gator unfollow \<feedurl>

this command unfollows a feed again.

### gator following

this command displays which feeds you are following.

### gator agg \<interval>

this command aggregates posts from all feeds the current user is following. you must provide a valid duration string as `interval`, specifying the interval at which to fetch new posts.

### gator browse [\<limit>]

this command displays the newest posts across all your feeds, up to `limit` if `limit` is not specified, it is set to 2.