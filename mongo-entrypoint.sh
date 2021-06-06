#!/bin/sh
# mkdir mongolog
# touch /mongolog/mongo.log
# mongod -v --logpath /mongolog/mongo.log

# sleep 60

chmod 600 /usr/mongo-secret
/usr/bin/mongod --keyFile /usr/mongo-secret --bind_ip_all --replSet rs1




