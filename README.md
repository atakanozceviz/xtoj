# xtoj
xtoj is a bi-directional XML, JSON converter api
If the given data is XML, it will be converted to JSON,
Or if the data is JSON, it will be converted to XML.
### How to use?
To read the body of given url and then convert to json or xml, add url as query string parameter like so:
example: https://xtoj.herokuapp.com/?url=https://www.ExampleRssFeed.com/rss

Or just add your XML or JSON as the body and make a POST request to https://xtoj.herokuapp.com/

### Deploy to heroku
Everything is pre-configured, Go Modules is used for vendoring.