#to setup the project -

https://docs.google.com/document/d/12oCEuxB6bklnEQhpuIcgt2NekFi0xaHguvGO8Ridxf4/edit?usp=sharing

#Remember that the database is setup like the following.
user -> users table with the columns

id | name | email | pass

#Remember to also have the following installed.

- npm install --save-dev del-cli
- Correction you may have to change the config.json if you are using windows to have “mv” be “move” in the line 
Instead of this

"build": "sapper export && mv ./__sapper__/export ./public",

You would have this

"build": "sapper export && move ./__sapper__/export ./public",

#commands to build and run

this is to install everything
- npm install 

this is to compile the files and send them to the running folders.
- npm run build

this is to have auto restart as you dev
- npm run dev

this is to actually run the program/ still need to work in the go build and executable
- npm start


