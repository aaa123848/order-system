mongo --eval "rs.initiate({_id:'rs1', members: [{_id: 0, host: 'order-mongo-1:27017'}]})"

mongo admin --eval "db.createUser({user: 'root', pwd: '1234', roles: ['root']})"