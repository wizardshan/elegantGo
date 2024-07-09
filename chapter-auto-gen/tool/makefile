gen:
	go run main.go -name ${name} -layers ent -layers repo -layers dom -layers req -layers resp -layers ctr

genReq:
	go run main.go -name ${name} -layers req

genEnt:
	go run main.go -name ${name} -layers ent

clearAll:
	go run main.go -name ${name} -layers ent -layers repo -layers dom -layers req -layers resp -layers ctr -c true

# make gen name=User
# make gen name=User -layer
# make clearAll name=User