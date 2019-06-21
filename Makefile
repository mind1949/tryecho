servername := iserver
hostname := localhost:1323
sid := mind1949

serve:
	go run $(servername)
testapi:
	curl $(hostname)/users/1
	curl -b "else.sid=$(sid)" $(hostname)/users/1
