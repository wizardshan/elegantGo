gen:
	go run main.go -name ${name} -layers entMapper -layers repo -layers dom -layers req -layers resp -layers respMapper -layers ctr

layer:
	go run main.go -name ${name} -layers ${layer}

ent:
	go run main.go -name ${name} -layers entMapper -layers dom -layers resp -layers respMapper -layers repo

clear:
	go run main.go -name ${name} -layers entMapper -layers dom -layers resp -layers respMapper -layers repo -layers req -layers ctr -c true

# make gen name=User
# make layer name=User -layer
# make clear name=User