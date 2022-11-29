# Capstone Project
## Summary
For my capstone project, I have decided to build a discord bot that will allow League of Legends players to keep track of their in-game statistics as well as communicate with Riot Games API to look up their match history.
I will also be working to develop open-API documentation for the api layer of my bot so that my friend who is a front-end developer will have the documentation that he needs to use the API I've created to make a front-facing website.

This project will utilize the following things that we've learned in class:
- Building a REST API
- Integration with existing services to retrieve game data
- Using GORM and PostgreSQL to store player statistics and profiles alongside Discord and in-game identifiers needed to make API calls
- Requesting, storing, and using API Keys and session tokens.

Alongside the things that I've learned in class, I will also demonstrate:
- Understanding of how to use [third party wrappers/packages](github.com/bwmarrin/discordgo) for external services
- Utilization of Railway.app to host and deploy my services
- Hosting and deploying kubernetes instances locally on my home server for development
- Understanding of how to utilize Redis as a caching layer
- Running timed go routines that will allow for leaderboards and statistics to be automatically updated at specific intervals.
- Using Discord's 0auth2 system to allow players who log in on the front facing website to view their specific statistics and profile information.

## User Stories
### Number One
### As a gamer, I find it difficult to manage multiple websites to track my statistics and also use Discord to communicate with my friends.
**Acceptance Criteria**  
 Create a way that allows for Discord and in-game statistics to be integrated into a single ecosystem.

Example:
```
Using discord "slash commands":
/player lookup TheJoshJ
```
This would return a list of statistics for "TheJoshJ" directly from Riot's API including recently played matches, and most played champions.

### Number Two
### As someone who doesn't enjoy playing League of Legends at a high & competitive, I find it difficult to find new ways to have fun with the game while also embracing friendly competition.
**Acceptance Critera**  
Have a weekly rotating objective and an hourly updating leaderboard to encourage friendly competition in a way that isn't directly associated with just winning the game.

Example:
```
Weekly Objective: Most Gold Earned
1. Dextri - 50.5k
2. 10ChrisRulz - 48.4k
3. NetBus - 48.3k
...
```
### Number Three
### As a developer, I would like to tap into the leaderboard system and allow my friends to compete alongside me
**Acceptance Criteria**  
Create documentation for my API layer using open-API so that other developers have the tools that they need to replicate similar systems without having to build out the backend themselves.

Example:
[Existing Documentation](https://api.lolqueue.com/docs/index.html#/accounts)
