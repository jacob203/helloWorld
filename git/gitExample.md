# git hub
if you want to contribute to an open source project, usually you need to fork the repository, it means to copy the 
repository in your account, actually it is another kind of branch. when your fix is done, check in your repository firstly,
then create a pull request to the owner of the project, it needs review, when the review is done, it is merged into the original.

what if there are some conflicts after creating a pull request?  
at this time, you need to sync the original to the forked, normally 
```
git pull https://github.com/ORIGINAL_OWNER/ORIGINAL_REPOSITORY.git BRANCH_NAM
```
or 
add a remote upstream for the original.
```
git remote add upstream git@github.com:myteksi/people.git
git checkout master
git rebase upstream/master
git push --force
```
"git remote add" is used to add another repository. 