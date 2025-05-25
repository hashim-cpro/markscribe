# markscribe

[![Latest Release](https://img.shields.io/github/release/muesli/markscribe.svg)](https://github.com/muesli/markscribe/releases)
[![Build Status](https://github.com/muesli/markscribe/workflows/build/badge.svg)](https://github.com/muesli/markscribe/actions)
[![Go ReportCard](https://goreportcard.com/badge/muesli/markscribe)](https://goreportcard.com/report/muesli/markscribe)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/muesli/markscribe)

Your personal markdown scribe with template-engine and Git(Hub) & RSS powers 📜

## Table of Contents

- [markscribe](#markscribe)
  - [Table of Contents](#table-of-contents)
  - [Usage](#usage)
  - [Installation](#installation)
    - [Packages \& Binaries](#packages--binaries)
    - [Build From Source](#build-from-source)
  - [Templates](#templates)
  - [Functions](#functions)
    - [RSS feed](#rss-feed)
    - [Your recent contributions](#your-recent-contributions)
    - [Your recent pull requests](#your-recent-pull-requests)
    - [Repositories you recently starred](#repositories-you-recently-starred)
    - [Repositories you recently created](#repositories-you-recently-created)
    - [Custom GitHub repository](#custom-github-repository)
    - [Forks you recently created](#forks-you-recently-created)
    - [Recent releases you contributed to](#recent-releases-you-contributed-to)
    - [Your published gists](#your-published-gists)
    - [Your latest followers](#your-latest-followers)
    - [Your sponsors](#your-sponsors)
    - [Your GoodReads reviews](#your-goodreads-reviews)
    - [Your GoodReads currently reading books](#your-goodreads-currently-reading-books)
    - [Your Literal.club currently reading books](#your-literalclub-currently-reading-books)
    - [Hackatime Stats and Visualizations](#hackatime-stats-and-visualizations)
      - [Basic Stats](#basic-stats)
      - [Language Statistics with Progress Bars](#language-statistics-with-progress-bars)
      - [Individual Language Data](#individual-language-data)
      - [Available Data](#available-data)
  - [Template Engine](#template-engine)
  - [Template Helpers](#template-helpers)
  - [GitHub Authentication](#github-authentication)
  - [GoodReads API key](#goodreads-api-key)
  - [Hackatime API Setup](#hackatime-api-setup)
  - [FAQ](#faq)

You can run markscribe as a GitHub Action: [readme-scribe](https://github.com/muesli/readme-scribe/)

## Usage

Render a template to stdout:

    markscribe template.tpl

Render to a file:

    markscribe -write /tmp/output.md template.tpl

## Installation

### Packages & Binaries

If you use Brew, you can simply install the package:

    brew install muesli/tap/markscribe

Or download a binary from the [releases](https://github.com/muesli/markscribe/releases)
page. Linux (including ARM) binaries are available, as well as Debian and RPM
packages.

### Build From Source

Alternatively you can also build `markscribe` from source. Make sure you have a
working Go environment (Go 1.16 or higher is required). See the
[install instructions](https://golang.org/doc/install.html).

To install markscribe, simply run:

    go get github.com/muesli/markscribe

## Templates

You can find an example template to generate a GitHub profile README under
[`templates/github-profile.tpl`](templates/github-profile.tpl). Make sure to fill in (or remove) placeholders,
like the RSS-feed or social media URLs.

Rendered it looks a little like my own profile page: https://github.com/muesli

## Functions

### RSS feed

```
{{range rss "https://domain.tld/feed.xml" 5}}
Title: {{.Title}}
URL: {{.URL}}
Published: {{humanize .PublishedAt}}
{{end}}
```

### Your recent contributions

```
{{range recentContributions 10}}
Name: {{.Repo.Name}}
Description: {{.Repo.Description}}
URL: {{.Repo.URL}})
Occurred: {{humanize .OccurredAt}}
{{end}}
```

This function requires GitHub authentication with the following API scopes:
`repo:status`, `public_repo`, `read:user`.

### Your recent pull requests

```
{{range recentPullRequests 10}}
Title: {{.Title}}
URL: {{.URL}}
State: {{.State}}
CreatedAt: {{humanize .CreatedAt}}
Repository name: {{.Repo.Name}}
Repository description: {{.Repo.Description}}
Repository URL: {{.Repo.URL}}
{{end}}
```

This function requires GitHub authentication with the following API scopes:
`repo:status`, `public_repo`, `read:user`.

### Repositories you recently starred

```
{{range recentStars 10}}
Name: {{.Repo.Name}}
Description: {{.Repo.Description}}
URL: {{.Repo.URL}})
Stars: {{.Repo.Stargazers}}
{{end}}
```

This function requires GitHub authentication with the following API scopes:
`repo:status`, `public_repo`, `read:user`.

### Repositories you recently created

```
{{range recentRepos 10}}
Name: {{.Name}}
Description: {{.Description}}
URL: {{.URL}})
Stars: {{.Stargazers}}
{{end}}
```

This function requires GitHub authentication with the following API scopes:
`repo:status`, `public_repo`, `read:user`.

### Custom GitHub repository

```
{{with repo "muesli" "markscribe"}}
Name: {{.Name}}
Description: {{.Description}}
URL: {{.URL}}
Stars: {{.Stargazers}}
Is Private: {{.IsPrivate}}
Last Git Tag: {{.LastRelease.TagName}}
Last Release: {{humanize .LastRelease.PublishedAt}}
{{end}}
```

This function requires GitHub authentication with the following API scopes:
`repo:status`, `public_repo`, `read:user`.

### Forks you recently created

```
{{range recentForks 10}}
Name: {{.Name}}
Description: {{.Description}}
URL: {{.URL}})
Stars: {{.Stargazers}}
{{end}}
```

This function requires GitHub authentication with the following API scopes:
`repo:status`, `public_repo`, `read:user`.

### Recent releases you contributed to

```
{{range recentReleases 10}}
Name: {{.Name}}
Git Tag: {{.LastRelease.TagName}}
URL: {{.LastRelease.URL}}
Published: {{humanize .LastRelease.PublishedAt}}
{{end}}
```

This function requires GitHub authentication with the following API scopes:
`repo:status`, `public_repo`, `read:user`.

### Your published gists

```
{{range gists 10}}
Name: {{.Name}}
Description: {{.Description}}
URL: {{.URL}}
Created: {{humanize .CreatedAt}}
{{end}}
```

This function requires GitHub authentication with the following API scopes:
`repo:status`, `public_repo`, `read:user`.

### Your latest followers

```
{{range followers 5}}
Username: {{.Login}}
Name: {{.Name}}
Avatar: {{.AvatarURL}}
URL: {{.URL}}
{{end}}
```

This function requires GitHub authentication with the following API scopes:
`read:user`.

### Your sponsors

```
{{range sponsors 5}}
Username: {{.User.Login}}
Name: {{.User.Name}}
Avatar: {{.User.AvatarURL}}
URL: {{.User.URL}}
Created: {{humanize .CreatedAt}}
{{end}}
```

This function requires GitHub authentication with the following API scopes:
`repo:status`, `public_repo`, `read:user`, `read:org`.

### Your GoodReads reviews

```
{{range goodReadsReviews 5}}
- {{.Book.Title}} - {{.Book.Link}} - {{.Rating}} - {{humanize .DateUpdated}}
{{- end}}
```

This function requires GoodReads API key!

### Your GoodReads currently reading books

```
{{range goodReadsCurrentlyReading 5}}
- {{.Book.Title}} - {{.Book.Link}} - {{humanize .DateUpdated}}
{{- end}}
```

This function requires GoodReads API key!

### Your Literal.club currently reading books

```
{{range literalClubCurrentlyReading 5}}
- {{.Title}} - {{.Subtitle}} - {{.Description}} - https://literal.club/_YOUR_USERNAME_/book/{{.Slug}}
  {{- range .Authors }}{{ .Name }}{{ end }}
{{- end}}
```

This function requires a `LITERAL_EMAIL` and `LITERAL_PASSWORD`.

### Hackatime Stats and Visualizations

To use Hackatime integration, you need to set up the following environment variables:

- `WAKATIME_API_KEY`: Your Hackatime API key
- `WAKATIME_USERNAME`: Your Hackatime username

#### Basic Stats

Get your basic coding statistics:

```
{{with hackatimeStats}}
Status: {{.Data.Status}}
Total Time: {{.Data.HumanReadableTotal}}
Daily Average: {{.Data.HumanReadableDailyAvg}}
{{end}}
```

#### Language Statistics with Progress Bars

Display your most used languages with beautiful progress bars:

```
{{with hackatimeStats}}
{{ wakatimeLanguages "💾 Languages:" .Data.Languages 5 .Data.HumanReadableTotal }}
{{end}}
```

This will output something like:

```
💾 Languages:
TypeScript      5h 52m 21s   ████████░░░░░░░░░░░░░░░░░  30.19%
HTML            2h 15m 23s   ███░░░░░░░░░░░░░░░░░░░░░░  11.60%
JavaScript      2h 14m 48s   ███░░░░░░░░░░░░░░░░░░░░░░  11.55%
Ruby            1h 51m 43s   ███░░░░░░░░░░░░░░░░░░░░░░  9.57%
Nix             1h 24m 0s    ██░░░░░░░░░░░░░░░░░░░░░░░  7.20%

Total: 19 hrs 28 mins
```

#### Individual Language Data

Access detailed information about each language:

```
{{with hackatimeStats}}
{{range .Data.Languages}}
- {{.Name}}: {{.Text}} ({{.Percent}}%)
{{end}}
{{end}}
```

Each language entry includes:

- Name of the programming language
- Total coding time in human-readable format
- Hours and minutes breakdown
- Percentage of total coding time

#### Available Data

The Hackatime integration provides access to:

- Coding activity visibility status
- Time range information
- Total coding time
- Daily averages
- Language statistics

Each language entry includes:

- Name
- Total seconds
- Human-readable time
- Percentage
- Digital format time

> Currently [Hackatime API](https://github.com/hackclub/hackatime) doesn't has any direct way of extracting projects related data. So can't access that. Also, I tried to filter the data for specific intervals on the API but that ain't working for some reason as well.(API issue)

## Template Engine

markscribe uses Go's powerful template engine. You can find its documentation
here: https://golang.org/pkg/text/template/

## Template Helpers

markscribe comes with a few handy template helpers:

To format timestamps, call `humanize`:

```
{{humanize .Timestamp}}
```

To reverse the order of a slice, call `reverse`:

```
{{reverse (rss "https://domain.tld/feed.xml" 5)}}
```

## GitHub Authentication

In order to access some of GitHub's API, markscribe requires you to provide a
valid GitHub token in an environment variable called `GITHUB_TOKEN`. You can
create a new token by going to your profile settings:

`Developer settings` > `Personal access tokens` > `Generate new token`

## GoodReads API key

In order to access some of GoodReads' API, markscribe requires you to provide a
valid GoodReads key in an environment variable called `GOODREADS_TOKEN`. You can
create a new token by going [here](https://www.goodreads.com/api/keys).
Then you need to go to your repository and add it, `Settings -> Secrets -> New secret`.
You also need to set your GoodReads user ID in your secrets as `GOODREADS_USER_ID`.

## Hackatime API Setup

To use the Hackatime integration, you need to set up two environment variables:

1. `WAKATIME_API_KEY`: Your Hackatime API key

   - Log into your Hackatime account
   - Go to your settings page
   - Generate a new API key or copy your existing one
   - Set the environment variable: `WAKATIME_API_KEY="your-api-key"`

2. `WAKATIME_USERNAME`: Your Hackatime username
   - This is your slack id in [Hackclub Slack](https://hackclub.com/slack/)
   - Set the environment variable: `WAKATIME_USERNAME="your-username"`

## FAQ

Q: That's awesome, but can you expose more APIs and data?  
A: Of course, just open a new issue and let me know what you'd like to do with markscribe!

Q: That's awesome, but I don't have my own server to run this on. Can you help?  
A: Check out [readme-scribe](https://github.com/muesli/readme-scribe/), a GitHub Action that runs markscribe for you!
