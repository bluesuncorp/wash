# wash

steps to get going
------------------

1. Clone repo eg. git clone https://github.com/bluesuncorp.wash $GOPATH/src/github.com/<user or org name>/<project name>
2. remove .git folder in clone repo, run git init and then add remote
3. do a global replace on the code eg. replace: github.com/bluesuncorp/wash with github.com/<user or org name>/<project name>
4. run ./bin/init.sh to download any dependencies
5. run ./bin/dev.sh to start developing your app


COMPILING TO STATIC BINARY
--------------------------
- NOTES coming soon, just creating the Dockerfile etc...

NOTES:
------
- a docker solution is in the works which will make everything easier
- ENV variables are locaed in ./bin/dev.sh for you manipulation