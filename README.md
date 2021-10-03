# committer
An auto-committer written in Go.

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