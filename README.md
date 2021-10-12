# commit-message-fixer
A small command line tool written in Go to fix the commit message by adding `JIRA` issue id at the beginning

# Get started #

## Windows
For `Windows`, put `cmf.exe` inside the root directory of the repository, then put the `./commit-msg` script in `.git/hooks` directory, and commit as you wish and JIRA issue no. will be added at the beginning if you forget to add it yourself. 

## Mac OSX
For `Mac OS X`, you need to put the file `cmf` inside the root directory of the repository instead of `cmf.exe`. Then run below command in the root directory of the repository to make it an executable
```
chmod 755 ./cmf
```
For recent version of `Mac`, to walkaround the security check of the OS, we need to run it from command line like below
```
./cmf 
```
then allow the application to run by clicking `Open` button in the popup warning dialog; otherwise, it will crash at runtime when we try to execute from the script

then put the `./commit-msg` script in `.git/hooks` directory, run below command to make it executable
```
chmod 755 ./commit-msg
```
The rest are the same as `Windows`.

## Linux
For `Linux`, about the same as `Mac OSX`, but you need to build the executable yourself
