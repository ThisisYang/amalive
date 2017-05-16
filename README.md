# amalive

This is a golang practice.
Trying to build a client/server module that similar to server heartbeat.

Main purpose of this is to check accessibility of target server or virtual machine.

The accessibility means if the server on internet can reach the client that installed in target server or virtual machine.

Ths reason of creating this project is because of facing some connectivity issue on a VPS provider that the virtual machine can response icmp packets, but dropped all tcp/udp packets sometimes.

If happened again, wants to use this app to send notification email.

## push module

If push module, client send alive signal to server. Server provide keep track the all the available servers.

If server didn't receive 3 heatbeats consecutively, notify admin (create a handle interface, provide email or other notification.)


## pull module

Server pull the client healthcheck endpoint.

## flag

-m, --mode: pull or push
-n, --notification: email
-e, --email: email address to receive email notification
-i, --intervel: pull/push intervel
-c, --client: run as client. by default run as server
-e, --endpoint: if pull mode, this is the endpoint of client healthcheck. If push mode, this is the endpoint on server side
-p, --protocol: tcp/udp/icmp/all
