# arc
arc diff is to push the local branch to the remote diff branch forcely, you can use 
merge and rebase to update the branch, it doesn't matter, it always create the diff between the latest and the crossing point.
* if commit in the same branch by merging or something else except rebasing, it will create the diff between the latest when you execute arc diff and the crossing point you create branch  
so from the diff website, you can see that each commit, the base commit doesn't change.
* if use rebase, the crossing point changes, so from the diff website, you can see that the base is changed to the commit you start to rebase.  

arc diff is always to create the diff between the crossing point you create branch and the latest you execute arc diff.