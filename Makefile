servername := iserver
hostname := localhost:1323
sid := else1561172677501-M_crt7Bzr1mMZRF3RWfzrVrlZBce0gnV 

serve:
	go run $(servername)
testapi:
	curl $(hostname)/users/1
	curl -b "else.sid=$(sid)" $(hostname)/users/1
