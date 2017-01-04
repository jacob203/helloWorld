git is a distributed open source version control software, there are a lot of
strengths:
* **no center**: no need to have a control center, each developer's repo is a
 center, you don't have to connect a center, you are allowed to connect anyone's
 repo. so even when you are offline, you also can work.
* open source and free

although it is famous for its no center feature, there is always a code center, all codes should be put there.

## important concepts
#### three checkin states
there are three states in git, working directory->stage->repo, it is also the order we control.  
* working directory: the directory you are working on
* stage: use "git add" to add changes to the stage, which are waiting to be confirmed to put in repo
* repo: the state in which all files are controlled by the version control system

#### git status
this command should be used mostly, it is used to do two things:
1. check file changes in the working directory
2. check file changes in stage

#### git-add put files from working directory to stage
when we need to put files to stage, need to use git add to add those files.  
if you changes a file , later you put it to stage, later you changes 
git add is only used to add changes.
there are some options you need to know
* git add filePath
git add hello.md git/learninggit.md
* git add .
git add all changes in the current directory including sub directories
* git add -u
update the stage files changes, for example, you put hello.txt in the stage, then you modify hello.txt again,
then you can use only git add -u to update changes, don't need to use git add file path
git add file

#### HEAD
it refers to the last commit in the currently checked-out branch

#### Working tree
it is working directory, it is the files that you are currently working on.

#### Git index
it is also known as stage, it is where you put files you want to commit to git repository

## git-reset used to drop a commit
it is divided into three parts:
1. drop commits from stage
git add paths is used to add all changes in that path to stage  
git reset paths is used to remove all changes in that path from stage
2. drop patchs
git add -p is used to apply a patch to the working directory
git reset -p is used to remove the patch from the working directory
3. drop commits
acctually it doesn't drop commits, it only resets the current branch head to a specific commit
git reset --soft HEAD^
then it puts all commits between the wanted commit and the current commit into stage

git reset --mix HEAD^^
reseet to the HEAD^^ commit, undo all files in stage. so when just wantting to drop all changes in stage  
you can just use git reset --mix

git reset --hard
drop all changes in stage and working directory.


# rebase
rebase is used to merge commits by applying the branch commits to the other branch, generally rebase feature branch to master.  
for example, you have a feature branch, it is finished, you want to merge your feature branch to master, and there are other developers working on the master, there migth be some conficts, so you want to fix the conficts.  
rebasing feature to master is to apply all feature commits to the latest master in your repo.  


how rebase works?
if you rebase feature to master, firstly it find the crossing point which you create feature branch at, then create a shaddow branch feature at the current master you have, apply all old feature branch's commits to the shaddow one, there might be some conficts, fix them. when rebase is done, your old feature branch will be removed, the shaddow one will be the feature branch, all the commits on the shaddow branch are new. at this time, generally you have to do force push, use your local branch to replace your remote, as the commit history is different.  
so be careful that rebase will change log history.  
what if you don't rebase on the same branch?  
for example "git rebase -i HEAD~2", there are no crossing points, then it use HEAD~2 as the crossing point, and use HEAD~2 as the base, apply HEAD~1 and HEAD to the base, if you don't do any squash, fixup or drop and so on, it does nothing. so rebase it on the same branch is used to clean up your feature branch.

_**one golden rule**_ is that it is not allowed to do rebase to a public branch which is not only owned by you but also can be committed by others. if you do that, you will re-write the history of the branch, it will affect others and make them confused. only allowed to do rebase on your own branch.

another senario to use rebase is that you and your coworker are working on the same branch, you do two commits and your coworker does another two commits, the next thing is to merge your coworker's commits to the local, let's assume your coworker syncs with remote repo first.  
at this time, if you do merge, the remote branch's commits will be merged with yours as one commit, then you will have to do force-push your local branch, then your coworker see what you did, but he will lose his commit history. it is good to do that.
the better way is to do rebase, add your changes to the remote, and push your local branch, then your coworker can see what you did, and he will not lose his commit history.
it doesn't break the golden rule.

when you use pull request work flow to review codes, it is not good to do rebase when a pull request is created.
clean up the local branch with rebase before submitting a pull request, and use merge to add other developers' commits, don't use rebase, it will changes history, it is hard for reviewers to see what changes.

both rebase and merge can be used to do merge branches, if you prefer a clean linear history, you should use rebase, if you want a complete history and avoid the risk of re-writing public branches' history, use merge.


the senario is reproduced:
1. create a branch such as feature
2. do three commits in a file such as world.txt , just add three lines, on the feature branch
3. checkout master branch, then add another two commits on the same file world.txt
4. checkout feature branch, then rebase -i master, you will see you have to solve three conflicts  
you have three commits in feature branch,you have three conflicts.it is wired?



#### local and remote
as said, there is always a version control center, such as github, so we need to sync 
## important commands
#### git add
the first commands we need to learn is "git add", it is used to add changes to a stage.