# commit-message-fixer
A small command line tool written in Go to fix the commit message by adding JIRA issue id at the start

# Get started #
For `Windows`, put `cmf.exe` inside the root directory of the repository, then put `commit-msg` in `.git/hooks` directory, and commit as you wish and JIRA issue no. will be added at the start if you forget to add it yourself. For `Mac OS X`, you need to put the file `cmf` inside the root directory of the repository instead of `cmf.exe`. For `Linux`, build the executable yourself and put it inside the root directory of the repository, the rest are the same.  