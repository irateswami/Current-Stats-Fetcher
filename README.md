# Current-Stats-Fetcher

## What's the point
The is a piece of software that will fetch the daily statistics of all MLB players for a fun data-anayltics project 

## How's it done
The dependency is a license to mysportsfeeds v2.1, this product hits this api daily for player statistics. It then pumps them into a database which we will query later. 

## How do I run it
With your license placed in "secrets.json" using the key "apiKey" use go run, everything else is handled for you
