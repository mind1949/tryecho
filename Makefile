servername := iserver

serve:
	go run $(servername)
testapi:
	curl localhost:1323/users/1

