# devops

autocrlf in git

## Simple Usage

[git help page](https://help.github.com/en/github/using-git/configuring-git-to-handle-line-endings) For windows:

```bash
$ git config --global core.autocrlf true
# Configure Git to ensure line endings in files you checkout are correct for Windows.
# For compatibility, line endings are converted to Unix style when you commit files.
```

For Mac/Linux:

```bash
$ git config --global core.autocrlf input
# Configure Git to ensure line endings in files you checkout are correct for OS X/Linux, when commit files, windows style end of line will be changed to Linux style (lf).
```

## Detail

TODO... [stackoverflow](https://stackoverflow.com/questions/1967370/git-replacing-lf-with-crlf/20653073#20653073) see this for windows setting.

