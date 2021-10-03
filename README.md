# committer
An auto-committer written in Go.

## Usage
### Setup
```
git clone https://github.com/zp4rker/committer
cd committer
rm -rf .git/ README.md LICENSE
```
**Note:** You'll have to set up git on your machine and ensure it is able to successfully commit. Usually, you'll have to config the git `user.name` and `user.email`.

### Start
```
go run . --amount X --final-commit "this is the final message"
```