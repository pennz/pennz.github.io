Some advices and useful commands for git.
This is a compilation of these three web-pages. For intermediate git users.

<!--# gist for these changes/github (just put to my dotfiles) -->
{% github_embed "https://github.com/v5analytics/gitbook-plugin-github-embed/blob/1cd16ac/index.js#L3-L8" %}{% endgithub_embed %}

# [Great advices](https://www.infoworld.com/article/3205884/27-essential-tips-for-git-and-github-users.html)
- Coding/Actions/CI
	- Clone almost anything
	- Pull frequently
	- Commit early and often
	- Comment your commits as you would have others comment theirs
	- Push when your changes are **tested**
	- Branch freely
	- Merge carefully
	- Stash before switching branches
	- Use editors and IDEs that “git it” (so you can do above things easily)
	- Automate your workflow with GitHub Actions
	- Build and publish packages
	- Check and resolve your security advisories and alerts
	- Scan your code for vulnerabilities
	- Publish your documentation pages
- Sharing
	- Fork a repo
	- Use gists to share snippets and pastes
	- Explore GitHub
	- Contribute to open source projects
	- Watch projects
	- Follow friends
	- Send pull requests
	- Create and resolve issues
	- Write informative README pages
	- Use Markdown
	- Convert your older repos to Git
	- Use GitHub project boards
	- Collaborate on documentation in wikis
- My suggestions
	- tag your commits (adding more information to your code)
	- Include batteries (tests and examples)

# [gitlab 15 Git tips to improve your workflow](https://about.gitlab.com/blog/2020/04/07/15-git-tips-improve-workflow/)
- git aliases
- See the repository status in your terminal’s prompt (edit PS1)
- Compare commits from the command line
	```shell
	git diff $start_commit..$end_commit -- path/to/file  # compare changes between the commits, you can also specify HEAD@{yesterday}
	```
- Stashing uncommitted changes
- Pull frequently
- Autocomplete commands (Tab), so your loop can be faster
- Set a global .gitignore
	```shell
	touch ~/.gitignore
	git config --global core.excludesFile ~/.gitignore
	```
- Enable Git’s autosquash feature by default
	```shell
	git rebase -i --autosquash
	git config --global rebase.autosquash true
	```
	`git rebase --interactive --autosquash` only picks up on commits with a
message that begins **fixup!** or **squash!**, and Git still gives you the
chance to to move things around in your editor like a regular interactive
rebase.
	[**You can check more details in this thoughtbot blog post**](https://thoughtbot.com/blog/autosquashing-git-commits)
- git blame
	```shell
	git blame -w  # ignores white space
	git blame -M  # ignores moving text
	git blame -C  # ignores moving text into other files
	```
- Add an alias to check out merge requests locally
In your .gitconfig file
	```
	[alias]
	  mr = !sh -c 'git fetch $1 merge-requests/$2/head:mr-$1-$2 && git checkout mr-$1-$2' -
	```
	```shell
	git mr upstream 5
	```

- An alias of HEAD
	Breaking news: @ is the same as HEAD. Using it during a rebase is a lifesaver:
	```shell
	git rebase -i @~2 # rebase from the second pevious to HEAD
	```
- `git reset` / `git checkout` to undo changes


# [opensource.com 13 Git tips for Git's 13th birthday](https://opensource.com/article/18/4/git-tips)

- .gitconfig file
	- your ~/.gitconfig file (you need set git config your name and email for first use of git. it is saved to this file)
	- your repo's .gitconfig
- aliases (faster command)
	- `git config --global --add alias.st status`, so `git st` will do the same as `git status`
	- previous command will save to the ~/.gitconfig file
	- you can also aliases to shell commands, e.g., put the following command to the alias section in your .gitconfig.
	```
	upstream-merge = !"git fetch origin -v && git fetch upstream -v && git merge upstream/master && git push"
	```
- visializing the commit graph
	using this git alias
	```
	[alias]
	    logp = log --pretty=oneline --graph --decorate=full
	    logt = log --pretty=oneline --graph --topo-order
	    logd = log --pretty=oneline --graph --date-order
	    loga = log --oneline --decorate=full
	    lg1 = log --graph --abbrev-commit --decorate --format=format:'%C(bold blue)%h%C(reset) - %C(bold green)(%ar)%C(reset) %C(white)%s%C(reset) %C(dim white)- %an%C(reset)%C(bold yellow)%d%C(reset)' --all
	    lg2 = log --graph --abbrev-commit --decorate --format=format:'%C(bold blue)%h%C(reset) - %C(bold cyan)%aD%C(reset) %C(bold green)(%ar)%C(reset)%C(bold yellow)%d%C(reset)%n''          %C(white)%s%C(reset) %C(dim white)- %an%C(reset)' --all
	    lg = !"git lg1"
	```

- a nicer force-push
	Use `git push --force-with-lease`, it will not allow you to force-push if the remote branch has been updated. So you won't throw away someone else's work
- more fine-grained add/change
	```shell
	git add -N # then you can use git diff show differences before you use git add -a, you can check it
	git add -p # more fine grained
	git checkout -p
	git rebase -x/ git rebase --exec YOUR_COMMAND_TO_TEST # you can run test suit after each rebase commit
	```
- time-based revision references
	```shell
	git diff HEAD@{yesterday} # compare current HEAD with the HEAD of yesterday
	git diff HEAD@{'2 months ago'}
	git diff HEAD@{'2010-01-01 12:00:00'}
	```

- `git rebase` and `git reflog`
	if you ever found you rebased away a committed changes, you can use `git reflog` to find it.
- keep it clean
	`git branch --merged` to get the list of merged branches. Then you can clean it up

# Version history
- v1 20200720 draft
