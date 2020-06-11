If you want just svelte then do the following
Git Clone https://github.com/acim/erinus
Run the command 
Cd into directory
npm install
npm run build
npm run dev
The package.json runs "start": "go run main.go"
You may have to change this to "start": "go run *.go"
And thats it. You now have a running svelte project with go.
If you want sapper then do the following.
Git clone https://github.com/golangast/go_sapper/tree/92e3b4ff57d773ccea6afc7c1b0ff9231edb3852
If you want to do this by scratch then install sapper and add the main.go and go.mod files.
Then change the package.json to include the following.
{ "scripts": { "build": "sapper export && mv ./__sapper__/export ./public" } }
"start": "go run main.go"
And go will pull the files that are built after you do npm run build and will run when you do npm start.
Technically you will have two servers running but you will have the power of go behind the nodejs server.
This way you get all the nice libs of go backend and the sapper qualities that are nice on the frontend like restart and rebuild and passing imports.
