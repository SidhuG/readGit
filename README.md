# readGit
Checkouts Git repo at specific tag,parses yaml configuration files in the git repo, and creates key/value endpoints in consul or etcd command options can be either specifid in command line or in config file in yaml forma

Please see an example

**For cloning private repos, it works by adding your key to ssh-agent:**
```
eval "$(ssh-agent -s)"
ssh-add -K ~/.ssh/<your key>
```