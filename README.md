## Historian: Generate "Release Notes" in 5 seconds.

### Installation

```
go get github.com/javidgon/historian
```
In case of errors, please have a look at the `git2go` documentation [here](https://github.com/libgit2/git2go).
You might need to follow the instuctions as described in the `install_requirements.sh` file

### Getting started

```
# The first argument is the path to the git repository
# The second argument is the last commit we want to include in the "release notes"

historian /path/to/the/repo 4481c5c8b3b548490e0cc55a18fb599a61d98131

Processing...
err %v <nil>
err %v <nil>
err %v <nil>
err %v <nil>
err %v <nil>
err %v <nil>
err %v <nil>
err %v <nil>
err %v <nil>
err %v <nil>
err %v <nil>
ret %v 0
********************************************
release_notes.txt file successfully created!
********************************************
```

Having a look at the `release_notes.txt` file:

```
==========================================================
** Release Notes (Sunday, 26-Jul-15 09:21:29 UTC):
0) <you@example.com> (2015-07-25 19:10:03 +0000 +0000) This is my first commit
1) <you@example.com> (2015-07-25 19:20:16 +0000 +0000) This my second commit
2) <you@example.com> (2015-07-26 08:14:10 +0000 +0000) This is my third commit

This commit has body with points as well:

* Point 1
* Point 2

3) <you@example.com> (2015-07-26 08:20:14 +0000 +0000) This is my forth commit
4) <you@example.com> (2015-07-26 08:24:10 +0000 +0000) This is my fifth comment

This one has a smaller body
==========================================================
```
