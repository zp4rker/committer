# committer
An auto-committer written in Go.

## Milestones
Some great feats reached with this tool.
- 1 million @ https://github.com/zp4rker/1mil-commits
- 5 million @ https://github.com/zp4rker/5mil-commits
- 10 million @ https://github.com/zp4rker/10mil-commits
- 25 million @ https://github.com/zp4rker/25mil-commits
- 50 million @ https://github.com/zp4rker/50mil-commits
- ~~100 million @ https://github.com/zp4rker/100mil-commits~~ (requested to stop and remove it by Github at the 61.5 million mark)

## Usage
### Setup
```
git clone https://github.com/zp4rker/committer
cd committer
rm -rf .git/ README.md LICENSE
```
**Notes** 
- You'll have to set up git on your machine and ensure it is able to successfully commit. Usually, you'll have to config the git `user.name` and `user.email`
- If you want any other files to be in the repository, add them now; the app will automatically add and commit them

### Start
```
go run . --amount X --final-commit "this is the final message"
```

## Features
- Custom amount of commits to go up to
- Custom final commit message
- Live progress output
- Ability to cancel/abort and resume where left off
