# RSS Aggregator Project using Golang

This is a simple RSS aggregator project implemented using the Go programming language. The application fetches and aggregates RSS feeds from various sources, providing users with a consolidated view of the latest news and updates.

## Features

- Fetches RSS feeds from multiple sources.
- Aggregates the latest news items from different feeds.
- Provides a clean and organized view of the aggregated news.

## Installation

1. Make sure you have Go installed on your system. If not, download and install it from the official Go website: https://golang.org/

2. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/rss-aggregator.git
   cd rss-aggregator
   ```

3. Build the project:

   ```bash
   go build
   ```

4. Run the executable:

   ```bash
   ./rss-aggregator
   ```

## Configuration

The project supports adding new RSS feed sources and customizing the behavior. To configure the aggregator, modify the `config.json` file:

```json
{
  "feeds": [
    {
      "name": "TechCrunch",
      "url": "https://techcrunch.com/rss"
    },
    {
      "name": "The Verge",
      "url": "https://www.theverge.com/rss/index.xml"
    },
    {
      "name": "BBC News",
      "url": "http://feeds.bbci.co.uk/news/rss.xml"
    }
  ],
  "refresh_interval": "30m"
}
```

- `"feeds"`: Add your desired RSS feed sources with a name and URL. You can include as many feeds as you like.

- `"refresh_interval"`: Set the time interval for refreshing the feeds. The value should be in the format of Go's `time.Duration`. For example, `"30m"` stands for 30 minutes.

## Usage

1. Run the application as described in the Installation section.

2. Open your web browser and go to `http://localhost:8080` (or the configured address if you changed it).

3. The application will fetch and display the latest news items from the configured RSS feeds.

4. Enjoy staying up-to-date with the latest news!

