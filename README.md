# dirlist
recursive list of directory with size and child count

uses depth-first search 

_could be done with a shell-script too, but this appeared 
to be faster and more platform independent..._

## Syntax

`dirlist -dir <folder> -csv <output file>`

```
PS C:\daten\go\src> ..\bin\dirlist.exe -dir .\github.com\je4\dirlist\ -csv C:\temp\output.csv
github.com/je4/dirlist/.git/branches: Size:0 / Folders:0 / Files:0
github.com/je4/dirlist/.git/hooks: Size:18848 / Folders:0 / Files:11
github.com/je4/dirlist/.git/info: Size:240 / Folders:0 / Files:1
github.com/je4/dirlist/.git/logs/refs/heads: Size:336 / Folders:0 / Files:1
github.com/je4/dirlist/.git/logs/refs/remotes/origin: Size:328 / Folders:0 / Files:2
github.com/je4/dirlist/.git/logs/refs/remotes: Size:328 / Folders:1 / Files:2
github.com/je4/dirlist/.git/logs/refs: Size:664 / Folders:3 / Files:3
github.com/je4/dirlist/.git/logs: Size:1000 / Folders:4 / Files:4
github.com/je4/dirlist/.git/objects/05: Size:188 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/0a: Size:180 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/17: Size:166 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/32: Size:712 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/3e: Size:49 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/74: Size:208 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/78: Size:176 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/81: Size:55 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/86: Size:183 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/96: Size:157 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/bf: Size:209 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/ef: Size:151 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/fa: Size:169 / Folders:0 / Files:1
github.com/je4/dirlist/.git/objects/info: Size:0 / Folders:0 / Files:0
github.com/je4/dirlist/.git/objects/pack: Size:14212 / Folders:0 / Files:2
github.com/je4/dirlist/.git/objects: Size:16815 / Folders:15 / Files:15
github.com/je4/dirlist/.git/refs/heads: Size:41 / Folders:0 / Files:1
github.com/je4/dirlist/.git/refs/remotes/origin: Size:73 / Folders:0 / Files:2
github.com/je4/dirlist/.git/refs/remotes: Size:73 / Folders:1 / Files:2
github.com/je4/dirlist/.git/refs/tags: Size:0 / Folders:0 / Files:0
github.com/je4/dirlist/.git/refs: Size:114 / Folders:4 / Files:3
github.com/je4/dirlist/.git: Size:38579 / Folders:29 / Files:41
github.com/je4/dirlist/.idea: Size:4645 / Folders:0 / Files:6
github.com/je4/dirlist: Size:80515 / Folders:31 / Files:52
PS C:\daten\go\src>
