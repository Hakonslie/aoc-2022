## Day 

### Solution

At first I over-engineered this one so much. I was trying to build a filesystem
and put everything in there. I spent maybe 2 hours this morning trying that.
But then I realized if I can somehow keep track of my folders while traversing the lines 
then I don't need it.


My original plan was to create a filesystem, add all the files there
and then recursively determine the size of the folders.

Wow, I liked the task though.

Basically my solution is this:

Keep track of which folders I am "inside". Do to this I need to give every folder a unique name since the names are not unique without their parent.
For this I add a number to their name. Whenever I traverse files I add the size to all of the folders I am inside. When I go one step up I close the most recently opened folder. When I go to the top I close all folders(Looks to me like this was uneccessary because it is only used once, kind of sad about that.

In the end I just do some simple calculations on my list of folders with sizes.
